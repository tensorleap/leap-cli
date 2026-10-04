package mcp

import "regexp"

var secretPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)\b(AKIA|ASIA)[A-Z0-9]{16}\b`),
	regexp.MustCompile(`(?i)(aws_secret_access_key|aws_session_token|secret[_-]?key|api[_-]?key|access[_-]?token|auth[_-]?secret|password|passwd|token)(["']?\s*[:=]\s*["']?)[^\s"',}]+`),
	regexp.MustCompile(`(?i)\b(bearer|kbearer)\s+[A-Za-z0-9._~+/=-]{16,}`),
	regexp.MustCompile(`\beyJ[A-Za-z0-9_-]{8,}\.[A-Za-z0-9_-]{8,}\.[A-Za-z0-9_-]{8,}\b`),
	regexp.MustCompile(`-----BEGIN [A-Z ]*PRIVATE KEY-----[\s\S]*?-----END [A-Z ]*PRIVATE KEY-----`),
	regexp.MustCompile(`(?i)(X-Amz-Credential|X-Amz-Signature|X-Amz-Security-Token)=[^&\s"]+`),
}

func Scrub(s string) string {
	for i, re := range secretPatterns {
		switch i {
		case 1:
			s = re.ReplaceAllString(s, "${1}${2}[redacted]")
		case 2:
			s = re.ReplaceAllString(s, "${1} [redacted]")
		case 5:
			s = re.ReplaceAllString(s, "${1}=[redacted]")
		default:
			s = re.ReplaceAllString(s, "[redacted]")
		}
	}
	return s
}
