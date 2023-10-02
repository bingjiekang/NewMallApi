package main

import (
	mallrouter "NewMallApi/router/mallRouter"
	// mallservice "NewMallApi/service/mallService"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./static/assets")
	router.LoadHTMLGlob("template/html/*")
	mallrouter.Router(&router.RouterGroup)
	router.Run(":8083")
}
