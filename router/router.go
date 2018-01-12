package router

import (
	"exercise/eaciit/controller"
	"exercise/eaciit/variable"
	"os"
	"path/filepath"

	knot "github.com/eaciit/knot/knot.v1"
)

type Router struct {
}

func (r Router) Init() {

	basePath, _ := os.Getwd()

	app := knot.NewApp("Hello")
	app.ViewsPath = variable.WebHTMLDir
	app.Static("static", filepath.Join(basePath, "views"))
	app.Register(new(controller.EaciitController))
	app.Register(new(controller.DBoxController))
	// app.Register(new(controller.PerlinController))
	knot.RegisterApp(app)

	knot.StartApp(app, variable.WebHost+":"+variable.WebPort)
}
