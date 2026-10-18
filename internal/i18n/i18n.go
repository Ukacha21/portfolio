package i18n

import (
	"net/http"
	"strings"
)

// Language represents a supported language configuration
type Language struct {
	Code      string
	Name      string
	Native    string
	Flag      string
	IsRTL     bool
}

// SupportedLanguages list
var SupportedLanguages = []Language{
	{Code: "en", Name: "English", Native: "English", Flag: "🇬🇧", IsRTL: false},
	{Code: "tr", Name: "Turkish", Native: "Türkçe", Flag: "🇹🇷", IsRTL: false},
	{Code: "ar", Name: "Arabic", Native: "العربية", Flag: "🇸🇦", IsRTL: true},
	{Code: "pt", Name: "Portuguese", Native: "Português", Flag: "🇵🇹", IsRTL: false},
	{Code: "es", Name: "Spanish", Native: "Español", Flag: "🇪🇸", IsRTL: false},
	{Code: "de", Name: "German", Native: "Deutsch", Flag: "🇩🇪", IsRTL: false},
	{Code: "fr", Name: "French", Native: "Français", Flag: "🇫🇷", IsRTL: false},
}

// TranslationData holds all text tokens for pages
type TranslationData struct {
	LangCode       string
	LangName       string
	IsRTL          bool
	Dir            string
	Languages      []Language

	// Navigation
	NavBrand       string
	NavHome        string
	NavProjects    string
	NavExperience  string
	NavAbout       string
	NavSkills      string
	NavContact     string
	NavEducation        string
	NavHonors      string
	NavCertificates     string
	NavActivities       string
	NavLanguages        string
	TabEducation        string
	TabAwards           string
	TabCertificates     string
	TabActivities       string
	TabLanguages        string
	TabSoloProjects     string
	TabTeamProjects     string
	TabContractProjects string
	ContractEmptyBadge  string
	ContractEmptyTitle  string
	ContractEmptyDesc   string
	CP1Title            string
	CP1Desc             string
	CP1Domain           string
	CP2Title            string
	CP2Desc             string
	CP2Domain           string
	BtnInstagramPost    string
	BtnLinkedInPost     string
	CP3Title            string
	CP3Desc             string
	CP3Domain           string
	CP4Title            string
	CP4Desc             string
	CP4Domain           string
	CP5Title            string
	CP5Desc             string
	CP5Domain           string
	CPNoticeInfra       string
	BtnOfficialWebsite  string
	BadgeVolunteered    string
	ExcavatorTitle      string
	ExcavatorDesc       string
	ExcavatorDomain     string
	SafrenTitle         string
	SafrenDesc          string
	SafrenDomain        string
	STM32Title          string
	STM32Desc           string
	STM32Domain         string
	ROSTitle            string
	ROSDesc             string
	ROSDomain           string
	BtnLinkedInVideo    string
	BtnGitHubRepo       string
	ActivitiesTitle     string
	ActivitiesSubtitle  string
	TeknoselTitle       string
	TeknoselDesc        string
	TeknoselMayorBadge  string
	AiResearchTitle     string
	AiResearchDesc      string
	AiResearchBadge     string
	IfmWorkshopTitle    string
	IfmWorkshopDesc     string
	IfmWorkshopBadge    string
	IfmBtnLinkedIn      string
	AwardsTag      string
	AwardsTitle    string
	AwardsSubtitle string
	AwardsCardTitle string
	AwardsCardDesc  string
	AwardsViewPost  string
	AwardsDeanTitle string
	AwardsDeanName  string
	AwardsTeamName  string
	AwardsOlympiad  string
	AwardsCeremony  string
	AwardsQuoteLabel string
	NavGitHub      string
	NavVelum       string

	// Home Hero
	HeroBadge      string
	AuthorRole     string
	HeroTitle      string
	HeroSubtitle   string
	HeroBtnExplore string
	HeroBtnContact string
	StatProjects         string
	StatSoloProjects     string
	StatTeamProjects     string
	StatContractProjects string
	StatWorkshops        string
	StatAward            string
	StatLanguages  string
	StatFocus      string

	// Project Grid
	ProjectsTitle     string
	ProjectsSubtitle  string
	StatusActive      string
	StatusQueue       string
	BtnViewProject    string
	BtnPreview        string
	BadgeFeatured     string
	BadgeBiomedical   string
	BadgeHardware     string
	BadgeRobotics     string

	// Projects descriptions
	P1Title string
	P1Desc  string
	P1Domain string

	P2Title string
	P2Desc  string
	P2Domain string

	P3Title string
	P3Desc  string
	P3Domain string

	P4Title string
	P4Desc  string
	P4Domain string

	P5Title string
	P5Desc  string
	P5Domain string

	P6Title string
	P6Desc  string
	P6Domain string

	P7Title string
	P7Desc  string
	P7Domain string

	P8Title string
	P8Desc  string
	P8Domain string

	P9Title string
	P9Desc  string
	P9Domain string

	P10Title string
	P10Desc  string
	P10Domain string

	P11Title string
	P11Desc  string
	P11Domain string

	P12Title string
	P12Desc  string
	P12Domain string

	P13Title string
	P13Desc  string
	P13Domain string

	P14Title string
	P14Desc  string
	P14Domain string

	// About & Skills
	AboutTitle       string
	AboutSubtitle    string
	AboutP1          string
	AboutP2          string
	AboutPolyglot    string
	AboutPolyglotDesc string
	SkillsTitle      string
	ExpTitle         string
	ExpSubtitle      string
	EduTitle         string
	CertsTitle       string

	// Velum Page
	VHeroBadge       string
	VHeroTitle       string
	VHeroSubtitle    string
	VStatCost        string
	VStatCostDesc    string
	VStatEarly       string
	VStatEarlyDesc   string
	VStatBOM         string
	VStatBOMDesc     string
	VStatEquity      string
	VStatEquityDesc  string

	VProblemTitle    string
	VProblemSubtitle string
	VProblemP1       string
	VProblemP2       string
	VProblemPoint1   string
	VProblemPoint2   string
	VProblemPoint3   string

	VSolutionTitle   string
	VSolutionSubtitle string
	VSolutionP1      string
	VSolutionPoint1  string
	VSolutionPoint2  string
	VSolutionPoint3  string

	VSimTitle        string
	VSimSubtitle     string
	VSimNormal       string
	VSimIschemia     string
	VSimOedema       string
	VSimReset        string
	VSimStatusAlert  string
	VSimStatusNormal string

	VTechTitle       string
	VTechSubtitle    string
	VLayer1Title     string
	VLayer1Desc      string
	VLayer2Title     string
	VLayer2Desc      string
	VLayer3Title     string
	VLayer3Desc      string
	VLayer4Title     string
	VLayer4Desc      string

	VCompTitle       string
	VCompSubtitle    string
	VCompColMetric   string
	VCompColVisual   string
	VCompColSEM      string
	VCompColMattress string
	VCompColVelum    string

	VRoadmapTitle    string
	VRoadmapSubtitle string

	VGalleryTitle    string
	VGallerySubtitle string

	VDownloadPitch   string
	VViewGitHub      string

	// Footer
	FooterRights     string
	FooterBuiltWith  string
}

