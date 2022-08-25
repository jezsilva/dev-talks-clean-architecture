package adapter_web

import (
	"devtalks/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinHttpAdapter(fn controller.FuncController) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := c.GetRawData()
		headers := make(map[string][]string)
		params := make(map[string][]string)

		for key, value := range c.Request.Header {
			headers[key] = value
		}

		for key, value := range c.Request.URL.Query() {
			params[key] = value
		}

		for _, value := range c.Params {
			params[value.Key] = []string{value.Value}
		}

		response, err := fn(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, response)
	}
}
