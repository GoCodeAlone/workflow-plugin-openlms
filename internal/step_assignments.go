package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// assignGetAssignmentsStep implements step.openlms_assign_get_assignments → mod_assign_get_assignments
type assignGetAssignmentsStep struct {
	name       string
	moduleName string
}

func newAssignGetAssignmentsStep(name string, config map[string]any) (*assignGetAssignmentsStep, error) {
	return &assignGetAssignmentsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *assignGetAssignmentsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseids[0]"] = courseID
	}
	result, err := client.callToMap(ctx, "mod_assign_get_assignments", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	courses, _ := result["courses"].([]any)
	return &sdk.StepResult{Output: map[string]any{"courses": courses, "warnings": result["warnings"]}}, nil
}

// assignGetSubmissionsStep implements step.openlms_assign_get_submissions → mod_assign_get_submissions
type assignGetSubmissionsStep struct {
	name       string
	moduleName string
}

func newAssignGetSubmissionsStep(name string, config map[string]any) (*assignGetSubmissionsStep, error) {
	return &assignGetSubmissionsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *assignGetSubmissionsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	assignID := resolveValue("assignid", current, config)
	if assignID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "assignid is required"}}, nil
	}
	params := map[string]string{"assignmentids[0]": assignID}
	result, err := client.callToMap(ctx, "mod_assign_get_submissions", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	assignments, _ := result["assignments"].([]any)
	return &sdk.StepResult{Output: map[string]any{"assignments": assignments, "warnings": result["warnings"]}}, nil
}

// assignGetGradesStep implements step.openlms_assign_get_grades → mod_assign_get_grades
type assignGetGradesStep struct {
	name       string
	moduleName string
}

func newAssignGetGradesStep(name string, config map[string]any) (*assignGetGradesStep, error) {
	return &assignGetGradesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *assignGetGradesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	assignID := resolveValue("assignid", current, config)
	if assignID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "assignid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_assign_get_grades", map[string]string{"assignmentids[0]": assignID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	assignments, _ := result["assignments"].([]any)
	return &sdk.StepResult{Output: map[string]any{"assignments": assignments}}, nil
}

// assignSaveSubmissionStep implements step.openlms_assign_save_submission → mod_assign_save_submission
type assignSaveSubmissionStep struct {
	name       string
	moduleName string
}

func newAssignSaveSubmissionStep(name string, config map[string]any) (*assignSaveSubmissionStep, error) {
	return &assignSaveSubmissionStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *assignSaveSubmissionStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	assignID := resolveValue("assignid", current, config)
	if assignID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "assignid is required"}}, nil
	}
	params := map[string]string{
		"assignmentid": assignID,
		"plugindata[onlinetext_editor][text]":   resolveValue("text", current, config),
		"plugindata[onlinetext_editor][format]": "1",
	}
	_, err := client.call(ctx, "mod_assign_save_submission", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"saved": true, "assignid": assignID}}, nil
}

// assignSubmitForGradingStep implements step.openlms_assign_submit_for_grading → mod_assign_submit_for_grading
type assignSubmitForGradingStep struct {
	name       string
	moduleName string
}

func newAssignSubmitForGradingStep(name string, config map[string]any) (*assignSubmitForGradingStep, error) {
	return &assignSubmitForGradingStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *assignSubmitForGradingStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	assignID := resolveValue("assignid", current, config)
	if assignID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "assignid is required"}}, nil
	}
	params := map[string]string{
		"assignmentid":       assignID,
		"acceptsubmissionstatement": "1",
	}
	_, err := client.call(ctx, "mod_assign_submit_for_grading", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"submitted": true, "assignid": assignID}}, nil
}

// assignSaveGradeStep implements step.openlms_assign_save_grade → mod_assign_save_grade
type assignSaveGradeStep struct {
	name       string
	moduleName string
}

func newAssignSaveGradeStep(name string, config map[string]any) (*assignSaveGradeStep, error) {
	return &assignSaveGradeStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *assignSaveGradeStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	assignID := resolveValue("assignid", current, config)
	userID := resolveValue("userid", current, config)
	grade := resolveValue("grade", current, config)
	if assignID == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "assignid and userid are required"}}, nil
	}
	params := map[string]string{
		"assignmentid":          assignID,
		"userid":                userID,
		"grade":                 grade,
		"attemptnumber":         "-1",
		"addattempt":            "0",
		"workflowstate":         resolveValue("workflowstate", current, config),
		"applytoall":            "0",
		"plugindata[assignfeedbackcomments_editor][text]": resolveValue("feedbackcomment", current, config),
		"plugindata[assignfeedbackcomments_editor][format]": "1",
	}
	_, err := client.call(ctx, "mod_assign_save_grade", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"graded": true, "assignid": assignID, "userid": userID}}, nil
}
