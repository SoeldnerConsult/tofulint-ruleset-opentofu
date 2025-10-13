package rules

import (
	"testing"

	"github.com/arsiba/tofulint-plugin-sdk/helper"
	"github.com/arsiba/tofulint-ruleset-opentofu/terraform"
)

func testRunner(t *testing.T, files map[string]string) *terraform.Runner {
	return terraform.NewRunner(helper.TestRunner(t, files))
}
