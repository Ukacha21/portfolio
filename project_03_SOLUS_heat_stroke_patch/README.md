**SAFETY & SECURITY** — Heat Stroke Early Warning Patch for Outdoor Workers

---

## 🧪 Minimum Viable Prototype — What You Need to Build

> **Rule:** Dyson rejects "wholly conceptual" entries. You do NOT need a finished product — you need verifiable evidence of iterative development. Here is the minimum credible proof of work for SOLUS.

### What counts as enough:

| # | Action | What it proves | Difficulty |
|---|--------|---------------|------------|
| 1 | **Purchase thermochromic pigment or ink** (available on AliExpress/Etsy — "thermochromic pigment 37°C" or "body heat activated") and paint or print it onto a strip of medical tape or flexible film | You have a physically responding material in hand — this is the core of the invention | Low — 1–2 days shipping |
| 2 | **Trigger the colour change** using your hand, a warm water bath (37–39°C), or a heat gun at low setting — photograph the transition clearly | Visual proof that the thermochromic mechanism works at body-temperature range | Trivial once material arrives |
| 3 | **Create a layered patch mockup** — base medical adhesive backing, thermochromic indicator layer, clear protective film — photograph it on a forearm | Shows the integrated patch form factor, not just a loose material test | Low |
| 4 | **Demonstrate the three-stage colour concept** — even if using three separate material samples at different temperatures (e.g., cold water, warm water, hand contact) — photograph each stage | Proves the green → amber → red warning logic is physically realisable | Low |
| 5 | **Build the RF side separately** — a 433 MHz TX/RX module pair (HC-12 or FS1000A, ~£3 each on AliExpress) triggered by a manual button standing in for the patch circuit, received on a wristband-sized receiver with a buzzer | Proves the alert mechanism works independently; the full integration is the stated next step | Medium — Arduino/breadboard work |
| 6 | **Photograph and write up** what temperature triggered the colour change, how sharp the transition was, any calibration issues | Shows iterative testing; honesty about calibration challenges is exactly what judges want to see |

### What you do NOT need:
- A fully integrated patch with embedded RF transmitter (describe integration as the next development phase)
- A leuco dye synthesis or custom microencapsulation (commercial thermochromic materials are the starting point for all real products in this space)
- Clinical temperature accuracy (±1–2°C is sufficient for a proof-of-concept prototype)
- A conductive trace circuit inside the patch (demonstrate RF trigger with a separate button for submission)

### Minimum time to credible prototype: **3–5 days** (2 days material delivery + 1–2 days assembly and testing)

---

# Project 03 — Heat Stroke Early Warning Patch

**James Dyson Award 2026**
**Category:** Occupational Health / Wearable Engineering
**Scope:** Regional focus: Latin America, South Asia, Sub-Saharan Africa
**Deadline:** 15 July 2026

---

## The Problem

Agricultural and construction workers in hot climates suffer heat stroke silently. Core temperature thresholds are crossed before visible symptoms appear. Existing health wearables are designed for fitness consumers, not informal-sector workers.

## The Solution

A low-cost disposable skin patch using colorimetric chemistry (no battery) that visibly changes colour near dangerous thresholds. A supervisor's wristband receives a loud audible alert via short-range RF when the patch triggers.

---

## James Dyson Award Application — Draft Answers

> 📝 **Character limits** based on the official Dyson Award application form.

---

### 📛 Project Name
*(Limit: 50 characters)*

Heat Stroke Early Warning Patch

---

### ⚡ Quick Description
*(Limit: 193 characters)*

A battery-free disposable skin patch using colorimetric chemistry that visibly changes colour near dangerous body temperature thresholds, paired with a supervisor's wristband alarm.

---

### ❓ What my invention does
*(Limit: 260 characters)*

The patch monitors body temperature via a colorimetric skin indicator. When core temperature nears a dangerous threshold, the patch visibly changes colour — simultaneously triggering an audible alarm on the supervisor's wristband via short-range RF.

---

