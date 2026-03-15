package main

import "strings"

func parseStatus(input string) (Status, bool) {

	formatted := strings.ToLower(strings.ReplaceAll(input, " ", "-"))

	switch formatted {
	case "todo":
		return Todo, true
	case "in-progress":
		return InProgress, true
	case "done":
		return Done, true
	default:
		return "", false
	}
}
