**SUSTAINABILITY** — Smart Waste Management System with Automated Sorting & Compaction

---

## 🧪 Minimum Viable Prototype — What You Need to Build

> **Rule:** Dyson rejects "wholly conceptual" entries. You do NOT need a finished product — you need verifiable evidence of iterative development. Here is the minimum credible proof of work for SMARTWASTE.

### What counts as enough:

| # | Action | What it proves | Difficulty |
|---|--------|---------------|------------|
| 1 | **Wire up an ESP32 + ToF (Time-of-Flight) laser distance sensor on a breadboard** — point it into a test bin and read fill-level values over Serial monitor | The core sensing mechanism works; fill level is measurable in real time | Low — weekend project |
| 2 | **Add a material classification module** — use a small servo-actuated gate triggered by a conductivity test or IR sensor to separate a metallic can from a plastic bottle in a split-bin mock-up | Proves the sorting gate concept responds to material type | Medium — 2–3 days |
| 3 | **Build a basic compaction actuator** — a servo or small linear actuator pressing down on a foam block inside a box, triggered when fill reaches 70% | Demonstrates the compression mechanism in a scaled-down form | Medium |
| 4 | **Photograph and log the bench setup** — breadboard, sensors, actuator, serial output on screen showing fill level and material type | Visual proof of physical prototyping for judges | Trivial |
| 5 | **Plot fill-level vs. time** (Python + matplotlib) across a simulated waste insertion cycle, with a compaction event marked on the graph | Shows iterative data analysis and system behaviour, not just a one-time test | Low |
| 6 | **Print a scaled layout** of the sensor module, sorting gate, and compaction piston inside a bin cross-section — photograph it | Shows form-factor thinking; judges want to see the mechanism, not just a circuit | Low |
| 7 | **One paragraph of written test notes** in `prototyping/` — what classification method worked, what failed, compression ratios achieved | Demonstrates design iteration; honest notes are worth more than polished renders | |

### What you do NOT need:
- A working PCB (stripboard or breadboard is fine)
- A fully enclosed weatherproof housing (mention IP67 as the next step)
- A live cloud dashboard deployed (a local web server log is sufficient)
- A real municipal deployment

### Minimum time to credible prototype: **3–5 days**

---

# Project 09 — SMARTWASTE: Smart Waste Sorting & Compaction System

**James Dyson Award 2026**
**Category:** Sustainability / Resource Efficiency
**Scope:** Urban municipalities, industrial zones, university campuses
**Deadline:** 15 July 2026

---

## The Problem

Municipal waste collection vehicles follow fixed, static routes regardless of whether bins are full or empty — wasting fuel, time, and labour. At the same time, unsorted waste (plastics, cardboard, metals) is sent to landfill because mixed materials are uneconomical to process. Compressible waste (plastics, cartons, cardboard) occupies far more volume than its mass justifies, dramatically reducing the effective storage capacity of each bin before collection.

## The Solution

An integrated smart bin unit that: (1) continuously monitors fill levels via laser distance sensing, (2) automatically separates waste by material type at the point of deposit using a sensor-triggered sorting gate, and (3) compacts compressible materials (plastic bottles, cardboard, cartons) in-situ using a motorised press — maximising storage volume between collections. Fill-level and route data are transmitted over Wi-Fi to a central municipal dashboard for dynamic, demand-driven collection routing.

---

## James Dyson Award Application — Draft Answers

> 📝 **Character limits** based on the official Dyson Award application form.

---

### 📛 Project Name
*(Limit: 50 characters)*

SMARTWASTE — Smart Sorting & Compaction Bin

---

### ⚡ Quick Description
*(Limit: 193 characters)*

A smart bin that sorts waste by material type, compacts plastics and cardboard to triple storage, and transmits real-time fill data so fleets only visit bins that are genuinely full.

---

### ❓ What my invention does
*(Limit: 260 characters)*

SMARTWASTE sorts waste into material streams via sensor-triggered gates, compacts plastics, cartons and cardboard with an in-bin press to maximise storage, and reports live fill levels to enable data-driven collection routing.

---

### ❓ Where I got my inspiration
*(Limit: 650 characters)*

Waste collection accounts for 70–80% of solid waste management costs globally, yet vehicles follow static routes and bins accept mixed, uncompressed waste. Field observations in Konya revealed none of the city's containers had real-time monitoring, sorting, or compaction — trucks visited half-empty and overflowing bins alike. Compressible materials like plastic bottles and cardboard occupy enormous volume relative to their mass. Intelligent fill sensing, automatic sorting, and mechanical compaction was the logical solution to tackle route inefficiency, landfill contamination, and premature overflow.

