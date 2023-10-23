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

type MediaFolder struct {
	Name string
	Files []string
}