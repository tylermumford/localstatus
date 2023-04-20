package checks

// Implements something that can be checked in the local dev environment.
// Common examples include checking a URL for HTTP 200 OK, checking a local
// database service, and running scripts that do custom checks.
type Check interface {
	Run(options map[string]any) (CheckResult, error)
}

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
