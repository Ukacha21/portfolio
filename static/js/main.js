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
        sendit: {
            title: "SENDIT — Autonomous Airport Baggage & Post Logistics",
            domain: "Autonomous Robotics · Logistics Automation",
            status: "In Queue (Development Phase)",
            tags: ["AGV Fleet", "LiDAR SLAM", "Terminal Routing", "ROS 2"],
            desc: "Autonomous guided vehicle (AGV) mechanism tailored for apron and terminal transfer of airmail and time-critical transfer baggage.",
            specs: "Payload capacity: 250 kg · Navigation: 2D LiDAR + AprilTag visual inertial odometry · Multi-agent fleet dispatcher · 12-hour continuous runtime on LiFePO4 pack.",
            link: "#"
        },
        smartwaste: {
            title: "SMARTWASTE — Autonomous AI Waste Segregator",
            domain: "Computer Vision · Mechatronic Actuation",
            status: "In Queue (Development Phase)",
            tags: ["Edge AI", "Near-IR Spectrometry", "Pneumatic Diverter", "Recycling"],
            desc: "Mechatronic conveyor sorting mechanism integrating near-infrared (NIR) spectrometry and high-speed computer vision to distinguish PET, HDPE, aluminum, and compostable organics.",
            specs: "Sorting throughput: 120 items/minute · Detection accuracy: 96.2% · High-speed pneumatic air jets · Edge TPU inference latency < 15 ms.",
            link: "#"
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
