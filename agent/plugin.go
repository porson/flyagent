// 2017_06_21 下午9:16
// Create by Porson

package agent

type FunctionArgs []interface{}
type Function map[string]FunctionArgs

// Plugins
type Plugin struct {
	Name        string
	Description string
	Functions   []Function
	Switch      bool
	Result      map[string]interface{}
}

func NewPlugin() *Plugin {
	return new(Plugin)
}
