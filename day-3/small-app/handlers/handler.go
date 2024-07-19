package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"small-app/models"
)

type handler struct {
	models.Conn
}

func API() *gin.Engine {

	r := gin.New()

	//apply middleware to all the endpoints using r.Use
	c := models.NewConn()
	h := handler{Conn: c}
	r.GET("/check", h.check)
	r.POST("/signup", h.Signup)

	return r
}

func (h handler) check(c *gin.Context) {

	u := struct {
		Status string
	}{
		Status: "Ok",
	}

	//JSON serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".
	c.JSON(http.StatusOK, u)

}
