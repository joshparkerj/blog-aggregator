package main

type TestPayload struct {
	Prompt      string `json:"prompt"`
	Suggestions string `json:"suggestion"`
}

var GlobalTestPayload TestPayload = TestPayload{
	Prompt:      "smile",
	Suggestions: "say cheese!",
}
