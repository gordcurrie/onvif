// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"github.com/gordcurrie/onvif"
	"github.com/gordcurrie/onvif/sdk"
	"github.com/gordcurrie/onvif/media"
)

// Call_RemovePTZConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemovePTZConfigurationResponse.
func Call_RemovePTZConfiguration(ctx context.Context, dev *onvif.Device, request media.RemovePTZConfiguration) (media.RemovePTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemovePTZConfigurationResponse media.RemovePTZConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemovePTZConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemovePTZConfiguration")
		return reply.Body.RemovePTZConfigurationResponse, err
	}
}
