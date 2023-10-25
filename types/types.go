package types

type Comic struct {
	Name     string
	Cover    string
	Chapters []Chapter
}

type Chapter struct {
	Name   string
	Images []string
}

type PictureFolder struct {
	Name string
	Files []string
}

type Video struct {
	Name     string
	Url      string
	Category string
	// Thumbnail string
}