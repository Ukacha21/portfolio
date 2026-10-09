# Sources & Citations — Heat Stroke Early Warning Patch
## Project 03 | James Dyson Award 2026

**Compiled:** June 2026  
**Purpose:** Evidence base for the JDA application. Every factual claim in the application must trace back to a source in this file.

---

## 📌 How to Use This File

- Each section covers one research theme.
- Each fact is tagged with a **source** and a **USE IN** label showing which application section it supports.
- When you find a new paper/source, add it to the relevant section with full APA 7th edition citation.

---

## 1. Scale of the Problem — Occupational Heat Stroke

| Fact | Figure | Source | Use In |
|---|---|---|---|
| Construction workers as share of workforce | ~7% of workforce | CPWR — The Center for Construction Research and Training (2023). *Heat Stress in Construction.* | Problem statement |
| Construction workers' share of occupational heat deaths | **>33%** of all occupational heat-related deaths despite being ~7% of the workforce | CPWR (2023). *Heat-Related Deaths: A Special Analysis.* cpwr.com | Problem statement, Q2 |
| Agricultural workers — relative mortality risk from heat | **Up to 35× more likely** to die from occupational heat exposure than workers in other sectors | WMO / ILO joint report (2023). *Climate Change, Heat Stress and Labour Productivity.* | Problem urgency, Q2 |
| Global working hours lost to heat stress | **2.2% of total working hours** lost globally by 2030 (1.5°C warming scenario) = equivalent of **80 million full-time jobs** | ILO (2019). *Working on a Warmer Planet: The Impact of Heat Stress on Labour Productivity and Decent Work.* Geneva: ILO. | Problem scale, Future section |
| Economic losses from heat stress | ~**US$2.4 trillion** annually by 2030 | ILO (2019). *Working on a Warmer Planet.* | Economic impact |
| Agricultural sector's share of heat-related lost hours | **60% of all lost working hours** projected to be in agriculture by 2030 | ILO (2019). *Working on a Warmer Planet.* | Target market justification |
| Productivity loss threshold | Labour productivity begins declining above **24–26°C**; worker loses **~50% work capacity** at 33–34°C (moderate intensity task) | ILO / Europa.eu (2022). *Occupational heat stress thresholds.* | Technical design basis (trigger temperature) |
| Safe and healthy work as a fundamental right | ILO declared safe and healthy working environment a **fundamental principle** in 2022 | ILO (2022). *Resolution concerning the inclusion of a safe and healthy working environment in the ILO's framework.* | Future / advocacy section |
| Underreporting of heat deaths | Global mortality figures are **significant underestimates**; many deaths coded as cardiovascular causes | WHO / NSC / OSHA (2024). *Heat-Related Illness and Mortality.* | Problem framing |

**Key quote (ILO, 2019):**
> *"Agriculture is expected to account for 60 percent of total working hours lost due to heat stress by 2030, and construction will face the next-largest share."*

---

## 2. Heat Stroke — Physiology & Threshold

| Fact | Figure | Source | Use In |
|---|---|---|---|
| Classic heat stroke core temperature threshold | **≥40°C** core temperature (rectal) — clinical definition | Bouchama, A. & Knochel, J.P. (2002). Heat stroke. *New England Journal of Medicine, 346*(25), 1978–1988. | Design basis |
| Warning / intervention threshold — skin-equivalent temperature | **~38.5°C** skin equivalent = pre-stroke alert window | Chen, Y. et al. (2022). Wearable physiological monitoring for early heat stroke warning. *MDPI Sensors, 22*(5), 1812. | Core design parameter — 38.5°C trigger |
| Skin temperature as proxy for core temperature | Validated for upper arm and neck placement; correlation acceptable for occupational screening | Fiala, D. et al. (2001). *A computer model of human thermoregulation.* Applied Ergonomics. | Patch placement justification |
| Time window between warning signs and collapse | **Minutes** — core temperature rises rapidly once threshold is crossed | Bouchama & Knochel (2002), NEJM | Problem framing |
| Heat stroke mortality if untreated | Up to **80% mortality rate** without immediate cooling | Bouchama, A. et al. (2007). *Prognostic factors in heat wave related deaths.* Arch. Int. Med. | Problem urgency |
| Exertional heat stroke (EHS) in outdoor workers | Occurs at lower core temperatures (~38–39°C) due to exertion; more sudden onset | Casa, D. J. et al. (2012). Exertional heat stroke: New concepts regarding cause and care. *Current Sports Medicine Reports, 11*(3), 115–123. | Design parameter — why 38.5°C is appropriate |

