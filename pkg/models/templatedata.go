package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	CSRFToken string
}
