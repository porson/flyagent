// 2017_06_21 下午9:15
// Create by Porson

package agent

import "time"

// Messages sending model
type Sender struct {
	Commander []IpAddress
	Protocol  Protocol
	Gap       time.Duration
	Timeout   time.Duration
}