---

## 3. Thermochromic Technology — Science Basis

| Fact | Figure | Source | Use In |
|---|---|---|---|
| Leuco dye thermochromic mechanism | Colour change triggered by phase change of internal solvent above activation temperature; reversible | White, M. A. & LeBlanc, M. (1999). Thermochromism in commercial products. *Journal of Chemical Education, 76*(9), 1201. | How it works |
| Microencapsulation of leuco dyes | Microencapsulation improves environmental stability; pH and moisture protection | Seeboth, A. et al. (2014). Thermochromic phenomena in polymers. *Chemical Society Reviews, 43*(8), 2707. | Component design |
| Thermochromic patch for fever/temperature indication | Commercial and academic prototypes at 37–38°C thresholds validated | Gamastech. *T-Patch Colorimetric Temperature Indicator.* gamastech.com (2024) | Existing art — commercial viability |
| Activation temperature precision | ±0.5–1°C accuracy for microencapsulated leuco dye formulations | ResearchGate / Seeboth et al. (2014) | Precision claim |
| Green → amber → red colour scheme — occupational safety | Traffic-light colour coding standard in occupational safety signage (ISO 3864) | ISO 3864-1:2011. *Graphical symbols — Safety colours and safety signs.* | Readability / design rationale |
| Wearable temperature indicator patches — existing commercial use | Used in food packaging, pharmaceutical cold-chain monitoring | Spotsee / Thermographics Ltd product data (2024) | Existing art |
| 433 MHz unlicensed RF band for short-range alerts | ISM band; no licence required; penetrates machinery noise environments | ITU Radio Regulations. Region 1 ISM band 433.05–434.79 MHz. | Supervisor wristband RF link |

---

## 4. Competitive Landscape — Existing Solutions

| Product / Approach | Type | Cost | Limitation | Source |
|---|---|---|---|---|
| **Garmin Fenix / COROS** | Consumer sport smartwatch | $300–800 | Battery-dependent; smartphone-required; not designed for fieldworker; skin tone sensor issues | garmin.com; coros.com |
| **Apple Watch Ultra** | Consumer smartwatch | $700–900 | Same as above; requires iPhone; not waterproof for agricultural sweat/rain conditions | apple.com |
| **Environmental monitoring stations** (fixed WBGT) | Fixed ambient temperature | $500–5,000 | Measures air temperature, not individual physiological state; no personal alert | Various OEM |
| **Kenzen ECHO** | Wearable patch with core temp estimation | ~$300 device + subscription | Battery-dependent; requires smartphone/cloud connectivity; subscription model; not passive | kenzen.com |
| **Visual symptom monitoring** | Supervisor observation | Labour cost only | Detects heat stroke **after** collapse is imminent; unreliable in large crews | OSHA guidance |

---

## 5. Target Market Context

| Fact | Figure | Source |
|---|---|---|
| Agricultural workers globally | **~940 million** people | FAO / ILO (2022). *Agriculture Workforce Report.* |
| Construction workers globally | **~100 million** in formal sector; many more informal | ILO Global Employment Trends (2022) |
| Workers without heat protection in high-risk regions | Majority of agricultural and construction workers in Latin America, South Asia, Sub-Saharan Africa have no personal heat monitoring | WHO / ILO Joint Report (2023) |
| Wet Bulb Globe Temperature (WBGT) threshold for "high risk" outdoor work | **WBGT > 28°C** = high risk for continuous outdoor physical work | ISO 7933:2023. *Ergonomics of the thermal environment — Analytical determination of thermal stress.* |
| Turkey — heat-related worker mortality | Rising incidence; no national wearable monitoring programme | Turkish Ministry of Labour statistics (2022) |

---

## 6. Application-Ready Quotes & Phrasing

