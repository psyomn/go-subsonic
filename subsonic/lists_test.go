package subsonic

import (
	"encoding/xml"
	"testing"
	"time"
)

const albumList2JSON = `{"subsonic-response":{"status":"ok","version":"1.16.1","type":"navidrome","serverVersion":"0.60.3 (34c6f12a)","openSubsonic":true,"albumList2":{"album":[{"id":"2cahnhu6UPen2wbYDm4CHK","name":"8-bit lagerfeuer","artist":"pornophonique","artistId":"24jpYAI4N8TG0cCAYyzZk5","coverArt":"al-2cahnhu6UPen2wbYDm4CHK_640a93a8","songCount":8,"duration":1954,"playCount":3678,"created":"2025-03-26T22:26:48.822020234Z","year":2007,"played":"2026-03-11T03:01:26.176Z","userRating":5,"averageRating":5,"genres":[],"musicBrainzId":"","isCompilation":false,"sortName":"8-bit lagerfeuer","discTitles":[],"originalReleaseDate":{},"releaseDate":{},"releaseTypes":[],"recordLabels":[],"moods":[],"artists":[{"id":"24jpYAI4N8TG0cCAYyzZk5","name":"pornophonique"}],"displayArtist":"pornophonique","explicitStatus":"","version":""},{"id":"6oq0g9th4iTM8a0y0T88tr","name":"netBloc Vol. 24: tiuqottigeloot","artist":"Various Artists","artistId":"63sqASlAfjbGMuLP4JhnZU","coverArt":"al-6oq0g9th4iTM8a0y0T88tr_640a939d","songCount":12,"duration":2615,"playCount":1359,"created":"2025-03-26T22:26:49.228824721Z","year":2009,"genre":"Netlabel","played":"2026-03-11T02:56:45.641Z","userRating":3,"averageRating":3,"genres":[{"name":"Netlabel"}],"musicBrainzId":"","isCompilation":true,"sortName":"netbloc vol. 24: tiuqottigeloot","discTitles":[],"originalReleaseDate":{},"releaseDate":{},"releaseTypes":[],"recordLabels":[],"moods":[],"artists":[{"id":"63sqASlAfjbGMuLP4JhnZU","name":"Various Artists"}],"displayArtist":"Various Artists","explicitStatus":"","version":""},{"id":"3ZE7pthjfyWFhRvT4JKKqJ","name":"Diplomatic Immunity","artist":"The Polish Ambassador","artistId":"3Vy6szDT2gBLGKc8yCcf5r","coverArt":"al-3ZE7pthjfyWFhRvT4JKKqJ_640a9360","songCount":20,"duration":3411,"playCount":1644,"created":"2025-03-26T22:26:47.074844665Z","starred":"2026-02-27T12:01:13.99028623Z","year":2014,"genre":"Electronic","played":"2026-03-11T02:45:21.466Z","userRating":5,"averageRating":5,"genres":[{"name":"Electronic"}],"musicBrainzId":"","isCompilation":false,"sortName":"diplomatic immunity","discTitles":[],"originalReleaseDate":{},"releaseDate":{},"releaseTypes":[],"recordLabels":[{"name":"http://www.jamendo.com"}],"moods":[],"artists":[{"id":"3Vy6szDT2gBLGKc8yCcf5r","name":"The Polish Ambassador"}],"displayArtist":"The Polish Ambassador","explicitStatus":"","version":""},{"id":"3eQTMlSmKBB63P4U2fEELP","name":"Ghosts V: Together","artist":"Nine Inch Nails","artistId":"3GSnSEURz17ltddsamzmSD","coverArt":"al-3eQTMlSmKBB63P4U2fEELP_640a931d","songCount":8,"duration":4218,"playCount":496,"created":"2025-03-26T22:26:45.003988183Z","year":2020,"played":"2026-03-11T02:44:44.59Z","userRating":4,"averageRating":4,"genres":[],"musicBrainzId":"","isCompilation":false,"sortName":"ghosts v: together","discTitles":[],"originalReleaseDate":{},"releaseDate":{},"releaseTypes":[],"recordLabels":[],"moods":[],"artists":[{"id":"3GSnSEURz17ltddsamzmSD","name":"Nine Inch Nails"}],"displayArtist":"Nine Inch Nails","explicitStatus":"","version":""},{"id":"6X1IJN63h7skU4qjlBVM5R","name":"Deflonesoundsystem","artist":"Deflone","artistId":"3eeD9vic30dCDcUq0X2ICg","coverArt":"al-6X1IJN63h7skU4qjlBVM5R_640a92c4","songCount":13,"duration":2922,"playCount":2725,"created":"2025-03-26T22:26:39.580618167Z","year":2008,"genre":"Hip-Hop","played":"2026-03-11T02:43:20.985Z","userRating":4,"averageRating":4,"genres":[{"name":"Hip-Hop"}],"musicBrainzId":"","isCompilation":false,"sortName":"deflonesoundsystem","discTitles":[],"originalReleaseDate":{},"releaseDate":{},"releaseTypes":[],"recordLabels":[],"moods":[],"artists":[{"id":"3eeD9vic30dCDcUq0X2ICg","name":"Deflone"}],"displayArtist":"Deflone","explicitStatus":"","version":""},{"id":"6dGMsiesuYw9lti9yGIGWY","name":"netBloc Vol. 42: Live, The Universe \u0026 Everything","artist":"Various Artists","artistId":"63sqASlAfjbGMuLP4JhnZU","coverArt":"al-6dGMsiesuYw9lti9yGIGWY_640a9396","songCount":11,"duration":2954,"playCount":5320,"created":"2025-03-26T22:26:49.315688846Z","year":2013,"played":"2026-03-11T02:34:50.41Z","userRating":2,"averageRating":2,"genres":[],"musicBrainzId":"","isCompilation":true,"sortName":"netbloc vol. 42: live, the universe \u0026 everything","discTitles":[],"originalReleaseDate":{},"releaseDate":{},"releaseTypes":[],"recordLabels":[],"moods":[],"artists":[{"id":"63sqASlAfjbGMuLP4JhnZU","name":"Various Artists"}],"displayArtist":"Various Artists","explicitStatus":"","version":""},{"id":"4QZO6nBGty9odiq7pUok5n","name":"Love (All I Was Able To Say)","artist":"Back On Earth","artistId":"04HrSORpypcLGNUdQp37gn","coverArt":"al-4QZO6nBGty9odiq7pUok5n_640a9269","songCount":2,"duration":372,"playCount":547,"created":"2025-03-26T22:26:35.129501286Z","starred":"2026-01-07T02:54:34.286241272Z","year":2010,"played":"2026-03-11T02:20:27.790010715Z","userRating":3,"averageRating":3,"genres":[],"musicBrainzId":"","isCompilation":false,"sortName":"love (all i was able to say)","discTitles":[],"originalReleaseDate":{},"releaseDate":{},"releaseTypes":[],"recordLabels":[{"name":"http://www.jamendo.com"}],"moods":[],"artists":[{"id":"04HrSORpypcLGNUdQp37gn","name":"Back On Earth"}],"displayArtist":"Back On Earth","explicitStatus":"","version":""},{"id":"7GZYpfn8wypaTR00LWeNI4","name":"Retroconnaissance","artist":"Ugress","artistId":"5xcMPJdeEgNrGtnzYbzAqb","coverArt":"al-7GZYpfn8wypaTR00LWeNI4_640a9384","songCount":4,"duration":801,"playCount":1477,"created":"2025-03-26T22:26:47.749422855Z","year":2006,"genre":"Electronic","played":"2026-03-11T01:39:26.276Z","userRating":5,"averageRating":5,"genres":[{"name":"Electronic"}],"musicBrainzId":"eb34953a-89a9-4c1e-ab30-7cd24abed2b6","isCompilation":false,"sortName":"retroconnaissance","discTitles":[],"originalReleaseDate":{"year":2006,"month":5,"day":14},"releaseDate":{},"releaseTypes":["ep"],"recordLabels":[{"name":"Uncanny Planet Records"}],"moods":[],"artists":[{"id":"5xcMPJdeEgNrGtnzYbzAqb","name":"Ugress"}],"displayArtist":"Ugress","explicitStatus":"","version":""},{"id":"6IBIiCkWECDiUQQUaQ8X7u","name":"Ghosts VI: Locusts","artist":"Nine Inch Nails","artistId":"3GSnSEURz17ltddsamzmSD","coverArt":"al-6IBIiCkWECDiUQQUaQ8X7u_640a932f","songCount":15,"duration":4984,"playCount":871,"created":"2025-03-26T22:26:47.617732281Z","year":2020,"played":"2026-03-11T01:01:46.267420611Z","userRating":5,"averageRating":5,"genres":[],"musicBrainzId":"","isCompilation":false,"sortName":"ghosts vi: locusts","discTitles":[],"originalReleaseDate":{},"releaseDate":{},"releaseTypes":[],"recordLabels":[],"moods":[],"artists":[{"id":"3GSnSEURz17ltddsamzmSD","name":"Nine Inch Nails"}],"displayArtist":"Nine Inch Nails","explicitStatus":"","version":""},{"id":"2yETGa7bDCWkc5UgZnZ4vm","name":"Detox Static EP","artist":"Front 242","artistId":"5r1GjGktwEMBF5CxujSJRf","coverArt":"al-2yETGa7bDCWkc5UgZnZ4vm_665fc7a7","songCount":7,"duration":1997,"playCount":399,"created":"2025-03-26T22:26:39.54532612Z","starred":"2026-02-01T03:43:52.932177165Z","year":2016,"genre":"Ebm","played":"2026-03-11T00:51:08.662Z","userRating":0,"genres":[{"name":"Ebm"},{"name":"Electro"}],"musicBrainzId":"1628bd05-fdcf-4b88-9c1a-aa6d79b6a23c","isCompilation":false,"sortName":"detox static ep","discTitles":[],"originalReleaseDate":{"year":2016,"month":1,"day":26},"releaseDate":{},"releaseTypes":["ep"],"recordLabels":[{"name":"Alfa Matrix"}],"moods":[],"artists":[{"id":"5r1GjGktwEMBF5CxujSJRf","name":"Front 242"}],"displayArtist":"Front 242","explicitStatus":"","version":""}]}}}`

