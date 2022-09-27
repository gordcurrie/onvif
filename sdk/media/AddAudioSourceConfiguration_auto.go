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

// Call_AddAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioSourceConfigurationResponse.
func Call_AddAudioSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.AddAudioSourceConfiguration) (media.AddAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioSourceConfigurationResponse media.AddAudioSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddAudioSourceConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddAudioSourceConfiguration")
		return reply.Body.AddAudioSourceConfigurationResponse, err
	}
}
