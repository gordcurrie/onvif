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

// Call_GetAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputConfigurationResponse.
func Call_GetAudioOutputConfiguration(ctx context.Context, dev *onvif.Device, request media.GetAudioOutputConfiguration) (media.GetAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationResponse media.GetAudioOutputConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioOutputConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioOutputConfiguration")
		return reply.Body.GetAudioOutputConfigurationResponse, err
	}
}
