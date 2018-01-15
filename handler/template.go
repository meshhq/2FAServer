package handler

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type HTMLTemplate struct {
	Templates *template.Template
}

func (t *HTMLTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