const openSubsonicAlbumList = `<?xml version="1.0" encoding="utf-8"?>
<subsonic-response openSubsonic="true" serverVersion="1" status="ok" type="lms" version="1.16.0"><albumList2><album artist="Au5, Danyka Nadeau" coverArt="al-5" created="2023-10-12T18:23:30.000" duration="36" genre="Unknown" id="al-5" isCompilation="false" musicBrainzId="7a4d48b2-66b7-4a93-afbf-2b176a0c92a6" name="Follow You (The Remixes)" originalReleaseDate="2014-07-11" songCount="7" year="2014"><artists id="ar-52" name="Au5"/><artists id="ar-53" name="Danyka Nadeau"/><genres name="Unknown"/><releaseTypes>ep</releaseTypes></album><album artist="Iron Maiden" artistId="ar-45" coverArt="al-4" created="2023-10-12T18:13:19.000" duration="120" genre="Heavy Metal" id="al-4" isCompilation="true" musicBrainzId="1e84b11c-6ea1-4685-87b9-ae35591809d7" name="A Real Live Dead One" originalReleaseDate="1993-03-23" songCount="23" year="1993"><artists id="ar-45" name="Iron Maiden"/><discTitles disc="1" title="Dead One"/><discTitles disc="2" title="Live One"/><genres name="Instrumental"/><genres name="Heavy Metal"/><genres name="Metal"/><releaseTypes>album</releaseTypes><releaseTypes>compilation</releaseTypes><releaseTypes>live</releaseTypes></album><album artist="DJ Banana" artistId="ar-44" coverArt="al-3" created="2023-05-28T18:35:21.000" duration="57" genre="Techno-Industrial" id="al-3" isCompilation="false" musicBrainzId="667851cb-0f84-3fdd-8882-33902fa16aef" name="Fruit Salad Mixtape" originalReleaseDate="" songCount="11" year="2001"><artists id="ar-44" name="DJ Banana"/><genres name="Techno-Industrial"/></album><album artist="Willbe" artistId="ar-2" coverArt="al-2" created="2020-11-19T13:17:03.000" duration="4797" genre="Unknown" id="al-2" isCompilation="true" musicBrainzId="31b3b92b-b212-43de-b57d-7f2961b7cc4a" name="Demovibes 3: Pixels in sequence" originalReleaseDate="2004-06-23" songCount="22" year="2004"><artists id="ar-2" name="Willbe"/><genres name="Electronic"/><genres name="Hip-Hop"/><genres name="Industrial"/><genres name="Rap"/><genres name="Dance"/><genres name="Happy Hardcore"/><genres name="Rave"/><genres name="Trance"/><genres name="Downtempo"/><genres name="Idm"/><genres name="Unknown"/><genres name="Lounge"/><genres name="Breaks"/><genres name="Minimal"/><genres name="Breakbeat"/><genres name="Gothic Metal"/><genres name="Hardcore"/><genres name="House"/><genres name="Instrumental"/><genres name="Ambient"/><genres name="Jazz"/><genres name="Indie"/><genres name="Techno"/><genres name="Stoner Rock"/><genres name="Heavy Metal"/><releaseTypes>album</releaseTypes><releaseTypes>compilation</releaseTypes></album><album artist="Willbe" artistId="ar-2" coverArt="al-1" created="2019-11-27T20:07:27.000" duration="4797" genre="Unknown" id="al-1" isCompilation="true" musicBrainzId="91eaab8e-b472-42f9-a6e4-9e06d0464f31" name="Demovibes 5: The mod inside" originalReleaseDate="2006-04-16" songCount="20" year="2006"><artists id="ar-2" name="Willbe"/><genres name="Disco"/><genres name="Electronic"/><genres name="New Wave"/><genres name="Christian"/><genres name="Hip-Hop"/><genres name="Industrial"/><genres name="Rap"/><genres name="Dance"/><genres name="Happy Hardcore"/><genres name="Rave"/><genres name="Trance"/><genres name="Downtempo"/><genres name="Idm"/><genres name="Unknown"/><genres name="Lounge"/><genres name="Progressive Rock"/><genres name="Psychedelic Rock"/><genres name="Avant-Garde"/><genres name="Drum and Bass"/><genres name="Breaks"/><genres name="Minimal"/><genres name="Breakbeat"/><releaseTypes>album</releaseTypes><releaseTypes>compilation</releaseTypes></album></albumList2></subsonic-response>`

