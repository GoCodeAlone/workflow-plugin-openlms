package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type cohortCreateStep struct {
	name       string
	moduleName string
}

func newCohortCreateStep(name string, config map[string]any) (*cohortCreateStep, error) {
	return &cohortCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *cohortCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	cohortName := resolveValue("name", current, config)
	idnumber := resolveValue("idnumber", current, config)
	if cohortName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "name is required"}}, nil
	}
	params := map[string]string{
		"cohorts[0][categorytype][type]":  "id",
		"cohorts[0][categorytype][value]": resolveValue("contextid", current, config),
		"cohorts[0][name]":                cohortName,
		"cohorts[0][idnumber]":            idnumber,
	}
	if desc := resolveValue("description", current, config); desc != "" {
		params["cohorts[0][description]"] = desc
	}
	result, err := client.callToSlice(ctx, "core_cohort_create_cohorts", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var cohort map[string]any
	if len(result) > 0 {
		cohort, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"cohort": cohort, "cohorts": result}}, nil
}

type cohortDeleteStep struct {
	name       string
	moduleName string
}

func newCohortDeleteStep(name string, config map[string]any) (*cohortDeleteStep, error) {
	return &cohortDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *cohortDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	cohortID := resolveValue("cohortid", current, config)
	if cohortID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "cohortid is required"}}, nil
	}
	_, err := client.call(ctx, "core_cohort_delete_cohorts", map[string]string{"cohortids[0]": cohortID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "cohortid": cohortID}}, nil
}

type cohortGetStep struct {
	name       string
	moduleName string
}

func newCohortGetStep(name string, config map[string]any) (*cohortGetStep, error) {
	return &cohortGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *cohortGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	cohortID := resolveValue("cohortid", current, config)
	if cohortID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "cohortid is required"}}, nil
	}
	result, err := client.callToSlice(ctx, "core_cohort_get_cohorts", map[string]string{"cohortids[0]": cohortID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var cohort map[string]any
	if len(result) > 0 {
		cohort, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"cohort": cohort, "cohorts": result}}, nil
}

type cohortSearchStep struct {
	name       string
	moduleName string
}

func newCohortSearchStep(name string, config map[string]any) (*cohortSearchStep, error) {
	return &cohortSearchStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *cohortSearchStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	query := resolveValue("query", current, config)
	contextID := resolveValue("contextid", current, config)
	if contextID == "" {
		contextID = "1"
	}
	result, err := client.callToMap(ctx, "core_cohort_search_cohorts", map[string]string{
		"query":     query,
		"context[contextid]": contextID,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	cohorts, _ := result["cohorts"].([]any)
	return &sdk.StepResult{Output: map[string]any{"cohorts": cohorts}}, nil
}

type cohortAddMembersStep struct {
	name       string
	moduleName string
}

func newCohortAddMembersStep(name string, config map[string]any) (*cohortAddMembersStep, error) {
	return &cohortAddMembersStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *cohortAddMembersStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	cohortID := resolveValue("cohortid", current, config)
	userID := resolveValue("userid", current, config)
	if cohortID == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "cohortid and userid are required"}}, nil
	}
	_, err := client.call(ctx, "core_cohort_add_cohort_members", map[string]string{
		"members[0][cohorttype][type]":  "id",
		"members[0][cohorttype][value]": cohortID,
		"members[0][usertype][type]":    "id",
		"members[0][usertype][value]":   userID,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"added": true, "cohortid": cohortID, "userid": userID}}, nil
}

type cohortDeleteMembersStep struct {
	name       string
	moduleName string
}

func newCohortDeleteMembersStep(name string, config map[string]any) (*cohortDeleteMembersStep, error) {
	return &cohortDeleteMembersStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *cohortDeleteMembersStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	cohortID := resolveValue("cohortid", current, config)
	userID := resolveValue("userid", current, config)
	if cohortID == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "cohortid and userid are required"}}, nil
	}
	_, err := client.call(ctx, "core_cohort_delete_cohort_members", map[string]string{
		"members[0][cohortid]": cohortID,
		"members[0][userid]":   userID,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "cohortid": cohortID, "userid": userID}}, nil
}
