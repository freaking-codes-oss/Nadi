package tools

import (
	"context"
	"fmt"

	"github.com/freaking-codes-oss/Nadi/internal/approval"
	"github.com/freaking-codes-oss/Nadi/internal/policy"
)

type Call struct { Name string; Input map[string]any; ApprovalID string }
type Result struct { Output string; Error string; ApprovalID string; ApprovalRequired bool }
type Tool interface { Definition() Definition; Execute(context.Context, Call) Result }
type Executor struct { Registry *Registry; Policy policy.Engine; Approvals *approval.Manager }

func (e Executor) Execute(ctx context.Context, call Call) Result {
	if e.Registry == nil { return Result{Error:"tool registry is unavailable"} }
	t, ok := e.Registry.Implementation(call.Name); if !ok { return Result{Error:"unknown tool: "+call.Name} }
	decision := e.Policy.Decide(call.Name)
	if decision == policy.Deny { return Result{Error:"tool denied by policy: "+call.Name} }
	if decision == policy.ApproveRequired {
		if e.Approvals == nil { return Result{Error:"approval manager is unavailable"} }
		if call.ApprovalID == "" { return Result{Error:"approval required", ApprovalRequired:true} }
		if !e.Approvals.ConsumeApproved(call.ApprovalID) { return Result{Error:fmt.Sprintf("approval required: %s", call.ApprovalID), ApprovalID:call.ApprovalID, ApprovalRequired:true} }
	}
	return t.Execute(ctx, call)
}
