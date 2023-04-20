package checks

// Holds a bunch of Checks, and can dish them out based on string keys.
type CheckRegistry struct {
	entries map[string]Check
}

func NewCheckRegistry() *CheckRegistry {
	r := CheckRegistry{}
	r.entries = make(map[string]Check)
	return &r
}

func (c *CheckRegistry) Add(key string, check Check) {
	c.entries[key] = check
}

func (c *CheckRegistry) Get(key string) Check {
	result := c.entries[key]
	if result == nil {
		panic("get check with key '" + key + "': not found")
	}
	return result
}

func (c *CheckRegistry) AddAllChecks() {
	c.Add("http.ok", CheckHttpOk{})
	c.Add("env", CheckEnv{})
	c.Add("tcp.open", CheckTcpOpen{})
}
