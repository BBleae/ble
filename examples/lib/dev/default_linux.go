package dev

import (
	"github.com/BBleae/ble"
	"github.com/BBleae/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
