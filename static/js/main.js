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
            status: "Complete Showcase (James Dyson Award 2026 Nominee)",
            tags: ["AI Vision Classification", "QR Contact Auto-Fill", "Zebra Thermal Labels", "SumUp Contactless", "Zero Airport Capex"],
            desc: "Smart self-service kiosk at airport and train station security checkpoints that packages, classifies, labels, weighs, and dispatches travellers' surrendered items home in under 2 minutes.",
            specs: "Interaction time: <120 seconds · SBC: Raspberry Pi 4/5 + Coral TPU · Zebra ZD421 203dpi thermal printer · Integrated 12V nichrome heat-sealer · SumUp Pro contactless reader.",
            link: "/project/sendit"
        },
        smartwaste: {
            title: "SMARTWASTE — Autonomous Waste Sorting & 3:1 Compactor",
            domain: "Autonomous Mechatronics · Smart Municipal IoT",
            status: "Complete Showcase (James Dyson Award '26 Nominee)",
            tags: ["ToF Laser Telemetry", "Bimodal Sorting Gate", "3:1 Lead-Screw Press", "Deep Sleep ESP32", "Konya Selçuklu Deployment"],
            desc: "Retrofittable IoT bin module combining millimetre-accurate VCSEL laser ToF fill-level telemetry, bimodal inductive/optical sorting gates, and an automated 1.2 kN lead-screw 3:1 compactor to slash municipal collection mileage at under $80 BOM.",
            specs: "Compaction ratio: 3:1 · BOM cost: < $80 · Sensor: VL53L1X Laser ToF (±3mm) · Battery autonomy: 180+ days (15 µA deep sleep) · Municipal route reduction: 34%.",
            link: "/project/smartwaste"
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
});
