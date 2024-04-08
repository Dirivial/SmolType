package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/Dirivial/SmolType/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.templ")),
	}
}

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	e.Static("/images", "images")
	e.Static("/css", "css")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", render(c, views.Home("Alexander")))
	})

	fmt.Println("Listening on :8000")
	e.Logger.Fatal(e.Start(":8000"))
}
