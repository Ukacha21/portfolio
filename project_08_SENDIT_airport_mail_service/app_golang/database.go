package main

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

var db *sql.DB

// Shipment represents a shipment record with timeline
type Shipment struct {
	ID           int64              `json:"id"`
	Tracking     string             `json:"tracking"`
	ItemType     string             `json:"item_type"`
	ItemDesc     string             `json:"item_desc"`
	ItemValue    float64            `json:"item_value"`
	Reason       string             `json:"reason"`
	SenderName   string             `json:"sender_name"`
	SenderEmail  string             `json:"sender_email"`
	SenderPhone  string             `json:"sender_phone"`
	RecvName     string             `json:"recv_name"`
	RecvAddress  string             `json:"recv_address"`
	RecvCity     string             `json:"recv_city"`
	RecvPostcode string             `json:"recv_postcode"`
	RecvCountry  string             `json:"recv_country"`
	Dest         string             `json:"dest"`
	Courier      string             `json:"courier"`
	PkgSize      string             `json:"pkg_size"`
	Total        float64            `json:"total"`
	Insurance    int                `json:"insurance"`
	Status       string             `json:"status"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
	Timeline     []TimelineEvent    `json:"timeline"`
}

type TimelineEvent struct {
	Status string `json:"status"`
	Note   string `json:"note,omitempty"`
	Ts     string `json:"ts"`
}

type Stats struct {
	Total   int     `json:"total"`
	Pending int     `json:"pending"`
	Posted  int     `json:"posted"`
	Mailed  int     `json:"mailed"`
	Failed  int     `json:"failed"`
	Revenue float64 `json:"revenue"`
}

// InitDB creates tables and seeds default data
func InitDB() error {
	dbPath := filepath.Join(".", "sendit.db")

	var err error
	db, err = sql.Open("sqlite", dbPath+"?_pragma=journal_mode(WAL)&_pragma=foreign_keys(ON)")
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	if err = db.Ping(); err != nil {
		return fmt.Errorf("ping db: %w", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS shipments (
		id            INTEGER PRIMARY KEY AUTOINCREMENT,
		tracking      TEXT    UNIQUE NOT NULL,
		item_type     TEXT    NOT NULL,
		item_desc     TEXT    NOT NULL,
		item_value    REAL    DEFAULT 0,
		reason        TEXT    DEFAULT 'confiscated',
		sender_name   TEXT    NOT NULL,
		sender_email  TEXT,
		sender_phone  TEXT,
		recv_name     TEXT    NOT NULL,
		recv_address  TEXT,
		recv_city     TEXT,
		recv_postcode TEXT,
		recv_country  TEXT,
		dest          TEXT,
		courier       TEXT    NOT NULL,
		pkg_size      TEXT    NOT NULL,
		total         REAL    NOT NULL,
		insurance     INTEGER DEFAULT 0,
		status        TEXT    NOT NULL DEFAULT 'pending',
		created_at    TEXT    NOT NULL,
		updated_at    TEXT    NOT NULL
	);
	CREATE TABLE IF NOT EXISTS timeline_events (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		shipment_id INTEGER NOT NULL REFERENCES shipments(id) ON DELETE CASCADE,
		status      TEXT    NOT NULL,
		note        TEXT,
		ts          TEXT    NOT NULL
	);
	CREATE TABLE IF NOT EXISTS settings (
		key   TEXT PRIMARY KEY,
		value TEXT
	);
	CREATE TABLE IF NOT EXISTS mobile_prefill (
		code       TEXT PRIMARY KEY,
		payload    TEXT NOT NULL,
		created_at TEXT NOT NULL,
		expires_at TEXT NOT NULL,
		used       INTEGER DEFAULT 0
	);`

	if _, err = db.Exec(schema); err != nil {
		return fmt.Errorf("create tables: %w", err)
	}

	// Default settings
	for _, kv := range [][2]string{
		{"admin_pin", "1234"},
		{"kiosk_name", "Airport Kiosk"},
		{"currency", "GBP"},
		{"mobile_base_url", ""},
	} {
		db.Exec("INSERT OR IGNORE INTO settings (key, value) VALUES (?, ?)", kv[0], kv[1])
	}

	// Seed if empty
	var count int
	db.QueryRow("SELECT COUNT(*) FROM shipments").Scan(&count)
	if count == 0 {
		seedDemoData()
	}

	log.Printf("✅ Database ready: %s", dbPath)
	return nil
}

func daysAgo(n int) string {
	t := time.Now().UTC().Add(-time.Duration(n) * 24 * time.Hour)
	return t.Format("2006-01-02T15:04:05.000Z")
}

