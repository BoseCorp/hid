// +build required

// Package dummy prevents the C dependency stripping for go mod vendor.
package dummy

import (
	// Import the C dependencies.
	_ "github.com/Fatih-Cetinkaya-Bose/hid/libusb/libusb/os"
)
