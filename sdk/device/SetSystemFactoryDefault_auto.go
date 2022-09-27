// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/gordcurrie/onvif"
	"github.com/gordcurrie/onvif/sdk"
	"github.com/gordcurrie/onvif/device"
)

// Call_SetSystemFactoryDefault forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSystemFactoryDefaultResponse.
func Call_SetSystemFactoryDefault(ctx context.Context, dev *onvif.Device, request device.SetSystemFactoryDefault) (device.SetSystemFactoryDefaultResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemFactoryDefaultResponse device.SetSystemFactoryDefaultResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetSystemFactoryDefaultResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetSystemFactoryDefault")
		return reply.Body.SetSystemFactoryDefaultResponse, err
	}
}