---

### ❓ How it works
*(Limit: 975 characters)*

A waterproof ToF laser sensor inside the bin lid measures distance to the waste surface, giving a volumetric fill percentage. At the inlet gate, an IR reflectance and conductivity array classifies material: metals trigger one chute, plastics and cardboard another. A servo gate directs each item to the correct chamber. In the compressible-materials chamber, a motorised piston press activates at 70% fill, compressing plastics, bottles, and cartons to one-third of their volume — tripling effective storage. An ESP32 manages sensors and actuators, transmitting fill data via Wi-Fi only when the 80% threshold is crossed or a keep-alive signal is due, preserving battery via Deep Sleep. Data reaches a cloud dashboard where fleet software generates optimised collection routes.

---

### ❓ How you developed your design | Design Process
*(Limit: 1300 characters)*

Design began with field observation and review of Konya municipal waste data, confirming no smart monitoring, sorting, or compaction existed — collection ran on static timetables.

The sensing challenge was fill-level accuracy inside containers with irregular waste. Ultrasonic sensors were eliminated due to multipath echo errors on uneven waste surfaces and degraded performance in humid, odour-rich environments. Waterproof laser ToF sensors were selected for millimetre precision under the same conditions.

For material classification, IR reflectance combined with a conductivity bridge distinguishes metals from non-metals; a density step separates rigid plastics from soft films. The sorting gate is a servo-actuated dual-chute — mechanically simple and field-repairable.

The compaction mechanism targets the compressible-materials chamber only. A lead-screw piston driven by a geared DC motor presses plastics, cartons, and cardboard — materials that yield predictably under compression. A 3:1 compression target was validated against published packaging data. The 3D-printed IP67 housing consolidates all electronics and actuators into a retrofit lid module for standard 240L and 1100L containers.

---

### ❓ What sets it apart
*(Limit: 975 characters)*

Existing smart bin solutions address only fill-level monitoring — none integrate sorting and in-bin compaction into a single retrofit unit. SMARTWASTE's three-function combination is the differentiator: monitoring reduces unnecessary collections, sorting prevents recyclable stream contamination, and compaction delays the collection trigger — multiplying the efficiency gain.

Laser ToF over ultrasonic is a technical differentiator: ultrasonic units produce systematic echo errors inside waste containers, a documented limitation in IoT waste literature. ToF eliminates this entirely.

At under $80 per unit using 3D-printed housings and off-the-shelf electronics, the system is accessible to mid-sized municipalities — not just well-funded pilot programmes. The SaaS dashboard model means municipalities pay per active bin rather than bearing upfront infrastructure costs.

---

### ❓ What the future holds for it
*(Limit: 650 characters)*

The immediate next step is a two-week pilot deployment on real municipal containers in Konya's Selçuklu district, measuring collection-route distance reduction and compaction ratio accuracy. A machine learning classification layer will improve material sorting accuracy beyond binary metal/non-metal detection — adding paper, glass, and organic streams. The platform's IoT backbone is extensible: air quality sensors, flood drain monitoring, and smart streetlight control can be added as modules. A SaaS licensing model will allow municipalities and Organised Industrial Zones (OIZ) to subscribe to the dashboard without purchasing hardware outright.

---

## Folder Structure

```
project_09_SMARTWASTE/
├── README.md
├── images/
├── research/
├── sketches/
├── cad/
├── simulation/
├── scripts/
├── prototyping/
├── application/
└── presentation/
```

---

## TekoSel 2026 — Report Sections (Serbest Kategori)

> Mapped to the UNIVERSITE_TEKNOSEL_2026_SERBEST_KATEGORI_RAPOR_SABLONU.pdf structure.
> **Proje No:** SÜ2686 | **Proje Adı:** KonyaWaste | **Konu:** Kaynak Verimliliği ve Sürdürülebilirlik

---

### 1. PROJE ÖZETİ (20 Puan)

**Proje konusu:** Bu projenin konusu, kentsel atık yönetimini optimize etmek amacıyla Nesnelerin İnterneti (IoT) teknolojisi kullanılarak geliştirilen, çöp ve atık konteynerlerinin doluluk oranını uzaktan takip eden; ayrıca atıkları türüne göre ayıran ve plastik, karton gibi sıkıştırılabilir malzemeleri mekanik baskı ile hacimsel olarak küçülten akıllı bir donanım/haberleşme sisteminin tasarlanmasıdır.

