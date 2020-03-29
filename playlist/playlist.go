package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type song struct {
	title    string
	path     string
	duration int
}

func main() {
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".m3u") {
		fmt.Printf("usage: <%s> filename1.m3u > filename2.pls\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		songs := parseM3Ufile(string(rawBytes))
		generatePLS(songs)
	}
}

func parseM3Ufile(rawData string) (songs []song) {
	var songData song
	for _, line := range strings.Split(rawData, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasSuffix(line, "#EXTM3U") {
			continue
		}

		if strings.HasPrefix(line, "#EXTINF") {
			songData.duration, songData.title = parseEXTINF(line)
		} else {
			songData.path = line
		}

		if songData.duration != 0 && songData.path != "" && songData.title != "" {
			songs = append(songs, songData)
			songData = song{}
		}
	}
	return songs
}

func generatePLS(songs []song) {
	fmt.Println("[playlist]")
	for index, s := range songs {
		index++
		fmt.Printf("File%v=%s\n", index, s.path)
		fmt.Printf("Title%v=%s\n", index, s.title)
		fmt.Printf("Length%v=%v\n", index, s.duration)
	}
	fmt.Printf("NumberOfEntries=%v\nVersion=2", len(songs))
}

func parseEXTINF(line string) (duration int, title string) {
	if i := strings.Index(line, ":"); i > -1 {
		line = line[i+1:]
		const separator = ","
		if j := strings.Index(line, ","); j > -1 {
			title = line[j+len(separator):]
			var err error
			if duration, err = strconv.Atoi(line[:j]); err != nil {
				log.Printf("Duration cannot be parsed for %s : %v", title, err)
				duration = -1
			}
		}
	}
	return duration, title
}
