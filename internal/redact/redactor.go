package redact

import "regexp"

var secret = regexp.MustCompile(`(?i)(api[_-]?key|token|password|secret)([=: ]+)[^\s]+`)

func Text(s string) string { return secret.ReplaceAllString(s, `$1$2[REDACTED]`) }
