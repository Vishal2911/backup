package server

import (
	"github.com/gin-gonic/gin"
)

func (server Server) readQueryParams(c *gin.Context)map[string]string {
	queryParams := c.Request.URL.Query()
	queryes := make(map[string]string)
	for key, values := range queryParams {
		// paramsValuse := []string{}
		for _, value := range values {
			// paramsValuse = append(paramsValuse, value)
			queryes[key] = value
		}
		// queryes[key] = values
	}
	return queryes
}
