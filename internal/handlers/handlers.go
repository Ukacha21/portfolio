package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"
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
	if r.URL.Path != "/" && !strings.HasPrefix(r.URL.Path, "/project/") {
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

// HandleHealthz handles health checks for container platforms (Render, Docker, Kubernetes)
func (h *Handler) HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
