package main

/*
SENDIT — Airport Mail Kiosk
Go + Gin + SQLite Backend Server v1.0

Run: go run .
Serves: http://localhost:3000

100% API-compatible with the Node.js and Python versions.
The shared frontend (index.html, mobile.html) works with this backend.
*/

import (
	"encoding/json"
	"fmt"
	"log"

	// "net/http"
	// "os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const PORT = 3000

// ─── In-memory QR sessions ──────────────────────────────────────
type Session struct {
	Data      map[string]interface{}
	CreatedAt time.Time
	ExpiresAt time.Time
}

var (
	sessions   = make(map[string]*Session)
	sessionsMu sync.RWMutex
)

const CODE_CHARS = "ACDEFGHJKMNPQRTUVWXY3469"

func generateCode() string {
	code := make([]byte, 6)
	for i := range code {
		code[i] = CODE_CHARS[time.Now().UnixNano()%int64(len(CODE_CHARS))]
		time.Sleep(time.Nanosecond)
	}
	return string(code)
}

func cleanupSession(code string) {
	time.Sleep(10 * time.Minute)
	sessionsMu.Lock()
	delete(sessions, code)
	sessionsMu.Unlock()
}

// ─── Main ────────────────────────────────────────────────────────
func main() {

	// ... inside main()
	//r := gin.Default()
	//r.Use(cors.Default()) // Allows your GitHub Pages site to talk to this API

	// Init DB
	if err := InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//r.Use(cors.Default()) // Allows your GitHub Pages site to talk to this API

	// CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Static frontend files
	r.StaticFile("/", "./index.html")
	r.StaticFile("/mobile", "./mobile.html")
	r.StaticFile("/style.css", "./style.css")
	r.StaticFile("/app.js", "./app.js")

	// ── API Routes ──────────────────────────────────────────────
	api := r.Group("/api")

	// Stats
	api.GET("/stats", handleGetStats)

	// Shipments
	api.GET("/shipments", handleGetShipments)
	api.POST("/shipments", handleCreateShipment)
	api.GET("/shipments/track/:code", handleTrackShipment)
	api.GET("/shipments/:id", handleGetShipment)
	api.PUT("/shipments/:id/status", handleUpdateStatus)
	api.DELETE("/shipments/:id", handleDeleteShipment)

	// Auth
	api.POST("/auth/verify-pin", handleVerifyPin)

	// Settings
	api.GET("/settings", handleGetSettings)
	api.PUT("/settings/pin", handleChangePin)
	api.PUT("/settings/:key", handleUpdateSetting)

	// QR Sessions
	api.POST("/session/create", handleCreateSession)
	api.POST("/session/submit/:code", handleSubmitSession)
	api.GET("/session/poll/:code", handlePollSession)

	// Mobile Pre-fill (fill from phone before reaching kiosk)
	api.POST("/prefill/create", handleCreatePrefill)
	api.POST("/prefill/redeem/:code", handleRedeemPrefill)
	api.GET("/prefill/check/:code", handleCheckPrefill)

	// Start cleanup goroutine for expired prefills
	go func() {
		for {
			time.Sleep(30 * time.Minute)
			CleanupExpiredPrefills()
		}
	}()

	// Start
	fmt.Printf("\n🚀 SENDIT Kiosk Server (Go) running at http://localhost:%d\n", PORT)
	fmt.Printf("📦 Database: sendit.db\n")
	fmt.Printf("📱 Mobile form: http://localhost:%d/mobile\n\n", PORT)

	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", PORT)); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// ─── Handlers ────────────────────────────────────────────────────

func handleGetStats(c *gin.Context) {
	stats, err := GetStats()
	if err != nil {
		c.JSON(500, gin.H{"ok": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"ok": true, "data": stats})
}

func handleGetShipments(c *gin.Context) {
	shipments, err := GetAllShipments()
	if err != nil {
		c.JSON(500, gin.H{"ok": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"ok": true, "data": shipments})
}

func handleGetShipment(c *gin.Context) {
	id := c.Param("id")
	s, err := GetShipmentByID(id)
	if err != nil || s == nil {
		c.JSON(404, gin.H{"ok": false, "error": "Not found"})
		return
	}
	c.JSON(200, gin.H{"ok": true, "data": s})
}

func handleTrackShipment(c *gin.Context) {
	code := c.Param("code")
	s, err := GetShipmentByTracking(code)
	if err != nil || s == nil {
		c.JSON(404, gin.H{"ok": false, "error": "Not found"})
		return
	}
	c.JSON(200, gin.H{"ok": true, "data": s})
}

func handleCreateShipment(c *gin.Context) {
	var body CreateShipmentReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"ok": false, "error": err.Error()})
		return
	}
	created, err := CreateShipment(body)
	if err != nil {
		c.JSON(500, gin.H{"ok": false, "error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"ok": true, "data": created})
}

func handleUpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var body UpdateStatusReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"ok": false, "error": err.Error()})
		return
	}
	valid := map[string]bool{"pending": true, "posted": true, "mailed": true, "failed": true}
	if !valid[body.Status] {
		c.JSON(400, gin.H{"ok": false, "error": "Invalid status"})
		return
	}
	updated, err := UpdateShipmentStatus(id, body.Status, body.Note)
	if err != nil {
		c.JSON(500, gin.H{"ok": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"ok": true, "data": updated})
}

func handleDeleteShipment(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteShipment(id); err != nil {
		c.JSON(500, gin.H{"ok": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"ok": true})
}

func handleVerifyPin(c *gin.Context) {
	var body VerifyPinReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"ok": false, "error": err.Error()})
		return
	}
	stored, _ := GetSetting("admin_pin")
	if stored == "" {
		stored = "1234"
	}
	if body.Pin == stored {
		c.JSON(200, gin.H{"ok": true})
	} else {
		c.JSON(401, gin.H{"ok": false, "error": "Invalid PIN"})
	}
}

