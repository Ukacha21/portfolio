// VELUM — Bioimpedance Mat & AD5933 Frequency Sweep Interactive Simulator
// Ukacha Juma Mechatronics Engineering

(function() {
    // DOM Elements
    const btnNormal = document.getElementById('simBtnNormal');
    const btnIschemia = document.getElementById('simBtnIschemia');
    const btnOedema = document.getElementById('simBtnOedema');
    const btnReset = document.getElementById('simBtnReset');

    const telemFreq = document.getElementById('telemFreq');
    const telemZ = document.getElementById('telemZ');
    const telemPhase = document.getElementById('telemPhase');
    const telemDelta = document.getElementById('telemDelta');

    const alertBanner = document.getElementById('simAlertBanner');
    const alertIcon = document.getElementById('simAlertIcon');
    const alertText = document.getElementById('simAlertText');

    const sacrumNode = document.getElementById('node-sacrum');
    const sacrumVal = document.getElementById('sacrumVal');
    const canvas = document.getElementById('sweepCanvas');

    if (!sacrumNode || !canvas) return;

    // Simulation State
    let currentScenario = 'normal'; // normal, ischemia, oedema, reset
    let currentFrequency = 50.0;
    let targetZ = 1450;
    let currentZ = 1450;
    let baselineZ = 1450;
    let sweepPhase = 0;

    const scenarios = {
        normal: {
            z: 1450,
            phase: -14.2,
            delta: 0.0,
            statusText: "NOMINAL: Tissue perfusion optimal across all 16 anatomical zones",
            nodeClass: "",
            sacrumDisplay: "1.45 kΩ",
            isAlert: false,
            isWarn: false
        },
        ischemia: {
            z: 1330,
            phase: -17.5,
            delta: -8.3,
            statusText: "CAUTION: Sustained capillary occlusion detected in sacral zone",
            nodeClass: "node-warn",
            sacrumDisplay: "1.33 kΩ",
            isAlert: false,
            isWarn: true
        },
        oedema: {
            z: 1180,
            phase: -24.8,
            delta: -18.6,
            statusText: "CRITICAL: Sub-Epidermal Oedema Detected (ΔZ ≥ 15%)! Reposition Sacrum Immediately",
            nodeClass: "node-alert",
            sacrumDisplay: "1.18 kΩ",
            isAlert: true,
            isWarn: false
        },
        reset: {
            z: 1440,
            phase: -14.6,
            delta: -0.7,
            statusText: "RECOVERY: 30° Lateral turn confirmed · Capillary re-perfusion restored",
            nodeClass: "",
            sacrumDisplay: "1.44 kΩ",
            isAlert: false,
            isWarn: false
        }
    };

    function setScenario(scKey) {
        currentScenario = scKey;
        const sc = scenarios[scKey];
        targetZ = sc.z;

        // Button active class
        [btnNormal, btnIschemia, btnOedema, btnReset].forEach(b => {
            if (b) b.classList.remove('active');
        });

        if (scKey === 'normal' && btnNormal) btnNormal.classList.add('active');
        if (scKey === 'ischemia' && btnIschemia) btnIschemia.classList.add('active');
        if (scKey === 'oedema' && btnOedema) btnOedema.classList.add('active');
        if (scKey === 'reset' && btnReset) btnReset.classList.add('active');

        // Update Sacrum Node State
        sacrumNode.classList.remove('node-warn', 'node-alert');
        if (sc.nodeClass) {
            sacrumNode.classList.add(sc.nodeClass);
        }
        if (sacrumVal) sacrumVal.textContent = sc.sacrumDisplay;

        // Update Alert Banner
        if (alertBanner) {
            alertBanner.classList.remove('status-normal', 'status-alert');
            if (sc.isAlert) {
                alertBanner.classList.add('status-alert');
                if (alertIcon) alertIcon.textContent = "⚠";
            } else {
                alertBanner.classList.add('status-normal');
                if (alertIcon) alertIcon.textContent = "✓";
            }
            if (alertText) alertText.textContent = sc.statusText;
        }

        // Telemetry values
        if (telemPhase) telemPhase.textContent = sc.phase.toFixed(1) + "°";
        if (telemDelta) {
            telemDelta.textContent = (sc.delta > 0 ? "+" : "") + sc.delta.toFixed(1) + "%";
            telemDelta.style.color = sc.isAlert ? "#ff6b6b" : (sc.isWarn ? "#f4c542" : "var(--cyan-accent)");
        }

        // Trigger Audio Chime on Alert (subtle Web Audio beep)
        if (sc.isAlert) {
            playAlertChime();
        }
    }

    // Audio chime using Web Audio API
    function playAlertChime() {
        try {
            const AudioContext = window.AudioContext || window.webkitAudioContext;
            if (!AudioContext) return;
            const ctx = new AudioContext();
            const osc = ctx.createOscillator();
            const gain = ctx.createGain();

            osc.type = 'sine';
            osc.frequency.setValueAtTime(880, ctx.currentTime); // A5 note
            osc.frequency.exponentialRampToValueAtTime(440, ctx.currentTime + 0.3);

            gain.gain.setValueAtTime(0.08, ctx.currentTime);
            gain.gain.exponentialRampToValueAtTime(0.001, ctx.currentTime + 0.35);

            osc.connect(gain);
            gain.connect(ctx.destination);
            osc.start();
            osc.stop(ctx.currentTime + 0.35);
        } catch (e) {
            // Audio context policy blocked or unsupported
        }
    }

    // Button event listeners
    if (btnNormal) btnNormal.addEventListener('click', () => setScenario('normal'));
    if (btnIschemia) btnIschemia.addEventListener('click', () => setScenario('ischemia'));
    if (btnOedema) btnOedema.addEventListener('click', () => setScenario('oedema'));
    if (btnReset) btnReset.addEventListener('click', () => setScenario('reset'));

    // Allow clicking on any sensor node in the bed
    document.querySelectorAll('.sensor-node').forEach(node => {
        node.addEventListener('click', () => {
            const zoneName = node.dataset.zone || "Zone";
            const valElem = node.querySelector('.node-val');
            const valStr = valElem ? valElem.textContent : "1.40 kΩ";
            if (telemZ) telemZ.textContent = valStr;
            if (telemDelta) telemDelta.textContent = "0.0%";
            if (telemDelta) telemDelta.style.color = "var(--cyan-accent)";
        });
    });

    // Canvas Frequency Sweep Animation
    const ctx = canvas.getContext('2d');
    let animationFrameId;

    function drawSweep() {
        // Smoothly interpolate currentZ to targetZ
        currentZ += (targetZ - currentZ) * 0.08;
        if (telemZ) {
            telemZ.textContent = Math.round(currentZ) + " Ω";
        }

        // Sweep animation parameter
        sweepPhase = (sweepPhase + 1.2) % canvas.width;
        currentFrequency = 1.0 + (sweepPhase / canvas.width) * 99.0;
        if (telemFreq) {
            telemFreq.textContent = currentFrequency.toFixed(1) + " kHz";
        }

        const w = canvas.width;
        const h = canvas.height;

        ctx.clearRect(0, 0, w, h);

        // Draw Grid Lines
        ctx.strokeStyle = "rgba(0, 201, 177, 0.08)";
        ctx.lineWidth = 1;

        for (let x = 0; x < w; x += 50) {
            ctx.beginPath();
            ctx.moveTo(x, 0);
            ctx.lineTo(x, h);
            ctx.stroke();
        }
        for (let y = 0; y < h; y += 25) {
            ctx.beginPath();
            ctx.moveTo(0, y);
            ctx.lineTo(w, y);
            ctx.stroke();
        }

        // Frequency sweep response curve (Cole-Cole dispersion curve simulation)
        // High frequency => lower impedance; oedema drops the curve vertically
        ctx.beginPath();
        const baseOffset = (1500 - currentZ) * 0.12;

        for (let i = 0; i < w; i++) {
            const f = 1 + (i / w) * 99; // 1 to 100 kHz
            // Bioimpedance dispersion curve
            const y = h * 0.45 + Math.sin(i * 0.04) * 4 + (baseOffset) + Math.log(f) * 6;
            if (i === 0) {
                ctx.moveTo(i, y);
            } else {
                ctx.lineTo(i, y);
            }
        }

        // Gradient Stroke for Sweep Curve
        const grad = ctx.createLinearGradient(0, 0, w, 0);
        if (currentScenario === 'oedema') {
            grad.addColorStop(0, '#ff6b6b');
            grad.addColorStop(1, '#ff9544');
        } else if (currentScenario === 'ischemia') {
            grad.addColorStop(0, '#f4c542');
            grad.addColorStop(1, '#00c9b1');
        } else {
            grad.addColorStop(0, '#00c9b1');
            grad.addColorStop(1, '#38c6ff');
        }

        ctx.strokeStyle = grad;
        ctx.lineWidth = 2.5;
        ctx.shadowColor = currentScenario === 'oedema' ? '#ff6b6b' : '#00c9b1';
        ctx.shadowBlur = 10;
        ctx.stroke();
        ctx.shadowBlur = 0;

        // Current Sweep Laser Cursor
        ctx.beginPath();
        ctx.moveTo(sweepPhase, 0);
        ctx.lineTo(sweepPhase, h);
        ctx.strokeStyle = currentScenario === 'oedema' ? 'rgba(255, 107, 107, 0.8)' : 'rgba(56, 198, 255, 0.8)';
        ctx.lineWidth = 1.5;
        ctx.stroke();

        // Laser Cursor Intersection Glow Point
        const currentY = h * 0.45 + Math.sin(sweepPhase * 0.04) * 4 + (baseOffset) + Math.log(currentFrequency) * 6;
        ctx.beginPath();
        ctx.arc(sweepPhase, currentY, 4.5, 0, Math.PI * 2);
        ctx.fillStyle = "#ffffff";
        ctx.shadowColor = "#38c6ff";
        ctx.shadowBlur = 12;
        ctx.fill();
        ctx.shadowBlur = 0;

        animationFrameId = requestAnimationFrame(drawSweep);
    }

    drawSweep();

    // Responsive Canvas Resize
    function resizeCanvas() {
        const rect = canvas.getBoundingClientRect();
        if (rect.width > 0) {
            canvas.width = rect.width;
        }
    }

    window.addEventListener('resize', resizeCanvas);
    resizeCanvas();
})();