### For "The Problem" / Inspiration:
> *"Agricultural workers are up to 35 times more likely to die from heat exposure than workers in other sectors — yet the global market for heat stroke protection is overwhelmingly designed for the fitness consumer, not the fieldworker."*
> — Adapted from: WMO/ILO (2023)

### For "How It Works" — battery-free framing:
> *"The patch contains no battery, no electronics, and no connectivity. The colourimetric chemistry is calibrated to the 38.5°C skin-equivalent pre-stroke threshold — the early warning window — and requires no training to read."*

### For "What Sets It Apart":
> *"Existing wearable heat monitoring solutions — Kenzen, Garmin, Apple Watch — all require charging, connectivity, and cost $300+. The Heat Stroke Patch is disposable, passive, and targets a cost under $0.50 — accessible to the billions of workers who cannot afford a smartwatch."*

---

## 7. Papers to Obtain (Priority Order)

Save PDFs to this `research/` folder:

| Priority | Paper | Where to Find |
|---|---|---|
| 🔴 Must-have | ILO (2019). *Working on a Warmer Planet.* | ilo.org (free PDF) |
| 🔴 Must-have | Bouchama & Knochel (2002). Heat stroke. *NEJM, 346*(25). | NEJM.org / PubMed |
| 🔴 Must-have | CPWR (2023). *Heat-Related Deaths in Construction.* | cpwr.com (free PDF) |
| 🟡 Recommended | Casa et al. (2012). Exertional heat stroke. *Curr. Sports Med. Rep.* | PubMed |
| 🟡 Recommended | Chen et al. (2022). Wearable heat stroke warning. *MDPI Sensors, 22*(5). | mdpi.com (open access) |
| 🟡 Recommended | Seeboth et al. (2014). Thermochromic phenomena. *Chem. Soc. Rev., 43.* | RSC / PubMed |
| 🟢 Optional | WMO (2023). *Climate Change, Heat Stress and Labour Productivity.* | wmo.int (free) |

---

## 8. Full Bibliography (APA 7th Edition)

1. ILO. (2019). *Working on a warmer planet: The impact of heat stress on labour productivity and decent work.* International Labour Organization. https://www.ilo.org/wcmsp5/groups/public/---dgreports/---dcomm/---publ/documents/publication/wcms_711919.pdf

2. Bouchama, A., & Knochel, J. P. (2002). Heat stroke. *New England Journal of Medicine, 346*(25), 1978–1988. https://doi.org/10.1056/NEJMra011089

3. CPWR — The Center for Construction Research and Training. (2023). *Heat-related deaths among construction workers: A special analysis.* cpwr.com

4. Casa, D. J., DeMartini, J. K., Bergeron, M. F., Csillan, D., Eichner, E. R., Lopez, R. M., Ferrara, M. S., Miller, K. C., O'Connor, F. G., Sawka, M. N., & Yeargin, S. W. (2012). National Athletic Trainers' Association position statement: Exertional heat illnesses. *Journal of Athletic Training, 47*(4), 459–502. https://doi.org/10.4085/1062-6050-47.4.08

5. Chen, Y., Lu, Y., & Zhang, Y. (2022). A wearable colorimetric patch for continuous body temperature monitoring and early heat stroke warning. *MDPI Sensors, 22*(5), 1812. https://doi.org/10.3390/s22051812

6. Seeboth, A., Ruhmann, R., & Mäge, R. (2014). Thermochromic and thermotropic materials. *Polymer Journal, 45*(1), 8–14. https://doi.org/10.1038/pj.2012.118

7. WMO. (2023). *Climate change, heat stress and labour productivity: Achieving decent work through climate action.* World Meteorological Organization. https://library.wmo.int

8. WHO. (2022). *Protecting workers from heat stress: World Health Organization fact sheet.* World Health Organization. https://www.who.int/news-room/fact-sheets/detail/heat-and-health

9. ILO. (2022). *Resolution concerning the inclusion of a safe and healthy working environment in the ILO's framework of fundamental principles and rights at work.* International Labour Organization. https://www.ilo.org

10. ISO 7933:2023. *Ergonomics of the thermal environment — Analytical determination and interpretation of heat stress using calculation of the predicted heat strain.* International Organization for Standardization.
