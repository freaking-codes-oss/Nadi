package policy

import "testing"

func TestEngineDecide(t *testing.T) {
	engine := Engine{}
	for _, name := range []string{"secret.read", "policy.update", "git.push", "file.delete"} {
		if got := engine.Decide(name); got != Deny {
			t.Fatalf("Decide(%q) = %v, want Deny", name, got)
		}
	}
	if got := engine.Decide("shell.run"); got != ApproveRequired {
		t.Fatalf("Decide(shell.run) = %v, want ApproveRequired", got)
	}
}
