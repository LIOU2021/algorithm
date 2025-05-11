package ip

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsInternalIP(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"192.168.1.1", true},           // Private IPv4
		{"10.0.0.1", true},              // Private IPv4
		{"172.16.0.1", true},            // Private IPv4
		{"8.8.8.8", false},              // Public IPv4
		{"127.0.0.1", true},             // Loopback IPv4
		{"::1", true},                   // Loopback IPv6
		{"fd00::1", true},               // Unique local IPv6
		{"2001:4860:4860::8888", false}, // Public IPv6
		{"162.120.184.43", false},       // Public IPv4
	}

	for _, test := range tests {
		ip := net.ParseIP(test.ip)
		assert.NotNil(t, ip, "Invalid IP address format")
		result := IsInternalIP(test.ip)

		assert.Equal(t, test.expected, result, "IP: %s", test.ip)
	}
}
