package cmd

import (
	"errors"
	"testing"

	"github.com/cli/go-gh/v2/pkg/api"
)

type fakeRESTClient struct {
	getFunc func(path string, response interface{}) error
}

func (f *fakeRESTClient) Get(path string, response interface{}) error {
	return f.getFunc(path, response)
}

func TestFormatAPIError_NotFound(t *testing.T) {
	err := formatAPIError(&api.HTTPError{StatusCode: 404}, "license", "mit")
	want := `license "mit" not found`
	if err == nil || err.Error() != want {
		t.Errorf("formatAPIError() = %v, want %q", err, want)
	}
}

func TestFormatAPIError_OtherError(t *testing.T) {
	original := errors.New("network unreachable")
	err := formatAPIError(original, "license", "mit")
	if !errors.Is(err, original) {
		t.Errorf("formatAPIError() = %v, want wrapped %v", err, original)
	}
}
