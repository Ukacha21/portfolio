package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ukacha21/portfolio/internal/assets"
	"github.com/ukacha21/portfolio/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3008"
	}

	h, err := handlers.NewHandler()
	if err != nil {
		log.Fatalf("Fatal error initializing handlers: %v", err)
	}

	mux := http.NewServeMux()

	// Static Assets Server (CSS & JS) with cache headers
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") || strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Cache-Control", "public, max-age=86400")
		} else {
			w.Header().Set("Cache-Control", "public, max-age=604800")
		}
		fs.ServeHTTP(w, r)
	})))

	// Assets Server (Images & Docs fetched from GitHub / local assets)
	assetHandler := assets.NewAssetHandler()
	mux.Handle("/assets/", assetHandler)

	// Favicon fallback
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/assets/img/velum/51153107113199203185801911829216068233242248243.ico"
		assetHandler.ServeHTTP(w, r)
	})

	// Health Check for Render
	mux.HandleFunc("/healthz", h.HandleHealthz)

	// Routes
	mux.HandleFunc("/project/velum", h.HandleVelum)
	mux.HandleFunc("/project/flowguard", h.HandleFlowguard)
	mux.HandleFunc("/project/solus", h.HandleSolus)
	mux.HandleFunc("/project/denova", h.HandleDenova)
	mux.HandleFunc("/project/renovo", h.HandleRenovo)
	mux.HandleFunc("/project/curasero", h.HandleCurasero)
	mux.HandleFunc("/project/welstech", h.HandleWelstech)
	mux.HandleFunc("/project/sendit", h.HandleSendit)
	mux.HandleFunc("/project/smartwaste", h.HandleSmartwaste)
	mux.HandleFunc("/project/classifier", h.HandleClassifier)
	mux.HandleFunc("/project/lumiere", h.HandleLumiere)
	mux.HandleFunc("/project/notcha", h.HandleNotcha)
	mux.HandleFunc("/project/spo2", h.HandleSpo2)
	mux.HandleFunc("/project/excavator", h.HandleExcavator)
	mux.HandleFunc("/project/safren", h.HandleSafren)
	mux.HandleFunc("/project/stm32", h.HandleSTM32)
	mux.HandleFunc("/", h.HandleIndex)

	log.Printf("=====================================================")
	log.Printf(" UKACHA JUMA — MECHATRONICS PORTFOLIO SERVER")
	log.Printf(" Server running on port :%s", port)
	log.Printf(" Home Page: http://localhost:%s", port)
	log.Printf(" VELUM Showcase: http://localhost:%s/project/velum", port)
	log.Printf(" FLOWGUARD Showcase: http://localhost:%s/project/flowguard", port)
	log.Printf(" SOLUS Showcase: http://localhost:%s/project/solus", port)
	log.Printf(" DENOVA Showcase: http://localhost:%s/project/denova", port)
	log.Printf(" RENOVO Showcase: http://localhost:%s/project/renovo", port)
	log.Printf(" CURA SERO Showcase: http://localhost:%s/project/curasero", port)
	log.Printf(" WELSTECH Showcase: http://localhost:%s/project/welstech", port)
	log.Printf(" SENDIT Showcase: http://localhost:%s/project/sendit", port)
	log.Printf(" SMARTWASTE Showcase: http://localhost:%s/project/smartwaste", port)
	log.Printf(" CLASSIFIER Showcase: http://localhost:%s/project/classifier", port)
	log.Printf(" LUMIERE Showcase: http://localhost:%s/project/lumiere", port)
	log.Printf(" NOTCHA Showcase: http://localhost:%s/project/notcha", port)
	log.Printf(" SPO2 PCB Showcase: http://localhost:%s/project/spo2", port)
	log.Printf(" EXCAVATOR Showcase: http://localhost:%s/project/excavator", port)
	log.Printf(" SAFREN Showcase: http://localhost:%s/project/safren", port)
	log.Printf(" STM32 Showcase: http://localhost:%s/project/stm32", port)
	log.Printf(" Health check: http://localhost:%s/healthz", port)
	log.Printf("=====================================================")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server stopped unexpectedly: %v", err)
	}
}
