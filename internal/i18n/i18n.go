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
	NavGitHub      string
	NavVelum       string

	// Home Hero
	HeroBadge      string
	HeroTitle      string
	HeroSubtitle   string
	HeroBtnExplore string
	HeroBtnContact string
	StatProjects   string
	StatAward      string
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
		NavHome: "Home", NavProjects: "Projects", NavExperience: "Experience", NavAbout: "About", NavSkills: "Skills", NavContact: "Contact", NavGitHub: "GitHub", NavVelum: "VELUM Project",
		HeroBadge: "Mechatronics Engineer · Biomedical Hardware · Embedded Systems · Polyglot",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Designing intelligent biomedical mechatronics, precision sensor instrumentation, and edge IoT devices that protect vulnerable human lives.",
		HeroBtnExplore: "Explore Projects", HeroBtnContact: "Get in Touch",
		StatProjects: "10 Engineered Projects", StatAward: "Mechatronics Engineering", StatLanguages: "7 Fluent Languages", StatFocus: "ESP32 / Bioimpedance / CAD",
		ProjectsTitle: "Engineering Portfolio",
		ProjectsSubtitle: "Proactive biomedical hardware, assistive devices, fluidic automation, and mechatronic systems.",
		StatusActive: "Complete & Interactive", StatusQueue: "In Queue",
		BtnViewProject: "Explore VELUM", BtnPreview: "Project Details",
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
		P4Title: "DENOVA",
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
		P10Title: "RESPIRO",
		P10Desc: "Low-cost, high-reliability biomedical adaptive assistive ventilator engineered for emergency deployment and resource-constrained clinics.",
		P10Domain: "Biomedical Mechatronics · Critical Care",
		AboutTitle: "About Ukacha Juma",
		AboutSubtitle: "Mechatronics Engineer · Hardware-Software Architect · Polyglot",
		AboutP1: "I am a Mechatronics Engineer driven by the conviction that modern hardware engineering should prevent suffering rather than merely treat it. My work bridges mechanical CAD/CAM design, precision analog sensor circuits, firmware on bare-metal and FreeRTOS microcontrollers, and real-time telemetry systems.",
		AboutP2: "With projects spanning proactive pressure-injury mats (VELUM), microplastic filtration (FLOWGUARD), and thermal biosensors (SOLUS), I build physical systems from early mathematical principles through benchtop circuit validation to production-ready design.",
		AboutPolyglot: "Multilingual Communication & Global Collaboration",
		AboutPolyglotDesc: "Fluent in English, Turkish, Arabic, Portuguese, Spanish, German, and French. Being a polyglot empowers me to bridge technical teams across continents, read global clinical literature in original sources, and design universally accessible technologies.",
		SkillsTitle: "Technical Proficiencies & Toolchain",
		ExpTitle: "Professional Experience", ExpSubtitle: "Hands-on industrial engineering, CNC diagnostics, automation maintenance, and international translation.", EduTitle: "Education & Academic Background", CertsTitle: "Certifications & Honors",
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
		FooterRights: "All rights reserved. Designed and developed by Ukacha Juma.",
		FooterBuiltWith: "Powered by Go (Golang) · Docker Support · Render Ready",
	},
	"tr": {
		LangCode: "tr", LangName: "Turkish", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Ana Sayfa", NavProjects: "Projeler", NavExperience: "Deneyim", NavAbout: "Hakkımda", NavSkills: "Yetenekler", NavContact: "İletişim", NavGitHub: "GitHub", NavVelum: "VELUM Projesi",
		HeroBadge: "Mekatronik Mühendisi · Biyomedikal Donanım · Gömülü Sistemler · Çok Dilli",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "İnsan hayatını koruyan akıllı biyomedikal mekatronik sistemler, hassas sensör enstrümantasyonu ve uç IoT cihazları tasarlıyorum.",
		HeroBtnExplore: "Projeleri Keşfet", HeroBtnContact: "İletişime Geç",
		StatProjects: "10 Mekatronik Proje", StatAward: "Mekatronik Mühendisliği", StatLanguages: "7 Akıcı Yabancı Dil", StatFocus: "ESP32 / Biyoempedans / CAD",
		ProjectsTitle: "Mühendislik Portföyü",
		ProjectsSubtitle: "Önleyici biyomedikal donanım, yardımcı cihazlar, akışkan otomasyonu ve mekatronik sistemler.",
		StatusActive: "Tamamlandı & Etkileşimli", StatusQueue: "Sırada",
		BtnViewProject: "VELUM'u İncele", BtnPreview: "Proje Detayları",
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
		P4Title: "DENOVA",
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
		P10Title: "RESPIRO",
		P10Desc: "Acil durum dağıtımı ve kaynak kısıtlı klinikler için tasarlanmış düşük maliyetli, yüksek güvenilirlikli biyomedikal adaptif ventilatör.",
		P10Domain: "Biyomedikal Mekatronik · Yoğun Bakım",
		AboutTitle: "Ukacha Juma Hakkında",
		AboutSubtitle: "Mekatronik Mühendisi · Donanım-Yazılım Mimarı · Poliglot",
		AboutP1: "Modern donanım mühendisliğinin acıyı sadece tedavi etmekle kalmayıp önlemesi gerektiği inancıyla çalışan bir Mekatronik Mühendisiyim. Çalışmalarım mekanik CAD/CAM tasarımını, hassas analog sensör devrelerini, bare-metal ve FreeRTOS mikrodenetleyici ürün yazılımlarını ve gerçek zamanlı telemetri sistemlerini birleştirmektedir.",
		AboutP2: "Yatak yarası önleyici matlardan (VELUM), mikroplastik filtrelemeye (FLOWGUARD) ve termal biyosensörlere (SOLUS) kadar, fiziksel sistemleri matematiksel prensiplerden laboratuvar doğrulamasına ve üretime hazır tasarıma kadar sıfırdan inşa ediyorum.",
		AboutPolyglot: "Çok Dilli İletişim & Küresel İşbirliği",
		AboutPolyglotDesc: "İngilizce, Türkçe, Arapça, Portekizce, İspanyolca, Almanca ve Fransızca dillerini akıcı bir şekilde konuşuyorum. Poliglot olmak, kıtalararası teknik ekipler arasında köprü kurmamı ve küresel klinik literatürü orijinal kaynağından takip etmemi sağlıyor.",
		SkillsTitle: "Teknik Yetkinlikler & Araç Takımı",
		ExpTitle: "Mesleki Deneyim", ExpSubtitle: "Uygulamalı endüstriyel mühendislik, CNC arıza tespiti, otomasyon bakımı ve uluslararası tercümanlık.", EduTitle: "Eğitim ve Akademik Geçmiş", CertsTitle: "Sertifikalar ve Başarılar",
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
		FooterRights: "Tüm hakları saklıdır. Ukacha Juma tarafından tasarlanmış ve geliştirilmiştir.",
		FooterBuiltWith: "Go (Golang) Altyapısı · Docker Desteği · Render Dağıtımına Hazır",
	},
	"ar": {
		LangCode: "ar", LangName: "Arabic", IsRTL: true, Dir: "rtl",
		NavBrand: "أوكاشا جمعة",
		NavHome: "الرئيسية", NavProjects: "المشاريع", NavExperience: "الخبرات", NavAbout: "عني", NavSkills: "المهارات", NavContact: "اتصل بي", NavGitHub: "جيت هاب", NavVelum: "مشروع فيلوم",
		HeroBadge: "مهندس ميكاترونكس · أجهزة طبية حيوية · أنظمة مدمجة · متعدد اللغات",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "تصميم أنظمة ميكاترونكس طبية ذكية، وأجهزة استشعار حيوية دقيقة، وتقنيات إنترنت الأشياء الطرفية لحماية حياة المرضى.",
		HeroBtnExplore: "استكشاف المشاريع", HeroBtnContact: "تواصل معي",
		StatProjects: "١٠ مشاريع هندسية", StatAward: "هندسة الميكاترونكس", StatLanguages: "٧ لغات متقنة", StatFocus: "ESP32 / المعاوقة الحيوية / CAD",
		ProjectsTitle: "معرض المشاريع الهندسية",
		ProjectsSubtitle: "أجهزة طبية وقائية، أنظمة مساعدة، أتمتة موائع، وتطبيقات ميكاترونكس متقدمة.",
		StatusActive: "مكتمل وتفاعلي", StatusQueue: "قيد التطوير",
		BtnViewProject: "استكشاف فيلوم (VELUM)", BtnPreview: "تفاصيل المشروع",
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
		P4Title: "دينوفا - DENOVA",
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
		P10Title: "ريسبيرو - RESPIRO",
		P10Desc: "جهاز تنفس صناعي تكيفي طبي منخفض التكلفة وعالي الموثوقية مخصص لغرف الطوارئ والمستشفيات الميدانية.",
		P10Domain: "ميكاترونكس طبية حيوية · العناية المركزة",
		AboutTitle: "نبذة عن أوكاشا جمعة",
		AboutSubtitle: "مهندس ميكاترونكس · مهندس نظم العتاد والبرمجيات · متعدد اللغات",
		AboutP1: "أنا مهندس ميكاترونكس أؤمن بأن الهندسة الحديثة يجب أن تكرس لمنع المعاناة الإنسانية واستباقها بدلاً من الاكتفاء بعلاجها. تجمع خبرتي بين التصميم الميكانيكي CAD/CAM، وتصميم الدوائر التناظرية الدقيقة، وبرمجة الأنظمة المدمجة مع FreeRTOS، وأنظمة الاتصال اللاسلكي الفوري.",
		AboutP2: "من خلال مشاريع مبتكرة مثل بساط الوقاية من قرح الفراش (VELUM)، وترشيح البلاستيك الدقيق (FLOWGUARD)، والحساسات الفسيولوجية (SOLUS)، أقوم ببناء النظم الفيزيائية المتكاملة من النماذج الحسابية إلى النماذج المختبرية والتصنيع الفعلي.",
		AboutPolyglot: "تواصل عالمي وطلاقة متعددة اللغات",
		AboutPolyglotDesc: "أتحدث بطلاقة ٧ لغات: الإنجليزية، التركية، العربية، البرتغالية، الإسبانية، الألمانية، والفرنسية. هذا التنوع يمنحني القدرة على قيادة الفرق الهندسية الدولية، وقراءة الأبحاث الطبية بلغاتها الأصلية، وتصميم تقنيات تصل للعالم بأسره.",
		SkillsTitle: "المهارات الفنية والأدوات الهندسية",
		ExpTitle: "الخبرة المهنية والعملية", ExpSubtitle: "هندسة صناعية تطبيقية، صيانة أنظمة CNC، التحكم الآلي، والترجمة الدولية.", EduTitle: "التعليم والخلفية الأكاديمية", CertsTitle: "الشهادات والاعتمادات",
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
		FooterRights: "جميع الحقوق محفوظة. صُمم وطُوّر بواسطة أوكاشا جمعة.",
		FooterBuiltWith: "خادم بلغة Go (Golang) · دعم Docker · جاهز للنشر على Render",
	},
	"pt": {
		LangCode: "pt", LangName: "Portuguese", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Início", NavProjects: "Projetos", NavExperience: "Experiência", NavAbout: "Sobre", NavSkills: "Habilidades", NavContact: "Contato", NavGitHub: "GitHub", NavVelum: "Projeto VELUM",
		HeroBadge: "Engenheiro Mecatrônico · Hardware Biomédico · Sistemas Embarcados · Poliglota",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Projetando sistemas mecatrônicos biomédicos inteligentes, instrumentação sensorial de precisão e dispositivos IoT para proteger vidas humanas.",
		HeroBtnExplore: "Explorar Projetos", HeroBtnContact: "Entrar em Contato",
		StatProjects: "10 Projetos Mecatrônicos", StatAward: "Engenharia Mecatrônica", StatLanguages: "7 Idiomas Fluentes", StatFocus: "ESP32 / Bioimpedância / CAD",
		ProjectsTitle: "Portfólio de Engenharia",
		ProjectsSubtitle: "Hardware biomédico preventivo, dispositivos assistivos, automação fluídica e sistemas mecatrônicos.",
		StatusActive: "Concluído e Interativo", StatusQueue: "Na Fila",
		BtnViewProject: "Explorar VELUM", BtnPreview: "Detalhes do Projeto",
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
		P4Title: "DENOVA",
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
		P10Title: "RESPIRO",
		P10Desc: "Ventilador assistivo adaptativo biomédico de baixo custo e alta confiabilidade projetado para emergências e clínicas com recursos limitados.",
		P10Domain: "Mecatrônica Biomédica · Cuidados Intensivos",
		AboutTitle: "Sobre Ukacha Juma",
		AboutSubtitle: "Engenheiro Mecatrônico · Arquiteto Hardware-Software · Poliglota",
		AboutP1: "Sou um Engenheiro Mecatrônico impulsionado pela convicção de que a engenharia de hardware moderna deve prevenir o sofrimento humano em vez de apenas remediá-lo. Meu trabalho conecta design mecânico em CAD/CAM, circuitos sensores analógicos de precisão, firmware em microcontroladores bare-metal e FreeRTOS, e telemetria em tempo real.",
		AboutP2: "Com projetos que abrangem tapetes de prevenção de escaras (VELUM), filtragem de microplásticos (FLOWGUARD) e biossensores térmicos (SOLUS), construo sistemas físicos desde os princípios matemáticos fundamentais até a validação em bancada e projetos prontos para produção.",
		AboutPolyglot: "Comunicação Multilíngue e Colaboração Global",
		AboutPolyglotDesc: "Fluente em inglês, turco, árabe, português, espanhol, alemão e francês. Ser poliglota me permite conectar equipes de engenharia ao redor do mundo, consultar literatura clínica nas fontes originais e projetar tecnologias universalmente acessíveis.",
		SkillsTitle: "Competências Técnicas e Ferramentas",
		ExpTitle: "Experiência Profissional", ExpSubtitle: "Engenharia industrial prática, manutenção e diagnóstico de CNC, automação e tradução internacional.", EduTitle: "Educação e Formação", CertsTitle: "Certificações e Distinções",
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
		FooterRights: "Todos os direitos reservados. Projetado e desenvolvido por Ukacha Juma.",
		FooterBuiltWith: "Desenvolvido em Go (Golang) · Suporte a Docker · Pronto para o Render",
	},
	"es": {
		LangCode: "es", LangName: "Spanish", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Inicio", NavProjects: "Proyectos", NavExperience: "Experiencia", NavAbout: "Sobre mí", NavSkills: "Habilidades", NavContact: "Contacto", NavGitHub: "GitHub", NavVelum: "Proyecto VELUM",
		HeroBadge: "Ingeniero Mecatrónico · Hardware Biomédico · Sistemas Embebidos · Políglota",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Diseño de sistemas mecatrónicos biomédicos inteligentes, instrumentación sensorial de precisión y dispositivos IoT para proteger vidas humanas.",
		HeroBtnExplore: "Explorar Proyectos", HeroBtnContact: "Contactar",
		StatProjects: "10 Proyectos Mecatrónicos", StatAward: "Ingeniería Mecatrónica", StatLanguages: "7 Idiomas Fluidos", StatFocus: "ESP32 / Bioimpedancia / CAD",
		ProjectsTitle: "Portafolio de Ingeniería",
		ProjectsSubtitle: "Hardware biomédico preventivo, dispositivos de asistencia, automatización fluídica y sistemas mecatrónicos.",
		StatusActive: "Completado e Interactivo", StatusQueue: "En Cola",
		BtnViewProject: "Explorar VELUM", BtnPreview: "Detalles del Proyecto",
		BadgeFeatured: "Proyecto Destacado", BadgeBiomedical: "Hardware Biomédico", BadgeHardware: "Mecatrónica", BadgeRobotics: "Robótica y Automatización",
		P1Title: "VELUM (anteriormente SentIO)",
		P1Desc: "Esterilla sensora de bioimpedancia eléctrica y humedad subepidérmica continua 24/7 que detecta úlceras por presión de 2 a 5 días antes de que aparezcan heridas.",
		P1Domain: "Mecatrónica Biomédica · Espectroscopia de Bioimpedancia",
		P2Title: "FLOWGUARD",
		P2Desc: "Sistema mecatrónico automatizado de filtración de microplásticos para lavadoras que evita la contaminación de los ecosistemas acuáticos.",
		P2Domain: "Mecatrónica Fluídica · Separación de Microplásticos",
		P3Title: "SOLUS",
		P3Desc: "Parche vestible de monitorización continua de temperatura central para la alerta temprana de golpes de calor en trabajadores expuestos.",
		P3Domain: "Biosensores Vestibles · Fisiología Preventiva",
		P4Title: "DENOVA",
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
		P10Title: "RESPIRO",
		P10Desc: "Ventilador asistencial adaptativo biomédico de bajo coste y alta fiabilidad diseñado para emergencias y hospitales de campaña.",
		P10Domain: "Mecatrónica Biomédica · Cuidados Críticos",
		AboutTitle: "Sobre Ukacha Juma",
		AboutSubtitle: "Ingeniero Mecatrónico · Arquitecto Hardware-Software · Políglota",
		AboutP1: "Soy un Ingeniero Mecatrónico motivado por la convicción de que la ingeniería moderna de hardware debe prevenir el sufrimiento humano en lugar de simplemente tratarlo. Mi trabajo integra diseño mecánico CAD/CAM, circuitos sensores analógicos de precisión, firmware en microcontroladores bare-metal y FreeRTOS, y telemetría en tiempo real.",
		AboutP2: "Con proyectos que van desde esterillas de prevención de úlceras (VELUM), filtración de microplásticos (FLOWGUARD) y biosensores térmicos (SOLUS), construyo sistemas físicos desde principios matemáticos hasta pruebas de laboratorio y producción.",
		AboutPolyglot: "Comunicación Multilingüe y Colaboración Global",
		AboutPolyglotDesc: "Hablo con fluidez inglés, turco, árabe, portugués, español, alemán y francés. Esta capacidad me permite liderar equipos internacionales, revisar literatura médica en sus idiomas originales y concebir tecnologías de alcance universal.",
		SkillsTitle: "Competencias Técnicas y Herramientas",
		ExpTitle: "Experiencia Profesional", ExpSubtitle: "Ingeniería industrial aplicada, diagnóstico de maquinaria CNC, mantenimiento de automatización y traducción internacional.", EduTitle: "Educación y Formación", CertsTitle: "Certificaciones y Reconocimientos",
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
		FooterRights: "Todos los derechos reservados. Diseñado y desarrollado por Ukacha Juma.",
		FooterBuiltWith: "Desarrollado en Go (Golang) · Compatible con Docker · Listo para Render",
	},
	"de": {
		LangCode: "de", LangName: "German", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Startseite", NavProjects: "Projekte", NavExperience: "Erfahrung", NavAbout: "Über mich", NavSkills: "Kompetenzen", NavContact: "Kontakt", NavGitHub: "GitHub", NavVelum: "VELUM Projekt",
		HeroBadge: "Mechatronik-Ingenieur · Biomedizinische Hardware · Eingebettete Systeme · Polyglott",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Entwicklung intelligenter biomedizinischer Mechatronik, hochpräziser Sensoren und Edge-IoT-Systeme zum Schutz von Menschenleben.",
		HeroBtnExplore: "Projekte Erkunden", HeroBtnContact: "Kontakt Aufnehmen",
		StatProjects: "10 Mechatronik-Projekte", StatAward: "Mechatronik-Ingenieurwesen", StatLanguages: "7 Fließende Sprachen", StatFocus: "ESP32 / Bioimpedanz / CAD",
		ProjectsTitle: "Ingenieur-Portfolio",
		ProjectsSubtitle: "Präventive biomedizinische Hardware, assistive Medizingeräte, Fluidik und mechatronische Systeme.",
		StatusActive: "Abgeschlossen & Interaktiv", StatusQueue: "In Vorbereitung",
		BtnViewProject: "VELUM Entdecken", BtnPreview: "Projektdetails",
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
		P4Title: "DENOVA",
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
		P10Title: "RESPIRO",
		P10Desc: "Kostengünstiges, hochzuverlässiges adaptives Beatmungsgerät für Notfalleinsätze und ressourcenarme Kliniken.",
		P10Domain: "Biomedizinische Mechatronik · Intensivmedizin",
		AboutTitle: "Über Ukacha Juma",
		AboutSubtitle: "Mechatronik-Ingenieur · Hardware-Software-Architekt · Polyglott",
		AboutP1: "Als Mechatronik-Ingenieur treibt mich die Überzeugung an, dass moderne Hardware-Entwicklung menschliches Leid verhindern sollte, anstatt es nur im Nachhinein zu behandeln. Meine Arbeit vereint mechanisches CAD/CAM-Design, analoge Präzisionsschaltungen, Firmware auf Bare-Metal und FreeRTOS sowie Telemetriesysteme.",
		AboutP2: "Mit Projekten wie Dekubitus-Präventionsmatten (VELUM), Mikroplastikfiltration (FLOWGUARD) und thermischen Biosensoren (SOLUS) entwickle ich mechatronische Systeme von mathematischen Grundlagen bis zu serienreifen Prototypen.",
		AboutPolyglot: "Mehrsprachige Kommunikation & Globale Zusammenarbeit",
		AboutPolyglotDesc: "Ich spreche fließend Englisch, Türkisch, Arabisch, Portugiesisch, Spanisch, Deutsch und Französisch. Als Polyglott kann ich internationale Entwicklerteams leiten und klinische Studien in ihren Originalfassungen analysieren.",
		SkillsTitle: "Technische Kompetenzen & Toolchain",
		ExpTitle: "Berufserfahrung", ExpSubtitle: "Praktische Industrieerfahrung, CNC-Diagnostik, Automatisierungswartung und internationale Übersetzung.", EduTitle: "Ausbildung & Bildungsweg", CertsTitle: "Zertifikate & Auszeichnungen",
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
		FooterRights: "Alle Rechte vorbehalten. Entwickelt von Ukacha Juma.",
		FooterBuiltWith: "Powered by Go (Golang) · Docker-Unterstützung · Bereit für Render",
	},
	"fr": {
		LangCode: "fr", LangName: "French", IsRTL: false, Dir: "ltr",
		NavBrand: "Ukacha JUMA",
		NavHome: "Accueil", NavProjects: "Projets", NavExperience: "Expérience", NavAbout: "À propos", NavSkills: "Compétences", NavContact: "Contact", NavGitHub: "GitHub", NavVelum: "Projet VELUM",
		HeroBadge: "Ingénieur Mécatronique · Matériel Biomédical · Systèmes Embarqués · Polyglotte",
		HeroTitle: "“I get things done!”",
		HeroSubtitle: "Conception de systèmes mécatroniques biomédicaux intelligents, d'instrumentation sensorielle de précision et d'appareils IoT pour sauver des vies.",
		HeroBtnExplore: "Explorer les Projets", HeroBtnContact: "Me Contacter",
		StatProjects: "10 Projets Mécatroniques", StatAward: "Génie Mécatronique", StatLanguages: "7 Langues Maîtrisées", StatFocus: "ESP32 / Bioimpédance / CAD",
		ProjectsTitle: "Portfolio d'Ingénierie",
		ProjectsSubtitle: "Matériel biomédical préventif, dispositifs d'assistance, automatisation fluidique et systèmes mécatroniques.",
		StatusActive: "Terminé et Interactif", StatusQueue: "En File d'Attente",
		BtnViewProject: "Découvrir VELUM", BtnPreview: "Détails du Projet",
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
		P4Title: "DENOVA",
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
		P10Title: "RESPIRO",
		P10Desc: "Ventilateur biomédical adaptatif d'assistance respiratoire à faible coût et haute fiabilité pour les soins d'urgence.",
		P10Domain: "Mécatronique Biomédicale · Soins Intensifs",
		AboutTitle: "À Propos d'Ukacha Juma",
		AboutSubtitle: "Ingénieur Mécatronique · Architecte Matériel-Logiciel · Polyglotte",
		AboutP1: "Je suis un Ingénieur Mécatronique animé par la certitude que l'ingénierie matérielle doit prévenir la souffrance humaine plutôt que simplement la soigner. Mon travail réunit la CAO/FAO mécanique, les circuits capteurs analogiques de haute précision, le microprogramme sur microcontrôleurs bare-metal et FreeRTOS, et les systèmes de télémétrie en temps réel.",
		AboutP2: "De la prévention proactive des escarres (VELUM) à la filtration des microplastiques (FLOWGUARD) et aux biocapteurs thermiques (SOLUS), je conçois des systèmes physiques complets, des équations initiales jusqu'à la mise en production.",
		AboutPolyglot: "Communication Multilingue et Synergie Mondiale",
		AboutPolyglotDesc: "Je parle couramment anglais, turc, arabe, portugais, espagnol, allemand et français. Cette polyvalence linguistique me permet de diriger des équipes internationales et d'étudier la littérature clinique dans ses langues originales.",
		SkillsTitle: "Compétences Techniques & Outils",
		ExpTitle: "Expérience Professionnelle", ExpSubtitle: "Ingénierie industrielle pratique, maintenance et diagnostic CNC, automatismes et interprétation internationale.", EduTitle: "Formation & Parcours", CertsTitle: "Certifications & Distinctions",
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
		FooterRights: "Tous droits réservés. Conçu et développé par Ukacha Juma.",
		FooterBuiltWith: "Propulsé par Go (Golang) · Support Docker · Prêt pour déploiement Render",
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
