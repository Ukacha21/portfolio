// Ukacha Juma Portfolio — Main JavaScript Interactions

document.addEventListener('DOMContentLoaded', () => {
    // 1. Language Dropdown Handling
    const langBtn = document.getElementById('langDropdownBtn');
    const langMenu = document.getElementById('langMenu');

    if (langBtn && langMenu) {
        langBtn.addEventListener('click', (e) => {
            e.stopPropagation();
            langMenu.classList.toggle('open');
        });

        document.addEventListener('click', () => {
            langMenu.classList.remove('open');
        });
    }

    // 2. Mobile Menu Toggle
    const mobileToggle = document.getElementById('mobileToggle');
    const mainNav = document.getElementById('mainNav');

    if (mobileToggle && mainNav) {
        mobileToggle.addEventListener('click', () => {
            mainNav.classList.toggle('open');
        });

        mainNav.querySelectorAll('.nav-link').forEach(link => {
            link.addEventListener('click', () => {
                mainNav.classList.remove('open');
            });
        });
    }

    // 3. Project Preview Modals
    const projectData = {
        flowguard: {
            title: "FLOWGUARD — Microplastic Laundry Filtration",
            domain: "Fluidic Mechatronics · Microplastic Separation",
            status: "In Queue (Development Phase)",
            tags: ["Hydro-Cyclonic", "Automated Backwash", "Continuous Pressure Sensing", "ESP32 C3"],
            desc: "Automated mechatronic microplastic filtration system for residential and commercial washing machines. Features dual-stage hydro-cyclonic separation and a self-cleaning micro-mesh filter, intercepting synthetic microfibers before municipal discharge.",
            specs: "Filtration efficiency: 94%+ down to 10 µm · Automated backwash pressure trigger: 1.2 bar · Low-power solenoid purge valve · Embedded turbidity sensor.",
            link: "#"
        },
        solus: {
            title: "SOLUS — Heat Stroke Continuous Biosensor Patch",
            domain: "Wearable Biosensors · Preventative Physiology",
            status: "In Queue (Development Phase)",
            tags: ["Dual-Heat Flux", "BLE 5.0", "Predictive Heat Exhaustion", "ISO 10993"],
            desc: "Wearable biomedical patch that calculates real-time core body temperature using non-invasive dual-heat flux sensors. Built for high-heat industrial laborers, soldiers, and athletes.",
            specs: "Core temp accuracy: ±0.15°C vs ingestible telemetry pill · Battery runtime: 72 hours continuous · BLE telemetry · Early prognostic warning prior to clinical collapse.",
            link: "#"
        },
        denova: {
            title: "DENOVA (formerly Avicenna) — Non-Invasive Optical Tooth Decay Detector",
            domain: "Biophotonics · Optical Diagnostics",
            status: "In Queue (Development Phase)",
            tags: ["405nm Fluorescence", "Multispectral LED", "Early Demineralization", "CMOS Probe"],
            desc: "Diagnostic dental handpiece utilizing 405nm light-induced fluorescence to detect early sub-surface enamel demineralization before visible cavitation occurs.",
            specs: "Narrowband 405nm excitation · Custom dichroic filter assembly · Micro-CMOS imaging sensor · Embedded fluorescence ratio analysis algorithm.",
            link: "#"
        },
        renovo: {
            title: "RENOVO — Rubble-to-Brick Hydraulic Press",
            domain: "Heavy Mechanical & Hydraulics · Disaster Relief",
            status: "In Queue (Development Phase)",
            tags: ["Hydraulic Compaction", "Interlocking Masonry", "Disaster Recovery", "Off-Grid"],
            desc: "Rugged, portable hydraulic brick press engineered for rapid post-earthquake reconstruction. Crushes masonry rubble and compacts it into interlocking structural bricks without mortar.",
            specs: "Compaction pressure: 15 MPa · Hydraulic cylinder: 20-ton rated · Cycle time: 45 seconds per brick · Compatible with gasoline or solar-electric power units.",
            link: "#"
        },
        curasero: {
            title: "CURA SERO (AVICENNA) — Wet-Surface Wound Closure Strip",
            domain: "Bio-Mechanical Device · Surgical Technology",
            status: "🔬 Unfinished Research (James Dyson Award 2026 Submission)",
            tags: ["Biomimetic Gecko Adhesion", "Micro-Ratcheting Tension", "Wet-Skin Grip", "Slow-Release Chlorhexidine", "Unfinished Research"],
            desc: "Biomimetic wound closure strip using gecko-inspired micro-suction that bonds to wet, bleeding skin, applies suture-equivalent wound tension via micro-ratchet, and delivers slow-release antimicrobial protection in last-mile healthcare.",
            specs: "Tensile distribution: 0.8 N/cm² uniform (3-0 suture equivalent) · 10–50 µm hexagonal silicone micro-pillars · 48–72h chlorhexidine hydrogel slow-release · Pre-clinical validation stage.",
            link: "/project/curasero"
        },
        welstech: {
            title: "WELSTECH — Wearable Life Saver Technology",
            domain: "Biomedical Wearable IoT · Cardiac Biosensing",
            status: "Complete Hardware Prototype",
            tags: ["ESP32-S3 Zero", "Optical PPG", "852 ADC Threshold", "Thermal Relief CAD", "Flutter Telemetry", "192.168.4.1 AP"],
            desc: "Ultra-compact wrist-worn vital signs monitor integrating optical PPG biosensing, ESP32-S3 processing, local Wi-Fi telemetry, and acoustic alarm escalation for cardiac anomaly detection.",
            specs: "Microcontroller: ESP32-S3 Zero · Sampling: 200 Hz PPG via ADC A0 · Threshold: Schmitt-trigger 852 ADC · Dimensions: 43 mm diameter CAD enclosure · Local HTTP AP: 192.168.4.1/data.",
            link: "/project/welstech"
        },
        sendit: {
            title: "SENDIT — Airport Item Mail-Home Kiosk",
            domain: "Smart Airport Infrastructure · Courier Automation",
            status: "Complete Showcase (Autonomous Airport Logistics)",
            tags: ["AI Vision Classification", "QR Contact Auto-Fill", "Zebra Thermal Labels", "SumUp Contactless", "Zero Airport Capex"],
            desc: "Smart self-service kiosk at airport and train station security checkpoints that packages, classifies, labels, weighs, and dispatches travellers' surrendered items home in under 2 minutes.",
            specs: "Interaction time: <120 seconds · SBC: Raspberry Pi 4/5 + Coral TPU · Zebra ZD421 203dpi thermal printer · Integrated 12V nichrome heat-sealer · SumUp Pro contactless reader.",
            link: "/project/sendit"
        },
        smartwaste: {
            title: "SMARTWASTE — Autonomous Waste Sorting & 3:1 Compactor",
            domain: "Autonomous Mechatronics · Smart Municipal IoT",
            status: "Complete Showcase (James Dyson Award '26 Submission)",
            tags: ["James Dyson Award 2026", "ToF Laser Telemetry", "Bimodal Sorting Gate", "3:1 Lead-Screw Press", "Deep Sleep ESP32", "Konya Selçuklu Deployment"],
            desc: "Retrofittable IoT bin module combining millimetre-accurate VCSEL laser ToF fill-level telemetry, bimodal inductive/optical sorting gates, and an automated 1.2 kN lead-screw 3:1 compactor to slash municipal collection mileage at under $80 BOM.",
            specs: "Compaction ratio: 3:1 · BOM cost: < $80 · Sensor: VL53L1X Laser ToF (±3mm) · Battery autonomy: 180+ days (15 µA deep sleep) · Municipal route reduction: 34%.",
            link: "/project/smartwaste"
        },
        classifier: {
            title: "IMAGE CLASSIFIER — Deep Convolutional Neural Network",
            domain: "Deep Learning · Computer Vision & Edge AI",
            status: "Complete Showcase (Active Deep Learning Model)",
            tags: ["TensorFlow / Keras", "Conv2D & MaxPool", "96×96 RGB Preprocessing", "40-Epoch Converged Weights", "Tkinter & OpenCV GUI"],
            desc: "Custom deep vision classifier trained from scratch in TensorFlow/Keras on 9,600+ curated images with dynamic aspect-preserving PIL preprocessing, three-stage Conv2D and Max-Pooling hierarchical topology, and desktop GUI validators.",
            specs: "Framework: TensorFlow 2.17 · Input: 96×96×3 RGB · Trainable parameters: 2,010,306 (~7.67 MB uncompressed) · Checkpoints: 3 epochs & 40 epochs (.keras) · Inference latency: <15ms.",
            link: "/project/classifier"
        },
        lumiere: {
            title: "LUMIERE — Smart Streetlight Failure & Geolocation Telemetry System",
            domain: "Smart City IoT · Embedded Firmware & Telemetry",
            status: "Complete Hardware Prototype",
            tags: ["ESP32 SoftAP", "GY-NEO6MV2 GPS (5m)", "CdS LDR Sensor", "UART1 NMEA Telemetry", "Escalating Piezo Alarm", "Sub-$15 BOM"],
            desc: "Compact IoT streetlight failure reporting device mounted directly beneath municipal luminaires, combining LDR photometric sensing, satellite GPS positioning, local Wi-Fi SoftAP telemetry, and escalating acoustic alarms to automate outage dispatch.",
            specs: "Microcontroller: ESP32 Dual-Core 240MHz · Sensor: CdS LDR 12-bit ADC (GPIO 34) · GPS: GY-NEO6MV2 (UART1 9600 baud, 5m fix) · Network: Wi-Fi SoftAP (SSID: Lumiere, 192.168.4.1) · Power: 78.3 mW active / 26.8 µW deep sleep · Target unit BOM: < $15 USD.",
            link: "/project/lumiere"
        },
        notcha: {
            title: "NOTCHA — Minimal Floating Dynamic Time Notch",
            domain: "Desktop Ergonomics · Custom UI Utilities",
            status: "Complete Desktop Application",
            tags: ["Python 3", "Tkinter GUI", "Auto-Hide & Hover", "Top Bezel Notch", "Win32 Transparency", "Zero DLL Injection"],
            desc: "Born out of ergonomic necessity in 2021/2022 when Windows 11 removed top taskbar placement, NOTCHA is an ultra-lightweight desktop utility providing an unobtrusive floating time notch at the top screen bezel with real-time updates and smart hover-expand mechanics.",
            specs: "Framework: Python 3 / Tkinter · Dimensions: 414×34 px (Active) / 414×1 px (Collapsed) · Memory: ~18 MB RAM · Binary: Notcha.exe (10.2 MB Standalone PyInstaller) · Overhead: <0.1% CPU · Refresh: 1 Hz Non-blocking threading.Timer.",
            link: "/project/notcha"
        },
        spo2: {
            title: "SPO2 — Pulse Oximeter 2-Layer PCB Design",
            domain: "PCB Design & EDA · Biomedical Hardware · KiCad 10.0",
            status: "Complete 2-Layer PCB Design",
            tags: ["KiCad 10.0 EDA", "2-Layer FR4", "MAX30102 Biosensor", "SSD1306 OLED", "Custom Footprints", "Zero Unrouted Nets", "DRC Verified"],
            desc: "A tested 2-layer pulse oximeter circuit board engineered in KiCad 10.0, bridging custom schematic symbols, footprints built from scratch, real manufacturer tolerances, MAX30102 biosensing, SSD1306 OLED telemetry, and an Arduino Nano core.",
            specs: "Layers: 2 (F.Cu + B.Cu) · EDA: KiCad 10.0 · Core: Arduino Nano v3 (ATmega328P) · Sensor: MAX30102 Optical PPG · Display: SSD1306 0.96\" I2C OLED · Alarm: D1 Red LED · Debounce: 1kΩ + 100nF RC networks · Nets: 31 (0 unrouted) · Target: JLCPCB / PCBWay manufacturing rules.",
            link: "/project/spo2"
        },
        respiro: {
            title: "RESPIRO — Biomedical Adaptive Assistive Ventilator (Discontinued)",
            domain: "Biomedical Mechatronics · Critical Care",
            status: "🚫 Discontinued (Archived R&D)",
            tags: ["Adaptive Ventilation", "Closed-Loop PID", "Discontinued", "Archived R&D"],
            desc: "Low-cost, high-reliability mechanical ventilator concept engineered for emergency deployment and resource-limited clinics. Discontinued during early prototyping in favor of primary clinical and hardware initiatives.",
            specs: "Tidal volume: 200–800 mL · PEEP regulation: 0–20 cmH2O · Concept archived; discontinued from development roadmap.",
            link: "#"
        }
    };

    const modal = document.getElementById('projectModal');
    const modalContent = document.getElementById('modalContent');
    const modalCloseBtn = document.getElementById('modalCloseBtn');

    if (modal && modalContent) {
        document.querySelectorAll('.preview-btn').forEach(btn => {
            btn.addEventListener('click', () => {
                const pKey = btn.dataset.project;
                const p = projectData[pKey];
                if (!p) return;

                modalContent.innerHTML = `
                    <div style="font-family: var(--font-mono); font-size: 11px; color: var(--teal-primary); letter-spacing: 1px; margin-bottom: 8px;">
                        ${p.domain}
                    </div>
                    <h3 style="font-size: 24px; font-weight: 800; color: #fff; margin-bottom: 8px;">
                        ${p.title}
                    </h3>
                    <div style="display: inline-block; font-family: var(--font-mono); font-size: 11px; padding: 3px 10px; background: rgba(244, 197, 66, 0.15); border: 1px solid rgba(244, 197, 66, 0.35); color: var(--status-warn); border-radius: var(--radius-full); margin-bottom: 18px;">
                        ${p.status}
                    </div>
                    <p style="font-size: 14.5px; color: var(--text-secondary); line-height: 1.6; margin-bottom: 18px;">
                        ${p.desc}
                    </p>
                    <div style="background: rgba(4, 9, 18, 0.7); border: 1px solid var(--border-subtle); border-radius: var(--radius-md); padding: 14px; margin-bottom: 20px;">
                        <strong style="display: block; font-size: 12.5px; color: var(--cyan-accent); font-family: var(--font-mono); margin-bottom: 4px;">ENGINEERING SPECIFICATIONS:</strong>
                        <span style="font-size: 13px; color: var(--text-secondary); line-height: 1.5;">${p.specs}</span>
                    </div>
                    <div style="display: flex; flex-wrap: wrap; gap: 6px; margin-bottom: 24px;">
                        ${p.tags.map(t => `<span class="tag">${t}</span>`).join('')}
                    </div>
                    <div style="display: flex; justify-content: flex-end; gap: 10px;">
                        <button class="btn btn-secondary" onclick="document.getElementById('projectModal').classList.remove('open')">
                            Close
                        </button>
                    </div>
                `;
                modal.classList.add('open');
            });
        });

        if (modalCloseBtn) {
            modalCloseBtn.addEventListener('click', () => {
                modal.classList.remove('open');
            });
        }

        modal.addEventListener('click', (e) => {
            if (e.target === modal) {
                modal.classList.remove('open');
            }
        });
    }

    // =========================================================================
    // IMAGE CLASSIFIER: Interactive Feature Activation & Inference Simulator
    // =========================================================================
    initClassifierSimulator();

    // =========================================================================
    // PROJECT LUMIERE: Interactive IoT Streetlight & Telemetry Simulator
    // =========================================================================
    initLumiereSimulator();

    // =========================================================================
    // HONORS & AWARDS: Interactive Ceremony Gallery & Plaque Inscription
    // =========================================================================
    initAwardsGallery();

    // =========================================================================
    // SECTION TABS: Honors, Certificates & Workshops
    // =========================================================================
    initSectionTabs();

    // =========================================================================
    // PROJECTS SUB-TABS: Solo, Team & Contract Projects
    // =========================================================================
    initProjectsTabs();
    initHobbyGallery();
});

