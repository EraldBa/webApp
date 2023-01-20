package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	Stats           *StatsSend
	CSRFToken       string
	Success         string
	Error           string
	Flash           string
	IsAuthenticated int
}
