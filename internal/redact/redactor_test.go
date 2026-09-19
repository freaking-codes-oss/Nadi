package redact

import "testing"

func TestTextRedactsSecrets(t *testing.T) {
	got := Text("api_key=abc123 token:xyz password hunter2")
	want := "api_key=[REDACTED] token:[REDACTED] password [REDACTED]"
	if got != want {
		t.Fatalf("Text() = %q, want %q", got, want)
	}
}
