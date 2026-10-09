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
            title: "DENOVA — Non-Invasive Optical Tooth Decay Detector",
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
            tags: ["ToF Laser Telemetry", "Bimodal Sorting Gate", "3:1 Lead-Screw Press", "Deep Sleep ESP32", "Konya Selçuklu Deployment"],
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
