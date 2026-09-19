package approval

import "testing"

func TestApprovalLifecycle(t *testing.T) {
	m := New()
	m.Add("request-1", Request{Tool: "shell.run", Summary: "run tests", Risk: "medium"})
	if r, ok := m.Get("request-1"); !ok || r.Status != Pending {
		t.Fatalf("new request = %#v, %v; want pending", r, ok)
	}
	if !m.Approve("request-1") || !m.ConsumeApproved("request-1") {
		t.Fatal("approved request was not consumed")
	}
	if _, ok := m.Get("request-1"); ok {
		t.Fatal("consumed request remains pending")
	}
}
