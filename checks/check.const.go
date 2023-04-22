package checks

/*
check = "const"

Simply returns the information
given when created.
Useful for debug messages
and grouping other checks.

  - okay: A bool indicating success or failure.
  - label: A string containing the message for the user.
*/
type CheckConst struct {
	Okay  bool
	Label string
}

func (c CheckConst) Run(p Params) CheckResult {
	if c.Label == "" {
		// c is probably not initialized.
		// Get Okay and Label from params.
		c.Okay = p.GetBool("okay")
		c.Label = p.GetString("label")
	}
	return newBasicResult(c.Okay, c.Label)
}