func TestUnmarshalOpenSubsonicAlbumList(t *testing.T) {
	parsed := Response{}
	err := xml.Unmarshal([]byte(openSubsonicAlbumList), &parsed)
	if err != nil {
		t.Fatalf("Failed: %v", err)
	}
	if parsed.AlbumList2 == nil {
		t.Fatal("No AlbumList2")
	}
	if len(parsed.AlbumList2.Album) < 1 {
		t.Error("No albums in AlbumList2")
	}
	if parsed.AlbumList2.Album[1].Artist == "" {
		t.Error("Did not parse album correctly")
	}
	if len(parsed.AlbumList2.Album[0].Artists) < 2 {
		t.Error("Did not parse OpenSubsonic Artists attribute correctly")
	}
	if !parsed.AlbumList2.Album[1].IsCompilation {
		t.Error("Did not parse isCompilation flag correctly")
	}
	if len(parsed.AlbumList2.Album[1].ReleaseTypes) != 3 {
		t.Error("Did not parse OpenSubsonci releaseTypes attribute correctly")
	}
	if parsed.AlbumList2.Album[1].ReleaseTypes[2] != "live" {
		t.Error("Did not parse OpenSubsonci releaseTypes attribute correctly")
	}
}

func TestUnmarshalJSONAlbumList2(t *testing.T) {
	parsed, err := unmarshalJSONResponse([]byte(albumList2JSON))
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Top-level response fields
	if parsed.Status != "ok" {
		t.Errorf("Wrong status: %s", parsed.Status)
	}
	if parsed.Version != "1.16.1" {
		t.Errorf("Wrong version: %s", parsed.Version)
	}
	if !parsed.OpenSubsonic {
		t.Error("OpenSubsonic should be true")
	}

	if parsed.AlbumList2 == nil {
		t.Fatal("AlbumList2 is nil")
	}
	albums := parsed.AlbumList2.Album
	if len(albums) != 10 {
		t.Fatalf("Expected 10 albums, got %d", len(albums))
	}

	// [0] "8-bit lagerfeuer" — basic field parsing
	first := albums[0]
	if first.ID != "2cahnhu6UPen2wbYDm4CHK" {
		t.Errorf("Wrong ID: %s", first.ID)
	}
	if first.Name != "8-bit lagerfeuer" {
		t.Errorf("Wrong name: %s", first.Name)
	}
	if first.Artist != "pornophonique" {
		t.Errorf("Wrong artist: %s", first.Artist)
	}
	if first.ArtistID != "24jpYAI4N8TG0cCAYyzZk5" {
		t.Errorf("Wrong artistId: %s", first.ArtistID)
	}
	if first.SongCount != 8 {
		t.Errorf("Wrong songCount: %d", first.SongCount)
	}
	if first.Duration != 1954 {
		t.Errorf("Wrong duration: %d", first.Duration)
	}
	if first.PlayCount != 3678 {
		t.Errorf("Wrong playCount: %d", first.PlayCount)
	}
	if first.Year != 2007 {
		t.Errorf("Wrong year: %d", first.Year)
	}
	if first.SortName != "8-bit lagerfeuer" {
		t.Errorf("Wrong sortName: %s", first.SortName)
	}

	// Created timestamp — exercises the xsdDateTime JSON unmarshaling path
	if first.Created.IsZero() {
		t.Error("Created timestamp was not parsed")
	}
	expectedCreated := time.Date(2025, time.March, 26, 22, 26, 48, 0, time.UTC)
	if first.Created.Year() != expectedCreated.Year() ||
		first.Created.Month() != expectedCreated.Month() ||
		first.Created.Day() != expectedCreated.Day() {
		t.Errorf("Created timestamp wrong: %v", first.Created)
	}

	// Starred absent on first album, present on third
	if !first.Starred.IsZero() {
		t.Error("First album Starred should be zero (not starred)")
	}
	starred := albums[2] // "Diplomatic Immunity"
	if starred.Starred.IsZero() {
		t.Error("Third album Starred should be set")
	}
	if starred.Starred.Year() != 2026 || starred.Starred.Month() != time.February {
		t.Errorf("Third album Starred timestamp wrong: %v", starred.Starred)
	}

	// IsCompilation
	if first.IsCompilation {
		t.Error("First album IsCompilation should be false")
	}
	if !albums[1].IsCompilation { // "netBloc Vol. 24"
		t.Error("Second album IsCompilation should be true")
	}

	// Artists (OpenSubsonic extension) — single artist
	if len(first.Artists) != 1 {
		t.Fatalf("Expected 1 artist on first album, got %d", len(first.Artists))
	}
	if first.Artists[0].Name != "pornophonique" {
		t.Errorf("Wrong artist name: %s", first.Artists[0].Name)
	}
	if first.Artists[0].ID != "24jpYAI4N8TG0cCAYyzZk5" {
		t.Errorf("Wrong artist ID: %s", first.Artists[0].ID)
	}

	// Genres — empty slice
	if len(first.Genres) != 0 {
		t.Errorf("Expected empty Genres on first album, got %d", len(first.Genres))
	}
	// Genres — single entry
	if len(albums[1].Genres) != 1 {
		t.Fatalf("Expected 1 genre on second album, got %d", len(albums[1].Genres))
	}
	if albums[1].Genres[0].Name != "Netlabel" {
		t.Errorf("Wrong genre name: %s", albums[1].Genres[0].Name)
	}
	// Genres — multiple entries
	last := albums[9] // "Detox Static EP"
	if len(last.Genres) != 2 {
		t.Fatalf("Expected 2 genres on last album, got %d", len(last.Genres))
	}
	if last.Genres[0].Name != "Ebm" || last.Genres[1].Name != "Electro" {
		t.Errorf("Wrong genres on last album: %v", last.Genres)
	}

	// OriginalReleaseDate with full year/month/day — [7] "Retroconnaissance"
	retro := albums[7]
	if retro.Name != "Retroconnaissance" {
		t.Fatalf("Unexpected album at index 7: %s", retro.Name)
	}
	if retro.OriginalReleaseDate == nil {
		t.Fatal("OriginalReleaseDate should not be nil for Retroconnaissance")
	}
	if retro.OriginalReleaseDate.Year == nil || *retro.OriginalReleaseDate.Year != 2006 {
		t.Errorf("Wrong OriginalReleaseDate year: %v", retro.OriginalReleaseDate.Year)
	}
	if retro.OriginalReleaseDate.Month == nil || *retro.OriginalReleaseDate.Month != 5 {
		t.Errorf("Wrong OriginalReleaseDate month: %v", retro.OriginalReleaseDate.Month)
	}
	if retro.OriginalReleaseDate.Day == nil || *retro.OriginalReleaseDate.Day != 14 {
		t.Errorf("Wrong OriginalReleaseDate day: %v", retro.OriginalReleaseDate.Day)
	}

	// ReleaseTypes
	if len(retro.ReleaseTypes) != 1 || retro.ReleaseTypes[0] != "ep" {
		t.Errorf("Wrong ReleaseTypes for Retroconnaissance: %v", retro.ReleaseTypes)
	}
	if len(first.ReleaseTypes) != 0 {
		t.Errorf("Expected empty ReleaseTypes for first album, got %v", first.ReleaseTypes)
	}

	// Unicode album name ("& Everything")
	netbloc42 := albums[5]
	if netbloc42.Name != "netBloc Vol. 42: Live, The Universe & Everything" {
		t.Errorf("Unicode escape in album name not decoded correctly: %s", netbloc42.Name)
	}
}

