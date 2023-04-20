package checks

// Implements something that can be checked in the local dev environment.
// Common examples include checking a URL for HTTP 200 OK, checking a local
// database service, and running scripts that do custom checks.
type Check interface {
	Run(options Params) CheckResult
}

// ------------

type Params map[string]any

func (p Params) GetString(key string) string {
	return p[key].(string)
}

func (p Params) GetStrings(key string) []string {
	result := []string{}
	given := p[key].([]any)
	for i := range given {
		result = append(result, given[i].(string))
	}
	return result
}

// -------------

// Reports what happened when a Check was run.
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
