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

// Call_AddAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioOutputConfigurationResponse.
func Call_AddAudioOutputConfiguration(ctx context.Context, dev *onvif.Device, request media.AddAudioOutputConfiguration) (media.AddAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioOutputConfigurationResponse media.AddAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddAudioOutputConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "AddAudioOutputConfiguration")
		return reply.Body.AddAudioOutputConfigurationResponse, err
	}
}
