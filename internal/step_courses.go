package internal

import (
	"context"
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// courseCreateStep implements step.openlms_course_create → core_course_create_courses
type courseCreateStep struct {
	name       string
	moduleName string
}

func newCourseCreateStep(name string, config map[string]any) (*courseCreateStep, error) {
	return &courseCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	shortname := resolveValue("shortname", current, config)
	fullname := resolveValue("fullname", current, config)
	categoryid := resolveValue("categoryid", current, config)
	if shortname == "" || fullname == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "shortname and fullname are required"}}, nil
	}
	if categoryid == "" {
		categoryid = "1"
	}
	params := map[string]string{
		"courses[0][shortname]":  shortname,
		"courses[0][fullname]":   fullname,
		"courses[0][categoryid]": categoryid,
	}
	for _, f := range []string{"summary", "format", "visible"} {
		if v := resolveValue(f, current, config); v != "" {
			params[fmt.Sprintf("courses[0][%s]", f)] = v
		}
	}
	result, err := client.callToSlice(ctx, "core_course_create_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var course map[string]any
	if len(result) > 0 {
		course, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"course": course, "courses": result}}, nil
}

// courseUpdateStep implements step.openlms_course_update → core_course_update_courses
type courseUpdateStep struct {
	name       string
	moduleName string
}

func newCourseUpdateStep(name string, config map[string]any) (*courseUpdateStep, error) {
	return &courseUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	params := map[string]string{"courses[0][id]": courseID}
	for _, f := range []string{"shortname", "fullname", "summary", "categoryid", "format", "visible"} {
		if v := resolveValue(f, current, config); v != "" {
			params[fmt.Sprintf("courses[0][%s]", f)] = v
		}
	}
	_, err := client.call(ctx, "core_course_update_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"updated": true, "courseid": courseID}}, nil
}

// courseDeleteStep implements step.openlms_course_delete → core_course_delete_courses
type courseDeleteStep struct {
	name       string
	moduleName string
}

func newCourseDeleteStep(name string, config map[string]any) (*courseDeleteStep, error) {
	return &courseDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	_, err := client.call(ctx, "core_course_delete_courses", map[string]string{"courseids[0]": courseID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "courseid": courseID}}, nil
}

// courseGetStep implements step.openlms_course_get → core_course_get_courses
type courseGetStep struct {
	name       string
	moduleName string
}

func newCourseGetStep(name string, config map[string]any) (*courseGetStep, error) {
	return &courseGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["options[ids][0]"] = courseID
	}
	result, err := client.callToSlice(ctx, "core_course_get_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var course map[string]any
	if len(result) > 0 {
		course, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"course": course, "courses": result}}, nil
}

// courseGetByFieldStep implements step.openlms_course_get_by_field → core_course_get_courses_by_field
type courseGetByFieldStep struct {
	name       string
	moduleName string
}

