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

// Call_GetAudioOutputs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputsResponse.
func Call_GetAudioOutputs(ctx context.Context, dev *onvif.Device, request media.GetAudioOutputs) (media.GetAudioOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputsResponse media.GetAudioOutputsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioOutputsResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioOutputs")
		return reply.Body.GetAudioOutputsResponse, err
	}
}
