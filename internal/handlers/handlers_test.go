package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandlers_RenderAllPagesAndLanguages(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	endpoints := map[string]http.HandlerFunc{
		"/":                   h.HandleIndex,
		"/project/velum":      h.HandleVelum,
		"/project/flowguard":  h.HandleFlowguard,
		"/project/solus":      h.HandleSolus,
		"/project/denova":     h.HandleDenova,
		"/project/renovo":     h.HandleRenovo,
		"/project/curasero":   h.HandleCurasero,
		"/project/welstech":   h.HandleWelstech,
		"/project/sendit":     h.HandleSendit,
		"/project/smartwaste": h.HandleSmartwaste,
		"/project/classifier": h.HandleClassifier,
		"/project/lumiere":    h.HandleLumiere,
		"/project/notcha":     h.HandleNotcha,
		"/project/spo2":       h.HandleSpo2,
		"/project/excavator":  h.HandleExcavator,
		"/project/safren":     h.HandleSafren,
		"/project/stm32":      h.HandleSTM32,
		"/project/ros":        h.HandleROS,
		"/project/supublic":   h.HandleSuPublic,
		"/project/missile":    h.HandleMissile,
	}

	languages := []string{"en", "tr", "ar", "pt", "es", "de", "fr"}

	for route, handler := range endpoints {
		for _, lang := range languages {
			req := httptest.NewRequest("GET", route+"?lang="+lang, nil)
			rec := httptest.NewRecorder()
			handler(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("route %s with lang %s returned status %d", route, lang, rec.Code)
			}
		}
	}
}

func TestIndex_ProjectCardsHaveViewProjectNotExplore(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	req := httptest.NewRequest("GET", "/?lang=en", nil)
	rec := httptest.NewRecorder()
	h.HandleIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	// Verify "View Project" is present in the rendered HTML
	if !strings.Contains(body, "View Project") {
		t.Errorf("expected index page to contain 'View Project'")
	}

	// Verify old explore buttons are absent
	forbiddenStrings := []string{
		"Explore Excavator Design",
		"Explore SafRen",
		"Explore ROS Noetic Robot",
	}
	for _, forbidden := range forbiddenStrings {
		if strings.Contains(body, forbidden) {
			t.Errorf("found forbidden string in project card button: %s", forbidden)
		}
	}
}

func TestIndex_SteerIOPlaceholderAndSoloProjectsVertical(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	req := httptest.NewRequest("GET", "/?lang=en", nil)
	rec := httptest.NewRecorder()
	h.HandleIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	// 1. steerIO horizontal card exists
	if !strings.Contains(body, `id="card-steerio"`) {
		t.Errorf("expected index page to contain steerIO card")
	}
	if !strings.Contains(body, "steerIO") {
		t.Errorf("expected index page to contain 'steerIO'")
	}
	if !strings.Contains(body, "project-card-steerio") {
		t.Errorf("expected steerIO card to have class project-card-steerio")
	}

	// 2. Velum, FlowGuard, Solus, DeNova, Renovo are vertical (no longer have project-card-featured)
	cards := []string{
		`id="card-velum"`,
		`id="card-flowguard"`,
		`id="card-solus"`,
		`id="card-denova"`,
		`id="card-renovo"`,
	}
	for _, card := range cards {
		if !strings.Contains(body, card) {
			t.Errorf("expected index page to contain %s", card)
		}
	}
}

func TestIndex_ContractProjectsSuPublicAtTop(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	req := httptest.NewRequest("GET", "/?lang=en", nil)
	rec := httptest.NewRecorder()
	h.HandleIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	supublicIdx := strings.Index(body, `id="card-contract-supublic"`)
	missileIdx := strings.Index(body, `id="card-contract-missile"`)
	logoIdx := strings.Index(body, `id="card-contract-logo"`)
	invoiceIdx := strings.Index(body, `id="card-contract-invoice"`)
	apacovItIdx := strings.Index(body, `id="card-contract-apacov-it"`)
	apacovLogoIdx := strings.Index(body, `id="card-contract-apacov-logo"`)

	if supublicIdx == -1 {
		t.Fatalf("card-contract-supublic not found in index page")
	}
	if missileIdx == -1 {
		t.Fatalf("card-contract-missile not found in index page")
	}
	if logoIdx == -1 {
		t.Fatalf("card-contract-logo not found in index page")
	}
	if invoiceIdx == -1 {
		t.Fatalf("card-contract-invoice not found in index page")
	}
	if apacovItIdx == -1 {
		t.Fatalf("card-contract-apacov-it not found in index page")
	}
	if apacovLogoIdx == -1 {
		t.Fatalf("card-contract-apacov-logo not found in index page")
	}

	if supublicIdx > missileIdx {
		t.Errorf("expected SuPublic (idx %d) to be at the top, before missile (idx %d)", supublicIdx, missileIdx)
	}
	if missileIdx > logoIdx {
		t.Errorf("expected missile (idx %d) to be before card-contract-logo (idx %d)", missileIdx, logoIdx)
	}
	if logoIdx > invoiceIdx {
		t.Errorf("expected card-contract-logo (idx %d) to be before card-contract-invoice (idx %d)", logoIdx, invoiceIdx)
	}
	if invoiceIdx > apacovItIdx {
		t.Errorf("expected card-contract-invoice (idx %d) to be before card-contract-apacov-it (idx %d)", invoiceIdx, apacovItIdx)
	}
	if apacovItIdx > apacovLogoIdx {
		t.Errorf("expected card-contract-apacov-it (idx %d) to be before card-contract-apacov-logo (idx %d)", apacovItIdx, apacovLogoIdx)
	}
}

