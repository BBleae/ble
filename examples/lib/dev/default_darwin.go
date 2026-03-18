package dev

import (
	"github.com/BBleae/ble"
	"github.com/BBleae/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
