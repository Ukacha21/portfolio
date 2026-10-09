/* ============================================================
   SENDIT — Airport Mail Kiosk
   Express + SQLite Backend Server v2.0
   Run: node server.js
   Serves: http://localhost:3000
   ============================================================ */

'use strict';

const express  = require('express');
const Database = require('better-sqlite3');
const cors     = require('cors');
const path     = require('path');
const crypto   = require('crypto');

const app  = express();
const PORT = process.env.PORT || 3000;

// ─── Database Setup ──────────────────────────────────────────
const DB_PATH = path.join(__dirname, 'sendit.db');
const db      = new Database(DB_PATH);

// Enable WAL for performance and concurrent reads
db.pragma('journal_mode = WAL');
db.pragma('foreign_keys = ON');

// Create tables
db.exec(`
  CREATE TABLE IF NOT EXISTS shipments (
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    tracking     TEXT    UNIQUE NOT NULL,
    item_type    TEXT    NOT NULL,
    item_desc    TEXT    NOT NULL,
    item_value   REAL    DEFAULT 0,
    reason       TEXT    DEFAULT 'confiscated',
    sender_name  TEXT    NOT NULL,
    sender_email TEXT,
    sender_phone TEXT,
    recv_name    TEXT    NOT NULL,
    recv_address TEXT,
    recv_city    TEXT,
    recv_postcode TEXT,
    recv_country TEXT,
    dest         TEXT,
    courier      TEXT    NOT NULL,
    pkg_size     TEXT    NOT NULL,
    total        REAL    NOT NULL,
    insurance    INTEGER DEFAULT 0,
    status       TEXT    NOT NULL DEFAULT 'pending',
    created_at   TEXT    NOT NULL,
    updated_at   TEXT    NOT NULL
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
`);

// Seed default settings
const insertSetting = db.prepare(`INSERT OR IGNORE INTO settings (key, value) VALUES (?, ?)`);
insertSetting.run('admin_pin',       '1234');
insertSetting.run('kiosk_name',      'Airport Kiosk');
insertSetting.run('currency',        'GBP');
insertSetting.run('mobile_base_url', '');  // if set, used for QR code URL instead of local IP

