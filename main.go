package main

import (
	"github.com/arsiba/tofulint-plugin-sdk/plugin"
	"github.com/arsiba/tofulint-plugin-sdk/tflint"
	"github.com/arsiba/tofulint-ruleset-opentofu/project"
	"github.com/arsiba/tofulint-ruleset-opentofu/rules"
	"github.com/arsiba/tofulint-ruleset-opentofu/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &terraform.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:       "opentofu",
				Version:    project.Version,
				Constraint: ">= 0.0.1",
			},
			PresetRules: rules.PresetRules,
		},
	})
}
