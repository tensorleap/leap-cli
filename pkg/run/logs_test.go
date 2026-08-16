package run

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestLogLineMarshalJSON(t *testing.T) {
	// A JSON engine log embeds verbatim: no second round of escaping, so the
	// traceback's own \" stays \" instead of becoming \\\".
	engineLog := `{"levelname": "ERROR", "message": "boom\n  File \"/app/x.py\"\n"}`
	encoded, err := json.Marshal([]LogLine{LogLine(engineLog)})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(encoded), `\\\"`) {
		t.Fatalf("engine log was double-escaped: %s", encoded)
	}
	var decoded []struct {
		LevelName string `json:"levelname"`
		Message   string `json:"message"`
	}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("not a nested object: %v (%s)", err, encoded)
	}
	if decoded[0].LevelName != "ERROR" || !strings.Contains(decoded[0].Message, `File "/app/x.py"`) {
		t.Fatalf("unexpected decode: %+v", decoded[0])
	}

	// A plain (non-JSON) line still round-trips as a string.
	for _, plain := range []string{`Traceback (most recent call last):`, `42`, `{"unterminated": `} {
		encoded, err := json.Marshal(LogLine(plain))
		if err != nil {
			t.Fatal(err)
		}
		var back string
		if err := json.Unmarshal(encoded, &back); err != nil {
			t.Fatalf("plain line %q did not encode as a string: %s", plain, encoded)
		}
		if back != plain {
			t.Fatalf("plain line %q round-tripped as %q", plain, back)
		}
	}
}
