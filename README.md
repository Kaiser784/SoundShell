# SoundShell

Spotify C2 in golang

## Introduction

SoundShell is a Command-and-Control (C2) tool built using the Go programming language, leveraging the Spotify Web API for its operations. It allows users to execute commands and generate playlists on Spotify based on encoded commands.

### Features

- Execute custom commands and map them to Spotify playlists.
- Encode commands into a format suitable for playlist generation.
- Generate playlists using random selection of tracks based on command encoding.

## Installation

To install SoundShell, ensure you have Go installed on your machine, then run the following command:

```bash
go get github.com/Kaiser784/SoundShell
```

## Usage

1. Run the main program:
```bash
go run main.go
```
2. Enter the command to execute: You will be prompted to enter a command, which will be encoded and used to generate a playlist.
3. Enter the playlist name: Provide a name for the new playlist.
4. View the generated playlist: The program will output the titles of the tracks included in the generated playlist.

## Encoding Example
### Command: `echo file.txt`

1. Command Input:
```bash
Enter the command to execute: echo file.txt
```
2. Encoding the Command:
The command `echo file.txt` is encoded using the `encodeCommand` function in the following way:

```bash
echospacefileedottxt
```

This is achieved by replacing each special character with a predefined string:
` `(space) -> space
`.` (dot) -> edot

Generating the Playlist:
The encoded command is then used to select songs from a predefined list (`songs/rock.json`). Each character of the encoded command is matched to the first character of the song titles.

For example:
- For e, a song starting with E is selected.
- For c, a song starting with C is selected.
- For h, a song starting with H is selected.
And so on...

Output:
The playlist generated from the encoded command might look like this:
```bash
Payload Playlist
- Echoes
- Come Together
- Hey Jude
- One
- Firework
- Imagine
- Like a Rolling Stone
- Enter Sandman
- Tom Sawyer
- X-Ray
- Titanium
```
The actual songs selected will vary as they are chosen randomly from the matching pool of songs for each character.

## Dependencies

The project dependencies are managed via Go modules, as listed in the go.mod file:

    github.com/spf13/cobra v1.8.0
    github.com/zmb3/spotify v1.3.0
    golang.org/x/oauth2 v0.13.0

## License

This project is licensed under the MIT License.
