from reportlab.lib.pagesizes import A4
from reportlab.lib.styles import getSampleStyleSheet, ParagraphStyle
from reportlab.lib.units import cm
from reportlab.lib import colors
from reportlab.platypus import (
    SimpleDocTemplate, Paragraph, Spacer, Table, TableStyle,
    HRFlowable, KeepTogether
)
from reportlab.lib.enums import TA_LEFT, TA_CENTER, TA_JUSTIFY
import os

OUTPUT = os.path.join(os.path.dirname(__file__), "description.pdf")

# ── Colour palette ──────────────────────────────────────────────────────────
DARK       = colors.HexColor("#1a1a2e")
ACCENT     = colors.HexColor("#e8a020")   # warm amber – earthy / construction
LIGHT_BG   = colors.HexColor("#f5f0e8")
MID_GREY   = colors.HexColor("#6b6b6b")
WHITE      = colors.white
TABLE_ALT  = colors.HexColor("#fdf6ec")

# ── Styles ───────────────────────────────────────────────────────────────────
base = getSampleStyleSheet()

def S(name, **kw):
    return ParagraphStyle(name, **kw)

styles = {
    "cover_tag": S("cover_tag",
        fontName="Helvetica", fontSize=9, textColor=ACCENT,
        spaceAfter=4, alignment=TA_CENTER, leading=12),
    "cover_title": S("cover_title",
        fontName="Helvetica-Bold", fontSize=36, textColor=WHITE,
        spaceAfter=6, alignment=TA_CENTER, leading=42),
    "cover_sub": S("cover_sub",
        fontName="Helvetica", fontSize=13, textColor=colors.HexColor("#cccccc"),
        spaceAfter=4, alignment=TA_CENTER, leading=18),
    "cover_meta": S("cover_meta",
        fontName="Helvetica", fontSize=9, textColor=colors.HexColor("#aaaaaa"),
        spaceAfter=2, alignment=TA_CENTER, leading=13),
    "section": S("section",
        fontName="Helvetica-Bold", fontSize=13, textColor=ACCENT,
        spaceBefore=18, spaceAfter=6, leading=16),
    "subsection": S("subsection",
        fontName="Helvetica-Bold", fontSize=10, textColor=DARK,
        spaceBefore=10, spaceAfter=4, leading=13),
    "body": S("body",
        fontName="Helvetica", fontSize=9.5, textColor=DARK,
        spaceAfter=5, leading=14, alignment=TA_JUSTIFY),
    "bullet": S("bullet",
        fontName="Helvetica", fontSize=9.5, textColor=DARK,
        spaceAfter=3, leading=14, leftIndent=14, firstLineIndent=-10),
    "label": S("label",
        fontName="Helvetica-Bold", fontSize=8, textColor=MID_GREY,
        spaceAfter=1, leading=11),
    "small": S("small",
        fontName="Helvetica", fontSize=8, textColor=MID_GREY,
        spaceAfter=2, leading=11),
    "quote": S("quote",
        fontName="Helvetica-Oblique", fontSize=9.5, textColor=DARK,
        spaceAfter=5, leading=14, leftIndent=18, rightIndent=18,
        alignment=TA_JUSTIFY),
}

W, H = A4
MARGIN = 2.0 * cm

# ── Helper ────────────────────────────────────────────────────────────────────
def rule(color=ACCENT, thickness=1):
    return HRFlowable(width="100%", thickness=thickness, color=color, spaceAfter=6, spaceBefore=2)

def section(title):
    return [Spacer(1, 4), rule(), Paragraph(title.upper(), styles["section"])]

def body(text):
    return Paragraph(text, styles["body"])

def bullet(text):
    return Paragraph(f"&#8226;  {text}", styles["bullet"])

