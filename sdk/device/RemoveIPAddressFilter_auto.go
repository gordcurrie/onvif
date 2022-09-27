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

// Call_RemoveIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveIPAddressFilterResponse.
func Call_RemoveIPAddressFilter(ctx context.Context, dev *onvif.Device, request device.RemoveIPAddressFilter) (device.RemoveIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveIPAddressFilterResponse device.RemoveIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveIPAddressFilterResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveIPAddressFilter")
		return reply.Body.RemoveIPAddressFilterResponse, err
	}
}
