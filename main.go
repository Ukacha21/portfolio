package main

import (
	"log"
	"net/http"
	"os"
	"strings"

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

	// Static Assets Server with cache headers
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") || strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Set("Cache-Control", "public, max-age=86400")
		} else if strings.HasSuffix(r.URL.Path, ".png") || strings.HasSuffix(r.URL.Path, ".jpg") || strings.HasSuffix(r.URL.Path, ".pdf") {
			w.Header().Set("Cache-Control", "public, max-age=604800")
		}
		fs.ServeHTTP(w, r)
	})))

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
