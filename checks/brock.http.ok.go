package checks

import "github.com/tylermumford/friendly-broccoli/brock"

type BrockHttpOk struct {
}

var _ brock.Check = BrockHttpOk{}

func (b BrockHttpOk) Run(options map[string]any) (brock.CheckResult, error) {
	panic("unimplmemented")
}
