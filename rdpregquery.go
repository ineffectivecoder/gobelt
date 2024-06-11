package gobelt

import (
	"fmt"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

// leveraging win32 api to enumerate list of successful RDP sessions from HKCurrentUser
func RDPRegQuery() ([]string, error) {
	var e error
	var k registry.Key
	var value []string
	// HKEY_CURRENT_USER\Software\Microsoft\Terminal Server Client\Default
	fmt.Print("[+] Querying RDP hosts Registry Entry\n")
	k, e = registry.OpenKey(registry.CURRENT_USER,
		filepath.Join("Software",
			"Microsoft",
			"Terminal Server Client",
			"Default"),
		registry.QUERY_VALUE,
	)
	if e != nil {
		return nil, e
	}
	keys, e := k.ReadValueNames(0)
	if e != nil {
		return nil, e
	}
	fmt.Printf("[+] The following values found %v\n", keys)

	for _, key := range keys {
		regvalue, _, err := k.GetStringValue(key)
		if err != nil {
			return nil, e
		}
		if regvalue == "" {

			return nil, fmt.Errorf("registry key is empty")

		}
		value = append(value, regvalue)
	}
	return value, nil

}
