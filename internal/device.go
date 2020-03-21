package internal

import (
	"net"
	"os/user"
	"runtime"
)

type Device struct {
	Macaddresses       []string
	Devicetype         string
	Devicename         string
	Username           string
	Userhomedir        string
	Os                 string
	Arch               string
	Machineusers       []string
	OsInstallationdate string
}

// Load all information of the device
func (dev *Device) LoadInfo() {
	dev.loadUserInfo()

	dev.Devicetype = "desktop" // Always desktop here (computer|desktop program)
	dev.Macaddresses, _ = dev.DiscoverMACAddresses()
	dev.Devicename, _ = dev.DiscoverDeviceName()
	dev.Os, _ = dev.DiscoverOS()
	dev.Arch, _ = dev.DiscoverARCH()
	dev.Machineusers, _ = dev.DiscoverMachineUsers()
	dev.OsInstallationdate, _ = dev.DiscoverOSInstallationDate()
}

// Discover MAC Address
func (dev *Device) DiscoverMACAddresses() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

//  Loads user information
func (dev *Device) loadUserInfo() (string, error) {

	user, err := user.Current()
	if err != nil {
		return "", err
	}

	dev.Username = user.Username
	dev.Userhomedir = user.HomeDir

	return user.Username, nil
}

// Discover device name
func (dev *Device) DiscoverDeviceName() (string, error) {
	// TODO: find a way multiplatform way to find the machine name
	return "", nil
}

// Discover OS
func (dev *Device) DiscoverOS() (string, error) {
	return runtime.GOOS, nil
}

// Discover ARCH
func (dev *Device) DiscoverARCH() (string, error) {
	return runtime.GOARCH, nil
}

// Discover all the machine users in this computer
func (dev *Device) DiscoverMachineUsers() ([]string, error) {
	// TODO: find a multiplatform way to discover machine users
	return nil, nil
}

// Discover installation date of the OS
func (dev *Device) DiscoverOSInstallationDate() (string, error) {
	// TODO: find a multiplatform ay to discover the installation date of the OS
	return "", nil
}

func NewDevice() Device {
	return Device{}
}
