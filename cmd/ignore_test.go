package cmd

import (
	"os"
	"reflect"
	"testing"

	"github.com/cli/go-gh/v2/pkg/api"
)

func TestFetchIgnoreTemplateNames(t *testing.T) {
	client := &fakeRESTClient{
		getFunc: func(path string, response interface{}) error {
			if path != "gitignore/templates" {
				t.Fatalf("unexpected path %q", path)
			}
			names := response.(*[]string)
			*names = []string{"Go", "Node", "Python"}
			return nil
		},
	}

	got, err := fetchIgnoreTemplateNames(client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"Go", "Node", "Python"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("fetchIgnoreTemplateNames() = %v, want %v", got, want)
	}
}

func TestFetchIgnoreTemplateSource_NotFound(t *testing.T) {
	client := &fakeRESTClient{
		getFunc: func(path string, response interface{}) error {
			return &api.HTTPError{StatusCode: 404}
		},
	}

	_, err := fetchIgnoreTemplateSource(client, "does-not-exist")
	want := `template "does-not-exist" not found`
	if err == nil || err.Error() != want {
		t.Fatalf("err = %v, want %q", err, want)
	}
}

func TestRunIgnoreCreate_Success(t *testing.T) {
	t.Chdir(t.TempDir())

	client := &fakeRESTClient{
		getFunc: func(path string, response interface{}) error {
			if path != "gitignore/templates/Go" {
				t.Fatalf("unexpected path %q", path)
			}
			resp := response.(*ignoreTemplateResponse)
			resp.Source = "*.o\n*.exe\n"
			return nil
		},
	}

	if err := runIgnoreCreate(client, "Go"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := os.ReadFile(".gitignore")
	if err != nil {
		t.Fatalf("reading .gitignore: %v", err)
	}
	want := "*.o\n*.exe\n"
	if string(got) != want {
		t.Errorf(".gitignore content = %q, want %q", got, want)
	}
}

func TestRunIgnoreCreate_ExistingFile(t *testing.T) {
	t.Chdir(t.TempDir())
	if err := os.WriteFile(".gitignore", []byte("existing"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}

	client := &fakeRESTClient{
		getFunc: func(path string, response interface{}) error {
			t.Fatal("API should not be called when .gitignore already exists")
			return nil
		},
	}

	err := runIgnoreCreate(client, "Go")
	want := ".gitignore already exists"
	if err == nil || err.Error() != want {
		t.Fatalf("err = %v, want %q", err, want)
	}
}
