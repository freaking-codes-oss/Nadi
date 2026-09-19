package shell

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/freaking-codes-oss/Nadi/internal/tools"
)

type Tool struct { Timeout time.Duration; Allowed map[string]bool }
func (t Tool) Definition() tools.Definition { return tools.Definition{Name:"shell.run", Description:"Run an explicitly approved command in a controlled working directory", InputSchema: map[string]any{"type":"object"}} }
func (t Tool) Execute(ctx context.Context, call tools.Call) tools.Result {
	command, _ := call.Input["command"].(string)
	if strings.TrimSpace(command) == "" { return tools.Result{Error:"command is required"} }
	if strings.ContainsAny(command, ";&|<>`\n") { return tools.Result{Error:"shell metacharacters are not allowed"} }
	name := strings.Fields(command)[0]
	if len(t.Allowed) > 0 && !t.Allowed[name] { return tools.Result{Error:fmt.Sprintf("command not allowed: %s", name)} }
	timeout := t.Timeout; if timeout <= 0 { timeout = 30*time.Second }
	runCtx, cancel := context.WithTimeout(ctx, timeout); defer cancel()
	args := strings.Fields(command)[1:]
	out, err := exec.CommandContext(runCtx, name, args...).CombinedOutput()
	if runCtx.Err() != nil { return tools.Result{Error:"command timed out or was cancelled"} }
	if err != nil { return tools.Result{Output:string(out), Error:err.Error()} }
	return tools.Result{Output:string(out)}
}
