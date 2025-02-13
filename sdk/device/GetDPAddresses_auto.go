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

// Call_GetDPAddresses forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDPAddressesResponse.
func Call_GetDPAddresses(ctx context.Context, dev *onvif.Device, request device.GetDPAddresses) (device.GetDPAddressesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDPAddressesResponse device.GetDPAddressesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDPAddressesResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDPAddresses")
		return reply.Body.GetDPAddressesResponse, err
	}
}
