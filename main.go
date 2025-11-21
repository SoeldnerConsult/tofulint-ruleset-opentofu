package main

import (
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/plugin"
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/tflint"
	"github.com/SoeldnerConsult/tofulint-ruleset-opentofu/project"
	"github.com/SoeldnerConsult/tofulint-ruleset-opentofu/rules"
	"github.com/SoeldnerConsult/tofulint-ruleset-opentofu/terraform"
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
