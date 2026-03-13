package internal

import (
	"context"
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type competencyCreateStep struct {
	name       string
	moduleName string
}

func newCompetencyCreateStep(name string, config map[string]any) (*competencyCreateStep, error) {
	return &competencyCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *competencyCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	shortname := resolveValue("shortname", current, config)
	frameworkID := resolveValue("competencyframeworkid", current, config)
	if shortname == "" || frameworkID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "shortname and competencyframeworkid are required"}}, nil
	}
	params := map[string]string{
		"competency[shortname]":              shortname,
		"competency[competencyframeworkid]":  frameworkID,
	}
	if desc := resolveValue("description", current, config); desc != "" {
		params["competency[description]"] = desc
	}
	if idnumber := resolveValue("idnumber", current, config); idnumber != "" {
		params["competency[idnumber]"] = idnumber
	}
	result, err := client.callToMap(ctx, "core_competency_create_competency", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type competencyListStep struct {
	name       string
	moduleName string
}

func newCompetencyListStep(name string, config map[string]any) (*competencyListStep, error) {
	return &competencyListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *competencyListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	frameworkID := resolveValue("competencyframeworkid", current, config)
	if frameworkID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "competencyframeworkid is required"}}, nil
	}
	result, err := client.callToSlice(ctx, "core_competency_list_competencies", map[string]string{
		"filters[competencyframeworkid]": frameworkID,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"competencies": result, "count": len(result)}}, nil
}

type competencyDeleteStep struct {
	name       string
	moduleName string
}

func newCompetencyDeleteStep(name string, config map[string]any) (*competencyDeleteStep, error) {
	return &competencyDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *competencyDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	compID := resolveValue("id", current, config)
	if compID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "id is required"}}, nil
	}
	_, err := client.call(ctx, "core_competency_delete_competency", map[string]string{"id": compID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "id": compID}}, nil
}

type competencyCreateFrameworkStep struct {
	name       string
	moduleName string
}

func newCompetencyCreateFrameworkStep(name string, config map[string]any) (*competencyCreateFrameworkStep, error) {
	return &competencyCreateFrameworkStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *competencyCreateFrameworkStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	shortname := resolveValue("shortname", current, config)
	idnumber := resolveValue("idnumber", current, config)
	if shortname == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "shortname is required"}}, nil
	}
	params := map[string]string{
		"competencyframework[shortname]": shortname,
		"competencyframework[idnumber]":  idnumber,
	}
	if scaleid := resolveValue("scaleid", current, config); scaleid != "" {
		params["competencyframework[scaleid]"] = scaleid
	}
	result, err := client.callToMap(ctx, "core_competency_create_competency_framework", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type competencyListFrameworksStep struct {
	name       string
	moduleName string
}

func newCompetencyListFrameworksStep(name string, config map[string]any) (*competencyListFrameworksStep, error) {
	return &competencyListFrameworksStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *competencyListFrameworksStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, _ map[string]any, _ map[string]any, _ map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	result, err := client.callToSlice(ctx, "core_competency_list_competency_frameworks", map[string]string{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"frameworks": result, "count": len(result)}}, nil
}

type competencyCreatePlanStep struct {
	name       string
	moduleName string
}

func newCompetencyCreatePlanStep(name string, config map[string]any) (*competencyCreatePlanStep, error) {
	return &competencyCreatePlanStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *competencyCreatePlanStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	planName := resolveValue("name", current, config)
	userID := resolveValue("userid", current, config)
	if planName == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "name and userid are required"}}, nil
	}
	params := map[string]string{
		"plan[name]":   planName,
		"plan[userid]": userID,
	}
	if desc := resolveValue("description", current, config); desc != "" {
		params["plan[description]"] = desc
	}
	result, err := client.callToMap(ctx, "core_competency_create_plan", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type competencyListPlansStep struct {
	name       string
	moduleName string
}

func newCompetencyListPlansStep(name string, config map[string]any) (*competencyListPlansStep, error) {
	return &competencyListPlansStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *competencyListPlansStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	if userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid is required"}}, nil
	}
	result, err := client.callToSlice(ctx, "core_competency_list_user_plans", map[string]string{"userid": userID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"plans": result, "count": len(result)}}, nil
}

type competencyAddToCourseStep struct {
	name       string
	moduleName string
}

func newCompetencyAddToCourseStep(name string, config map[string]any) (*competencyAddToCourseStep, error) {
	return &competencyAddToCourseStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *competencyAddToCourseStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	compID := resolveValue("competencyid", current, config)
	if courseID == "" || compID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid and competencyid are required"}}, nil
	}
	_, err := client.call(ctx, "core_competency_add_competency_to_course", map[string]string{
		"courseid": courseID, "competencyid": compID,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"added": true, "courseid": courseID, "competencyid": compID}}, nil
}

type competencyGradeStep struct {
	name       string
	moduleName string
}

func newCompetencyGradeStep(name string, config map[string]any) (*competencyGradeStep, error) {
	return &competencyGradeStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *competencyGradeStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	compID := resolveValue("competencyid", current, config)
	grade := resolveValue("grade", current, config)
	if userID == "" || compID == "" || grade == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid, competencyid, and grade are required"}}, nil
	}
	_, err := client.call(ctx, "core_competency_grade_competency_in_course", map[string]string{
		"courseid": resolveValue("courseid", current, config),
		"userid": userID, "competencyid": compID, "grade": grade,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"graded": true, "userid": userID, "competencyid": compID}}, nil
}

// suppress unused import
var _ = fmt.Sprintf
