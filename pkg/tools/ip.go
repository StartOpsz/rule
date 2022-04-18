package tools

import (
	"fmt"
	"net"
)

func CIDRToIpMask() {
	ip, ipNet, err := net.ParseCIDR("192.168.1.1/24")
	if err != nil {
		return
	}

	fmt.Println("ip: ", ip)

	fmt.Println("ipNet: ", ipNet)
	return
}

func ParseIp() {
	ip := net.ParseIP("192.168.1.1.1")
	fmt.Println("ip: ", ip)
}

func ParseMask() {
	hw, err := net.ParseMAC("ff:ff:ff:ff:ff:ff")
	if err != nil {

		return
	}

	fmt.Println("hw: ", hw)
}
