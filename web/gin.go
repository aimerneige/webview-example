package web

import "github.com/gin-gonic/gin"

func StartWebServer(port string) {
	r := gin.Default()

	r.LoadHTMLFiles(
		"./static/index.html",
		"./static/multi.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/multi", func(c *gin.Context) {
		c.HTML(200, "multi.html", nil)
	})

	r.Static("/img", "./static/img")

	panic(r.Run(port))
}
