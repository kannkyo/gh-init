package cmd

import (
	"errors"
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

// restGetter is the subset of api.RESTClient used by init commands.
type restGetter interface {
	Get(path string, response interface{}) error
}

// newRESTClient constructs the REST client used by commands. Overridden in
// tests to inject a fake client.
var newRESTClient = func() (restGetter, error) {
	return api.DefaultRESTClient()
}

// formatAPIError converts a 404 response from the GitHub API into a
// user-facing "<kind> \"<name>\" not found" error. Other errors are
// returned unchanged.
func formatAPIError(err error, kind, name string) error {
	var httpErr *api.HTTPError
	if errors.As(err, &httpErr) && httpErr.StatusCode == 404 {
		return fmt.Errorf("%s %q not found", kind, name)
	}
	return err
}