func runListsTests(client Client, t *testing.T) {
	sampleGenre := getSampleGenre(client)

	t.Run("GetAlbumList", func(t *testing.T) {
		_, err := client.GetAlbumList("foobar", nil)
		if err == nil {
			t.Error("No error was returned with an invalid listType argument")
		}
		_, err = client.GetAlbumList("byYear", nil)
		if err == nil {
			t.Error("Failed to validate byYear parameters")
		}
		_, err = client.GetAlbumList("byYear", map[string]string{"fromYear": "1990"})
		if err == nil {
			t.Error("Failed to validate partial byYear parameters")
		}
		_, err = client.GetAlbumList("byGenre", nil)
		if err == nil {
			t.Error("Failed to validate byGenre parameters")
		}
		albums, err := client.GetAlbumList("random", nil)
		if err != nil {
			t.Error(err)
		}
		if albums == nil {
			t.Error("No albums were returned in a call to random getAlbumList")
		}
		for _, album := range albums {
			if album.Title == "" {
				t.Errorf("Album %#v has an empty name :(", album)
			}
		}
		// Work out genre matching
		albums, err = client.GetAlbumList("byGenre", map[string]string{"genre": sampleGenre.Name})
		if err != nil {
			t.Error(err)
		}
		if albums == nil || len(albums) < 1 {
			t.Error("No albums were returned in a call to a byGenre getAlbumList")
		}
		var empty time.Time
		for _, album := range albums {
			if album.Created == empty {
				t.Errorf("Album %#v has empty created time", album)
			}
		}
	})

	t.Run("GetAlbumList2", func(t *testing.T) {
		// Test incorrect parameters
		_, err := client.GetAlbumList2("foobar", nil)
		if err == nil {
			t.Error("No error was returned with an invalid listType argument")
		}
		_, err = client.GetAlbumList2("byYear", nil)
		if err == nil {
			t.Error("Failed to validate byYear parameters")
		}
		_, err = client.GetAlbumList2("byYear", map[string]string{"fromYear": "1990"})
		if err == nil {
			t.Error("Failed to validate partial byYear parameters")
		}
		_, err = client.GetAlbumList2("byGenre", nil)
		if err == nil {
			t.Error("Failed to validate byGenre parameters")
		}
		// Test with proper parameters
		albums, err := client.GetAlbumList2("newest", nil)
		if err != nil {
			t.Error(err)
		}
		if albums == nil {
			t.Error("No albums were returned in a call to newest getAlbumList2")
		}
		var empty time.Time
		for _, album := range albums {
			if album.Name == "" {
				t.Errorf("Album %#v has an empty name", album)
			}
			if album.Created == empty {
				t.Errorf("Album %#v has empty created time", album)
			}
		}
	})

	t.Run("GetRandomSongs", func(t *testing.T) {
		songs, err := client.GetRandomSongs(nil)
		if err != nil || songs == nil {
			t.Error("Basic call to getRandomSongs failed")
		}
		var empty time.Time
		for _, song := range songs {
			if song.Created == empty {
				t.Errorf("Song %#v had an empty created", song)
			}
		}
		songs, err = client.GetRandomSongs(map[string]string{"size": "1"})
		if len(songs) != 1 {
			t.Errorf("Limiting songs returned by getRandomSongs failed: expected 1, length actual %d", len(songs))
		}
		for _, song := range songs {
			if song.Created == empty {
				t.Errorf("Song %#v had an empty created", song)
			}
		}
	})

	t.Run("GetSongsByGenre", func(t *testing.T) {
		songs, err := client.GetSongsByGenre(sampleGenre.Name, nil)
		if err != nil {
			t.Error(err)
		}
		if songs == nil {
			t.Errorf("No songs returned for genre %v", sampleGenre)
		}
		songs, err = client.GetSongsByGenre(sampleGenre.Name, map[string]string{"count": "1"})
		if err != nil {
			t.Error(err)
		}
		if len(songs) != 1 {
			t.Errorf("Limiting songs returned by GetSongsByGenre failed: expected 1, length actual %d", len(songs))
		}
		var empty time.Time
		for _, song := range songs {
			if song.Created == empty {
				t.Errorf("Song %#v had an empty created", song)
			}
		}
	})

	t.Run("GetNowPlaying", func(t *testing.T) {
		// This test is essentially a no-op because we can't depend on the state of playing something in a test environment
		entries, err := client.GetNowPlaying()
		if err != nil {
			t.Error(err)
		}
		var empty time.Time
		for _, nowPlaying := range entries {
			//t.Logf("NowPlaying %d minutes ago, created %v", nowPlaying.MinutesAgo, nowPlaying.Created.Format("2006-01-02T15:04:05.999999-07:00"))
			if nowPlaying.Created == empty {
				t.Errorf("NowPlayingEntry %#v had an empty created", nowPlaying)
			}
		}
	})

	t.Run("GetStarred", func(t *testing.T) {
		// State dependent test
		_, err := client.GetStarred(nil)
		if err != nil {
			t.Error(err)
		}
		_, err = client.GetStarred2(nil)
		if err != nil {
			t.Error(err)
		}
	})
}