var translations = map[string]TranslationData{
	"en": {
		LangCode: "en", LangName: "English", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Home", NavProjects: "Projects", NavExperience: "Experience", NavAbout: "About", NavSkills: "Skills", NavContact: "Contact", NavEducation: "Education", NavHonors: "Awards", NavCertificates: "Certificates", NavActivities: "Workshops", NavLanguages: "Languages", NavGitHub: "GitHub", NavVelum: "VELUM Project",
		HeroBadge: "Mechatronics Engineer · Biomedical Hardware · Embedded Systems · Polyglot",
		AuthorRole: "B.Sc. in Mechatronics Engineering · Embedded Systems & IoT Developer",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Designing intelligent biomedical mechatronics, precision sensor instrumentation, and edge IoT devices that protect vulnerable human lives.",
		HeroBtnExplore: "Explore Projects", HeroBtnContact: "Get in Touch",
		StatProjects: "18 Engineered Projects", StatSoloProjects: "Solo Projects", StatTeamProjects: "Team Projects", StatContractProjects: "Contract Projects", StatWorkshops: "Workshops", StatAward: "Mechatronics Engineering", StatLanguages: "7 Fluent Languages", StatFocus: "ESP32 / Bioimpedance / CAD",
		ProjectsTitle: "Engineering Portfolio",
		ProjectsSubtitle: "Proactive biomedical hardware, assistive devices, fluidic automation, and mechatronic systems.",
		StatusActive: "Complete & Interactive", StatusQueue: "In Queue",
		BtnViewProject: "View Project", BtnPreview: "Project Details",
		BadgeFeatured: "Featured Project", BadgeBiomedical: "Biomedical Hardware", BadgeHardware: "Mechatronics", BadgeRobotics: "Robotics & Automation",
		P1Title: "VELUM (formerly SentIO)",
		P1Desc: "Continuous 24/7 Sub-Epidermal Moisture & Bioelectrical Impedance Sensing Mat that detects pressure injuries 2–5 days before tissue damage appears.",
		P1Domain: "Biomedical Mechatronics · Bioimpedance Spectroscopy",
		P2Title: "FLOWGUARD",
		P2Desc: "Automated mechatronic microplastic filtration system for residential and commercial washing machines preventing oceanic synthetic micro-waste.",
		P2Domain: "Fluidic Mechatronics · Microplastic Separation",
		P3Title: "SOLUS",
		P3Desc: "Wearable continuous core temperature & physiological monitoring patch delivering early predictive warnings for heat stroke in extreme workers.",
		P3Domain: "Wearable Biosensors · Preventative Physiology",
		P4Title: "DENOVA (formerly Avicenna)",
		P4Desc: "Non-invasive optical tooth decay detector utilizing multispectral illumination and fluorescence imaging to diagnose early-stage demineralization.",
		P4Domain: "Biophotonics · Optical Diagnostics",
		P5Title: "RENOVO",
		P5Desc: "Portable disaster-relief hydraulic press converting urban earthquake demolition rubble into interlocking structural building blocks on site.",
		P5Domain: "Heavy Mechanical & Hydraulics · Disaster Relief",
		P6Title: "CURA SERO",
		P6Desc: "Micro-tension bio-adhesive surgical wound closure strip ensuring uniform epidermal tensile distribution and minimizing scar formation.",
		P6Domain: "Bio-Mechanical Device · Surgical Technology",
		P7Title: "WELSTECH — Wearable Life Saver Technology",
		P7Desc: "Ultra-compact wrist-worn vital signs monitor integrating optical PPG biosensing, ESP32-S3 processing, local Wi-Fi telemetry, and acoustic alarm escalation for cardiac anomaly detection.",
		P7Domain: "Biomedical Wearable IoT · Cardiac Biosensing",
		P8Title: "SENDIT — Airport Item Mail-Home Kiosk",
		P8Desc: "Autonomous airport security kiosk and courier logistics system allowing passengers to package, label, and courier-dispatch surrendered or over-limit items home in under 2 minutes.",
		P8Domain: "Smart Airport Infrastructure · Courier Automation",
		P9Title: "SMARTWASTE — Autonomous Waste Sorting & 3:1 Compactor",
		P9Desc: "Autonomous retrofittable smart lid combining millimeter ToF laser fill-level telemetry, bimodal inductive/optical sorting gates, and lead-screw 3:1 compaction for municipal containers.",
		P9Domain: "Autonomous Mechatronics · Smart Municipal IoT",
		P10Title: "IMAGE CLASSIFIER — Deep Convolutional Neural Network",
		P10Desc: "Custom deep convolutional neural network (CNN) trained with TensorFlow/Keras on 9,600+ curated image datasets, featuring automated 96×96 preprocessing, multi-layer feature extraction, and real-time inference GUI.",
		P10Domain: "Deep Learning · Computer Vision & Edge AI",
		P11Title: "LUMIERE — Smart Streetlight Failure & Geolocation Telemetry System",
		P11Desc: "Compact IoT failure-reporting device mounted on municipal streetlamps, combining LDR photometric sensing, GY-NEO6MV2 GPS geolocation, ESP32 Wi-Fi SoftAP telemetry, and escalating buzzer alarms to automate outage dispatch.",
		P11Domain: "Smart City IoT · Embedded Firmware & Telemetry",
		P12Title: "NOTCHA — Minimal Floating Dynamic Time Notch",
		P12Desc: "Born out of ergonomic necessity in 2021/2022 when Windows 11 removed top taskbar placement, NOTCHA is an ultra-lightweight desktop utility providing an unobtrusive floating time notch at the top screen bezel with real-time updates and smart hover-expand mechanics.",
		P12Domain: "Desktop Ergonomics · Custom UI Utilities",
		P13Title: "SPO2 — Pulse Oximeter 2-Layer PCB Design",
		P13Desc: "A robust 2-layer pulse oximeter circuit board engineered in KiCad 10.0, integrating custom schematic symbols, footprints built from scratch, real manufacturer tolerances, MAX30102 optical biosensing, SSD1306 OLED telemetry, and an Arduino Nano core.",
		P13Domain: "PCB Design & EDA · Biomedical Hardware · KiCad 10.0",
		P14Title: "RESPIRO",
		P14Desc: "Low-cost, high-reliability biomedical adaptive assistive ventilator engineered for emergency deployment and resource-constrained clinics.",
		P14Domain: "Biomedical Mechatronics · Critical Care",
		AboutTitle: "About Ukacha Juma",
		AboutSubtitle: "Mechatronics Engineer · Hardware-Software Architect · Polyglot",
		AboutP1: "I am a Mechatronics Engineer driven by the conviction that modern hardware engineering should prevent suffering rather than merely treat it. My work bridges mechanical CAD/CAM design, precision analog sensor circuits, firmware on bare-metal and FreeRTOS microcontrollers, and real-time telemetry systems.",
		AboutP2: "With projects spanning proactive pressure-injury mats (VELUM), microplastic filtration (FLOWGUARD), and thermal biosensors (SOLUS), I build physical systems from early mathematical principles through benchtop circuit validation to production-ready design.",
		AboutPolyglot: "Multilingual Communication & Global Collaboration",
		AboutPolyglotDesc: "Fluent in English, Turkish, Arabic, Portuguese, Spanish, German, and French. Being a polyglot empowers me to bridge technical teams across continents, read global clinical literature in original sources, and design universally accessible technologies.",
		SkillsTitle: "Technical Proficiencies & Toolchain",
		ExpTitle: "Professional Experience", ExpSubtitle: "Hands-on industrial engineering, CNC diagnostics, automation maintenance, and international translation.", EduTitle: "Education & Academic Background", AwardsTag: "ACADEMIC & ROBOTIC RECOGNITION", AwardsTitle: "Honors & Awards", AwardsSubtitle: "Distinguished faculty honors, national robotics olympiad rankings, and institutional engineering excellence.", AwardsCardTitle: "İTÜRO 2026 Ranking — Faculty Plaque Conferred by the Dean", AwardsCardDesc: "Honored with a commemorative plaque presented by esteemed Dean Prof. Dr. Humar Kahramanli Örnek at the Selçuk University Faculty of Technology graduation ceremony (2025–26), recognizing team STWolf's ranking at the prestigious Istanbul Technical University Robot Olympiad (İTÜRO 2026).", AwardsViewPost: "View LinkedIn Post & Recognition", AwardsDeanTitle: "Dean of the Faculty of Technology", AwardsDeanName: "Prof. Dr. Humar Kahramanlı Örnek", AwardsTeamName: "Team STWolf · Selçuk University", AwardsOlympiad: "İstanbul Technical University Robot Olympiad (İTÜRO 2026)", AwardsCeremony: "Graduation Ceremony 2025–26 · Selçuk University", AwardsQuoteLabel: "Official Plaque Inscription", TabEducation: "Education", TabAwards: "Honors & Awards", TabCertificates: "Certificates & Licenses", TabActivities: "Workshops", TabLanguages: "Languages", ActivitiesTitle: "Extracurricular & Technical Activities", ActivitiesSubtitle: "Festival project exhibitions, community leadership, and AI research methodologies.", TeknoselTitle: "Selçuklu Teknosel Technology Festival — University Category", TeknoselDesc: "Exhibited the KonyaWaste / SmartWaste mechatronics prototype at our booth in the Free Project - University Level category, engaging with fellow engineers and presenting our smart city hardware to the Mayor of Selçuklu in person.", TeknoselMayorBadge: "Personal Presentation to the Mayor of Selçuklu", AiResearchTitle: "Source-Grounded AI in Academic Research (TADD Workshop)", AiResearchDesc: "Completed an intensive 2-day workshop led by instructor Yunus Emre Kıymaz and TADD, mastering source-grounded AI techniques with Zotero, ResearchRabbit, and NotebookLM to eliminate AI hallucinations and ensure rigorous technical accuracy for engineering literature reviews.", AiResearchBadge: "Zero-Hallucination Source-Grounded Research", IfmWorkshopTitle: "Expanding my Automation Toolkit with ifm electronic & CoDeSys (48-Hour Intensive)", IfmWorkshopDesc: "Completed a specialized 48-hour intensive industrial PLC training program hosted on-site at Selçuk University by ifm electronic Türkiye headquarters. Gained hands-on experience in PLC architecture, CoDeSys V3.5 (SP18 & SP19), Modbus and CANbus industrial communication protocols, and high-pressure site troubleshooting.", IfmWorkshopBadge: "⚙️ ifm electronic Türkiye · Selçuk University", IfmBtnLinkedIn: "View ifm Workshop LinkedIn Post", TabSoloProjects: "Solo Projects", TabTeamProjects: "Team Projects", TabContractProjects: "Contract Projects", ContractEmptyBadge: "Available for Contracting", ContractEmptyTitle: "Client & Contract Engineering Engagements", ContractEmptyDesc: "Proprietary commercial contracts, hardware consulting, and industrial mechatronics deliverables. Portfolio case studies coming soon.",
		CP1Title: "Mec\u00e2nica Auto Moderna \u2014 Corporate Brand & Logo Design", CP1Desc: "Custom corporate visual brand identity and emblem designed for Mec\u00e2nica Auto Moderna. Engineered a distinctive automotive repair insignia combining precision tool symbolism, continuous 24h reliability, and dynamic industrial typography.", CP1Domain: "Corporate Brand Identity & Vector Design",
		CP2Title: "Mec\u00e2nica Auto Moderna \u2014 Invoice Moderna System", CP2Desc: "Proprietary enterprise billing and management platform engineered for automotive repair facility Mec\u00e2nica Auto Moderna. Streamlines automated electronic invoice generation (Fatura Proforma), VAT compliance, repair order tracking, and client records to maximize operational precision.", CP2Domain: "Enterprise Billing & Electronic Invoicing Software",
		BtnInstagramPost: "View Instagram Post", BtnLinkedInPost: "View LinkedIn Post",
		CP3Title: "APACOV \u2014 IT Infrastructure & Web Architecture", CP3Desc: "Volunteered full-stack institutional web portal and end-to-end IT infrastructure architecture for APACOV (Associa\u00e7\u00e3o Para Protec\u00e7\u00e3o e Amparo a Crian\u00e7as \u00d3rf\u00e3s e Vulner\u00e1veis), an NGO dedicated to sheltering and empowering vulnerable children. Serving as Lead IT Administrator managing server deployments, digital security, and international foreign affairs.", CP3Domain: "IT Infrastructure \u00b7 Web Architecture \u00b7 Foreign Affairs",
		CP4Title: "APACOV \u2014 Humanitarian Brand & Logo Design", CP4Desc: "Pro bono visual identity and emblem designed for humanitarian non-profit APACOV. Crafted in Adobe Illustrator and refined with Samsung PenUp, the insignia portrays protective hands reaching out in solidarity enclosed in an arc of continuous shelter and hope.", CP4Domain: "Humanitarian Brand Identity \u00b7 Vector Design",
		CP5Title: "SuPublic — Automated Drinking Water Kiosk Sheet Metal Engineering", CP5Desc: "Contracted mechanical engineering and SolidWorks parametric sheet metal design for SuPublic, an eco-innovation startup founded by Gonca Özata under the academic mentorship of Dr. Yavuz Selim Taşpınar. Initiated in Konya to map free, clean drinking water fountains across Turkey and globally, SuPublic eliminates single-use plastic bottle waste. Engineered the complete sheet metal enclosure, structural chassis, laser-cut bend patterns, and ergonomic bottle refill niche for the automated dispensing kiosk prototype.", CP5Domain: "Contract Engineering · Sheet Metal Design · SolidWorks CAD & Prototyping",
		CPNoticeInfra: "Server Migration Note: Active server transition and self-hosted server migration currently in progress for apacov.org, alongside institutional domain and mail provisioning.", BtnOfficialWebsite: "Official Website (apacov.org)", BadgeVolunteered: "Volunteered \u00b7 Pro Bono",
		ExcavatorTitle: "Excavator Design (Makine Elemanları Dersi) — Komatsu PC200 LC-7/LC-8 Inspired",
		ExcavatorDesc: "Comprehensive mechanical system design and structural FEA analysis of a 24-ton heavy hydraulic crawler excavator. Conducted rigorous analytical calculations and ANSYS validations across 5 core machine subsystems: heavy-duty pivot pins & shafts (Miller), slewing ring & track roller bearings (Rulmanlar), preloaded structural fastener joints (Civatalar), planetary swing gearboxes (Dişliler), and high-pressure hydraulic cylinders (Hidrolik).",
		ExcavatorDomain: "Heavy Machinery Design · Machine Elements · ANSYS FEA & Hydraulics",
		SafrenTitle: "SafRen (formerly Ce-Saf) — Smart Shoe Pedometer & Wearable IoT System",
		SafrenDesc: "A collaborative smart wearable shoe pedometer and gait telemetry system developed in cross-disciplinary partnership with Shoe Design majors Safina Najmudinova and Ceren Başaran, who provided the footwear and engineered the optimum ergonomic placement for the electronic hardware (OLED LCD, tactile ground switches, and microcontroller). Features a compact Waveshare ESP32-S3-Zero, MPU6050 6-DOF IMU for dynamic strike detection, on-chip SoftAP Wi-Fi web server, and seamless bidirectional synchronization with a custom Flutter companion mobile app.",
		SafrenDomain: "Wearable IoT · Gait Analysis · ESP32-S3 Firmware · Flutter Mobile",
		STM32Title: "STM32 Discovery Board Challenge — Autonomous Mobile Robotics",
		STM32Desc: "Dual-challenge autonomous mobile robotics powered by the STM32F407VGT6 ARM Cortex-M4 Discovery Board. Challenge 1 (Midterm): Multi-channel LDR sensor array with 12-bit ADC gradient tracking and TIM3 PWM differential drive to locate random light sources in a 1 m² arena and stop with millimeter accuracy without collision on the first run. Challenge 2 (Final): Autonomous candle extinguisher employing a 12V high-velocity blower fan, strategic arena-center homing, and 360° dynamic vortex rotation—extinguishing all 5 perimeter candles in under 10 seconds (1/3 of the allowed 30s limit).",
		STM32Domain: "Embedded Systems · STM32 ARM Cortex-M4 · Bare-Metal C / HAL · Autonomous Robotics",
		ROSTitle: "ROS Noetic Autonomous Mobile Robot — ESP32 TCP/IP Bridge & Differential Kinematics",
		ROSDesc: "Autonomous mobile robotics engineered in partnership with Mehmet Çakaroğlu under advisor Dr. Öğr. Üyesi Okan Uyar. Bridges ROS Noetic on Linux with an ESP32 dual-core microcontroller over a high-throughput, low-latency Wi-Fi TCP/IP socket bridge. Features differential drive kinematics, closed-loop odometry with dual 420-tick quadrature encoders, calibrated MPU6050 6-DOF IMU, a 500ms deadman failsafe, and Gazebo/RViz digital twin trajectory tracking meeting 100/100 benchmark criteria for circular and waypoint arena challenges.",
		ROSDomain: "Autonomous Mobile Robotics · ROS Noetic · Differential Kinematics · Gazebo & RViz",
		BtnLinkedInVideo: "LinkedIn Video Demo", BtnGitHubRepo: "GitHub Repository", CertsTitle: "Certifications & Honors",
		VHeroBadge: "James Dyson Award 2026 Submission · Healthcare Innovation",
		VHeroTitle: "VELUM — Pressure Injury Prevention Mat",
		VHeroSubtitle: "The world's first continuous, passive biomedical sensor mat that detects pressure injuries 2–5 days before they become visible — equitably, affordably, and without clinical disruption.",
		VStatCost: "$11B+ USD", VStatCostDesc: "Annual US hospital cost treating preventable pressure ulcers",
		VStatEarly: "2–5 Days Early", VStatEarlyDesc: "Pre-visual warning via Sub-Epidermal Moisture changes",
		VStatBOM: "< $50 USD", VStatBOMDesc: "Target unit BOM enabling deployment across every bed",
		VStatEquity: "100% Equitable", VStatEquityDesc: "Skin-tone agnostic bioimpedance sensing (Fitzpatrick I–VI)",
		VProblemTitle: "A Flawed System Built to Fail Patients",
		VProblemSubtitle: "Visual inspection detects wounds only after tissue necrosis is irreversible.",
		VProblemP1: "Pressure injuries affect over 12.8% of hospitalised patients worldwide. Current standard of care relies on manual nurse turning rounds every 2 hours and subjective visual skin inspections. However, cellular damage originates deep within sub-epidermal tissue 48 to 120 hours before any surface erythema appears.",
		VProblemP2: "Furthermore, visual erythema assessment systematically fails in patients with darker skin tones (Fitzpatrick Types IV–VI), creating severe health equity disparities. Commercial handheld SEM scanners cost $3,000–$10,000 and demand scheduled manual nurse labor, excluding care homes.",
		VProblemPoint1: "Reactive by Design: Surface redness indicates existing cellular tissue necrosis.",
		VProblemPoint2: "Documented Racial Disparity: Visual detection misses Stage 1 ulcers in dark skin tones.",
		VProblemPoint3: "Cost & Staffing Bottlenecks: Handheld probes cost thousands and worsen nurse alarm fatigue.",
		VSolutionTitle: "The VELUM Solution: Continuous Bioimpedance",
		VSolutionSubtitle: "Passive, 24/7 localized tissue monitoring right beneath the hospital bedsheet.",
		VSolutionP1: "VELUM is an ultra-thin, flexible polyimide sensor mat placed beneath bed-bound patients. A dense grid of screen-printed silver-carbon electrodes performs bioelectrical impedance spectroscopy (1–100 kHz) across 16 critical anatomical zones every 5 minutes.",
		VSolutionPoint1: "Zero-Setup Deployment: Slide under bedsheet, power on, zero patient disruption.",
		VSolutionPoint2: "Sub-Epidermal Moisture Detection: A 15%+ drop in impedance signals fluid accumulation days before wounds form.",
		VSolutionPoint3: "Instant BLE / Wi-Fi Alerts: Nurses receive actionable alerts identifying the exact anatomical zone requiring repositioning.",
		VSimTitle: "Interactive Bioimpedance Mat Simulator",
		VSimSubtitle: "Real-time 16-zone anatomical pressure mapping & AD5933 impedance sweep simulation.",
		VSimNormal: "1. Normal Baseline", VSimIschemia: "2. Sustained Sacral Pressure",
		VSimOedema: "3. Stage 0 Oedema Alert", VSimReset: "4. Caregiver Repositioning",
		VSimStatusAlert: "CRITICAL: Sub-Epidermal Oedema Detected! Reposition Sacrum Immediately",
		VSimStatusNormal: "NOMINAL: Tissue perfusion optimal across all anatomical zones",
		VTechTitle: "Exploded Mechatronic Architecture",
		VTechSubtitle: "Precision flexible electronics engineered for clinical disinfection and mass production.",
		VLayer1Title: "1. Medical TPU & Platinum Silicone Encapsulation",
		VLayer1Desc: "Vapor-permeable, liquid-impermeable biocompatible skin (ISO 10993) preserving microclimate and resisting daily hospital-grade disinfectant wipes.",
		VLayer2Title: "2. Screen-Printed Silver/Carbon Electrode Grid",
		VLayer2Desc: "Screen-printed conductive inks on 50 µm DuPont Kapton® polyimide film arranged in tetrapolar bioimpedance arrays under bony prominences.",
		VLayer3Title: "3. AD5933 Impedance Converter & 16-Ch MUX",
		VLayer3Desc: "Analog Devices AD5933 paired with an enhanced Howland current source (<1 mA AC excitation) and CD74HC4067 multiplexer scanning zones in seconds.",
		VLayer4Title: "4. ESP32-S3 Dual-Core Compute & Wireless Module",
		VLayer4Desc: "Xtensa LX7 @ 240MHz running FreeRTOS with rolling baseline calibration DSP, BLE 5.0 mesh, and hospital EMR encrypted Wi-Fi telemetry.",
		VCompTitle: "Competitive Clinical Advantage",
		VCompSubtitle: "Why VELUM creates an unmatched clinical and economic moat.",
		VCompColMetric: "Capability / Metric", VCompColVisual: "Visual Inspection", VCompColSEM: "Handheld SEM Probe", VCompColMattress: "Pressure Mattresses", VCompColVelum: "VELUM Mat",
		VRoadmapTitle: "Clinical & Commercial Roadmap",
		VRoadmapSubtitle: "From benchtop validation to ICU multi-site trials and FDA 510(k) pre-submission.",
		VGalleryTitle: "Hardware Renders & Design Artifacts",
		VGallerySubtitle: "High-resolution product illustrations and mechatronic exploded diagrams.",
		VDownloadPitch: "Download Pitch Deck PDF",
		VViewGitHub: "View GitHub Repository",
		FooterRights: "All rights reserved.",
		FooterBuiltWith: "",
	},
	"tr": {
		LangCode: "tr", LangName: "Turkish", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Ana Sayfa", NavProjects: "Projeler", NavExperience: "Deneyim", NavAbout: "Hakkımda", NavSkills: "Yetenekler", NavContact: "İletişim", NavEducation: "Eğitim", NavHonors: "Ödüller", NavCertificates: "Sertifikalar", NavActivities: "Çalıştaylar", NavLanguages: "Diller", NavGitHub: "GitHub", NavVelum: "VELUM Projesi",
		HeroBadge: "Mekatronik Mühendisi · Biyomedikal Donanım · Gömülü Sistemler · Çok Dilli",
		AuthorRole: "Mekatronik Mühendisliği (Lisans) · Gömülü Sistemler & IoT Geliştirici",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "İnsan hayatını koruyan akıllı biyomedikal mekatronik sistemler, hassas sensör enstrümantasyonu ve uç IoT cihazları tasarlıyorum.",
		HeroBtnExplore: "Projeleri Keşfet", HeroBtnContact: "İletişime Geç",
		StatProjects: "18 Mühendislik Projesi", StatSoloProjects: "Bireysel Projeler", StatTeamProjects: "Takım Projeleri", StatContractProjects: "Sözleşmeli Projeler", StatWorkshops: "Çalıştaylar & Etkinlikler", StatAward: "Mekatronik Mühendisliği", StatLanguages: "7 Akıcı Yabancı Dil", StatFocus: "ESP32 / Biyoempedans / CAD",
		ProjectsTitle: "Mühendislik Portföyü",
		ProjectsSubtitle: "Önleyici biyomedikal donanım, yardımcı cihazlar, akışkan otomasyonu ve mekatronik sistemler.",
		StatusActive: "Tamamlandı & Etkileşimli", StatusQueue: "Sırada",
		BtnViewProject: "Projeyi İncele", BtnPreview: "Proje Detayları",
		BadgeFeatured: "Öne Çıkan Proje", BadgeBiomedical: "Biyomedikal Donanım", BadgeHardware: "Mekatronik", BadgeRobotics: "Robotik & Otomasyon",
		P1Title: "VELUM (Eski Adıyla SentIO)",
		P1Desc: "Yatak yaralarını doku hasarı görünür olmadan 2–5 gün önce tespit eden, 7/24 sürekli sub-epidermal nem ve biyoelektriksel empedans algılama matı.",
		P1Domain: "Biyomedikal Mekatronik · Biyoempedans Spektroskopisi",
		P2Title: "FLOWGUARD",
		P2Desc: "Sentetik mikro-atıkların denizlere karışmasını önleyen, çamaşır makineleri için otomatik mekatronik mikroplastik filtreleme sistemi.",
		P2Domain: "Akışkan Mekatroniği · Mikroplastik Ayrıştırma",
		P3Title: "SOLUS",
		P3Desc: "Zorlu çalışma koşullarındaki bireylerde sıcak çarpmasını önceden haber veren giyilebilir sürekli çekirdek vücut ısısı ve fizyoloji yaması.",
		P3Domain: "Giyilebilir Biyosensörler · Önleyici Fizyoloji",
		P4Title: "DENOVA (Eski Adıyla Avicenna)",
		P4Desc: "Erken aşama demineralizasyonu teşhis etmek için çok spektrumlu aydınlatma ve floresan görüntüleme kullanan cerrahi olmayan optik diş çürüğü tespit cihazı.",
		P4Domain: "Biyofotonik · Optik Teşhis",
		P5Title: "RENOVO",
		P5Desc: "Deprem ve afet molozlarını sahada geçmeli yapı bloklarına dönüştüren taşınabilir hidrolik briket presi.",
		P5Domain: "Ağır Mekanik & Hidrolik · Afet Yardımı",
		P6Title: "CURA SERO",
		P6Desc: "Epidermal gerilimi eşit dağıtarak yara izini en aza indiren mikro-gerilimli biyo-yapışkan cerrahi yara kapatma şeridi.",
		P6Domain: "Biyo-Mekanik Cihaz · Cerrahi Teknoloji",
		P7Title: "WELSTECH — Giyilebilir Hayat Kurtarıcı Teknoloji",
		P7Desc: "Optik PPG biyosensörü, ESP32-S3 mikrodenetleyici, yerel Wi-Fi telemetrisi ve kardiyak anomali tespitinde sesli alarm yükseltmesi sunan ultra kompakt bileklik.",
		P7Domain: "Biyomedikal Giyilebilir IoT · Nabız Biyosensörü",
		P8Title: "SENDIT — Havalimanı Eşya Gönderim Kiosku",
		P8Desc: "Güvenlik noktalarında el konulan eşyaları 2 dakikada paketleyip ev adresine kargolayan otonom havalimanı kiosku ve lojistik sistemi.",
		P8Domain: "Akıllı Havalimanı Altyapısı · Kargo Otomasyonu",
		P9Title: "SMARTWASTE — Otonom Atık Ayrıştırma ve 3:1 Kompaktör",
		P9Desc: "Milimetrik ToF lazer doluluk telemetrisi, çift modlu endüktif/optik ayrıştırma kapağı ve 3:1 vidalı mil sıkıştırma presini birleştiren otonom akıllı konteyner kapağı.",
		P9Domain: "Otonom Mekatronik · Akıllı Belediye IoT",
		P10Title: "IMAGE CLASSIFIER — Derin Evrişimsel Sinir Ağı",
		P10Desc: "TensorFlow/Keras ile 9.600+ özel veri seti üzerinde eğitilmiş, otomatik 96×96 ön işleme, çok katmanlı öznitelik çıkarımı ve masaüstü çıkarım arayüzü sunan derin CNN modeli.",
		P10Domain: "Derin Öğrenme · Bilgisayarlı Görü & Uç Yapay Zeka",
		P11Title: "LUMIERE — Akıllı Sokak Lambası Arıza ve Konum Bildirim Sistemi",
		P11Desc: "Belediye sokak lambalarına monte edilen, LDR ışık sensörü, GY-NEO6MV2 GPS konumlandırma, ESP32 Wi-Fi SoftAP telemetrisi ve alarm zili ile lamba arızalarını anında yetkililere bildiren kompakt IoT cihazı.",
		P11Domain: "Akıllı Şehir IoT · Gömülü Yazılım ve Telemetri",
		P12Title: "NOTCHA — Minimal Dinamik Kayan Zaman Çentiği",
		P12Desc: "Windows 11'in görev çubuğunu ekranın üstüne yerleştirme seçeneğini kaldırmasının ardından 2021/2022'de ergonomik bir çözüm olarak geliştirilen NOTCHA; gerçek zamanlı saat göstergesi, şeffaf üst kenar çentiği ve akıllı fareyle üzerine gelince açılma (hover) mekanizması sunan hafif bir masaüstü aracıdır.",
		P12Domain: "Masaüstü Ergonomisi · Özel Arayüz Araçları",
		P13Title: "SPO2 — Nabız Oksimetresi 2-Katmanlı PCB Tasarımı",
		P13Desc: "KiCad 10.0 ile geliştirilen 2 katmanlı nabız oksimetresi devre kartı; sıfırdan oluşturulan özel ayak izleri, semboller, üretici toleransları, MAX30102 optik biyosensörü, SSD1306 OLED ekranı ve Arduino Nano çekirdeğini bir araya getiriyor.",
		P13Domain: "PCB Tasarımı ve EDA · Biyomedikal Donanım · KiCad 10.0",
		P14Title: "RESPIRO",
		P14Desc: "Acil durum konuşlandırması ve kaynak kısıtlı klinikler için tasarlanmış düşük maliyetli, yüksek güvenilirlikli biyomedikal uyarlanabilir yardımcı ventilatör.",
		P14Domain: "Biyomedikal Mekatronik · Yoğun Bakım",
		AboutTitle: "Ukacha Juma Hakkında",
		AboutSubtitle: "Mekatronik Mühendisi · Donanım-Yazılım Mimarı · Poliglot",
		AboutP1: "Modern donanım mühendisliğinin acıyı sadece tedavi etmekle kalmayıp önlemesi gerektiği inancıyla çalışan bir Mekatronik Mühendisiyim. Çalışmalarım mekanik CAD/CAM tasarımını, hassas analog sensör devrelerini, bare-metal ve FreeRTOS mikrodenetleyici ürün yazılımlarını ve gerçek zamanlı telemetri sistemlerini birleştirmektedir.",
		AboutP2: "Yatak yarası önleyici matlardan (VELUM), mikroplastik filtrelemeye (FLOWGUARD) ve termal biyosensörlere (SOLUS) kadar, fiziksel sistemleri matematiksel prensiplerden laboratuvar doğrulamasına ve üretime hazır tasarıma kadar sıfırdan inşa ediyorum.",
		AboutPolyglot: "Çok Dilli İletişim & Küresel İşbirliği",
		AboutPolyglotDesc: "İngilizce, Türkçe, Arapça, Portekizce, İspanyolca, Almanca ve Fransızca dillerini akıcı bir şekilde konuşuyorum. Poliglot olmak, kıtalararası teknik ekipler arasında köprü kurmamı ve küresel klinik literatürü orijinal kaynağından takip etmemi sağlıyor.",
		SkillsTitle: "Teknik Yetkinlikler & Araç Takımı",
		ExpTitle: "Mesleki Deneyim", ExpSubtitle: "Uygulamalı endüstriyel mühendislik, CNC arıza tespiti, otomasyon bakımı ve uluslararası tercümanlık.", EduTitle: "Eğitim ve Akademik Geçmiş", AwardsTag: "AKADEMİK VE ROBOTİK BAŞARI", AwardsTitle: "Ödüller ve Başarılar", AwardsSubtitle: "Fakülte dekanlığı onur plaketleri, ulusal robot olimpiyatları dereceleri ve kurumsal mühendislik takdirleri.", AwardsCardTitle: "İTÜRO 2026 Derecesi — Teknoloji Fakültesi Dekanı Onur Plaketi", AwardsCardDesc: "Selçuk Üniversitesi Teknoloji Fakültesi Dekanımız Sayın Prof. Dr. Humar Kahramanlı Örnek tarafından, STWolf takımımızın İstanbul Teknik Üniversitesi Robot Olimpiyatları (İTÜRO 2026)'da elde ettiği üstün başarı ve derece vesilesiyle takdim edilen onur plaketi.", AwardsViewPost: "LinkedIn Paylaşımını ve Duyuruyu Gör", AwardsDeanTitle: "Teknoloji Fakültesi Dekanı", AwardsDeanName: "Prof. Dr. Humar Kahramanlı Örnek", AwardsTeamName: "STWolf Takımı · Selçuk Üniversitesi", AwardsOlympiad: "İstanbul Teknik Üniversitesi Robot Olimpiyatları (İTÜRO 2026)", AwardsCeremony: "2025–26 Mezuniyet Töreni · Selçuk Üniversitesi", AwardsQuoteLabel: "Resmi Plaket Metni", TabEducation: "Eğitim", TabAwards: "Ödüller & Başarılar", TabCertificates: "Sertifikalar & Lisanslar", TabActivities: "Çalıştaylar", TabLanguages: "Diller", ActivitiesTitle: "Mühendislik Etkinlikleri & Çalıştaylar", ActivitiesSubtitle: "Teknoloji festivali sergileri, topluluk etkileşimi ve akademik yapay zeka araştırma yöntemleri.", TeknoselTitle: "Selçuklu Teknosel Teknoloji Festivali — Üniversite Kategorisi", TeknoselDesc: "Serbest Proje - Üniversite Seviyesi kategorisinde standımızda KonyaWaste akıllı atık mekatronik prototipimizi sergiledik; mühendis meslektaşlarımızla buluşup Selçuklu Belediye Başkanı Sayın Ahmet Pekyatırmacı'ya projemizi bizzat anlattık.", TeknoselMayorBadge: "Selçuklu Belediye Başkanı'na Birebir Sunum", AiResearchTitle: "Akademik Araştırmada Kaynak Odaklı Yapay Zeka (TADD Çalıştayı)", AiResearchDesc: "TADD ve eğitmenimiz Yunus Emre Kıymaz liderliğinde 2 günlük yoğun çalıştayı tamamlayarak; Zotero, ResearchRabbit ve NotebookLM ile yapay zeka halüsinasyonlarını ortadan kaldıran ve mühendislik literatür taramalarında teknik doğruluğu garanti eden kaynak odaklı araştırma metodolojisinde yetkinlik kazandım.", AiResearchBadge: "Halüsinasyonsuz Kaynak Odaklı Araştırma", IfmWorkshopTitle: "Otomasyon Yetkinliklerimi ifm electronic ve CoDeSys ile Genişletiyorum (48 Saatlik Yoğun Eğitim)", IfmWorkshopDesc: "ifm electronic Türkiye İstanbul genel merkezi tarafından Selçuk Üniversitemizde düzenlenen 48 saatlik yoğunlaştırılmış özel PLC eğitim programını başarıyla tamamladım. PLC mimarisi, CoDeSys V3.5 (SP18 ve SP19), Modbus ve CANbus haberleşme protokolleri ve sahada sorun giderme konularında uygulamalı yetkinlik kazandım.", IfmWorkshopBadge: "⚙️ ifm electronic Türkiye · Selçuk Üniversitesi", IfmBtnLinkedIn: "ifm Çalıştayı LinkedIn Paylaşımını Gör", TabSoloProjects: "Bireysel Projeler", TabTeamProjects: "Takım Projeleri", TabContractProjects: "Sözleşmeli Projeler", ContractEmptyBadge: "Sözleşmeli Çalışmaya Açık", ContractEmptyTitle: "Sözleşmeli ve Müşteri Mühendislik Projeleri", ContractEmptyDesc: "\u00d6zel ticari s\u00f6zle\u015fmeler, donan\u0131m dan\u0131\u015fmanl\u0131\u011f\u0131 ve end\u00fcstriyel mekatronik \u00e7\u00f6z\u00fcmleri. Portf\u00f6y detaylar\u0131 yak\u0131nda eklenecektir.",
		CP1Title: "Mec\u00e2nica Auto Moderna \u2014 Kurumsal Logo ve Marka Tasar\u0131m\u0131", CP1Desc: "Mec\u00e2nica Auto Moderna oto servis firmas\u0131 i\u00e7in \u00f6zel olarak tasarlanm\u0131\u015f kurumsal logo ve g\u00f6rsel marka kimli\u011fi. Tabela, i\u015f k\u0131yafetleri ve dijital mecralar i\u00e7in hassas tamir ara\u00e7lar\u0131, 24 saat kesintisiz hizmet sembolizmi ve dinamik tipografi i\u00e7eren \u00f6zg\u00fcn amblem.", CP1Domain: "Kurumsal Kimlik ve Vekt\u00f6rel Tasar\u0131m",
		CP2Title: "Mec\u00e2nica Auto Moderna \u2014 Invoice Moderna Fatura Sistemi", CP2Desc: "Mec\u00e2nica Auto Moderna oto bak\u0131m ve onar\u0131m i\u015fletmesi i\u00e7in geli\u015ftirilen \u00f6zel yaz\u0131l\u0131m \u00e7\u00f6z\u00fcm\u00fc. Elektronik fatura (Fatura Proforma) kesimi, KDV hesaplamalar\u0131, m\u00fc\u015fteri kayd\u0131 ve i\u015f emri s\u00fcre\u00e7lerini otomatikle\u015ftirerek at\u00f6lye operasyonel verimlili\u011fini ve hassasiyetini en \u00fcst d\u00fczeye \u00e7\u0131kar\u0131r.", CP2Domain: "Kurumsal Fatura ve \u0130\u015fletme Y\u00f6netim Yaz\u0131l\u0131m\u0131",
		BtnInstagramPost: "Instagram Payla\u015f\u0131m\u0131n\u0131 G\u00f6r", BtnLinkedInPost: "LinkedIn Payla\\u015f\\u0131m\\u0131n\\u0131 G\\u00f6r",
		CP3Title: "APACOV \u2014 BT Altyap\u0131s\u0131 ve Web Mimarisi", CP3Desc: "Yetim ve dezavantajl\u0131 \u00e7ocuklar\u0131 koruma ve bar\u0131nd\u0131rma misyonu \u00fcstlenen APACOV insani yard\u0131m vakf\u0131 i\u00e7in g\u00f6n\u00fcll\u00fc olarak geli\u015ftirilen kurumsal web portal\u0131 ve u\u00e7tan uca bili\u015fim altyap\u0131s\u0131. Sunucu y\u00f6netimi, dijital g\u00fcvenlik ve d\u0131\u015f ili\u015fkiler sorumlulu\u011fu \u00fcstlenilmektedir.", CP3Domain: "BT Altyap\u0131s\u0131 \u00b7 Web Mimarisi \u00b7 D\u0131\u015f \u0130li\u015fkiler",
		CP4Title: "APACOV \u2014 \u0130nsani Yard\u0131m Logosu ve Marka Kimli\u011fi", CP4Desc: "APACOV insani yard\u0131m derne\u011fi i\u00e7in pro bono olarak tasarlanan kurumsal logo ve g\u00f6rsel kimlik. Adobe Illustrator ve Samsung PenUp ile modellenen amblem, koruyucu bir yay i\u00e7erisinde birbirine uzanan dayan\u0131\u015fma ve umut ellerini temsil eder.", CP4Domain: "\u0130nsani Marka Kimli\u011fi \u00b7 Vekt\u00f6rel Tasar\u0131m",
		CP5Title: "SuPublic — Otomatik İçme Suyu Dolum İstasyonu ve Kiosk Sac Metal Tasarımı", CP5Desc: "Gonca Özata tarafından kurulan ve Dr. Yavuz Selim Taşpınar danışmanlığında sürdürülen SuPublic girişimi için sözleşmeli (contract) makine mühendisliği ve SolidWorks sac metal tasarımı. Konya'da başlayıp tüm Türkiye ve dünyadaki ücretsiz temiz içme suyu çeşmelerini haritalandırarak tek kullanımlık plastik şişe atıklarını önlemeyi amaçlayan proje kapsamında; otomatik su dolum cihazı ve kiosk prototipinin tüm sac metal gövde, büküm açılımları, lazer kesim detayları ve ergonomik matara dolum haznesi tasarlandı.", CP5Domain: "Sözleşmeli Mühendislik · Sac Metal Tasarımı · SolidWorks CAD ve Prototipleme",
		CPNoticeInfra: "Sunucu Ge\u00e7i\u015f Bildirimi: apacov.org i\u00e7in kendi sunucular\u0131m\u0131za altyap\u0131 ta\u015f\u0131ma s\u00fcreci, ek kurumsal alan adlar\u0131 ve e-posta yap\u0131land\u0131rmalar\u0131 aktif olarak devam etmektedir.", BtnOfficialWebsite: "Resmi Web Sitesi (apacov.org)", BadgeVolunteered: "G\u00f6n\u00fcll\u00fc \u00b7 Pro Bono",
		ExcavatorTitle: "Ekskavatör Tasarımı (Makine Elemanları Dersi) — Komatsu PC200 LC-7/LC-8",
		ExcavatorDesc: "24 ton sınıfı Komatsu PC200 LC-7/LC-8 ekskavatörden esinlenilerek Makine Elemanları dersi kapsamında gerçekleştirilen kapsamlı mekanik tasarım ve sonlu elemanlar analizi (FEA). Bom, arm ve kova milleri (yük, gerilme, sehim, yorulma), kule dönüş ve alt takım rulmanları (makaralar, istikamet tekeri ömür hesapları), kritik civata grupları mukavemet hesapları, tork gereksinimlerine uygun dişli mekanizmaları ve kazma/kaldırma kuvvetlerini optimize eden hidrolik silindir ve devre hesaplamaları ANSYS ile doğrulanmıştır.",
		ExcavatorDomain: "İş Makineleri Tasarımı · Makine Elemanları · ANSYS FEA & Hidrolik",
		SafrenTitle: "SafRen (Eski Adıyla Ce-Saf) — Akıllı Ayakkabı Adım Sayar ve Giyilebilir IoT Sistemi",
		SafrenDesc: "Ayakkabı Tasarımı bölümü öğrencileri Safina Najmudinova ve Ceren Başaran ile disiplinler arası ortaklıkla geliştirilen akıllı giyilebilir ayakkabı adım sayar ve adım telemetrisi sistemi. Prototip ayakkabıları temin eden ve donanımların (OLED LCD ekran, butonlar, yaylı zemin temas anahtarı, mikrodenetleyici) ayakkabı üzerindeki en ergonomik ve optimum yerleşim noktalarını belirleyen Safina ve Ceren Başaran ile birlikte hayata geçirildi. Waveshare ESP32-S3-Zero mikrodenetleyici, MPU6050 6-eksen IMU sensörü, yerleşik SoftAP Wi-Fi HTTP sunucusu ve Flutter mobil uygulaması ile gerçek zamanlı kablosuz adım senkronizasyonu sunar.",
		SafrenDomain: "Giyilebilir IoT · Adım Analizi · ESP32-S3 Yazılımı · Flutter Mobil",
		STM32Title: "STM32 Discovery Kartı Mücadelesi — Otonom Mobil Robotik",
		STM32Desc: "STM32F407VGT6 ARM Cortex-M4 Discovery kartı ile geliştirilen iki aşamalı otonom mobil robotik projesi. Vize Mücadelesi: 1 m² pistte rastgele konumlara yerleştirilen ışıklara çok kanallı LDR sensör dizilimi, 12-bit ADC okumaları ve TIM3 PWM diferansiyel motor kontrolü ile yönelip ilk denemede çarpmadan milimetrik mesafede durma. Final Mücadelesi: Aynı pistte 5 adet mumu 30 saniye içinde bulup söndürme görevi; sahada merkeze konumlanma, 12V fan entegrasyonu ve 360 derece kesintisiz girdap dönüş algoritması ile tüm mumları 10 saniyenin altında (verilen sürenin 3'te 1'inde) başarıyla söndürme.",
		STM32Domain: "Gömülü Sistemler · STM32 ARM Cortex-M4 · Bare-Metal C / HAL · Otonom Robotik",
		ROSTitle: "ROS Noetic Otonom Mobil Robot — ESP32 TCP/IP Köprüsü ve Diferansiyel Kinematik",
		ROSDesc: "Mehmet Çakaroğlu ile ortaklaşa ve Dr. Öğr. Üyesi Okan Uyar danışmanlığında geliştirilen otonom mobil robot projesi. Linux üzerinde koşan ROS Noetic ile ESP32 çift çekirdekli mikrodenetleyici arasında düşük gecikmeli, yüksek bant genişlikli Wi-Fi TCP/IP soket köprüsü kurar. Diferansiyel sürüş kinematiği, tekerlek başına 420 tick manyetik enkoderler ile kapalı çevrim odometri, kalibre edilmiş MPU6050 6-eksen IMU, 500 ms güvenlik zamanlayıcısı (deadman switch) ve Gazebo/RViz dijital ikiz simülasyonu ile dairesel yörünge ve git-dön görevlerinde 100/100 tam puan kriterlerini sağlamıştır.",
		ROSDomain: "Otonom Mobil Robotik · ROS Noetic · Diferansiyel Kinematik · Gazebo ve RViz",
		BtnLinkedInVideo: "LinkedIn Video Sunumu", BtnGitHubRepo: "GitHub Deposu", CertsTitle: "Sertifikalar ve Başarılar",
		VHeroBadge: "James Dyson Ödülü 2026 Başvurusu · Sağlık İnovasyonu",
		VHeroTitle: "VELUM — Yatak Yarası Önleme Matı",
		VHeroSubtitle: "Yatak yaralarını gözle görülür hale gelmeden 2–5 gün önce tespit eden dünyanın ilk sürekli, pasif, adil ve uygun maliyetli biyomedikal sensör matı.",
		VStatCost: "11 Milyar $+ USD", VStatCostDesc: "Önlenebilir yatak yaralarının ABD hastanelerine yıllık maliyeti",
		VStatEarly: "2–5 Gün Erken", VStatEarlyDesc: "Sub-Epidermal Nem değişimleri ile görsel öncesi erken uyarı",
		VStatBOM: "< 50 $ USD", VStatBOMDesc: "Her hasta yatağına uygulanabilir hedef parça maliyeti",
		VStatEquity: "%100 Eşitlikçi", VStatEquityDesc: "Ten renginden bağımsız biyoempedans algılama (Fitzpatrick I–VI)",
		VProblemTitle: "Hastalara Zarar Veren Yetersiz Mevcut Sistem",
		VProblemSubtitle: "Görsel inceleme, yaraları ancak doku nekrozu geri döndürülemez olduktan sonra fark eder.",
		VProblemP1: "Yatak yaraları dünya genelinde yatan hastaların %12.8'inden fazlasını etkilemektedir. Mevcut bakım standardı, hemşirelerin 2 saatte bir hastayı manuel çevirmesine ve öznel görsel cilt kontrollerine dayanmaktadır. Ancak hücre hasarı, yüzeyde kızarıklık oluşmadan 48 ila 120 saat önce derinin alt katmanlarında başlar.",
		VProblemP2: "Ayrıca görsel kızarıklık değerlendirmesi koyu ten rengine sahip hastalarda (Fitzpatrick Tip IV–VI) klinik olarak yetersiz kalmaktadır. Ticari el tipi SEM tarayıcıları 3.000–10.000 $ maliyetindedir ve manuel personel gerektirir.",
		VProblemPoint1: "Tepkisel Yapı: Cilt kızarıklığı, hücre ölümünün zaten başladığını gösterir.",
		VProblemPoint2: "Belgelenmiş Etnik Eşitsizlik: Görsel kontrol koyu tende 1. Aşama yaraları kaçırır.",
		VProblemPoint3: "Maliyet & Personel Yükü: Pahalı cihazlar hemşire yorgunluğunu çözemez.",
		VSolutionTitle: "VELUM Çözümü: Sürekli Biyoempedans",
		VSolutionSubtitle: "Çarşafın altında görünmez, 7/24 pasif lokalize doku sağlığı takibi.",
		VSolutionP1: "VELUM, yatalak hastaların altına yerleştirilen ultra ince, esnek bir poliimid sensör matıdır. Serigrafi baskılı gümüş-karbon elektrot matrisi, her 5 dakikada bir 16 kritik anatomik bölgede biyoelektriksel empedans spektroskopisi (1–100 kHz) gerçekleştirir.",
		VSolutionPoint1: "Kurulum Gerektirmez: Çarşafın altına koyun, çalıştırın, hasta konforu bozulmaz.",
		VSolutionPoint2: "Sub-Epidermal Nem Tespiti: Empedanstaki %15+ düşüş, yaradan günler önce sıvı birikimini haber verir.",
		VSolutionPoint3: "Anlık BLE / Wi-Fi Uyarıları: Hemşireler hastanın tam olarak hangi bölgesinin çevrilmesi gerektiğini anında görür.",
		VSimTitle: "Etkileşimli Biyoempedans Mat Simülatörü",
		VSimSubtitle: "Gerçek zamanlı 16 bölgeli anatomik basınç haritası ve AD5933 empedans tarama simülasyonu.",
		VSimNormal: "1. Normal Durum", VSimIschemia: "2. Sakral Basınç (İskemi)",
		VSimOedema: "3. Aşama 0 Ödem Uyarısı", VSimReset: "4. Hasta Pozisyonu Değiştirme",
		VSimStatusAlert: "KRİTİK: Doku Altı Ödem Tespit Edildi! Sakrum Bölgesini Derhal Çevirin",
		VSimStatusNormal: "NOMİNAL: Tüm anatomik bölgelerde doku perfüzyonu ideal seviyede",
		VTechTitle: "Patlatılmış Mekatronik Mimari",
		VTechSubtitle: "Klinik dezenfeksiyon ve seri üretim için tasarlanmış hassas esnek elektronik.",
		VLayer1Title: "1. Medikal TPU & Platin Silikon Kapsülleme",
		VLayer1Desc: "Mikroklimayı koruyan, hastane sınıfı dezenfektanlara dayanıklı biyouyumlu membran (ISO 10993).",
		VLayer2Title: "2. Baskılı Gümüş/Karbon Elektrot Matrisi",
		VLayer2Desc: "Kemik çıkıntıları altına yerleştirilmiş tetrapolar biyoempedans dizilimli 50 µm DuPont Kapton® poliimid film.",
		VLayer3Title: "3. AD5933 Empedans Dönüştürücü & 16-Kanallı MUX",
		VLayer3Desc: "Analog Devices AD5933 ve Howland akım kaynağı (<1 mA AC) ile saniyeler içinde tüm bölgeleri tarayan CD74HC4067 çoklayıcı.",
		VLayer4Title: "4. ESP32-S3 Çift Çekirdekli Hesaplama & Kablosuz Modül",
		VLayer4Desc: "FreeRTOS üzerinde temel çizgi kalibrasyonu, BLE 5.0 mesh ve hastane EMR şifreli Wi-Fi telemetrisi yürüten 240MHz işlemci.",
		VCompTitle: "Klinik & Rekabet Üstünlüğü",
		VCompSubtitle: "VELUM'un benzersiz klinik ve ekonomik avantajları.",
		VCompColMetric: "Kriter / Özellik", VCompColVisual: "Görsel Kontrol", VCompColSEM: "El Tipi SEM Probu", VCompColMattress: "Basınç Yatakları", VCompColVelum: "VELUM Mat",
		VRoadmapTitle: "Klinik & Ticari Yol Haritası",
		VRoadmapSubtitle: "Laboratuvar prototipinden yoğun bakım klinik denemelerine ve FDA 510(k) onayına.",
		VGalleryTitle: "Donanım Renderları & Tasarım Çizimleri",
		VGallerySubtitle: "Yüksek çözünürlüklü ürün görselleri ve mekatronik patlatılmış montaj diyagramları.",
		VDownloadPitch: "Yatırımcı Sunumunu İndir (PDF)",
		VViewGitHub: "GitHub Deposunu Görüntüle",
		FooterRights: "Tüm hakları saklıdır.",
		FooterBuiltWith: "",
	},
	"ar": {
		LangCode: "ar", LangName: "Arabic", IsRTL: true, Dir: "rtl",
		NavBrand: "أوكاشا جمعة",
		NavHome: "الرئيسية", NavProjects: "المشاريع", NavExperience: "الخبرات", NavAbout: "عني", NavSkills: "المهارات", NavContact: "اتصل بي", NavEducation: "التعليم", NavHonors: "الجوائز", NavCertificates: "الشهادات", NavActivities: "ورش العمل", NavLanguages: "اللغات", NavGitHub: "جيت هاب", NavVelum: "مشروع فيلوم",
		HeroBadge: "مهندس ميكاترونكس · أجهزة طبية حيوية · أنظمة مدمجة · متعدد اللغات",
		AuthorRole: "بكالوريوس هندسة الميكاترونكس · مطور الأنظمة المدمجة وإنترنت الأشياء",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "تصميم أنظمة ميكاترونكس طبية ذكية، وأجهزة استشعار حيوية دقيقة، وتقنيات إنترنت الأشياء الطرفية لحماية حياة المرضى.",
		HeroBtnExplore: "استكشاف المشاريع", HeroBtnContact: "تواصل معي",
		StatProjects: "١٨ مشروعاً هندسياً وبرمجياً", StatSoloProjects: "المشاريع الفردية", StatTeamProjects: "مشاريع الفرق", StatContractProjects: "المشاريع التعاقدية", StatWorkshops: "ورش العمل والأنشطة", StatAward: "هندسة الميكاترونكس", StatLanguages: "٧ لغات متقنة", StatFocus: "ESP32 / المعاوقة الحيوية / CAD",
		ProjectsTitle: "معرض المشاريع الهندسية",
		ProjectsSubtitle: "أجهزة طبية وقائية، أنظمة مساعدة، أتمتة موائع، وتطبيقات ميكاترونكس متقدمة.",
		StatusActive: "مكتمل وتفاعلي", StatusQueue: "قيد التطوير",
		BtnViewProject: "عرض المشروع", BtnPreview: "تفاصيل المشروع",
		BadgeFeatured: "مشروع رئيسي", BadgeBiomedical: "أجهزة طبية حيوية", BadgeHardware: "ميكاترونكس", BadgeRobotics: "روبوتات وأتمتة",
		P1Title: "فيلوم - VELUM (سابقاً SentIO)",
		P1Desc: "بساط استشعار مرن يقيس الرطوبة تحت الجلد والمعاوقة الحيوية الكهربائية على مدار الساعة لاكتشاف قرح الفراش قبل ظهورها بـ ٢-٥ أيام.",
		P1Domain: "ميكاترونكس طبية حيوية · مطيافية المعاوقة الحيوية",
		P2Title: "فلوجارد - FLOWGUARD",
		P2Desc: "نظام ترشيح ميكاترونيكي ذكي لغسالات الملابس يمنع وصول الجزيئات البلاستيكية الدقيقة الاصطناعية إلى مياه المحيطات.",
		P2Domain: "ميكاترونكس الموائع · فصل البلاستيك الدقيق",
		P3Title: "سولوس - SOLUS",
		P3Desc: "لاصقة طبية حيوية قابلة للارتداء تقيس درجة الحرارة الأساسية باستمرار للإنذار المبكر من ضربات الشمس في البيئات القاسية.",
		P3Domain: "حساسات حيوية قابلة للارتداء · وقاية فسيولوجية",
		P4Title: "دينوفا - DENOVA (سابقاً Avicenna)",
		P4Desc: "كاشف ضوئي غير جراحي لتسوس الأسنان يعتمد على الإضاءة متعددة الأطياف لتشخيص نقص المعادن في مراحله المبكرة جداً.",
		P4Domain: "فوتونيات حيوية · تشخيص بصري",
		P5Title: "رينوفو - RENOVO",
		P5Desc: "مكبس هيدروليكي محمول لإغاثة الكوارث يحول ركام المباني المنهارة في الزلازل إلى قوالب بناء متشابكة في الموقع.",
		P5Domain: "ميكانيكا ثقيلة وهيدروليك · إغاثة الكوارث",
		P6Title: "كورا سيرو - CURA SERO",
		P6Desc: "شريط حيوي لاصق لإغلاق الجروح الجراحية الدقيقة يوزع قوى الشد بانتظام ويمنع تشكل الندبات الجلدية المشوهة.",
		P6Domain: "أجهزة ميكانيكية حيوية · جراحة متقدمة",
		P7Title: "ويلستيك - WELSTECH",
		P7Desc: "سوار معصم طبي حيوي فائق الصغر لمراقبة النبض والعلامات الحيوية عبر مستشعر PPG ومعالج ESP32-S3 مع نظام إنذار صوتي وتطبيق فلاتر.",
		P7Domain: "إنترنت الأشياء الطبي · استشعار النبض الحيوي",
		P8Title: "سيند إت - SENDIT",
		P8Desc: "آلية لوجستية ذاتية الحركة لفرز ونقل الحقائب والبريد السريع في المطارات لتقليل أوقات المناولة الجوية.",
		P8Domain: "روبوتات ذاتية الحركة · أتمتة لوجستية",
		P9Title: "سمارت ويست - SMARTWASTE",
		P9Desc: "نظام فرز نفايات ميكاترونيكي آلي يدمج المطيافية الضوئية والرؤية الحاسوبية لفرز النفايات الحضرية تلقائياً.",
		P9Domain: "رؤية حاسوبية · مشغلات ميكاترونيكية",
		P10Title: "مصنف الصور - شبكة عصبية تلافيفية عميقة (CNN)",
		P10Desc: "نموذج رؤية حاسوبية مدرب عبر TensorFlow/Keras على أكثر من ٩٦٠٠ صورة مخصصة بدقة ٩٦×٩٦ مع واجهة مستخدم رسومية لتصنيف الصور لحظياً.",
		P10Domain: "التعلم العميق · الرؤية الحاسوبية والذكاء الاصطناعي",
		P11Title: "LUMIERE — نظام ذكي للإبلاغ عن أعطال إنارة الشوارع وتحديد المواقع",
		P11Desc: "جهاز إنترنت أشياء (IoT) مدمج يُثبت تحت أعمدة الإنارة البلدية، يجمع بين مستشعر المقاوم الضوئي LDR وتحديد الموقع بدقة عبر GPS ونقطة اتصال ESP32 Wi-Fi لإرسال تنبيهات فورية بالانقطاع للبلديات.",
		P11Domain: "إنترنت الأشياء للمدن الذكية · الأنظمة المدمجة والاتصالات",
		P12Title: "نوتشا - NOTCHA: شريط الوقت الديناميكي العائم المينيمالي",
		P12Desc: "تم تطويره كحل هندسي إرجونومي في 2021/2022 بعد أن ألغت مايكروسوفت إمكانية تثبيت شريط مهام ويندوز في أعلى الشاشة؛ يقدم NOTCHA شريط وقت عائم أنيق أعلى الحافة بشفافية تامة، تحديثات فورية وخاصية الإظهار التلقائي عند تمرير المؤشر دون استهلاك للموارد.",
		P12Domain: "إرجونوميا سطح المكتب · أدوات واجهة المستخدم المخصصة",
		P13Title: "SPO2 — تصميم لوحة دوائر مطبوعة ثنائية الطبقات لمقياس التأكسج",
		P13Desc: "لوحة دوائر مطبوعة ثنائية الطبقات لمقياس التأكسج النبضي مصممة ببرنامج KiCad 10.0، تجمع بين رموز ومخططات أقدام مخصصة من الصفر، وتفاوتات تصنيع حقيقية، ومستشعر MAX30102 وشاشة SSD1306 ونواة Arduino Nano.",
		P13Domain: "تصميم اللوحات المطبوعة · الأجهزة الطبية الحيوية · KiCad 10.0",
		P14Title: "RESPIRO",
		P14Desc: "جهاز تنفس اصطناعي مساعد تكيفي منخفض التكلفة وعالي الموثوقية مصمم للنشر في حالات الطوارئ والعيادات محدودة الموارد.",
		P14Domain: "ميكاترونكس طبية حيوية · العناية المركزة",
		AboutTitle: "نبذة عن أوكاشا جمعة",
		AboutSubtitle: "مهندس ميكاترونكس · مهندس نظم العتاد والبرمجيات · متعدد اللغات",
		AboutP1: "أنا مهندس ميكاترونكس أؤمن بأن الهندسة الحديثة يجب أن تكرس لمنع المعاناة الإنسانية واستباقها بدلاً من الاكتفاء بعلاجها. تجمع خبرتي بين التصميم الميكانيكي CAD/CAM، وتصميم الدوائر التناظرية الدقيقة، وبرمجة الأنظمة المدمجة مع FreeRTOS، وأنظمة الاتصال اللاسلكي الفوري.",
		AboutP2: "من خلال مشاريع مبتكرة مثل بساط الوقاية من قرح الفراش (VELUM)، وترشيح البلاستيك الدقيق (FLOWGUARD)، والحساسات الفسيولوجية (SOLUS)، أقوم ببناء النظم الفيزيائية المتكاملة من النماذج الحسابية إلى النماذج المختبرية والتصنيع الفعلي.",
		AboutPolyglot: "تواصل عالمي وطلاقة متعددة اللغات",
		AboutPolyglotDesc: "أتحدث بطلاقة ٧ لغات: الإنجليزية، التركية، العربية، البرتغالية، الإسبانية، الألمانية، والفرنسية. هذا التنوع يمنحني القدرة على قيادة الفرق الهندسية الدولية، وقراءة الأبحاث الطبية بلغاتها الأصلية، وتصميم تقنيات تصل للعالم بأسره.",
		SkillsTitle: "المهارات الفنية والأدوات الهندسية",
		ExpTitle: "الخبرة المهنية والعملية", ExpSubtitle: "هندسة صناعية تطبيقية، صيانة أنظمة CNC، التحكم الآلي، والترجمة الدولية.", EduTitle: "التعليم والخلفية الأكاديمية", AwardsTag: "تقدير أكاديمي وروبوتي", AwardsTitle: "الجوائز والتكريمات", AwardsSubtitle: "أوسمة التقدير الأكاديمي، وتصنيفات أولمبياد الروبوتات الوطنية، والتفوق الهندسي المؤسسي.", AwardsCardTitle: "تصنيف أولمبياد إيتورو 2026 — درع تكريم عميد كلية التكنولوجيا", AwardsCardDesc: "التكريم بدرع التميز المقدم من عميدة كلية التكنولوجيا بجامعة سلجوق، أ.د. هومار كهرمانلي أورنيك، خلال حفل التخرج (2025-26)، تقديراً لتصنيف فريق STWolf في أولمبياد الروبوت بجامعة إسطنبول التقنية (İTÜRO 2026).", AwardsViewPost: "عرض منشور التكريم على لينكد إن", AwardsDeanTitle: "عميدة كلية التكنولوجيا", AwardsDeanName: "أ.د. هومار كهرمانلي أورنيك", AwardsTeamName: "فريق STWolf · جامعة سلجوق", AwardsOlympiad: "أولمبياد الروبوت بجامعة إسطنبول التقنية (İTÜRO 2026)", AwardsCeremony: "حفل التخرج الأكاديمي 2025–26 · جامعة سلجوق", AwardsQuoteLabel: "النص الرسمي للدرع التكريمي", TabEducation: "التعليم الأكاديمي", TabAwards: "الجوائز والتكريمات", TabCertificates: "الشهادات والتراخيص", TabActivities: "ورش العمل", TabLanguages: "اللغات", ActivitiesTitle: "الأنشطة الهندسية وورش العمل", ActivitiesSubtitle: "معارض المهرجانات التكنولوجية، والمشاركة المجتمعية، وأساليب البحث بالذكاء الاصطناعي.", TeknoselTitle: "مهرجان سلجوقلو تكنوسيل للتكنولوجيا — الفئة الجامعية", TeknoselDesc: "عرض النموذج الأولي لمشروع إدارة النفايات الذكية KonyaWaste في جناحنا ضمن فئة المشاريع الحرة الجامعية، وتقديم شرح تفصيلي ومباشر لرئيس بلدية سلجوقلو.", TeknoselMayorBadge: "عرض مباشر أمام رئيس بلدية سلجوقلو", AiResearchTitle: "الذكاء الاصطناعي القائم على المصادر في البحث الأكاديمي (TADD)", AiResearchDesc: "إتمام ورشة عمل مكثفة نظمتها جمعية أصدقاء إفريقيا (TADD) بقيادة يونس إمري قيماز، لإتقان أدوات Zotero وResearchRabbit وNotebookLM للقضاء على هلوسة الذكاء الاصطناعي وضمان الدقة العلمية للأبحاث الهندسية.", AiResearchBadge: "بحث موثوق خالي من هلوسة الذكاء الاصطناعي", IfmWorkshopTitle: "تطوير مهارات الأتمتة الصناعية مع ifm electronic و CoDeSys (تدريب مكثف 48 ساعة)", IfmWorkshopDesc: "إتمام برنامج تدريبي مكثف لمدة 48 ساعة في أنظمة التحكم المنطقي المبرمج (PLC) أقامته شركة ifm electronic التركية في جامعة سلجوق. خبرة عملية في بنية الـ PLC وبرمجة CoDeSys V3.5 وبروتوكولات Modbus و CANbus ومعالجة الأعطال الميدانية.", IfmWorkshopBadge: "⚙️ شركة ifm electronic تركيا · جامعة سلجوق", IfmBtnLinkedIn: "عرض منشور ورشة ifm على لينكد إن", TabSoloProjects: "المشاريع الفردية", TabTeamProjects: "مشاريع الفرق", TabContractProjects: "المشاريع التعاقدية", ContractEmptyBadge: "متاح للتعاقد والاستشارات", ContractEmptyTitle: "مشاريع العقود والاستشارات الهندسية", ContractEmptyDesc: "\u0639\u0642\u0648\u062f \u062a\u062c\u0627\u0631\u064a\u0629 \u062e\u0627\u0635\u0629\u060c \u0627\u0633\u062a\u0634\u0627\u0631\u0627\u062a \u0641\u064a \u0627\u0644\u0639\u062a\u0627\u062f \u0648\u0627\u0644\u0623\u0646\u0638\u0645\u0629 \u0627\u0644\u0645\u062f\u0645\u062c\u0629\u060c \u0648\u062d\u0644\u0648\u0644 \u0647\u0646\u062f\u0633\u064a\u0629 \u0635\u0646\u0627\u0639\u064a\u0629. \u0633\u062a\u062a\u0645 \u0625\u0636\u0627\u0641\u0629 \u062f\u0631\u0627\u0633\u0627\u062a \u0627\u0644\u062d\u0627\u0644\u0629 \u0642\u0631\u064a\u0628\u0627\u064b.",
		CP1Title: "\u0645\u064a\u0643\u0627\u0646\u064a\u0643\u0627 \u0623\u0648\u062a\u0648 \u0645\u0648\u062f\u064a\u0631\u0646\u0627 \u2014 \u062a\u0635\u0645\u064a\u0645 \u0627\u0644\u0634\u0639\u0627\u0631 \u0648\u0627\u0644\u0647\u0648\u064a\u0629 \u0627\u0644\u0628\u0635\u0631\u064a\u0629", CP1Desc: "\u062a\u0635\u0645\u064a\u0645 \u0634\u0639\u0627\u0631 \u0648\u0647\u0648\u064a\u0629 \u0628\u0635\u0631\u064a\u0629 \u0645\u0624\u0633\u0633\u064a\u0629 \u0645\u062e\u0635\u0635\u0629 \u0644\u0634\u0631\u0643\u0629 \u0645\u064a\u0643\u0627\u0646\u064a\u0643\u0627 \u0623\u0648\u062a\u0648 \u0645\u0648\u062f\u064a\u0631\u0646\u0627 (Mec\u00e2nica Auto Moderna). \u0627\u0628\u062a\u0643\u0627\u0631 \u0634\u0639\u0627\u0631 \u0633\u064a\u0627\u0631\u0627\u062a \u0641\u0631\u064a\u062f \u064a\u062f\u0645\u062c \u0623\u062f\u0648\u0627\u062a \u0627\u0644\u0635\u064a\u0627\u0646\u0629 \u0627\u0644\u062f\u0642\u064a\u0642\u0629 \u0648\u0631\u0645\u0632\u064a\u0629 \u0627\u0644\u062e\u062f\u0645\u0629 \u0639\u0644\u0649 \u0645\u062f\u0627\u0631 24 \u0633\u0627\u0639\u0629 \u0644\u0644\u0648\u062d\u0627\u062a \u0627\u0644\u0625\u0639\u0644\u0627\u0646\u064a\u0629 \u0648\u0627\u0644\u0623\u0632\u064a\u0627\u0621 \u0648\u0627\u0644\u0648\u0633\u0627\u0626\u0637 \u0627\u0644\u0631\u0642\u0645\u064a\u0629.", CP1Domain: "\u0627\u0644\u0647\u0648\u064a\u0629 \u0627\u0644\u0628\u0635\u0631\u064a\u0629 \u0627\u0644\u0645\u0624\u0633\u0633\u064a\u0629 \u0648\u0627\u0644\u062a\u0635\u0645\u064a\u0645 \u0627\u0644\u0645\u062a\u062c\u0647\u064a",
		CP2Title: "\u0645\u064a\u0643\u0627\u0646\u064a\u0643\u0627 \u0623\u0648\u062a\u0648 \u0645\u0648\u062f\u064a\u0631\u0646\u0627 \u2014 \u0646\u0638\u0627\u0645 \u0625\u062f\u0627\u0631\u0629 \u0648\u0641\u0648\u0627\u062a\u064a\u0631 Invoice Moderna", CP2Desc: "\u062d\u0644 \u0628\u0631\u0645\u062c\u064a \u0645\u062e\u0635\u0635 \u062a\u0645 \u062a\u0637\u0648\u064a\u0631\u0647 \u0644\u0648\u0631\u0634\u0629 \u0645\u064a\u0643\u0627\u0646\u064a\u0643\u0627 \u0623\u0648\u062a\u0648 \u0645\u0648\u062f\u064a\u0631\u0646\u0627 \u0644\u0644\u0633\u064a\u0627\u0631\u0627\u062a. \u064a\u0639\u0645\u0644 \u0639\u0644\u0649 \u0623\u062a\u0645\u062a\u0629 \u0625\u0635\u062f\u0627\u0631 \u0627\u0644\u0641\u0648\u0627\u062a\u064a\u0631 \u0627\u0644\u0625\u0644\u0643\u062a\u0631\u0648\u0646\u064a\u0629 \u0627\u0644\u0645\u0628\u062f\u0626\u064a\u0629 \u0648\u062d\u0633\u0627\u0628 \u0627\u0644\u0636\u0631\u0627\u0626\u0628 \u0648\u0625\u062f\u0627\u0631\u0629 \u0637\u0644\u0628\u0627\u062a \u0627\u0644\u0635\u064a\u0627\u0646\u0629 \u0648\u0627\u0644\u0639\u0645\u0644\u0627\u0621 \u0628\u0623\u0639\u0644\u0649 \u062f\u0631\u062c\u0627\u062a \u0627\u0644\u062f\u0642\u0629 \u0648\u0627\u0644\u0641\u0639\u0627\u0644\u064a\u0629 \u0641\u064a \u0633\u064a\u0631 \u0627\u0644\u0639\u0645\u0644.", CP2Domain: "\u0628\u0631\u0645\u062c\u064a\u0627\u062a \u0627\u0644\u0641\u0648\u062a\u0631\u0629 \u0627\u0644\u0625\u0644\u0643\u062a\u0631\u0648\u0646\u064a\u0629 \u0648\u0625\u062f\u0627\u0631\u0629 \u0627\u0644\u0648\u0631\u0634",
		BtnInstagramPost: "\u0639\u0631\u0636 \u0645\u0646\u0634\u0648\u0631 \u0625\u0646\u0633\u062a\u063a\u0631\u0627\u0645", BtnLinkedInPost: "\\u0639\\u0631\\u0636 \\u0645\\u0646\\u0634\\u0648\\u0631 \\u0644\\u064a\\u0646\\u0643\\u062f \\u0625\\u0646",
		CP3Title: "\u0623\u0628\u0627\u0643\u0648\u0641 \u2014 \u0627\u0644\u0628\u0646\u064a\u0629 \u0627\u0644\u062a\u062d\u062a\u064a\u0629 \u0644\u062a\u0643\u0646\u0648\u0644\u0648\u062c\u064a\u0627 \u0627\u0644\u0645\u0639\u0644\u0648\u0645\u0627\u062a \u0648\u0647\u0646\u062f\u0633\u0629 \u0627\u0644\u0648\u064a\u0628", CP3Desc: "\u062a\u0637\u0648\u064a\u0631 \u062a\u0637\u0648\u0639\u064a \u0644\u0644\u0628\u0648\u0627\u0628\u0629 \u0627\u0644\u0625\u0644\u0643\u062a\u0631\u0648\u0646\u064a\u0629 \u0627\u0644\u0645\u0624\u0633\u0633\u064a\u0629 \u0627\u0644\u0643\u0627\u0645\u0644\u0629 \u0648\u0647\u0646\u062f\u0633\u0629 \u0627\u0644\u0628\u0646\u064a\u0629 \u0627\u0644\u062a\u062d\u062a\u064a\u0629 \u0644\u062a\u0642\u0646\u064a\u0629 \u0627\u0644\u0645\u0639\u0644\u0648\u0645\u0627\u062a \u0644\u0645\u0646\u0638\u0645\u0629 APACOV \u0627\u0644\u0625\u0646\u0633\u0627\u0646\u064a\u0629 \u0644\u0631\u0639\u0627\u064a\u0629 \u0627\u0644\u0623\u064a\u062a\u0627\u0645 \u0648\u0627\u0644\u0623\u0637\u0641\u0627\u0644 \u0627\u0644\u0645\u0639\u0631\u0636\u064a\u0646 \u0644\u0644\u062e\u0637\u0631. \u062a\u0648\u0644\u064a \u0625\u062f\u0627\u0631\u0629 \u0627\u0644\u062e\u0648\u0627\u062f\u0645 \u0648\u0627\u0644\u0623\u0645\u0627\u0646 \u0627\u0644\u0631\u0642\u0645\u064a \u0648\u0634\u0624\u0648\u0646 \u0627\u0644\u0639\u0644\u0627\u0642\u0627\u062a \u0627\u0644\u062f\u0648\u0644\u064a\u0629.", CP3Domain: "\u0627\u0644\u0628\u0646\u064a\u0629 \u0627\u0644\u062a\u062d\u062a\u064a\u0629 \u0644\u062a\u0643\u0646\u0648\u0644\u0648\u062c\u064a\u0627 \u0627\u0644\u0645\u0639\u0644\u0648\u0645\u0627\u062a \u00b7 \u0647\u0646\u062f\u0633\u0629 \u0627\u0644\u0648\u064a\u0628 \u00b7 \u0627\u0644\u0639\u0644\u0627\u0642\u0627\u062a \u0627\u0644\u062e\u0627\u0631\u062c\u064a\u0629",
		CP4Title: "\u0623\u0628\u0627\u0643\u0648\u0641 \u2014 \u062a\u0635\u0645\u064a\u0645 \u0627\u0644\u0634\u0639\u0627\u0631 \u0627\u0644\u0625\u0646\u0633\u0627\u0646\u064a \u0648\u0627\u0644\u0647\u0648\u064a\u0629 \u0627\u0644\u0628\u0635\u0631\u064a\u0629", CP4Desc: "\u062a\u0635\u0645\u064a\u0645 \u0647\u0648\u064a\u0629 \u0628\u0635\u0631\u064a\u0629 \u062a\u0637\u0648\u0639\u064a\u0629 (Pro Bono) \u0644\u0645\u0646\u0638\u0645\u0629 APACOV \u0627\u0644\u062e\u064a\u0631\u064a\u0629. \u062a\u0645 \u062a\u0646\u0641\u064a\u0630\u0647 \u0639\u0628\u0631 Adobe Illustrator \u0648\u062a\u0637\u0628\u064a\u0642 Samsung PenUp\u060c \u0648\u064a\u0631\u0645\u0632 \u0627\u0644\u0634\u0639\u0627\u0631 \u0625\u0644\u0649 \u0627\u0644\u0623\u064a\u062f\u064a \u0627\u0644\u0645\u062a\u0643\u0627\u062a\u0641\u0629 \u0648\u0627\u0644\u0645\u0645\u062a\u062f\u0629 \u0628\u0631\u0648\u062d \u0627\u0644\u062a\u0636\u0627\u0645\u0646 \u062f\u0627\u062e\u0644 \u0642\u0648\u0633 \u0627\u0644\u0623\u0645\u0627\u0646 \u0648\u0627\u0644\u0623\u0645\u0644.", CP4Domain: "\u0627\u0644\u0647\u0648\u064a\u0629 \u0627\u0644\u0625\u0646\u0633\u0627\u0646\u064a\u0629 \u0627\u0644\u0645\u0624\u0633\u0633\u064a\u0629 \u00b7 \u0627\u0644\u062a\u0635\u0645\u064a\u0645 \u0627\u0644\u0645\u062a\u062c\u0647\u064a",
		CP5Title: "SuPublic — هندسة الصفائح المعدنية لكشك تعبئة مياه الشرب الآلي", CP5Desc: "خدمات هندسة ميكانيكية وتصميم الصفائح المعدنية ثلاثية الأبعاد عبر SolidWorks لصالح مشروع SuPublic البيئي المبتكر الذي طورته غونجا أوزاتا تحت إشراف د. يافوز سليم تاشبينار. يهدف المشروع، الذي انطلق من قونية، إلى تحديد مواقع نوافير مياه الشرب المجانية والنظيفة في تركيا والعالم للحد من النفايات البلاستيكية لزجاجات المياه ذات الاستخدام الواحد. تم تصميم الهيكل المعدني الكامل للكشك ونقاط الثني بالليزر وحجرة التعبئة الآلية للنموذج الأولي.", CP5Domain: "هندسة تعاقدية · تصميم الصفائح المعدنية · SolidWorks والنماذج الأولية",
		CPNoticeInfra: "\u0645\u0644\u0627\u062d\u0638\u0629 \u0646\u0642\u0644 \u0627\u0644\u062e\u0648\u0627\u062f\u0645: \u062a\u062c\u0631\u064a \u062d\u0627\u0644\u064a\u0627\u064b \u0639\u0645\u0644\u064a\u0629 \u0646\u0642\u0644 \u0627\u0644\u0628\u0646\u064a\u0629 \u0627\u0644\u062a\u062d\u062a\u064a\u0629 \u0648\u062e\u0648\u0627\u062f\u0645 apacov.org \u0625\u0644\u0649 \u062e\u0648\u0627\u062f\u0645\u0646\u0627 \u0627\u0644\u062e\u0627\u0635\u0629\u060c \u0625\u0644\u0649 \u062c\u0627\u0646\u0628 \u0634\u0631\u0627\u0621 \u0646\u0637\u0627\u0642\u0627\u062a \u0648\u0628\u0631\u064a\u062f \u0625\u0644\u0643\u062a\u0631\u0648\u0646\u064a \u0631\u0633\u0645\u064a \u0625\u0636\u0627\u0641\u064a.", BtnOfficialWebsite: "\u0627\u0644\u0645\u0648\u0642\u0639 \u0627\u0644\u0631\u0633\u0645\u064a (apacov.org)", BadgeVolunteered: "\u0639\u0645\u0644 \u062a\u0637\u0648\u0639\u064a \u00b7 \u0645\u062c\u0627\u0646\u064a",
		ExcavatorTitle: "تصميم حفارة هيدروليكية (مشروع عناصر الآلات) — مستوحى من كوماتسو PC200",
		ExcavatorDesc: "تصميم ميكانيكي شامل وتحليل هيكلي بالعناصر المحدودة (FEA) لحفارة مجنزرة ثقيلة بوزن 24 طناً مستوحاة من كوماتسو PC200 LC-7/LC-8. شمل المشروع حسابات دقيقة ومحاكاة ANSYS لخمسة أنظمة رئيسية: المحاور والدبابيس الثقيلة للذراع والدلو، محامل الدوران الدائري ومحامل الجنزير، وصلات البراغي عالية الإجهاد، صناديق التروس الكوكبية، والأسطوانات والدوائر الهيدروليكية عالية الضغط.",
		ExcavatorDomain: "تصميم الآلات الثقيلة · عناصر الآلات · تحليل ANSYS والأنظمة الهيدروليكية",
		SafrenTitle: "سافرين SafRen (سابقاً Ce-Saf) — عداد خطوات ذكي بالحذاء ونظام إنترنت أشياء قابل للارتداء",
		SafrenDesc: "نظام عداد خطوات ذكي مدمج في الأحذية ونقل بيانات المشي لاسلكياً، طُوّر بالتعاون المشترك مع الزميلتين سافينا نجم الدينوفا وسيرين باشاران (Ceren Başaran) المتخصصتين في تصميم الأحذية (Shoe Design)؛ حيث قامتا بتوفير الأحذية وتحديد المواضع الهندسية المثلى لتثبيت العتاد الإلكتروني (شاشة OLED LCD، ومفاتيح الضغط الأرضية، والمتحكم الدقيق). يعتمد النظام على متحكم Waveshare ESP32-S3-Zero، وحساس MPU6050 سداسي المحاور، مع خادم Wi-Fi SoftAP محلي يتزامن مباشرة مع تطبيق هاتف محمول مخصص بتقنية Flutter.",
		SafrenDomain: "إنترنت الأشياء القابل للارتداء · تحليل المشي · برمجيات ESP32-S3 · تطبيق فلاتر للهواتف",
		STM32Title: "تحدي لوحة STM32 Discovery — روبوتات متحركة ذاتية القيادة",
		STM32Desc: "هندسة روبوتية ذاتية القيادة ثنائية المهام تعتمد على لوحة STM32F407VGT6 بمعالج ARM Cortex-M4. التحدي النصفي: مصفوفة مستشعرات LDR متعددة القنوات مع تحويل ADC بقناة 12 بت والتحكم بمحركات TIM3 PWM للتوجه نحو الأضواء العشوائية في حلبة 1 متر مربع والتوقف بدقة ملليمترية دون اصطدام من المحاولة الأولى. التحدي النهائي: روبوت إطفاء شموع ذاتي باستخدام مروحة 12V عالية التدفق وتمركز ذكي بمنتصف الحلبة ودوران 360 درجة لإطفاء جميع الشموع الخمس خلال أقل من 10 ثوانٍ (ثلث الوقت المحدد).",
		STM32Domain: "الأنظمة المدمجة · STM32 ARM Cortex-M4 · برمجة C / HAL عتادية · روبوتات ذاتية",
		ROSTitle: "روبوت متنقل ذاتي القيادة بنظام ROS Noetic — جسر ESP32 عبر TCP/IP وحركيات تفاضلية",
		ROSDesc: "مشروع روبوت متنقل ذاتي القيادة تم تطويره بالتعاون مع مهندس الميكاترونكس محمد شاكار أوغلو تحت إشراف د. أوكان أويار. يربط بين نظام ROS Noetic على لينكس ومتحكم ESP32 عبر جسر مقبس Wi-Fi TCP/IP منخفض التأخير. يتضمن حركيات الدفع التفاضلي، قياس مسار مغلق الحلقة بمشفرات تربيعية 420 نبضة، وحدة IMU بستة محاور MPU6050، قاطع أمان طوارئ 500ms، ومحاكاة توأم رقمي كاملة على Gazebo وRViz تلبي معايير التحديات بدرجة 100/100.",
		ROSDomain: "الروبوتات المتنقلة ذاتية القيادة · ROS Noetic · حركيات تفاضلية · Gazebo & RViz",
		BtnLinkedInVideo: "فيديو العرض على لينكد إن", BtnGitHubRepo: "مستودع GitHub", CertsTitle: "الشهادات والاعتمادات",
		VHeroBadge: "مشارك في جائزة جيمس دايسون ٢٠٢٦ · ابتكار الرعاية الصحية",
		VHeroTitle: "فيلوم (VELUM) — بساط الوقاية من قرح الفراش",
		VHeroSubtitle: "أول بساط استشعار طبي حيوي سلبي ومستمر في العالم يكشف قرح الفراش قبل ظهورها بـ ٢-٥ أيام — بعدالة تامة لجميع ألوان البشرة وبتكلفة منخفضة.",
		VStatCost: "+١١ مليار دولار", VStatCostDesc: "التكلفة السنوية لعلاج قرح الفراش في المستشفيات الأمريكية",
		VStatEarly: "٢-٥ أيام مبكراً", VStatEarlyDesc: "إنذار استباقي عبر قياس الرطوبة تحت الجلد قبل تلف السطح",
		VStatBOM: "< ٥٠ دولار", VStatBOMDesc: "تكلفة تصنيع منخفضة تتيح تزويد كل سرير مريض بالبساط",
		VStatEquity: "١٠٠٪ عادل", VStatEquityDesc: "قياس كهربائي حيوي مستقل تماماً عن لون البشرة (أنواع فيتزباتريك I–VI)",
		VProblemTitle: "نظام رعاية تقليدي يخذل المرضى",
		VProblemSubtitle: "الفحص البصري يكتشف الجروح فقط بعد أن يكون موت الأنسجة العميقة قد حدث بالفعل.",
		VProblemP1: "تصيب قرح الفراش أكثر من ١٢.٨٪ من المرضى المنومين في المستشفيات عالمياً. تعتمد الإجراءات الحالية على التقليب اليدوي للمريض كل ساعتين والفحص البصري للممرضين. ومع ذلك، فإن الضرر الخلوي يبدأ عميقاً في الأنسجة قبل ٤٨ إلى ١٢٠ ساعة من ظهور أي احمرار على الجلد.",
		VProblemP2: "علاوة على ذلك، يفشل الفحص البصري تماماً في اكتشاف الاحمرار المبكر لدى أصحاب البشرة الداكنة. كما أن الماسحات اليدوية التجارية تكلف بين ٣٠٠٠ و١٠٠٠٠ دولار وتتطلب تدريباً شاقاً للممرضين.",
		VProblemPoint1: "نظام تفاعلي متأخر: احمرار الجلد يعني أن موت الخلايا الداخلية قد بدأ بالفعل.",
		VProblemPoint2: "فجوة عرقية موثقة: الفحص بالعين المجردة يعجز عن رصد القرح في درجات البشرة الداكنة.",
		VProblemPoint3: "تكاليف باهظة وإرهاق للممرضين: الأجهزة اليدوية تتطلب وقتاً طويلاً ومكلفة للغاية.",
		VSolutionTitle: "حل فيلوم: المعاوقة الحيوية المستمرة",
		VSolutionSubtitle: "مراقبة مستمرة وسلبية ٢٤/٧ تحت ملاءة السرير مباشرة دون إزعاج المريض.",
		VSolutionP1: "فيلوم هو بساط بوليميد فائق النحافة والمرونة يوضع تحت المريض مباشرة. تقوم مصفوفة من أقطاب الفضة والكربون المطبوعة بإجراء قياسات مطيافية المعاوقة الحيوية عبر ١٦ منطقة تشريحية حرجة كل ٥ دقائق.",
		VSolutionPoint1: "تشغيل فوري: يُوضع تحت الملاءة ويعمل مباشرة دون حاجة لتدريب أو مقاطعة نوم المريض.",
		VSolutionPoint2: "استشعار الرطوبة الخلوية: هبوط المعاوقة بنسبة ١٥٪+ يشير لتجمع السوائل قبل تشكل الجرح بأيام.",
		VSolutionPoint3: "تنبيهات فورية للممرضين: إرسال إشعارات لاسلكية تحدد بدقة الجزء التشريحي الذي يحتاج إلى تقليب.",
		VSimTitle: "محاكي بساط المعاوقة الحيوية التفاعلي",
		VSimSubtitle: "خريطة ضغط تشريحية حية لـ ١٦ منطقة مع محاكاة مسح شريحة AD5933 في الوقت الفعلي.",
		VSimNormal: "١. الوضع الطبيعي", VSimIschemia: "٢. ضغط مستمر على العجز",
		VSimOedema: "٣. إنذار ارتشاح المرحلة ٠", VSimReset: "٤. إعادة توجيه المريض",
		VSimStatusAlert: "تنبيه حرج: تم اكتشاف ارتشاح تحت الجلد! يرجى تقليب منطقة العجز فوراً",
		VSimStatusNormal: "الحالة ممتازة: تروية الأنسجة الدموية طبيعية عبر جميع المناطق التشريحية",
		VTechTitle: "البنية الميكاترونيكية المفككة",
		VTechSubtitle: "إلكترونيات مرنة فائقة الدقة مصممة للتعقيم الطبي والتصنيع الكمي.",
		VLayer1Title: "١. غلاف طبي من السيليكون والبولي يوريثان الحراري TPU",
		VLayer1Desc: "غشاء متوافق حيوياً (ISO 10993) ينفذ البخار ويمنع السوائل، ويتحمل المعقمات اليومية في المستشفيات.",
		VLayer2Title: "٢. مصفوفة أقطاب الفضة والكربون المطبوعة",
		VLayer2Desc: "أحبار موصلة عالية الجودة مطبوعة على فيلم كابتون (Kapton®) بسمك ٥٠ ميكرومتر تحت البروزات العظمية.",
		VLayer3Title: "٣. محول المعاوقة AD5933 ومبدل القنوات التناظري",
		VLayer3Desc: "شريحة Analog Devices AD5933 مع مصدر تيار هولاند (<١ مللي أمبير) ومبدل يمسح النقاط في ثوانٍ.",
		VLayer4Title: "٤. معالج ESP32-S3 اللاسلكي ثنائي النواة",
		VLayer4Desc: "معالج بتردد ٢٤٠ ميجاهرتز يشغل نظام FreeRTOS لمعايرة القراءات وإرسال التنبيهات اللاسلكية المشفرة.",
		VCompTitle: "الميزة التنافسية السريرية",
		VCompSubtitle: "لماذا يتفوق فيلوم سريرياً واقتصادياً على كافة الحلول القائمة.",
		VCompColMetric: "المعيار / الخاصية", VCompColVisual: "الفحص البصري", VCompColSEM: "الماسح اليدوي", VCompColMattress: "مراتب تخفيف الضغط", VCompColVelum: "بساط فيلوم",
		VRoadmapTitle: "خارطة الطريق السريرية والتجارية",
		VRoadmapSubtitle: "من الاختبارات المعملية إلى التجارب السريرية في العناية المركزة واعتماد إدارة الغذاء والدواء FDA.",
		VGalleryTitle: "الرسومات الهندسية ونماذج العتاد",
		VGallerySubtitle: "رسومات ثلاثية الأبعاد عالية الدقة ومخططات التجميع الميكاترونيكي للمنتج.",
		VDownloadPitch: "تحميل العرض التقديمي (PDF)",
		VViewGitHub: "عرض مستودع جيت هاب",
		FooterRights: "جميع الحقوق محفوظة.",
		FooterBuiltWith: "",
	},
	"pt": {
		LangCode: "pt", LangName: "Portuguese", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Início", NavProjects: "Projetos", NavExperience: "Experiência", NavAbout: "Sobre", NavSkills: "Habilidades", NavContact: "Contato", NavEducation: "Formação", NavHonors: "Prêmios", NavCertificates: "Certificados", NavActivities: "Workshops", NavLanguages: "Idiomas", NavGitHub: "GitHub", NavVelum: "Projeto VELUM",
		HeroBadge: "Engenheiro Mecatrônico · Hardware Biomédico · Sistemas Embarcados · Poliglota",
		AuthorRole: "B.Sc. em Engenharia Mecatrônica · Desenvolvedor de Sistemas Embarcados & IoT",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Projetando sistemas mecatrônicos biomédicos inteligentes, instrumentação sensorial de precisão e dispositivos IoT para proteger vidas humanas.",
		HeroBtnExplore: "Explorar Projetos", HeroBtnContact: "Entrar em Contato",
		StatProjects: "18 Projetos de Engenharia", StatSoloProjects: "Projetos Individuais", StatTeamProjects: "Projetos em Equipe", StatContractProjects: "Projetos Contratuais", StatWorkshops: "Workshops & Atividades", StatAward: "Engenharia Mecatrônica", StatLanguages: "7 Idiomas Fluentes", StatFocus: "ESP32 / Bioimpedância / CAD",
		ProjectsTitle: "Portfólio de Engenharia",
		ProjectsSubtitle: "Hardware biomédico preventivo, dispositivos assistivos, automação fluídica e sistemas mecatrônicos.",
		StatusActive: "Concluído e Interativo", StatusQueue: "Na Fila",
		BtnViewProject: "Ver Projeto", BtnPreview: "Detalhes do Projeto",
		BadgeFeatured: "Projeto em Destaque", BadgeBiomedical: "Hardware Biomédico", BadgeHardware: "Mecatrônica", BadgeRobotics: "Robótica e Automação",
		P1Title: "VELUM (antigo SentIO)",
		P1Desc: "Tapete sensor de bioimpedância elétrica e umidade subepidérmica contínua 24/7 que detecta úlceras por pressão de 2 a 5 dias antes de se tornarem visíveis.",
		P1Domain: "Mecatrônica Biomédica · Espectroscopia de Bioimpedância",
		P2Title: "FLOWGUARD",
		P2Desc: "Sistema mecatrônico automatizado de filtragem de microplásticos para máquinas de lavar, impedindo a contaminação dos oceanos.",
		P2Domain: "Mecatrônica Fluídica · Separação de Microplásticos",
		P3Title: "SOLUS",
		P3Desc: "Adesivo vestível de monitoramento contínuo da temperatura corporal central para alerta precoce de insolação em trabalhadores sob estresse térmico.",
		P3Domain: "Biossensores Vestíveis · Fisiologia Preventiva",
		P4Title: "DENOVA (antigo Avicenna)",
		P4Desc: "Detector óptico não invasivo de cárie dentária com iluminação multiespectral e fluorescência para diagnóstico precoce de desmineralização.",
		P4Domain: "Biofotônica · Diagnóstico Óptico",
		P5Title: "RENOVO",
		P5Desc: "Prensa hidráulica portátil para auxílio em desastres que converte entulhos de demolição de terremotos em blocos estruturais no local.",
		P5Domain: "Mecânica Pesada e Hidráulica · Auxílio em Catástrofes",
		P6Title: "CURA SERO",
		P6Desc: "Tira bioadesiva cirúrgica de microtensão para fechamento de feridas, garantindo distribuição uniforme da tração e minimizando cicatrizes.",
		P6Domain: "Dispositivo Biomecânico · Tecnologia Cirúrgica",
		P7Title: "WELSTECH — Wearable Life Saver Technology",
		P7Desc: "Monitor de sinais vitais de pulso ultra-compacto com biossensor óptico PPG, processamento ESP32-S3, telemetria Wi-Fi local e alarme sonoro escalonável para anomalias cardíacas.",
		P7Domain: "IoT Biomédico Vestível · Biosensoriamento Cardíaco",
		P8Title: "SENDIT",
		P8Desc: "Mecanismo autônomo de logística e triagem de bagagens e encomendas em aeroportos, otimizando o fluxo em terminais aéreos.",
		P8Domain: "Robótica Autônoma · Automação Logística",
		P9Title: "SMARTWASTE",
		P9Desc: "Sistema robótico mecatrônico de triagem de resíduos que combina espectrometria embarcada e visão computacional para reciclagem automatizada.",
		P9Domain: "Visão Computacional · Atuação Mecatrônica",
		P10Title: "IMAGE CLASSIFIER — Rede Neural Convolucional Profunda",
		P10Desc: "Rede neural convolucional profunda (CNN) treinada em TensorFlow/Keras com mais de 9.600 imagens, pré-processamento automatizado de 96×96 e interface de inferência gráfica.",
		P10Domain: "Deep Learning · Visão Computacional e Edge AI",
		P11Title: "LUMIERE — Sistema Inteligente de Notificação de Falhas em Iluminação Pública",
		P11Desc: "Dispositivo IoT compacto instalado sob postes de iluminação pública, combinando sensor fotossensível LDR, geolocalização GY-NEO6MV2 GPS, telemetria ESP32 Wi-Fi SoftAP e alarme sonoro para automatizar o reparo municipal.",
		P11Domain: "IoT para Cidades Inteligentes · Firmware Embarcado & Telemetria",
		P12Title: "NOTCHA — Entalhe de Tempo Dinâmico Minimalista Flutuante",
		P12Desc: "Criado em 2021/2022 quando o Windows 11 removeu a opção de posicionar a barra de tarefas no topo da tela, o NOTCHA é um utilitário leve para desktop que traz um entalhe de relógio dinâmico na borda superior, com atualização em tempo real e expansão inteligente ao passar o mouse.",
		P12Domain: "Ergonomia de Desktop · Utilitários de Interface Personalizados",
		P13Title: "SPO2 — Design de PCB de 2 Camadas para Oxímetro de Pulso",
		P13Desc: "Placa de circuito impresso de 2 camadas projetada no KiCad 10.0, integrando símbolos e footprints personalizados do zero, tolerâncias reais de fabricação, biossensor óptico MAX30102, display OLED SSD1306 e núcleo Arduino Nano.",
		P13Domain: "Design de PCB e EDA · Hardware Biomédico · KiCad 10.0",
		P14Title: "RESPIRO",
		P14Desc: "Ventilador assistivo adaptativo biomédico de baixo custo e alta confiabilidade projetado para implantação de emergência e clínicas com recursos limitados.",
		P14Domain: "Mecatrônica Biomédica · Cuidados Críticos",
		AboutTitle: "Sobre Ukacha Juma",
		AboutSubtitle: "Engenheiro Mecatrônico · Arquiteto Hardware-Software · Poliglota",
		AboutP1: "Sou um Engenheiro Mecatrônico impulsionado pela convicção de que a engenharia de hardware moderna deve prevenir o sofrimento humano em vez de apenas remediá-lo. Meu trabalho conecta design mecânico em CAD/CAM, circuitos sensores analógicos de precisão, firmware em microcontroladores bare-metal e FreeRTOS, e telemetria em tempo real.",
		AboutP2: "Com projetos que abrangem tapetes de prevenção de escaras (VELUM), filtragem de microplásticos (FLOWGUARD) e biossensores térmicos (SOLUS), construo sistemas físicos desde os princípios matemáticos fundamentais até a validação em bancada e projetos prontos para produção.",
		AboutPolyglot: "Comunicação Multilíngue e Colaboração Global",
		AboutPolyglotDesc: "Fluente em inglês, turco, árabe, português, espanhol, alemão e francês. Ser poliglota me permite conectar equipes de engenharia ao redor do mundo, consultar literatura clínica nas fontes originais e projetar tecnologias universalmente acessíveis.",
		SkillsTitle: "Competências Técnicas e Ferramentas",
		ExpTitle: "Experiência Profissional", ExpSubtitle: "Engenharia industrial prática, manutenção e diagnóstico de CNC, automação e tradução internacional.", EduTitle: "Educação e Formação", AwardsTag: "RECONHECIMENTO ACADÊMICO E ROBÓTICO", AwardsTitle: "Honrarias e Prêmios", AwardsSubtitle: "Distinções acadêmicas institucionais, classificações em olimpíadas nacionais de robótica e mérito em engenharia.", AwardsCardTitle: "Classificação İTÜRO 2026 — Placa de Honra Concedida pela Reitora", AwardsCardDesc: "Homenageado com placa comemorativa entregue pela estimada Diretora Profa. Dra. Humar Kahramanli Örnek na formatura da Faculdade de Tecnologia da Universidade de Selçuk (2025–26), reconhecendo a classificação da equipe STWolf na Olimpíada de Robótica da Universidade Técnica de Istambul (İTÜRO 2026).", AwardsViewPost: "Ver Publicação e Reconhecimento no LinkedIn", AwardsDeanTitle: "Diretora da Faculdade de Tecnologia", AwardsDeanName: "Profa. Dra. Humar Kahramanlı Örnek", AwardsTeamName: "Equipe STWolf · Universidade de Selçuk", AwardsOlympiad: "Olimpíada de Robótica da Universidade Técnica de Istambul (İTÜRO 2026)", AwardsCeremony: "Cerimônia de Formatura 2025–26 · Universidade de Selçuk", AwardsQuoteLabel: "Inscrição Oficial da Placa", TabEducation: "Formação Acadêmica", TabAwards: "Honrarias e Prêmios", TabCertificates: "Certificados & Licenças", TabActivities: "Workshops", TabLanguages: "Idiomas", ActivitiesTitle: "Atividades Técnicas & Extracurriculares", ActivitiesSubtitle: "Exposições em festivais de tecnologia, liderança comunitária e metodologias de pesquisa com IA.", TeknoselTitle: "Festival de Tecnologia Selçuklu Teknosel — Categoria Universitária", TeknoselDesc: "Apresentação do protótipo mecatrônico KonyaWaste em nosso estande na categoria Projeto Livre - Nível Universitário, demonstrando o hardware ao vivo para o Prefeito de Selçuklu.", TeknoselMayorBadge: "Apresentação Pessoal ao Prefeito de Selçuklu", AiResearchTitle: "IA Baseada em Fontes na Pesquisa Acadêmica (Workshop TADD)", AiResearchDesc: "Conclusão de workshop intensivo de 2 dias organizado pela TADD e liderado por Yunus Emre Kıymaz, dominando ferramentas como Zotero, ResearchRabbit e NotebookLM para eliminar alucinações de IA e assegurar rigor técnico em revisões bibliográficas de engenharia.", AiResearchBadge: "Pesquisa Confiável sem Alucinações de IA", IfmWorkshopTitle: "Expandindo Minhas Competências em Automação com ifm electronic & CoDeSys (Intensivo de 48h)", IfmWorkshopDesc: "Conclusão de treinamento intensivo de 48 horas em PLC industrial ministrado na Universidade de Selçuk pela sede da ifm electronic Türkiye. Experiência prática com arquitetura de PLC, CoDeSys V3.5 (SP18 & SP19), protocolos Modbus e CANbus e diagnóstico em campo.", IfmWorkshopBadge: "⚙️ ifm electronic Türkiye · Universidade de Selçuk", IfmBtnLinkedIn: "Ver Publicação do Workshop no LinkedIn", TabSoloProjects: "Projetos Individuais", TabTeamProjects: "Projetos em Equipe", TabContractProjects: "Projetos Contratuais", ContractEmptyBadge: "Disponível para Contratos", ContractEmptyTitle: "Projetos Contratuais e Consultoria", ContractEmptyDesc: "Contratos comerciais propriet\u00e1rios, consultoria em hardware e mecatr\u00f4nica industrial. Estudos de caso em breve.",
		CP1Title: "Mec\u00e2nica Auto Moderna \u2014 Design de Logotipo e Identidade Visual", CP1Desc: "Design de logotipo corporativo e identidade visual sob medida para a Mec\u00e2nica Auto Moderna. Cria\u00e7\u00e3o de emblema automotivo exclusivo integrando ferramentas de precis\u00e3o, simbologia 24h e tipografia din\u00e2mica para sinaliza\u00e7\u00e3o, uniformes e canais digitais.", CP1Domain: "Identidade Corporativa e Design Vetorial",
		CP2Title: "Mec\u00e2nica Auto Moderna \u2014 Sistema Invoice Moderna", CP2Desc: "Solu\u00e7\u00e3o de software personalizada desenvolvida para a Mec\u00e2nica Auto Moderna. Simplifica e automatiza a emiss\u00e3o de faturas eletr\u00f4nicas e faturas proforma, c\u00e1lculo de IVA, cadastro de ordens de servi\u00e7o e clientes, aumentando a efici\u00eancia e a precis\u00e3o operacional no fluxo de trabalho da oficina.", CP2Domain: "Software de Gest\u00e3o e Faturamento Eletr\u00f4nico",
		BtnInstagramPost: "Ver Publica\u00e7\u00e3o no Instagram", BtnLinkedInPost: "Ver Publica\\u00e7\\u00e3o no LinkedIn",
		CP3Title: "APACOV \u2014 Infraestrutura de TI & Arquitetura Web", CP3Desc: "Desenvolvimento volunt\u00e1rio do portal web institucional completo e gest\u00e3o de infraestrutura de TI para a APACOV (Associa\u00e7\u00e3o Para Protec\u00e7\u00e3o e Amparo a Crian\u00e7as \u00d3rf\u00e3s e Vulner\u00e1veis). Atuando como Administrador de TI respons\u00e1vel pela implanta\u00e7\u00e3o de servidores, seguran\u00e7a digital e rela\u00e7\u00f5es internacionais/assuntos externos.", CP3Domain: "Infraestrutura de TI \u00b7 Arquitetura Web \u00b7 Rela\u00e7\u00f5es Externas",
		CP4Title: "APACOV \u2014 Design de Logotipo e Identidade Humanit\u00e1ria", CP4Desc: "Identidade visual pro bono desenvolvida para a organiza\u00e7\u00e3o humanit\u00e1ria APACOV. Criada no Adobe Illustrator com toques finais no Samsung PenUp, o emblema retrata m\u00e3os de apoio m\u00fatuo estendidas em solidariedade, envolvidas por um c\u00edrculo protetor de esperan\u00e7a.", CP4Domain: "Identidade Visual Humanit\u00e1ria \u00b7 Design Vetorial",
		CP5Title: "SuPublic — Engenharia de Chapas Metálicas para Quiosque de Água Potável", CP5Desc: "Serviço de engenharia mecânica contratada e modelagem paramétrica de chapas metálicas em SolidWorks para a iniciativa sustentável SuPublic, fundada por Gonca Özata com orientação do Dr. Yavuz Selim Taşpınar. Iniciado em Konya para mapear bebedouros públicos de água limpa na Turquia e no mundo e erradicar o descarte de garrafas plásticas descartáveis. Desenvolveu todo o gabinete em chapa metálica, dobras de precisão, suportes estruturais e nicho ergonômico para o protótipo de quiosque de recarga automatizada.", CP5Domain: "Engenharia Contratada · Design em Chapa Metálica · SolidWorks CAD e Prototipagem",
		CPNoticeInfra: "Aviso de Transi\u00e7\u00e3o de Servidores: Processo ativo de migra\u00e7\u00e3o da infraestrutura e transi\u00e7\u00e3o do apacov.org para servidores pr\u00f3prios em andamento, incluindo novos dom\u00ednios e e-mails institucionais.", BtnOfficialWebsite: "Website Oficial (apacov.org)", BadgeVolunteered: "Voluntariado \u00b7 Pro Bono",
		ExcavatorTitle: "Projeto de Escavadeira Hidráulica (Elementos de Máquinas) — Inspirado em Komatsu PC200",
		ExcavatorDesc: "Projeto de engenharia mecânica e análise estrutural FEA no ANSYS para escavadeira hidráulica de 24 toneladas inspirada na Komatsu PC200 LC-7/LC-8. Cálculos analíticos rigorosos e seleção por catálogo em 5 subsistemas fundamentais: eixos e pinos de articulação da lança/braço/caçamba, rolamentos de giro e esteiras, juntas aparafusadas pré-tensionadas, redutores planetários e cilindros hidráulicos de alta pressão.",
		ExcavatorDomain: "Projeto de Máquinas Pesadas · Elementos de Máquinas · FEA ANSYS & Hidráulica",
		SafrenTitle: "SafRen (anteriormente Ce-Saf) — Pedômetro Inteligente para Calçados e IoT Vestível",
		SafrenDesc: "Sistema vestível colaborativo de contagem de passos e telemetria de marcha em calçados, desenvolvido em parceria interdisciplinar com as alunas de Design de Calçados Safina Najmudinova e Ceren Başaran, responsáveis pelo fornecimento dos calçados e pela determinação dos pontos ergonômicos ideais para alocação dos componentes eletrônicos (display OLED LCD, botões de impacto e microcontrolador). Equipado com Waveshare ESP32-S3-Zero, sensor inercial MPU6050 de 6 graus de liberdade, servidor HTTP Wi-Fi SoftAP integrado e sincronização sem fio em tempo real com app móvel Flutter.",
		SafrenDomain: "IoT Vestível · Análise de Marcha · Firmware ESP32-S3 · Flutter Mobile",
		STM32Title: "Desafio da Placa STM32 Discovery — Robótica Móvel Autônoma",
		STM32Desc: "Projeto de robótica autônoma em duas etapas baseado na placa STM32F407VGT6 ARM Cortex-M4. Desafio 1 (Vize): Matriz de sensores LDR com ADC de 12 bits e controle diferencial de motores TIM3 PWM para localizar luzes aleatórias em arena de 1 m² e parar com precisão milimétrica sem colisão na primeira tentativa. Desafio 2 (Final): Robô extintor de velas com ventoinha de 12V, posicionamento central estratégico e varredura rotacional contínua de 360°—apagando todas as 5 velas em menos de 10 segundos (1/3 do tempo estipulado).",
		STM32Domain: "Sistemas Embarcados · STM32 ARM Cortex-M4 · C Bare-Metal / HAL · Robótica Autônoma",
		ROSTitle: "Robô Móvel Autônomo ROS Noetic — Ponte TCP/IP ESP32 e Cinemática Diferencial",
		ROSDesc: "Robô móvel autônomo desenvolvido em colaboração com o engenheiro mecatrônico Mehmet Çakaroğlu sob orientação do Dr. Okan Uyar. Faz a ponte entre o ROS Noetic no Linux e o microcontrolador ESP32 de núcleo duplo por meio de um socket Wi-Fi TCP/IP de baixa latência. Apresenta cinemática diferencial, odometria de malha fechada com encoders de 420 pulsos, IMU MPU6050 de 6 eixos calibrada, botão de segurança deadman de 500 ms e simulação digital twin em Gazebo/RViz atingindo pontuação máxima de 100/100.",
		ROSDomain: "Robótica Móvel Autônoma · ROS Noetic · Cinemática Diferencial · Gazebo e RViz",
		BtnLinkedInVideo: "Vídeo no LinkedIn", BtnGitHubRepo: "Repositório GitHub", CertsTitle: "Certificações e Distinções",
		VHeroBadge: "Submissão ao James Dyson Award 2026 · Inovação em Saúde",
		VHeroTitle: "VELUM — Tapete de Prevenção de Úlceras por Pressão",
		VHeroSubtitle: "O primeiro tapete sensor biomédico passivo e contínuo do mundo que detecta úlceras por pressão de 2 a 5 dias antes de sua visibilidade — de forma equitativa, acessível e sem perturbar o paciente.",
		VStatCost: "+11 Bilhões $ USD", VStatCostDesc: "Custo anual de tratamento de escaras hospitalares nos EUA",
		VStatEarly: "2 a 5 Dias Antes", VStatEarlyDesc: "Alerta precoce via detecção de umidade subepidérmica",
		VStatBOM: "< $50 USD", VStatBOMDesc: "Custo alvo de materiais permitindo adoção em todos os leitos",
		VStatEquity: "100% Equitativo", VStatEquityDesc: "Detecção elétrica agnóstica ao tom de pele (Fitzpatrick I–VI)",
		VProblemTitle: "Um Sistema Clínico Tradicional Falho",
		VProblemSubtitle: "A inspeção visual detecta feridas somente quando a necrose tecidual já é irreversível.",
		VProblemP1: "As lesões por pressão afetam mais de 12,8% dos pacientes hospitalizados globalmente. O padrão atual baseia-se em rondas manuais de enfermeiros para reposicionamento a cada 2 horas e inspeções visuais subjetivas da pele. Contudo, os danos celulares começam profundamente no tecido subepidérmico de 48 a 120 horas antes do aparecimento do eritema superficial.",
		VProblemP2: "Além disso, a avaliação visual do eritema falha sistematicamente em pacientes de pele escura (Fitzpatrick IV–VI). Scanners SEM manuais comerciais custam entre US$ 3.000 e US$ 10.000 e exigem treinamento especializado.",
		VProblemPoint1: "Reativo por Natureza: A vermelhidão na pele indica que a morte celular já está em andamento.",
		VProblemPoint2: "Disparidade Étnica Documentada: A avaliação visual não detecta lesões em tons de pele escuros.",
		VProblemPoint3: "Custos Elevados e Fadiga da Enfermagem: Scanners manuais são caros e aumentam a carga de trabalho.",
		VSolutionTitle: "A Solução VELUM: Bioimpedância Contínua",
		VSolutionSubtitle: "Monitoramento passivo de tecido 24 horas por dia, discretamente sob o lençol.",
		VSolutionP1: "O VELUM é um tapete sensor ultrafino e flexível de poliimida colocado sob o paciente acamado. Uma matriz de eletrodos de prata e carbono impressos realiza espectroscopia de bioimpedância elétrica (1–100 kHz) em 16 zonas anatômicas críticas a cada 5 minutos.",
		VSolutionPoint1: "Instalação Instantânea: Deslize sob o lençol e ligue. Zero perturbação ao paciente.",
		VSolutionPoint2: "Detecção de Umidade Subepidérmica: Uma queda de 15%+ na impedância acusa edema tecidual dias antes da ferida.",
		VSolutionPoint3: "Alertas Imediatos via BLE / Wi-Fi: Enfermeiros recebem notificações identificando a área anatômica exata em risco.",
		VSimTitle: "Simulador Interativo do Tapete de Bioimpedância",
		VSimSubtitle: "Mapeamento anatômico de pressão em 16 zonas e simulação da varredura de impedância do AD5933.",
		VSimNormal: "1. Estado Normal", VSimIschemia: "2. Pressão Sacral Sustentada",
		VSimOedema: "3. Alerta de Edema Estágio 0", VSimReset: "4. Reposicionamento do Paciente",
		VSimStatusAlert: "CRÍTICO: Edema Subepidérmico Detectado! Reposicione a Região Sacral Imediatamente",
		VSimStatusNormal: "NOMINAL: Perfusão tecidual ideal em todas as zonas anatômicas",
		VTechTitle: "Arquitetura Mecatrônica em Vista Explodida",
		VTechSubtitle: "Eletrônica flexível de precisão projetada para desinfecção hospitalar e produção em larga escala.",
		VLayer1Title: "1. Encapsulamento em TPU Médico e Silicone Platina",
		VLayer1Desc: "Membrana biocompatível (ISO 10993) permeável ao vapor e impermeável a líquidos, resistente a desinfetantes hospitalares.",
		VLayer2Title: "2. Matriz de Eletrodos Impressos de Prata e Carbono",
		VLayer2Desc: "Filme de poliimida DuPont Kapton® de 50 µm com circuitos serigrafados posicionados sob proeminências ósseas.",
		VLayer3Title: "3. Conversor de Impedância AD5933 e MUX de 16 Canais",
		VLayer3Desc: "Analog Devices AD5933 acoplado a fonte de corrente Howland (<1 mA CA) e multiplexador CD74HC4067 para varredura ultrarrápida.",
		VLayer4Title: "4. Módulo de Computação ESP32-S3 Dual-Core e Sem Fio",
		VLayer4Desc: "Processador Xtensa LX7 a 240MHz com FreeRTOS, calibração adaptativa, rede BLE 5.0 e telemetria Wi-Fi criptografada para prontuários eletrônicos.",
		VCompTitle: "Diferencial Clínico e Competitivo",
		VCompSubtitle: "Por que o VELUM cria uma vantagem clínica e econômica insuperável.",
		VCompColMetric: "Métrica / Recurso", VCompColVisual: "Inspeção Visual", VCompColSEM: "Sonda SEM Manual", VCompColMattress: "Colchões de Alívio", VCompColVelum: "Tapete VELUM",
		VRoadmapTitle: "Roteiro Clínico e Comercial",
		VRoadmapSubtitle: "Da validação em laboratório aos ensaios clínicos em UTI e submissão à FDA 510(k).",
		VGalleryTitle: "Renderizações de Hardware e Esquemas",
		VGallerySubtitle: "Ilustrações de produto em alta resolução e diagramas de montagem mecatrônica.",
		VDownloadPitch: "Baixar Apresentação para Investidores (PDF)",
		VViewGitHub: "Ver Repositório no GitHub",
		FooterRights: "Todos os direitos reservados.",
		FooterBuiltWith: "",
	},
	"es": {
		LangCode: "es", LangName: "Spanish", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Inicio", NavProjects: "Proyectos", NavExperience: "Experiencia", NavAbout: "Sobre mí", NavSkills: "Habilidades", NavContact: "Contacto", NavEducation: "Educación", NavHonors: "Premios", NavCertificates: "Certificados", NavActivities: "Workshops", NavLanguages: "Idiomas", NavGitHub: "GitHub", NavVelum: "Proyecto VELUM",
		HeroBadge: "Ingeniero Mecatrónico · Hardware Biomédico · Sistemas Embebidos · Políglota",
		AuthorRole: "Grado en Ingeniería Mecatrónica · Desarrollador de Sistemas Embebidos & IoT",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Diseño de sistemas mecatrónicos biomédicos inteligentes, instrumentación sensorial de precisión y dispositivos IoT para proteger vidas humanas.",
		HeroBtnExplore: "Explorar Proyectos", HeroBtnContact: "Contactar",
		StatProjects: "18 Proyectos de Ingeniería", StatSoloProjects: "Proyectos Individuales", StatTeamProjects: "Proyectos en Equipo", StatContractProjects: "Proyectos Contractuales", StatWorkshops: "Talleres y Workshops", StatAward: "Ingeniería Mecatrónica", StatLanguages: "7 Idiomas Fluidos", StatFocus: "ESP32 / Bioimpedancia / CAD",
		ProjectsTitle: "Portafolio de Ingeniería",
		ProjectsSubtitle: "Hardware biomédico preventivo, dispositivos de asistencia, automatización fluídica y sistemas mecatrónicos.",
		StatusActive: "Completado e Interactivo", StatusQueue: "En Cola",
		BtnViewProject: "Ver Proyecto", BtnPreview: "Detalles del Proyecto",
		BadgeFeatured: "Proyecto Destacado", BadgeBiomedical: "Hardware Biomédico", BadgeHardware: "Mecatrónica", BadgeRobotics: "Robótica y Automatización",
		P1Title: "VELUM (anteriormente SentIO)",
		P1Desc: "Esterilla sensora de bioimpedancia eléctrica y humedad subepidérmica continua 24/7 que detecta úlceras por presión de 2 a 5 días antes de que aparezcan heridas.",
		P1Domain: "Mecatrónica Biomédica · Espectroscopia de Bioimpedancia",
		P2Title: "FLOWGUARD",
		P2Desc: "Sistema mecatrónico automatizado de filtración de microplásticos para lavadoras que evita la contaminación de los ecosistemas acuáticos.",
		P2Domain: "Mecatrónica Fluídica · Separación de Microplásticos",
		P3Title: "SOLUS",
		P3Desc: "Parche vestible de monitorización continua de temperatura central para la alerta temprana de golpes de calor en trabajadores expuestos.",
		P3Domain: "Biosensores Vestibles · Fisiologia Preventiva",
		P4Title: "DENOVA (anteriormente Avicenna)",
		P4Desc: "Detector óptico no invasivo de caries dental con iluminación multiespectral y fluorescencia para el diagnóstico precoz de desmineralización.",
		P4Domain: "Biofotónica · Diagnóstico Óptico",
		P5Title: "RENOVO",
		P5Desc: "Prensa hidráulica portátil de socorro que transforma escombros de terremotos en bloques estructurales machihembrados in situ.",
		P5Domain: "Mecánica Pesada e Hidráulica · Ayuda en Desastres",
		P6Title: "CURA SERO",
		P6Desc: "Tira bioadhesiva quirúrgica de microtensión para cierre de heridas que distribuye uniformemente la tensión y minimiza cicatrices.",
		P6Domain: "Dispositivo Biomecánico · Tecnología Quirúrgica",
		P7Title: "WELSTECH — Wearable Life Saver Technology",
		P7Desc: "Monitor de constantes vitales de muñeca ultracompacto con biosensor óptico PPG, microcontrolador ESP32-S3, telemetría Wi-Fi local y alarma sonora escalable para anomalías cardíacas.",
		P7Domain: "IoT Biomédico Wearable · Biosensado Cardíaco",
		P8Title: "SENDIT",
		P8Desc: "Mecanismo autónomo de clasificación y tránsito logístico de equipajes y correo urgente en aeropuertos.",
		P8Domain: "Robótica Autónoma · Automatización Logística",
		P9Title: "SMARTWASTE",
		P9Desc: "Sistema robótico mecatrónico de clasificación de residuos que combina espectrometría embebida y visión artificial para reciclaje automatizado.",
		P9Domain: "Visión Artificial · Actuación Mecatrónica",
		P10Title: "IMAGE CLASSIFIER — Red Neuronal Convolucional Profunda",
		P10Desc: "Red neuronal convolucional profunda (CNN) desarrollada en TensorFlow/Keras con dataset de más de 9.600 imágenes, preprocesamiento dinámico de 96×96 y validadores de inferencia en tiempo real.",
		P10Domain: "Deep Learning · Visión Artificial e Inteligencia en el Borde",
		P11Title: "LUMIERE — Sistema Inteligente de Notificación de Averías en Alumbrado Público",
		P11Desc: "Dispositivo IoT compacto instalado bajo farolas municipales, que combina sensor fotométrico LDR, geolocalización GY-NEO6MV2 GPS, telemetría ESP32 Wi-Fi SoftAP y alarma acústica para automatizar el reporte de averías.",
		P11Domain: "IoT para Ciudades Inteligentes · Firmware Embebido & Telemetría",
		P12Title: "NOTCHA — Muesca de Tiempo Dinámica y Flotante Minimalista",
		P12Desc: "Diseñado en 2021/2022 cuando Windows 11 eliminó la capacidad de colocar la barra de tareas en la parte superior, NOTCHA es una utilidad ligera que ofrece una muesca de reloj flotante en el bisel superior con reloj en tiempo real y despliegue inteligente al pasar el cursor.",
		P12Domain: "Ergonomía de Escritorio · Utilidades de UI Personalizadas",
		P13Title: "SPO2 — Diseño de PCB de 2 Capas para Pulsioxímetro",
		P13Desc: "Placa de circuito impreso de 2 capas desarrollada en KiCad 10.0, integrando huellas y símbolos personalizados desde cero, tolerancias reales de fabricación, biosensor MAX30102, pantalla OLED SSD1306 y núcleo Arduino Nano.",
		P13Domain: "Diseño de PCB y EDA · Hardware Biomédico · KiCad 10.0",
		P14Title: "RESPIRO",
		P14Desc: "Ventilador asistencial adaptativo biomédico de bajo costo y alta confiabilidad diseñado para despliegue de emergencia y clínicas con recursos limitados.",
		P14Domain: "Mecatrónica Biomédica · Cuidados Críticos",
		AboutTitle: "Sobre Ukacha Juma",
		AboutSubtitle: "Ingeniero Mecatrónico · Arquitecto Hardware-Software · Políglota",
		AboutP1: "Soy un Ingeniero Mecatrónico motivado por la convicción de que la ingeniería moderna de hardware debe prevenir el sufrimiento humano en lugar de simplemente tratarlo. Mi trabajo integra diseño mecánico CAD/CAM, circuitos sensores analógicos de precisión, firmware en microcontroladores bare-metal y FreeRTOS, y telemetría en tiempo real.",
		AboutP2: "Con proyectos que van desde esterillas de prevención de úlceras (VELUM), filtración de microplásticos (FLOWGUARD) y biosensores térmicos (SOLUS), construyo sistemas físicos desde principios matemáticos hasta pruebas de laboratorio y producción.",
		AboutPolyglot: "Comunicación Multilingüe y Colaboración Global",
		AboutPolyglotDesc: "Hablo con fluidez inglés, turco, árabe, portugués, español, alemán y francés. Esta capacidad me permite liderar equipos internacionales, revisar literatura médica en sus idiomas originales y concebir tecnologías de alcance universal.",
		SkillsTitle: "Competencias Técnicas y Herramientas",
		ExpTitle: "Experiencia Profesional", ExpSubtitle: "Ingeniería industrial aplicada, diagnóstico de maquinaria CNC, mantenimiento de automatización y traducción internacional.", EduTitle: "Educación y Formación", AwardsTag: "RECONOCIMIENTO ACADÉMICO Y ROBÓTICO", AwardsTitle: "Honores y Premios", AwardsSubtitle: "Distinciones institucionales, clasificaciones en olimpiadas nacionales de robótica y excelencia en ingeniería.", AwardsCardTitle: "Clasificación İTÜRO 2026 — Placa de Honor Concedida por la Decana", AwardsCardDesc: "Homenajeado con placa honorífica entregada por la distinguida Decana Prof. Dra. Humar Kahramanli Örnek en la ceremonia de graduación de la Facultad de Tecnología de la Universidad de Selçuk (2025–26), reconociendo la clasificación del equipo STWolf en la Olimpiada de Robótica de la Universidad Técnica de Estambul (İTÜRO 2026).", AwardsViewPost: "Ver Publicación y Reconocimiento en LinkedIn", AwardsDeanTitle: "Decana de la Facultad de Tecnología", AwardsDeanName: "Prof. Dra. Humar Kahramanlı Örnek", AwardsTeamName: "Equipo STWolf · Universidad de Selçuk", AwardsOlympiad: "Olimpiada de Robótica de la Universidad Técnica de Estambul (İTÜRO 2026)", AwardsCeremony: "Ceremonia de Graduación 2025–26 · Universidad de Selçuk", AwardsQuoteLabel: "Inscripción Oficial de la Placa", TabEducation: "Educación Académica", TabAwards: "Honores y Premios", TabCertificates: "Certificados y Licencias", TabActivities: "Workshops", TabLanguages: "Idiomas", ActivitiesTitle: "Actividades Técnicas y Extracurriculares", ActivitiesSubtitle: "Exhibiciones en festivales tecnológicos, participación comunitaria y metodologías de investigación con IA.", TeknoselTitle: "Festival de Tecnología Selçuklu Teknosel — Categoría Universitaria", TeknoselDesc: "Exhibición del prototipo mecatrónico KonyaWaste en nuestro estand en la categoría Proyecto Libre - Nivel Universitario, explicando el funcionamiento directamente al Alcalde de Selçuklu.", TeknoselMayorBadge: "Presentación Personal al Alcalde de Selçuklu", AiResearchTitle: "IA Basada en Fuentes para Investigación Académica (Taller TADD)", AiResearchDesc: "Finalización del taller intensivo de 2 días impartido por TADD y Yunus Emre Kıymaz, dominando Zotero, ResearchRabbit y NotebookLM para erradicar alucinaciones de IA y garantizar máxima precisión técnica en ingeniería.", AiResearchBadge: "Investigación Verificada sin Alucinaciones de IA", IfmWorkshopTitle: "Ampliando Mis Competencias en Automatización con ifm electronic y CoDeSys (Intensivo de 48h)", IfmWorkshopDesc: "Finalización exitosa del programa intensivo de 48 horas en PLC industrial impartido en la Universidad de Selçuk por la sede central de ifm electronic Turquía. Práctica avanzada en arquitectura PLC, CoDeSys V3.5 (SP18 y SP19), protocolos Modbus y CANbus y resolución de incidencias en planta.", IfmWorkshopBadge: "⚙️ ifm electronic Turquía · Universidad de Selçuk", IfmBtnLinkedIn: "Ver Publicación del Taller en LinkedIn", TabSoloProjects: "Proyectos Individuales", TabTeamProjects: "Proyectos en Equipo", TabContractProjects: "Proyectos Contractuales", ContractEmptyBadge: "Disponible para Contratación", ContractEmptyTitle: "Proyectos Contractuales y Consultoría", ContractEmptyDesc: "Contratos comerciales confidenciales, consultor\u00eda en hardware y mecatr\u00f3nica industrial. Estudios de caso pr\u00f3ximamente.",
		CP1Title: "Mec\u00e2nica Auto Moderna \u2014 Dise\u00f1o de Logotipo e Identidad Visual", CP1Desc: "Dise\u00f1o de logotipo corporativo e identidad visual a medida para Mec\u00e2nica Auto Moderna. Emblema automotriz \u00fanico que combina herramientas mec\u00e1nicas, s\u00edmbolo de servicio 24h y tipograf\u00eda moderna para rotulaci\u00f3n, uniformes y medios digitales.", CP1Domain: "Identidad Corporativa y Dise\u00f1o Vectorial",
		CP2Title: "Mec\u00e2nica Auto Moderna \u2014 Sistema Invoice Moderna", CP2Desc: "Soluci\u00f3n de software personalizada desarrollada para Mec\u00e2nica Auto Moderna. Automatiza la emisi\u00f3n de facturas electr\u00f3nicas y proformas, c\u00e1lculo de impuestos, \u00f3rdenes de reparaci\u00f3n y gesti\u00f3n de clientes con m\u00e1xima precisi\u00f3n y eficiencia operativa.", CP2Domain: "Software de Facturaci\u00f3n Electr\u00f3nica y Gesti\u00f3n Empresarial",
		BtnInstagramPost: "Ver Publicaci\u00f3n en Instagram", BtnLinkedInPost: "Ver Publicaci\\u00f3n en LinkedIn",
		CP3Title: "APACOV \u2014 Infraestructura de TI y Arquitectura Web", CP3Desc: "Desarrollo voluntario del portal web institucional integral y gesti\u00f3n de infraestructura tecnol\u00f3gica para APACOV, ONG dedicada a la protecci\u00f3n y amparo de ni\u00f1os hu\u00e9rfanos y vulnerables. Desempe\u00f1ando el rol de Administrador de TI para servidores, ciberseguridad y asuntos exteriores.", CP3Domain: "Infraestructura de TI \u00b7 Arquitectura Web \u00b7 Asuntos Exteriores",
		CP4Title: "APACOV \u2014 Dise\u00f1o de Logotipo e Identidad Humanitaria", CP4Desc: "Identidad visual pro bono dise\u00f1ada para la organizaci\u00f3n humanitaria APACOV. Creado en Adobe Illustrator y afinado en Samsung PenUp, el emblema representa manos de auxilio mutuo en solidaridad dentro de un arco protector de esperanza.", CP4Domain: "Identidad Visual Humanitaria \u00b7 Dise\u00f1o Vectorial",
		CP5Title: "SuPublic — Ingeniería de Chapa Metálica para Quiosco Automatizado de Agua Potable", CP5Desc: "Ingeniería mecánica contratada y diseño paramétrico de chapa metálica en SolidWorks para la iniciativa ecológica SuPublic, fundada por Gonca Özata bajo la mentoría del Dr. Yavuz Selim Taşpınar. Iniciado en Konya para mapear fuentes públicas de agua potable gratuita en Turquía y el mundo, combatiendo el residuo de botellas plásticas de un solo uso. Diseñó el chasis completo en chapa metálica, patrones de corte por láser y pliegues, y el nicho de recarga para el prototipo del quiosco expendedor.", CP5Domain: "Ingeniería Contratada · Diseño de Chapa Metálica · SolidWorks CAD y Prototipado",
		CPNoticeInfra: "Nota de Migraci\u00f3n de Servidores: Proceso activo de transici\u00f3n de infraestructura y migraci\u00f3n de apacov.org hacia servidores propios en curso, con nuevos dominios y correos corporativos.", BtnOfficialWebsite: "Sitio Web Oficial (apacov.org)", BadgeVolunteered: "Voluntariado \u00b7 Pro Bono",
		ExcavatorTitle: "Diseño de Excavadora Hidráulica (Elementos de Máquinas) — Inspirado en Komatsu PC200",
		ExcavatorDesc: "Diseño de ingeniería mecánica y análisis estructural por elementos finitos (FEA en ANSYS) para una excavadora hidráulica sobre orugas de 24 toneladas inspirada en la Komatsu PC200 LC-7/LC-8. Cálculos analíticos y verificación en 5 subsistemas críticos: pasadores y ejes de pluma/brazo/cuchara, rodamientos de giro y rodillos de oruga, uniones empernadas pretensadas, reductores planetarios y cilindros hidráulicos de alta presión.",
		ExcavatorDomain: "Diseño de Maquinaria Pesada · Elementos de Máquinas · ANSYS FEA & Hidráulica",
		SafrenTitle: "SafRen (anteriormente Ce-Saf) — Podómetro Inteligente para Calzado e IoT Vestible",
		SafrenDesc: "Sistema vestible colaborativo de podómetro y telemetría de marcha para calzado deportivo, desarrollado en colaboración interdisciplinaria con las estudiantes de Diseño de Calzado Safina Najmudinova y Ceren Başaran, quienes aportaron el calzado y determinaron la ubicación ergonómica óptima para el hardware electrónico (pantalla OLED LCD, pulsadores táctiles de impacto y microcontrolador). Cuenta con microcontrolador Waveshare ESP32-S3-Zero, sensor inercial MPU6050 de 6 ejes, servidor web Wi-Fi SoftAP integrado y sincronización inalámbrica en vivo con aplicación móvil desarrollada en Flutter.",
		SafrenDomain: "IoT Vestible · Análisis de Marcha · Firmware ESP32-S3 · Flutter Móvil",
		STM32Title: "Desafío de la Placa STM32 Discovery — Robótica Móvil Autónoma",
		STM32Desc: "Ingeniería de robótica autónoma de doble desafío impulsada por la placa STM32F407VGT6 ARM Cortex-M4. Desafío 1 (Vize): Matriz de sensores LDR con ADC de 12 bits y control PWM diferencial TIM3 para rastrear luces aleatorias en un área de 1 m² y detenerse con precisión milimétrica sin colisión al primer intento. Desafío 2 (Final): Extintor de velas autónomo con ventilador de 12V de alto flujo, centrado estratégico y giros de vórtice de 360°—apagando las 5 velas perimetrales en menos de 10 segundos (un tercio del tiempo límite).",
		STM32Domain: "Sistemas Embebidos · STM32 ARM Cortex-M4 · C Bare-Metal / HAL · Robótica Autónoma",
		ROSTitle: "Robot Móvil Autónomo ROS Noetic — Puente TCP/IP ESP32 y Cinemática Diferencial",
		ROSDesc: "Robot móvil autónomo desarrollado en colaboración con el ingeniero mecatrónico Mehmet Çakaroğlu bajo la supervisión del Dr. Okan Uyar. Conecta ROS Noetic sobre Linux con el microcontrolador de doble núcleo ESP32 mediante un socket Wi-Fi TCP/IP de baja latencia. Cuenta con cinemática diferencial, odometría en bucle cerrado con encoders de 420 pulsos, IMU MPU6050 de 6 ejes calibrada, mecanismo de seguridad deadman de 500 ms y simulación de gemelo digital en Gazebo y RViz superando los desafíos con puntuación 100/100.",
		ROSDomain: "Robótica Móvil Autónoma · ROS Noetic · Cinemática Diferencial · Gazebo y RViz",
		BtnLinkedInVideo: "Video en LinkedIn", BtnGitHubRepo: "Repositorio GitHub", CertsTitle: "Certificaciones y Reconocimientos",
		VHeroBadge: "Candidatura al Premio James Dyson 2026 · Innovación Médica",
		VHeroTitle: "VELUM — Esterilla Preventiva de Úlceras por Presión",
		VHeroSubtitle: "La primera esterilla sensora biomédica pasiva y continua del mundo que detecta úlceras por presión de 2 a 5 días antes de ser visibles, de forma equitativa y económica.",
		VStatCost: "+11.000 M$ USD", VStatCostDesc: "Coste anual en hospitales de EE.UU. por úlceras por presión evitables",
		VStatEarly: "2 a 5 Días Antes", VStatEarlyDesc: "Alerta previa a daños visibles mediante humedad subepidérmica",
		VStatBOM: "< 50 $ USD", VStatBOMDesc: "Coste objetivo de componentes para equipar todas las camas",
		VStatEquity: "100% Equitativo", VStatEquityDesc: "Medición bioeléctrica independiente del tono de piel (Fitzpatrick I–VI)",
		VProblemTitle: "Un Protocolo Clínico Deficiente que Falla al Paciente",
		VProblemSubtitle: "La inspección visual detecta heridas cuando la necrosis tisular ya es irreversible.",
		VProblemP1: "Las úlceras por presión afectan a más del 12,8% de los pacientes hospitalizados en todo el mundo. El estándar actual depende de rondas manuales de enfermería cada 2 horas e inspecciones visuales subjetivas. Sin embargo, el daño celular comienza en el tejido subepidérmico entre 48 y 120 horas antes de que aparezca eritema en la piel.",
		VProblemP2: "Además, la detección visual falla clínicamente en pieles oscuras (Fitzpatrick IV–VI). Los escáneres SEM manuales comerciales cuestan entre 3.000 $ y 10.000 $ y requieren personal entrenado.",
		VProblemPoint1: "Reactivo por Definición: El enrojecimiento cutáneo confirma que la necrosis celular ya ha comenzado.",
		VProblemPoint2: "Disparidad Racial Documentada: La inspección visual no detecta úlceras en tonos de piel oscuros.",
		VProblemPoint3: "Sobrecarga de Personal y Costes: Los dispositivos manuales no evitan la fatiga de enfermería.",
		VSolutionTitle: "La Solución VELUM: Bioimpedancia Continua",
		VSolutionSubtitle: "Monitorización tisular pasiva 24/7 colocada discretamente bajo la sábana del paciente.",
		VSolutionP1: "VELUM es una esterilla sensora ultrafina y flexible de poliimida ubicada bajo el paciente. Una matriz de electrodos de plata y carbono serigrafiados realiza espectroscopia de bioimpedancia eléctrica (1–100 kHz) en 16 puntos anatómicos críticos cada 5 minutos.",
		VSolutionPoint1: "Despliegue Instantáneo: Se desliza bajo la sábana y se enciende, sin perturbar al paciente.",
		VSolutionPoint2: "Detección de Humedad Subepidérmica: Una caída de ≥ 15% en la impedancia avisa del edema días antes de la herida.",
		VSolutionPoint3: "Alertas Inmediatas BLE / Wi-Fi: Las enfermeras reciben avisos que indican la zona anatómica exacta que requiere giro.",
		VSimTitle: "Simulador Interactivo de la Esterilla de Bioimpedancia",
		VSimSubtitle: "Mapeo anatómico de presión en 16 zonas y simulación del barrido de frecuencia del AD5933.",
		VSimNormal: "1. Estado Normal", VSimIschemia: "2. Presión Sacra Sostenida",
		VSimOedema: "3. Alerta de Edema Fase 0", VSimReset: "4. Reposicionamiento del Paciente",
		VSimStatusAlert: "CRÍTICO: ¡Edema subepidérmico detectado! Reposicione la zona sacra de inmediato",
		VSimStatusNormal: "NOMINAL: Perfusión tisular adecuada en todas las zonas anatómicas",
		VTechTitle: "Arquitectura Mecatrónica en Vista Explosionada",
		VTechSubtitle: "Electrónica flexible de alta precisión diseñada para desinfección hospitalaria y fabricación masiva.",
		VLayer1Title: "1. Encapsulado en TPU Médico y Silicona de Platino",
		VLayer1Desc: "Membrana biocompatible (ISO 10993) permeable al vapor e impermeable a líquidos, resistente a desinfectantes clínicos.",
		VLayer2Title: "2. Matriz de Electrodos Serigrafiados de Plata y Carbono",
		VLayer2Desc: "Película de poliimida DuPont Kapton® de 50 µm con pistas conductoras serigrafiadas bajo puntos óseos críticos.",
		VLayer3Title: "3. Convertidor de Impedancia AD5933 y MUX de 16 Canales",
		VLayer3Desc: "Analog Devices AD5933 con fuente de corriente Howland (<1 mA CA) y multiplexor CD74HC4067 para barrido ultrarrápido.",
		VLayer4Title: "4. Módulo de Cómputo ESP32-S3 Dual-Core y Conectividad Inalámbrica",
		VLayer4Desc: "Microcontrolador a 240MHz con FreeRTOS, calibración dinámica, red BLE 5.0 y telemetría Wi-Fi cifrada para historiales clínicos.",
		VCompTitle: "Ventaja Competitiva y Clínica",
		VCompSubtitle: "Por qué VELUM establece un estándar médico y económico insuperable.",
		VCompColMetric: "Criterio / Métrica", VCompColVisual: "Inspección Visual", VCompColSEM: "Sonda SEM Manual", VCompColMattress: "Colchones de Alivio", VCompColVelum: "Esterilla VELUM",
		VRoadmapTitle: "Hoja de Ruta Clínica y Comercial",
		VRoadmapSubtitle: "Desde el prototipo de banco hasta ensayos clínicos en UCI y autorización FDA 510(k).",
		VGalleryTitle: "Renders de Hardware y Diseños Técnicos",
		VGallerySubtitle: "Ilustraciones de producto en alta resolución y diagramas de montaje mecatrónico.",
		VDownloadPitch: "Descargar Presentación para Inversores (PDF)",
		VViewGitHub: "Ver Repositorio en GitHub",
		FooterRights: "Todos los derechos reservados.",
		FooterBuiltWith: "",
	},
	"de": {
		LangCode: "de", LangName: "German", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Startseite", NavProjects: "Projekte", NavExperience: "Erfahrung", NavAbout: "Über mich", NavSkills: "Kompetenzen", NavContact: "Kontakt", NavEducation: "Ausbildung", NavHonors: "Auszeichnungen", NavCertificates: "Zertifikate", NavActivities: "Workshops", NavLanguages: "Sprachen", NavGitHub: "GitHub", NavVelum: "VELUM Projekt",
		HeroBadge: "Mechatronik-Ingenieur · Biomedizinische Hardware · Eingebettete Systeme · Polyglott",
		AuthorRole: "B.Sc. Mechatronik · Embedded- & IoT-Entwickler",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Entwicklung intelligenter biomedizinischer Mechatronik, hochpräziser Sensoren und Edge-IoT-Systeme zum Schutz von Menschenleben.",
		HeroBtnExplore: "Projekte Erkunden", HeroBtnContact: "Kontakt Aufnehmen",
		StatProjects: "18 Technische Projekte", StatSoloProjects: "Solo-Projekte", StatTeamProjects: "Team-Projekte", StatContractProjects: "Auftragsprojekte", StatWorkshops: "Workshops & Aktivitäten", StatAward: "Mechatronik-Ingenieurwesen", StatLanguages: "7 Fließende Sprachen", StatFocus: "ESP32 / Bioimpedanz / CAD",
		ProjectsTitle: "Ingenieur-Portfolio",
		ProjectsSubtitle: "Präventive biomedizinische Hardware, assistive Medizingeräte, Fluidik und mechatronische Systeme.",
		StatusActive: "Abgeschlossen & Interaktiv", StatusQueue: "In Vorbereitung",
		BtnViewProject: "Projekt Anzeigen", BtnPreview: "Projektdetails",
		BadgeFeatured: "Hauptprojekt", BadgeBiomedical: "Biomedizinische Hardware", BadgeHardware: "Mechatronik", BadgeRobotics: "Robotik & Automation",
		P1Title: "VELUM (vormals SentIO)",
		P1Desc: "Kontinuierliche 24/7 Sensormatte zur Messung von subepidermaler Feuchtigkeit und Bioimpedanz, die Dekubitus 2–5 Tage vor sichtbaren Wunden erkennt.",
		P1Domain: "Biomedizinische Mechatronik · Bioimpedanzspektroskopie",
		P2Title: "FLOWGUARD",
		P2Desc: "Automatisiertes mechatronisches Mikroplastik-Filtersystem für Waschmaschinen zur Verhinderung synthetischer Gewässerverschmutzung.",
		P2Domain: "Fluidische Mechatronik · Mikroplastik-Filtration",
		P3Title: "SOLUS",
		P3Desc: "Tragbares medizinisches Pflaster zur kontinuierlichen Überwachung der Kerntemperatur zur Früherkennung von Hitzschlägen bei extremen Arbeitsbedingungen.",
		P3Domain: "Tragbare Biosensoren · Präventive Physiologie",
		P4Title: "DENOVA (vormals Avicenna)",
		P4Desc: "Nicht-invasiver optischer Kariesdetektor mit multispektraler Fluoreszenzbeleuchtung zur Frühdiagnose von Schmelzentkalkungen.",
		P4Domain: "Biophotonik · Optische Diagnostik",
		P5Title: "RENOVO",
		P5Desc: "Tragbare hydraulische Katastrophenschutzpresse, die Trümmerschutt vor Ort in ineinandergreifende Bausteine umwandelt.",
		P5Domain: "Schwere Mechanik & Hydraulik · Katastrophenhilfe",
		P6Title: "CURA SERO",
		P6Desc: "Chirurgischer Mikrozug-Wundverschlussstreifen mit gleichmäßiger Zugkraftverteilung zur Vermeidung von Narbenbildung.",
		P6Domain: "Biomechanisches Gerät · Chirurgische Technik",
		P7Title: "WELSTECH — Wearable Life Saver Technology",
		P7Desc: "Ultrakompakter tragbarer Vitalzeichenmonitor mit optischem PPG-Biosensor, ESP32-S3-Recheneinheit, lokaler Wi-Fi-Telemetrie und eskalierendem akustischem Alarm.",
		P7Domain: "Biomedizinisches Wearable-IoT · Kardiale Biosensorik",
		P8Title: "SENDIT",
		P8Desc: "Autonomes Flughafen-Gepäck- und Luftpost-Sortier- und Transportsystem zur Optimierung der Abfertigungszeiten.",
		P8Domain: "Autonome Robotik · Logistik-Automation",
		P9Title: "SMARTWASTE",
		P9Desc: "Mechatronisches robotergestütztes Abfalltrenn-System mit integrierter Spektrometrie und Computer-Vision für automatisiertes Recycling.",
		P9Domain: "Computer Vision · Mechatronische Aktorik",
		P10Title: "IMAGE CLASSIFIER — Tiefes Faltungsnetzwerk (CNN)",
		P10Desc: "Tiefes Convolutional Neural Network (CNN) in TensorFlow/Keras, trainiert auf über 9.600 benutzerdefinierten Bildern mit 96×96 Vorverarbeitung und Desktop-Inferenz-GUI.",
		P10Domain: "Deep Learning · Computer Vision & Edge AI",
		P11Title: "LUMIERE — Intelligentes Straßenbeleuchtungs-Störungsmelde- & Ortungssystem",
		P11Desc: "Kompaktes IoT-Gerät für Straßenlaternen mit LDR-Helligkeitssensor, GY-NEO6MV2 GPS-Ortung, ESP32 Wi-Fi SoftAP-Telemetrie und Summer-Alarm zur vollautomatisierten Störungsmeldung an Kommunen.",
		P11Domain: "Smart-City IoT · Eingebettete Firmware & Telemetrie",
		P12Title: "NOTCHA — Minimalistische Schwebende Dynamische Zeitkerbe",
		P12Desc: "Entwickelt in den Jahren 2021/2022, nachdem Windows 11 das Verschieben der Taskleiste an den oberen Bildschirmrand entfernt hatte; NOTCHA ist ein ressourcenschonendes Desktop-Tool mit schwebender Zeitanzeige am oberen Rand, Echtzeituhr und intelligenter Hover-Erweiterung.",
		P12Domain: "Desktop-Ergonomie · Individuelle Benutzeroberflächen-Tools",
		P13Title: "SPO2 — 2-Lagen-Puls-Oximeter-Leiterplattendesign",
		P13Desc: "2-Lagen-Leiterplatte für ein Pulsoximeter entwickelt in KiCad 10.0; kombiniert maßgeschneiderte Footprints und Symbole, reale Fertigungstoleranzen, optische MAX30102-Biosensorik, SSD1306-OLED und einen Arduino Nano-Kern.",
		P13Domain: "PCB-Design & EDA · Biomedizinische Hardware · KiCad 10.0",
		P14Title: "RESPIRO",
		P14Desc: "Kostengünstiges, hochzuverlässiges biomedizinisches adaptives Beatmungsgerät für den Noteinsatz und ressourcenbeschränkte Kliniken.",
		P14Domain: "Biomedizinische Mechatronik · Intensivpflege",
		AboutTitle: "Über Ukacha Juma",
		AboutSubtitle: "Mechatronik-Ingenieur · Hardware-Software-Architekt · Polyglott",
		AboutP1: "Als Mechatronik-Ingenieur treibt mich die Überzeugung an, dass moderne Hardware-Entwicklung menschliches Leid verhindern sollte, anstatt es nur im Nachhinein zu behandeln. Meine Arbeit vereint mechanisches CAD/CAM-Design, analoge Präzisionsschaltungen, Firmware auf Bare-Metal und FreeRTOS sowie Telemetriesysteme.",
		AboutP2: "Mit Projekten wie Dekubitus-Präventionsmatten (VELUM), Mikroplastikfiltration (FLOWGUARD) und thermischen Biosensoren (SOLUS) entwickle ich mechatronische Systeme von mathematischen Grundlagen bis zu serienreifen Prototypen.",
		AboutPolyglot: "Mehrsprachige Kommunikation & Globale Zusammenarbeit",
		AboutPolyglotDesc: "Ich spreche fließend Englisch, Türkisch, Arabisch, Portugiesisch, Spanisch, Deutsch und Französisch. Als Polyglott kann ich internationale Entwicklerteams leiten und klinische Studien in ihren Originalfassungen analysieren.",
		SkillsTitle: "Technische Kompetenzen & Toolchain",
		ExpTitle: "Berufserfahrung", ExpSubtitle: "Praktische Industrieerfahrung, CNC-Diagnostik, Automatisierungswartung und internationale Übersetzung.", EduTitle: "Ausbildung & Bildungsweg", AwardsTag: "AKADEMISCHE & ROBOTIK-AUSZEICHNUNGEN", AwardsTitle: "Auszeichnungen & Ehrungen", AwardsSubtitle: "Institutionelle Würdigungen, Platzierungen bei nationalen Robotik-Olympiaden und Auszeichnungen für Ingenieurleistungen.", AwardsCardTitle: "İTÜRO 2026 Platzierung — Ehrenplakette der Fakultätsdekanin", AwardsCardDesc: "Ausgezeichnet mit einer Ehrenplakette durch die Dekanin Prof. Dr. Humar Kahramanli Örnek bei der Graduierungsfeier der Technologiefakultät der Selçuk-Universität (2025–26) in Anerkennung der Platzierung des STWolf-Teams bei der Robotik-Olympiade der TU Istanbul (İTÜRO 2026).", AwardsViewPost: "LinkedIn-Beitrag & Auszeichnung Ansehen", AwardsDeanTitle: "Dekanin der Fakultät für Technologie", AwardsDeanName: "Prof. Dr. Humar Kahramanlı Örnek", AwardsTeamName: "Team STWolf · Selçuk-Universität", AwardsOlympiad: "Robotik-Olympiade der Technischen Universität Istanbul (İTÜRO 2026)", AwardsCeremony: "Abschlussfeier 2025–26 · Selçuk-Universität", AwardsQuoteLabel: "Offizielle Plaketteninschrift", TabEducation: "Ausbildung & Studium", TabAwards: "Auszeichnungen", TabCertificates: "Zertifikate & Lizenzen", TabActivities: "Workshops", TabLanguages: "Sprachen", ActivitiesTitle: "Extrakurrikulare & Technische Aktivitäten", ActivitiesSubtitle: "Technologiefestival-Ausstellungen, Gemeinschaftsnetzwerke und KI-gestützte akademische Forschung.", TeknoselTitle: "Selçuklu Teknosel Technologie-Festival — Kategorie Universität", TeknoselDesc: "Präsentation des mechatronischen KonyaWaste-Prototyps an unserem Messestand in der Kategorie Freies Projekt - Universitätsebene mit persönlichem Rundgang für den Bürgermeister von Selçuklu.", TeknoselMayorBadge: "Persönliche Präsentation vor dem Bürgermeister", AiResearchTitle: "Quellengestützte KI in der akademischen Forschung (TADD-Workshop)", AiResearchDesc: "Absolvierung des zweitägigen Intensiv-Workshops von TADD unter Leitung von Yunus Emre Kıymaz zur Beherrschung von Zotero, ResearchRabbit und NotebookLM, um KI-Halluzinationen bei mechatronischen Literaturrecherchen vollständig auszuschließen.", AiResearchBadge: "Quellengestützte KI ohne Halluzinationen", IfmWorkshopTitle: "Erweiterung der Automatisierungskompetenz mit ifm electronic & CoDeSys (48h Intensivtraining)", IfmWorkshopDesc: "Erfolgreicher Abschluss des 48-stündigen SPS-Intensivtrainings an der Selçuk-Universität, veranstaltet von der türkischen Zentrale von ifm electronic. Praxiserfahrung in SPS-Architektur, CoDeSys V3.5 (SP18 & SP19), Modbus- und CANbus-Feldbusprotokollen sowie Vor-Ort-Fehlerbehebung.", IfmWorkshopBadge: "⚙️ ifm electronic Türkiye · Selçuk-Universität", IfmBtnLinkedIn: "ifm-Workshop auf LinkedIn ansehen", TabSoloProjects: "Solo-Projekte", TabTeamProjects: "Team-Projekte", TabContractProjects: "Auftragsprojekte", ContractEmptyBadge: "Verfügbar für Aufträge", ContractEmptyTitle: "Auftrags- & Industrieprojekte", ContractEmptyDesc: "Propriet\u00e4re Kundenprojekte, Hardware-Beratung und mechatronische Industrieentwicklungen. Fallstudien folgen in K\u00fcrze.",
		CP1Title: "Mec\u00e2nica Auto Moderna \u2014 Firmenlogo & Corporate Design", CP1Desc: "Individuelles Firmenlogo und visuelle Markenidentit\u00e4t f\u00fcr den Kfz-Fachbetrieb Mec\u00e2nica Auto Moderna. Speziell entwickeltes Emblem mit Kfz-Werkzeugen, 24-Stunden-Service-Symbolik und moderner Typografie f\u00fcr Beschilderung, Arbeitskleidung und Digitalmedien.", CP1Domain: "Corporate Identity & Vektorgrafik-Design",
		CP2Title: "Mec\u00e2nica Auto Moderna \u2014 Rechnungs- und Verwaltungssystem Invoice Moderna", CP2Desc: "Ma\u00dfgeschneiderte Softwarel\u00f6sung f\u00fcr den Kfz-Betrieb Mec\u00e2nica Auto Moderna. Automatisiert die Erstellung elektronischer Rechnungen (Proforma), Steuerberechnungen, Kundenverwaltung und Werkstattauftr\u00e4ge mit h\u00f6chster Pr\u00e4zision und Effizienz.", CP2Domain: "Unternehmenssoftware & Elektronische Rechnungsstellung",
		BtnInstagramPost: "Instagram-Beitrag ansehen", BtnLinkedInPost: "LinkedIn-Beitrag ansehen",
		CP3Title: "APACOV \u2014 IT-Infrastruktur & Web-Architektur", CP3Desc: "Ehrenamtliche Entwicklung des institutionellen Webportals und ganzheitliche Betreuung der IT-Infrastruktur f\u00fcr die Hilfsorganisation APACOV zum Schutz von Waisenkindern. Zust\u00e4ndig f\u00fcr Serverbereitstellung, Cybersicherheit und internationale Beziehungen.", CP3Domain: "IT-Infrastruktur \u00b7 Web-Architektur \u00b7 Au\u00dfenbeziehungen",
		CP4Title: "APACOV \u2014 Humanit\u00e4res Logo & Markenidentit\u00e4t", CP4Desc: "Pro-bono-Markenidentit\u00e4t f\u00fcr die Hilfsorganisation APACOV. Gestaltet in Adobe Illustrator und finalisiert mit Samsung PenUp; das Emblem symbolisiert helfende H\u00e4nde der Solidarit\u00e4t innerhalb eines sch\u00fctzenden Bogens der Hoffnung.", CP4Domain: "Humanit\u00e4re Markenidentit\u00e4t \u00b7 Vektorgrafik",
		CP5Title: "SuPublic — Blechkonstruktion für automatisierten Trinkwasser-Kiosk", CP5Desc: "Vertragsbasierte mechanische Konstruktion und SolidWorks-Blechdesign für die Umweltinitiative SuPublic, gegründet von Gonca Özata unter der wissenschaftlichen Betreuung von Dr. Yavuz Selim Taşpınar. Das in Konya gestartete Projekt kartiert kostenlose, saubere Trinkwasserbrunnen in der Türkei und weltweit, um Einweg-Plastikflaschenmüll zu vermeiden. Konstruktion des gesamten Blechgehäuses, der Biege- und Laserschnittmuster sowie der ergonomischen Nachfüllnische für den Kiosk-Prototyp.", CP5Domain: "Vertragsingenieurwesen · Blechdesign · SolidWorks CAD & Prototyping",
		CPNoticeInfra: "Server-Migrationshinweis: Aktiver Infrastruktur- und Serverumzug von apacov.org auf eigene Server sowie Beschaffung zus\u00e4tzlicher Domains und E-Mail-Dienste laufen derzeit.", BtnOfficialWebsite: "Offizielle Website (apacov.org)", BadgeVolunteered: "Ehrenamtlich \u00b7 Pro Bono",
		ExcavatorTitle: "Baggerkonstruktion (Maschinenelemente) — Inspiriert von Komatsu PC200 LC-7/LC-8",
		ExcavatorDesc: "Umfassende mechanische Systemkonstruktion und strukturelle FEA-Analyse in ANSYS für einen 24-Tonnen-Hydraulikbagger, inspiriert vom Komatsu PC200 LC-7/LC-8. Analytische Berechnungen und Katalogauswahl für 5 zentrale Maschinensubsysteme: Bolzen und Wellen für Ausleger, Stiel und Löffel, Großwälzlager und Laufwerksrollen, hochfeste Schraubenverbindungen, Planetengetriebe und Hochdruckhydraulikzylinder.",
		ExcavatorDomain: "Baumaschinenkonstruktion · Maschinenelemente · ANSYS FEA & Hydraulik",
		SafrenTitle: "SafRen (ehemals Ce-Saf) — Intelligenter Schuh-Schrittzähler & Wearable IoT-System",
		SafrenDesc: "Kollaboratives Wearable-Projekt für intelligente Schrittzählung und Ganganalyse am Schuh, entwickelt in interdisziplinärer Zusammenarbeit mit den Schuhdesign-Studentinnen Safina Najmudinova und Ceren Başaran. Sie stellten die Schuhe bereit und optimierten die ergonomische Platzierung der Hardwarekomponenten (OLED LCD, Taster, Mikrocontroller). Ausgestattet mit Waveshare ESP32-S3-Zero, MPU6050 6-Achs-Trägheitssensor, integriertem SoftAP Wi-Fi-Server und bidirektionaler Echtzeitsynchronisation mit einer maßgeschneiderten Flutter-App.",
		SafrenDomain: "Wearable IoT · Ganganalyse · ESP32-S3 Firmware · Flutter Mobile",
		STM32Title: "STM32 Discovery Board Challenge — Autonome Mobile Robotik",
		STM32Desc: "Autonome Robotik-Entwicklung auf dem STM32F407VGT6 ARM Cortex-M4 Discovery Board. Herausforderung 1: Mehrkanal-LDR-Sensorarray mit 12-Bit-ADC-Gradientenerfassung und TIM3-PWM-Differentialantrieb zur punktgenauen Annäherung an Lichtquellen im 1 m² Areal ohne Kollision im ersten Versuch. Herausforderung 2: Autonomer Kerzenlöscher mit 12V-Hochleistungslüfter, zentraler Arena-Positionierung und 360°-Wirbeldrehung – Löschen aller 5 Kerzen in unter 10 Sekunden (1/3 des 30s-Zeitlimits).",
		STM32Domain: "Embedded Systems · STM32 ARM Cortex-M4 · Bare-Metal C / HAL · Autonome Robotik",
		ROSTitle: "Autonomer mobiler Roboter mit ROS Noetic — ESP32 TCP/IP-Bridge & Differentialantrieb",
		ROSDesc: "Autonomer mobiler Roboter, entwickelt in Kooperation mit Mechatronikingenieur Mehmet Çakaroğlu unter Leitung von Dr. Okan Uyar. Verbindet ROS Noetic unter Linux mit einem ESP32-Dual-Core-Mikrocontroller über eine latenzarme Wi-Fi TCP/IP-Socket-Bridge. Verfügt über Differentialantriebskinematik, Closed-Loop-Odometrie mit 420-Pulse-Quadratur-Encodern, kalibriertes MPU6050-6-DOF-IMU, einen 500-ms-Deadman-Schalter und vollständige Gazebo/RViz-Digital-Twin-Trajektorienverfolgung mit 100/100 Bewertungspunkten.",
		ROSDomain: "Autonome mobile Robotik · ROS Noetic · Differentialkinematik · Gazebo & RViz",
		BtnLinkedInVideo: "LinkedIn Video-Präsentation", BtnGitHubRepo: "GitHub Repository", CertsTitle: "Zertifikate & Auszeichnungen",
		VHeroBadge: "James Dyson Award 2026 Einreichung · Healthcare Innovation",
		VHeroTitle: "VELUM — Dekubitus-Präventionsmatte",
		VHeroSubtitle: "Die weltweit erste kontinuierliche, passive biomedizinische Sensormatte, die Druckgeschwüre 2–5 Tage vor ihrer Sichtbarkeit erkennt — für alle Hauttöne und bezahlbar.",
		VStatCost: "11+ Mrd. $ USD", VStatCostDesc: "Jährliche US-Krankenhauskosten für vermeidbare Druckgeschwüre",
		VStatEarly: "2–5 Tage Früher", VStatEarlyDesc: "Frühwarnung durch Erfassung subepidermaler Feuchtigkeit",
		VStatBOM: "< 50 $ USD", VStatBOMDesc: "Ziel-Stücklistenkosten für flächendeckenden Betteneinsatz",
		VStatEquity: "100% Chancengleich", VStatEquityDesc: "Hauttonunabhängige elektrische Bioimpedanzmessung (Fitzpatrick I–VI)",
		VProblemTitle: "Ein Unzureichendes System, das Patienten Gefährdet",
		VProblemSubtitle: "Die visuelle Inspektion erkennt Wunden erst, wenn Gewebsnekrosen bereits irreversibel sind.",
		VProblemP1: "Druckgeschwüre (Dekubitus) betreffen über 12,8% aller stationären Patienten weltweit. Die derzeitige Standardversorgung beruht auf manuellem Umlagern alle 2 Stunden und subjektiven Sichtkontrollen. Gewebeschäden beginnen jedoch 48 bis 120 Stunden vor sichtbarer Rötung in tieferen Schichten.",
		VProblemP2: "Zudem versagt die visuelle Rötungserkennung bei dunkleren Hauttönen (Fitzpatrick IV–VI) regelmäßig. Kommerzielle Hand-SEM-Scanner kosten 3.000–10.000 $ und erfordern manuelle Bedienung.",
		VProblemPoint1: "Reaktiv: Hautrötung zeigt an, dass zelluläre Schäden bereits eingetreten sind.",
		VProblemPoint2: "Dokumentierte Disparität: Sichtkontrollen übersehen Dekubitus Stadium 1 bei dunkler Haut.",
		VProblemPoint3: "Kosten & Pflegekräftemangel: Manuelle Scanner sind teuer und überlasten das Personal.",
		VSolutionTitle: "Die VELUM-Lösung: Kontinuierliche Bioimpedanz",
		VSolutionSubtitle: "Passive, permanente Gewebeüberwachung 24/7 unsichtbar unter dem Bettlaken.",
		VSolutionP1: "VELUM ist eine ultradünne, flexible Polyimid-Sensormatte unter dem Patienten. Ein siebgedrucktes Silber-Kohlenstoff-Elektrodennetz misst alle 5 Minuten die Bioimpedanz (1–100 kHz) an 16 kritischen Körperstellen.",
		VSolutionPoint1: "Sofort Einsatzbereit: Unter das Laken legen, einschalten, ungestörter Patientenschlaf.",
		VSolutionPoint2: "Subepidermale Feuchtigkeitserkennung: Ein Impedanzabfall von 15%+ meldet Zellschwellung Tage vor der Wunde.",
		VSolutionPoint3: "Drahtlose Alarme via BLE / Wi-Fi: Pflegekräfte erhalten zielgenaue Hinweise, welche Stelle umzulagern ist.",
		VSimTitle: "Interaktiver Bioimpedanz-Matten-Simulator",
		VSimSubtitle: "Echtzeit-16-Zonen-Druckverteilung und Simulation des AD5933-Frequenzsweeps.",
		VSimNormal: "1. Normalzustand", VSimIschemia: "2. Anhaltender Kreuzbeindruck (Ischämie)",
		VSimOedema: "3. Ödem-Alarm Stadium 0", VSimReset: "4. Pflege-Umlagerung",
		VSimStatusAlert: "KRITISCH: Subepidermales Ödem erkannt! Kreuzbein sofort umlagern",
		VSimStatusNormal: "NOMINAL: Gewebeperfusion in allen anatomischen Zonen optimal",
		VTechTitle: "Explosionsdarstellung der Mechatronischen Architektur",
		VTechSubtitle: "Präzise flexible Elektronik, ausgelegt für klinische Desinfektion und Serienfertigung.",
		VLayer1Title: "1. Medizinisches TPU & Platin-Silikon",
		VLayer1Desc: "Dampfdurchlässige, flüssigkeitsdichte biokompatible Hülle (ISO 10993), beständig gegen Klinik-Desinfektionsmittel.",
		VLayer2Title: "2. Siebgedrucktes Silber/Kohlenstoff-Elektrodennetz",
		VLayer2Desc: "Leiterbahnen auf 50 µm DuPont Kapton® Polyimidfolie, angeordnet in Vier-Punkt-Arrays unter Knochenvorsprüngen.",
		VLayer3Title: "3. AD5933 Impedanzwandler & 16-Kanal-Multiplexer",
		VLayer3Desc: "Analog Devices AD5933 mit Howland-Stromquelle (<1 mA Wechselstrom) und CD74HC4067 für sekundenschnelle Scans.",
		VLayer4Title: "4. ESP32-S3 Dual-Core Compute & Funkmodul",
		VLayer4Desc: "240MHz Xtensa-Prozessor mit FreeRTOS, Basiskalibrierung, BLE 5.0 Mesh und verschlüsselter Wi-Fi-Krankenhaus-Telemetrie.",
		VCompTitle: "Klinischer & Wirtschaftlicher Wettbewerbsvorteil",
		VCompSubtitle: "Warum VELUM einen unübertroffenen medizinischen und ökonomischen Mehrwert schafft.",
		VCompColMetric: "Merkmal / Kriterium", VCompColVisual: "Sichtprüfung", VCompColSEM: "Hand-SEM-Sonde", VCompColMattress: "Druckentlastungsmatratzen", VCompColVelum: "VELUM Matte",
		VRoadmapTitle: "Klinische & Kommerzielle Roadmap",
		VRoadmapSubtitle: "Vom Laborprototyp über klinische ICU-Studien bis zur FDA 510(k)-Zulassung.",
		VGalleryTitle: "Hardware-Renderings & Technische Zeichnungen",
		VGallerySubtitle: "Hochauflösende Produktdarstellungen und mechatronische Explosionszeichnungen.",
		VDownloadPitch: "Pitch-Deck herunterladen (PDF)",
		VViewGitHub: "GitHub-Repository aufrufen",
		FooterRights: "Alle Rechte vorbehalten.",
		FooterBuiltWith: "",
	},
	"fr": {
		LangCode: "fr", LangName: "French", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Accueil", NavProjects: "Projets", NavExperience: "Expérience", NavAbout: "À propos", NavSkills: "Compétences", NavContact: "Contact", NavEducation: "Formation", NavHonors: "Prix", NavCertificates: "Certificats", NavActivities: "Workshops", NavLanguages: "Langues", NavGitHub: "GitHub", NavVelum: "Projet VELUM",
		HeroBadge: "Ingénieur Mécatronique · Matériel Biomédical · Systèmes Embarqués · Polyglotte",
		AuthorRole: "B.Sc. en Génie Mécatronique · Développeur Systèmes Embarqués & IoT",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Conception de systèmes mécatroniques biomédicaux intelligents, d'instrumentation sensorielle de précision et d'appareils IoT pour sauver des vies.",
		HeroBtnExplore: "Explorer les Projets", HeroBtnContact: "Me Contacter",
		StatProjects: "18 Projets d'Ingénierie", StatSoloProjects: "Projets Solo", StatTeamProjects: "Projets en Équipe", StatContractProjects: "Projets sous Contrat", StatWorkshops: "Workshops & Activités", StatAward: "Génie Mécatronique", StatLanguages: "7 Langues Maîtrisées", StatFocus: "ESP32 / Bioimpédance / CAD",
		ProjectsTitle: "Portfolio d'Ingénierie",
		ProjectsSubtitle: "Matériel biomédical préventif, dispositifs d'assistance, automatisation fluidique et systèmes mécatroniques.",
		StatusActive: "Terminé et Interactif", StatusQueue: "En File d'Attente",
		BtnViewProject: "Voir le Projet", BtnPreview: "Détails du Projet",
		BadgeFeatured: "Projet Phare", BadgeBiomedical: "Matériel Biomédical", BadgeHardware: "Mécatronique", BadgeRobotics: "Robotique & Automatisation",
		P1Title: "VELUM (anciennement SentIO)",
		P1Desc: "Tapis capteur continu 24/7 d'humidité sous-épidermique et de bioimpédance électrique détectant les escarres 2 à 5 jours avant les lésions visibles.",
		P1Domain: "Mécatronique Biomédicale · Spectroscopie de Bioimpédance",
		P2Title: "FLOWGUARD",
		P2Desc: "Système mécatronique automatisé de filtration des microplastiques pour lave-linge évitant la pollution des écosystèmes marins.",
		P2Domain: "Mécatronique Fluidique · Séparation des Microplastiques",
		P3Title: "SOLUS",
		P3Desc: "Patch portable de surveillance continue de la température corporelle centrale pour l'alerte précoce des coups de chaleur en milieu extrême.",
		P3Domain: "Biocapteurs Portables · Physiologie Préventive",
		P4Title: "DENOVA (anciennement Avicenna)",
		P4Desc: "Détecteur optique non invasif de caries dentaires utilisant l'éclairage multispectral et la fluorescence pour le diagnostic précoce.",
		P4Domain: "Biophotonique · Diagnostic Optique",
		P5Title: "RENOVO",
		P5Desc: "Presse hydraulique portable de secours transformant les gravats de séismes en briques de construction emboîtables sur place.",
		P5Domain: "Mécanique Lourde & Hydraulique · Secours d'Urgence",
		P6Title: "CURA SERO",
		P6Desc: "Bandelette chirurgicale bio-adhésive à micro-tension assurant une répartition uniforme des contraintes et limitant les cicatrices.",
		P6Domain: "Dispositif Biomécanique · Technologie Chirurgicale",
		P7Title: "WELSTECH — Wearable Life Saver Technology",
		P7Desc: "Moniteur de signes vitaux au poignet ultra-compact avec biocapteur optique PPG, processeur ESP32-S3, télémétrie Wi-Fi locale et alarme sonore progressive.",
		P7Domain: "IoT Biomédical Portable · Biocapteur Cardiaque",
		P8Title: "SENDIT",
		P8Desc: "Mécanisme autonome de tri et de transit logistique des bagages et du fret postal dans les aéroports.",
		P8Domain: "Robotique Autonome · Automatisation Logistique",
		P9Title: "SMARTWASTE",
		P9Desc: "Système robotisé mécatronique de tri sélectif combinant spectrométrie embarquée et vision par ordinateur pour le recyclage municipal.",
		P9Domain: "Vision par Ordinateur · Actionneurs Mécatroniques",
		P10Title: "IMAGE CLASSIFIER — Réseau de Neurones Convolutif Profond",
		P10Desc: "Réseau de neurones convolutif (CNN) entraîné sous TensorFlow/Keras sur plus de 9 600 images, intégrant un prétraitement 96×96 et une interface graphique d'inférence en temps réel.",
		P10Domain: "Deep Learning · Vision par Ordinateur & IA Embarquée",
		P11Title: "LUMIERE — Système Intelligent de Signalement des Pannes d'Éclairage Public",
		P11Desc: "Dispositif IoT compact installé sous les lampadaires municipaux, associant capteur photométrique LDR, géolocalisation GY-NEO6MV2 GPS, télémétrie ESP32 Wi-Fi SoftAP et alarme sonore pour automatiser les interventions.",
		P11Domain: "IoT Villes Intelligentes · Firmware Embarqué & Télémétrie",
		P12Title: "NOTCHA — Encoche Temporelle Dynamique et Flottante Minimaliste",
		P12Desc: "Conçu en 2021/2022 lorsque Windows 11 a supprimé la possibilité de placer la barre des tâches en haut de l'écran, NOTCHA est un utilitaire de bureau ultra-léger intégrant une encoche horaire discrète au sommet de l'écran, avec mise à jour en temps réel et affichage fluide au survol.",
		P12Domain: "Ergonomie de Bureau · Utilitaires d'Interface Personnalisés",
		P13Title: "SPO2 — Conception de Circuit Imprimé 2 Couches pour Oxymètre de Pouls",
		P13Desc: "Circuit imprimé 2 couches conçu avec KiCad 10.0, intégrant des empreintes et symboles personnalisés créés de zéro, les tolérances réelles de fabrication, le biocapteur optique MAX30102, l'écran OLED SSD1306 et un microcontrôleur Arduino Nano.",
		P13Domain: "Conception de PCB & EDA · Matériel Biomédical · KiCad 10.0",
		P14Title: "RESPIRO",
		P14Desc: "Ventilateur d'assistance biomédical adaptatif à faible coût et haute fiabilité conçu pour les déploiements d'urgence et les dispensaires à ressources limitées.",
		P14Domain: "Mécatronique Biomédicale · Soins d'Urgence",
		AboutTitle: "À Propos d'Ukacha Juma",
		AboutSubtitle: "Ingénieur Mécatronique · Architecte Matériel-Logiciel · Polyglotte",
		AboutP1: "Je suis un Ingénieur Mécatronique animé par la certitude que l'ingénierie matérielle doit prévenir la souffrance humaine plutôt que simplement la soigner. Mon travail réunit la CAO/FAO mécanique, les circuits capteurs analogiques de haute précision, le microprogramme sur microcontrôleurs bare-metal et FreeRTOS, et les systèmes de télémétrie en temps réel.",
		AboutP2: "De la prévention proactive des escarres (VELUM) à la filtration des microplastiques (FLOWGUARD) et aux biocapteurs thermiques (SOLUS), je conçois des systèmes physiques complets, des équations initiales jusqu'à la mise en production.",
		AboutPolyglot: "Communication Multilingue et Synergie Mondiale",
		AboutPolyglotDesc: "Je parle couramment anglais, turc, arabe, portugais, espagnol, allemand et français. Cette polyvalence linguistique me permet de diriger des équipes internationales et d'étudier la littérature clinique dans ses langues originales.",
		SkillsTitle: "Compétences Techniques & Outils",
		ExpTitle: "Expérience Professionnelle", ExpSubtitle: "Ingénierie industrielle pratique, maintenance et diagnostic CNC, automatismes et interprétation internationale.", EduTitle: "Formation & Parcours", AwardsTag: "RECONNAISSANCE ACADÉMIQUE ET ROBOTIQUE", AwardsTitle: "Prix et Distinctions", AwardsSubtitle: "Distinctions académiques institutionnelles, classements aux olympiades nationales de robotique et excellence en ingénierie.", AwardsCardTitle: "Classement İTÜRO 2026 — Plaque d'Honneur Décernée par la Doyenne", AwardsCardDesc: "Honoré par une plaque commémorative remise par la doyenne Prof. Dr. Humar Kahramanli Örnek lors de la cérémonie de remise des diplômes de la Faculté de Technologie de l'Université de Selçuk (2025–26), célébrant le classement de l'équipe STWolf aux Olympiades de Robotique de l'Université Technique d'Istanbul (İTÜRO 2026).", AwardsViewPost: "Voir la Publication et Reconnaissance sur LinkedIn", AwardsDeanTitle: "Doyenne de la Faculté de Technologie", AwardsDeanName: "Prof. Dr. Humar Kahramanlı Örnek", AwardsTeamName: "Équipe STWolf · Université de Selçuk", AwardsOlympiad: "Olympiades de Robotique de l'Université Technique d'Istanbul (İTÜRO 2026)", AwardsCeremony: "Cérémonie de Remise des Diplômes 2025–26 · Université de Selçuk", AwardsQuoteLabel: "Inscription Officielle de la Plaque", TabEducation: "Formation Académique", TabAwards: "Prix & Distinctions", TabCertificates: "Certificats & Licences", TabActivities: "Workshops", TabLanguages: "Langues", ActivitiesTitle: "Activités Techniques & Extracurriculaires", ActivitiesSubtitle: "Expositions en festivals technologiques, engagement communautaire et méthodologies de recherche avec l'IA.", TeknoselTitle: "Festival Technologique Selçuklu Teknosel — Catégorie Universitaire", TeknoselDesc: "Présentation du prototype mécatronique KonyaWaste sur notre stand dans la catégorie Projet Libre - Niveau Universitaire, avec démonstration en personne devant le maire de Selçuklu.", TeknoselMayorBadge: "Présentation Personnelle au Maire de Selçuklu", AiResearchTitle: "IA Fondée sur les Sources dans la Recherche Académique (Atelier TADD)", AiResearchDesc: "Atelier intensif de 2 jours organisé par TADD et dirigé par Yunus Emre Kıymaz, maîtrisant Zotero, ResearchRabbit et NotebookLM pour éliminer les hallucinations d'IA et garantir la rigueur scientifique des revues de littérature en mécatronique.", AiResearchBadge: "Recherche Rigoureuse sans Hallucinations d'IA", IfmWorkshopTitle: "Développement des Compétences en Automatismes avec ifm electronic & CoDeSys (Intensif 48h)", IfmWorkshopDesc: "Formation intensive spécialisée de 48 heures sur automates programmables industriels (PLC) organisée à l'Université de Selçuk par le siège d'ifm electronic Turquie. Expérience pratique sur l'architecture PLC, CoDeSys V3.5 (SP18 & SP19), protocoles Modbus et CANbus et dépannage sur site.", IfmWorkshopBadge: "⚙️ ifm electronic Turquie · Université de Selçuk", IfmBtnLinkedIn: "Voir la Publication de l'Atelier sur LinkedIn", TabSoloProjects: "Projets Solo", TabTeamProjects: "Projets en Équipe", TabContractProjects: "Projets sous Contrat", ContractEmptyBadge: "Disponible pour Contrats", ContractEmptyTitle: "Projets sous Contrat & Conseil", ContractEmptyDesc: "Contrats commerciaux confidentiels, conseil en ing\u00e9nierie mat\u00e9rielle et m\u00e9catronique industrielle. \u00c9tudes de cas prochainement.",
		CP1Title: "Mec\u00e2nica Auto Moderna \u2014 Design de Logo & Identit\u00e9 Visuelle", CP1Desc: "Cr\u00e9ation de logo d'entreprise et charte graphique sur mesure pour l'atelier automobile Mec\u00e2nica Auto Moderna. Embl\u00e8me combinant outils m\u00e9caniques de pr\u00e9cision, symbole de service 24h/24 et typographie moderne pour enseignes, uniformes et supports num\u00e9riques.", CP1Domain: "Identit\u00e9 Visuelle d'Entreprise & Design Vectoriel",
		CP2Title: "Mec\u00e2nica Auto Moderna \u2014 Syst\u00e8me de Facturation Invoice Moderna", CP2Desc: "Solution logicielle personnalis\u00e9e d\u00e9velopp\u00e9e pour Mec\u00e2nica Auto Moderna. Automatise l'\u00e9mission de factures \u00e9lectroniques proforma, le calcul des taxes, la gestion des clients et des ordres de r\u00e9paration avec une pr\u00e9cision et une fluidit\u00e9 accrues.", CP2Domain: "Logiciel de Facturation \u00c9lectronique & Gestion d'Entreprise",
		BtnInstagramPost: "Voir la publication Instagram", BtnLinkedInPost: "Voir la publication LinkedIn",
		CP3Title: "APACOV \u2014 Infrastructure Informatique & Architecture Web", CP3Desc: "D\u00e9veloppement b\u00e9n\u00e9vole du portail web institutionnel et gestion de l'infrastructure informatique pour l'ONG humanitaire APACOV, d\u00e9di\u00e9e \u00e0 la protection des enfants orphelins et vuln\u00e9rables. Responsable du d\u00e9ploiement serveur, de la s\u00e9curit\u00e9 num\u00e9rique et des affaires \u00e9trang\u00e8res.", CP3Domain: "Infrastructure IT \u00b7 Architecture Web \u00b7 Affaires \u00c9trang\u00e8res",
		CP4Title: "APACOV \u2014 Logo & Identit\u00e9 Visuelle Humanitaire", CP4Desc: "Identit\u00e9 visuelle b\u00e9n\u00e9vole pour l'ONG humanitaire APACOV. Con\u00e7ue sur Adobe Illustrator et finalis\u00e9e sur Samsung PenUp, le symbole illustre des mains tendues en signe de solidarit\u00e9 au c\u0153ur d'un cercle protecteur d'espoir.", CP4Domain: "Identit\u00e9 Visuelle Humanitaire \u00b7 Design Vectoriel",
		CP5Title: "SuPublic — Ingénierie Tôlerie pour Kiosque de Distribution d'Eau Potable", CP5Desc: "Prestation d'ingénierie mécanique sous contrat et conception tôlerie SolidWorks pour le projet éco-responsable SuPublic, fondé par Gonca Özata sous la direction du Dr Yavuz Selim Taşpınar. Initié à Konya pour cartographier les fontaines d'eau potable gratuites en Turquie et dans le monde afin d'éliminer les bouteilles en plastique à usage unique. Conception intégrale du châssis en tôle pliée, des plans de découpe laser et de la niche ergonomique de remplissage pour le prototype du kiosque automatique.", CP5Domain: "Ingénierie sous Contrat · Conception Tôlerie · SolidWorks CAD & Prototypage",
		CPNoticeInfra: "Note de Transition Serveur : Migration active de l'infrastructure et transfert d'apacov.org vers nos propres serveurs en cours, avec acquisition de nouveaux domaines et messageries institutionnelles.", BtnOfficialWebsite: "Site Web Officiel (apacov.org)", BadgeVolunteered: "B\u00e9n\u00e9volat \u00b7 Pro Bono",
		ExcavatorTitle: "Conception d'une Pelle Hydraulique (Éléments de Machines) — Inspirée du Komatsu PC200",
		ExcavatorDesc: "Conception mécanique approfondie et analyse structurelle par éléments finis (FEA ANSYS) d'une pelle hydraulique sur chenilles de 24 tonnes inspirée du Komatsu PC200 LC-7/LC-8. Calculs analytiques et vérifications sur 5 sous-systèmes critiques : axes et tourillons de flèche/balancier/godet, couronne d'orientation et galets de chenille, assemblages boulonnés précontraints, réducteurs planétaires et vérins hydrauliques haute pression.",
		ExcavatorDomain: "Conception d'Engins Lourds · Éléments de Machines · ANSYS FEA & Hydraulique",
		SafrenTitle: "SafRen (anciennement Ce-Saf) — Podomètre Intelligent pour Chaussures & IoT Portable",
		SafrenDesc: "Système connecté portable de comptage de pas et de télémétrie de marche développé en partenariat interdisciplinaire avec les étudiantes en Design de Chaussures Safina Najmudinova et Ceren Başaran, qui ont fourni les chaussures et défini l'emplacement ergonomique optimal des composants électroniques (écran OLED LCD, boutons et microcontrôleur). Intègre un microcontrôleur Waveshare ESP32-S3-Zero, une centrale inertielle 6 axes MPU6050, un serveur Wi-Fi SoftAP embarqué et une synchronisation en direct avec une application mobile Flutter.",
		SafrenDomain: "IoT Portable · Analyse de la Marche · Firmware ESP32-S3 · Application Flutter",
		STM32Title: "Défi Carte STM32 Discovery — Robotique Mobile Autonome",
		STM32Desc: "Ingénierie robotique autonome à double défi animée par la carte STM32F407VGT6 ARM Cortex-M4. Défi 1: Réseau de capteurs LDR avec conversion ADC 12 bits et commande différentielle PWM TIM3 pour cibler des sources lumineuses aléatoires dans une arène d'1 m² et s'arrêter au millimètre sans collision dès le premier essai. Défi 2: Robot extincteur de bougies avec ventilateur 12V à haut débit, centrage stratégique et balayage rotatif 360°—éteignant les 5 bougies en moins de 10 secondes (1/3 du temps alloué).",
		STM32Domain: "Systèmes Embarqués · STM32 ARM Cortex-M4 · C Bare-Metal / HAL · Robotique Autonome",
		ROSTitle: "Robot Mobile Autonome ROS Noetic — Pont TCP/IP ESP32 et Cinématique Différentielle",
		ROSDesc: "Robot mobile autonome développé en partenariat avec l'ingénieur mécatronique Mehmet Çakaroğlu sous la direction du Dr Okan Uyar. Relie ROS Noetic sous Linux au microcontrôleur double cœur ESP32 via un pont socket Wi-Fi TCP/IP à faible latence. Comprend une cinématique différentielle, une odométrie en boucle fermée avec encodeurs 420 impulsions, une centrale IMU 6 axes MPU6050 calibrée, un arrêt d'urgence deadman de 500 ms et une simulation jumeau numérique Gazebo/RViz validant les défis avec 100/100 points.",
		ROSDomain: "Robotique Mobile Autonome · ROS Noetic · Cinématique Différentielle · Gazebo & RViz",
		BtnLinkedInVideo: "Vidéo de Présentation LinkedIn", BtnGitHubRepo: "Dépôt GitHub", CertsTitle: "Certifications & Distinctions",
		VHeroBadge: "Candidature au James Dyson Award 2026 · Innovation en Santé",
		VHeroTitle: "VELUM — Tapis de Prévention des Escarres",
		VHeroSubtitle: "Le premier tapis capteur biomédical passif et continu au monde détectant les escarres 2 à 5 jours avant leur visibilité — équitablement, à bas coût et sans déranger le patient.",
		VStatCost: "+11 Milliards $ USD", VStatCostDesc: "Coût annuel du traitement des escarres évitables dans les hôpitaux américains",
		VStatEarly: "2 à 5 Jours Plus Tôt", VStatEarlyDesc: "Alerte pré-visuelle via la mesure de l'humidité sous-épidermique",
		VStatBOM: "< 50 $ USD", VStatBOMDesc: "Coût cible unitaire permettant d'équiper chaque lit d'hôpital",
		VStatEquity: "100% Équitable", VStatEquityDesc: "Mesure électrique indépendante de la pigmentation cutanée (Fitzpatrick I–VI)",
		VProblemTitle: "Un Protocole Hospitalier Défaillant",
		VProblemSubtitle: "L'inspection visuelle ne détecte les plaies qu'après l'apparition d'une nécrose tissulaire irréversible.",
		VProblemP1: "Les escarres touchent plus de 12,8 % des patients hospitalisés dans le monde. Le protocole actuel repose sur des retournements manuels toutes les 2 heures et des inspections visuelles subjectives. Pourtant, la nécrose cellulaire commence dans le tissu sous-épidermique 48 à 120 heures avant la rougeur superficielle.",
		VProblemP2: "De plus, l'examen visuel échoue sur les peaux foncées (Fitzpatrick IV–VI). Les scanners manuels commerciaux coûtent entre 3 000 $ et 10 000 $ et exigent un temps infirmier considérable.",
		VProblemPoint1: "Réactif par Nature: L'érythème cutané prouve que les cellules profondes sont déjà mortes.",
		VProblemPoint2: "Disparité Raciale Documentée: Le contrôle visuel ignore les escarres de stade 1 sur peaux mates.",
		VProblemPoint3: "Coûts Élevés & Fatigue du Personnel: Les sondes manuelles sont chères et alourdissent la charge soignante.",
		VSolutionTitle: "La Solution VELUM: Bioimpédance Continue",
		VSolutionSubtitle: "Surveillance tissulaire passive 24/7 installée discrètement sous le drap d'hôpital.",
		VSolutionP1: "VELUM est un tapis ultra-fin et souple en polyimide placé sous le patient. Une grille d'électrodes sérigraphiées argent-carbone réalise une spectroscopie de bioimpédance électrique (1–100 kHz) sur 16 zones anatomiques critiques toutes les 5 minutes.",
		VSolutionPoint1: "Mise en Place Immédiate: Glissez sous le drap, allumez, zéro perturbation pour le patient.",
		VSolutionPoint2: "Détection de l'Humidité Sous-Épidermique: Une chute de 15%+ d'impédance signale l'œdème plusieurs jours avant la plaie.",
		VSolutionPoint3: "Alertes Instantanées BLE / Wi-Fi: Les infirmiers reçoivent des alertes précises désignant la zone anatomique à soulager.",
		VSimTitle: "Simulateur Interactif de Bioimpédance",
		VSimSubtitle: "Cartographie anatomique en 16 zones et simulation du balayage de fréquence de l'AD5933.",
		VSimNormal: "1. État Normal", VSimIschemia: "2. Pression Sacrée Soutenue (Ischémie)",
		VSimOedema: "3. Alerte d'Œdème Stade 0", VSimReset: "4. Repositionnement du Patient",
		VSimStatusAlert: "CRITIQUE: Œdème sous-épidermique détecté! Repositionnez immédiatement la région sacrée",
		VSimStatusNormal: "NOMINAL: Perfusion tissulaire optimale sur l'ensemble des zones anatomiques",
		VTechTitle: "Architecture Mécatronique en Vue Éclatée",
		VTechSubtitle: "Électronique flexible de précision conçue pour la désinfection clinique et la production en série.",
		VLayer1Title: "1. Encapsulation en TPU Médical et Silicone Platine",
		VLayer1Desc: "Membrane biocompatible (ISO 10993) perméable à la vapeur et étanche aux liquides, résistante aux désinfectants hospitaliers.",
		VLayer2Title: "2. Grille d'Électrodes Sérigraphiées Argent/Carbone",
		VLayer2Desc: "Pistes conductrices sur film polyimide DuPont Kapton® de 50 µm disposées en réseaux tétrapolaires sous les saillies osseuses.",
		VLayer3Title: "3. Convertisseur d'Impédance AD5933 & MUX 16 Canaux",
		VLayer3Desc: "Composant Analog Devices AD5933 associé à une source Howland (<1 mA CA) et un multiplexeur CD74HC4067 pour un balayage rapide.",
		VLayer4Title: "4. Module de Calcul ESP32-S3 Dual-Core et Sans Fil",
		VLayer4Desc: "Processeur Xtensa 240MHz sous FreeRTOS avec filtrage d'étalonnage dynamique, BLE 5.0 mesh et télémétrie Wi-Fi sécurisée pour dossiers patients.",
		VCompTitle: "Avantage Clinique & Concurrentiel",
		VCompSubtitle: "Pourquoi VELUM offre une valeur clinique et économique sans équivalent.",
		VCompColMetric: "Critère / Fonction", VCompColVisual: "Inspection Visuelle", VCompColSEM: "Sonde SEM Manuelle", VCompColMattress: "Matelas de Décharge", VCompColVelum: "Tapis VELUM",
		VRoadmapTitle: "Feuille de Route Clinique et Commerciale",
		VRoadmapSubtitle: "Du banc d'essai aux essais cliniques en réanimation et à la soumission FDA 510(k).",
		VGalleryTitle: "Rendus 3D et Schémas Mécatroniques",
		VGallerySubtitle: "Illustrations de haute précision et vues éclatées du dispositif matériel.",
		VDownloadPitch: "Télécharger le Pitch Deck (PDF)",
		VViewGitHub: "Voir le Répertoire GitHub",
		FooterRights: "Tous droits réservés.",
		FooterBuiltWith: "",
	},
}

