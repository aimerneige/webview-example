package web

import "github.com/gin-gonic/gin"

func StartWebServer(port string) {
	r := gin.Default()

	r.LoadHTMLFiles(
		"./static/index.html",
		"./static/multi.html",
		"./static/image.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/multi", func(c *gin.Context) {
		c.HTML(200, "multi.html", nil)
	})
	r.GET("/image", func(c *gin.Context) {
		src := c.Query("src")
		alt := c.Query("alt")
		c.HTML(200, "image.html", gin.H{
			"src": src,
			"alt": alt,
		})
	})

	r.Static("/img", "./static/img")

	panic(r.Run(port))
}
