package gobelt

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows/registry"
)

// not including all of the fields for AntivirusProduct, just 
// what we need
type AntiVirusProduct struct {
	DisplayName string
	InstanceGuid string
	PathToSignedProductExe string
	PathToSignedReportingExe string
}

// not using all of the fields at the moment, but leaving them 
// in here in case we want to use them later
type Win32_NetworkConnection struct {
	Caption string
	Description string
	InstallDate string
	Status string
	AccessMask uint32
	Comment string
	ConnectionState string
	ConnectionType string
	DisplayType string
	LocalName string
	Name string
	Persistent bool
	ProviderName string
	RemoteName string
	RemotePath string
	ResourceType string
	UserName string
}

func AntiVirus() Result {
	var dst []AntiVirusProduct
	var err error
	var q string
	var value []string

	value = append(
		value,
		"[+] Retrieving list of antivirus products installed",
	)

	q = wmi.CreateQuery(&dst, "", "AntiVirusProduct")
	err = wmi.QueryNamespace(q, &dst, `root\SecurityCenter2`)
	if err != nil {
		return Result {
			Kind: KindError,
			Error: err,
		}
	}

	if len(dst) == 0 {
		return Result {
			Kind: KindError,
			Error: errors.New("No antivirus found"),
		}
	}

	value = append(value, "[+] Installed Antivirus Products:\n")
	for _, av := range(dst) {
		value = append(
			value,
			"Name:                  " + av.DisplayName,
			"Instance GUID:         " + av.InstanceGuid,
			"Path to Product EXE:   " + av.PathToSignedProductExe,
			"Path to Reporting EXE: " + av.PathToSignedReportingExe,
			"\n",
		)
	}

	return Result{
		Kind: KindInfo,
		Data: value,
	}
}

// retrieves list of mapped drives using WMI to query 
// Win32_NetworkConnection
func MappedDrives() Result {
	var dst []Win32_NetworkConnection
	var err error
	var q string
	var value []string

	value = append(
		value, 
		"[+] Retrieving list of mapped drives using WMI",
	)

	q = wmi.CreateQuery(&dst, "")
	err = wmi.Query(q, &dst)
	if err != nil {
		return Result {
			Kind: KindError,
			Error: err,
		}
	}

	if len(dst) == 0 {
		return Result {
			Kind: KindError,
			Error: errors.New("No mapped drives found"),
		}
	}

	value = append(value, "[+] Mapped Drives:\n")
	for _, mappeddrive := range(dst) {
		value = append(
			value,
			"Local Name:          " + mappeddrive.LocalName,
			"Remote Name:         " + mappeddrive.RemoteName,
			"Remote Path:         " + mappeddrive.RemotePath,
			"Status:              " + mappeddrive.Status,
			"ConnectionState:     " + mappeddrive.ConnectionState,
			"Persistent:          " + 
				fmt.Sprintf("%v",mappeddrive.Persistent),
			"UserName:            " + mappeddrive.UserName,
			"Description:         " + mappeddrive.Description,
			"\n",
		)
	}

	return Result{
		Kind: KindInfo,
		Data: value,
	}
}

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

