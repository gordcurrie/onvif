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

// Call_GetWsdlUrl forwards the call to dev.CallMethod() then parses the payload of the reply as a GetWsdlUrlResponse.
func Call_GetWsdlUrl(ctx context.Context, dev *onvif.Device, request device.GetWsdlUrl) (device.GetWsdlUrlResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetWsdlUrlResponse device.GetWsdlUrlResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetWsdlUrlResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetWsdlUrl")
		return reply.Body.GetWsdlUrlResponse, err
	}
}
