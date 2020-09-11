package zoom

import "fmt"

type (

	// ListEndedMeetingsOptions contains options for ListEndedMeetings
	ListEndedMeetingsOptions struct {
		MeetingID int `url:"-"`
	}

	// ListEndedMeetingsResponse container the response from a call to ListEndedMeetings
	ListEndedMeetingsResponse struct {
		Meetings []ListMeeting `json:"meetings"`
	}

	// ListEndedMeeting represents a meeting object returned by ListEndedMeetings endpoint
	ListEndedMeeting struct {
		UUID      string `json:"uuid"`
		StartTime *Time  `json:"start_time"`
	}
)

const (
	// ListEndedMeetingsPath - v2 list all ended instances of a meeting
	ListEndedMeetingsPath = "/past_meetings/%d/instances"
)

// ListMeetings calls /users/ID/meetings
func ListEndedMeetings(opts ListEndedMeetingsOptions) (ListEndedMeetingsResponse, error) {
	return defaultClient.ListEndedMeetings(opts)
}

// ListMeetings calls /users/ID/meetings
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetings
func (c *Client) ListEndedMeetings(opts ListEndedMeetingsOptions) (ListEndedMeetingsResponse, error) {
	var ret = ListEndedMeetingsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(ListEndedMeetingsPath, opts.MeetingID),
		URLParameters: &opts,
		Ret:           &ret,
	})
}