// Seed demo data if empty
const shipmentCount = db.prepare('SELECT COUNT(*) as cnt FROM shipments').get().cnt;
if (shipmentCount === 0) {
  const seedShipments = [
    { tracking:'SD-2026-001920', item_type:'liquid',      item_desc:'Chanel No. 5 Perfume 100ml',   sender_name:'Maria García',    sender_email:'maria@example.com', recv_name:'Maria García',    recv_address:'42 Notting Hill Gate', recv_city:'London',    recv_postcode:'W11 3HX', recv_country:'GB', dest:'London, GB',     courier:'dhl', pkg_size:'small',    total:10.90, insurance:0, status:'mailed',  days_ago_created:3, days_ago_updated:1 },
    { tracking:'SD-2026-001919', item_type:'sharp',       item_desc:"Victorinox Swiss Army Knife",   sender_name:"James O'Brien",   sender_email:'james@example.com', recv_name:"James O'Brien",   recv_address:'12 O\'Connell St',     recv_city:'Dublin',    recv_postcode:'D01',     recv_country:'IE', dest:'Dublin, IE',     courier:'dpd', pkg_size:'small',    total:8.70,  insurance:0, status:'posted',  days_ago_created:2, days_ago_updated:1 },
    { tracking:'SD-2026-001918', item_type:'electronics', item_desc:'Anker PowerBank 26800mAh',      sender_name:'Yuki Tanaka',     sender_email:'yuki@example.com',  recv_name:'Yuki Tanaka',     recv_address:'3-1 Shibuya',         recv_city:'Tokyo',     recv_postcode:'150-0001',recv_country:'JP', dest:'Tokyo, JP',      courier:'dhl', pkg_size:'small',    total:16.40, insurance:1, status:'mailed',  days_ago_created:4, days_ago_updated:2 },
    { tracking:'SD-2026-001917', item_type:'lighter',     item_desc:'Zippo Lighter Silver',          sender_name:'Ahmed Al-Farsi',  sender_email:'ahmed@example.com', recv_name:'Ahmed Al-Farsi',  recv_address:'Al Barsha St',        recv_city:'Dubai',     recv_postcode:'00000',   recv_country:'AE', dest:'Dubai, AE',      courier:'rm',  pkg_size:'envelope', total:7.00,  insurance:0, status:'mailed',  days_ago_created:5, days_ago_updated:3 },
    { tracking:'SD-2026-001916', item_type:'cosmetics',   item_desc:'Jo Malone Scented Candle',      sender_name:'Sophie Laurent',  sender_email:'sophie@example.com',recv_name:'Sophie Laurent',  recv_address:'15 Rue de Rivoli',    recv_city:'Paris',     recv_postcode:'75001',   recv_country:'FR', dest:'Paris, FR',      courier:'dpd', pkg_size:'small',    total:8.70,  insurance:0, status:'mailed',  days_ago_created:6, days_ago_updated:4 },
    { tracking:'SD-2026-001915', item_type:'tool',        item_desc:'Kitchen Scissors',              sender_name:'Karan Mehta',     sender_email:'karan@example.com', recv_name:'Karan Mehta',     recv_address:'MG Road',             recv_city:'Mumbai',    recv_postcode:'400001',  recv_country:'IN', dest:'Mumbai, IN',     courier:'dhl', pkg_size:'small',    total:15.90, insurance:0, status:'pending', days_ago_created:1, days_ago_updated:1 },
    { tracking:'SD-2026-001914', item_type:'liquid',      item_desc:'Gel Toothpaste 200ml',          sender_name:'Lena Mueller',    sender_email:'lena@example.com',  recv_name:'Lena Mueller',    recv_address:'Unter den Linden 10', recv_city:'Berlin',    recv_postcode:'10117',   recv_country:'DE', dest:'Berlin, DE',     courier:'rm',  pkg_size:'envelope', total:6.00,  insurance:0, status:'pending', days_ago_created:1, days_ago_updated:1 },
    { tracking:'SD-2026-001913', item_type:'food',        item_desc:'Whisky Miniatures x4',          sender_name:'Callum MacLeod',  sender_email:'cal@example.com',   recv_name:'Callum MacLeod',  recv_address:'Royal Mile',          recv_city:'Edinburgh', recv_postcode:'EH1 2AB', recv_country:'GB', dest:'Edinburgh, GB',  courier:'rm',  pkg_size:'envelope', total:7.00,  insurance:0, status:'failed',  days_ago_created:7, days_ago_updated:5 },
  ];

  const insertShipment = db.prepare(`
    INSERT INTO shipments (tracking, item_type, item_desc, sender_name, sender_email, recv_name, recv_address, recv_city, recv_postcode, recv_country, dest, courier, pkg_size, total, insurance, status, created_at, updated_at)
    VALUES (@tracking, @item_type, @item_desc, @sender_name, @sender_email, @recv_name, @recv_address, @recv_city, @recv_postcode, @recv_country, @dest, @courier, @pkg_size, @total, @insurance, @status, @created_at, @updated_at)
  `);
  const insertEvent = db.prepare(`INSERT INTO timeline_events (shipment_id, status, ts) VALUES (?, ?, ?)`);

  const now = Date.now();
  const daysAgo = (n) => new Date(now - n * 86400000).toISOString();

  const seedTx = db.transaction(() => {
    for (const s of seedShipments) {
      const created_at = daysAgo(s.days_ago_created);
      const updated_at = daysAgo(s.days_ago_updated);
      const info = insertShipment.run({ ...s, created_at, updated_at });
      const sid  = info.lastInsertRowid;

      // Build timeline
      if (['posted','mailed','failed'].includes(s.status)) {
        insertEvent.run(sid, 'pending', daysAgo(s.days_ago_created));
        if (s.status === 'mailed') {
          insertEvent.run(sid, 'posted', daysAgo(s.days_ago_created - 1));
          insertEvent.run(sid, 'mailed', daysAgo(s.days_ago_updated));
        } else if (s.status === 'posted') {
          insertEvent.run(sid, 'posted', daysAgo(s.days_ago_updated));
        } else if (s.status === 'failed') {
          insertEvent.run(sid, 'posted', daysAgo(s.days_ago_created - 1));
          insertEvent.run(sid, 'failed', daysAgo(s.days_ago_updated));
        }
      } else {
        insertEvent.run(sid, 'pending', created_at);
      }
    }
  });
  seedTx();
  console.log('✅ Demo data seeded into sendit.db');
}

