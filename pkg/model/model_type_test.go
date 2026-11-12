package model

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func writeArchive(t *testing.T, name string, gz bool, members ...string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), name)
	f, err := os.Create(p)
	assert.NoError(t, err)
	var w io.Writer = f
	var gzw *gzip.Writer
	if gz {
		gzw = gzip.NewWriter(f)
		w = gzw
	}
	tw := tar.NewWriter(w)
	for _, m := range members {
		body := []byte("x")
		assert.NoError(t, tw.WriteHeader(&tar.Header{Name: m, Mode: 0o644, Size: int64(len(body))}))
		_, err := tw.Write(body)
		assert.NoError(t, err)
	}
	assert.NoError(t, tw.Close())
	if gzw != nil {
		assert.NoError(t, gzw.Close())
	}
	assert.NoError(t, f.Close())
	return p
}

var (
	onnxType = string(tensorleapapi.IMPORTMODELTYPE_ONNX)
	pbType   = string(tensorleapapi.IMPORTMODELTYPE_PB_TF2)
)

func TestDetectModelType_PlainExtensions(t *testing.T) {
	assert.Equal(t, onnxType, DetectModelType("m/model.ONNX"))
	assert.Equal(t, string(tensorleapapi.IMPORTMODELTYPE_H5_TF2), DetectModelType("model.h5"))
	assert.Equal(t, string(tensorleapapi.IMPORTMODELTYPE_JSON_TF2), DetectModelType("model.json"))
	assert.Equal(t, "", DetectModelType("model.pt2"))
}

func TestDetectModelType_OnnxArchives(t *testing.T) {
	for _, tc := range []struct {
		name string
		gz   bool
	}{{"model.tar", false}, {"model.tar.gz", true}, {"model.tgz", true}} {
		p := writeArchive(t, tc.name, tc.gz, "._model.onnx", "model.onnx", "model.onnx.data")
		assert.Equal(t, onnxType, DetectModelType(p), tc.name)
	}
}

func TestDetectModelType_OnnxInSubdirAfterData(t *testing.T) {
	p := writeArchive(t, "model.tar.gz", true, "big/model.onnx.data", "big/model.onnx")
	assert.Equal(t, onnxType, DetectModelType(p))
}

func TestDetectModelType_SavedModelArchive(t *testing.T) {
	p := writeArchive(t, "model.tar.gz", true, "export/saved_model.pb", "export/variables/variables.index")
	assert.Equal(t, pbType, DetectModelType(p))
}

func TestDetectModelType_UnrecognisedArchives(t *testing.T) {
	// .tar.gz keeps its historical SavedModel meaning; a plain .tar is left to the prompt
	assert.Equal(t, pbType, DetectModelType(writeArchive(t, "x.tar.gz", true, "readme.txt")))
	assert.Equal(t, "", DetectModelType(writeArchive(t, "x.tar", false, "readme.txt")))
	assert.Equal(t, pbType, DetectModelType(filepath.Join(t.TempDir(), "missing.tar.gz")))
}

func TestSelectModelType_ExplicitTypeWins(t *testing.T) {
	p := writeArchive(t, "model.tar", false, "model.onnx")
	mt := pbType
	assert.NoError(t, SelectModelType(&mt, p))
	assert.Equal(t, pbType, mt)

	mt = ""
	assert.NoError(t, SelectModelType(&mt, p))
	assert.Equal(t, onnxType, mt)
}
