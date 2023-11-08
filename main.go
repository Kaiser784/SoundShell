package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
    "github.com/zmb3/spotify"
    "golang.org/x/oauth2/clientcredentials"
)

const (
    spotifyClientID     = "your-spotify-client-id"
    spotifyClientSecret = "your-spotify-client-secret"
)

var rootCmd = &cobra.Command{
    Use:   "spotifycli",
    Short: "spotifycli controls Spotify through command line",
}

func init() {
    rootCmd.AddCommand(createPlaylistCmd)
    // Add other commands here
}

var createPlaylistCmd = &cobra.Command{
    Use:   "create-playlist [name] [public]",
    Short: "Creates a new Spotify playlist",
    Args:  cobra.MinimumNArgs(2),
    Run:   createPlaylist,
}

func createPlaylist(cmd *cobra.Command, args []string) {
    playlistName := args[0]
    public := args[1] == "true"

    client := getSpotifyClient()
    user, err := client.CurrentUser()
    if err != nil {
        log.Fatalf("couldn't get current user: %v", err)
    }

    playlist, err := client.CreatePlaylistForUser(user.ID, playlistName, "Created by CLI", public)
    if err != nil {
        log.Fatalf("couldn't create playlist: %v", err)
    }

    fmt.Printf("Created playlist! ID: %v\n", playlist.ID)
}

func getSpotifyClient() spotify.Client {
    config := &clientcredentials.Config{
        ClientID:     spotifyClientID,
        ClientSecret: spotifyClientSecret,
        TokenURL:     spotify.TokenURL,
    }

    token, err := config.Token(context.Background())
    if err != nil {
        log.Fatalf("couldn't get token: %v", err)
    }

    httpClient := spotify.Authenticator{}.NewClient(token)
    return spotify.NewClient(&httpClient)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
