package gobelt

import (
	"errors"
	"fmt"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

// leveraging win32 api to enumerate list of successful RDP sessions from HKCurrentUser

func RDPRegQuery() Result {
	var k registry.Key
	var value []string
	// HKEY_CURRENT_USER\Software\Microsoft\Terminal Server Client\Default
	fmt.Print("[+] Querying RDP hosts Registry Entry\n")
	k, err := registry.OpenKey(registry.CURRENT_USER,
		filepath.Join("Software",
			"Microsoft",
			"Terminal Server Client",
			"Default"),
		registry.QUERY_VALUE,
	)
	if err != nil {
		return Result{
			Kind:  KindError,
			Error: err,
		}
	}
	keys, err := k.ReadValueNames(0)
	if err != nil {
		return Result{
			Kind:  KindError,
			Error: err,
		}
	}
	fmt.Printf("[+] The following values found %v\n", keys)

	for _, key := range keys {
		regvalue, _, err := k.GetStringValue(key)
		if err != nil {
			return Result{
				Kind:  KindError,
				Error: err,
			}
		}
		if regvalue == "" {
			return Result{
				Kind:  KindError,
				Error: errors.New("registry key is empty"),
			}
		}
		value = append(value, regvalue)
	}
	return Result{
		Kind: KindInfo,
		Data: value,
	}
}
