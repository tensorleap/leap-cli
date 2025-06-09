package project

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func TestBuildProjectContext_NoBgImage(t *testing.T) {
	p := &tensorleapapi.Project{
		Cid:  "cid",
		Name: "TestProject",
	}

	_, err := BuildProjectContext(context.Background(), p, 1)
	assert.Error(t, err)
	assert.EqualError(t, err, "project TestProject has no background image configured")
}
