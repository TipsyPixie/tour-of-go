package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ipAddress IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipAddress[0], ipAddress[1], ipAddress[2], ipAddress[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

