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

// Call_CreateStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateStorageConfigurationResponse.
func Call_CreateStorageConfiguration(ctx context.Context, dev *onvif.Device, request device.CreateStorageConfiguration) (device.CreateStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateStorageConfigurationResponse device.CreateStorageConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateStorageConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateStorageConfiguration")
		return reply.Body.CreateStorageConfigurationResponse, err
	}
}
