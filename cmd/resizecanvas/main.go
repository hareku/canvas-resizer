package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	canvasresizer "github.com/hareku/canvas-resizer"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "resizecanvas",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "src",
				Aliases:  []string{"s"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "width-ratio",
				Aliases: []string{"width", "w", "wr"},
				Value:   "1",
			},
			&cli.StringFlag{
				Name:    "height-ratio",
				Aliases: []string{"height", "hr"}, // "h" is already used by help
				Value:   "1",
			},
		},
		Action: func(cCtx *cli.Context) error {
			wrStr := cCtx.String("width-ratio")
			wr, err := strconv.ParseFloat(wrStr, 64)
			if err != nil {
				return fmt.Errorf("parse width ratio: %w", err)
			}
			hrStr := cCtx.String("height-ratio")
			hr, err := strconv.ParseFloat(hrStr, 64)
			if err != nil {
				return fmt.Errorf("parse height ratio: %w", err)
			}

			src, err := os.Open(cCtx.String("src"))
			if err != nil {
				return fmt.Errorf("open source file: %w", err)
			}
			defer src.Close()

			dstName := fmt.Sprintf("%s_w%s_h%s.png", cCtx.String("src"), wrStr, hrStr)
			dst, err := os.Create(dstName)
			if err != nil {
				return fmt.Errorf("create destination file: %w", err)
			}
			defer dst.Close()

			if err := canvasresizer.ResizeCanvas(canvasresizer.ResizeCanvasInput{
				Src:         src,
				Dst:         dst,
				WidthRatio:  wr,
				HeightRatio: hr,
			}); err != nil {
				return fmt.Errorf("resize canvas: %w", err)
			}
			slog.Info("resize canvas", "src", cCtx.String("src"), "dst", dstName)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
