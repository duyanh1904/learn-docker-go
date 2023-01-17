package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"io"
	"net/http"
)

type CurlController struct{}

func (h CurlController) GetListPartner(c *gin.Context) {
	headers := req.Header{
		"Content-Type": "application/json",
		"token":        "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6InNCeCJ9.eyJpc3MiOiJrdnNzand0Iiwic3ViIjo2NzY3MiwiaWF0IjoxNjczODM2MzY2LCJleHAiOjE2NzY0MjgzNjYsInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluIiwicm9sZXMiOlsiVXNlciJdLCJrdnVzZXRmYSI6MCwia3Z3YWl0b3RwIjowLCJrdnNlcyI6ImRiZTljYTRlMDFlYjQ2NmY4YzIwMWMxNDAzYmY3NzM4Iiwia3Z1aWQiOjY3NjcyLCJrdnV0eXBlIjowLCJrdnVsaW1pdCI6IkZhbHNlIiwia3Z1YWRtaW4iOiJUcnVlIiwia3Z1YWN0IjoiVHJ1ZSIsImt2dWxpbWl0dHJhbnMiOiJGYWxzZSIsImt2dXNob3dzdW0iOiJUcnVlIiwia3ZiaWQiOjcyNDAsImt2cmluZGlkIjo3LCJrdnJjb2RlIjoiYmluaGR0NiIsImt2cmlkIjoxMDE2MTc2LCJrdnVyaWQiOjEwMTYxNzYsImt2cmdpZCI6MSwicGVybXMiOiIifQ.d8ZhPZ-9e60NYzr4HgkQAlw5u0EKnDrwEzRnpkalMhPFKji-U9Sv11ALaCn7pPYu0jOijCV8Itbt5cEUFrGhStZvMmW8wAZYQzgiBLkUNCyoVKBW_-02Mb5lhb_o2PeNPN7P7UDGa1BVY5gGAqlvJleKJYYJ9S1GxWWE21b22g9JiFW70aAcxb9MosAh4qxfkcx2OYazpkmTfQOiZWmjkmcRhhWKaoXGjL_IKH2rGP9fcULQ7BxjHrvXBukAlnIECp5CFUvsGiyJ4cRzCU93TlB6_wntEVgbHDzzcL3ecudaMOlOaHgH3HqZxf2WZ4j0SdJekVuUHBvWOsH3jBYJ9g",
		"branch-id":    "7240",
	}
	resp, err := req.Get("http://host.docker.internal/api/v3/partner-delivery", headers)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError, "Error making request to Laravel server")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Response().Body)

	c.String(http.StatusOK, resp.String())
}
