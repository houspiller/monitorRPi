package route

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/houspiller/monitorRPi/monitor"
)

//CreateRouters cria as rotas
func CreateRouters() *gin.Engine {

	r := gin.Default()

	r.Use(cors.Default())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", ping)
		v1.GET("/temperatura", getTemperatura)
	}

	return r
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func getTemperatura(c *gin.Context) {
	monitor.ExecCmd("")
}
