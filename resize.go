package canvasresizer

import (
	"fmt"
	"image"
	"image/color"
	"io"

	"github.com/disintegration/imaging"
)

type ResizeCanvasInput struct {
	Src         io.Reader
	Dst         io.Writer
	WidthRatio  float64
	HeightRatio float64
}

func ResizeCanvas(in ResizeCanvasInput) error {
	img, err := imaging.Decode(in.Src)
	if err != nil {
		return fmt.Errorf("decode source image: %w", err)
	}

	bounds := img.Bounds()
	canvas := imaging.New(int(float64(bounds.Dx())*in.WidthRatio), int(float64(bounds.Dy())*in.HeightRatio), color.Black)

	bgBounds := canvas.Bounds()
	bgW := bgBounds.Dx()
	bgMinX := bgBounds.Min.X
	centerX := bgMinX + bgW/2
	x0 := centerX - img.Bounds().Dx()/2
	out := imaging.Paste(canvas, img, image.Pt(x0, 0))

	if err := imaging.Encode(in.Dst, out, imaging.PNG, imaging.PNGCompressionLevel(0)); err != nil {
		return fmt.Errorf("encode image: %w", err)
	}
	return nil
}
