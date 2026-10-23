package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/ukacha21/portfolio/internal/i18n"
)

type Handler struct {
	pages map[string]*template.Template
}

func NewHandler() (*Handler, error) {
	pages := make(map[string]*template.Template)

	indexTmpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		return nil, err
	}
	pages["index"] = indexTmpl

	velumTmpl, err := template.ParseFiles("templates/base.html", "templates/velum.html")
	if err != nil {
		return nil, err
	}
	pages["velum"] = velumTmpl

	flowguardTmpl, err := template.ParseFiles("templates/base.html", "templates/flowguard.html")
	if err != nil {
		return nil, err
	}
	pages["flowguard"] = flowguardTmpl

	solusTmpl, err := template.ParseFiles("templates/base.html", "templates/solus.html")
	if err != nil {
		return nil, err
	}
	pages["solus"] = solusTmpl

	denovaTmpl, err := template.ParseFiles("templates/base.html", "templates/denova.html")
	if err != nil {
		return nil, err
	}
	pages["denova"] = denovaTmpl

	renovoTmpl, err := template.ParseFiles("templates/base.html", "templates/renovo.html")
	if err != nil {
		return nil, err
	}
	pages["renovo"] = renovoTmpl

	curaseroTmpl, err := template.ParseFiles("templates/base.html", "templates/curasero.html")
	if err != nil {
		return nil, err
	}
	pages["curasero"] = curaseroTmpl

	senditTmpl, err := template.ParseFiles("templates/base.html", "templates/sendit.html")
	if err != nil {
		return nil, err
	}
	pages["sendit"] = senditTmpl

	smartwasteTmpl, err := template.ParseFiles("templates/base.html", "templates/smartwaste.html")
	if err != nil {
		return nil, err
	}
	pages["smartwaste"] = smartwasteTmpl

	welstechTmpl, err := template.ParseFiles("templates/base.html", "templates/welstech.html")
	if err != nil {
		return nil, err
	}
	pages["welstech"] = welstechTmpl

	classifierTmpl, err := template.ParseFiles("templates/base.html", "templates/classifier.html")
	if err != nil {
		return nil, err
	}
	pages["classifier"] = classifierTmpl

	lumiereTmpl, err := template.ParseFiles("templates/base.html", "templates/lumiere.html")
	if err != nil {
		return nil, err
	}
	pages["lumiere"] = lumiereTmpl

	notchaTmpl, err := template.ParseFiles("templates/base.html", "templates/notcha.html")
	if err != nil {
		return nil, err
	}
	pages["notcha"] = notchaTmpl

	spo2Tmpl, err := template.ParseFiles("templates/base.html", "templates/spo2.html")
	if err != nil {
		return nil, err
	}
	pages["spo2"] = spo2Tmpl

	excavatorTmpl, err := template.ParseFiles("templates/base.html", "templates/excavator.html")
	if err != nil {
		return nil, err
	}
	pages["excavator"] = excavatorTmpl

	safrenTmpl, err := template.ParseFiles("templates/base.html", "templates/safren.html")
	if err != nil {
		return nil, err
	}
	pages["safren"] = safrenTmpl

	stm32Tmpl, err := template.ParseFiles("templates/base.html", "templates/stm32.html")
	if err != nil {
		return nil, err
	}
	pages["stm32"] = stm32Tmpl

	rosTmpl, err := template.ParseFiles("templates/base.html", "templates/ros.html")
	if err != nil {
		return nil, err
	}
	pages["ros"] = rosTmpl

	supublicTmpl, err := template.ParseFiles("templates/base.html", "templates/supublic.html")
	if err != nil {
		return nil, err
	}
	pages["supublic"] = supublicTmpl

	missileTmpl, err := template.ParseFiles("templates/base.html", "templates/missile.html")
	if err != nil {
		return nil, err
	}
	pages["missile"] = missileTmpl

	steerioTmpl, err := template.ParseFiles("templates/base.html", "templates/steerio.html")
	if err != nil {
		return nil, err
	}
	pages["steerio"] = steerioTmpl

	return &Handler{pages: pages}, nil
}

