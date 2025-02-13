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

// Call_GetServices forwards the call to dev.CallMethod() then parses the payload of the reply as a GetServicesResponse.
func Call_GetServices(ctx context.Context, dev *onvif.Device, request device.GetServices) (device.GetServicesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServicesResponse device.GetServicesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetServicesResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetServices")
		return reply.Body.GetServicesResponse, err
	}
}
