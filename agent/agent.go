package agent

// Agent main struct
type FlyAgent struct {
	Auth     Auth
	Sender   Sender
	Listener Listener
	Plugins  []Plugin
}
