package api

import (
	"encoding/json"
	"fmt"
	. "go-server/types"
	"net/http"
	"os"
)

func GetVideos(w http.ResponseWriter, r *http.Request) {
	GetMedia(w, r, "videos")
}

func GetPictures(w http.ResponseWriter, r *http.Request) {
	GetMedia(w, r, "pictures")
}

func GetMedia(w http.ResponseWriter, r *http.Request, folder string) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("All Videos Endpoint Hit")
	entries, err := os.ReadDir(fmt.Sprintf("./public/%s", folder))
    if err != nil {
        fmt.Fprintln(w,err)
		return
    }
	var mediaFolders []MediaFolder
	for _, entry := range entries {
		if entry.IsDir() == false {
			continue
		}
		mediaDir, err := os.ReadDir(fmt.Sprintf("./public/%s/%s", folder, entry.Name()))
		if err != nil {
			fmt.Fprintln(w,err)
			continue
		}
		var media []string
		for _, m := range mediaDir {
			media = append(media, fmt.Sprintf("/%s/%s/%s", folder, entry.Name(), m.Name()))
		}
		
		
		mediaFolders = append(mediaFolders, MediaFolder{
			Name: entry.Name(),
			Files: media,
		})
	}
	x, err := json.Marshal(mediaFolders)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintln(w, string(x))
}