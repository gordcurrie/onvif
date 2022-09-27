// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"git.taservs.net/federal-devices/go-onvif"
	"git.taservs.net/federal-devices/go-onvif/sdk"
	"git.taservs.net/federal-devices/go-onvif/device"
)

// Call_AddIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a AddIPAddressFilterResponse.
func Call_AddIPAddressFilter(ctx context.Context, dev *onvif.Device, request device.AddIPAddressFilter) (device.AddIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddIPAddressFilterResponse device.AddIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddIPAddressFilterResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddIPAddressFilter")
		return reply.Body.AddIPAddressFilterResponse, err
	}
}
