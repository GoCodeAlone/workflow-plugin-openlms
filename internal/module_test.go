package internal

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestModuleInit_RegistersClient(t *testing.T) {
	m, err := newOpenLMSModule("test-init", map[string]any{
		"siteUrl": "https://lms.example.com",
		"token":   "testtoken",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err != nil {
		t.Fatal(err)
	}
	c, ok := GetClient("test-init")
	if !ok || c == nil {
		t.Error("expected client to be registered")
	}
	UnregisterClient("test-init")
}

func TestModuleStop_UnregistersClient(t *testing.T) {
	m, _ := newOpenLMSModule("test-stop", map[string]any{
		"siteUrl": "https://lms.example.com",
		"token":   "testtoken",
	})
	_ = m.Init()
	_ = m.Stop(context.Background())
	_, ok := GetClient("test-stop")
	if ok {
		t.Error("expected client to be unregistered after stop")
	}
}

func TestModuleInit_MissingSiteURL(t *testing.T) {
	m, err := newOpenLMSModule("test-missing-url", map[string]any{
		"token": "testtoken",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err == nil {
		t.Error("expected error for missing siteUrl")
		UnregisterClient("test-missing-url")
	}
}

func TestModuleInit_MissingToken(t *testing.T) {
	m, err := newOpenLMSModule("test-missing-token", map[string]any{
		"siteUrl": "https://lms.example.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err == nil {
		t.Error("expected error for missing token")
		UnregisterClient("test-missing-token")
	}
}

func TestModuleInit_RESTful(t *testing.T) {
	m, err := newOpenLMSModule("test-restful", map[string]any{
		"siteUrl": "https://lms.example.com",
		"token":   "testtoken",
		"restful": true,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Init(); err != nil {
		t.Fatal(err)
	}
	c, ok := GetClient("test-restful")
	if !ok || c == nil {
		t.Error("expected client to be registered")
	}
	if !c.restful {
		t.Error("expected restful mode to be enabled")
	}
	UnregisterClient("test-restful")
}

func TestMoodleClient_StandardMode(t *testing.T) {
	// Mock Moodle server
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Fatal(err)
		}
		if r.FormValue("wsfunction") != "core_webservice_get_site_info" {
			t.Errorf("unexpected wsfunction: %s", r.FormValue("wsfunction"))
		}
		if r.FormValue("wstoken") != "testtoken" {
			t.Errorf("unexpected wstoken: %s", r.FormValue("wstoken"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"sitename": "Test Moodle",
			"release":  "4.1",
		})
	}))
	defer srv.Close()

	client := newMoodleClient(srv.URL, "testtoken", false)
	result, err := client.callToMap(context.Background(), "core_webservice_get_site_info", map[string]string{})
	if err != nil {
		t.Fatal(err)
	}
	if result["sitename"] != "Test Moodle" {
		t.Errorf("unexpected sitename: %v", result["sitename"])
	}
}

func TestMoodleClient_RESTfulMode(t *testing.T) {
	var capturedPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"sitename": "RESTful Moodle"})
	}))
	defer srv.Close()

	client := newMoodleClient(srv.URL, "testtoken", true)
	result, err := client.callToMap(context.Background(), "core_webservice_get_site_info", map[string]string{})
	if err != nil {
		t.Fatal(err)
	}
	if capturedPath != "/webservice/restful/server.php/core_webservice_get_site_info" {
		t.Errorf("unexpected path: %s", capturedPath)
	}
	if result["sitename"] != "RESTful Moodle" {
		t.Errorf("unexpected sitename: %v", result["sitename"])
	}
}

func TestMoodleClient_MoodleError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"exception": "moodle_exception",
			"errorcode": "invalidtoken",
			"message":   "Invalid token - token not found",
		})
	}))
	defer srv.Close()

	client := newMoodleClient(srv.URL, "badtoken", false)
	_, err := client.callToMap(context.Background(), "core_webservice_get_site_info", map[string]string{})
	if err == nil {
		t.Error("expected error for Moodle exception")
	}
}

func TestMoodleClient_SliceResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]any{
			map[string]any{"id": float64(1), "username": "student1"},
			map[string]any{"id": float64(2), "username": "student2"},
		})
	}))
	defer srv.Close()

	client := newMoodleClient(srv.URL, "testtoken", false)
	result, err := client.callToSlice(context.Background(), "core_user_get_users_by_field", map[string]string{
		"field":     "username",
		"values[0]": "student1",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 users, got %d", len(result))
	}
}

func TestStepRegistry_AllStepTypes(t *testing.T) {
	types := allStepTypes()
	if len(types) < 120 {
		t.Errorf("expected at least 120 step types, got %d", len(types))
	}
	// Spot-check a few
	required := []string{
		"step.openlms_user_create",
		"step.openlms_course_create",
		"step.openlms_enrol_manual_enrol",
		"step.openlms_quiz_get_by_course",
		"step.openlms_call_function",
		"step.openlms_xapi_statement_post",
		"step.openlms_scorm_get_by_course",
	}
	typeSet := make(map[string]bool, len(types))
	for _, t := range types {
		typeSet[t] = true
	}
	for _, req := range required {
		if !typeSet[req] {
			t.Errorf("missing step type: %s", req)
		}
	}
}

func TestUserCreateStep_MissingClient(t *testing.T) {
	step, _ := newUserCreateStep("test", map[string]any{"module": "nonexistent"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestUserCreateStep_MissingRequired(t *testing.T) {
	// Register a dummy client
	RegisterClient("test-create", newMoodleClient("http://localhost", "tok", false))
	defer UnregisterClient("test-create")

	step, _ := newUserCreateStep("test", map[string]any{"module": "test-create"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing required fields")
	}
}

func TestCallFunctionStep_MissingClient(t *testing.T) {
	step, _ := newCallFunctionStep("test", map[string]any{"module": "nonexistent-callf"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["error"] == nil {
		t.Error("expected error for missing client")
	}
}

func TestSiteGetInfoStep_WithMockServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"sitename": "My LMS",
			"version":  "2023100900",
		})
	}))
	defer srv.Close()

	RegisterClient("test-siteinfo", newMoodleClient(srv.URL, "tok", false))
	defer UnregisterClient("test-siteinfo")

	step, _ := newSiteGetInfoStep("test", map[string]any{"module": "test-siteinfo"})
	result, err := step.Execute(context.Background(), nil, nil, map[string]any{}, nil, map[string]any{})
	if err != nil {
		t.Fatal(err)
	}
	if result.Output["sitename"] != "My LMS" {
		t.Errorf("unexpected sitename: %v", result.Output["sitename"])
	}
}