**Proje amacı:** Projenin temel amacı, çöp toplama araçlarının rotalarını veri odaklı ve dinamik hale getirerek, yalnızca doluluk eşiğini aşan konteynerlere yönlendirilmelerini sağlamaktır. Entegre sıralama mekanizması, geri dönüştürülebilir malzemelerin (plastik, karton, metal) karışmasını önleyerek geri dönüşüm verimliliğini artırır. Sıkıştırma ünitesi ise plastikleri ve kartonları yaklaşık 3:1 oranında preslayerek depolama hacmini üçe katlar ve koleksiyon sıklığını azaltır. Bu sayede yerel yönetimlerin atık yönetim süreçlerinde gereksiz yakıt tüketimini, personel zaman kaybını ve karbon emisyonunu en aza indirmek hedeflenmektedir.

**Proje kapsamı ve yöntemi:** Proje kapsamı; çöp konteynerlerinin doluluk oranlarını anlık tespit eden akıllı bir donanım sistemi geliştirmeyi, sensör tetiklemeli ayırma kapıları ile atık sınıflandırmasını ve motorlu piston basıncıyla sıkıştırma işlemini tek bir retrofit modülde birleştirmeyi içermektedir. Şehir kaynaklarının israfını önleyen bu yaklaşım, "Kaynak Verimliliği ve Sürdürülebilirlik" kategorisiyle doğrudan ilişkilidir.

---

### 2. ÇÖZÜM ÜRETTİĞİ SORUN / İHTİYAÇ (25 Puan)

**Problem tanımı / Projenin Gerekçesi:** Geleneksel kentsel atık yönetimi sistemlerinde, yerel yönetimlere ait çöp toplama araçları, konteynerlerin anlık doluluk durumlarını bilmeksizin önceden belirlenmiş sabit (statik) rotaları takip etmektedir. Bu durum, henüz dolmamış konteynerlerin gereksiz yere ziyaret edilmesine yol açarak ciddi oranlarda yakıt, zaman ve iş gücü israfına neden olmaktadır. Diğer yandan, insan sirkülasyonunun yüksek olduğu bölgelerde hızlı dolan konteynerler zamanında boşaltılamamakta; bu da çöplerin taşmasına, çevreye kötü koku yayılmasına, karbon emisyonunun artmasına ve halk sağlığını tehdit eden hijyenik sorunlara zemin hazırlamaktadır. Ek olarak, sıkıştırılabilir malzemeler (plastik şişeler, karton kutular) orijinal hacimlerinin çok üzerinde alan kaplayarak konteynerlerin erken dolmasına yol açmakta; karışık atık kabul edilmesi ise geri dönüşüm verimliliğini düşürmektedir.

**Mevcut çözümler:** Günümüzde mevcut teknolojik çözümler genellikle "ultrasonik sensörler" kullanan sistemlerdir. Ancak ultrasonik sensörler, çöp konteyneri içindeki düzensiz atık yüzeylerinden ve dar alanlardan yansıyan yankılar nedeniyle sıklıkla hatalı ölçümler yapmakta, fiziksel çevre şartlarından (nem, koku gazları) olumsuz etkilenmektedir. Ayrıca mevcut akıllı konteyner çözümleri yalnızca doluluk takibi sunmakta; atık ayırma veya sıkıştırma işlevi içermemektedir.

**Çözüm fikri:** Önerilen çözüm; ESP32 mikrodenetleyici tabanlı, ölçüm için su geçirmez lazer mesafe (ToF) sensörleri kullanan, servo tetiklemeli çift bölmeli ayırma kapısıyla atıkları türüne göre (metal / sıkıştırılabilir) yönlendiren ve sıkıştırılabilir malzemeler bölmesinde vidalı mil tahrikli motorlu piston ile 3:1 oranında hacim azaltımı gerçekleştiren akıllı bir IoT donanımıdır. Veriler Wi-Fi üzerinden bulut sunucuya aktarılarak belediye araç filolarının rotaları dinamik olarak optimize edilir.

---

### 3. YERLİLİK VE ÖZGÜNLÜK TARAFI (10 Puan)

