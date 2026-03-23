package internal_test

import (
	"testing"

	"github.com/GoCodeAlone/workflow/wftest"
)

// TestIntegration_UserGet verifies a pipeline that fetches a user by ID and
// confirms the result is propagated to subsequent steps.
func TestIntegration_UserGet(t *testing.T) {
	h := wftest.New(t,
		wftest.WithYAML(`
pipelines:
  lookup_user:
    steps:
      - name: get_user
        type: step.openlms_user_get
        config:
          key: id
          value: "42"
      - name: confirm
        type: step.set
        config:
          values:
            found: true
`),
		wftest.MockStep("step.openlms_user_get", wftest.Returns(map[string]any{
			"user":  map[string]any{"id": "42", "username": "jdoe", "email": "jdoe@example.com"},
			"users": []any{map[string]any{"id": "42", "username": "jdoe"}},
			"total": "1",
		})),
	)

	result := h.ExecutePipeline("lookup_user", nil)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if result.Output["found"] != true {
		t.Errorf("expected found=true, got %v", result.Output["found"])
	}
}

// TestIntegration_CourseGet verifies a pipeline that fetches a course and
// records how many times the step was invoked.
func TestIntegration_CourseGet(t *testing.T) {
	rec := wftest.RecordStep("step.openlms_course_get")
	rec.WithOutput(map[string]any{
		"course":  map[string]any{"id": "10", "shortname": "CS101", "fullname": "Intro to CS"},
		"courses": []any{},
	})

	h := wftest.New(t,
		wftest.WithYAML(`
pipelines:
  lookup_course:
    steps:
      - name: get_course
        type: step.openlms_course_get
        config:
          key: id
          value: "10"
      - name: mark_done
        type: step.set
        config:
          values:
            fetched: true
`),
		rec,
	)

	result := h.ExecutePipeline("lookup_course", nil)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if result.Output["fetched"] != true {
		t.Errorf("expected fetched=true, got %v", result.Output["fetched"])
	}
	if rec.CallCount() != 1 {
		t.Errorf("expected step called once, got %d", rec.CallCount())
	}
}

// TestIntegration_EnrolManualEnrol verifies a pipeline that enrols a user in a
// course and exposes the enrolment result in the pipeline output.
func TestIntegration_EnrolManualEnrol(t *testing.T) {
	h := wftest.New(t,
		wftest.WithYAML(`
pipelines:
  enrol_user:
    steps:
      - name: enrol
        type: step.openlms_enrol_manual_enrol
        config:
          userid: "7"
          courseid: "10"
          roleid: "5"
      - name: done
        type: step.set
        config:
          values:
            enrolled: true
`),
		wftest.MockStep("step.openlms_enrol_manual_enrol", wftest.Returns(map[string]any{
			"enrolled": true,
			"userid":   "7",
			"courseid": "10",
		})),
	)

	result := h.ExecutePipeline("enrol_user", nil)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if result.Output["enrolled"] != true {
		t.Errorf("expected enrolled=true, got %v", result.Output["enrolled"])
	}
}
