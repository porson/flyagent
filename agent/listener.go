// 2017_06_21 下午9:15
// Create by Porson

package agent

import "time"

// Command listener
type Listener struct {
	IpAddress []IpAddress
	Gap       time.Duration
}