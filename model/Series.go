package model

type Series struct {
	Title          string
	Cover          string
	Alias          string
	Rank           string
	Authors        string
	Genres         string
	OriginLang     string
	TranslatedLang string
	Status         string
	Release        string
	Sinopsis       string
	Chapter        []struct {
		Title     string
		ChapterId string
	}
}
