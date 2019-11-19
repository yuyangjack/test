package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(brokerEntrySeemsValid("baas-dfdsdfsfs-aaa:9911"))
}

// This does just a barebones sanity check.
func brokerEntrySeemsValid(broker string) bool {
	if !strings.Contains(broker, ":") {
		return false
	}

	parts := strings.Split(broker, ":")
	if len(parts) > 2 {
		return false
	}

	host := parts[0]
	port := parts[1]

	if _, err := strconv.ParseUint(port, 10, 16); err != nil {
		return false
	}

	// Valid hostnames may contain only the ASCII letters 'a' through 'z' (in a
	// case-insensitive manner), the digits '0' through '9', and the hyphen. IP
	// v4 addresses are  represented in dot-decimal notation, which consists of
	// four decimal numbers, each ranging from 0 to 255, separated by dots,
	// e.g., 172.16.254.1
	// The following regular expression:
	// 1. allows just a-z (case-insensitive), 0-9, and the dot and hyphen characters
	// 2. does not allow leading trailing dots or hyphens
	re, _ := regexp.Compile("^([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9.-]*[a-zA-Z0-9])$")
	matched := re.FindString(host)
	return len(matched) == len(host)
}