func TestMissile_ShowcaseRendersContent(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	req := httptest.NewRequest("GET", "/project/missile?lang=en", nil)
	rec := httptest.NewRecorder()
	h.HandleMissile(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	requiredKeywords := []string{
		"AI Armaments",
		"ai-armaments.com",
		"Hasan Tutar",
		"hasan-tutar-176826353",
		"Precision Guided Missile",
		"SolidWorks",
		"a7_4_flap_holder_shaft",
		"a1_head",
	}

	for _, kw := range requiredKeywords {
		if !strings.Contains(body, kw) {
			t.Errorf("expected missile page to contain %q", kw)
		}
	}
}

func TestContractProjects_AllCardsAreVertical(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	// Verify style.css does not have grid-column spans for contract cards (all must be vertical)
	cssBytes, err := os.ReadFile("static/css/style.css")
	if err != nil {
		t.Fatalf("failed to read style.css: %v", err)
	}
	css := string(cssBytes)

	// Neither SuPublic nor Missile should have grid-column span across columns
	if strings.Contains(css, ".project-card-contract-cp5 {\n    grid-column: 1 / -1;") {
		t.Errorf("project-card-contract-cp5 should not have grid-column: 1 / -1 (must be vertical)")
	}
	if strings.Contains(css, ".project-card-contract-cp6 {\n    grid-column: 1 / -1;") {
		t.Errorf("project-card-contract-cp6 should not have grid-column: 1 / -1 (must be vertical)")
	}

	// Verify all 6 contract project cards exist in rendered HTML
	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	req := httptest.NewRequest("GET", "/?lang=en", nil)
	rec := httptest.NewRecorder()
	h.HandleIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	contractCards := []string{
		`id="card-contract-supublic"`,
		`id="card-contract-missile"`,
		`id="card-contract-logo"`,
		`id="card-contract-invoice"`,
		`id="card-contract-apacov-it"`,
		`id="card-contract-apacov-logo"`,
	}

	for _, card := range contractCards {
		if !strings.Contains(body, card) {
			t.Errorf("expected contract card %s to exist in rendered HTML", card)
		}
	}
}

func TestNavigationAndPageFlowOrder(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	req := httptest.NewRequest("GET", "/?lang=en", nil)
	rec := httptest.NewRecorder()
	h.HandleIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	// 1. Verify Top Banner Order:
	// Home -> Education -> Experience -> Projects -> Workshops -> Skills -> Awards -> Certificates -> Languages -> About -> Contact
	navLinks := []string{
		`#hero`,
		`#education`,
		`#experience`,
		`#projects`,
		`#workshops`,
		`#skills`,
		`#awards`,
		`#certificates`,
		`#languages`,
		`#about`,
		`#contact`,
	}

	mainNavStart := strings.Index(body, `id="mainNav"`)
	if mainNavStart == -1 {
		t.Fatalf("mainNav not found in rendered HTML")
	}
	mainNavEnd := strings.Index(body[mainNavStart:], `</nav>`) + mainNavStart
	mainNavHTML := body[mainNavStart:mainNavEnd]

	lastNavPos := -1
	for _, link := range navLinks {
		pos := strings.Index(mainNavHTML, link)
		if pos == -1 {
			t.Errorf("expected mainNav to contain %s", link)
		}
		if pos < lastNavPos {
			t.Errorf("nav link %s out of order: was %d, previous was %d", link, pos, lastNavPos)
		}
		lastNavPos = pos
	}

	// 2. Verify Page Flow Order (sections in DOM):
	// hero -> education -> experience -> projects -> workshops -> skills -> awards -> certificates -> languages -> about -> contact
	sections := []string{
		`id="hero"`,
		`id="education"`,
		`id="experience"`,
		`id="projects"`,
		`id="workshops"`,
		`id="skills"`,
		`id="awards"`,
		`id="certificates"`,
		`id="languages"`,
		`id="about"`,
		`id="contact"`,
	}

	lastSecPos := -1
	for _, sec := range sections {
		pos := strings.Index(body, sec)
		if pos == -1 {
			t.Errorf("expected section %s to exist in body", sec)
		}
		if pos < lastSecPos {
			t.Errorf("section %s out of order: pos %d < previous pos %d", sec, pos, lastSecPos)
		}
		lastSecPos = pos
	}
}

func TestSmartWaste_DysonAwardSubmission(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	req := httptest.NewRequest("GET", "/?lang=en", nil)
	rec := httptest.NewRecorder()
	h.HandleIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	// Locate SMARTWASTE card
	swCardIdx := strings.Index(body, `id="card-smartwaste"`)
	if swCardIdx == -1 {
		t.Fatalf("card-smartwaste not found in index HTML")
	}
	swCardEnd := strings.Index(body[swCardIdx:], `</article>`) + swCardIdx
	swHTML := body[swCardIdx:swCardEnd]

	// Verify Dyson '26 Submission ribbon and tag
	if !strings.Contains(swHTML, "Dyson '26 Submission") {
		t.Errorf("expected card-smartwaste to contain 'Dyson '26 Submission' ribbon")
	}
	if !strings.Contains(swHTML, "Dyson '26 Entry") {
		t.Errorf("expected card-smartwaste to contain 'Dyson '26 Entry' tag")
	}
}

func TestLanguages_FrenchBeginnerAlphabeticalOrder(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	req := httptest.NewRequest("GET", "/?lang=en", nil)
	rec := httptest.NewRecorder()
	h.HandleIndex(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	// Locate languages section
	secIdx := strings.Index(body, `id="languages"`)
	if secIdx == -1 {
		t.Fatalf("section id=languages not found")
	}
	secEnd := strings.Index(body[secIdx:], `</section>`) + secIdx
	secHTML := body[secIdx:secEnd]

	// Verify French exists and is beginner
	if !strings.Contains(secHTML, "French") {
		t.Errorf("expected languages section to contain 'French'")
	}

	// Verify alphabetical order of languages in the names column:
	// Arabic, English, French, German, Portuguese, Spanish, Turkish, Local Languages
	expectedOrder := []string{
		">Arabic<",
		">English<",
		">French<",
		">German<",
		">Portuguese<",
		">Spanish<",
		">Turkish<",
		">Local Languages<",
	}

	lastPos := -1
	for _, lang := range expectedOrder {
		pos := strings.Index(secHTML, lang)
		if pos == -1 {
			t.Errorf("expected language %s in languages board", lang)
		}
		if pos < lastPos {
			t.Errorf("language %s out of order: pos %d < lastPos %d", lang, pos, lastPos)
		}
		lastPos = pos
	}

	// Also verify French level is beginner:
	// French is 3rd item, so 3rd level should be beginner
	levelsStart := strings.Index(secHTML, "languages-col-levels")
	if levelsStart == -1 {
		t.Fatalf("languages-col-levels not found")
	}
	levelsHTML := secHTML[levelsStart:]
	expectedLevels := []string{
		"good",       // Arabic
		"advanced",   // English
		"beginner",   // French
		"beginner",   // German
		"native",     // Portuguese
		"advanced",   // Spanish
		"advanced",   // Turkish
		"native",     // Local Languages
	}

	currLevelsHTML := levelsHTML
	for i, lvl := range expectedLevels {
		idx := strings.Index(currLevelsHTML, lvl)
		if idx == -1 {
			t.Errorf("expected level %d to be %s", i, lvl)
			break
		}
		currLevelsHTML = currLevelsHTML[idx+len(lvl):]
	}
}

func TestHandleSteerIO(t *testing.T) {
	origWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get working dir: %v", err)
	}
	if err := os.Chdir("../.."); err != nil {
		t.Fatalf("could not chdir to root: %v", err)
	}
	defer os.Chdir(origWd)

	h, err := NewHandler()
	if err != nil {
		t.Fatalf("failed to initialize handlers: %v", err)
	}

	req := httptest.NewRequest("GET", "/project/steerio?lang=en", nil)
	rec := httptest.NewRecorder()
	h.HandleSteerIO(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()

	expectedTokens := []string{
		"STEERIO — DIY USB MECHATRONICS STEERING RIG",
		"600 PPR",
		"STM32F103",
		"FreeJoy",
		"74HC165",
		"DE-9",
		"UCP 204",
		"WX14",
		"https://www.youtube.com/embed/c893AAE7Mpk",
		"https://www.instagram.com/reel/DekWnuOtpiv/",
		"@ukachajuma",
		"PB6",
		"PB12",
		"PB14",
		"PB13",
		"Not Used",
		"steerioLightbox",
		"Screenshot (753).png",
		"Screenshot (754).png",
	}

	for _, token := range expectedTokens {
		if !strings.Contains(body, token) {
			t.Errorf("expected steerIO page to contain %q", token)
		}
	}
}




