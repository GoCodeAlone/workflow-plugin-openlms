package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type completionGetActivitiesStatusStep struct {
	name       string
	moduleName string
}

func newCompletionGetActivitiesStatusStep(name string, config map[string]any) (*completionGetActivitiesStatusStep, error) {
	return &completionGetActivitiesStatusStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *completionGetActivitiesStatusStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	userID := resolveValue("userid", current, config)
	if courseID == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid and userid are required"}}, nil
	}
	result, err := client.callToMap(ctx, "core_completion_get_activities_completion_status", map[string]string{
		"courseid": courseID, "userid": userID,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	statuses, _ := result["statuses"].([]any)
	return &sdk.StepResult{Output: map[string]any{"statuses": statuses}}, nil
}

type completionGetCourseStatusStep struct {
	name       string
	moduleName string
}

func newCompletionGetCourseStatusStep(name string, config map[string]any) (*completionGetCourseStatusStep, error) {
	return &completionGetCourseStatusStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *completionGetCourseStatusStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	userID := resolveValue("userid", current, config)
	if courseID == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid and userid are required"}}, nil
	}
	result, err := client.callToMap(ctx, "core_completion_get_course_completion_status", map[string]string{
		"courseid": courseID, "userid": userID,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type completionUpdateActivityStep struct {
	name       string
	moduleName string
}

func newCompletionUpdateActivityStep(name string, config map[string]any) (*completionUpdateActivityStep, error) {
	return &completionUpdateActivityStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *completionUpdateActivityStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	cmid := resolveValue("cmid", current, config)
	completed := resolveValue("completed", current, config)
	if cmid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "cmid is required"}}, nil
	}
	if completed == "" {
		completed = "1"
	}
	_, err := client.call(ctx, "core_completion_update_activity_completion_status_manually", map[string]string{
		"cmid": cmid, "completed": completed,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"updated": true, "cmid": cmid}}, nil
}

type completionMarkSelfCompletedStep struct {
	name       string
	moduleName string
}

func newCompletionMarkSelfCompletedStep(name string, config map[string]any) (*completionMarkSelfCompletedStep, error) {
	return &completionMarkSelfCompletedStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *completionMarkSelfCompletedStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	_, err := client.call(ctx, "core_completion_mark_course_self_completed", map[string]string{"courseid": courseID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"completed": true, "courseid": courseID}}, nil
}