func handleGetSettings(c *gin.Context) {
	settings, err := GetAllSettings()
	if err != nil {
		c.JSON(500, gin.H{"ok": false, "error": err.Error()})
		return
	}
	delete(settings, "admin_pin") // never expose PIN
	c.JSON(200, gin.H{"ok": true, "data": settings})
}

func handleChangePin(c *gin.Context) {
	var body ChangePinReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"ok": false, "error": err.Error()})
		return
	}
	stored, _ := GetSetting("admin_pin")
	if stored == "" {
		stored = "1234"
	}
	if body.CurrentPin != stored {
		c.JSON(401, gin.H{"ok": false, "error": "Wrong current PIN"})
		return
	}
	if len(body.NewPin) < 4 || len(body.NewPin) > 8 {
		c.JSON(400, gin.H{"ok": false, "error": "PIN must be 4–8 digits"})
		return
	}
	SetSetting("admin_pin", body.NewPin)
	c.JSON(200, gin.H{"ok": true})
}

func handleUpdateSetting(c *gin.Context) {
	key := c.Param("key")
	allowed := map[string]bool{"currency": true, "kiosk_name": true, "mobile_base_url": true}
	if !allowed[key] {
		c.JSON(400, gin.H{"ok": false, "error": "Setting not allowed"})
		return
	}
	var body UpdateSettingReq
	c.ShouldBindJSON(&body)
	SetSetting(key, body.Value)
	c.JSON(200, gin.H{"ok": true})
}

func handleCreateSession(c *gin.Context) {
	code := generateCode()
	sessionsMu.Lock()
	attempts := 0
	for _, ok := sessions[code]; ok && attempts < 20; attempts++ {
		code = generateCode()
		_, ok = sessions[code]
	}
	now := time.Now()
	expires := now.Add(10 * time.Minute)
	sessions[code] = &Session{Data: nil, CreatedAt: now, ExpiresAt: expires}
	sessionsMu.Unlock()

	go cleanupSession(code)

	mobileBase, _ := GetSetting("mobile_base_url")
	host := mobileBase
	if host == "" {
		scheme := "http"
		hostHeader := c.Request.Host
		if hostHeader == "" {
			hostHeader = fmt.Sprintf("localhost:%d", PORT)
		}
		host = fmt.Sprintf("%s://%s", scheme, hostHeader)
	}
	mobileURL := fmt.Sprintf("%s/mobile?code=%s", host, code)

	c.JSON(200, gin.H{
		"ok":         true,
		"code":       code,
		"mobile_url": mobileURL,
		"expires_at": expires.UTC().Format(time.RFC3339),
	})
}

