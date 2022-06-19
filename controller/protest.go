package controller

import (
	"log"

	"gitlab.sicepat.tech/muhfaris/protest/models"
)

type protest struct {
	url string
}

func NewProtest(address string) *protest {
	return &protest{url: address}
}

func (this *protest) Parse() {
	templates, err := models.NewTemplates()
	if err != nil {
		log.Printf("error parse template = %+v\n", err)
		return
	}

	for _, template := range templates {
		log.Println("Name:", template.Name)
	}
}
