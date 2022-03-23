package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/file/:digest", download)
	server.GET("/image/:digest", download)

	err := server.Run()
	if err != nil {
		return
	}
}

func download(context *gin.Context) {
	resp, err := http.Get("https://oldupload.naimi.me" + context.Request.RequestURI)
	if err != nil {
		context.String(http.StatusNotFound, err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	context.Data(200, resp.Header.Get("Content-Type"), body)
}
