# SENDIT — Go Backend (Gin + SQLite)

> **Backend API & database management using Go + Gin framework**
> This is an alternative backend to the Node.js version. The frontend (`index.html`, `style.css`) is shared and works with any backend by pointing to `http://localhost:3000/api`.

---

## 📁 Project Structure

```
app_golang/
├── main.go          ← Gin HTTP server & all API routes
├── database.go      ← SQLite DB setup, queries & helpers
├── models.go        ← Request/response struct definitions
├── go.mod           ← Go module definition
├── go.sum           ← Dependency checksums (auto-generated)
└── README.md        ← This file
```

---

## 🖥️ Running on PC (Testing / Development)

### Prerequisites
- **Go 1.21+** → https://go.dev/dl/ (pure Go SQLite driver is used, no compiler required)

### Step 1 — Navigate to this folder
```bash
cd app_golang
```

### Step 2 — Download dependencies
```bash
go mod tidy
```

### Step 3 — Run the server
```bash
go run .
```

Or build a binary first:
```bash
go build -o sendit-server .
./sendit-server
```

The API will be available at: **http://localhost:3000**

---

## 🍓 Running on Raspberry Pi 4/5 (Kiosk / Production)

Go compiles to a **single native binary** — no runtime needed on the Pi. Just cross-compile or build directly on the Pi.

### Option A — Cross-compile from PC (faster)
On your **Windows PC**:
```powershell
# Set target to Raspberry Pi 4/5 (64-bit ARM)
$env:GOOS = "linux"
$env:GOARCH = "arm64"
$env:CGO_ENABLED = "0"
go build -o sendit-server .
```

> 💡 **CGO note**: Because we use a pure Go SQLite library (`modernc.org/sqlite`), cross-compiling works out of the box on Windows without a cross-compiler.

### Option B — Compile directly on the Raspberry Pi
```bash
# On the Pi: install Go
wget https://go.dev/dl/go1.21.0.linux-arm64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-arm64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# (No GCC or compiler compilation tools are needed since this is pure Go)

# Copy project files to Pi first
# From your PC:
# scp -r ./app_golang pi@<PI_IP>:/home/pi/sendit/app_golang

# Then on the Pi:
cd /home/pi/sendit/app_golang
go mod tidy
go build -o sendit-server .
./sendit-server
```

### Step — Set up systemd for auto-start on boot
```bash
sudo nano /etc/systemd/system/sendit.service
```

Paste:
```ini
[Unit]
Description=SENDIT Kiosk API Server (Go)
After=network.target

[Service]
User=pi
WorkingDirectory=/home/pi/sendit/app_golang
ExecStart=/home/pi/sendit/app_golang/sendit-server
Restart=always
RestartSec=5
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
```

Enable and start:
```bash
sudo systemctl daemon-reload
sudo systemctl enable sendit
sudo systemctl start sendit
sudo systemctl status sendit
```

### Serving the frontend on the Pi
The Go server automatically serves `index.html`, `mobile.html`, and `style.css` from its working directory. Point Chromium at `http://localhost:3000`:

```bash
chromium-browser --kiosk --noerrdialogs --disable-infobars http://localhost:3000
```

---

## 📦 Go Dependencies (`go.mod`)

```
github.com/gin-gonic/gin        v1.9.1   ← HTTP router
modernc.org/sqlite              v1.54.0  ← SQLite driver (Pure Go)
```

To install:
```bash
go mod tidy
```

---

## 🔌 API Endpoints

All endpoints match the Node.js and Python versions — the frontend is fully compatible.

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/stats` | Server health & package stats |
| GET | `/api/shipments` | List all shipments |
| POST | `/api/shipments` | Create new shipment |
| GET | `/api/shipments/:id` | Get shipment by ID |
| PUT | `/api/shipments/:id/status` | Update shipment status |
| DELETE | `/api/shipments/:id` | Delete shipment |
| GET | `/api/shipments/track/:code` | Track by tracking number |
| POST | `/api/auth/verify-pin` | Verify admin PIN |
| PUT | `/api/settings/pin` | Change admin PIN |
| GET | `/api/settings` | Get app settings |
| PUT | `/api/settings/:key` | Update a setting |
| POST | `/api/session/create` | Create QR mobile session |
| GET | `/api/session/poll/:code` | Poll session data |
| POST | `/api/session/submit/:code` | Submit mobile form data |

---

## 🛠️ Notes

- **Database file**: `sendit.db` is created automatically in the `app_golang/` directory.
- **Default admin PIN**: `1234` — change from the admin dashboard after first login.
- **Binary size**: The compiled Go binary is typically ~15–25MB and includes everything.
- **Performance**: Go is ideal for the Raspberry Pi — very low memory footprint (~10MB RAM), near-instant startup, and excellent SQLite performance.
- **No runtime required**: Unlike Python/Node.js, Go compiles to a native binary — no interpreter needed on the Pi.
