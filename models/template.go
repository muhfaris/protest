package models

import (
	"encoding/json"
	"fmt"

	"gitlab.sicepat.tech/muhfaris/protest/pkg/utils"
)

const (
	pathTemplateFile = "template/template.json"
)

type Template struct {
	Name    string       `json:"name"`
	Code    string       `json:"code"`
	Payload interface{}  `json:"payload"`
	GenData GenerateData `json:"gen_data"`
}

type Templates []Template

func NewTemplates() (Templates, error) {
	pwd, err := utils.CurrentDir()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/%s", pwd, pathTemplateFile)
	data, err := utils.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var t = Templates{}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return t, nil
}
