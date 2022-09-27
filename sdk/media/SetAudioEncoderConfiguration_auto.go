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

// Call_SetAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioEncoderConfigurationResponse.
func Call_SetAudioEncoderConfiguration(ctx context.Context, dev *onvif.Device, request media.SetAudioEncoderConfiguration) (media.SetAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioEncoderConfigurationResponse media.SetAudioEncoderConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetAudioEncoderConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetAudioEncoderConfiguration")
		return reply.Body.SetAudioEncoderConfigurationResponse, err
	}
}
