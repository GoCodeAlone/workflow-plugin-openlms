package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type calendarCreateEventsStep struct {
	name       string
	moduleName string
}

func newCalendarCreateEventsStep(name string, config map[string]any) (*calendarCreateEventsStep, error) {
	return &calendarCreateEventsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *calendarCreateEventsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	eventName := resolveValue("name", current, config)
	if eventName == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "name is required"}}, nil
	}
	params := map[string]string{"events[0][name]": eventName}
	for _, f := range []string{"description", "courseid", "timestart", "timeduration", "eventtype"} {
		if v := resolveValue(f, current, config); v != "" {
			params["events[0]["+f+"]"] = v
		}
	}
	result, err := client.callToMap(ctx, "core_calendar_create_calendar_events", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	events, _ := result["events"].([]any)
	return &sdk.StepResult{Output: map[string]any{"events": events}}, nil
}

type calendarDeleteEventsStep struct {
	name       string
	moduleName string
}

func newCalendarDeleteEventsStep(name string, config map[string]any) (*calendarDeleteEventsStep, error) {
	return &calendarDeleteEventsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *calendarDeleteEventsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	eventID := resolveValue("eventid", current, config)
	if eventID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "eventid is required"}}, nil
	}
	_, err := client.call(ctx, "core_calendar_delete_calendar_events", map[string]string{
		"events[0][eventid]": eventID, "events[0][repeat]": "0",
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "eventid": eventID}}, nil
}

type calendarGetEventsStep struct {
	name       string
	moduleName string
}

func newCalendarGetEventsStep(name string, config map[string]any) (*calendarGetEventsStep, error) {
	return &calendarGetEventsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *calendarGetEventsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["events[courseids][0]"] = courseID
	}
	result, err := client.callToMap(ctx, "core_calendar_get_calendar_events", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	events, _ := result["events"].([]any)
	return &sdk.StepResult{Output: map[string]any{"events": events}}, nil
}

type calendarGetDayViewStep struct {
	name       string
	moduleName string
}

func newCalendarGetDayViewStep(name string, config map[string]any) (*calendarGetDayViewStep, error) {
	return &calendarGetDayViewStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *calendarGetDayViewStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	year := resolveValue("year", current, config)
	month := resolveValue("month", current, config)
	day := resolveValue("day", current, config)
	if year == "" || month == "" || day == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "year, month, and day are required"}}, nil
	}
	params := map[string]string{"year": year, "month": month, "day": day}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseid"] = courseID
	}
	result, err := client.callToMap(ctx, "core_calendar_get_calendar_day_view", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type calendarGetMonthlyViewStep struct {
	name       string
	moduleName string
}

func newCalendarGetMonthlyViewStep(name string, config map[string]any) (*calendarGetMonthlyViewStep, error) {
	return &calendarGetMonthlyViewStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *calendarGetMonthlyViewStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	year := resolveValue("year", current, config)
	month := resolveValue("month", current, config)
	if year == "" || month == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "year and month are required"}}, nil
	}
	params := map[string]string{"year": year, "month": month}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseid"] = courseID
	}
	result, err := client.callToMap(ctx, "core_calendar_get_calendar_monthly_view", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
