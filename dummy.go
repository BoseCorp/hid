// +build required

package hid

import (
	// Import the C dependencies.
	_ "github.com/Fatih-Cetinkaya-Bose/hid/hidapi"
	_ "github.com/Fatih-Cetinkaya-Bose/hid/libusb"
)
