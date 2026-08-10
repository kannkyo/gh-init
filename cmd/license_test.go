package cmd

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/cli/go-gh/v2/pkg/api"
)

func TestRenderLicenseBody(t *testing.T) {
	got := renderLicenseBody("Copyright (c) [year] [fullname]\nSome [year] text.", 2026, "Ada Lovelace")
	want := "Copyright (c) 2026 Ada Lovelace\nSome 2026 text."
	if got != want {
		t.Errorf("renderLicenseBody() = %q, want %q", got, want)
	}
}

func TestRunLicenseCreate_Success(t *testing.T) {
	t.Chdir(t.TempDir())

	client := &fakeRESTClient{
		getFunc: func(path string, response interface{}) error {
			if path != "licenses/mit" {
				t.Fatalf("unexpected path %q", path)
			}
			resp := response.(*licenseResponse)
			resp.Body = "Copyright (c) [year] [fullname]\n"
			return nil
		},
	}

	err := runLicenseCreate(client, "mit", func() (string, error) { return "Ada Lovelace", nil })
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := os.ReadFile("LICENSE")
	if err != nil {
		t.Fatalf("reading LICENSE: %v", err)
	}

	want := "Copyright (c) " + strconv.Itoa(time.Now().Year()) + " Ada Lovelace\n"
	if string(got) != want {
		t.Errorf("LICENSE content = %q, want %q", got, want)
	}
}

func TestRunLicenseCreate_ExistingFile(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := os.WriteFile("LICENSE", []byte("existing"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	client := &fakeRESTClient{
		getFunc: func(path string, response interface{}) error {
			t.Fatal("API should not be called when LICENSE already exists")
			return nil
		},
	}

	err := runLicenseCreate(client, "mit", func() (string, error) { return "Ada Lovelace", nil })
	want := "LICENSE already exists"
	if err == nil || err.Error() != want {
		t.Fatalf("err = %v, want %q", err, want)
	}
}

func TestRunLicenseCreate_NotFound(t *testing.T) {
	t.Chdir(t.TempDir())

	client := &fakeRESTClient{
		getFunc: func(path string, response interface{}) error {
			return &api.HTTPError{StatusCode: 404}
		},
	}

	err := runLicenseCreate(client, "does-not-exist", func() (string, error) { return "Ada Lovelace", nil })
	want := `license "does-not-exist" not found`
	if err == nil || err.Error() != want {
		t.Fatalf("err = %v, want %q", err, want)
	}
}
