// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upnp

const (
	ProductName = "go-net-upnp"

	SupportVersion      = "1.1"
	SupportVersionMajor = 1
	SupportVersionMinor = 1

	ControlPointDefaultPortBase  = 5004
	ControlPointDefaultPortRange = 1024
	ControlPointDefaultPortMax   = ControlPointDefaultPortBase + ControlPointDefaultPortRange

	DeviceDefaultPortBase  = 6004
	DeviceDefaultPortRange = 1024
	DeviceDefaultPortMax   = DeviceDefaultPortBase + DeviceDefaultPortRange

	XmlMarshallIndent = " "
)
