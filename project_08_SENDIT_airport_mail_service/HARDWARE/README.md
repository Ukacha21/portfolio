# SENDIT Hardware Integration Guide

**Project 08 — Airport Mail Kiosk**
*Running continuously at 110V/220V mains power*

---

## Table of Contents

1. [Recommended Hardware](#1-recommended-hardware)
2. [Power Supply (Mains to Board)](#2-power-supply)
3. [Display & Touchscreen](#3-display--touchscreen)
4. [QR Code Scanner](#4-qr-code-scanner)
5. [Label Printer](#5-label-printer)
6. [Network Connectivity](#6-network-connectivity)
7. [Peripheral Wiring Summary](#7-peripheral-wiring-summary)
8. [OS Setup — Raspberry Pi](#8-os-setup--raspberry-pi)
9. [Software Installation](#9-software-installation)
10. [Kiosk Mode (Auto-Start Browser)](#10-kiosk-mode-auto-start-browser)
11. [Watchdog — Auto-Restart on Crash](#11-watchdog--auto-restart-on-crash)
12. [24/7 Reliability Checklist](#12-247-reliability-checklist)
13. [Optional: ESP32 for Peripheral Control](#13-optional-esp32-for-peripheral-control)
14. [Enclosure & Physical Installation](#14-enclosure--physical-installation)
15. [Troubleshooting](#15-troubleshooting)

---

## 1. Recommended Hardware

### Primary Controller: **Raspberry Pi 4 Model B (4GB RAM)**

| Component | Recommended Model | Notes |
|---|---|---|
| **Main Board** | Raspberry Pi 4B — 4GB | Quad-core 1.8GHz, WiFi/BT built-in |
| **Storage** | 64GB A2-class microSD (e.g. SanDisk Extreme) | A2 = faster random I/O for SQLite |
| **Display** | 15.6" 1920×1080 HDMI Touchscreen | Look for HMTECH or SunFounder models |
| **QR Scanner** | Honeywell Voyager 1250g (USB HID) | Plug-and-play, no drivers needed |
| **Label Printer** | Zebra ZD421 or DYMO LabelWriter 550 | USB, works on Pi natively |
| **Cooling** | Argon ONE M.2 case OR GeeekPi heatsink+fan | Required for continuous operation |
| **UPS Hat** | Waveshare UPS HAT (C) | Protects against power outages |

### Why Raspberry Pi 4 over ESP32?

| | Raspberry Pi 4 | ESP32 |
|---|---|---|
| **OS** | Full Linux (Raspberry Pi OS) | No OS (bare metal or RTOS) |
| **Browser** | Chromium (full kiosk browser) | ❌ No browser |
| **Node.js** | ✅ Native support | ❌ Not supported |
| **SQLite** | ✅ Full support | ❌ No filesystem |
| **Touchscreen** | ✅ Plug-and-play via HDMI | Requires custom TFT |
| **Use case** | Full kiosk application | Sensor/actuator peripheral |

> **ESP32 is best used as a secondary controller** — for reading hardware sensors, controlling a locker latch, buzzer, LEDs, or RFID reader. It communicates with the Pi via USB serial or UART.

---

## 2. Power Supply

### From Wall (110V/220V AC) → Raspberry Pi 4 (5V/3A DC)

**DO NOT use cheap phone chargers.** The Pi 4 draws up to 3A peak and needs a stable 5V. Undervoltage causes filesystem corruption and crashes.

### Recommended:

```
Wall Outlet (110V or 220V AC)
       │
       ▼
┌──────────────────────────────────┐
│ Official Raspberry Pi 4 PSU      │
│ Input:  100–240V AC, 50/60Hz     │
│ Output: 5.1V DC, 3.0A            │
│ Connector: USB-C                 │
│ Part: SC0218 (Raspberry Pi Ltd.) │
└──────────────────────────────────┘
       │
       ▼ USB-C
 Raspberry Pi 4
```

**If your location uses 110V (e.g. USA/Turkey):** The official PSU is universal (100–240V) — no adapter needed, just the right plug type.

**If you use the Argon ONE case** (recommended): it includes its own power connector and fan controller — still powered by the official USB-C PSU.

### For the Display (if not USB-powered):

Most 15.6" HDMI touchscreens include their own 12V/2A power adapter (universal voltage). Connect to wall separately or via a power strip inside the kiosk enclosure.

### Recommended Setup Inside Enclosure:

```
Wall Socket (C13 IEC inlet)
       │
       ▼
Power Strip / IEC Hub (inside enclosure)
   │           │             │
   ▼           ▼             ▼
Pi 4 PSU   Display PSU   USB Hub PSU (if needed)
(5.1V 3A) (12V 2A)      (5V 2.5A)
```

---

## 3. Display & Touchscreen

### Recommended: 15.6" 1080p HDMI Capacitive Touchscreen

**Example models:**
- HMTECH HM156WT01 15.6" IPS Touchscreen (~$120)
- SunFounder 15.6" Raspberry Pi Monitor

**Connection:**
```
Pi 4 HDMI port (micro-HDMI)
    │
    ▼ (via micro-HDMI to HDMI cable)
Display HDMI IN

Display USB Touch OUT
    │
    ▼ (via USB-A cable)
Pi 4 USB-A port (touch input)
```

**Configure `/boot/config.txt` for proper resolution:**
```ini
hdmi_group=2
hdmi_mode=82          # 1920x1080 @ 60Hz
display_rotate=0      # 0=landscape, 1=portrait
```

**Touchscreen calibration** (one-time):
```bash
sudo apt install xinput-calibrator
DISPLAY=:0 xinput_calibrator
```

### For a wall-mounted or angled kiosk:
- Use **VESA 75×75 mount** behind the display
- Pi can be mounted on the back of the display with a Pi case that has VESA holes

---

## 4. QR Code Scanner

### Option A: USB HID QR Scanner (Recommended — simplest)

**Model:** Honeywell Voyager 1250g or any generic USB HID 1D/2D barcode scanner

- Plug into Pi USB port
- Appears as a keyboard (HID device)
- No drivers needed on Linux
- Scans QR codes and types the data directly into the focused text field

**How it works with SENDIT:**
1. Kiosk displays QR code (the mobile session code shown on passenger's phone)
2. Staff/passenger scans it with the handheld scanner
3. Scanner "types" the session ID into the tracking input field
4. SENDIT app polls `/api/mobile/poll/:session_id` and auto-fills the form

> [!TIP]
> For fully hands-free kiosk: mount a fixed-angle 2D scanner (e.g. Zebra DS457) in the kiosk frame pointing at a marked scanning zone. No staff interaction needed.

### Option B: Pi Camera + ZBar (Software QR Decode)

For a camera-based approach:
```bash
sudo apt install zbar-tools python3-picamera2
```

Script to capture and decode:
```python
import subprocess, time
from picamera2 import Picamera2

cam = Picamera2()
cam.start()
while True:
    cam.capture_file('/tmp/qr.jpg')
    result = subprocess.run(['zbarimg', '--raw', '/tmp/qr.jpg'], capture_output=True, text=True)
    if result.stdout.strip():
        print('QR:', result.stdout.strip())
    time.sleep(0.5)
```

---

## 5. Label Printer

### Recommended: Zebra ZD421 (USB) or DYMO LabelWriter 550

**Setup on Raspberry Pi:**
```bash
sudo apt install cups
sudo systemctl enable cups
sudo usermod -aG lpadmin pi

# Add printer via CUPS web interface
# Navigate to: http://localhost:631
```

**For Zebra ZD421** — install CUPS driver:
```bash
# Download Zebra CUPS driver from zebra.com
# Or use ZPLII labels via direct USB raw print:
echo "^XA^FO50,50^A0N,40,40^FDTracking: SD-2026-001921^FS^XZ" > /dev/usb/lp0
```

**In the SENDIT app**, the `printLabel()` function sends a `window.print()` call or can be extended to POST to a `/api/print` endpoint that pipes ZPL to the printer.

---

## 6. Network Connectivity

### Recommended: **Wired Ethernet (Primary) + WiFi (Backup)**

The Raspberry Pi 4 has both built-in:

```
Airport Network Switch/Router
         │
         ▼ (Cat5e/Cat6 Ethernet cable)
   Pi 4 Ethernet Port (eth0)
         │
         ▼
   SENDIT Server (localhost:3000)
```

**Configure static IP** (so the kiosk URL never changes):

Edit `/etc/dhcpcd.conf`:
```
interface eth0
static ip_address=192.168.1.50/24
static routers=192.168.1.1
static domain_name_servers=8.8.8.8 1.1.1.1
```

**Mobile QR form** accessible at: `http://192.168.1.50:3000/mobile`

> [!IMPORTANT]
> Display the kiosk's local IP on-screen so passengers and staff know the mobile URL. You can show it in the kiosk footer or print it on a sign beside the machine.

---

## 7. Peripheral Wiring Summary

```
┌─────────────────────────────────────────────────────────────┐
│                   SENDIT Kiosk Enclosure                    │
│                                                             │
│  ┌─────────────────┐                                        │
│  │  15.6" Display  │◄── HDMI ──┐   USB Touch ──┐           │
│  └─────────────────┘           │               │           │
│                                │               │           │
│  ┌──────────────────────────────▼───────────────▼────────┐  │
│  │               Raspberry Pi 4B (4GB)                   │  │
│  │  GPIO 40-pin header (optional ESP32 UART)             │  │
│  │  USB-A × 4:  [QR Scanner] [Label Printer] [USB Hub]  │  │
│  │  Ethernet:   [Airport LAN]                            │  │
│  │  micro-HDMI: [Display]                                │  │
│  │  USB-C:      [5.1V PSU]                               │  │
│  └───────────────────────────────────────────────────────┘  │
│                                                             │
│  ┌───────────────┐  ┌──────────────┐  ┌──────────────────┐  │
│  │  QR Scanner   │  │ Label Printer│  │ UPS HAT (optional│  │
│  │  (USB HID)    │  │ (Zebra/DYMO) │  │ for power safety)│  │
│  └───────────────┘  └──────────────┘  └──────────────────┘  │
│                                                             │
│  Power Strip ◄── IEC C13 ◄── Wall Socket (110V/220V)       │
└─────────────────────────────────────────────────────────────┘
```

---

## 8. OS Setup — Raspberry Pi

### Step 1: Flash OS

1. Download **Raspberry Pi OS Lite (64-bit)** from [raspberrypi.com/software](https://raspberrypi.com/software)
2. Use **Raspberry Pi Imager** to flash to microSD
3. In Imager settings (gear icon):
   - Set hostname: `sendit-kiosk`
   - Enable SSH
   - Set WiFi credentials (backup network)
   - Set username: `pi`, password of your choice

### Step 2: First Boot

```bash
# SSH in (or connect keyboard/monitor)
ssh pi@sendit-kiosk.local

# Update system
sudo apt update && sudo apt upgrade -y

# Install essential packages
sudo apt install -y \
  git curl wget \
  xorg openbox chromium-browser \
  unclutter \
  cups cups-bsd \
  ufw
```

### Step 3: Firewall

```bash
sudo ufw allow 22/tcp    # SSH
sudo ufw allow 3000/tcp  # SENDIT server
sudo ufw enable
```

---

## 9. Software Installation

### Install Node.js (v20 LTS)

```bash
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install -y nodejs
node --version   # should show v20.x.x
npm --version    # should show 10.x.x
```

### Install SENDIT App

```bash
# Copy project files to Pi (via scp, USB, or git)
# Assuming files are at ~/sendit/app/

cd ~/sendit/app
npm install

# Test the server manually
node server.js
# You should see: 🚀 SENDIT Kiosk Server running at http://localhost:3000
```

### Create Systemd Service (Auto-Start on Boot)

```bash
sudo nano /etc/systemd/system/sendit.service
```

Paste:
```ini
[Unit]
Description=SENDIT Kiosk Server
After=network.target

[Service]
Type=simple
User=pi
WorkingDirectory=/home/pi/sendit/app
ExecStart=/usr/bin/node /home/pi/sendit/app/server.js
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
sudo systemctl status sendit  # should show "active (running)"
```

---

## 10. Kiosk Mode (Auto-Start Browser)

### Configure Openbox to Launch Chromium on Boot

```bash
mkdir -p ~/.config/openbox
nano ~/.config/openbox/autostart
```

Paste:
```bash
# Hide mouse cursor after 1 second of inactivity
unclutter -idle 1 &

# Launch Chromium in kiosk mode
chromium-browser \
  --kiosk \
  --no-sandbox \
  --disable-infobars \
  --disable-session-crashed-bubble \
  --disable-restore-session-state \
  --noerrdialogs \
  --disable-pinch \
  --overscroll-history-navigation=0 \
  --check-for-update-interval=31536000 \
  http://localhost:3000 &
```

### Auto-Login to Desktop (for Openbox)

```bash
sudo raspi-config
# → System Options → Boot / Auto Login → Desktop Autologin
```

Or manually edit `/etc/lightdm/lightdm.conf`:
```ini
[SeatDefaults]
autologin-user=pi
autologin-session=openbox
```

### Set Openbox as default session:

```bash
sudo update-alternatives --config x-session-manager
# Select openbox-session
```

---

## 11. Watchdog — Auto-Restart on Crash

### Enable Pi Hardware Watchdog

```bash
sudo nano /etc/systemd/system.conf
```

Add/uncomment:
```
RuntimeWatchdogSec=15
ShutdownWatchdogSec=2min
```

### Enable bcm2835 watchdog module:

```bash
echo 'dtparam=watchdog=on' | sudo tee -a /boot/config.txt
echo 'bcm2835_wdt' | sudo tee -a /etc/modules
sudo modprobe bcm2835_wdt
```

### Watchdog daemon:

```bash
sudo apt install watchdog
sudo nano /etc/watchdog.conf
```

Add:
```
watchdog-device = /dev/watchdog
watchdog-timeout = 15
interval = 10
realtime = yes
priority = 1
min-memory = 1
```

```bash
sudo systemctl enable watchdog
sudo systemctl start watchdog
```

> If the Pi hangs or the process locks up for more than 15 seconds, it **automatically reboots**.

### Chromium Watchdog Script

```bash
nano ~/sendit/watchdog_browser.sh
```

```bash
#!/bin/bash
# Restart Chromium if it crashes
while true; do
  if ! pgrep -x "chromium-browse" > /dev/null; then
    DISPLAY=:0 chromium-browser \
      --kiosk \
      --no-sandbox \
      http://localhost:3000 &
    echo "$(date): Chromium restarted" >> /var/log/sendit_watchdog.log
  fi
  sleep 10
done
```

```bash
chmod +x ~/sendit/watchdog_browser.sh
```

Add to autostart:
```bash
~/sendit/watchdog_browser.sh &
```

---

## 12. 24/7 Reliability Checklist

| Item | Action |
|------|--------|
| **SD Card Quality** | Use A2-class (SanDisk Extreme, Samsung PRO Endurance) |
| **Power** | Official Pi PSU only — no cheap chargers |
| **UPS** | Waveshare UPS HAT prevents DB corruption on power cuts |
| **Cooling** | Argon ONE case or active cooling — Pi throttles at 85°C |
| **SQLite WAL mode** | ✅ Already enabled in server.js |
| **Database backup** | Daily cron: `cp sendit.db sendit_backup_$(date +%Y%m%d).db` |
| **Log rotation** | `sudo logrotate` configured for journald |
| **Auto-updates** | `sudo apt install unattended-upgrades` (security only) |
| **Screen timeout** | Disable screen blanking: `xset s off; xset -dpms` |
| **Read-only FS** | Optional: overlay filesystem to protect SD from corruption |

### Daily Backup Cron

```bash
crontab -e
```

Add:
```cron
0 3 * * * cp /home/pi/sendit/app/sendit.db /home/pi/sendit/backups/sendit_$(date +\%Y\%m\%d).db
# Keep last 30 days only
0 4 * * * find /home/pi/sendit/backups -name "*.db" -mtime +30 -delete
```

---

## 13. Optional: ESP32 for Peripheral Control

Use an **ESP32 DevKit** for low-level hardware:
- Buzzer (beep on scan)
- Status LEDs (green/red/amber)
- Servo/solenoid (locker door release)
- RFID reader (NXP MFRC522)
- PIR sensor (wake screen when someone approaches)

### ESP32 ↔ Pi Communication (USB Serial)

**ESP32 code (Arduino):**
```cpp
void loop() {
  if (Serial.available()) {
    String cmd = Serial.readStringUntil('\n');
    if (cmd == "BEEP") {
      tone(BUZZER_PIN, 1000, 200);
    }
    if (cmd == "LED_GREEN") {
      digitalWrite(LED_G, HIGH);
    }
  }
}
```

**Pi Node.js side (in server.js):**
```javascript
const { SerialPort } = require('serialport');
const port = new SerialPort({ path: '/dev/ttyUSB0', baudRate: 115200 });

// Call from API endpoint
function sendToESP32(command) {
  port.write(command + '\n');
}
```

### ESP32 Power

- Power from Pi USB port (5V via USB-A) OR
- Dedicated 5V/1A adapter (recommended for continuous use)

---

## 14. Enclosure & Physical Installation

### Recommended Kiosk Enclosure

- **Metalwork/Acrylic custom box** OR
- **Off-the-shelf kiosk enclosure**: Armodilo, DSE, or custom fabricated
- Screen mounted at **eye level for standing adults** (~140–160cm from floor)
- Anti-glare screen protector (airports have bright lights)
- **Touchscreen orientation**: Landscape (16:9)

### Cable Management Inside Enclosure

```
Back panel:
├── C13 IEC inlet (screwed in)
├── Internal power strip (DIN rail mount)
├── Cable ties for all internal cables
└── Ventilation slots for Pi cooling
```

### Security

- Physical lock on enclosure access panel
- Anti-tamper VESA screws for display
- Pi secured with standoffs inside enclosure
- USB ports on Pi taped/blocked (except the ones in use) to prevent unauthorized USB devices

---

## 15. Troubleshooting

| Problem | Cause | Fix |
|---|---|---|
| Chromium shows blank page | Server not started | `sudo systemctl status sendit` |
| "Server Offline" toast | Node not running | `sudo systemctl restart sendit` |
| Touchscreen not responding | USB touch not connected | Check USB cable from display to Pi |
| Display wrong resolution | HDMI config | Edit `/boot/config.txt`, set `hdmi_mode` |
| Database not persisting | Wrong path or permissions | Check `ls -la ~/sendit/app/sendit.db` |
| Pi overheating (throttling) | No cooling | Add heatsink + fan |
| Screen goes blank | Screen blanking enabled | Run `xset s off; xset -dpms` in autostart |
| QR scanner not detected | USB not recognized | Run `lsusb` to verify device seen |

---

## Quick Start Summary

```bash
# 1. Flash Raspberry Pi OS Lite (64-bit) with Pi Imager
# 2. Boot Pi, SSH in, update system
# 3. Install Node.js v20
# 4. Copy SENDIT app files to ~/sendit/app/
# 5. cd ~/sendit/app && npm install
# 6. Create systemd service (see above)
# 7. Configure Openbox kiosk autostart
# 8. Reboot → Pi boots to Chromium showing SENDIT kiosk
```

---

*Last updated: 2026 · SENDIT Project 08 · Dyson Award Entry*
