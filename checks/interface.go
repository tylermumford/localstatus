/*
Package checks contains the individual checks (actions)
which can be performed by LocalStatus.
Each check implements the [Check] interface,
and each check is referenced in the [CheckRegistry] type.

(I plan to add support for custom script-based checks
once I figure out how I want to deal with shells across platforms.)
*/
package checks

// Implements something that can be checked in the local dev environment.
// Common examples include checking a URL for HTTP 200 OK, checking a local
// database service, and running scripts that do custom checks.
type Check interface {
	Run(options Params) CheckResult
}

// ------------

// Params hold(s) information needed
// to customize a Check's behavior.
// The Params come straight from the parsed TOML,
// and they are typed as interface{}.
//
// Helper methods such as [GetString]
// are a convenient way to get the data
// into more useful types.
type Params map[string]any

// Returns an empty string if key not found.
func (p Params) GetString(key string) string {
	if p[key] == nil {
		return ""
	}
	return p[key].(string)
}

// Returns an empty slice if key is not found.
func (p Params) GetStrings(key string) []string {
	result := []string{}

	if p[key] == nil {
		return result
	}

	given := p[key].([]any)
	for i := range given {
		result = append(result, given[i].(string))
	}
	return result
}

func (p Params) GetBool(key string) bool {
	return p[key].(bool)
}

// ------------

// A CheckResult reports what happened when a Check was run.
type CheckResult interface {
	IsOkay() bool
	Label() string
}

type basicResult struct {
	okay  bool
	label string
}

func (b basicResult) IsOkay() bool {
	return b.okay
}

func (b basicResult) Label() string {
	return b.label
}

func newBasicResult(okay bool, label string) CheckResult {
	return basicResult{okay: okay, label: label}
}
