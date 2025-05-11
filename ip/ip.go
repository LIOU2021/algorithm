package ip

import (
	"net"
)

var privateCIDRs []*net.IPNet

func init() {
	// RFC1918 + 本地範圍 (IPv4 + IPv6)
	cidrStrings := []string{
		// IPv4 Private
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",

		// IPv4 Special Local
		"127.0.0.0/8",    // Loopback
		"169.254.0.0/16", // Link-local

		// IPv6 Loopback, Link-local, ULA
		"::1/128",   // Loopback
		"fe80::/10", // Link-local
		"fc00::/7",  // Unique local address (ULA)
	}

	for _, cidr := range cidrStrings {
		_, block, err := net.ParseCIDR(cidr)
		if err == nil {
			privateCIDRs = append(privateCIDRs, block)
		}
	}
}

// IsInternalIP returns true if ipStr is in a private/internal network range.
func IsInternalIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	for _, block := range privateCIDRs {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
