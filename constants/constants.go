package constants

func ClientActivityView() []string {
	return []string{"respond", "rejected", "mark_as_resolved"}
}

func AdminActivityView() []string {
	return []string{"respond", "rejected", "mark_as_resolved", "reviewer_reassigned", "escalated", "reviewer_assigned", "resolution_request", "esolution_rejected"}
}

func DateTimeFormat() string {
	return "2006-01-02"
}