### ❓ Where I got my inspiration
*(Limit: 650 characters)*

Agricultural workers in tropical climates face life-threatening heat stroke with no warning. Core temperature thresholds are crossed in minutes, but visible symptoms appear only after collapse is imminent. Existing wearables are designed for fitness consumers — expensive, battery-dependent, and inaccessible to informal-sector workers. A solution needed to cost almost nothing, require no electricity or smartphone, and work under direct sun in high humidity. Colorimetric chemistry offered exactly that: a passive, instant, readable colour change triggered by temperature — requiring no training to interpret.

---

### ❓ How it works
*(Limit: 975 characters)*

The patch contains a temperature-sensitive colorimetric indicator layer — a microencapsulated thermochromic compound calibrated to the clinical heat stroke risk threshold (38.5°C skin-equivalent). Below threshold, the patch appears green. As temperature approaches danger, it shifts to amber, then bright red — a three-stage warning readable from several metres away.

The patch is worn on the upper arm or back of the neck, areas with reliable correlation to core temperature. A thin conductive trace on the patch surface completes a circuit only when the colour-change layer reaches red-stage temperature. This triggers a micro-RF transmitter (433 MHz, coin-cell powered) embedded in the patch edge — sending a signal to the supervisor's wristband, which emits a loud audible alert and vibration.

The patch itself contains no battery. It is a single-use disposable worn for one shift and replaced daily at a target material cost under $0.30 per unit.

---

### ❓ How you developed your design | Design Process
*(Limit: 1300 characters)*

The design began with field research into working conditions of agricultural and construction workers in high-heat climates. Three constraints emerged as non-negotiable: battery-free on the worker's body, under $0.50 per patch, and readable without training — since supervisors may be monitoring dozens of workers simultaneously.

Colorimetric thermochromic compounds — used in consumer products like temperature-sensitive packaging — were identified as the core technology. The challenge was calibration: commercial thermochromic inks typically shift at 45°C+. A lower-threshold formulation was developed through microencapsulation of a leuco dye system calibrated to a 38.5°C trigger point.

The three-stage colour system (green → amber → red) was designed for unambiguous field readability under direct sunlight. Patch geometry was optimised for upper-arm placement, balancing blood-flow proximity with visibility to supervisors at distance.

The supervisor wristband uses a passive 433 MHz RF receiver requiring no pairing or setup. Alert volume was set at 85 dB minimum to cut through machinery noise and wind. The wristband is reusable and solar-recharged, while only the worker patch is disposable — minimising ongoing cost.

---

### ❓ What sets it apart
*(Limit: 975 characters)*

No passive, battery-free, disposable heat stroke warning system currently exists for field workers. Existing solutions fall into two categories: expensive consumer wearables (Garmin, Apple Watch) requiring daily charging, smartphone pairing, and costing $200–$500; or fixed environmental monitoring stations that measure air temperature, not individual physiological state.

The Heat Stroke Patch addresses a population — informal agricultural and construction workers — that neither product reaches. It is worn and forgotten, requiring zero interaction from the worker. The colour-change is legible without any device, training, or literacy. The supervisor alert is audible and ambient — no screen to check. At under $0.50 per patch, it is viable for bulk distribution by employers and governments in high-risk regions.

---

### ❓ What the future holds for it
*(Limit: 650 characters)*

The next phase involves field trials with agricultural cooperatives in Turkey and India, validating temperature-response accuracy across real working conditions and varied skin tones. A bulk manufacturing partnership will be sought with existing medical consumable producers to achieve the $0.30/patch cost target at scale. Regulatory pathway will target CE marking as personal protective equipment rather than a medical device, avoiding lengthy approval timelines. Long-term: a hydration indicator layer will be integrated into the same patch, adding a second protective dimension to the same disposable format.

---

## Folder Structure

```
project_03_heat_stroke_patch/
├── README.md
├── research/
├── sketches/
├── cad/
├── simulation/
├── scripts/
├── prototyping/
├── application/
└── presentation/
```