function initClassifierSimulator() {
    const simSelector = document.getElementById('ic-sample-selector');
    const modelToggle = document.getElementById('ic-model-toggle');
    if (!simSelector || !modelToggle) return;

    const sampleSvgs = {
        cat1: `<rect width="96" height="96" fill="#0b1329"/>
               <path d="M28 28 L40 44 L16 44 Z" fill="#38bdf8" opacity="0.85"/>
               <path d="M68 28 L80 44 L56 44 Z" fill="#38bdf8" opacity="0.85"/>
               <circle cx="48" cy="56" r="28" fill="#1e293b" stroke="#38bdf8" stroke-width="2"/>
               <ellipse cx="38" cy="52" rx="4" ry="6" fill="#00f2fe"/>
               <ellipse cx="58" cy="52" rx="4" ry="6" fill="#00f2fe"/>
               <polygon points="48,60 45,64 51,64" fill="#f43f5e"/>
               <path d="M48 64 L48 68 M45 68 Q48 71 51 68" stroke="#94a3b8" stroke-width="1.5" fill="none"/>
               <line x1="26" y1="62" x2="10" y2="60" stroke="#06b6d4" stroke-width="1.2"/>
               <line x1="26" y1="66" x2="10" y2="67" stroke="#06b6d4" stroke-width="1.2"/>
               <line x1="70" y1="62" x2="86" y2="60" stroke="#06b6d4" stroke-width="1.2"/>
               <line x1="70" y1="66" x2="86" y2="67" stroke="#06b6d4" stroke-width="1.2"/>`,
        dog1: `<rect width="96" height="96" fill="#0b1329"/>
               <path d="M18 36 Q12 56 24 64 Q30 52 26 36 Z" fill="#a855f7" opacity="0.85"/>
               <path d="M78 36 Q84 56 72 64 Q66 52 70 36 Z" fill="#a855f7" opacity="0.85"/>
               <circle cx="48" cy="50" r="26" fill="#1e293b" stroke="#a855f7" stroke-width="2"/>
               <ellipse cx="38" cy="46" rx="4.5" ry="5.5" fill="#c084fc"/>
               <ellipse cx="58" cy="46" rx="4.5" ry="5.5" fill="#c084fc"/>
               <ellipse cx="48" cy="62" rx="14" ry="11" fill="#0f172a" stroke="#a855f7" stroke-width="1.5"/>
               <ellipse cx="48" cy="58" rx="6" ry="4" fill="#06b6d4"/>
               <path d="M48 62 L48 67 Q48 70 52 69" stroke="#94a3b8" stroke-width="1.5" fill="none"/>`,
        cat2: `<rect width="96" height="96" fill="#0b1329"/>
               <circle cx="48" cy="52" r="30" fill="#1e293b" stroke="#06b6d4" stroke-width="2"/>
               <path d="M26 32 L38 46 L18 46 Z" fill="#38bdf8" opacity="0.7"/>
               <path d="M70 32 L82 46 L62 46 Z" fill="#38bdf8" opacity="0.7"/>
               <path d="M32 54 Q38 60 44 54" stroke="#38bdf8" stroke-width="2" fill="none"/>
               <path d="M52 54 Q58 60 64 54" stroke="#38bdf8" stroke-width="2" fill="none"/>
               <polygon points="48,60 45,63 51,63" fill="#f43f5e"/>
               <path d="M48 63 Q48 67 44 66 M48 63 Q48 67 52 66" stroke="#94a3b8" stroke-width="1.2" fill="none"/>`,
        dog2: `<rect width="96" height="96" fill="#062e24"/>
               <path d="M22 28 L34 44 L16 48 Z" fill="#10b981" opacity="0.85"/>
               <path d="M74 28 L86 44 L68 48 Z" fill="#10b981" opacity="0.85"/>
               <circle cx="48" cy="52" r="27" fill="#0f172a" stroke="#10b981" stroke-width="2"/>
               <ellipse cx="38" cy="48" rx="4" ry="5" fill="#34d399"/>
               <ellipse cx="58" cy="48" rx="4" ry="5" fill="#34d399"/>
               <ellipse cx="48" cy="64" rx="13" ry="10" fill="#022c22" stroke="#10b981" stroke-width="1.5"/>
               <ellipse cx="48" cy="61" rx="5.5" ry="3.5" fill="#6ee7b7"/>
               <path d="M48 64 L48 70" stroke="#94a3b8" stroke-width="1.5"/>`
    };

    let activeSample = 'cat1';
    let activeCheckpoint = 'epoch40';

    const vectorSvg = document.getElementById('sim-vector-svg');
    const verdictCard = document.getElementById('sim-verdict-card');
    const verdictIcon = document.getElementById('sim-verdict-icon');
    const verdictName = document.getElementById('sim-verdict-name');
    const verdictConf = document.getElementById('sim-verdict-conf');
    const catBar = document.getElementById('sim-cat-bar');
    const catVal = document.getElementById('sim-cat-val');
    const dogBar = document.getElementById('sim-dog-bar');
    const dogVal = document.getElementById('sim-dog-val');
    const latencyVal = document.getElementById('sim-latency');
    const activeCkptVal = document.getElementById('sim-active-checkpoint');
    const lossVal = document.getElementById('sim-loss-val');
    const previewFrame = document.getElementById('sim-preview-frame');

    function updateSim() {
        const sampleBtn = simSelector.querySelector(`[data-sample="${activeSample}"]`);
        if (!sampleBtn) return;

        const isCat = sampleBtn.dataset.label === 'Cat';
        const confKey = activeCheckpoint === 'epoch40' ? 'conf40' : 'conf3';
        const winnerConf = parseFloat(sampleBtn.dataset[confKey]);
        const otherConf = (100 - winnerConf).toFixed(1);

        const catProb = isCat ? winnerConf.toFixed(1) : otherConf;
        const dogProb = isCat ? otherConf : winnerConf.toFixed(1);

        // Vector art injection
        if (vectorSvg && sampleSvgs[activeSample]) {
            vectorSvg.innerHTML = sampleSvgs[activeSample];
        }

        // Pulse frame & feature kernels
        if (previewFrame) {
            previewFrame.classList.remove('pulse-scan');
            void previewFrame.offsetWidth;
            previewFrame.classList.add('pulse-scan');
        }

        document.querySelectorAll('.feat-kernel').forEach(k => {
            k.style.opacity = '0.3';
            setTimeout(() => { k.style.opacity = '1'; }, 120);
        });

        // Verdict output
        if (verdictIcon) verdictIcon.textContent = isCat ? '🐱' : '🐶';
        if (verdictName) verdictName.textContent = isCat ? 'CAT' : 'DOG';
        if (verdictConf) verdictConf.textContent = `Confidence: ${winnerConf.toFixed(1)}%`;

        if (catBar) catBar.style.width = `${catProb}%`;
        if (catVal) catVal.textContent = `${catProb}%`;
        if (dogBar) dogBar.style.width = `${dogProb}%`;
        if (dogVal) dogVal.textContent = `${dogProb}%`;

        if (activeCkptVal) {
            activeCkptVal.textContent = activeCheckpoint === 'epoch40' ? '40-Epochs (24.1 MB)' : '3-Epochs (24.1 MB)';
        }

        if (lossVal) {
            const baseLoss = parseFloat(sampleBtn.dataset.loss || '0.082');
            lossVal.textContent = activeCheckpoint === 'epoch40' ? baseLoss.toFixed(3) : (baseLoss * 2.85).toFixed(3);
        }

        if (latencyVal) {
            const lat = (10.5 + Math.random() * 2.4).toFixed(1);
            latencyVal.textContent = `${lat} ms (M1 / Intel)`;
        }
    }

    // Attach sample listeners
    simSelector.querySelectorAll('.ic-sample-btn').forEach(btn => {
        btn.addEventListener('click', () => {
            simSelector.querySelectorAll('.ic-sample-btn').forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            activeSample = btn.dataset.sample;
            updateSim();
        });
    });

    // Attach model toggle listeners
    modelToggle.querySelectorAll('.ic-toggle-btn').forEach(btn => {
        btn.addEventListener('click', () => {
            modelToggle.querySelectorAll('.ic-toggle-btn').forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            activeCheckpoint = btn.dataset.checkpoint;
            updateSim();
        });
    });

    // Initial render
    updateSim();
}

