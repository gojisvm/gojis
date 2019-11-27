package config

// Cfg is the config struct that the VM can use.
type Cfg struct {
	// Debug indicates whether the VM should print additional debug information.
	// This flag can have a negative impact on the VM's performance.
	Debug bool
}
