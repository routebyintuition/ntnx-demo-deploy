package main

import (
	"html/template"
	"net/http"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/common-nighthawk/go-figure"
)

// HomeResponse is the response to the home request
type HomeResponse struct {
	BS             string
	Company        string
	CompanyGraphic string
	Emoji          string
	Hacker         string
	Animal         string
	App            string
	Game           string
	Color          string
	Beer           string
	Sentence1      string
	Sentence2      string
	Sentence3      string
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	Info.Println("new visitor from: ", getIP(r))

	company := gofakeit.Company()

	hr := &HomeResponse{
		BS:             gofakeit.BS(),
		Company:        gofakeit.Company(),
		CompanyGraphic: figure.NewFigure(company, "", true).String(),
		Emoji:          gofakeit.Emoji(),
		Hacker:         gofakeit.HackerPhrase(),
		Animal:         gofakeit.Animal(),
		App:            gofakeit.AppName(),
		Game:           gofakeit.Gamertag(),
		Color:          gofakeit.Color(),
		Beer:           gofakeit.BeerName(),
		Sentence1:      gofakeit.Sentence(7),
		Sentence2:      gofakeit.Sentence(6),
		Sentence3:      gofakeit.Sentence(5),
	}

	homeSource, _ := template.ParseFiles("web/welcome.html")
	homeSource.Execute(w, hr)
	return
}
