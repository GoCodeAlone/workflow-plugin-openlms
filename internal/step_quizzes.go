package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// quizGetByCourseStep implements step.openlms_quiz_get_by_course → mod_quiz_get_quizzes_by_courses
type quizGetByCourseStep struct {
	name       string
	moduleName string
}

func newQuizGetByCourseStep(name string, config map[string]any) (*quizGetByCourseStep, error) {
	return &quizGetByCourseStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *quizGetByCourseStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseids[0]"] = courseID
	}
	result, err := client.callToMap(ctx, "mod_quiz_get_quizzes_by_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	quizzes, _ := result["quizzes"].([]any)
	return &sdk.StepResult{Output: map[string]any{"quizzes": quizzes, "count": len(quizzes)}}, nil
}

// quizGetAttemptsStep implements step.openlms_quiz_get_attempts → mod_quiz_get_attempts
type quizGetAttemptsStep struct {
	name       string
	moduleName string
}

func newQuizGetAttemptsStep(name string, config map[string]any) (*quizGetAttemptsStep, error) {
	return &quizGetAttemptsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *quizGetAttemptsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	quizID := resolveValue("quizid", current, config)
	if quizID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "quizid is required"}}, nil
	}
	params := map[string]string{"quizid": quizID}
	if userID := resolveValue("userid", current, config); userID != "" {
		params["userid"] = userID
	}
	result, err := client.callToMap(ctx, "mod_quiz_get_attempts", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	attempts, _ := result["attempts"].([]any)
	return &sdk.StepResult{Output: map[string]any{"attempts": attempts}}, nil
}

// quizGetAttemptDataStep implements step.openlms_quiz_get_attempt_data → mod_quiz_get_attempt_data
type quizGetAttemptDataStep struct {
	name       string
	moduleName string
}

func newQuizGetAttemptDataStep(name string, config map[string]any) (*quizGetAttemptDataStep, error) {
	return &quizGetAttemptDataStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *quizGetAttemptDataStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	attemptID := resolveValue("attemptid", current, config)
	page := resolveValue("page", current, config)
	if attemptID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "attemptid is required"}}, nil
	}
	if page == "" {
		page = "0"
	}
	result, err := client.callToMap(ctx, "mod_quiz_get_attempt_data", map[string]string{"attemptid": attemptID, "page": page})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// quizGetAttemptReviewStep implements step.openlms_quiz_get_attempt_review → mod_quiz_get_attempt_review
type quizGetAttemptReviewStep struct {
	name       string
	moduleName string
}

func newQuizGetAttemptReviewStep(name string, config map[string]any) (*quizGetAttemptReviewStep, error) {
	return &quizGetAttemptReviewStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *quizGetAttemptReviewStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	attemptID := resolveValue("attemptid", current, config)
	if attemptID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "attemptid is required"}}, nil
	}
	params := map[string]string{"attemptid": attemptID}
	if page := resolveValue("page", current, config); page != "" {
		params["page"] = page
	}
	result, err := client.callToMap(ctx, "mod_quiz_get_attempt_review", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// quizStartAttemptStep implements step.openlms_quiz_start_attempt → mod_quiz_start_attempt
type quizStartAttemptStep struct {
	name       string
	moduleName string
}

func newQuizStartAttemptStep(name string, config map[string]any) (*quizStartAttemptStep, error) {
	return &quizStartAttemptStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *quizStartAttemptStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	quizID := resolveValue("quizid", current, config)
	if quizID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "quizid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_quiz_start_attempt", map[string]string{"quizid": quizID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	attempt, _ := result["attempt"].(map[string]any)
	return &sdk.StepResult{Output: map[string]any{"attempt": attempt}}, nil
}

// quizSaveAttemptStep implements step.openlms_quiz_save_attempt → mod_quiz_save_attempt
type quizSaveAttemptStep struct {
	name       string
	moduleName string
}

func newQuizSaveAttemptStep(name string, config map[string]any) (*quizSaveAttemptStep, error) {
	return &quizSaveAttemptStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *quizSaveAttemptStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	attemptID := resolveValue("attemptid", current, config)
	if attemptID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "attemptid is required"}}, nil
	}
	params := map[string]string{"attemptid": attemptID}
	if timeup := resolveValue("timeup", current, config); timeup != "" {
		params["timeup"] = timeup
	}
	_, err := client.call(ctx, "mod_quiz_save_attempt", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"saved": true, "attemptid": attemptID}}, nil
}

// quizProcessAttemptStep implements step.openlms_quiz_process_attempt → mod_quiz_process_attempt
type quizProcessAttemptStep struct {
	name       string
	moduleName string
}

func newQuizProcessAttemptStep(name string, config map[string]any) (*quizProcessAttemptStep, error) {
	return &quizProcessAttemptStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *quizProcessAttemptStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	attemptID := resolveValue("attemptid", current, config)
	if attemptID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "attemptid is required"}}, nil
	}
	params := map[string]string{
		"attemptid": attemptID,
		"finishattempt": resolveValue("finishattempt", current, config),
		"timeup":        resolveValue("timeup", current, config),
	}
	result, err := client.callToMap(ctx, "mod_quiz_process_attempt", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
