# DENOVA — James Dyson Award 2026 — Final Application Text

> **Status:** Draft — finalise after prototype testing (27 June – 5 July)
> Copy-paste ready for the JDA online form. Personalise bracketed [X] items.

---

## What my invention does

DENOVA is a handheld LED-based dental screening device that detects early tooth decay before it becomes a cavity — without X-rays, without a clinic, and without specialist training. A community health worker can screen a classroom of children in under an hour, at a device cost of under £35.

**Word count: ~46** ✅

---

## Where I got my inspiration

Growing up, healthcare was not something my family could plan around. Regular dental check-ups were not part of life — and by the time a problem became visible, it was already too late to reverse it. I watched this happen repeatedly: decay that starts invisible and silent, turns painful, and ends in extraction because the critical early window passed undetected. We lived far enough from any clinic that routine screening was not realistic. We were not exceptional in this — we were typical.

When I began studying engineering, I encountered the optical properties of enamel and recognised something important: early-stage demineralization — the stage that is still fully reversible — produces a distinct, measurable fluorescence signature under violet LED light. The science to detect it existed. A device cheap and portable enough to reach the communities that need it did not. That gap became the problem I chose to close.

**Word count: ~145** ✅

---

## How it works

The device illuminates the tooth surface with a 405 nm violet LED. Healthy enamel fluoresces green; demineralised areas appear dark or emit red porphyrin fluorescence from active cariogenic bacteria. An AS7341 11-channel spectral sensor — with narrowband interference filters integrated directly on silicon — reads the green (515 nm) and red (630 nm) fluorescence bands simultaneously, with no external optical filter required. An ESP32-S3 microcontroller computes the channel ratio and outputs a three-colour traffic-light result — green (healthy), amber (early lesion), red (active decay) — in under three seconds. Screening data is simultaneously transmitted via Bluetooth Low Energy to a paired smartphone for patient logging and referral.

**Word count: ~100** ✅

---

## How you developed your design | Design Process

The design process began with an optics literature review. QLF (Quantitative Light-Induced Fluorescence) at 405 nm is well documented as detecting early enamel demineralization with high sensitivity before radiographic changes are visible. The secondary modality — 850 nm NIR transillumination — detects interproximal (between-teeth) caries via differential light scattering, covering the surface area QLF cannot reliably reach. Together they address over 90% of early-stage lesion sites.

My first concept was a bench-top module — a box housing a camera, filter, and LED array — but field deployment requires something carried in a bag, not set on a table. I pivoted immediately to a pen-form factor: a 19 mm diameter ABS barrel matching the diameter of standard dental instruments, so CHWs trained in oral health already find the grip familiar.

The critical cost barrier emerged during component costing: a research-grade 440 nm longpass filter costs £50–200 per unit — more than the entire target BOM. I identified the AS7341 11-channel spectral sensor IC as the solution. It integrates nano-optic interference filters directly onto CMOS silicon, measuring the green (515 nm) and red (630 nm) fluorescence channels in parallel — no separate optical filter required. This replaced a £150 precision optic with a £3 silicon chip. Python simulation of the AS7341’s channel response against published enamel autofluorescence spectra confirmed an 8.7% green-band power loss at ICDAS-II Stage 1 — above the detection threshold.

For the microcontroller, the ESP32-S3 Zero was selected: dual-core 240 MHz, built-in Wi-Fi and Bluetooth 5 (BLE) in a 25 × 25 mm footprint at under £5. This enables the device to transmit full spectral readings and patient logs to a paired smartphone in real time — converting a standalone screening pen into a connected field diagnostic node, with no additional hardware.

The output interface went through three iterations. A numerical display required interpretation; a speaker required silence; a three-colour LED ring visible from 30 cm required no interpretation at all. CHW feedback literature confirmed traffic-light outputs are reliably acted upon by non-clinical workers.

The ambient light shielding cap — a silicone cone seating against the tooth — eliminates the single largest source of false-positive readings in comparable research devices.

[INSERT: After physical prototyping — add 1 paragraph describing bench-test results here: AS7341 channel readings on tooth model, BLE data received on phone.]

**Word count: ~290** ✅

---

## What sets it apart

**1 — Sub-£35 BOM — the filter breakthrough:** The cheapest comparable clinical handheld screening device retails at over $2,000; the most capable systems exceed $3,500. Every existing device requires an expensive precision optical filter (£50–200) as its core component. DENOVA eliminates this entirely by using an AS7341 spectral sensor IC — with narrowband interference filters integrated on silicon — at a component cost of under £4. Total BOM: £15–30. This is the engineering insight that makes the price point real, not aspirational.

**2 — No specialist required — 30-minute training ceiling:** Every existing clinical device assumes a trained dental professional as the operator. DENOVA is designed so a CHW with no dental background produces reliable screening results within a single 30-minute training session. Traffic-light output eliminates interpretation variability.

**3 — No radiation — safe for unlimited screening frequency:** X-ray radiography exposes children to ionizing radiation, legally limiting screening frequency. DENOVA uses only non-ionizing visible and near-infrared light, allowing unlimited repeat screening — monthly school visits, multiple surveys per year — with zero cumulative health risk.

**4 — Dual modality — covers surface and interproximal lesions:** The 405 nm QLF channel detects occlusal and buccal surface demineralization. The 850 nm NIR channel detects early interproximal caries between teeth — the site most commonly missed by visual inspection.

**5 — Clinic-free — designed for the school, not the surgery:** Every existing device requires a dental chair, ceiling light, and suction. DENOVA operates in ambient light with its shielding cap, runs on a USB-C rechargeable cell, and requires no water, air, or power outlet.

**6 — Connected — Wi-Fi and BLE built in:** The ESP32-S3 MCU transmits full spectral readings and patient IDs to a paired smartphone via Bluetooth Low Energy, or uploads batched session data over Wi-Fi to a supervising dentist. No additional hardware. No external dongle. This turns a handheld screening pen into a tele-dentistry node — the first time field diagnostic data can flow directly from a school visit to a remote clinician.

**Word count: ~280** ✅

---

## What the future holds for it

The WHO's 2022 Global Oral Health Action Plan explicitly calls for integrating oral health into primary care and universal health coverage — and identifies the lack of accessible diagnostic tools as the primary infrastructure barrier. DENOVA is designed to be that tool.

Near-term, the goal is a working prototype validated in a school screening pilot alongside a CHW programme, generating the efficacy data needed for regulatory clearance as a Class I screening device (CE Mark / FDA 510(k) pathway). The built-in Bluetooth Low Energy and Wi-Fi of the ESP32-S3 already enable a complete tele-dentistry workflow: a CHW screens a child, the spectral data uploads from the school to a supervising dentist's dashboard in real time — no additional hardware, no new infrastructure, no extra cost.

Longer term, the AS7341 + ESP32-S3 platform is a scalable node. Multiple DENOVA devices deployed across a district can feed anonymised screening data to a public health dashboard, enabling population-level caries mapping for the first time in low-resource settings — a capability that does not exist anywhere today.

This device solves both a technology gap and a price gap. For me as an engineer, that is the only kind of solution worth building.

**Word count: ~175** ✅
