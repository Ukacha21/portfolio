"""
DENOVA - 3D Product Visualization
Project 04: Tooth Decay Detector (James Dyson Award 2026)

Renders a 3D mesh model of the DENOVA handheld screening pen using
matplotlib surface plots. Fixed camera angle — no editor camera,
no interactive controls required. Saves a high-resolution PNG.

Components shown:
  - ABS barrel (main body)
  - Silicone shielding cap (conical tip, teal)
  - AS7341 sensor window (tip face, dark)
  - Traffic-light LED ring (red / amber / green)
  - USB-C port slot
  - Top end cap

Run: python denova_product_visualization.py
Requires: numpy, matplotlib  (pip install numpy matplotlib)
"""

import sys
import os
import numpy as np
import matplotlib
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D
from mpl_toolkits.mplot3d.art3d import Poly3DCollection
import matplotlib.patches as mpatches

# ── Output path ───────────────────────────────────────────────────────────────
sys.stdout.reconfigure(encoding="utf-8")
SCRIPT_DIR     = os.path.dirname(os.path.abspath(__file__))
SIMULATION_DIR = os.path.join(SCRIPT_DIR, "..", "simulation")
OUTPUT_PNG     = os.path.join(SIMULATION_DIR, "DENOVA_product_visualization.png")

# ── Colour palette ────────────────────────────────────────────────────────────
C_BARREL    = "#2d3037"   # dark ABS body
C_BARREL_LT = "#3d4047"   # lighter highlight face
C_GRIP      = "#23262d"   # darker grip band
C_CAP_TIP   = "#00b4a0"   # teal silicone cap
C_CAP_FACE  = "#00c8b4"   # brighter cap face
C_SENSOR    = "#111827"   # near-black sensor window
C_LED_G     = "#32de50"   # green LED
C_LED_A     = "#ffa500"   # amber LED
C_LED_R     = "#de3232"   # red LED
C_USB       = "#111418"   # USB-C slot
C_TOP_CAP   = "#3a3f4a"   # top end cap
BG_DARK     = "#0a0c14"   # figure background

# ── Geometry helpers ──────────────────────────────────────────────────────────

def cylinder_mesh(r, z_bottom, z_top, n=64):
    """Return (X, Y, Z) surface arrays for a cylinder along Z-axis."""
    theta = np.linspace(0, 2 * np.pi, n)
    z     = np.array([z_bottom, z_top])
    T, Z  = np.meshgrid(theta, z)
    X = r * np.cos(T)
    Y = r * np.sin(T)
    return X, Y, Z

def cone_mesh(r_base, r_tip, z_bottom, z_top, n=64):
    """Return (X, Y, Z) surface arrays for a truncated cone along Z-axis."""
    theta  = np.linspace(0, 2 * np.pi, n)
    z_vals = np.array([z_bottom, z_top])
    r_vals = np.array([r_base,   r_tip])
    T, _   = np.meshgrid(theta, z_vals)        # (2, n)
    _, Zg  = np.meshgrid(theta, z_vals)        # (2, n)  — Z values
    # Radius varies linearly with z row
    R = np.column_stack([r_vals] * n)          # (2, n)
    X = R * np.cos(T)
    Y = R * np.sin(T)
    return X, Y, Zg

def disk_mesh(r, z_pos, n=64):
    """Return (X, Y, Z) surface arrays for a filled disk at z_pos."""
    rr     = np.linspace(0, r, 8)
    theta  = np.linspace(0, 2 * np.pi, n)
    R, T   = np.meshgrid(rr, theta)
    X = R * np.cos(T)
    Y = R * np.sin(T)
    Z = np.full_like(X, z_pos)
    return X, Y, Z

