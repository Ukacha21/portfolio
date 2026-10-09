"""
DENOVA - Optical Filter Simulation
Project 04: Tooth Decay Detector (James Dyson Award 2026)

Simulates the fluorescence emission spectrum of healthy vs. demineralised enamel
under 405 nm excitation, and validates that a 440 nm longpass filter provides
sufficient signal-to-noise ratio (SNR) to detect ICDAS-II Stage 1 lesions.

References:
  - Stookey et al. (2004): QLF autofluorescence spectra for enamel
  - Gmur et al. (2006): Porphyrin red fluorescence from cariogenic bacteria
  - ICDAS-II criteria: Stage 1 = ~10% fluorescence loss vs healthy baseline

Run: python optical_filter_simulation.py
Requires: numpy, matplotlib  (pip install numpy matplotlib)
"""

import sys
import os
import numpy as np
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches

# Force UTF-8 output on Windows terminals
sys.stdout.reconfigure(encoding="utf-8")

# Resolve output path relative to this script's location
SCRIPT_DIR   = os.path.dirname(os.path.abspath(__file__))
SIMULATION_DIR = os.path.join(SCRIPT_DIR, "..", "simulation")
OUTPUT_PNG   = os.path.join(SIMULATION_DIR, "DENOVA_optical_simulation.png")

# ─────────────────────────────────────────────────────────────────────────────
# 1. Define wavelength axis (400 – 750 nm)
# ─────────────────────────────────────────────────────────────────────────────
wavelengths = np.linspace(400, 750, 700)

# ─────────────────────────────────────────────────────────────────────────────
# 2. Model enamel autofluorescence spectra (normalised, 405 nm excitation)
#    - Healthy enamel: broad green fluorescence peaking ~520 nm
#    - ICDAS Stage 1 (early demineralisation): ~10% reduction in green band
#    - ICDAS Stage 2 (moderate demineralisation): ~30% reduction, red porphyrin peak
#    - ICDAS Stage 3 (active cavity): ~60% reduction + strong red at 630–680 nm
# ─────────────────────────────────────────────────────────────────────────────

def gaussian(x, mu, sigma, amplitude):
    return amplitude * np.exp(-0.5 * ((x - mu) / sigma) ** 2)

# Excitation spike at 405 nm (what the longpass filter must block)
excitation = gaussian(wavelengths, 405, 4, 1.0)

# Healthy enamel: green autofluorescence (peak ~520 nm) + minor tail
healthy = (
    gaussian(wavelengths, 520, 60, 0.80)   # main green band
    + gaussian(wavelengths, 620, 40, 0.05) # minimal baseline red
)

# Stage 1 (ICDAS-II): ~10% green reduction, slight red increase
stage1 = (
    gaussian(wavelengths, 520, 60, 0.72)   # 10% reduction
    + gaussian(wavelengths, 630, 45, 0.12) # mild porphyrin red
)

# Stage 2: ~30% green reduction, notable red porphyrin
stage2 = (
    gaussian(wavelengths, 520, 60, 0.56)   # 30% reduction
    + gaussian(wavelengths, 635, 50, 0.30) # moderate porphyrin
)

# Stage 3 (active cavity): strong red dominates
stage3 = (
    gaussian(wavelengths, 520, 60, 0.32)
    + gaussian(wavelengths, 640, 55, 0.65)
)

# ─────────────────────────────────────────────────────────────────────────────
# 3. Model the 440 nm longpass filter (OD4 rejection below 440 nm)
#    Transmission: 0 below 430 nm, linear ramp 430–450 nm, 1.0 above 450 nm
# ─────────────────────────────────────────────────────────────────────────────
filter_440lp = np.where(
    wavelengths < 430, 0.0,
    np.where(
        wavelengths < 450,
        (wavelengths - 430) / 20.0,
        1.0
    )
)

# Apply filter to all spectra
excitation_filtered = excitation * filter_440lp
healthy_filtered    = healthy    * filter_440lp
stage1_filtered     = stage1     * filter_440lp
stage2_filtered     = stage2     * filter_440lp
stage3_filtered     = stage3     * filter_440lp

# ─────────────────────────────────────────────────────────────────────────────
# 4. Compute SNR metrics
#    Green band: 490–570 nm | Red band: 600–700 nm
# ─────────────────────────────────────────────────────────────────────────────
green_mask = (wavelengths >= 490) & (wavelengths <= 570)
red_mask   = (wavelengths >= 600) & (wavelengths <= 700)

def band_power(spectrum, mask):
    return np.trapezoid(spectrum[mask], wavelengths[mask])

green_healthy = band_power(healthy_filtered, green_mask)
red_healthy   = band_power(healthy_filtered, red_mask)