func seedDemoData() {
	type demo struct {
		tracking, itemType, itemDesc, senderName, senderEmail string
		recvName, recvAddress, recvCity, recvPostcode, recvCountry, dest string
		courier, pkgSize, status string
		total float64; insurance, created, updated int
	}
	demos := []demo{
		{"SD-2026-001920","liquid","Chanel No. 5 Perfume 100ml","Maria García","maria@example.com","Maria García","42 Notting Hill Gate","London","W11 3HX","GB","London, GB","dhl","small","mailed",10.90,0,3,1},
		{"SD-2026-001919","sharp","Victorinox Swiss Army Knife","James O'Brien","james@example.com","James O'Brien","12 O'Connell St","Dublin","D01","IE","Dublin, IE","dpd","small","posted",8.70,0,2,1},
		{"SD-2026-001915","tool","Kitchen Scissors","Karan Mehta","karan@example.com","Karan Mehta","MG Road","Mumbai","400001","IN","Mumbai, IN","dhl","small","pending",15.90,0,1,1},
		{"SD-2026-001913","food","Whisky Miniatures x4","Callum MacLeod","cal@example.com","Callum MacLeod","Royal Mile","Edinburgh","EH1 2AB","GB","Edinburgh, GB","rm","envelope","failed",7.00,0,7,5},
	}
	for _, d := range demos {
		ca, ua := daysAgo(d.created), daysAgo(d.updated)
		res, err := db.Exec(`
			INSERT INTO shipments (tracking,item_type,item_desc,sender_name,sender_email,recv_name,recv_address,recv_city,recv_postcode,recv_country,dest,courier,pkg_size,total,insurance,status,created_at,updated_at)
			VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
			d.tracking,d.itemType,d.itemDesc,d.senderName,d.senderEmail,
			d.recvName,d.recvAddress,d.recvCity,d.recvPostcode,d.recvCountry,
			d.dest,d.courier,d.pkgSize,d.total,d.insurance,d.status,ca,ua)
		if err != nil { continue }
		sid, _ := res.LastInsertId()
		switch d.status {
		case "mailed":
			db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "pending", ca)
			db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "posted", daysAgo(d.created-1))
			db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "mailed", ua)
		case "posted":
			db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "pending", ca)
			db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "posted", ua)
		case "failed":
			db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "pending", ca)
			db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "posted", daysAgo(d.created-1))
			db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "failed", ua)
		default:
			db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "pending", ca)
		}
	}
	log.Println("✅ Demo data seeded")
}

func getTimeline(id int64) []TimelineEvent {
	rows, err := db.Query("SELECT status, COALESCE(note,'') as note, ts FROM timeline_events WHERE shipment_id=? ORDER BY ts ASC", id)
	if err != nil { return nil }
	defer rows.Close()
	var tl []TimelineEvent
	for rows.Next() {
		var e TimelineEvent
		rows.Scan(&e.Status, &e.Note, &e.Ts)
		tl = append(tl, e)
	}
	return tl
}

func scanShipment(row interface{ Scan(...interface{}) error }) (*Shipment, error) {
	var s Shipment
	err := row.Scan(&s.ID, &s.Tracking, &s.ItemType, &s.ItemDesc, &s.ItemValue, &s.Reason,
		&s.SenderName, &s.SenderEmail, &s.SenderPhone,
		&s.RecvName, &s.RecvAddress, &s.RecvCity, &s.RecvPostcode, &s.RecvCountry, &s.Dest,
		&s.Courier, &s.PkgSize, &s.Total, &s.Insurance, &s.Status, &s.CreatedAt, &s.UpdatedAt)
	if err != nil { return nil, err }
	s.Timeline = getTimeline(s.ID)
	return &s, nil
}

const selectShipment = `SELECT id,tracking,item_type,item_desc,item_value,reason,
	sender_name,sender_email,sender_phone,recv_name,recv_address,recv_city,recv_postcode,recv_country,dest,
	courier,pkg_size,total,insurance,status,created_at,updated_at FROM shipments`

func GetAllShipments() ([]Shipment, error) {
	rows, err := db.Query(selectShipment + " ORDER BY created_at DESC")
	if err != nil { return nil, err }
	defer rows.Close()
	var result []Shipment
	for rows.Next() {
		s, err := scanShipment(rows)
		if err != nil { continue }
		result = append(result, *s)
	}
	if result == nil { result = []Shipment{} }
	return result, nil
}

func GetShipmentByID(id string) (*Shipment, error) {
	row := db.QueryRow(selectShipment+" WHERE id=?", id)
	return scanShipment(row)
}

func GetShipmentByTracking(tracking string) (*Shipment, error) {
	row := db.QueryRow(selectShipment+" WHERE tracking=?", strings.ToUpper(tracking))
	return scanShipment(row)
}

func GetStats() (*Stats, error) {
	var s Stats
	err := db.QueryRow(`SELECT COUNT(*),
		SUM(CASE WHEN status='pending' THEN 1 ELSE 0 END),
		SUM(CASE WHEN status='posted'  THEN 1 ELSE 0 END),
		SUM(CASE WHEN status='mailed'  THEN 1 ELSE 0 END),
		SUM(CASE WHEN status='failed'  THEN 1 ELSE 0 END),
		COALESCE(SUM(total),0) FROM shipments`).
		Scan(&s.Total, &s.Pending, &s.Posted, &s.Mailed, &s.Failed, &s.Revenue)
	return &s, err
}

func generateTracking() string {
	year := time.Now().UTC().Year()
	var mx sql.NullInt64
	db.QueryRow(fmt.Sprintf("SELECT MAX(CAST(SUBSTR(tracking,9) AS INTEGER)) FROM shipments WHERE tracking LIKE 'SD-%d-%%'", year)).Scan(&mx)
	next := int64(1919)
	if mx.Valid { next = mx.Int64 }
	return fmt.Sprintf("SD-%d-%06d", year, next+1)
}

func CreateShipment(body CreateShipmentReq) (*Shipment, error) {
	now := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	tracking := generateTracking()
	ins := 0; if body.Insurance { ins = 1 }
	res, err := db.Exec(`
		INSERT INTO shipments (tracking,item_type,item_desc,item_value,reason,
		sender_name,sender_email,sender_phone,recv_name,recv_address,recv_city,recv_postcode,recv_country,dest,
		courier,pkg_size,total,insurance,status,created_at,updated_at)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,'pending',?,?)`,
		tracking, body.ItemType, body.ItemDesc, body.ItemValue, body.Reason,
		body.SenderName, body.SenderEmail, body.SenderPhone,
		body.RecvName, body.RecvAddress, body.RecvCity, body.RecvPostcode, body.RecvCountry, body.Dest,
		body.Courier, body.PkgSize, body.Total, ins, now, now)
	if err != nil { return nil, err }
	sid, _ := res.LastInsertId()
	db.Exec("INSERT INTO timeline_events (shipment_id,status,ts) VALUES (?,?,?)", sid, "pending", now)
	return GetShipmentByID(fmt.Sprintf("%d", sid))
}

func UpdateShipmentStatus(id, status, note string) (*Shipment, error) {
	now := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	db.Exec("UPDATE shipments SET status=?, updated_at=? WHERE id=?", status, now, id)
	db.Exec("INSERT INTO timeline_events (shipment_id,status,note,ts) VALUES (?,?,?,?)", id, status, note, now)
	return GetShipmentByID(id)
}

func DeleteShipment(id string) error {
	_, err := db.Exec("DELETE FROM shipments WHERE id=?", id)
	return err
}

func GetSetting(key string) (string, error) {
	var v string
	err := db.QueryRow("SELECT value FROM settings WHERE key=?", key).Scan(&v)
	return v, err
}

func SetSetting(key, value string) {
	db.Exec("INSERT OR REPLACE INTO settings (key, value) VALUES (?,?)", key, value)
}

func GetAllSettings() (map[string]string, error) {
	rows, err := db.Query("SELECT key, value FROM settings")
	if err != nil { return nil, err }
	defer rows.Close()
	result := make(map[string]string)
	for rows.Next() {
		var k, v string
		rows.Scan(&k, &v)
		result[k] = v
	}
	return result, nil
}

// ─── Mobile Pre-fill ─────────────────────────────────────────

const PREFILL_TTL = 3 * time.Hour

func CreatePrefill(code, payload string) (time.Time, error) {
	now := time.Now().UTC()
	expires := now.Add(PREFILL_TTL)
	_, err := db.Exec(
		`INSERT OR REPLACE INTO mobile_prefill (code, payload, created_at, expires_at, used) VALUES (?,?,?,?,0)`,
		code, payload,
		now.Format(time.RFC3339),
		expires.Format(time.RFC3339),
	)
	return expires, err
}

type PrefillResult struct {
	Payload   string
	ExpiresAt time.Time
	Used      bool
}

func RedeemPrefill(code string) (*PrefillResult, error) {
	var payload, expiresStr string
	var used int
	err := db.QueryRow(
		`SELECT payload, expires_at, used FROM mobile_prefill WHERE code=?`, code,
	).Scan(&payload, &expiresStr, &used)
	if err != nil {
		return nil, err // not found
	}
	expiresAt, _ := time.Parse(time.RFC3339, expiresStr)
	if time.Now().UTC().After(expiresAt) {
		db.Exec(`DELETE FROM mobile_prefill WHERE code=?`, code)
		return nil, fmt.Errorf("expired")
	}
	if used != 0 {
		return nil, fmt.Errorf("already_used")
	}
	// Mark as used
	db.Exec(`UPDATE mobile_prefill SET used=1 WHERE code=?`, code)
	return &PrefillResult{Payload: payload, ExpiresAt: expiresAt, Used: false}, nil
}

func CleanupExpiredPrefills() {
	db.Exec(`DELETE FROM mobile_prefill WHERE expires_at < ?`, time.Now().UTC().Format(time.RFC3339))
}
