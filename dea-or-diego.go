package main

import (
  "github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	//remove the following line to not have your deployment tracker
	router := gin.Default()

	cfInstanceGUID := os.Getenv("CF_INSTANCE_GUID")
		if cfInstanceGUID == "" {
	  router.LoadHTMLGlob("templates/*")

		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "dea.tmpl", gin.H{
            "title": "DEA",
        })
		})
	} else {
		router.LoadHTMLGlob("templates/*")
		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "diego.tmpl", gin.H{
            "title": "DIEGO",
        })
		})
	}


	port := os.Getenv("VCAP_APP_PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
