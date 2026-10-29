package server

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	servercmd "github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/config"
)

type sentEvent struct {
	Event analytics.EventType
	Props map[string]interface{}
}

// runs install or reinstall with every step that touches the machine replaced, and returns the events
// it sent and the token it would apply to the new server
func runWithStubs(t *testing.T, newCmd func() *cobra.Command, failAt string, args ...string) ([]sentEvent, string) {
	var events []sentEvent
	var applied string
	oldSend, oldVersion, oldInit, oldInstall, oldReinstall, oldApply := sendEvent, installedVersion, initDataDir, runInstall, runReinstall, applyLicense
	t.Cleanup(func() {
		sendEvent, installedVersion, initDataDir, runInstall, runReinstall, applyLicense = oldSend, oldVersion, oldInit, oldInstall, oldReinstall, oldApply
	})
	sendEvent = func(e analytics.EventType, props map[string]interface{}) {
		events = append(events, sentEvent{e, props})
	}
	installedVersion = func(context.Context) string { return "" }
	initDataDir = func(_ context.Context, _ string) (bool, error) {
		if failAt == "init_data_dir" {
			return false, errors.New("data dir failed")
		}
		return false, nil
	}
	failRun := func() (*server.InstallationResult, error) {
		if failAt == "run" {
			return nil, errors.New("helm failed")
		}
		return &server.InstallationResult{}, nil
	}
	runInstall = func(*cobra.Command, *servercmd.InstallFlags) (*server.InstallationResult, error) { return failRun() }
	runReinstall = func(*cobra.Command, *servercmd.ReinstallFlags, bool) (*server.InstallationResult, error) {
		return failRun()
	}
	applyLicense = func(_ *cobra.Command, f *auth.LicenseFlag) error {
		if f.HasLicense() {
			applied, _ = f.PrepareLicenseToken()
		}
		return nil
	}

	cmd := newCmd()
	cmd.SetOut(io.Discard)
	cmd.SetErr(io.Discard)
	// --airgap skips the internet check; the file is never opened because the install itself is stubbed
	cmd.SetArgs(append([]string{"--airgap", "/nonexistent/airgap.tar", "--data-dir", "/tmp/tl-test"}, args...))
	err := cmd.Execute()
	if (err != nil) != (failAt != "") {
		t.Fatalf("failAt=%q, got err %v", failAt, err)
	}
	return events, applied
}

func TestSavedLicenseSendsTheSameEventsAsTheFlags(t *testing.T) {
	viper.Reset()
	t.Cleanup(viper.Reset)
	dir := t.TempDir()
	if err := config.InitConfig(filepath.Join(dir, "config.yaml")); err != nil {
		t.Fatal(err)
	}
	enc := base64.RawURLEncoding.EncodeToString
	token := enc([]byte(`{"alg":"RS256"}`)) + "." + enc([]byte(`{"activated":false,"user":{"email":"a@b.c"}}`)) + ".sig"
	file := filepath.Join(dir, "license.jwt")
	_ = os.WriteFile(file, []byte(token+"\n"), 0o600)

	commands := map[string]func() *cobra.Command{"install": NewInstallCmd, "reinstall": NewReinstallCmd}
	for name, newCmd := range commands {
		for _, failAt := range []string{"", "init_data_dir", "run"} {
			t.Run(fmt.Sprintf("%s/fail=%q", name, failAt), func(t *testing.T) {
				viper.Set("license.token", "")
				byToken, appliedToken := runWithStubs(t, newCmd, failAt, "--license-token", token)
				byFile, appliedFile := runWithStubs(t, newCmd, failAt, "--license-file", file)
				if err := auth.SaveLicense(token); err != nil {
					t.Fatal(err)
				}
				bySaved, appliedSaved := runWithStubs(t, newCmd, failAt)

				if len(byToken) < 2 {
					t.Fatalf("expected started and a result event, got %v", byToken)
				}
				for _, e := range byToken {
					if e.Props["email"] != "a@b.c" {
						t.Errorf("%s: email %v", e.Event, e.Props["email"])
					}
				}
				if !reflect.DeepEqual(byToken, byFile) || !reflect.DeepEqual(byToken, bySaved) {
					t.Fatalf("events differ:\n--license-token %v\n--license-file  %v\nsaved           %v", byToken, byFile, bySaved)
				}
				if strings.Contains(fmt.Sprint(bySaved), token) {
					t.Fatal("the token itself must never reach an event")
				}
				wantApplied := token
				if failAt != "" {
					wantApplied = ""
				}
				if appliedToken != wantApplied || appliedFile != wantApplied || appliedSaved != wantApplied {
					t.Fatalf("license applied after install: token=%v file=%v saved=%v", appliedToken != "", appliedFile != "", appliedSaved != "")
				}
				for i := range byToken {
					t.Logf("%-28s --license-token email=%v | saved email=%v", byToken[i].Event, byToken[i].Props["email"], bySaved[i].Props["email"])
				}
			})
		}
	}
}
