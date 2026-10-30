package mcp

import (
	"bytes"
	"encoding/base64"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/config"
)

// no test in this package sends a real Mixpanel event
func TestMain(m *testing.M) {
	sendEvent = func(analytics.EventType, map[string]interface{}) {}
	os.Exit(m.Run())
}

func TestConfigSendsOneEventWhetherItWorkedOrNot(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("HOME", dir)
	viper.Reset()
	t.Cleanup(viper.Reset)
	if err := config.InitConfig(filepath.Join(dir, "config.yaml")); err != nil {
		t.Fatal(err)
	}
	var events []map[string]interface{}
	t.Cleanup(func() { sendEvent = func(analytics.EventType, map[string]interface{}) {} })
	sendEvent = func(e analytics.EventType, props map[string]interface{}) {
		if e != analytics.EventMcpConfig {
			t.Errorf("event %s", e)
		}
		events = append(events, props)
	}
	enc := base64.RawURLEncoding.EncodeToString
	token := enc([]byte(`{}`)) + "." + enc([]byte(`{"activated":false,"user":{"email":"a+b@c.d"}}`)) + ".sig"
	run := func(args ...string) map[string]interface{} {
		t.Helper()
		events = nil
		cmd := newConfigCmd()
		cmd.SetOut(&bytes.Buffer{})
		cmd.SetArgs(args)
		_ = cmd.Execute()
		if len(events) != 1 {
			t.Fatalf("%v: want one event, got %v", args, events)
		}
		return events[0]
	}

	cases := []struct {
		args []string
		want map[string]interface{}
	}{
		{[]string{"codex", "--no-skills", "--license-token", token},
			map[string]interface{}{"assistant": "codex", "success": true, "license_saved": true, "print_only": false, "email": "a+b@c.d"}},
		{[]string{"cursor", "--no-skills"},
			map[string]interface{}{"assistant": "cursor", "success": true, "license_saved": false, "print_only": false}},
		{[]string{"codex", "--print", "--license-token", token},
			map[string]interface{}{"assistant": "codex", "success": true, "license_saved": false, "print_only": true, "email": "a+b@c.d"}},
		{[]string{"codex", "--license-token", "garbage"},
			map[string]interface{}{"success": false, "license_saved": false, "print_only": false}},
		{[]string{"no-such-assistant"},
			map[string]interface{}{"success": false, "license_saved": false, "print_only": false}},
	}
	for _, c := range cases {
		if got := run(c.args...); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%v:\ngot  %v\nwant %v", c.args, got, c.want)
		}
	}
	if viper.GetString("license.token") != token {
		t.Fatal("the license was saved by the first run")
	}
}
