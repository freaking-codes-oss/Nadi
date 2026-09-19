package policy

import "strings"

type Decision int
const ( Deny Decision = iota; ApproveRequired; Allow )
type Engine struct{}
func (Engine) Decide(tool string) Decision {
	if strings.Contains(tool, "secret") || strings.Contains(tool, "policy") || strings.Contains(tool, "push") || strings.Contains(tool, "delete") { return Deny }
	return ApproveRequired
}
