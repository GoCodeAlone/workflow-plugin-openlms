package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// enrolGetEnrolledUsersStep implements step.openlms_enrol_get_enrolled_users → core_enrol_get_enrolled_users
type enrolGetEnrolledUsersStep struct {
	name       string
	moduleName string
}

func newEnrolGetEnrolledUsersStep(name string, config map[string]any) (*enrolGetEnrolledUsersStep, error) {
	return &enrolGetEnrolledUsersStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *enrolGetEnrolledUsersStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	result, err := client.callToSlice(ctx, "core_enrol_get_enrolled_users", map[string]string{"courseid": courseID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"users": result, "count": len(result)}}, nil
}

// enrolGetUserCoursesStep implements step.openlms_enrol_get_user_courses → core_enrol_get_users_courses
type enrolGetUserCoursesStep struct {
	name       string
	moduleName string
}

func newEnrolGetUserCoursesStep(name string, config map[string]any) (*enrolGetUserCoursesStep, error) {
	return &enrolGetUserCoursesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *enrolGetUserCoursesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	if userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid is required"}}, nil
	}
	result, err := client.callToSlice(ctx, "core_enrol_get_users_courses", map[string]string{"userid": userID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"courses": result, "count": len(result)}}, nil
}

// enrolManualEnrolStep implements step.openlms_enrol_manual_enrol → enrol_manual_enrol_users
type enrolManualEnrolStep struct {
	name       string
	moduleName string
}

func newEnrolManualEnrolStep(name string, config map[string]any) (*enrolManualEnrolStep, error) {
	return &enrolManualEnrolStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *enrolManualEnrolStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	courseID := resolveValue("courseid", current, config)
	if userID == "" || courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid and courseid are required"}}, nil
	}
	params := map[string]string{
		"enrolments[0][userid]":   userID,
		"enrolments[0][courseid]": courseID,
	}
	if roleID := resolveValue("roleid", current, config); roleID != "" {
		params["enrolments[0][roleid]"] = roleID
	}
	_, err := client.call(ctx, "enrol_manual_enrol_users", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"enrolled": true, "userid": userID, "courseid": courseID}}, nil
}

// enrolManualUnenrolStep implements step.openlms_enrol_manual_unenrol → enrol_manual_unenrol_users
type enrolManualUnenrolStep struct {
	name       string
	moduleName string
}

func newEnrolManualUnenrolStep(name string, config map[string]any) (*enrolManualUnenrolStep, error) {
	return &enrolManualUnenrolStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *enrolManualUnenrolStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	courseID := resolveValue("courseid", current, config)
	if userID == "" || courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid and courseid are required"}}, nil
	}
	params := map[string]string{
		"enrolments[0][userid]":   userID,
		"enrolments[0][courseid]": courseID,
	}
	_, err := client.call(ctx, "enrol_manual_unenrol_users", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"unenrolled": true, "userid": userID, "courseid": courseID}}, nil
}

// enrolSelfEnrolStep implements step.openlms_enrol_self_enrol → enrol_self_enrol_user
type enrolSelfEnrolStep struct {
	name       string
	moduleName string
}

func newEnrolSelfEnrolStep(name string, config map[string]any) (*enrolSelfEnrolStep, error) {
	return &enrolSelfEnrolStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *enrolSelfEnrolStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	params := map[string]string{"courseid": courseID}
	if password := resolveValue("password", current, config); password != "" {
		params["password"] = password
	}
	result, err := client.callToMap(ctx, "enrol_self_enrol_user", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

// enrolGetCourseMethodsStep implements step.openlms_enrol_get_course_methods → core_enrol_get_course_enrolment_methods
type enrolGetCourseMethodsStep struct {
	name       string
	moduleName string
}

func newEnrolGetCourseMethodsStep(name string, config map[string]any) (*enrolGetCourseMethodsStep, error) {
	return &enrolGetCourseMethodsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *enrolGetCourseMethodsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	result, err := client.callToSlice(ctx, "core_enrol_get_course_enrolment_methods", map[string]string{"courseid": courseID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"methods": result, "count": len(result)}}, nil
}
