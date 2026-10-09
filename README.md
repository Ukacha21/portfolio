<div align="center">

# Ukacha JUMA — Mechatronics Engineering Portfolio
### Biomedical Hardware · Bioimpedance Instrumentation · Embedded Systems · Polyglot

[![Go Version](https://img.shields.io/badge/Go-1.23-00ADD8.svg?logo=go)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED.svg?logo=docker)](https://www.docker.com/)
[![Render](https://img.shields.io/badge/Render-Deployable-46E3B7.svg?logo=render)](https://render.com/)
[![Languages](https://img.shields.io/badge/i18n-7%20Languages-success.svg)](#-internationalization--multilingual-support)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

<br/>

**[🌐 Live Portfolio Web Application](http://localhost:8080)** | **[📄 Project Tracker (`portf_done_web.md`)](./portf_done_web.md)**

</div>

---

## 👨‍💻 About Ukacha Juma

I am a **Mechatronics Engineer** dedicated to creating physical-digital hardware systems that prevent human suffering and solve life-critical challenges. My engineering disciplines span:

- **Biomedical Instrumentation:** Sub-epidermal moisture sensing, bioelectrical impedance spectroscopy (AD5933), tetrapolar electrode designs, and physiological heat-flux monitoring.
- **Embedded Computing & Firmware:** ESP32-S3 (Xtensa dual-core), STM32 ARM Cortex, FreeRTOS multi-threading, low-noise analog signal conditioning, and BLE 5.0 mesh telemetry.
- **Mechanical CAD & Precision Manufacturing:** Autodesk Fusion 360, SolidWorks, flexible printed circuits (FPC on Kapton® polyimide), biocompatible TPU/silicone overmolding, and rapid prototyping.
- **Polyglot & Global Communication:** Fluent in **English, Turkish, Arabic, Portuguese, Spanish, German, and French**.

---

## 📁 Repository Structure

```
.
├── cmd/
│   └── server/
├── internal/
│   ├── handlers/         # Template rendering & HTTP route handlers
│   └── i18n/             # Multilingual dictionary (7 languages + RTL engine)
├── projects/             # Sub-directories for each engineering portfolio project
│   └── project_01_velum/ # Project 01: VELUM Pressure Injury Prevention Mat
│       ├── README.md     # Full technical documentation, BOM, schematics, FreeRTOS
│       └── assets/       # High-resolution CAD renders & exploded illustrations
├── static/
│   ├── css/              # Dark obsidian & neon bio-teal design system
│   ├── js/               # Interactive 16-zone bioimpedance heatmap simulator
│   ├── img/              # Product renders, exploded diagrams, icons
│   └── docs/             # Official VELUM investor pitch deck PDF
├── templates/
│   ├── base.html         # Base HTML layout with RTL support & language switcher
│   ├── index.html        # Main portfolio portal with project cards
│   └── velum.html        # Interactive VELUM biomedical showcase
├── Dockerfile            # Multi-stage production container for Render
├── render.yaml           # Render Infrastructure-as-Code blueprint
├── portf_done_web.md     # Project completion roadmap & double-tick tracker
├── main.go               # Golang server entrypoint
└── go.mod                # Go 1.23 module
```

---

## 🚀 Projects Overview

| # | Project | Domain | Web Portfolio | GitHub Portfolio |
|---|---|---|:---:|:---:|
| **01** | [**VELUM**](./projects/project_01_velum/) | Continuous 24/7 Sub-Epidermal Bioimpedance Mat | [✅ Live Demo](/project/velum) | [✅ Readme](./projects/project_01_velum/README.md) |
| **02** | **FLOWGUARD** | Automated Microplastic Washing Machine Filtration | ⏳ In Queue | ⏳ In Queue |
| **03** | **SOLUS** | Wearable Core Temperature & Heat Stroke Patch | ⏳ In Queue | ⏳ In Queue |
| **04** | **DENOVA** | Non-Invasive Optical Tooth Decay Early Detector | ⏳ In Queue | ⏳ In Queue |
| **05** | **RENOVO** | Disaster Relief Demolition Rubble-to-Brick Press | ⏳ In Queue | ⏳ In Queue |
| **06** | **CURA SERO** | Micro-Tension Bio-Adhesive Wound Closure Device | ⏳ In Queue | ⏳ In Queue |
| **08** | **SENDIT** | Autonomous Airport Baggage & Post Logistics | ⏳ In Queue | ⏳ In Queue |
| **09** | **SMARTWASTE**| Autonomous Vision & Spectrometry Waste Sorter | ⏳ In Queue | ⏳ In Queue |
| **10** | **RESPIRO** | Mechatronic Adaptive Assistive Ventilator | ⏳ In Queue | ⏳ In Queue |

---

## 🌐 Internationalization & Multilingual Support

The web portfolio includes complete native translations across 7 languages, featuring automatic Right-to-Left (RTL) mirror layouts:

- 🇬🇧 **English** (`en`)
- 🇹🇷 **Türkçe** (`tr`)
- 🇸🇦 **العربية** (`ar` — RTL enabled)
- 🇵🇹 **Português** (`pt`)
- 🇪🇸 **Español** (`es`)
- 🇩🇪 **Deutsch** (`de`)
- 🇫🇷 **Français** (`fr`)

Language selection is persisted across requests using session cookies and query parameter overrides (`?lang=...`).

---

## 🛠️ Running Locally

### Option 1: Native Go (Zero Dependencies)

Ensure [Go 1.22+](https://go.dev/) is installed:

```bash
# Clone the repository
git clone https://github.com/ukacha21/portfolio.git
cd portfolio

# Run server
go run main.go
```

Visit `http://localhost:8080` in your browser.

### Option 2: Docker

```bash
# Build Docker image
docker build -t ukacha-portfolio .

# Run container
docker run -p 8080:8080 ukacha-portfolio
```

---

## ☁️ Deployment to Render

This repository is pre-configured with a multi-stage `Dockerfile` and `render.yaml` for automatic zero-downtime deployment on [Render](https://render.com):

1. Connect your GitHub repository `ukacha21/portfolio` to Render.
2. Select **Web Service** with the **Docker** runtime.
3. Render automatically picks up `Dockerfile` and binds the container to `$PORT`.
4. Health check endpoint `/healthz` provides continuous uptime monitoring.

---

## 📄 License & Attribution

Designed and developed by **Ukacha Juma**.  
Released under the [MIT License](LICENSE).
