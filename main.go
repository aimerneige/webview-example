package main

import (
	"webview-example/bind"
	"webview-example/config"
	"webview-example/database"
	"webview-example/web"

	"github.com/webview/webview"
)

const (
	title  = "Webview Example"
	width  = 1024
	height = 768
	port   = ":8080"
)

func main() {
	config.InitConfig()
	database.InitDatabase()

	go web.StartWebServer(port)

	debug := true
	w := webview.New(debug)
	defer w.Destroy()

	w.SetTitle(title)
	w.SetSize(width, height, webview.HintNone)

	bind.AllBindCollection(w)
	w.Navigate("http://localhost" + port)
	w.Run()
}
