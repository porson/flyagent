// 2017_06_21 下午9:12
// Create by Porson

package agent

import (
	"strings"
	"errors"
	"strconv"
	"net"
)

const (
	HTTP  = Protocol("HTTP")
	HTTPS = Protocol("HTTPS")
)

type Protocol string
type IpAddress string

func NewIpAddress(ip string, port int) (IpAddress, error) {
	str := IpAddress(ip + ":" + string(port))
	err := str.IpCheck()
	if err != nil {
		str = ""
	}
	return str, err
}

func (ip IpAddress) IpCheck() error {
	if count := strings.Count(string(ip), ":"); count != 1 {
		return errors.New("Too many ':' in this IpAddress")
	}
	str := strings.Split(string(ip), ":")
	port, err := strconv.Atoi(str[1])
	if net.ParseIP(str[0]) == nil {
		return errors.New("Wrong IP address :" + str[0])
	} else {
		if err != nil {
			return errors.New("Wrong port ：" + str[1])
		} else {
			if port > 65535 || port < 0 {
				return errors.New("Port out of range (1-65535)：" + string(port))
			} else {
				return nil
			}
		}
	}
}
