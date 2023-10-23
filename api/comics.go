package api

import (
	"encoding/json"
	"fmt"
	. "go-server/types"
	"net/http"
	"os"
)


func GetComics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("All Comics Endpoint Hit")
	entries, err := os.ReadDir("./public/comics")
    if err != nil {
        fmt.Fprintln(w,err)
		return
    }
	var comics []Comic
	for _, entry := range entries {
		if entry.IsDir() == false {
			continue
		}
		chaptersDir, err := os.ReadDir(fmt.Sprintf("./public/comics/%s", entry.Name()))
		if err != nil {
			fmt.Fprintln(w,err)
			return
		}
		var chapters []Chapter
		for _, chapter := range chaptersDir {
			if chapter.IsDir() == false {
				continue
			}
			imagesDir, err := os.ReadDir(fmt.Sprintf("./public/comics/%s/%s", entry.Name(), chapter.Name()))
			if err != nil {
				fmt.Fprintln(w,err)
				continue
			}
			var images []string
			for _, image := range imagesDir {
				images = append(images, fmt.Sprintf("/comics/%s/%s/%s", entry.Name(), chapter.Name(), image.Name()))
			}
			chapters = append(chapters, Chapter{
				Name: chapter.Name(),
				Images: images,
			})
		}
		comics = append(comics, Comic{
			Name: entry.Name(),
			Cover: fmt.Sprintf("/comics/%s/cover.jpg", entry.Name()),
			Chapters: chapters,
		})
	}
	c, err := json.Marshal(comics)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintln(w, string(c))
}
