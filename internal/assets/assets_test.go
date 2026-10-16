package assets

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAssetHandler_LocalAsset(t *testing.T) {
	h := NewAssetHandler()
	h.localDir = "../../assets"

	req := httptest.NewRequest("GET", "/assets/img/ukacha_avatar.png", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	ct := rec.Header().Get("Content-Type")
	if ct != "image/png" {
		t.Fatalf("expected image/png content-type, got %s", ct)
	}

	cc := rec.Header().Get("Cache-Control")
	if cc == "" {
		t.Fatalf("expected Cache-Control header to be set")
	}
}

func TestAssetHandler_DirectoryTraversal(t *testing.T) {
	h := NewAssetHandler()
	h.localDir = "../../assets"

	req := httptest.NewRequest("GET", "/assets/../../main.go", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected status 403 Forbidden, got %d", rec.Code)
	}
}