def kv_table(rows, col_widths=None):
    """Two-column label / value table."""
    if col_widths is None:
        col_widths = [5*cm, 11.5*cm]
    data = [[Paragraph(f"<b>{k}</b>", styles["small"]),
             Paragraph(str(v), styles["body"])] for k, v in rows]
    t = Table(data, colWidths=col_widths)
    t.setStyle(TableStyle([
        ("BACKGROUND", (0,0), (0,-1), LIGHT_BG),
        ("BACKGROUND", (1,0), (1,-1), WHITE),
        ("TEXTCOLOR",  (0,0), (-1,-1), DARK),
        ("ROWBACKGROUNDS", (0,0), (-1,-1), [LIGHT_BG, TABLE_ALT]),
        ("GRID",  (0,0), (-1,-1), 0.3, colors.HexColor("#d8cfc0")),
        ("TOPPADDING",    (0,0), (-1,-1), 5),
        ("BOTTOMPADDING", (0,0), (-1,-1), 5),
        ("LEFTPADDING",   (0,0), (-1,-1), 7),
    ]))
    return t

def wide_table(headers, rows, col_widths=None):
    """Multi-column data table."""
    header_row = [Paragraph(f"<b>{h}</b>", styles["small"]) for h in headers]
    body_rows  = [[Paragraph(str(c), styles["body"]) for c in r] for r in rows]
    data = [header_row] + body_rows
    t = Table(data, colWidths=col_widths)
    t.setStyle(TableStyle([
        ("BACKGROUND",    (0,0), (-1,0),  DARK),
        ("TEXTCOLOR",     (0,0), (-1,0),  WHITE),
        ("ROWBACKGROUNDS",(0,1), (-1,-1), [WHITE, TABLE_ALT]),
        ("GRID",          (0,0), (-1,-1), 0.3, colors.HexColor("#d8cfc0")),
        ("TOPPADDING",    (0,0), (-1,-1), 5),
        ("BOTTOMPADDING", (0,0), (-1,-1), 5),
        ("LEFTPADDING",   (0,0), (-1,-1), 7),
    ]))
    return t

# ── Cover page elements ───────────────────────────────────────────────────────
def cover_block():
    elems = []
    # Dark banner
    banner = Table(
        [[Paragraph("SUSTAINABILITY &amp; DISASTER RELIEF", styles["cover_tag"])],
         [Paragraph("RENOVO", styles["cover_title"])],
         [Paragraph("Rubble-to-Brick Press", styles["cover_sub"])],
         [Spacer(1, 6)],
         [Paragraph("James Dyson Award 2026  ·  Project 05", styles["cover_meta"])],
         [Paragraph("Disaster Relief / Structural Engineering", styles["cover_meta"])],
         [Paragraph("Regional focus: Turkey · Morocco · Nepal · Central America", styles["cover_meta"])],
         [Paragraph("Deadline: 15 July 2026", styles["cover_meta"])],
        ],
        colWidths=[W - 2*MARGIN]
    )
    banner.setStyle(TableStyle([
        ("BACKGROUND", (0,0), (-1,-1), DARK),
        ("TOPPADDING",    (0,0), (-1,-1), 14),
        ("BOTTOMPADDING", (0,0), (-1,-1), 14),
        ("LEFTPADDING",   (0,0), (-1,-1), 20),
        ("RIGHTPADDING",  (0,0), (-1,-1), 20),
    ]))
    elems.append(banner)
    elems.append(Spacer(1, 14))
    return elems

