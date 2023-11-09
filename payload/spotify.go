package main

import (
	_ "context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/zmb3/spotify"
	_ "golang.org/x/oauth2"
)

var (
	auth  spotify.Authenticator
	ch    = make(chan *spotify.Client)
	state = "some-random-state"
)

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

func buildPayload(playlist string, URI []string) {
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	auth = spotify.NewAuthenticator("http://localhost:8080/callback", spotify.ScopePlaylistModifyPublic)
	auth.SetAuthInfo(spotifyClientID, spotifyClientSecret)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// Create a new playlist for a user.
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatalf("couldn't get current user: %v", err)
	}

	playlistID, err := client.CreatePlaylistForUser(user.ID, playlist, playlist, true)
	if err != nil {
		log.Fatalf("couldn't create playlist: %v", err)
	}

	fmt.Println("Created playlist:", playlistID.Name)

	// List of track IDs to add to the playlist.
	var trackIDs []spotify.ID
	for _, uri := range URI {
		// Extract the ID from the URI.
		parts := strings.Split(uri, ":")
		if len(parts) == 3 && parts[1] == "track" {
			trackIDs = append(trackIDs, spotify.ID(parts[2]))
		} else {
			fmt.Printf("Invalid Spotify URI: %s\n", uri)
		}
	}

	// Add tracks to the playlist.
	_, err = client.AddTracksToPlaylist(playlistID.ID, trackIDs...)
	if err != nil {
		log.Fatalf("couldn't add tracks to playlist: %v", err)
	}

	fmt.Println("Added tracks to playlist: ", playlistID.ID)
}