func handleSubmitSession(c *gin.Context) {
	code := c.Param("code")
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"ok": false, "error": err.Error()})
		return
	}
	sessionsMu.Lock()
	defer sessionsMu.Unlock()
	sess, ok := sessions[code]
	if !ok {
		c.JSON(404, gin.H{"ok": false, "error": "Code not found or expired"})
		return
	}
	if sess.Data != nil {
		c.JSON(409, gin.H{"ok": false, "error": "Code already used"})
		return
	}
	if time.Now().After(sess.ExpiresAt) {
		delete(sessions, code)
		c.JSON(410, gin.H{"ok": false, "error": "Code expired"})
		return
	}
	sess.Data = body
	c.JSON(200, gin.H{"ok": true, "message": "Data received — check the kiosk screen"})
}

func handlePollSession(c *gin.Context) {
	code := c.Param("code")
	sessionsMu.Lock()
	defer sessionsMu.Unlock()
	sess, ok := sessions[code]
	if !ok {
		c.JSON(200, gin.H{"ok": false, "status": "expired"})
		return
	}
	if time.Now().After(sess.ExpiresAt) {
		delete(sessions, code)
		c.JSON(200, gin.H{"ok": false, "status": "expired"})
		return
	}
	if sess.Data != nil {
		data := sess.Data
		delete(sessions, code)
		c.JSON(200, gin.H{"ok": true, "status": "received", "data": data})
		return
	}
	c.JSON(200, gin.H{"ok": false, "status": "waiting"})
}

// ─── Mobile Pre-fill Handlers ───────────────────────────────────────

func handleCreatePrefill(c *gin.Context) {
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"ok": false, "error": err.Error()})
		return
	}

	// Generate a unique 6-char code
	code := generateCode()
	for {
		// Ensure uniqueness (simple retry)
		var cnt int
		db.QueryRow(`SELECT COUNT(*) FROM mobile_prefill WHERE code=? AND used=0 AND expires_at > ?`,
			code, time.Now().UTC().Format(time.RFC3339)).Scan(&cnt)
		if cnt == 0 {
			break
		}
		code = generateCode()
	}

	payloadBytes, err := json.Marshal(body)
	if err != nil {
		c.JSON(500, gin.H{"ok": false, "error": "failed to encode payload"})
		return
	}

	expiresAt, err := CreatePrefill(code, string(payloadBytes))
	if err != nil {
		c.JSON(500, gin.H{"ok": false, "error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"ok":         true,
		"code":       code,
		"expires_at": expiresAt.UTC().Format(time.RFC3339),
		"ttl_hours":  3,
	})
}

func handleRedeemPrefill(c *gin.Context) {
	code := c.Param("code")
	result, err := RedeemPrefill(code)
	if err != nil {
		switch err.Error() {
		case "expired":
			c.JSON(410, gin.H{"ok": false, "error": "Code expired — prefill codes are valid for 3 hours"})
		case "already_used":
			c.JSON(409, gin.H{"ok": false, "error": "This code has already been used at the kiosk"})
		default:
			c.JSON(404, gin.H{"ok": false, "error": "Code not found"})
		}
		return
	}

	// Decode payload back to map
	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(result.Payload), &payload); err != nil {
		c.JSON(500, gin.H{"ok": false, "error": "corrupted payload"})
		return
	}

	c.JSON(200, gin.H{
		"ok":   true,
		"data": payload,
	})
}

func handleCheckPrefill(c *gin.Context) {
	code := c.Param("code")
	var expiresStr string
	var used int
	err := db.QueryRow(
		`SELECT expires_at, used FROM mobile_prefill WHERE code=?`, code,
	).Scan(&expiresStr, &used)
	if err != nil {
		c.JSON(200, gin.H{"ok": false, "status": "not_found"})
		return
	}
	expiresAt, _ := time.Parse(time.RFC3339, expiresStr)
	if time.Now().UTC().After(expiresAt) {
		c.JSON(200, gin.H{"ok": false, "status": "expired"})
		return
	}
	if used != 0 {
		c.JSON(200, gin.H{"ok": false, "status": "used"})
		return
	}
	c.JSON(200, gin.H{
		"ok":         true,
		"status":     "ready",
		"expires_at": expiresAt.UTC().Format(time.RFC3339),
	})
}
