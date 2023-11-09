package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)


func main() {
    fmt.Print("\nEnter the command to execute: ")
    cmd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    cmd = strings.TrimSpace(cmd)

    fmt.Print("\nEnter Playlist name: ")
    playlist, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    playlist = strings.TrimSpace(playlist)

    cmdenc := encodeCommand(cmd)

    songs, err := load()
    if err != nil {
        fmt.Println(err)
        return
    }

    payloadTitles, URI := buildplaylist(songs, cmdenc)

    fmt.Println("\n\nPayload Playlist")
    for _, t := range payloadTitles {
        fmt.Println(t)
    }

    buildPayload(playlist, URI)
    // addTracksToPlaylist(token, playlistID, URI)

}

func load() ([]map[string]interface{}, error)  {
    file, err := os.ReadFile("songs/rock.json")

    if err != nil {
        return nil, err
    }

    var songs []map[string]interface{}

    json.Unmarshal(file, &songs)
    
    return songs, nil
}

func buildplaylist(songs []map[string]interface{}, cmdenc string) ([]string, []string) {
    var payloadTitles []string
    var URI []string

    for _, c := range cmdenc {
        var track []map[string]interface{}
        for _, s := range songs {
            title, ok := s["title"].(string)
            if !ok {
                continue
            }  
            if strings.ToLower(string(title[0])) == strings.ToLower(string(c)) {
                track = append(track, s)
            }
        }
        selected := track[rand.Intn(len(track))]
        title, _ := selected["title"].(string)
        uri, _ := selected["uri"].(string)
        payloadTitles = append(payloadTitles, title)
        URI = append(URI, uri)
    }
    return payloadTitles, URI
}



func encodeCommand(command string) string {
	replacer := strings.NewReplacer(
		" ", "space",
		"-", "hiphen",
		"/", "fslash",
		"\\", "bslash",
		"!", "ebang",
		"#", "epound",
		"$", "edollar",
		"@", "eatsym",
		"%", "eperc",
		"^", "ecarr",
		"&", "eand",
		"*", "estar",
		"(", "eopar",
		")", "ecpar",
		"+", "eplus",
		"=", "eeq",
		".", "edot",
		",", "ecoma",
		"?", "eques",
		"1", "eone",
		"2", "etwo",
		"3", "ethree",
		"4", "efour",
		"5", "efive",
		"6", "esix",
		"7", "eseven",
		"8", "eeight",
		"9", "enine",
		"0", "ezero",
	)
	return replacer.Replace(command)
}