func newCourseGetByFieldStep(name string, config map[string]any) (*courseGetByFieldStep, error) {
	return &courseGetByFieldStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseGetByFieldStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	field := resolveValue("field", current, config)
	value := resolveValue("value", current, config)
	params := map[string]string{}
	if field != "" {
		params["field"] = field
	}
	if value != "" {
		params["value"] = value
	}
	result, err := client.callToMap(ctx, "core_course_get_courses_by_field", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	courses, _ := result["courses"].([]any)
	var course map[string]any
	if len(courses) > 0 {
		course, _ = courses[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"course": course, "courses": courses}}, nil
}

// courseSearchStep implements step.openlms_course_search → core_course_search_courses
type courseSearchStep struct {
	name       string
	moduleName string
}

func newCourseSearchStep(name string, config map[string]any) (*courseSearchStep, error) {
	return &courseSearchStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseSearchStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	criterianame := resolveValue("criterianame", current, config)
	criteriavalue := resolveValue("criteriavalue", current, config)
	if criterianame == "" {
		criterianame = "search"
	}
	params := map[string]string{
		"criterianame":  criterianame,
		"criteriavalue": criteriavalue,
	}
	result, err := client.callToMap(ctx, "core_course_search_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	courses, _ := result["courses"].([]any)
	return &sdk.StepResult{Output: map[string]any{"courses": courses, "total": result["total"]}}, nil
}

// courseGetContentsStep implements step.openlms_course_get_contents → core_course_get_contents
type courseGetContentsStep struct {
	name       string
	moduleName string
}

func newCourseGetContentsStep(name string, config map[string]any) (*courseGetContentsStep, error) {
	return &courseGetContentsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseGetContentsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	result, err := client.callToSlice(ctx, "core_course_get_contents", map[string]string{"courseid": courseID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"sections": result, "count": len(result)}}, nil
}

// courseGetCategoriesStep implements step.openlms_course_get_categories → core_course_get_categories
type courseGetCategoriesStep struct {
	name       string
	moduleName string
}

func newCourseGetCategoriesStep(name string, config map[string]any) (*courseGetCategoriesStep, error) {
	return &courseGetCategoriesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseGetCategoriesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if catID := resolveValue("categoryid", current, config); catID != "" {
		params["criteria[0][key]"] = "id"
		params["criteria[0][value]"] = catID
	}
	result, err := client.callToSlice(ctx, "core_course_get_categories", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"categories": result, "count": len(result)}}, nil
}

// courseCreateCategoriesStep implements step.openlms_course_create_categories → core_course_create_categories
type courseCreateCategoriesStep struct {
	name       string
	moduleName string
}

func newCourseCreateCategoriesStep(name string, config map[string]any) (*courseCreateCategoriesStep, error) {
	return &courseCreateCategoriesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseCreateCategoriesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	catName := resolveValue("name", current, config)
	if catName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "name is required"}}, nil
	}
	params := map[string]string{
		"categories[0][name]":     catName,
		"categories[0][parent]":   resolveValue("parent", current, config),
		"categories[0][idnumber]": resolveValue("idnumber", current, config),
	}
	result, err := client.callToSlice(ctx, "core_course_create_categories", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var category map[string]any
	if len(result) > 0 {
		category, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"category": category, "categories": result}}, nil
}

// courseDeleteCategoriesStep implements step.openlms_course_delete_categories → core_course_delete_categories
type courseDeleteCategoriesStep struct {
	name       string
	moduleName string
}

func newCourseDeleteCategoriesStep(name string, config map[string]any) (*courseDeleteCategoriesStep, error) {
	return &courseDeleteCategoriesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseDeleteCategoriesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	catID := resolveValue("categoryid", current, config)
	if catID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "categoryid is required"}}, nil
	}
	params := map[string]string{"categories[0][id]": catID}
	if newParent := resolveValue("newparent", current, config); newParent != "" {
		params["categories[0][newparent]"] = newParent
	}
	_, err := client.call(ctx, "core_course_delete_categories", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "categoryid": catID}}, nil
}

// courseDuplicateStep implements step.openlms_course_duplicate → core_course_duplicate_course
type courseDuplicateStep struct {
	name       string
	moduleName string
}

func newCourseDuplicateStep(name string, config map[string]any) (*courseDuplicateStep, error) {
	return &courseDuplicateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *courseDuplicateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	fullname := resolveValue("fullname", current, config)
	shortname := resolveValue("shortname", current, config)
	categoryid := resolveValue("categoryid", current, config)
	if courseID == "" || fullname == "" || shortname == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid, fullname, and shortname are required"}}, nil
	}
	if categoryid == "" {
		categoryid = "1"
	}
	params := map[string]string{
		"courseid":   courseID,
		"fullname":   fullname,
		"shortname":  shortname,
		"categoryid": categoryid,
	}
	result, err := client.callToMap(ctx, "core_course_duplicate_course", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