// =========================================================================
// PROJECT LUMIERE: Interactive Live IoT Streetlight & Telemetry Simulator
// =========================================================================
function initLumiereSimulator() {
    const presetSelector = document.getElementById('lum-preset-selector');
    const adcSlider = document.getElementById('lum-adc-slider');
    const btnInterrupt = document.getElementById('btn-admin-interrupt');
    const btnPollHttp = document.getElementById('btn-poll-http');
    if (!presetSelector || !adcSlider) return;

    let currentAdc = 3420;
    let isSilenced = false;
    let activePreset = 'night-on';

    const lightCone = document.getElementById('svg-light-cone');
    const bulbGlow = document.getElementById('svg-bulb-glow');
    const bulbCore = document.getElementById('svg-bulb-core');
    const skyBg = document.getElementById('svg-sky-bg');
    const statusLed = document.getElementById('svg-status-led');
    const buzzerWaves = document.getElementById('svg-buzzer-waves');

    const hudLampState = document.getElementById('lum-hud-lamp-state');
    const hudBuzzerState = document.getElementById('lum-hud-buzzer-state');
    const hudAdcDisplay = document.getElementById('hud-adc-display');
    const hudAdcBar = document.getElementById('hud-adc-bar');
    const hudLuxDisplay = document.getElementById('hud-lux-display');
    const hudAdcVerdict = document.getElementById('hud-adc-verdict');
    const hudHttpResponse = document.getElementById('hud-http-response');
    const lblInterruptBtn = document.getElementById('lbl-interrupt-btn');

    function updateSimulation() {
        const isOutage = currentAdc < 3000;
        const lux = Math.max(0.1, (Math.pow(currentAdc / 4095, 2.5) * 1000)).toFixed(1);

        // Update Slider & Number
        adcSlider.value = currentAdc;
        if (hudAdcDisplay) hudAdcDisplay.textContent = `${currentAdc} / 4095 ADC`;
        if (hudLuxDisplay) hudLuxDisplay.textContent = `~${lux} Lux`;

        // Update progress bar
        const pct = ((currentAdc / 4095) * 100).toFixed(1);
        if (hudAdcBar) {
            hudAdcBar.style.width = `${pct}%`;
            if (isOutage) {
                hudAdcBar.style.background = 'linear-gradient(90deg, #ef4444, #f97316)';
            } else if (currentAdc < 3300) {
                hudAdcBar.style.background = 'linear-gradient(90deg, #f59e0b, #eab308)';
            } else {
                hudAdcBar.style.background = 'linear-gradient(90deg, #10b981, #06b6d4)';
            }
        }

        // SVG Visual states
        if (activePreset === 'daylight') {
            if (skyBg) skyBg.setAttribute('fill', '#0284c7');
            if (lightCone) lightCone.setAttribute('opacity', '0.05');
            if (bulbGlow) bulbGlow.setAttribute('fill', '#94a3b8');
            if (bulbCore) bulbCore.setAttribute('fill', '#cbd5e1');
            if (hudLampState) {
                hudLampState.textContent = 'DAYLIGHT AMBIENT (LAMP STANDBY / OFF)';
                hudLampState.className = 'lum-state-day';
            }
        } else if (activePreset === 'twilight') {
            if (skyBg) skyBg.setAttribute('fill', '#451a03');
            if (lightCone) lightCone.setAttribute('opacity', '0.45');
            if (bulbGlow) bulbGlow.setAttribute('fill', '#f59e0b');
            if (bulbCore) bulbCore.setAttribute('fill', '#fef08a');
            if (hudLampState) {
                hudLampState.textContent = 'DUSK TRANSITION (WARMING UP)';
                hudLampState.className = 'lum-state-twilight';
            }
        } else if (isOutage) {
            if (skyBg) skyBg.setAttribute('fill', '#02040a');
            if (lightCone) lightCone.setAttribute('opacity', '0');
            if (bulbGlow) bulbGlow.setAttribute('fill', '#334155');
            if (bulbCore) bulbCore.setAttribute('fill', '#1e293b');
            if (hudLampState) {
                hudLampState.textContent = 'CRITICAL OUTAGE (LAMP BULB DEFECT!)';
                hudLampState.className = 'lum-state-alert';
            }
        } else {
            // Normal night
            if (skyBg) skyBg.setAttribute('fill', '#060913');
            if (lightCone) lightCone.setAttribute('opacity', '0.9');
            if (bulbGlow) bulbGlow.setAttribute('fill', '#fef08a');
            if (bulbCore) bulbCore.setAttribute('fill', '#ffffff');
            if (hudLampState) {
                hudLampState.textContent = 'NORMAL NIGHT ILLUMINATION (LAMP ON)';
                hudLampState.className = 'lum-state-normal';
            }
        }

        // Status LED on device box
        if (isOutage) {
            if (statusLed) statusLed.setAttribute('fill', '#ef4444');
            if (hudAdcVerdict) {
                hudAdcVerdict.textContent = 'LAMP OUTAGE DETECTED';
                hudAdcVerdict.className = 'hud-status-badge badge-alert';
            }
        } else {
            if (statusLed) statusLed.setAttribute('fill', '#10b981');
            if (hudAdcVerdict) {
                hudAdcVerdict.textContent = 'NOMINAL LIGHT LEVEL';
                hudAdcVerdict.className = 'hud-status-badge badge-normal';
            }
        }

        // Buzzer state
        const buzzerActive = isOutage && !isSilenced;
        if (buzzerActive) {
            if (buzzerWaves) buzzerWaves.setAttribute('opacity', '1');
            if (hudBuzzerState) {
                hudBuzzerState.textContent = 'ALARM ACTIVE (ESCALATING 300ms BEEP)';
                hudBuzzerState.className = 'lum-state-alert-beep';
            }
        } else {
            if (buzzerWaves) buzzerWaves.setAttribute('opacity', '0');
            if (hudBuzzerState) {
                hudBuzzerState.textContent = isSilenced ? 'SILENCED (ADMIN OVERRIDE ACTIVE)' : 'SILENT (INACTIVE)';
                hudBuzzerState.className = 'lum-state-silent';
            }
        }

        // HTTP payload
        const statusStr = isOutage ? 'OUTAGE' : 'OK';
        if (hudHttpResponse) {
            hudHttpResponse.textContent = `HTTP/1.1 200 OK
Content-Type: text/plain
Connection: close

${currentAdc} | STATUS: ${statusStr} | [coordinates](https://maps.google.com/?q=37.871500,32.493200)`;
        }
    }

    // Attach preset buttons
    presetSelector.querySelectorAll('.lum-preset-btn').forEach(btn => {
        btn.addEventListener('click', () => {
            presetSelector.querySelectorAll('.lum-preset-btn').forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            activePreset = btn.dataset.preset;
            currentAdc = parseInt(btn.dataset.adc, 10);
            isSilenced = false;
            if (lblInterruptBtn) lblInterruptBtn.textContent = 'Press Admin Silence Override (GPIO 21)';
            btnInterrupt.classList.remove('active-override');
            updateSimulation();
        });
    });

    // Slider listener
    adcSlider.addEventListener('input', (e) => {
        currentAdc = parseInt(e.target.value, 10);
        presetSelector.querySelectorAll('.lum-preset-btn').forEach(b => b.classList.remove('active'));
        if (currentAdc < 3000) {
            activePreset = 'night-outage';
        } else if (currentAdc > 3800) {
            activePreset = 'daylight';
        } else {
            activePreset = 'night-on';
        }
        updateSimulation();
    });

    // Admin Interrupt Silence Button
    if (btnInterrupt) {
        btnInterrupt.addEventListener('click', () => {
            isSilenced = !isSilenced;
            if (isSilenced) {
                btnInterrupt.classList.add('active-override');
                if (lblInterruptBtn) lblInterruptBtn.textContent = '✓ Admin Silenced (Active Override)';
            } else {
                btnInterrupt.classList.remove('active-override');
                if (lblInterruptBtn) lblInterruptBtn.textContent = 'Press Admin Silence Override (GPIO 21)';
            }
            updateSimulation();
        });
    }

    // Poll HTTP Server button
    if (btnPollHttp) {
        btnPollHttp.addEventListener('click', () => {
            if (hudHttpResponse) {
                hudHttpResponse.style.opacity = '0.3';
                setTimeout(() => {
                    hudHttpResponse.style.opacity = '1';
                }, 180);
            }
        });
    }

    updateSimulation();
}

