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

// Call_GetVideoAnalyticsConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoAnalyticsConfigurationResponse.
func Call_GetVideoAnalyticsConfiguration(ctx context.Context, dev *onvif.Device, request media.GetVideoAnalyticsConfiguration) (media.GetVideoAnalyticsConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoAnalyticsConfigurationResponse media.GetVideoAnalyticsConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoAnalyticsConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoAnalyticsConfiguration")
		return reply.Body.GetVideoAnalyticsConfigurationResponse, err
	}
}
