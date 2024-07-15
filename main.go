package main

import (
	myparser "parser/parser"
	"parser/shemas"
)

func main() {
	// создаем инстанс парсера с изначальными значениями
	pars := myparser.New(shemas.Month, shemas.Views, 0, "month")

	pars.Get()

	pars.SetPeriod(shemas.Day)
	pars.SetSort(shemas.Viewers)
	pars.SetOffset(30)
	pars.SetFileName("day")
	pars.Get()

	pars.SetPeriod(shemas.Week)
	pars.SetSort(shemas.Popular)
	pars.SetOffset(50000)
	pars.SetFileName("week")
	pars.Get()
}
