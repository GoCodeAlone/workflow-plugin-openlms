package internal

import (
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// stepConstructor is a function that creates a StepInstance.
type stepConstructor func(name string, config map[string]any) (sdk.StepInstance, error)

// stepRegistry maps step type strings to constructor functions.
var stepRegistry = map[string]stepConstructor{
	// Users
	"step.openlms_user_create":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newUserCreateStep(n, c) },
	"step.openlms_user_update":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newUserUpdateStep(n, c) },
	"step.openlms_user_delete":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newUserDeleteStep(n, c) },
	"step.openlms_user_get":          func(n string, c map[string]any) (sdk.StepInstance, error) { return newUserGetStep(n, c) },
	"step.openlms_user_get_by_field": func(n string, c map[string]any) (sdk.StepInstance, error) { return newUserGetByFieldStep(n, c) },
	"step.openlms_user_search":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newUserSearchStep(n, c) },

	// Courses
	"step.openlms_course_create":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseCreateStep(n, c) },
	"step.openlms_course_update":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseUpdateStep(n, c) },
	"step.openlms_course_delete":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseDeleteStep(n, c) },
	"step.openlms_course_get":               func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseGetStep(n, c) },
	"step.openlms_course_get_by_field":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseGetByFieldStep(n, c) },
	"step.openlms_course_search":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseSearchStep(n, c) },
	"step.openlms_course_get_contents":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseGetContentsStep(n, c) },
	"step.openlms_course_get_categories":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseGetCategoriesStep(n, c) },
	"step.openlms_course_create_categories": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseCreateCategoriesStep(n, c) },
	"step.openlms_course_delete_categories": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseDeleteCategoriesStep(n, c) },
	"step.openlms_course_duplicate":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newCourseDuplicateStep(n, c) },

	// Enrollments
	"step.openlms_enrol_get_enrolled_users": func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnrolGetEnrolledUsersStep(n, c) },
	"step.openlms_enrol_get_user_courses":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnrolGetUserCoursesStep(n, c) },
	"step.openlms_enrol_manual_enrol":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnrolManualEnrolStep(n, c) },
	"step.openlms_enrol_manual_unenrol":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnrolManualUnenrolStep(n, c) },
	"step.openlms_enrol_self_enrol":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnrolSelfEnrolStep(n, c) },
	"step.openlms_enrol_get_course_methods": func(n string, c map[string]any) (sdk.StepInstance, error) { return newEnrolGetCourseMethodsStep(n, c) },

	// Grades
	"step.openlms_grade_get_grades":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newGradeGetGradesStep(n, c) },
	"step.openlms_grade_update_grades":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newGradeUpdateGradesStep(n, c) },
	"step.openlms_grade_get_grade_items":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newGradeGetGradeItemsStep(n, c) },
	"step.openlms_grade_get_grades_table": func(n string, c map[string]any) (sdk.StepInstance, error) { return newGradeGetGradesTableStep(n, c) },

	// Assignments
	"step.openlms_assign_get_assignments":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newAssignGetAssignmentsStep(n, c) },
	"step.openlms_assign_get_submissions":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newAssignGetSubmissionsStep(n, c) },
	"step.openlms_assign_get_grades":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newAssignGetGradesStep(n, c) },
	"step.openlms_assign_save_submission":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newAssignSaveSubmissionStep(n, c) },
	"step.openlms_assign_submit_for_grading": func(n string, c map[string]any) (sdk.StepInstance, error) { return newAssignSubmitForGradingStep(n, c) },
	"step.openlms_assign_save_grade":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newAssignSaveGradeStep(n, c) },

	// Quizzes
	"step.openlms_quiz_get_by_course":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newQuizGetByCourseStep(n, c) },
	"step.openlms_quiz_get_attempts":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newQuizGetAttemptsStep(n, c) },
	"step.openlms_quiz_get_attempt_data":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newQuizGetAttemptDataStep(n, c) },
	"step.openlms_quiz_get_attempt_review": func(n string, c map[string]any) (sdk.StepInstance, error) { return newQuizGetAttemptReviewStep(n, c) },
	"step.openlms_quiz_start_attempt":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newQuizStartAttemptStep(n, c) },
	"step.openlms_quiz_save_attempt":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newQuizSaveAttemptStep(n, c) },
	"step.openlms_quiz_process_attempt":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newQuizProcessAttemptStep(n, c) },

	// Forums
	"step.openlms_forum_get_by_course":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newForumGetByCourseStep(n, c) },
	"step.openlms_forum_get_discussions": func(n string, c map[string]any) (sdk.StepInstance, error) { return newForumGetDiscussionsStep(n, c) },
	"step.openlms_forum_get_posts":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newForumGetPostsStep(n, c) },
	"step.openlms_forum_add_discussion":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newForumAddDiscussionStep(n, c) },
	"step.openlms_forum_add_post":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newForumAddPostStep(n, c) },
	"step.openlms_forum_delete_post":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newForumDeletePostStep(n, c) },

	// Groups
	"step.openlms_group_create":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newGroupCreateStep(n, c) },
	"step.openlms_group_delete":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newGroupDeleteStep(n, c) },
	"step.openlms_group_get_course_groups": func(n string, c map[string]any) (sdk.StepInstance, error) { return newGroupGetCourseGroupsStep(n, c) },
	"step.openlms_group_get_members":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newGroupGetMembersStep(n, c) },
	"step.openlms_group_add_members":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newGroupAddMembersStep(n, c) },
	"step.openlms_group_delete_members":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newGroupDeleteMembersStep(n, c) },

	// Messages
	"step.openlms_message_send":             func(n string, c map[string]any) (sdk.StepInstance, error) { return newMessageSendStep(n, c) },
	"step.openlms_message_get_messages":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newMessageGetMessagesStep(n, c) },
	"step.openlms_message_get_conversations": func(n string, c map[string]any) (sdk.StepInstance, error) { return newMessageGetConversationsStep(n, c) },
	"step.openlms_message_get_unread_count": func(n string, c map[string]any) (sdk.StepInstance, error) { return newMessageGetUnreadCountStep(n, c) },
	"step.openlms_message_mark_read":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newMessageMarkReadStep(n, c) },
	"step.openlms_message_block_user":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newMessageBlockUserStep(n, c) },
	"step.openlms_message_unblock_user":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newMessageUnblockUserStep(n, c) },

	// Calendar
	"step.openlms_calendar_create_events":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newCalendarCreateEventsStep(n, c) },
	"step.openlms_calendar_delete_events":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newCalendarDeleteEventsStep(n, c) },
	"step.openlms_calendar_get_events":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newCalendarGetEventsStep(n, c) },
	"step.openlms_calendar_get_day_view":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newCalendarGetDayViewStep(n, c) },
	"step.openlms_calendar_get_monthly_view": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCalendarGetMonthlyViewStep(n, c) },

	// Competencies
	"step.openlms_competency_create":           func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompetencyCreateStep(n, c) },
	"step.openlms_competency_list":             func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompetencyListStep(n, c) },
	"step.openlms_competency_delete":           func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompetencyDeleteStep(n, c) },
	"step.openlms_competency_create_framework": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompetencyCreateFrameworkStep(n, c) },
	"step.openlms_competency_list_frameworks":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompetencyListFrameworksStep(n, c) },
	"step.openlms_competency_create_plan":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompetencyCreatePlanStep(n, c) },
	"step.openlms_competency_list_plans":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompetencyListPlansStep(n, c) },
	"step.openlms_competency_add_to_course":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompetencyAddToCourseStep(n, c) },
	"step.openlms_competency_grade":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompetencyGradeStep(n, c) },

	// Completion
	"step.openlms_completion_get_activities_status": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompletionGetActivitiesStatusStep(n, c) },
	"step.openlms_completion_get_course_status":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompletionGetCourseStatusStep(n, c) },
	"step.openlms_completion_update_activity":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompletionUpdateActivityStep(n, c) },
	"step.openlms_completion_mark_self_completed":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newCompletionMarkSelfCompletedStep(n, c) },

	// Files
	"step.openlms_file_get_files": func(n string, c map[string]any) (sdk.StepInstance, error) { return newFileGetFilesStep(n, c) },
	"step.openlms_file_upload":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newFileUploadStep(n, c) },

	// Badges
	"step.openlms_badge_get_user_badges": func(n string, c map[string]any) (sdk.StepInstance, error) { return newBadgeGetUserBadgesStep(n, c) },

	// Cohorts
	"step.openlms_cohort_create":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newCohortCreateStep(n, c) },
	"step.openlms_cohort_delete":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newCohortDeleteStep(n, c) },
	"step.openlms_cohort_get":            func(n string, c map[string]any) (sdk.StepInstance, error) { return newCohortGetStep(n, c) },
	"step.openlms_cohort_search":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newCohortSearchStep(n, c) },
	"step.openlms_cohort_add_members":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newCohortAddMembersStep(n, c) },
	"step.openlms_cohort_delete_members": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCohortDeleteMembersStep(n, c) },

	// Roles
	"step.openlms_role_assign":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newRoleAssignStep(n, c) },
	"step.openlms_role_unassign": func(n string, c map[string]any) (sdk.StepInstance, error) { return newRoleUnassignStep(n, c) },

	// Notes
	"step.openlms_note_create": func(n string, c map[string]any) (sdk.StepInstance, error) { return newNoteCreateStep(n, c) },
	"step.openlms_note_get":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newNoteGetStep(n, c) },
	"step.openlms_note_delete": func(n string, c map[string]any) (sdk.StepInstance, error) { return newNoteDeleteStep(n, c) },

	// SCORM
	"step.openlms_scorm_get_by_course":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newScormGetByCourseStep(n, c) },
	"step.openlms_scorm_get_attempt_count": func(n string, c map[string]any) (sdk.StepInstance, error) { return newScormGetAttemptCountStep(n, c) },
	"step.openlms_scorm_get_scos":         func(n string, c map[string]any) (sdk.StepInstance, error) { return newScormGetScosStep(n, c) },
	"step.openlms_scorm_get_user_data":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newScormGetUserDataStep(n, c) },
	"step.openlms_scorm_insert_tracks":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newScormInsertTracksStep(n, c) },
	"step.openlms_scorm_launch_sco":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newScormLaunchScoStep(n, c) },

	// H5P
	"step.openlms_h5p_get_by_course": func(n string, c map[string]any) (sdk.StepInstance, error) { return newH5PGetByCourseStep(n, c) },
	"step.openlms_h5p_get_attempts":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newH5PGetAttemptsStep(n, c) },
	"step.openlms_h5p_get_results":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newH5PGetResultsStep(n, c) },

	// Reports
	"step.openlms_reportbuilder_list":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newReportBuilderListStep(n, c) },
	"step.openlms_reportbuilder_get":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newReportBuilderGetStep(n, c) },
	"step.openlms_reportbuilder_retrieve": func(n string, c map[string]any) (sdk.StepInstance, error) { return newReportBuilderRetrieveStep(n, c) },

	// Site Info
	"step.openlms_site_get_info":          func(n string, c map[string]any) (sdk.StepInstance, error) { return newSiteGetInfoStep(n, c) },
	"step.openlms_webservice_get_site_info": func(n string, c map[string]any) (sdk.StepInstance, error) { return newWebserviceGetSiteInfoStep(n, c) },

	// Lessons
	"step.openlms_lesson_get_by_course":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newLessonGetByCourseStep(n, c) },
	"step.openlms_lesson_get_pages":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newLessonGetPagesStep(n, c) },
	"step.openlms_lesson_get_page_data":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newLessonGetPageDataStep(n, c) },
	"step.openlms_lesson_launch_attempt":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newLessonLaunchAttemptStep(n, c) },
	"step.openlms_lesson_process_page":    func(n string, c map[string]any) (sdk.StepInstance, error) { return newLessonProcessPageStep(n, c) },
	"step.openlms_lesson_finish_attempt":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newLessonFinishAttemptStep(n, c) },

	// Glossary
	"step.openlms_glossary_get_by_course": func(n string, c map[string]any) (sdk.StepInstance, error) { return newGlossaryGetByCourseStep(n, c) },
	"step.openlms_glossary_get_entries":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newGlossaryGetEntriesStep(n, c) },
	"step.openlms_glossary_add_entry":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newGlossaryAddEntryStep(n, c) },
	"step.openlms_glossary_delete_entry":  func(n string, c map[string]any) (sdk.StepInstance, error) { return newGlossaryDeleteEntryStep(n, c) },

	// Search
	"step.openlms_search_get_results": func(n string, c map[string]any) (sdk.StepInstance, error) { return newSearchGetResultsStep(n, c) },

	// Tags
	"step.openlms_tag_get_tags": func(n string, c map[string]any) (sdk.StepInstance, error) { return newTagGetTagsStep(n, c) },
	"step.openlms_tag_update":   func(n string, c map[string]any) (sdk.StepInstance, error) { return newTagUpdateStep(n, c) },

	// LTI
	"step.openlms_lti_get_by_course":        func(n string, c map[string]any) (sdk.StepInstance, error) { return newLtiGetByCourseStep(n, c) },
	"step.openlms_lti_get_tool_launch_data": func(n string, c map[string]any) (sdk.StepInstance, error) { return newLtiGetToolLaunchDataStep(n, c) },
	"step.openlms_lti_get_tool_types":       func(n string, c map[string]any) (sdk.StepInstance, error) { return newLtiGetToolTypesStep(n, c) },

	// xAPI
	"step.openlms_xapi_statement_post": func(n string, c map[string]any) (sdk.StepInstance, error) { return newXAPIStatementPostStep(n, c) },
	"step.openlms_xapi_get_state":      func(n string, c map[string]any) (sdk.StepInstance, error) { return newXAPIGetStateStep(n, c) },
	"step.openlms_xapi_post_state":     func(n string, c map[string]any) (sdk.StepInstance, error) { return newXAPIPostStateStep(n, c) },

	// Generic
	"step.openlms_call_function": func(n string, c map[string]any) (sdk.StepInstance, error) { return newCallFunctionStep(n, c) },
}

// createStep dispatches to the appropriate step constructor.
func createStep(typeName, name string, config map[string]any) (sdk.StepInstance, error) {
	constructor, ok := stepRegistry[typeName]
	if !ok {
		return nil, fmt.Errorf("openlms plugin: unknown step type %q", typeName)
	}
	return constructor(name, config)
}

// allStepTypes returns all registered step type strings.
func allStepTypes() []string {
	types := make([]string, 0, len(stepRegistry))
	for k := range stepRegistry {
		types = append(types, k)
	}
	return types
}
