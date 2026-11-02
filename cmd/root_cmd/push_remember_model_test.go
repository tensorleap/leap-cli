package root_cmd

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

// G90: the picker highlights the model of the last push that finished, never one that failed
func TestPushRemembersTheModelOnlyWhenThePushFinished(t *testing.T) {
	for _, c := range []struct {
		name    string
		jobErr  error
		wantErr bool
		want    string
	}{
		{name: "push finished", want: "model/resnet_ls.h5"},
		{name: "push failed", jobErr: errors.New("parsing model failed"), wantErr: true, want: "model/resnet.h5"},
	} {
		t.Run(c.name, func(t *testing.T) {
			t.Chdir(t.TempDir())
			for _, f := range []string{"model/resnet.h5", "model/resnet_ls.h5"} {
				if err := os.MkdirAll(filepath.Dir(f), 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(f, []byte("m"), 0o644); err != nil {
					t.Fatal(err)
				}
			}
			// an earlier push of resnet.h5 finished
			if err := workspace.SaveLastPushedModelPath(".", "model/resnet.h5"); err != nil {
				t.Fatal(err)
			}

			prev := waitForPushJob
			waitForPushJob = func(context.Context, string, string, string) error { return c.jobErr }
			t.Cleanup(func() { waitForPushJob = prev })

			s := &pushState{ctx: context.Background(), project: &tensorleapapi.Project{},
				inputs: &pushInputs{modelPath: "./model/resnet_ls.h5"}}
			if err := s.waitForModelPush("v1", "j1"); (err != nil) != c.wantErr {
				t.Fatalf("waitForModelPush error = %v, wantErr %v", err, c.wantErr)
			}
			if got := workspace.LastPushedModelPath("."); got != c.want {
				t.Fatalf("remembered %q, want %q", got, c.want)
			}
		})
	}
}
