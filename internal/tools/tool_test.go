package tools

import (
	"context"
	"testing"

	"github.com/freaking-codes-oss/Nadi/internal/approval"
)

type testTool struct{}
func (testTool) Definition() Definition { return Definition{Name: "test.run"} }
func (testTool) Execute(context.Context, Call) Result { return Result{Output: "executed"} }

func TestExecutorRequiresApproval(t *testing.T) {
	registry := NewRegistry()
	registry.RegisterTool(testTool{})
	approvals := approval.New()
	executor := Executor{Registry: registry, Approvals: approvals}
	result := executor.Execute(context.Background(), Call{Name: "test.run"})
	if !result.ApprovalRequired || result.Error == "" {
		t.Fatalf("result = %#v, want approval requirement", result)
	}
}

func TestExecutorConsumesApproval(t *testing.T) {
	registry := NewRegistry()
	registry.RegisterTool(testTool{})
	approvals := approval.New()
	approvals.Add("a1", approval.Request{Tool: "test.run"})
	approvals.Approve("a1")
	executor := Executor{Registry: registry, Approvals: approvals}
	result := executor.Execute(context.Background(), Call{Name: "test.run", ApprovalID: "a1"})
	if result.Error != "" || result.Output != "executed" {
		t.Fatalf("result = %#v, want execution", result)
	}
}
