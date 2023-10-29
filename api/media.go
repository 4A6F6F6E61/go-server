package api

import (
	"encoding/json"
	"fmt"
	. "go-server/types"
	"log"
	"net/http"
	"os"
)

func GetVideos(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /api/videos")
	w.Header().Set("Content-Type", "application/json")

	entries, err := os.ReadDir("./public/videos")
    if err != nil {
        fmt.Fprintln(w,err)
		return
    }
	var retVideos []Video
	for _, cat := range entries {
		if !cat.IsDir() {
			continue
		}
		videos, err := os.ReadDir(fmt.Sprintf("./public/videos/%s", cat.Name()))
		if err != nil {
			continue
		}
		
		for _, video := range videos {
			files, err := os.ReadDir(fmt.Sprintf("./public/videos/%s/%s", cat.Name(), video.Name()))
			if err != nil {
				continue
			}
			if len(files) == 0 {
				log.Println("No video found")
				continue
			}
			firstVideo := files[0].Name()
			retVideos = append(retVideos, Video{
				Name: video.Name(),
				Url: fmt.Sprintf("/videos/%s/%s/%s", cat.Name(), video.Name(), firstVideo),
				Category: cat.Name(),
			})
		}
	}
	x, err := json.Marshal(retVideos)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintln(w, string(x))
}

func GetPictures(w http.ResponseWriter, r *http.Request) {
	folder := "pictures"
	w.Header().Set("Content-Type", "application/json")

	entries, err := os.ReadDir(fmt.Sprintf("./public/%s", folder))
    if err != nil {
		return
    }
	var pictureFolders []PictureFolder
	for _, entry := range entries {
		if entry.IsDir() == false {
			continue
		}
		mediaDir, err := os.ReadDir(fmt.Sprintf("./public/%s/%s", folder, entry.Name()))
		if err != nil {
			continue
		}
		var media []string
		for _, m := range mediaDir {
			media = append(media, fmt.Sprintf("/%s/%s/%s", folder, entry.Name(), m.Name()))
		}
		
		
		pictureFolders = append(pictureFolders, PictureFolder{
			Name: entry.Name(),
			Files: media,
		})
	}
	x, err := json.Marshal(pictureFolders)
    if err != nil {
        return
    }
	fmt.Fprintln(w, string(x))
}