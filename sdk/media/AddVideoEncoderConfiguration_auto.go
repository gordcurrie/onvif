// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"git.taservs.net/federal-devices/go-onvif"
	"git.taservs.net/federal-devices/go-onvif/sdk"
	"git.taservs.net/federal-devices/go-onvif/media"
)

// Call_AddVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddVideoEncoderConfigurationResponse.
func Call_AddVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.AddVideoEncoderConfiguration) (media.AddVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoEncoderConfigurationResponse media.AddVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddVideoEncoderConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddVideoEncoderConfiguration")
		return reply.Body.AddVideoEncoderConfigurationResponse, err
	}
}
