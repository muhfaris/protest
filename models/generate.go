package models

type GenerateData struct {
	Field   string          `json:"field"`
	Options GenerateOptions `json:"options"`
}

type GenerateOption struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GenerateOptions []GenerateOption
