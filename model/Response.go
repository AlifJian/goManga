package model

type ResponseManga struct {
	Status  int
	Message string
	Data    []Manga
}

type ResponseSeries struct {
	Status  int
	Message string
	Data    Series
}
