// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/use-go/onvif"
	"github.com/use-go/onvif/sdk"
	"github.com/use-go/onvif/media"
)

// Call_GetAudioOutputConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputConfigurationsResponse.
func Call_GetAudioOutputConfigurations(ctx context.Context, dev *onvif.Device, request media.GetAudioOutputConfigurations) (media.GetAudioOutputConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationsResponse media.GetAudioOutputConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAudioOutputConfigurationsResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAudioOutputConfigurations")
		return reply.Body.GetAudioOutputConfigurationsResponse, err
	}
}