**Özgünlük ve Yerlilik:** Geliştirilen çözüm fikrinin özgünlüğü, pilot uygulama alanı olarak hedeflenen Konya ilindeki mevcut atık yönetimi altyapısına getirdiği yenilikçi ve dönüştürücü yaklaşımdan kaynaklanmaktadır. Yapılan saha gözlemleri ve incelemeler sonucunda, mevcut çöp konteynerlerinde herhangi bir anlık doluluk tespit cihazı, IoT tabanlı sensör altyapısı, atık ayırma mekanizması veya sıkıştırma sistemi bulunmadığı; toplama operasyonlarının tamamen geleneksel, manuel ve statik rotalar üzerinden yürütüldüğü tespit edilmiştir.

Sistemin teknolojik özgünlüğü üç katmanlıdır: (1) Literatürdeki benzer çalışmalarda sıkça rastlanan ve hatalı ölçümlere neden olan standart ultrasonik sensörler yerine, milimetrik hassasiyete sahip su geçirmez Lazer (ToF) sensörlerin kullanılması; (2) sensör tetiklemeli ayırma kapısıyla atık akışlarının konteyner içinde otomatik sınıflandırılması; (3) sıkıştırılabilir malzemelere (plastik, karton) uygulanan in-situ mekanik pres sayesinde depolama hacminin artırılması. Bu üç fonksiyonun tek bir retrofit modülde birleştirilmesi literatürde eşdeğeri bulunmayan bir kombinasyonu temsil etmektedir.

Projenin yerlilik boyutu değerlendirildiğinde; mecburi dış kaynaklı mikroelektronik bileşenler (ESP32, lazer diyot, DC motor sürücü vb.) haricinde kalan tüm sistem mimarisi, veri işleme algoritmaları, 3D yazıcı ile üretilecek IP67 standartlarındaki dış koruma muhafazası ve sunucu tarafındaki yönetim paneli (dashboard) tamamen yerli imkânlarla tasarlanmaktadır.

---

### 4. YÖNTEM VE HEDEF KİTLE (15 Puan)

**Verimlilik ve etkinlik:** Sistemin en büyük verimliliği üç birleşik etki üzerinden ölçülür: (1) çöp toplama filolarının kat ettiği mesafeyi ve harcadığı zamanı optimize etmesi; (2) sıkıştırma mekanizması sayesinde sıkıştırılabilir malzemelerin hacmini yaklaşık 3:1 oranında azaltarak konteynerlerin fiilen dolmadan önce taşıma kapasitesini üçe katlaması; (3) ayırma kapısı sayesinde geri dönüştürülebilir malzemelerin kirletilmeden toplanması. Konteynerlere yerleştirilen cihazlar, gereksiz veri trafiğini önlemek adına sadece çöp seviyesi belirli bir eşiği (%80 ve üzeri) aştığında veya günlük rutin bir "hayatta kalma" (keep-alive) sinyali gönderilmesi gerektiğinde uyanarak Wi-Fi üzerinden haberleşir.

**İzlenecek yöntem:** Proje; prototipleme, yazılım geliştirme ve saha testleri olmak üzere üç ana metodolojik aşamadan oluşmaktadır.
- **I. Kullanılacak Malzemeler:** Veri işleme ve haberleşme için ESP32 mikrodenetleyici, mesafe ölçümü için IP67 standartlarında su geçirmez ToF Lazer Sensör, atık sınıflandırması için IR reflektans ve iletkenlik sensör dizisi, ayırma için servo kapı mekanizması, sıkıştırma için vidalı mil tahrikli DC motor + piston ünitesi, sistemin güç ihtiyacı için Li-ion batarya modülü ve dış ortam koşullarına dayanıklı 3D baskı muhafaza kullanılacaktır.
- **II. İzlenecek Metot ve Deney:** İlk aşamada sensör kalibrasyonları ve ayırma mekanizması laboratuvar (breadboard) ortamında test edilecektir. Sıkıştırma pistonu, plastik şişeler ve karton kutular üzerinde kuvvet-hacim azalması ölçümleri ile doğrulanacaktır.
- **III. Analizler:** Hazırlanan prototip, yetkililer tarafından uygun görülürse, gerçek bir belediye çöp konteynerine monte edilerek 2 haftalık saha testine (pilot uygulama) tabi tutulacak; doluluk doğruluğu, sıkıştırma oranı ve sınıflandırma isabeti analiz edilecektir.