// =========================================================================
// HONORS & AWARDS: Interactive Ceremony Gallery & Plaque Inscription
// =========================================================================
function initAwardsGallery() {
    const mainImg = document.getElementById('awardMainImg');
    const thumbBtns = document.querySelectorAll('.award-thumb-btn');
    const zoomBtn = document.getElementById('awardImgZoomBtn');
    const modal = document.getElementById('awardModal');
    const modalImg = document.getElementById('awardModalImg');
    const modalCaption = document.getElementById('awardModalCaption');
    const modalCloseBtn = document.getElementById('awardModalCloseBtn');

    // Tab buttons for inscription
    const tabTr = document.getElementById('tabInscrTr');
    const tabEn = document.getElementById('tabInscrEn');
    const inscrTr = document.getElementById('inscrTextTr');
    const inscrEn = document.getElementById('inscrTextEn');

    if (tabTr && tabEn && inscrTr && inscrEn) {
        tabTr.addEventListener('click', () => {
            tabTr.classList.add('active');
            tabEn.classList.remove('active');
            inscrTr.classList.add('active');
            inscrEn.classList.remove('active');
        });

        tabEn.addEventListener('click', () => {
            tabEn.classList.add('active');
            tabTr.classList.remove('active');
            inscrEn.classList.add('active');
            inscrTr.classList.remove('active');
        });
    }

    if (thumbBtns.length > 0 && mainImg) {
        thumbBtns.forEach(btn => {
            btn.addEventListener('click', () => {
                thumbBtns.forEach(b => b.classList.remove('active'));
                btn.classList.add('active');
                const src = btn.getAttribute('data-src');
                const caption = btn.getAttribute('data-caption') || '';
                mainImg.style.opacity = '0.3';
                setTimeout(() => {
                    mainImg.src = src;
                    mainImg.style.opacity = '1';
                }, 150);
            });
        });
    }

    function openModal() {
        if (!modal || !modalImg) return;
        const activeThumb = document.querySelector('.award-thumb-btn.active');
        const src = activeThumb ? activeThumb.getAttribute('data-src') : (mainImg ? mainImg.src : '');
        const caption = activeThumb ? activeThumb.getAttribute('data-caption') : '';
        modalImg.src = src;
        if (modalCaption) modalCaption.textContent = caption;
        modal.classList.add('open');
        document.body.style.overflow = 'hidden';
    }

    function closeModal() {
        if (!modal) return;
        modal.classList.remove('open');
        document.body.style.overflow = '';
    }

    if (mainImg && mainImg.parentElement) {
        mainImg.parentElement.addEventListener('click', openModal);
    }
    if (zoomBtn) {
        zoomBtn.addEventListener('click', (e) => {
            e.stopPropagation();
            openModal();
        });
    }
    if (modalCloseBtn) {
        modalCloseBtn.addEventListener('click', closeModal);
    }
    if (modal) {
        modal.addEventListener('click', (e) => {
            if (e.target === modal) closeModal();
        });
    }
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape' && modal && modal.classList.contains('open')) {
            closeModal();
        }
    });
}

