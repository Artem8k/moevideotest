package shemas

// Структура json ответа
type Response struct {
	Project_count int      `json:"project_count"` // кол-во сайтов в топе
	Result        []Result `json:"result"`
}

// Структура с названием сайта и вложенной структурой с рейтингом
type Result struct {
	Name   string `json:"name"` // название
	Rating Rating `json:"rating"`
}

type Rating struct {
	Popularity Popularity `json:"popularity"`
	Viewers    int        `json:"viewers"` // посетители
	Views      int        `json:"views"`   // просмотры
}

type Popularity struct {
	Current  int `json:"current"`  // текущая популярность
	Previous int `json:"previous"` // популярность за пред. период?
}
