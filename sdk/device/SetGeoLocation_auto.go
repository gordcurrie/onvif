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

// Call_SetGeoLocation forwards the call to dev.CallMethod() then parses the payload of the reply as a SetGeoLocationResponse.
func Call_SetGeoLocation(ctx context.Context, dev *onvif.Device, request device.SetGeoLocation) (device.SetGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetGeoLocationResponse device.SetGeoLocationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetGeoLocationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetGeoLocation")
		return reply.Body.SetGeoLocationResponse, err
	}
}
