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

// Call_GetVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoEncoderConfigurationResponse.
func Call_GetVideoEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.GetVideoEncoderConfiguration) (media.GetVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoEncoderConfigurationResponse media.GetVideoEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoEncoderConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoEncoderConfiguration")
		return reply.Body.GetVideoEncoderConfigurationResponse, err
	}
}
