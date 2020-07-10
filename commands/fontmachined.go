package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/jpbede/fontmachine"
	"net/http"
	"strings"
)

func main() {
	fm := fontmachine.NewFontMachinery(fontmachine.WithFontPath("./fonts"))
	router := gin.Default()
	router.GET("/:fontstack/:range", func(context *gin.Context) {
		fontstack := context.Param("fontstack")
		fontrange := strings.Trim(context.Param("range"), ".pbf")
		pbf, err := fm.ComposeFontstack(fontstack, fontrange)
		spew.Dump(err)
		context.Data(http.StatusOK, "application/x-protobuf", pbf)
	})
	router.Run(":8080")
}
