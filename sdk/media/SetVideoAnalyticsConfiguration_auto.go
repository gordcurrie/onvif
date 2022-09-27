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

// Call_SetVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoAnalyticsConfigurationResponse.
func Call_SetVideoAnalyticsConfiguration(ctx context.Context, dev *onvif.Device, request media.SetVideoAnalyticsConfiguration) (media.SetVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoAnalyticsConfigurationResponse media.SetVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetVideoAnalyticsConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetVideoAnalyticsConfiguration")
		return reply.Body.SetVideoAnalyticsConfigurationResponse, err
	}
}
