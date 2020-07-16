package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jpbede/fontmachine/machinery"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	app := &cli.App{
		Name:  "fontmachined",
		Usage: "serve SDF font glyphs on-the-fly",
		Action: func(c *cli.Context) error {
			fm := machinery.NewFontMachinery(machinery.WithFontPath(c.String("path")), machinery.WithFontSize(24))
			router := gin.Default()
			router.GET("/:fontstack/:range", func(context *gin.Context) {
				fontstack := context.Param("fontstack")
				fontrange := strings.Trim(context.Param("range"), ".pbf")
				pbf, err := fm.ComposeFontstack(fontstack, fontrange)
				if err != nil {
					context.String(http.StatusInternalServerError, "a error occurred")
				} else {
					context.Header("Content-Length", strconv.Itoa(len(pbf)))
					context.Data(http.StatusOK, "application/x-protobuf", pbf)
				}
			})
			return router.Run(c.String("listen"))
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				Usage:       "Sets the font path",
				DefaultText: "/fonts",
			},
			&cli.StringFlag{
				Name:        "listen",
				Aliases:     []string{"l"},
				Usage:       "Sets the ip and port listen for",
				DefaultText: ":8080",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