// setLanguageCookie sets the selected language in an HTTP cookie
func setLanguageCookie(w http.ResponseWriter, lang string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "portfolio_lang",
		Value:    lang,
		Path:     "/",
		Expires:  time.Now().Add(365 * 24 * time.Hour),
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	})
}

// HandleIndex handles the home page (Ukacha JUMA portfolio)
func (h *Handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["index"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering index: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleVelum handles the dedicated VELUM project showcase
func (h *Handler) HandleVelum(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["velum"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering velum: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleFlowguard handles the dedicated FLOWGUARD project showcase
func (h *Handler) HandleFlowguard(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["flowguard"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering flowguard: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleSolus handles the dedicated SOLUS project showcase
func (h *Handler) HandleSolus(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["solus"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering solus: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleDenova handles the dedicated DENOVA project showcase
func (h *Handler) HandleDenova(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["denova"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering denova: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleRenovo handles the dedicated RENOVO project showcase
func (h *Handler) HandleRenovo(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["renovo"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering renovo: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleCurasero handles the dedicated CURA SERO project showcase
func (h *Handler) HandleCurasero(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["curasero"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering curasero: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleSendit handles the dedicated SENDIT project showcase
func (h *Handler) HandleSendit(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["sendit"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering sendit: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleSmartwaste handles the dedicated SMARTWASTE project showcase
func (h *Handler) HandleSmartwaste(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["smartwaste"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering smartwaste: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleWelstech handles the dedicated WELSTECH wearable biosensor showcase
func (h *Handler) HandleWelstech(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["welstech"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering welstech: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleClassifier handles the dedicated IMAGE CLASSIFIER deep learning model showcase
func (h *Handler) HandleClassifier(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["classifier"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering classifier: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleLumiere handles the dedicated LUMIERE IoT smart streetlight failure reporting showcase
func (h *Handler) HandleLumiere(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["lumiere"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering lumiere: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleNotcha handles the dedicated NOTCHA dynamic time notch desktop utility showcase
func (h *Handler) HandleNotcha(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["notcha"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering notcha: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleSpo2 handles the dedicated 2-layer SPO2 Pulse Oximeter KiCad PCB Design showcase
func (h *Handler) HandleSpo2(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["spo2"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering spo2: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleExcavator handles the dedicated Excavator Design 24-ton mechanical capstone showcase
func (h *Handler) HandleExcavator(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["excavator"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering excavator: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleSafren handles the dedicated SafRen (formerly Ce-Saf) smart footwear pedometer & IoT system showcase
func (h *Handler) HandleSafren(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["safren"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering safren: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleSTM32 handles the dedicated STM32 Discovery Board challenge showcase
func (h *Handler) HandleSTM32(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["stm32"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering stm32: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleROS handles the dedicated ROS Noetic autonomous mobile robot showcase
func (h *Handler) HandleROS(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["ros"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering ros: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleSuPublic handles the dedicated SuPublic drinking water kiosk sheet metal showcase
func (h *Handler) HandleSuPublic(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["supublic"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering supublic: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleMissile handles the dedicated AI Armaments precision guided missile airframe showcase
func (h *Handler) HandleMissile(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["missile"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering missile: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleSteerIO handles the dedicated steerIO DIY USB mechatronics steering rig showcase
func (h *Handler) HandleSteerIO(w http.ResponseWriter, r *http.Request) {
	langCode := i18n.DetectLanguage(r)
	if r.URL.Query().Get("lang") != "" {
		setLanguageCookie(w, langCode)
	}

	data := i18n.GetLanguage(langCode)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.pages["steerio"].ExecuteTemplate(w, "base.html", data); err != nil {
		log.Printf("Template error rendering steerio: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// HandleHealthz handles health checks for container platforms (Render, Docker, Kubernetes)
func (h *Handler) HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
