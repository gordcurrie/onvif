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

// Call_GetDeviceInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDeviceInformationResponse.
func Call_GetDeviceInformation(ctx context.Context, dev *onvif.Device, request device.GetDeviceInformation) (device.GetDeviceInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDeviceInformationResponse device.GetDeviceInformationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDeviceInformationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDeviceInformation")
		return reply.Body.GetDeviceInformationResponse, err
	}
}