// =========================================================================
// PROJECTS SUB-TABS: Solo, Team & Contract Projects
// =========================================================================
function initProjectsTabs() {
    const tabBtns = document.querySelectorAll('#projectsTabs .section-tab-btn');
    const tabPanes = document.querySelectorAll('.projects-tab-pane');
    if (!tabBtns.length || !tabPanes.length) return;

    function switchProjectTab(tabId) {
        if (!tabId) return;
        tabBtns.forEach(btn => {
            if (btn.getAttribute('data-tab') === tabId) {
                btn.classList.add('active');
            } else {
                btn.classList.remove('active');
            }
        });
        tabPanes.forEach(pane => {
            if (pane.id === tabId) {
                pane.classList.add('active');
            } else {
                pane.classList.remove('active');
            }
        });

        // Ensure Instagram embedded media renders properly when tab becomes visible
        if (tabId === 'tab-contract-projects' && window.instgrm && window.instgrm.Widgets) {
            try {
                window.instgrm.Widgets.load();
            } catch (err) {
                console.warn('Instagram widget reload:', err);
            }
        }
    }

    // Expose globally for header stats and direct interactions
    window.switchProjectsTab = switchProjectTab;
    window.switchProjectTabAndScroll = function(tabId) {
        switchProjectTab(tabId);
        const target = document.getElementById('projectsTabs') || document.getElementById('projects');
        if (target) {
            target.scrollIntoView({ behavior: 'smooth', block: 'start' });
        }
    };

    tabBtns.forEach(btn => {
        btn.addEventListener('click', () => {
            const tabId = btn.getAttribute('data-tab');
            switchProjectTab(tabId);
        });
    });
}