// GetLanguage returns the translation data for a language code, falling back to English
func GetLanguage(code string) TranslationData {
	code = strings.ToLower(strings.TrimSpace(code))
	if data, ok := translations[code]; ok {
		data.Languages = SupportedLanguages
		return data
	}
	data := translations["en"]
	data.Languages = SupportedLanguages
	return data
}

// DetectLanguage determines the best language from URL query, cookie, or Accept-Language header
func DetectLanguage(r *http.Request) string {
	// 1. URL Query ?lang=...
	if lang := r.URL.Query().Get("lang"); lang != "" {
		lang = strings.ToLower(lang)
		if _, ok := translations[lang]; ok {
			return lang
		}
	}

	// 2. Cookie
	if cookie, err := r.Cookie("portfolio_lang"); err == nil {
		lang := strings.ToLower(cookie.Value)
		if _, ok := translations[lang]; ok {
			return lang
		}
	}

	// 3. Accept-Language header
	accept := r.Header.Get("Accept-Language")
	if accept != "" {
		for _, part := range strings.Split(accept, ",") {
			code := strings.TrimSpace(strings.Split(part, ";")[0])
			if len(code) >= 2 {
				prefix := strings.ToLower(code[:2])
				if _, ok := translations[prefix]; ok {
					return prefix
				}
			}
		}
	}

	return "en"
}
