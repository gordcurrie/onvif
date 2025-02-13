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

// Call_GetMetadataConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMetadataConfigurationOptionsResponse.
func Call_GetMetadataConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetMetadataConfigurationOptions) (media.GetMetadataConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationOptionsResponse media.GetMetadataConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetadataConfigurationOptionsResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetMetadataConfigurationOptions")
		return reply.Body.GetMetadataConfigurationOptionsResponse, err
	}
}
