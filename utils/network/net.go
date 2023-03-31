package network

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

// checkURL check url isValid
func checkURL(host string) bool {
	addr := host
	if !strings.Contains(host, "http://") && !strings.Contains(host, "https://") {
		addr = fmt.Sprintf("https://%s", host)
	}
	_, err := url.ParseRequestURI(addr)
	return err == nil
}

// checkIP check IP address isValid
func checkIP(host string) bool {
	return net.ParseIP(host) != nil
}

// CheckHost check the host isValid
func CheckHost(addr string) bool {
	return checkURL(addr) || checkIP(addr)
}
