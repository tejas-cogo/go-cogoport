package constants

func ClientActivityView() []string {
	return []string{"respond", "rejected", "mark_as_resolved"}
}

func AdminActivityView() []string {
	return []string{"respond", "rejected", "mark_as_resolved","reassigned","escalated","assigned"}
}
