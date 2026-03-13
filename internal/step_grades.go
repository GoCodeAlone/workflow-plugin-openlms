package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// gradeGetGradesStep implements step.openlms_grade_get_grades → core_grades_get_grades
type gradeGetGradesStep struct {
	name       string
	moduleName string
}

func newGradeGetGradesStep(name string, config map[string]any) (*gradeGetGradesStep, error) {
	return &gradeGetGradesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *gradeGetGradesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	params := map[string]string{"courseid": courseID}
	if component := resolveValue("component", current, config); component != "" {
		params["component"] = component
	}
	if activityID := resolveValue("activityid", current, config); activityID != "" {
		params["activityid"] = activityID
	}
	if userID := resolveValue("userid", current, config); userID != "" {
		params["userids[0]"] = userID
	}
	result, err := client.callToMap(ctx, "core_grades_get_grades", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// gradeUpdateGradesStep implements step.openlms_grade_update_grades → core_grades_update_grades
type gradeUpdateGradesStep struct {
	name       string
	moduleName string
}

func newGradeUpdateGradesStep(name string, config map[string]any) (*gradeUpdateGradesStep, error) {
	return &gradeUpdateGradesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *gradeUpdateGradesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	source := resolveValue("source", current, config)
	courseID := resolveValue("courseid", current, config)
	component := resolveValue("component", current, config)
	activityID := resolveValue("activityid", current, config)
	if source == "" || courseID == "" || component == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "source, courseid, and component are required"}}, nil
	}
	params := map[string]string{
		"source":                       source,
		"courseid":                     courseID,
		"component":                    component,
		"activityid":                   activityID,
		"grades[0][studentid]":         resolveValue("studentid", current, config),
		"grades[0][rawgrade]":          resolveValue("rawgrade", current, config),
	}
	result, err := client.call(ctx, "core_grades_update_grades", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	if m, ok := result.(map[string]any); ok {
		return &sdk.StepResult{Output: m}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"updated": true}}, nil
}

// gradeGetGradeItemsStep implements step.openlms_grade_get_grade_items → gradereport_user_get_grade_items
type gradeGetGradeItemsStep struct {
	name       string
	moduleName string
}

func newGradeGetGradeItemsStep(name string, config map[string]any) (*gradeGetGradeItemsStep, error) {
	return &gradeGetGradeItemsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *gradeGetGradeItemsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	params := map[string]string{"courseid": courseID}
	if userID := resolveValue("userid", current, config); userID != "" {
		params["userid"] = userID
	}
	result, err := client.callToMap(ctx, "gradereport_user_get_grade_items", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// gradeGetGradesTableStep implements step.openlms_grade_get_grades_table → gradereport_user_get_grades_table
type gradeGetGradesTableStep struct {
	name       string
	moduleName string
}

func newGradeGetGradesTableStep(name string, config map[string]any) (*gradeGetGradesTableStep, error) {
	return &gradeGetGradesTableStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *gradeGetGradesTableStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	params := map[string]string{"courseid": courseID}
	if userID := resolveValue("userid", current, config); userID != "" {
		params["userid"] = userID
	}
	result, err := client.callToMap(ctx, "gradereport_user_get_grades_table", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
