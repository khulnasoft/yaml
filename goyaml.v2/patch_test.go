package yaml_test

import (
	"testing"

	yaml "go.khulnasoft.com/yaml/goyaml.v2"
)

func TestIssue117(t *testing.T) {
	data := []byte(`
a:
<<:
-
?
-
`)

	x := map[string]interface{}{}
	err := yaml.Unmarshal([]byte(data), &x)
	if err == nil {
		t.Errorf("expected error, got none")
	}
}
