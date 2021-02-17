package api

import "github.com/gin-gonic/gin"

type response struct {
	Data   interface{} `json:"data"`
	Status status      `json:"status"`
}

type status struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func (r response) send(c *gin.Context) {
	c.JSON(r.Status.Code, r)
}

func newRes() response {
	return response{
		Status: status{
			Code: 200,
		},
	}
}
