package tools

import "context"

type Call struct {
	Name  string
	Input map[string]any
}

type Result struct {
	Output string
	Error  string
}

type Tool interface {
	Definition() Definition
	Execute(ctx context.Context, call Call) Result
}

type Executor struct {
	Registry *Registry
}

func (e Executor) Execute(ctx context.Context, call Call) Result {
	if e.Registry == nil {
		return Result{Error: "tool registry is unavailable"}
	}
	t, ok := e.Registry.Implementation(call.Name)
	if !ok {
		return Result{Error: "unknown tool: " + call.Name}
	}
	return t.Execute(ctx, call)
}
