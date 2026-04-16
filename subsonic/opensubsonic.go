package subsonic

import (
	"fmt"
	"net/url"
)

// OpenSubsonic extension names
const (
	SongLyricsExtension  = "songLyrics"
	TranscodeOffset      = "transcodeOffset"
	IndexBasedQueue      = "indexBasedQueue"
	HTTPFormPost         = "formPost"
	PlaybackReport       = "playbackReport"
)

// PlaybackState represents the playback state for a ReportPlayback call.
type PlaybackState string

const (
	PlaybackStateStarting PlaybackState = "starting"
	PlaybackStatePlaying  PlaybackState = "playing"
	PlaybackStatePaused   PlaybackState = "paused"
	PlaybackStateStopped  PlaybackState = "stopped"
)

// PlaybackMediaType represents the type of media being reported.
type PlaybackMediaType string

const (
	PlaybackMediaTypeSong    PlaybackMediaType = "song"
	PlaybackMediaTypePodcast PlaybackMediaType = "podcast"
)

// ReportPlaybackParameters holds the parameters for a ReportPlayback call.
type ReportPlaybackParameters struct {
	// MediaID is the ID of the media being reported. Required.
	MediaID string
	// MediaType is either "song" or "podcast". Required.
	MediaType PlaybackMediaType
	// PositionMs is the playback position in milliseconds. Required.
	PositionMs int64
	// State is the current playback state. Required.
	State PlaybackState
	// PlaybackRate is the playback speed multiplier. Optional, defaults to 1.0.
	PlaybackRate *float64
	// IgnoreScrobble if true, the server should only update now-playing state
	// and not trigger scrobble/playcount side effects. Optional, defaults to false.
	IgnoreScrobble *bool
}

// Get the list of supported OpenSubsonic extensions for this server.
func (c *Client) GetOpenSubsonicExtensions() ([]*OpenSubsonicExtension, error) {
	resp, err := c.Get("getOpenSubsonicExtensions", nil)
	if err != nil {
		return nil, err
	}
	if resp.OpenSubsonicExtensions == nil {
		return nil, nil
	}
	return resp.OpenSubsonicExtensions, nil
}

// Get structured lyrics for a track.
//
// Server must support OpenSubsonic songLyrics extension
func (c *Client) GetLyricsBySongId(songID string) (*LyricsList, error) {
	resp, err := c.Get("getLyricsBySongId", map[string]string{"id": songID})
	if err != nil {
		return nil, err
	}
	return resp.LyricsList, nil
}

// GetPlayQueueByIndex returns the state of the play queue for this user
// (as set by savePlayQueueByIndex). This includes the tracks in the
// play queue, the currently playing track by index, and the position
// within this track.
//
// Server must support OpenSubsonic playQueueByIndex extension.
func (c *Client) GetPlayQueueByIndex() (*PlayQueueByIndex, error) {
	resp, err := c.Get("getPlayQueueByIndex", nil)
	if err != nil {
		return nil, err
	}
	return resp.PlayQueueByIndex, nil
}

// SavePlayQueueByIndex saves the state of the play queue for this user.
// This includes the tracks in the play queue, the currently playing
// track by index, and the position within this track.
//
// Parameters:
//
//	songIDs: IDs of the songs in the play queue
//
// Optional parameters:
//
//	currentIndex: the index of the currently playing song (0-based)
//	position: The position in milliseconds within the currently playing song
//
// Server must support OpenSubsonic playQueueByIndex extension.
func (c *Client) SavePlayQueueByIndex(songIDs []string, params map[string]string) error {
	values := url.Values{}
	for _, trID := range songIDs {
		values.Add("id", trID)
	}
	for k, v := range params {
		values.Add(k, v)
	}
	_, err := c.getValues("savePlayQueueByIndex", values)
	return err
}

// ReportPlayback reports the playback timeline state for a media item.
// Clients should call this at least on each state change.
//
// Server must support OpenSubsonic playbackReport extension.
func (c *Client) ReportPlayback(params ReportPlaybackParameters) error {
	values := url.Values{}
	values.Set("mediaId", params.MediaID)
	values.Set("mediaType", string(params.MediaType))
	values.Set("positionMs", fmt.Sprintf("%d", params.PositionMs))
	values.Set("state", string(params.State))
	if params.PlaybackRate != nil {
		values.Set("playbackRate", fmt.Sprintf("%g", *params.PlaybackRate))
	}
	if params.IgnoreScrobble != nil {
		if *params.IgnoreScrobble {
			values.Set("ignoreScrobble", "true")
		} else {
			values.Set("ignoreScrobble", "false")
		}
	}
	_, err := c.getValues("reportPlayback", values)
	return err
}