// =========================================================================
// SECTION TABS: Honors, Certificates & Workshops
// =========================================================================
function initSectionTabs() {
    const tabBtns = document.querySelectorAll('#awardsTabs .section-tab-btn');
    const tabPanes = document.querySelectorAll('#awards .tab-pane');

    function switchTab(tabId) {
        if (!tabId) return;
        tabBtns.forEach(btn => {
            if (btn.getAttribute('data-tab') === tabId) {
                btn.classList.add('active');
            } else {
                btn.classList.remove('active');
            }
        });
        tabPanes.forEach(pane => {
            if (pane.id === tabId) {
                pane.classList.add('active');
            } else {
                pane.classList.remove('active');
            }
        });
    }

    // Expose globally for header stats and direct interactions
    window.switchAwardsTab = switchTab;
    window.switchAwardsTabAndScroll = function(tabId) {
        if (tabId === 'tab-education') {
            const el = document.getElementById('education');
            if (el) { el.scrollIntoView({ behavior: 'smooth', block: 'start' }); return; }
        } else if (tabId === 'tab-activities') {
            const el = document.getElementById('workshops') || document.getElementById('activities');
            if (el) { el.scrollIntoView({ behavior: 'smooth', block: 'start' }); return; }
        } else if (tabId === 'tab-certificates') {
            const el = document.getElementById('certificates');
            if (el) { el.scrollIntoView({ behavior: 'smooth', block: 'start' }); return; }
        } else if (tabId === 'tab-languages') {
            const el = document.getElementById('languages');
            if (el) { el.scrollIntoView({ behavior: 'smooth', block: 'start' }); return; }
        } else if (tabId === 'tab-awards') {
            const el = document.getElementById('awards');
            if (el) { el.scrollIntoView({ behavior: 'smooth', block: 'start' }); return; }
        }
        switchTab(tabId);
        const target = document.getElementById('awardsTabs') || document.getElementById('awards');
        if (target) {
            target.scrollIntoView({ behavior: 'smooth', block: 'start' });
        }
    };

    tabBtns.forEach(btn => {
        btn.addEventListener('click', () => {
            const tabId = btn.getAttribute('data-tab');
            switchTab(tabId);
            if (history.pushState) {
                if (tabId === 'tab-education') history.pushState(null, null, '#education');
                else if (tabId === 'tab-awards') history.pushState(null, null, '#awards');
                else if (tabId === 'tab-certificates') history.pushState(null, null, '#certificates');
                else if (tabId === 'tab-activities') history.pushState(null, null, '#workshops');
                else if (tabId === 'tab-languages') history.pushState(null, null, '#languages');
            }
        });
    });

    // Check hash on load and on hashchange
    function handleHash() {
        const hash = window.location.hash;
        if (!hash) return;
        let targetId = hash.replace('#', '');
        if (targetId === 'activities') targetId = 'workshops';
        const target = document.getElementById(targetId);
        if (target) {
            target.scrollIntoView({ behavior: 'smooth', block: 'start' });
        }
    }

    window.addEventListener('hashchange', handleHash);
    if (window.location.hash) {
        handleHash();
    }

    // Activity sub-thumbnails switcher (e.g. Teknosel photos)
    const teknoselThumbs = document.querySelectorAll('.activity-sub-thumb-btn');
    const teknoselMainImg = document.getElementById('teknoselMainImg');
    if (teknoselThumbs.length > 0 && teknoselMainImg) {
        teknoselThumbs.forEach(btn => {
            btn.addEventListener('click', () => {
                teknoselThumbs.forEach(b => b.classList.remove('active'));
                btn.classList.add('active');
                const src = btn.getAttribute('data-src');
                teknoselMainImg.style.opacity = '0.3';
                setTimeout(() => {
                    teknoselMainImg.src = src;
                    teknoselMainImg.style.opacity = '1';
                }, 150);
            });
        });
    }

    // SuPublic contract project gallery thumbnails switcher
    const supublicThumbs = document.querySelectorAll('.supublic-thumb-btn');
    const supublicMainImg = document.getElementById('supublicMainMediaImg');
    if (supublicThumbs.length > 0 && supublicMainImg) {
        supublicThumbs.forEach(btn => {
            btn.addEventListener('click', () => {
                supublicThumbs.forEach(b => b.classList.remove('active'));
                btn.classList.add('active');
                const src = btn.getAttribute('data-src');
                const alt = btn.getAttribute('data-alt');
                supublicMainImg.style.opacity = '0.3';
                setTimeout(() => {
                    supublicMainImg.src = src;
                    if (alt) supublicMainImg.alt = alt;
                    supublicMainImg.style.opacity = '1';
                }, 150);
            });
        });
    }

    // AI Armaments missile contract project gallery thumbnails switcher
    const missileThumbs = document.querySelectorAll('.missile-thumb-btn');
    const missileMainImg = document.getElementById('missileCardMainMediaImg');
    if (missileThumbs.length > 0 && missileMainImg) {
        missileThumbs.forEach(btn => {
            btn.addEventListener('click', () => {
                missileThumbs.forEach(b => b.classList.remove('active'));
                btn.classList.add('active');
                const src = btn.getAttribute('data-src');
                const alt = btn.getAttribute('data-alt');
                missileMainImg.style.opacity = '0.3';
                setTimeout(() => {
                    missileMainImg.src = src;
                    if (alt) missileMainImg.alt = alt;
                    missileMainImg.style.opacity = '1';
                }, 150);
            });
        });
    }

    // 3D Shaped Trash Bins contract project gallery thumbnails switcher
    const binThumbs = document.querySelectorAll('.bin-thumb-btn');
    const binMainImg = document.getElementById('binCardMainMediaImg');
    if (binThumbs.length > 0 && binMainImg) {
        binThumbs.forEach(btn => {
            btn.addEventListener('click', () => {
                binThumbs.forEach(b => b.classList.remove('active'));
                btn.classList.add('active');
                const src = btn.getAttribute('data-src');
                const alt = btn.getAttribute('data-alt');
                binMainImg.style.opacity = '0.3';
                setTimeout(() => {
                    binMainImg.src = src;
                    if (alt) binMainImg.alt = alt;
                    binMainImg.style.opacity = '1';
                }, 150);
            });
        });
    }

    // Aquarium Filter contract project gallery thumbnails switcher
    const aquariumThumbs = document.querySelectorAll('.aquarium-thumb-btn');
    const aquariumMainImg = document.getElementById('aquariumCardMainMediaImg');
    if (aquariumThumbs.length > 0 && aquariumMainImg) {
        aquariumThumbs.forEach(btn => {
            btn.addEventListener('click', () => {
                aquariumThumbs.forEach(b => b.classList.remove('active'));
                btn.classList.add('active');
                const src = btn.getAttribute('data-src');
                const alt = btn.getAttribute('data-alt');
                aquariumMainImg.style.opacity = '0.3';
                setTimeout(() => {
                    aquariumMainImg.src = src;
                    if (alt) aquariumMainImg.alt = alt;
                    aquariumMainImg.style.opacity = '1';
                }, 150);
            });
        });
    }
}

