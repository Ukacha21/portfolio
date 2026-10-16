package assets

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// AssetHandler serves assets from local disk if available,
// or fetches and caches them from the GitHub repository in production.
type AssetHandler struct {
	localDir   string
	githubBase string
	token      string
	cacheDir   string
	client     *http.Client
}

// NewAssetHandler initializes the asset proxy and cache
func NewAssetHandler() *AssetHandler {
	githubBase := os.Getenv("GITHUB_ASSETS_BASE_URL")
	if githubBase == "" {
		githubBase = "https://raw.githubusercontent.com/Ukacha21/portfolio/main/assets"
	}
	githubBase = strings.TrimSuffix(githubBase, "/")

	cacheDir := os.Getenv("ASSET_CACHE_DIR")
	if cacheDir == "" {
		cacheDir = filepath.Join(os.TempDir(), "ukacha_portfolio_cache")
	}
	_ = os.MkdirAll(cacheDir, 0755)

	return &AssetHandler{
		localDir:   "assets",
		githubBase: githubBase,
		token:      os.Getenv("GITHUB_TOKEN"),
		cacheDir:   cacheDir,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *AssetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	relPath := strings.TrimPrefix(r.URL.Path, "/assets/")
	relPath = strings.TrimPrefix(relPath, "/")
	cleanPath := filepath.Clean(relPath)

	// Prevent directory traversal attacks
	if strings.HasPrefix(cleanPath, "..") || strings.HasPrefix(cleanPath, "/") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// 1. Try local assets directory first (fast offline/dev execution)
	localFilePath := filepath.Join(h.localDir, cleanPath)
	if fi, err := os.Stat(localFilePath); err == nil && !fi.IsDir() {
		setAssetHeaders(w, cleanPath)
		http.ServeFile(w, r, localFilePath)
		return
	}

	// 2. Check local disk cache (for production containers after first fetch)
	cachedFilePath := filepath.Join(h.cacheDir, cleanPath)
	if fi, err := os.Stat(cachedFilePath); err == nil && !fi.IsDir() {
		setAssetHeaders(w, cleanPath)
		http.ServeFile(w, r, cachedFilePath)
		return
	}

	// 3. Fetch from GitHub repository
	remoteURL := h.githubBase + "/" + strings.ReplaceAll(cleanPath, "\\", "/")
	req, err := http.NewRequestWithContext(r.Context(), "GET", remoteURL, nil)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if h.token != "" {
		req.Header.Set("Authorization", "token "+h.token)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		log.Printf("[Assets] Error fetching from GitHub (%s): %v", remoteURL, err)
		http.Error(w, "Asset Unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		http.NotFound(w, r)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("[Assets] GitHub returned status %d for %s", resp.StatusCode, remoteURL)
		http.Error(w, "Asset fetch failed", resp.StatusCode)
		return
	}

	// Read content
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[Assets] Error reading response body: %v", err)
		http.Error(w, "Failed to read asset", http.StatusInternalServerError)
		return
	}

	// Save to disk cache asynchronously
	go func(path string, content []byte) {
		destDir := filepath.Dir(path)
		if err := os.MkdirAll(destDir, 0755); err == nil {
			_ = os.WriteFile(path, content, 0644)
		}
	}(cachedFilePath, data)

	// Set content-type and caching headers
	setAssetHeaders(w, cleanPath)
	if cType := resp.Header.Get("Content-Type"); cType != "" && !strings.Contains(cType, "text/plain") {
		w.Header().Set("Content-Type", cType)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func setAssetHeaders(w http.ResponseWriter, path string) {
	lower := strings.ToLower(path)
	switch {
	case strings.HasSuffix(lower, ".png"):
		w.Header().Set("Content-Type", "image/png")
	case strings.HasSuffix(lower, ".jpg"), strings.HasSuffix(lower, ".jpeg"):
		w.Header().Set("Content-Type", "image/jpeg")
	case strings.HasSuffix(lower, ".ico"):
		w.Header().Set("Content-Type", "image/x-icon")
	case strings.HasSuffix(lower, ".pdf"):
		w.Header().Set("Content-Type", "application/pdf")
	case strings.HasSuffix(lower, ".svg"):
		w.Header().Set("Content-Type", "image/svg+xml")
	}
	w.Header().Set("Cache-Control", "public, max-age=604800, immutable")
}