# ── Document build ────────────────────────────────────────────────────────────
def build():
    doc = SimpleDocTemplate(
        OUTPUT, pagesize=A4,
        leftMargin=MARGIN, rightMargin=MARGIN,
        topMargin=MARGIN, bottomMargin=MARGIN,
        title="RENOVO — Rubble-to-Brick Press | Project Brief",
        author="James Dyson Award 2026",
    )

    story = []

    # ── COVER ─────────────────────────────────────────────────────────────────
    story += cover_block()

    # ── OVERVIEW ──────────────────────────────────────────────────────────────
    story += section("1. Project Overview")
    story.append(kv_table([
        ("Project Name",    "RENOVO — Rubble-to-Brick Press"),
        ("Award",           "James Dyson Award 2026"),
        ("Category",        "Sustainability / Disaster Relief / Structural Engineering"),
        ("Tag",             "SUSTAINABILITY"),
        ("Regional Focus",  "Turkey, Morocco, Nepal, Central America"),
        ("Target Cost",     "Near-zero per brick (energy + binder only)"),
        ("Deadline",        "15 July 2026"),
        ("Status",          "🔄 In Progress"),
    ]))

    # ── THE PROBLEM ───────────────────────────────────────────────────────────
    story += section("2. The Problem")
    story.append(body(
        "After major earthquakes — Turkey 2023, Morocco 2023, Nepal 2015 — two crises unfold "
        "simultaneously. Communities sit surrounded by millions of tonnes of rubble that must be "
        "expensively cleared, while simultaneously importing new building materials at prices they "
        "cannot afford. The waste and the need exist side by side, completely disconnected."
    ))
    story.append(body(
        "Standard rubble management has two options: pay for expensive removal and landfill, or "
        "leave rubble in place for years. Some NGOs use manual rubble-sorting for road fill, but "
        "no field-deployable system exists to upcycle concrete rubble into structural bricks."
    ))
    story.append(Spacer(1, 4))
    story.append(wide_table(
        ["Earthquake", "Year", "Estimated Rubble Volume", "Reconstruction Cost"],
        [
            ["Turkey (Kahramanmaraş)", "2023", "~210 million tonnes", ">$100 billion USD"],
            ["Morocco (Al Haouz)", "2023", "~5 million tonnes", ">$10 billion USD"],
            ["Nepal (Gorkha)", "2015", "~100 million tonnes", ">$7 billion USD"],
            ["Syria (ongoing)", "2011–", ">900 million tonnes", ">$400 billion USD"],
        ],
        col_widths=[5*cm, 2*cm, 4.5*cm, 5*cm]
    ))

    # ── THE SOLUTION ──────────────────────────────────────────────────────────
    story += section("3. The Solution")
    story.append(body(
        "RENOVO is a compact, modular, field-deployable press that crushes concrete rubble, "
        "mixes it with geopolymer binder (no kiln required), and compresses it into standard-dimension "
        "structural bricks curing at ambient temperature. A full pallet of bricks is produced "
        "on-site within 48 hours, using solar power or manual operation — requiring no grid, "
        "no fuel, and no imported materials."
    ))
    story.append(Spacer(1, 6))
    for b in [
        "Transportable by small truck — modular, bolts together on-site",
        "Operable by unskilled labour — minimal training required",
        "Manufacturable from standard steel sections available in any country",
        "Dual power: 1.5 kW electric (200 W solar panel) or manual flywheel",
        "Brick dimensions: ISO 228 × 108 × 76 mm — drop-in compatible with conventional masonry",
        "Target compressive strength: 15–25 MPa (meets single-storey residential standard)",
    ]:
        story.append(bullet(b))

    # ── HOW IT WORKS ──────────────────────────────────────────────────────────
    story += section("4. How It Works — Three Stages")
    stages = [
        ["Stage 1\nCrushing", "Jaw-crusher attachment (electric or manual flywheel) reduces concrete rubble to 5–20 mm aggregate. Integral sieve removes rebar and oversized debris before material enters the mixer."],
        ["Stage 2\nMixing", "Aggregate enters a mixing chamber with geopolymer binder solution. Sodium silicate activates fly ash or GGBS (ground granulated blast-furnace slag) into a cementitious matrix. No firing required. Mix ratio: 75% crushed rubble / 25% geopolymer paste by volume."],
        ["Stage 3\nCompression", "Mixture loads into a compression mould and is pressed at 10–15 MPa via hydraulic ram — solar-battery or manual lever. Brick is ejected and stacked for 24–48 hour ambient curing. Achieves 15–25 MPa compressive strength."],
    ]
    t = Table(stages, colWidths=[3.2*cm, 13.3*cm])
    t.setStyle(TableStyle([
        ("BACKGROUND",    (0,0), (0,-1), DARK),
        ("TEXTCOLOR",     (0,0), (0,-1), ACCENT),
        ("FONTNAME",      (0,0), (0,-1), "Helvetica-Bold"),
        ("FONTSIZE",      (0,0), (0,-1), 8),
        ("ROWBACKGROUNDS",(1,0), (1,-1), [WHITE, TABLE_ALT, WHITE]),
        ("GRID",          (0,0), (-1,-1), 0.3, colors.HexColor("#d8cfc0")),
        ("TOPPADDING",    (0,0), (-1,-1), 8),
        ("BOTTOMPADDING", (0,0), (-1,-1), 8),
        ("LEFTPADDING",   (0,0), (-1,-1), 8),
        ("VALIGN",        (0,0), (-1,-1), "TOP"),
    ]))
    story.append(t)

    # ── GEOPOLYMER CHEMISTRY ──────────────────────────────────────────────────
    story += section("5. Geopolymer Binder Chemistry")
    story.append(body(
        "Portland cement was eliminated — it requires controlled curing conditions and does not "
        "bind recycled aggregate reliably under field conditions. Alkali-activated geopolymers "
        "cure at ambient temperature and achieve equivalent strength within 24 hours."
    ))
    story.append(Spacer(1, 4))
    story.append(wide_table(
        ["Component", "Role", "Source / Availability", "Cost"],
        [
            ["Sodium silicate (Na₂SiO₃)", "Alkaline activator — triggers polymerisation of alumino-silicate precursor", "Hardware shops (waterglass), industrial chemical suppliers", "~£5–15 / kg"],
            ["Fly ash (Class F)", "Primary geopolymer precursor — provides SiO₂ and Al₂O₃ for network formation", "Coal power plant byproduct — free or near-free in most regions", "~£0–5 / tonne"],
            ["GGBS (Ground Granulated Blast-Furnace Slag)", "Alternative/supplementary precursor — faster strength gain", "Steel industry byproduct — widely available globally", "~£30–80 / tonne"],
            ["Crushed concrete aggregate", "Filler / structural matrix — 75% of brick volume", "On-site rubble — zero cost", "Free"],
            ["Water", "Mix workability adjustment", "Local supply", "Free"],
        ],
        col_widths=[4*cm, 5.5*cm, 4*cm, 3*cm]
    ))
    story.append(Spacer(1, 6))
    story.append(body(
        "<b>Reaction mechanism:</b> The alkali solution (NaOH + sodium silicate) dissolves "
        "Si and Al from the precursor surface, forming Si–O–Al oligomers that condense into a "
        "3D aluminosilicate network — geopolymer gel. This gel fills voids between aggregate "
        "particles and hardens progressively at room temperature, reaching structural strength "
        "in 24–48 hours and continuing to gain strength over 28 days."
    ))

    # ── MECHANICAL ARCHITECTURE ───────────────────────────────────────────────
    story += section("6. Mechanical Architecture")
    story.append(body(
        "The press is modular — jaw crusher, mixing drum, and compression press are separate "
        "units that bolt together in the field, allowing damaged components to be replaced "
        "individually without losing the whole machine."
    ))
    story.append(wide_table(
        ["Module", "Function", "Power", "Key Spec"],
        [
            ["Jaw Crusher", "Reduces rubble to 5–20 mm aggregate", "1.5 kW electric or manual flywheel", "Throughput: ~50 kg/hr at manual operation"],
            ["Integral Sieve", "Removes rebar and oversized debris", "Passive (gravity)", "Mesh aperture: 20 mm"],
            ["Mixing Drum", "Combines aggregate + binder at 75:25 ratio", "0.5 kW electric or manual paddle", "Batch size: ~10 kg per cycle (~5 bricks)"],
            ["Hydraulic Press", "Compresses mixture at 10–15 MPa", "Solar battery bank (200 W panel) or manual lever", "Mould: ISO 228×108×76 mm"],
            ["Ejection System", "Releases cured brick without damage", "Mechanical (spring-loaded)", "Cycle time: ~2 min per brick"],
        ],
        col_widths=[3.5*cm, 4.5*cm, 4.5*cm, 4*cm]
    ))

    # ── BRICK SPECIFICATION ───────────────────────────────────────────────────
    story += section("7. Brick Specification")
    story.append(kv_table([
        ("Dimensions",              "228 × 108 × 76 mm (ISO standard — drop-in compatible with conventional masonry)"),
        ("Compressive Strength",    "15–25 MPa (published geopolymer literature at 75% recycled aggregate, 10–15 MPa press pressure)"),
        ("Standard Equivalent",     "Meets requirements for single-storey residential construction in most national codes"),
        ("Curing Time",             "24–48 hours ambient — no kiln, no heat, no controlled environment"),
        ("Aggregate Composition",   "75% crushed concrete rubble / 25% geopolymer paste by volume"),
        ("Cost per Brick",          "Energy + binder only — rubble is free on-site. Binder cost: ~£0.05–0.15 per brick at scale"),
        ("vs. Commercial Brick",    "Post-earthquake brick prices spike 300–500% due to demand surge — rubble brick eliminates import cost entirely"),
    ]))

    # ── COMPETITIVE LANDSCAPE ─────────────────────────────────────────────────
    story += section("8. Competitive Landscape")
    story.append(body(
        "No field-deployable integrated system exists combining rubble crushing, geopolymer "
        "mixing, and compression in a single transportable unit. The closest adjacent products:"
    ))
    story.append(wide_table(
        ["Category", "Existing Products", "Gap vs. RENOVO"],
        [
            ["Compressed Earth Block Presses", "CINVA-Ram, Vermeer BP714, Auram 3000", "Use virgin soil — not rubble. Not designed for disaster deployment or rubble pre-processing."],
            ["Rubble Removal Equipment", "Standard excavators, dump trucks", "Move rubble to landfill — do not process or reuse it."],
            ["NGO Manual Sorting", "Manual labour, simple sieves", "Produces road fill only — not structural bricks. Extremely low output."],
            ["Geopolymer Research", "Academic prototypes (Davidovits et al.)", "Laboratory scale — not field-deployable, not integrated with crushing stage."],
        ],
        col_widths=[4*cm, 5.5*cm, 7*cm]
    ))

    # ── MINIMUM VIABLE PROTOTYPE ───────────────────────────────────────────────
    story += section("9. Minimum Viable Prototype Checklist")
    story.append(body(
        "Dyson Award requires evidence of iterative development — not a finished machine. "
        "The following steps constitute a credible submission prototype:"
    ))
    mvp_rows = [
        ["1", "Mix geopolymer binder (sodium silicate + fly ash/GGBS) in a paper cup. Cure 24 hr at room temp.",
         "Proves no-kiln binder chemistry hardens without heat — core chemistry claim.", "Low — ~£10 materials"],
        ["2", "Hammer-crush a concrete paving slab offcut to 5–20 mm aggregate. Mix with binder at 75:25. Press into any mould (bottle cap, cardboard box). Cure 48 hr.",
         "Physical compressed rubble brick exists — photograph it.", "Low–Medium"],
        ["3", "Photograph brick next to a standard brick for scale.",
         "Physical artifact > any render.", "Trivial"],
        ["4", "Load-test: stand on it, hit with a hammer. Document qualitatively.",
         "Strength evidence without lab equipment.", "Low"],
        ["5", "Build tabletop press: G-clamp + two wooden boards. Photograph in use.",
         "Proves compression achievable with basic tooling.", "Low"],
        ["6", "Write 1-page test log: mix ratio, cure time, ambient temp, result.",
         "Design iteration evidence.", "Trivial"],
    ]
    story.append(wide_table(
        ["#", "Action", "What It Proves", "Difficulty"],
        mvp_rows,
        col_widths=[0.6*cm, 5.5*cm, 6.5*cm, 3.9*cm]
    ))
    story.append(Spacer(1, 4))
    story.append(body("<b>Minimum time to credible prototype:</b> 3–4 days (1 day sourcing + 1 day mixing + 2 days curing)"))

    # ── DYSON APPLICATION ANSWERS ─────────────────────────────────────────────
    story += section("10. James Dyson Award — Application Answers")

    qa = [
        ("Project Name (50 chars)", "Rubble-to-Brick Press"),
        ("Quick Description (193 chars)",
         "A compact press that crushes post-earthquake concrete rubble, mixes it with geopolymer binder, "
         "and produces standard compressed bricks on-site within 48 hours — no kiln, near-zero cost."),
        ("What it does (260 chars)",
         "The press crushes rubble, mixes it with a no-kiln geopolymer binder, and compresses it into "
         "standard bricks curing at ambient temperature. A full pallet of structural bricks is produced "
         "on-site within 48 hours using solar power or manual operation."),
        ("Inspiration (650 chars)",
         "After every major earthquake — Turkey 2023, Morocco 2023, Nepal 2015 — two crises unfold "
         "simultaneously. Communities sit surrounded by millions of tonnes of rubble that must be "
         "expensively removed, while importing new building materials at prices they cannot afford. "
         "The waste and the need exist side by side, completely disconnected. The insight: rubble is "
         "not waste — it is unprocessed raw material. Geopolymer chemistry allows concrete aggregate "
         "to be rebound without heat, through chemical activation alone. Combining these facts pointed "
         "to a field-deployable press that turns disaster debris into reconstruction materials."),
        ("What sets it apart (975 chars)",
         "Existing post-earthquake rubble management has two options: pay for expensive removal and "
         "landfill, or leave rubble in place for years. Some NGOs use manual rubble-sorting for road "
         "fill, but no field-deployable system exists to upcycle concrete rubble into structural bricks. "
         "Geopolymer compressed earth block presses exist for new-build construction but use virgin "
         "soil, not recycled concrete rubble, and are not designed for disaster-site deployment. "
         "RENOVO is the first integrated system combining on-site rubble crushing, geopolymer mixing, "
         "and compression in a single deployable unit. It eliminates material transportation costs in "
         "both directions: rubble stays on-site and new bricks need not be imported. Post-earthquake "
         "brick prices spike 300–500% — rubble bricks cost only the energy and binder to produce."),
        ("Future (650 chars)",
         "First prototype to be validated with concrete rubble from demolition sites, testing "
         "compressive strength across varying aggregate types and contamination levels. Field trials "
         "planned with disaster relief NGOs in earthquake-prone regions. A lightweight flat-pack version "
         "will be developed for air-freight delivery in the first 72 hours after a disaster. Long-term, "
         "the binder chemistry will be refined to work with broader rubble types including brick, tile, "
         "and mixed debris, expanding the addressable waste stream significantly."),
    ]
    for label, text in qa:
        story.append(KeepTogether([
            Paragraph(label, styles["subsection"]),
            Paragraph(text, styles["quote"]),
            Spacer(1, 4),
        ]))

    # ── FUTURE ROADMAP ────────────────────────────────────────────────────────
    story += section("11. Future Roadmap")
    story.append(wide_table(
        ["Phase", "Timeline", "Milestone"],
        [
            ["Phase 0 — Proof of Concept", "Now (Jun 2026)", "Geopolymer binder sample + hammer-crushed rubble brick. Tabletop G-clamp press. Photograph + test log."],
            ["Phase 1 — Prototype", "Jul–Sep 2026", "Steel-framed tabletop jaw crusher + mixing drum + hydraulic jack press. Test with construction site rubble."],
            ["Phase 2 — Field Validation", "2027", "NGO partnership trial in earthquake-affected region. Systematic compressive strength testing across rubble types."],
            ["Phase 3 — Flat-Pack Version", "2027–2028", "Air-freight-deployable variant. 72-hour post-disaster deployment target."],
            ["Phase 4 — Scale", "2028+", "OEM manufacturing partnership. Binder chemistry extended to brick, tile, mixed debris."],
        ],
        col_widths=[4.5*cm, 3.5*cm, 8.5*cm]
    ))

    # ── FOOTER NOTE ───────────────────────────────────────────────────────────
    story.append(Spacer(1, 20))
    story.append(rule(MID_GREY, 0.5))
    story.append(Paragraph(
        "RENOVO — Project 05 | James Dyson Award 2026 | Generated June 2026",
        styles["small"]
    ))

    doc.build(story)
    print(f"PDF saved → {OUTPUT}")

if __name__ == "__main__":
    build()