// =========================================================================
// HOBBY PROJECTS: SolidWorks CAD Interactive Gallery & Modal Zoom
// =========================================================================
function openHobbyModal(src, caption) {
    const modal = document.getElementById('awardModal');
    const modalImg = document.getElementById('awardModalImg');
    const modalCaption = document.getElementById('awardModalCaption');
    if (!modal || !modalImg) return;
    modalImg.src = src;
    if (modalCaption) modalCaption.textContent = caption || '';
    modal.classList.add('open');
    document.body.style.overflow = 'hidden';
}
window.openHobbyModal = openHobbyModal;

function initHobbyGallery() {
    const thumbBtns = document.querySelectorAll('.hobby-thumb-btn');
    if (!thumbBtns.length) return;

    thumbBtns.forEach(btn => {
        btn.addEventListener('click', (e) => {
            e.stopPropagation();
            const targetImgId = btn.getAttribute('data-target');
            const newSrc = btn.getAttribute('data-src');
            const caption = btn.getAttribute('data-caption') || '';
            const parentBar = btn.closest('.hobby-gallery-bar');

            if (parentBar) {
                parentBar.querySelectorAll('.hobby-thumb-btn').forEach(b => b.classList.remove('active'));
            }
            btn.classList.add('active');

            const targetImg = document.getElementById(targetImgId);
            if (targetImg && newSrc) {
                targetImg.style.opacity = '0.3';
                setTimeout(() => {
                    targetImg.src = newSrc;
                    if (caption) targetImg.alt = caption;
                    targetImg.style.opacity = '1';
                }, 140);
            }

            // Sync active state in expanded grid if present
            const card = btn.closest('.project-card-hobby');
            if (card && newSrc) {
                const gridItems = card.querySelectorAll('.hobby-grid-item');
                gridItems.forEach(item => {
                    const itemImg = item.querySelector('img');
                    if (itemImg && itemImg.getAttribute('src') === newSrc) {
                        item.classList.add('active');
                    } else {
                        item.classList.remove('active');
                    }
                });
            }
        });
    });
}

