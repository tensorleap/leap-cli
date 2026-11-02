package local

import (
	"reflect"
	"testing"
)

func TestMoveToFront(t *testing.T) {
	files := []string{"./model/resnet.h5", "./model/resnet_ls.h5", "./other.onnx"}
	cases := []struct {
		name, preferred string
		want            []string
	}{
		{"remembered file comes first", "model/resnet_ls.h5", []string{"./model/resnet_ls.h5", "./model/resnet.h5", "./other.onnx"}},
		{"already first", "./model/resnet.h5", files},
		{"nothing remembered", "", files},
		{"remembered file not in the list", "model/gone.h5", files},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := moveToFront(files, c.preferred); !reflect.DeepEqual(got, c.want) {
				t.Fatalf("moveToFront(%q) = %v, want %v", c.preferred, got, c.want)
			}
		})
	}
}