// ─── Helpers ─────────────────────────────────────────────────
function getShipmentWithTimeline(id) {
  const s = db.prepare('SELECT * FROM shipments WHERE id = ?').get(id);
  if (!s) return null;
  s.timeline = db.prepare('SELECT * FROM timeline_events WHERE shipment_id = ? ORDER BY ts ASC').all(id);
  return s;
}

function generateTracking() {
  const year = new Date().getFullYear();
  const num  = db.prepare("SELECT MAX(CAST(SUBSTR(tracking, 9) AS INTEGER)) as max FROM shipments WHERE tracking LIKE 'SD-' || ? || '-%'").get(String(year));
  const next = (num?.max || 1919) + 1;
  return `SD-${year}-${String(next).padStart(6, '0')}`;
}

// ─── Middleware ──────────────────────────────────────────────
app.use(cors());
app.use(express.json());
app.use(express.static(path.join(__dirname)));  // serve frontend files

// ─── API Routes ──────────────────────────────────────────────

// GET all shipments (with timeline)
app.get('/api/shipments', (req, res) => {
  try {
    const shipments = db.prepare('SELECT * FROM shipments ORDER BY created_at DESC').all();
    const withTimeline = shipments.map(s => {
      s.timeline = db.prepare('SELECT * FROM timeline_events WHERE shipment_id = ? ORDER BY ts ASC').all(s.id);
      return s;
    });
    res.json({ ok: true, data: withTimeline });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// GET stats summary for dashboard
app.get('/api/stats', (req, res) => {
  try {
    const counts = db.prepare(`
      SELECT 
        COUNT(*) as total,
        SUM(CASE WHEN status='pending' THEN 1 ELSE 0 END) as pending,
        SUM(CASE WHEN status='posted'  THEN 1 ELSE 0 END) as posted,
        SUM(CASE WHEN status='mailed'  THEN 1 ELSE 0 END) as mailed,
        SUM(CASE WHEN status='failed'  THEN 1 ELSE 0 END) as failed,
        SUM(total) as revenue
      FROM shipments
    `).get();
    res.json({ ok: true, data: counts });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// GET single shipment by tracking code
app.get('/api/shipments/track/:code', (req, res) => {
  try {
    const s = db.prepare('SELECT * FROM shipments WHERE tracking = ?').get(req.params.code.toUpperCase());
    if (!s) return res.status(404).json({ ok: false, error: 'Not found' });
    s.timeline = db.prepare('SELECT * FROM timeline_events WHERE shipment_id = ? ORDER BY ts ASC').all(s.id);
    res.json({ ok: true, data: s });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// GET single shipment by id
app.get('/api/shipments/:id', (req, res) => {
  try {
    const s = getShipmentWithTimeline(parseInt(req.params.id));
    if (!s) return res.status(404).json({ ok: false, error: 'Not found' });
    res.json({ ok: true, data: s });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// POST create new shipment
app.post('/api/shipments', (req, res) => {
  try {
    const b   = req.body;
    const now = new Date().toISOString();
    const tracking = generateTracking();

    const insert = db.prepare(`
      INSERT INTO shipments (
        tracking, item_type, item_desc, item_value, reason,
        sender_name, sender_email, sender_phone,
        recv_name, recv_address, recv_city, recv_postcode, recv_country, dest,
        courier, pkg_size, total, insurance,
        status, created_at, updated_at
      ) VALUES (
        @tracking, @item_type, @item_desc, @item_value, @reason,
        @sender_name, @sender_email, @sender_phone,
        @recv_name, @recv_address, @recv_city, @recv_postcode, @recv_country, @dest,
        @courier, @pkg_size, @total, @insurance,
        'pending', @now, @now
      )
    `);

    const info = insert.run({
      tracking,
      item_type:    b.item_type    || 'other',
      item_desc:    b.item_desc    || '',
      item_value:   b.item_value   || 0,
      reason:       b.reason       || 'confiscated',
      sender_name:  b.sender_name  || '',
      sender_email: b.sender_email || '',
      sender_phone: b.sender_phone || '',
      recv_name:    b.recv_name    || '',
      recv_address: b.recv_address || '',
      recv_city:    b.recv_city    || '',
      recv_postcode:b.recv_postcode|| '',
      recv_country: b.recv_country || '',
      dest:         b.dest         || '',
      courier:      b.courier      || 'dhl',
      pkg_size:     b.pkg_size     || 'small',
      total:        b.total        || 0,
      insurance:    b.insurance    ? 1 : 0,
      now,
    });

    // Insert initial timeline event
    db.prepare('INSERT INTO timeline_events (shipment_id, status, ts) VALUES (?, ?, ?)').run(info.lastInsertRowid, 'pending', now);

    const created = getShipmentWithTimeline(info.lastInsertRowid);
    res.status(201).json({ ok: true, data: created });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// PUT update shipment status
app.put('/api/shipments/:id/status', (req, res) => {
  try {
    const id     = parseInt(req.params.id);
    const status = req.body.status;
    const note   = req.body.note || null;
    const valid  = ['pending', 'posted', 'mailed', 'failed'];
    if (!valid.includes(status)) return res.status(400).json({ ok: false, error: 'Invalid status' });

    const now = new Date().toISOString();
    db.prepare('UPDATE shipments SET status = ?, updated_at = ? WHERE id = ?').run(status, now, id);
    db.prepare('INSERT INTO timeline_events (shipment_id, status, note, ts) VALUES (?, ?, ?, ?)').run(id, status, note, now);

    const updated = getShipmentWithTimeline(id);
    res.json({ ok: true, data: updated });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// DELETE a shipment
app.delete('/api/shipments/:id', (req, res) => {
  try {
    const id = parseInt(req.params.id);
    db.prepare('DELETE FROM shipments WHERE id = ?').run(id);
    res.json({ ok: true });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// POST verify admin PIN
app.post('/api/auth/verify-pin', (req, res) => {
  try {
    const { pin } = req.body;
    const stored  = db.prepare("SELECT value FROM settings WHERE key = 'admin_pin'").get()?.value || '1234';
    if (pin === stored) {
      res.json({ ok: true });
    } else {
      res.status(401).json({ ok: false, error: 'Invalid PIN' });
    }
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// GET/PUT settings
app.get('/api/settings', (req, res) => {
  try {
    const rows = db.prepare('SELECT key, value FROM settings').all();
    const settings = {};
    rows.forEach(r => { settings[r.key] = r.value; });
    // Never expose the PIN to the frontend
    delete settings.admin_pin;
    res.json({ ok: true, data: settings });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

app.put('/api/settings/pin', (req, res) => {
  try {
    const { current_pin, new_pin } = req.body;
    const stored = db.prepare("SELECT value FROM settings WHERE key = 'admin_pin'").get()?.value || '1234';
    if (current_pin !== stored) return res.status(401).json({ ok: false, error: 'Wrong current PIN' });
    if (!/^\d{4,8}$/.test(new_pin)) return res.status(400).json({ ok: false, error: 'PIN must be 4–8 digits' });
    db.prepare("UPDATE settings SET value = ? WHERE key = 'admin_pin'").run(new_pin);
    res.json({ ok: true });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// ─── Mobile QR Form route ─────────────────────────────────────
app.get('/mobile', (req, res) => {
  res.sendFile(path.join(__dirname, 'mobile.html'));
});

// ─── One-Time Code Session System ────────────────────────────
// Flow:
//  1. Kiosk calls POST /api/session/create  → gets { code, expires_at }
//  2. Kiosk shows code + QR on screen, starts polling GET /api/session/poll/:code
//  3. Passenger opens URL (any network), enters code, fills form
//  4. Passenger calls POST /api/session/submit/:code with their data
//  5. Kiosk poll returns { ok: true, data: {...} } → auto-fills form

const sessions = new Map(); // code => { data: null|{...}, created_at, expires_at }

function generateCode() {
  // 6 uppercase alphanumeric, avoiding ambiguous chars (0/O, 1/I/L)
  const chars = 'ACDEFGHJKMNPQRTUVWXY3469';
  let code = '';
  for (let i = 0; i < 6; i++) {
    code += chars[Math.floor(Math.random() * chars.length)];
  }
  return code;
}

// POST — kiosk creates a session, gets back the one-time code
app.post('/api/session/create', (req, res) => {
  try {
    // Generate a unique code
    let code;
    let attempts = 0;
    do {
      code = generateCode();
      attempts++;
    } while (sessions.has(code) && attempts < 20);

    const now       = Date.now();
    const expires   = now + 10 * 60 * 1000; // 10 minutes
    sessions.set(code, { data: null, created_at: now, expires_at: expires });

    // Auto-expire
    setTimeout(() => sessions.delete(code), 10 * 60 * 1000);

    // Determine the mobile URL base
    const mobileBase = db.prepare("SELECT value FROM settings WHERE key = 'mobile_base_url'").get()?.value || '';
    const host = mobileBase.trim() || `http://localhost:${PORT}`;
    const mobileUrl = `${host}/mobile?code=${code}`;

    res.json({ ok: true, code, mobile_url: mobileUrl, expires_at: new Date(expires).toISOString() });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// POST — passenger submits their form data using the code
app.post('/api/session/submit/:code', (req, res) => {
  try {
    const code = req.params.code.toUpperCase();
    const sess = sessions.get(code);
    if (!sess) return res.status(404).json({ ok: false, error: 'Code not found or expired' });
    if (sess.data) return res.status(409).json({ ok: false, error: 'Code already used' });
    if (Date.now() > sess.expires_at) {
      sessions.delete(code);
      return res.status(410).json({ ok: false, error: 'Code expired' });
    }
    sess.data = req.body;
    res.json({ ok: true, message: 'Data received — check the kiosk screen' });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// GET — kiosk polls for data (long-poll every 2s from frontend)
app.get('/api/session/poll/:code', (req, res) => {
  try {
    const code = req.params.code.toUpperCase();
    const sess = sessions.get(code);
    if (!sess) return res.json({ ok: false, status: 'expired' });
    if (Date.now() > sess.expires_at) {
      sessions.delete(code);
      return res.json({ ok: false, status: 'expired' });
    }
    if (sess.data) {
      const data = sess.data;
      sessions.delete(code); // consume it
      return res.json({ ok: true, status: 'received', data });
    }
    res.json({ ok: false, status: 'waiting' });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// PUT — update a general setting (currency, kiosk_name, mobile_base_url)
app.put('/api/settings/:key', (req, res) => {
  try {
    const allowed = ['currency', 'kiosk_name', 'mobile_base_url'];
    const key = req.params.key;
    if (!allowed.includes(key)) return res.status(400).json({ ok: false, error: 'Setting not allowed' });
    const value = req.body.value ?? '';
    db.prepare('INSERT OR REPLACE INTO settings (key, value) VALUES (?, ?)').run(key, String(value));
    res.json({ ok: true });
  } catch (e) {
    res.status(500).json({ ok: false, error: e.message });
  }
});

// ─── Start Server ────────────────────────────────────────────
app.listen(PORT, () => {
  console.log(`\n🚀 SENDIT Kiosk Server running at http://localhost:${PORT}`);
  console.log(`📦 Database: ${DB_PATH}`);
  console.log(`🔐 Admin PIN: check sendit.db settings table`);
  console.log(`📱 Mobile form: http://localhost:${PORT}/mobile\n`);
  console.log(`💡 To allow phones on ANY network: set 'mobile_base_url' in settings\n`);
});

// Graceful shutdown
process.on('SIGINT', () => {
  db.close();
  console.log('\n🔒 Database closed. Server stopped.');
  process.exit(0);
});
