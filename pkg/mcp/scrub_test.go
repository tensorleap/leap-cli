package mcp

import (
	"strings"
	"testing"
)

func TestScrubRemovesSecrets(t *testing.T) {
	secrets := []string{
		"AKIAABCDEFGHIJKLMNOP",
		"wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		"FwoGZXIvYXdzEJr//////////session-token-value",
		"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZXN0In0.c2lnbmF0dXJlc2lnbmF0dXJl",
		"sk-test-0123456789abcdef0123",
		"Zm9vYmFyYmF6c2lnbmF0dXJl",
	}
	log := strings.Join([]string{
		"export AWS_ACCESS_KEY_ID=" + secrets[0],
		`aws_secret_access_key = "` + secrets[1] + `"`,
		"AWS_SESSION_TOKEN: " + secrets[2],
		"Authorization: Bearer " + secrets[3],
		`{"api_key": "` + secrets[4] + `"}`,
		"GET /session/x.png?X-Amz-Signature=" + secrets[5] + "&X-Amz-Expires=3600",
		"ValueError: expected float32, got int64",
	}, "\n")
	got := Scrub(log)
	for _, s := range secrets {
		if strings.Contains(got, s) {
			t.Fatalf("secret %q survived scrubbing:\n%s", s, got)
		}
	}
	if !strings.Contains(got, "ValueError: expected float32, got int64") {
		t.Fatalf("ordinary log lines must be kept:\n%s", got)
	}
}

func TestMeasureBatchesSplitsSameField(t *testing.T) {
	b := measureBatches([]Measure{{"sample_id", "Count"}, {"metrics.loss", "Average"}, {"metrics.loss", "Max"}, {"metadata.x", "Average"}})
	if len(b) != 2 || len(b[0]) != 3 || b[1][0].Aggregation != "Max" {
		t.Fatalf("same-field measures must go to separate requests: %+v", b)
	}
}

func TestErrorLinesIgnoreRoutineShutdown(t *testing.T) {
	noise := "2026-09-27 13:13:12,815 INFO [FAILMODE] pika.connection: AMQP stack terminated, failed to connect, or aborted: error-arg=None"
	real := []string{"Traceback (most recent call last):", "ValueError: expected float32", "2026-09-27 ERROR engine: sample generator failed", "reason: OOMKilled"}
	if isErrorLine(noise) {
		t.Fatalf("routine INFO line flagged as an error: %s", noise)
	}
	for _, l := range real {
		if !isErrorLine(l) {
			t.Fatalf("missed error line: %s", l)
		}
	}
}

func TestScrubRemovesUserContactDetailsFromEngineLogs(t *testing.T) {
	line := `{"user.id": "6a9f", "user.email": "dana@acme.com", "user.full_name": "Dana Levi", "user.phone_number": "+972500000000", "team.name": "Vision"}`
	out := Scrub(line)
	for _, v := range []string{"dana@acme.com", "Dana Levi", "+972500000000", "Vision"} {
		if strings.Contains(out, v) {
			t.Fatalf("%q leaked: %s", v, out)
		}
	}
	if !strings.Contains(out, `"user.id": "6a9f"`) {
		t.Fatalf("ids are not personal data: %s", out)
	}
}
