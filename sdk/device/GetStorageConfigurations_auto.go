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

// Call_GetStorageConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStorageConfigurationsResponse.
func Call_GetStorageConfigurations(ctx context.Context, dev *onvif.Device, request device.GetStorageConfigurations) (device.GetStorageConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStorageConfigurationsResponse device.GetStorageConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetStorageConfigurationsResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetStorageConfigurations")
		return reply.Body.GetStorageConfigurationsResponse, err
	}
}
