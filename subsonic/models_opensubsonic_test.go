package subsonic

import (
	"bytes"
	"encoding/xml"
	"testing"
)

const (
	songLyricsResp string = `<subsonic-response status="ok" version="1.16.1" type="AwesomeServerName" serverVersion="0.1.3 (tag)" openSubsonic="true">
    <lyricsList>
      <structuredLyrics displayArtist="Muse" displayTitle="Hysteria" lang="en" offset="-100" synced="true">
        <line start="0">It's bugging me</line>
        <line start="2000">Grating me</line>
        <line start="3001">And twisting me around...</line>
      </structuredLyrics>
      <structuredLyrics displayArtist="Muse" displayTitle="Hysteria" lang="en" offset="100" synced="false">
        <line>It's bugging me</line>
        <line>Grating me</line>
        <line>And twisting me around...</line>
      </structuredLyrics>
    </lyricsList>
  </subsonic-response>`

	songLyricsNavidromeIncorrectResp string = `<subsonic-response status="ok" version="1.16.1" type="AwesomeServerName" serverVersion="0.1.3 (tag)" openSubsonic="true">
    <lyricsList>
      <structuredLyrics displayArtist="Muse" displayTitle="Hysteria" lang="en" offset="-100" synced="true">
        <line start="0">
		  <value>It's bugging me</value>
		</line>
        <line start="2000">
		  <value>Grating me</value>
		</line>
        <line start="3001">
		  <value>And twisting me around...</value>
		</line>
      </structuredLyrics>
      <structuredLyrics displayArtist="Muse" displayTitle="Hysteria" lang="en" offset="100" synced="false">
        <line>
		  <value>It's bugging me</value>
		</line>
        <line><value>Grating me</value></line>
        <line><value>And twisting me around...</value></line>
      </structuredLyrics>
    </lyricsList>
  </subsonic-response>`
)

func TestLyricsBySongId(t *testing.T) {
	var response Response
	err := xml.Unmarshal([]byte(songLyricsResp), &response)
	if err != nil {
		t.Errorf("Error unmarshaling lyrics XML: %v", err)
	}
	if response.OpenSubsonic == false {
		t.Error("wrong OpenSubsonic")
	}
	if response.LyricsList == nil {
		t.Error("nil response.LyricsList")
	}
	if l := len(response.LyricsList.StructuredLyrics); l != 2 {
		t.Errorf("wrong length of StructuredLyrics (want 2, got %d)", l)
	}
	if o := response.LyricsList.StructuredLyrics[0].Offset; o != -100 {
		t.Errorf("wrong lyric offset (want -100, got %d)", o)
	}
	if l := len(response.LyricsList.StructuredLyrics[0].Lines); l != 3 {
		t.Errorf("wrong line count (want 3, got %d)", l)
	}
	line := response.LyricsList.StructuredLyrics[0].Lines[1]
	if line.Start != 2000 {
		t.Errorf("wrong line start (want 2000, got %d)", line.Start)
	}
	if line.Text != "Grating me" {
		t.Errorf("wrong line text, want %q, got %q)", "Grating me", line.Text)
	}
	if line.Value != "" {
		t.Errorf("Value (Navidrome incorrect field) should be empty for compliant response")
	}
}

const nowPlayingWithPlaybackReportResp = `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.16.1" type="AwesomeServerName" serverVersion="0.1.3 (tag)" openSubsonic="true">
  <nowPlaying>
    <entry id="123" username="user" minutesAgo="0" playerId="0" isDir="false" title="Take the Home"
           state="playing" positionMs="120000" playbackRate="1.5"/>
  </nowPlaying>
</subsonic-response>`

func TestNowPlayingEntry_PlaybackReportFields(t *testing.T) {
	resp := bytes.NewReader([]byte(nowPlayingWithPlaybackReportResp))
	unmarshaled, err := unmarshalResponse(resp)
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if unmarshaled.NowPlaying == nil {
		t.Fatal("nil NowPlaying")
	}
	if l := len(unmarshaled.NowPlaying.Entry); l != 1 {
		t.Fatalf("wrong number of NowPlaying entries: %d", l)
	}
	entry := unmarshaled.NowPlaying.Entry[0]
	if entry.State != "playing" {
		t.Errorf("wrong State: want %q, got %q", "playing", entry.State)
	}
	if entry.PositionMs != 120000 {
		t.Errorf("wrong PositionMs: want %d, got %d", 120000, entry.PositionMs)
	}
	if entry.PlaybackRate != 1.5 {
		t.Errorf("wrong PlaybackRate: want %v, got %v", 1.5, entry.PlaybackRate)
	}
}

func TestLyricsBySongId_NavidromeIncorrectResponse(t *testing.T) {
	var response Response
	err := xml.Unmarshal([]byte(songLyricsNavidromeIncorrectResp), &response)
	if err != nil {
		t.Errorf("Error unmarshaling lyrics XML: %v", err)
	}
	if response.OpenSubsonic == false {
		t.Error("wrong OpenSubsonic")
	}
	if response.LyricsList == nil {
		t.Error("nil response.LyricsList")
	}
	if l := len(response.LyricsList.StructuredLyrics); l != 2 {
		t.Errorf("wrong length of StructuredLyrics (want 2, got %d)", l)
	}
	if o := response.LyricsList.StructuredLyrics[0].Offset; o != -100 {
		t.Errorf("wrong lyric offset (want -100, got %d)", o)
	}
	if l := len(response.LyricsList.StructuredLyrics[0].Lines); l != 3 {
		t.Errorf("wrong line count (want 3, got %d)", l)
	}
	line := response.LyricsList.StructuredLyrics[0].Lines[1]
	if line.Start != 2000 {
		t.Errorf("wrong line start (want 2000, got %d)", line.Start)
	}
	if line.Value != "Grating me" {
		t.Errorf("wrong line text, want %q, got %q)", "Grating me", line.Text)
	}
}
