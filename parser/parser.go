package myparser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"parser/shemas"
)

// Методы которые парсер уметь делать
type Parser interface {
	Get()
	SetPeriod(shemas.Period)
	SetSort(shemas.Sort)
	SetOffset(int)
	SetFileName(string)
}

type RamblerParser struct {
	Period   shemas.Period
	Sort     shemas.Sort
	Offset   int
	FileName string
}

func New(p shemas.Period, s shemas.Sort, o int, f string) Parser {
	return &RamblerParser{
		Period:   p,
		Sort:     s,
		Offset:   o,
		FileName: f,
	}
}

func (r *RamblerParser) Get() {
	// собираем строку url
	url := fmt.Sprintf("https://top100.rambler.ru/api/catalogue/v2.0/table?period=%s&sort=%s&offset=%d", r.Period, r.Sort, r.Offset)

	// Получаем ответ от rambler
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Читаем из тела ответа
	s, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Анмаршалим json из ответа в структуру
	sr := shemas.Response{}
	err = json.Unmarshal(s, &sr)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем файлик в который будем записывать
	fileName := fmt.Sprintf("%s.txt", r.FileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Пробегаемся по всему json'у, форматируем в нужный формат и записываем в созданный файл
	for _, e := range sr.Result {
		s := fmt.Sprintf("Название: %s\n Посетители: %d\n Просмотры: %d\n Популярность: %d\n\n", e.Name, e.Rating.Viewers, e.Rating.Views, e.Rating.Popularity.Current)

		_, err := file.WriteString(s)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (r *RamblerParser) SetPeriod(s shemas.Period) {
	r.Period = s
}

func (r *RamblerParser) SetSort(s shemas.Sort) {
	r.Sort = s
}

func (r *RamblerParser) SetOffset(o int) {
	r.Offset = o
}

func (r *RamblerParser) SetFileName(f string) {
	r.FileName = f
}
