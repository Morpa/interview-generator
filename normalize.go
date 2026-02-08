package main

import "strings"

func normalizeQuestion(q string) string {
	q = strings.TrimSpace(q)
	q = strings.TrimSuffix(q, "?")
	return "What is " + q + "?"
}

func normalizeExplanation(text string) string {
	replacements := map[string]string{
		"javascript": "JavaScript",
		"js":         "JavaScript",
		"async":      "asynchronous",
	}

	for k, v := range replacements {
		text = strings.ReplaceAll(text, k, v)
	}

	return strings.TrimSpace(text)
}

func buildSeniorExplanation(question string) string {
	return "In JavaScript, " + question +
		" This concept is important because it directly impacts application correctness, performance, and maintainability. " +
		"In real-world scenarios, it is commonly used to manage state, control side effects, or handle asynchronous operations. " +
		"A common mistake is misunderstanding how it behaves under different execution contexts."
}

func isEliminatory(q string) bool {
	keywords := []string{
		"closure",
		"event loop",
		"promise",
		"async",
		"hook",
		"useeffect",
		"state",
	}

	q = strings.ToLower(q)
	for _, k := range keywords {
		if strings.Contains(q, k) {
			return true
		}
	}
	return false
}
