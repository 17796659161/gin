package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:=gin.Default()
	r.GET("/test", func(c *gin.Context) {
		first_name := c.Query("first_name")
		last_name := c.DefaultQuery("last_name","last_default_name")
		c.String(http.StatusOK,"%s,%s",first_name,last_name)
	})
	r.Run(":8080")
}