**Hedef kitle:** Projenin birincil hedef kitlesi, çöp toplama operasyonlarını yürüten Konya Büyükşehir, Selçuklu ve diğer ilçe belediyeleridir. İkincil hedef kitle ise; kendi içinde büyük çaplı atık yönetimi yapan Organize Sanayi Bölgeleri (OSB), geniş üniversite kampüsleri, büyük hastane kompleksleri ve özel atık yönetimi firmalarıdır.

---

### 5. YENİLİK VE TİCARİLEŞME POTANSİYELİ (10 Puan)

**Teknoloji ve yenilik:** Ürünün içerdiği en temel ileri teknolojik yenilik, mevcut doluluk takip sistemlerinin hiçbirinde bulunmayan atık ayırma + in-situ sıkıştırma fonksiyonlarının IoT bağlantılı bir konteyner sistemine entegrasyonudur. Mesafe ölçümü için kullanılan Lazer (ToF) teknolojisi, ışık fotonlarının gidiş-dönüş süresini hesaplayarak milimetrik ölçüm yapan ileri düzey bir sensör sistemidir. Servo kapı + motorlu piston kombinasyonu, tamamen mekanik olarak güvenilir, alanda onarılabilir ve 3D baskı parçalarla düşük maliyetle üretilecek şekilde tasarlanmıştır.

**Ticarileştirme potansiyeli:** Geliştirilen çözüm, düşük donanım maliyetleri ve yerli 3D üretim olanakları sayesinde seri üretime son derece uygundur. Proje, "Hizmet Olarak Yazılım" (SaaS) modeliyle ticarileştirilebilir; yerel yönetimlere veya OSB'lere cihaz satışı/kiralaması yapılarak merkezi yönetim paneli üzerinden abonelik tabanlı bir gelir modeli oluşturulabilir. Sıkıştırma özelliği sayesinde müşteriler aynı konteyner sayısıyla üç katına kadar daha fazla atık saklayabilmekte, bu da koleksiyon sıklığını düşürerek somut bir mali tasarruf sağlamaktadır.

---

### 6. PROJE TAKVİMİ (5 Puan)

İş paketleri, zamanlama ve iş bölümü: Proje geliştirme süreci, yarışma final tarihine kadar eksiksiz tamamlanacak şekilde 6 temel iş paketine (İP) bölünmüştür. Her bir iş paketi, donanım tedarikinden başlayarak Selçuklu ilçesindeki pilot saha testlerine ve nihai raporlamaya kadar olan süreci kapsayacak biçimde detaylandırılmış ve görev dağılımlarıyla birlikte haftalık bazda görselleştirilmiştir.

---

### 7. KAYNAKÇA (5 Puan)

- T.C. Çevre, Şehircilik ve İklim Değişikliği Bakanlığı. (2017). *Ulusal Atık Yönetimi ve Eylem Planı (2023-2035)*. Ankara: Çevre Yönetimi Genel Müdürlüğü. https://cygm.csb.gov.tr/
- Pardini, K., Rodrigues, J. J., Kozlov, S. A., Kumar, N., & Furtado, V. (2019). IoT-based solid waste management solutions: A survey. *Journal of Sensor and Actuator Networks*, 8(1), 5. https://doi.org/10.3390/jsan8010005
- Anagnostopoulos, T., Zaslavsky, A., Kolomvatsos, K., et al. (2017). Challenges and opportunities of waste management in smart cities. *IEEE Transactions on Industrial Informatics*, 13(4), 275–289. https://doi.org/10.1109/TII.2016.2620685
- Espressif Systems. (2024). *ESP32 Series Datasheet*. https://www.espressif.com/sites/default/files/documentation/esp32_datasheet_en.pdf
- Zanella, A., Bui, N., Castellani, A., Vangelista, L., & Zorzi, M. (2014). Internet of Things for Smart Cities. *IEEE Internet of Things Journal*, 1(1), 22–32. https://doi.org/10.1109/JIOT.2014.2306328
- Espressif Systems. (2023). *ESP32 Deep Sleep & Wakeup Strategies*. https://docs.espressif.com/projects/esp-idf/en/latest/esp32/api-reference/system/sleep_modes.html
- Folianto, F., Low, Y. S., & Yeow, W. L. (2015). Smart trash bin: Garbage management system using sensor motes. *2015 IEEE ISSNIP*. https://doi.org/10.1109/ISSNIP.2015.7106974