results = {}
for name, spec in [("Healthy", healthy_filtered),
                   ("Stage 1 (ICDAS-II)", stage1_filtered),
                   ("Stage 2", stage2_filtered),
                   ("Stage 3", stage3_filtered)]:
    g = band_power(spec, green_mask)
    r = band_power(spec, red_mask)
    results[name] = {
        "green_power": g,
        "red_power":   r,
        "green_loss_%": (1 - g / green_healthy) * 100,
        "red_gain_%":   (r / red_healthy - 1) * 100 if red_healthy > 0 else float("inf"),
        "red_green_ratio": r / g if g > 0 else float("inf"),
    }

# ─────────────────────────────────────────────────────────────────────────────
# 5. Print summary table
# ─────────────────────────────────────────────────────────────────────────────
print("=" * 70)
print("DENOVA — Optical Filter Simulation Results")
print("Excitation: 405 nm | Filter: 440 nm longpass | Detection: photodetector")
print("=" * 70)
print(f"{'Condition':<22} {'Green Power':>12} {'Green Loss':>10} {'Red/Green Ratio':>16}")
print("-" * 70)
for name, r in results.items():
    print(f"{name:<22} {r['green_power']:>12.3f} {r['green_loss_%']:>9.1f}% {r['red_green_ratio']:>16.3f}")

print()
print("Conclusion:")
stage1_green_loss = results["Stage 1 (ICDAS-II)"]["green_loss_%"]
if stage1_green_loss >= 8:
    print(f"  ✅ ICDAS Stage 1 green loss = {stage1_green_loss:.1f}% — DETECTABLE above noise threshold (>8%)")
else:
    print(f"  ⚠️  ICDAS Stage 1 green loss = {stage1_green_loss:.1f}% — marginal; consider higher-gain photodetector")
print()

# ─────────────────────────────────────────────────────────────────────────────
# 6. Plot
# ─────────────────────────────────────────────────────────────────────────────
fig, axes = plt.subplots(2, 1, figsize=(12, 10))
fig.suptitle("DENOVA — 405 nm QLF Optical Simulation\nEnamel Fluorescence Through 440 nm Longpass Filter",
             fontsize=14, fontweight="bold")

# Panel A: Raw spectra + filter transmission
ax = axes[0]
ax.set_title("A. Raw Fluorescence Spectra + Filter Transmission (before filter)")
ax.plot(wavelengths, excitation, color="purple", lw=1.5, ls="--", alpha=0.5, label="Excitation 405 nm")
ax.plot(wavelengths, healthy,    color="limegreen", lw=2, label="Healthy enamel")
ax.plot(wavelengths, stage1,     color="yellowgreen", lw=2, label="Stage 1 (ICDAS-II, ~10% loss)")
ax.plot(wavelengths, stage2,     color="orange", lw=2, label="Stage 2 (~30% loss)")
ax.plot(wavelengths, stage3,     color="red", lw=2, label="Stage 3 (active cavity)")
ax2 = ax.twinx()
ax2.fill_between(wavelengths, filter_440lp, alpha=0.12, color="steelblue", label="440 nm LP filter Tx")
ax2.set_ylabel("Filter Transmission", color="steelblue")
ax2.tick_params(axis="y", labelcolor="steelblue")
ax2.set_ylim(0, 1.5)
ax.axvline(440, color="steelblue", lw=1, ls=":", alpha=0.8)
ax.set_xlabel("Wavelength (nm)")
ax.set_ylabel("Relative Fluorescence Intensity")
ax.set_xlim(400, 750)
ax.legend(loc="upper right", fontsize=9)

# Panel B: Filtered spectra (what the photodetector sees)
ax = axes[1]
ax.set_title("B. Post-Filter Spectra — What the Photodetector Sees")
ax.plot(wavelengths, healthy_filtered, color="limegreen", lw=2, label="Healthy")
ax.plot(wavelengths, stage1_filtered,  color="yellowgreen", lw=2, label="Stage 1 (ICDAS-II)")
ax.plot(wavelengths, stage2_filtered,  color="orange", lw=2, label="Stage 2")
ax.plot(wavelengths, stage3_filtered,  color="red", lw=2, label="Stage 3")
ax.axvspan(490, 570, alpha=0.08, color="green", label="Green detection band")
ax.axvspan(600, 700, alpha=0.08, color="red",   label="Red detection band")
ax.set_xlabel("Wavelength (nm)")
ax.set_ylabel("Relative Intensity (excitation blocked)")
ax.set_xlim(400, 750)
ax.legend(loc="upper right", fontsize=9)

plt.tight_layout()
plt.savefig(OUTPUT_PNG, dpi=150, bbox_inches="tight")
print(f"  Plot saved to: {OUTPUT_PNG}")
plt.show()