def sphere_cap_mesh(r, z_center, z_sign=1, n=32):
    """Return (X, Y, Z) surface arrays for a hemisphere."""
    phi   = np.linspace(0, np.pi / 2, n // 2)
    theta = np.linspace(0, 2 * np.pi, n)
    P, T  = np.meshgrid(phi, theta)
    X = r * np.sin(P) * np.cos(T)
    Y = r * np.sin(P) * np.sin(T)
    Z = z_center + z_sign * r * np.cos(P)
    return X, Y, Z

def small_sphere(cx, cy, cz, r=0.012, n=20):
    """Return X, Y, Z for a small sphere — used for LEDs."""
    u = np.linspace(0, 2 * np.pi, n)
    v = np.linspace(0, np.pi, n)
    X = cx + r * np.outer(np.cos(u), np.sin(v))
    Y = cy + r * np.outer(np.sin(u), np.sin(v))
    Z = cz + r * np.outer(np.ones(n), np.cos(v))
    return X, Y, Z

# ── Figure setup ──────────────────────────────────────────────────────────────
fig = plt.figure(figsize=(16, 10), facecolor=BG_DARK)
fig.patch.set_facecolor(BG_DARK)

# Main 3D axes — fixed camera (elev, azim chosen for best product view)
ax = fig.add_axes([0.0, 0.0, 0.70, 1.0], projection="3d")
ax.set_facecolor(BG_DARK)
ax.xaxis.pane.fill = False
ax.yaxis.pane.fill = False
ax.zaxis.pane.fill = False
ax.xaxis.pane.set_edgecolor("none")
ax.yaxis.pane.set_edgecolor("none")
ax.zaxis.pane.set_edgecolor("none")
ax.set_axis_off()

# ── FIXED CAMERA — no EditorCamera, no mouse interaction ─────────────────────
ax.view_init(elev=18, azim=-55)   # locked presentation angle

# ── Pen dimensions (metres-scale, displayed vertically along Z) ───────────────
R_BARREL  = 0.095    # barrel radius (19mm dia → 9.5mm r, scaled ×10)
R_GRIP    = 0.105    # grip band radius (slightly wider)
Z_TOTAL   = 1.00     # total pen length
Z_TIP     = -0.50    # bottom of barrel (tip end)
Z_TOP     = 0.50     # top of barrel

# ── Main barrel ───────────────────────────────────────────────────────────────
X, Y, Z = cylinder_mesh(R_BARREL, Z_TIP + 0.18, Z_TOP - 0.05)
ax.plot_surface(X, Y, Z, color=C_BARREL, alpha=1.0, linewidth=0, antialiased=True, zorder=2)

# Top disk cap face
Xd, Yd, Zd = disk_mesh(R_BARREL, Z_TOP - 0.05)
ax.plot_surface(Xd, Yd, Zd, color=C_TOP_CAP, alpha=1.0, linewidth=0, antialiased=True, zorder=3)

# Hemisphere top cap
Xs, Ys, Zs = sphere_cap_mesh(R_BARREL, Z_TOP - 0.05, z_sign=1)
ax.plot_surface(Xs, Ys, Zs, color=C_TOP_CAP, alpha=1.0, linewidth=0, antialiased=True, zorder=3)

# ── Grip band (slightly wider, recessed texture colour) ───────────────────────
Xg, Yg, Zg = cylinder_mesh(R_GRIP, 0.10, 0.32)
ax.plot_surface(Xg, Yg, Zg, color=C_GRIP, alpha=1.0, linewidth=0, antialiased=True, zorder=2)

# ── Silicone shielding cap (teal cone, tip end) ───────────────────────────────
Xc, Yc, Zc = cone_mesh(R_BARREL, 0.025, Z_TIP + 0.18, Z_TIP)
ax.plot_surface(Xc, Yc, Zc, color=C_CAP_TIP, alpha=1.0, linewidth=0, antialiased=True, zorder=2)

# Cap tip face (small disk at the very tip)
Xtf, Ytf, Ztf = disk_mesh(0.025, Z_TIP)
ax.plot_surface(Xtf, Ytf, Ztf, color=C_CAP_FACE, alpha=1.0, linewidth=0, antialiased=True, zorder=3)

# ── AS7341 sensor window (dark inset disk just above shielding cap) ───────────
Xsw, Ysw, Zsw = disk_mesh(0.055, Z_TIP + 0.19)
ax.plot_surface(Xsw, Ysw, Zsw, color=C_SENSOR, alpha=1.0, linewidth=0, antialiased=True, zorder=4)

# ── LED indicator ring — three small spheres (G / A / R) ─────────────────────
led_z     = 0.18
led_r_pos = R_BARREL + 0.005    # just proud of the surface
led_colors = [(C_LED_G, 0), (C_LED_A, np.pi * 0.75), (C_LED_R, np.pi * 1.5)]
for col, angle in led_colors:
    lx = led_r_pos * np.cos(angle - np.pi * 0.25)
    ly = led_r_pos * np.sin(angle - np.pi * 0.25)
    Xl, Yl, Zl = small_sphere(lx, ly, led_z, r=0.014)
    ax.plot_surface(Xl, Yl, Zl, color=col, alpha=1.0, linewidth=0, antialiased=True, zorder=5)

# ── USB-C port (small dark rectangular slot on barrel side) ──────────────────
usb_theta = np.linspace(-0.12, 0.12, 4)
usb_z     = np.array([0.38, 0.42])
UT, UZ    = np.meshgrid(usb_theta, usb_z)
UX = (R_BARREL - 0.005) * np.cos(UT + np.pi * 1.1)
UY = (R_BARREL - 0.005) * np.sin(UT + np.pi * 1.1)
ax.plot_surface(UX, UY, UZ, color=C_USB, alpha=1.0, linewidth=0, zorder=6)

# ── Axis limits & final fixed view ───────────────────────────────────────────
ax.set_xlim(-0.4, 0.4)
ax.set_ylim(-0.4, 0.4)
ax.set_zlim(-0.65, 0.75)

# ── Annotation panel (right side, 2D text axes) ───────────────────────────────
ax2 = fig.add_axes([0.68, 0.0, 0.32, 1.0])
ax2.set_facecolor(BG_DARK)
ax2.set_axis_off()

def txt(x, y, s, size=11, color="#c8d0e0", bold=False, **kw):
    weight = "bold" if bold else "normal"
    ax2.text(x, y, s, transform=ax2.transAxes,
             fontsize=size, color=color, fontweight=weight,
             verticalalignment="top", **kw)

txt(0.05, 0.97, "DENOVA",          size=32, color="#00c8b4", bold=True)
txt(0.05, 0.90, "Handheld Early Caries Detector", size=13, color="#8090b0")
txt(0.05, 0.86, "James Dyson Award 2026",          size=11, color="#5060a0")

txt(0.05, 0.74, "─── COMPONENTS ───",              size=10, color="#405070")

specs = [
    ("#00c8b4", "Detection",  "AS7341 11-channel spectral sensor"),
    ("#00c8b4", "",           "Integrated Si interference filters"),
    ("#00c8b4", "",           "515 nm (green) + 630 nm (red) channels"),
    ("#8090b0", "Excitation", "405 nm violet LED  (QLF)"),
    ("#8090b0", "",           "850 nm NIR LED  (transillumination)"),
    ("#a0b8d8", "MCU",        "ESP32-S3 Zero  |  240 MHz dual-core"),
    ("#a0b8d8", "",           "Wi-Fi 802.11 b/g/n  +  BLE 5.0"),
    ("#c8a060", "Power",      "18650 Li-ion  |  USB-C charge (TP4056)"),
    ("#c8a060", "",           "8+ hours continuous use"),
    ("#32de50", "Output",     "Green / Amber / Red LED indicator"),
    ("#32de50", "",           "BLE packet  →  paired smartphone"),
    ("#d0d8e8", "Housing",    "ABS barrel  |  19 mm dia  |  200 mm"),
    ("#d0d8e8", "",           "Silicone shielding cap  (teal tip)"),
]

y_pos = 0.68
for col, label, value in specs:
    if label:
        txt(0.05, y_pos, f"{label:<13}", size=10, color="#607090", bold=True)
        txt(0.35, y_pos, value, size=10, color=col)
    else:
        txt(0.35, y_pos, value, size=10, color=col)
    y_pos -= 0.040

txt(0.05, 0.10, "─── COST TARGET ───",             size=10, color="#405070")
txt(0.05, 0.06, "BOM  £15 – 30  per unit",         size=13, color="#00c8b4", bold=True)
txt(0.05, 0.02, "100× cheaper than clinical equiv.", size=10, color="#8090b0")

# Legend for LED colours
legend_items = [
    mpatches.Patch(color=C_LED_G,   label="Green  — Healthy enamel"),
    mpatches.Patch(color=C_LED_A,   label="Amber  — Early lesion (ICDAS Stage 1–2)"),
    mpatches.Patch(color=C_LED_R,   label="Red    — Active decay (Stage 3)"),
    mpatches.Patch(color=C_CAP_TIP, label="Teal   — Silicone shielding cap"),
]
leg = ax2.legend(handles=legend_items, loc="lower left",
                 bbox_to_anchor=(0.03, 0.19),
                 framealpha=0.15, labelcolor="#c8d0e0",
                 facecolor="#1a1e2a", edgecolor="#405070", fontsize=9)

# ── Save (fixed output — no plt.show() blocking) ─────────────────────────────
plt.savefig(OUTPUT_PNG, dpi=180, bbox_inches="tight",
            facecolor=BG_DARK, edgecolor="none")
print(f"  DENOVA 3D product visualization saved to:")
print(f"  {OUTPUT_PNG}")

# Show interactively as well (comment out if running headlessly)
plt.show()
