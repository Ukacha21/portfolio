# SENDIT — Quick Start (Windows)

> **The `npm` / `node` error you saw is because Node.js is not installed yet.**
> Node.js is required to run the backend server. Follow Step 1 first.

---

## Step 1 — Install Node.js (one time only)

1. Go to **https://nodejs.org**
2. Download the **LTS version** (green button)
3. Run the installer — click **Next** through all steps, leave defaults as-is
4. **Restart PowerShell** after it finishes

**Verify it worked:**
```powershell
node --version   # should print something like v20.x.x
npm --version    # should print something like 10.x.x
```

---

## Step 2 — Install SENDIT Dependencies

Open PowerShell and run:

```powershell
cd "F:\Projects\Engineering\Yarismalar\Dyson\project_08_SENDIT_airport_mail_service\app"
npm install
```

This downloads: `express`, `better-sqlite3`, `cors` (~30 seconds, requires internet)

---

## Step 3 — Start the Server

```powershell
node server.js
```

Expected output:
```
🚀 SENDIT Kiosk Server running at http://localhost:3000
📦 Database: ...sendit.db
🔐 Admin PIN: check sendit.db settings table
📱 Mobile form: http://localhost:3000/mobile
💡 To allow phones on ANY network: set 'mobile_base_url' in settings
```

---

## Step 4 — Open the Kiosk

Open **http://localhost:3000** in Chrome or Edge.

| Mode | How to access |
|------|--------------|
| **User / Kiosk Mode** | Default — shown on startup |
| **Admin Mode** | Click `···` top-right → "Switch to Admin Mode" |
| **Admin PIN** | Default: `1234` |

---

## What's New

### 🔑 One-Time Code QR System (no WiFi sharing needed)

When creating a shipment (Step 2 — Addresses), click the **QR box**:

1. Kiosk generates a **6-character code** (e.g. `K7M3P4`) + QR code
2. Passenger scans the QR **or** manually goes to `localhost:3000/mobile`
3. Passenger enters the code → fills their address on their phone
4. Kiosk **auto-fills the form** in real time

> For true zero-WiFi operation (passenger uses mobile data):
> Deploy the app to a public server, then in Admin Settings set
> `Mobile URL` to your public domain, e.g. `https://sendit.yourapp.com`

### 💱 Multi-Currency

In Admin Mode → Appearance Settings → **Currency** — choose from:
`GBP · EUR · USD · CHF · AUD · TRY · QAR · SAR · BHD · KWD · EGP`

All prices convert automatically using approximate exchange rates.

---

## Files

| File | Purpose |
|------|---------|
| `app/server.js` | Node.js backend (Express + SQLite) |
| `app/package.json` | Node project manifest |
| `app/sendit.db` | SQLite database (auto-created on first run) |
| `app/mobile.html` | Passenger phone companion (one-time code flow) |
| `HARDWARE/README.md` | Raspberry Pi kiosk hardware integration guide |

---

## Database

`app/sendit.db` — SQLite, created automatically on first `node server.js`.

To view: install **DB Browser for SQLite** → https://sqlitebrowser.org
