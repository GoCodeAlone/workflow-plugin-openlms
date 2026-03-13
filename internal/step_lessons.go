package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type lessonGetByCourseStep struct {
	name       string
	moduleName string
}

func newLessonGetByCourseStep(name string, config map[string]any) (*lessonGetByCourseStep, error) {
	return &lessonGetByCourseStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *lessonGetByCourseStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseids[0]"] = courseID
	}
	result, err := client.callToMap(ctx, "mod_lesson_get_lessons_by_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	lessons, _ := result["lessons"].([]any)
	return &sdk.StepResult{Output: map[string]any{"lessons": lessons}}, nil
}

type lessonGetPagesStep struct {
	name       string
	moduleName string
}

func newLessonGetPagesStep(name string, config map[string]any) (*lessonGetPagesStep, error) {
	return &lessonGetPagesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *lessonGetPagesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	lessonID := resolveValue("lessonid", current, config)
	if lessonID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "lessonid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_lesson_get_pages", map[string]string{"lessonid": lessonID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	pages, _ := result["pages"].([]any)
	return &sdk.StepResult{Output: map[string]any{"pages": pages}}, nil
}

type lessonGetPageDataStep struct {
	name       string
	moduleName string
}

func newLessonGetPageDataStep(name string, config map[string]any) (*lessonGetPageDataStep, error) {
	return &lessonGetPageDataStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *lessonGetPageDataStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	lessonID := resolveValue("lessonid", current, config)
	pageid := resolveValue("pageid", current, config)
	if lessonID == "" || pageid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "lessonid and pageid are required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_lesson_get_page_data", map[string]string{
		"lessonid": lessonID, "pageid": pageid,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type lessonLaunchAttemptStep struct {
	name       string
	moduleName string
}

func newLessonLaunchAttemptStep(name string, config map[string]any) (*lessonLaunchAttemptStep, error) {
	return &lessonLaunchAttemptStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *lessonLaunchAttemptStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	lessonID := resolveValue("lessonid", current, config)
	if lessonID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "lessonid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_lesson_launch_attempt", map[string]string{"lessonid": lessonID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type lessonProcessPageStep struct {
	name       string
	moduleName string
}

func newLessonProcessPageStep(name string, config map[string]any) (*lessonProcessPageStep, error) {
	return &lessonProcessPageStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *lessonProcessPageStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	lessonID := resolveValue("lessonid", current, config)
	pageid := resolveValue("pageid", current, config)
	if lessonID == "" || pageid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "lessonid and pageid are required"}}, nil
	}
	params := map[string]string{"lessonid": lessonID, "pageid": pageid}
	if data := resolveValue("data", current, config); data != "" {
		params["data[0][name]"] = "answer"
		params["data[0][value]"] = data
	}
	result, err := client.callToMap(ctx, "mod_lesson_process_page", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type lessonFinishAttemptStep struct {
	name       string
	moduleName string
}

func newLessonFinishAttemptStep(name string, config map[string]any) (*lessonFinishAttemptStep, error) {
	return &lessonFinishAttemptStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *lessonFinishAttemptStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	lessonID := resolveValue("lessonid", current, config)
	if lessonID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "lessonid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_lesson_finish_attempt", map[string]string{"lessonid": lessonID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
