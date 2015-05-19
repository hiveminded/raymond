package raymond

import "testing"

//
// Those tests come from:
//   https://github.com/wycats/handlebars.js/blob/master/spec/builtin.js
//
var hbBuiltinsTests = []raymondTest{
	{
		"#if - if with boolean argument shows the contents when true",
		"{{#if goodbye}}GOODBYE {{/if}}cruel {{world}}!",
		map[string]interface{}{"goodbye": true, "world": "world"},
		nil,
		"GOODBYE cruel world!",
	},
	{
		"#if - if with string argument shows the contents",
		"{{#if goodbye}}GOODBYE {{/if}}cruel {{world}}!",
		map[string]interface{}{"goodbye": "dummy", "world": "world"},
		nil,
		"GOODBYE cruel world!",
	},
	{
		"#if - if with boolean argument does not show the contents when false",
		"{{#if goodbye}}GOODBYE {{/if}}cruel {{world}}!",
		map[string]interface{}{"goodbye": false, "world": "world"},
		nil,
		"cruel world!",
	},
	{
		"#if - if with undefined does not show the contents",
		"{{#if goodbye}}GOODBYE {{/if}}cruel {{world}}!",
		map[string]interface{}{"world": "world"},
		nil,
		"cruel world!",
	},
	{
		"#if - if with non-empty array shows the contents",
		"{{#if goodbye}}GOODBYE {{/if}}cruel {{world}}!",
		map[string]interface{}{"goodbye": []string{"foo"}, "world": "world"},
		nil,
		"GOODBYE cruel world!",
	},
	{
		"#if - if with empty array does not show the contents",
		"{{#if goodbye}}GOODBYE {{/if}}cruel {{world}}!",
		map[string]interface{}{"goodbye": []string{}, "world": "world"},
		nil,
		"cruel world!",
	},
	{
		"#if - if with zero does not show the contents",
		"{{#if goodbye}}GOODBYE {{/if}}cruel {{world}}!",
		map[string]interface{}{"goodbye": 0, "world": "world"},
		nil,
		"cruel world!",
	},
	{
		"#if - if with zero and includeZero option shows the contents'",
		"{{#if goodbye includeZero=true}}GOODBYE {{/if}}cruel {{world}}!",
		map[string]interface{}{"goodbye": 0, "world": "world"},
		nil,
		"GOODBYE cruel world!",
	},

	// {
	// 	"",
	// 	"",
	// 	map[string]interface{}{},
	// 	nil,
	// 	"",
	// },
}

func TestHandlebarsBuiltins(t *testing.T) {
	launchRaymondTests(t, hbBuiltinsTests)
}