// Toggle expansion of the hobby project full image showcase drawer
function toggleHobbyProject(drawerId, btn) {
    const drawer = document.getElementById(drawerId);
    if (!drawer) return;

    const isOpen = drawer.classList.contains('open');
    const labelSpan = btn ? btn.querySelector('.btn-hobby-label') : null;
    const viewText = btn ? (btn.getAttribute('data-view-text') || 'View Project') : 'View Project';
    const hideText = btn ? (btn.getAttribute('data-hide-text') || 'Hide CAD Views') : 'Hide CAD Views';

    if (isOpen) {
        drawer.classList.remove('open');
        if (btn) {
            btn.classList.remove('is-expanded');
            btn.setAttribute('aria-expanded', 'false');
            if (labelSpan) labelSpan.textContent = viewText;
        }
    } else {
        drawer.classList.add('open');
        if (btn) {
            btn.classList.add('is-expanded');
            btn.setAttribute('aria-expanded', 'true');
            if (labelSpan) labelSpan.textContent = hideText;
        }

        // Smooth scroll to ensure drawer visibility
        setTimeout(() => {
            drawer.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
        }, 120);
    }
}
window.toggleHobbyProject = toggleHobbyProject;

// Close the hobby project image showcase drawer
function closeHobbyProject(drawerId) {
    const drawer = document.getElementById(drawerId);
    if (!drawer) return;
    drawer.classList.remove('open');

    const card = drawer.closest('.project-card');
    if (card) {
        const btn = card.querySelector('.btn-hobby-expand');
        if (btn) {
            btn.classList.remove('is-expanded');
            btn.setAttribute('aria-expanded', 'false');
            const labelSpan = btn.querySelector('.btn-hobby-label');
            const viewText = btn.getAttribute('data-view-text') || 'View Project';
            if (labelSpan) labelSpan.textContent = viewText;
        }
    }
}
window.closeHobbyProject = closeHobbyProject;

// Select an image from the expanded drawer: update main card image & open lightbox modal
function selectHobbyImage(targetImgId, src, caption, element) {
    const targetImg = document.getElementById(targetImgId);
    if (targetImg && src) {
        targetImg.style.opacity = '0.3';
        setTimeout(() => {
            targetImg.src = src;
            if (caption) targetImg.alt = caption;
            targetImg.style.opacity = '1';
        }, 120);
    }

    const card = element ? element.closest('.project-card') : null;
    if (card) {
        // Sync active state on top gallery buttons
        const topBtns = card.querySelectorAll('.hobby-thumb-btn');
        topBtns.forEach(b => {
            if (b.getAttribute('data-src') === src) {
                b.classList.add('active');
            } else {
                b.classList.remove('active');
            }
        });

        // Sync active state in expanded grid
        const gridItems = card.querySelectorAll('.hobby-grid-item');
        gridItems.forEach(item => item.classList.remove('active'));
        if (element && element.classList.contains('hobby-grid-item')) {
            element.classList.add('active');
        }
    }

    // Open lightbox modal for high-resolution inspection
    openHobbyModal(src, caption);
}
window.selectHobbyImage = selectHobbyImage;


