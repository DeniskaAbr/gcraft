package gc_config

type Configuration struct {
	FullScreen      bool
	RunInBackground bool
	VsyncEnabled    bool
	TicksPerSecond  int
	path            string
}

func (c *Configuration) SetPath(p string) {
	c.path = p
}
