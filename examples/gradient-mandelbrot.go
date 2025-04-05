package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
)

const (
	outFile       = "gradient-mandelbrot.png"
	width, height = 1024, 512
	marginSize    = 16
)

func main() {
	dc := gg.NewContext(width, height)

	grad1 := gg.NewLinearGradient(0, height/2, width, height/2)
	grad1.AddColorStop(0.0, colors.RGBAF{1.0, 1.0, 1.0, 1.0})
	grad1.AddColorStop(0.15, colors.RGBAF{1.0, 0.8, 0.0, 1.0})
	grad1.AddColorStop(0.33, colors.RGBAF{0.52, 0.121, 0.074, 1.0})
	grad1.AddColorStop(0.67, colors.RGBAF{0.0, 0.0, 0.6, 1.0})
	grad1.AddColorStop(0.85, colors.RGBAF{0.0, 0.384, 1.0, 1.0})
	grad1.AddColorStop(1.0, colors.RGBAF{1.0, 1.0, 1.0, 1.0})

	grad2 := gg.NewLinearGradient(0, height/2, width, height/2)
	grad2.AddColorStop(0.0, colors.RGBAF{0.7909, 0.9961, 0.7630, 1.0})
	grad2.AddColorStop(0.16, colors.RGBAF{0.8974, 0.8953, 0.6565, 1.0})
	grad2.AddColorStop(0.33, colors.RGBAF{0.9465, 0.3161, 0.1267, 1.0})
	grad2.AddColorStop(0.5, colors.RGBAF{0.5184, 0.1109, 0.0917, 1.0})
	grad2.AddColorStop(0.66, colors.RGBAF{0.0198, 0.4563, 0.6839, 1.0})
	grad2.AddColorStop(0.83, colors.RGBAF{0.5385, 0.8259, 0.8177, 1.0})
	grad2.AddColorStop(1.0, colors.RGBAF{0.7909, 0.9961, 0.7630, 1.0})

	grad3 := gg.NewLinearGradient(0, height/2, width, height/2)
	grad3.AddColorStop(0.0, colors.RGBAF{0.80585, 0.81648, 0.8218, 1.0})
	grad3.AddColorStop(0.18, colors.RGBAF{0.43882, 0.52393, 1.0, 1.0})
	grad3.AddColorStop(0.42, colors.RGBAF{1.0, 0.35904, 0.58244, 1.0})
	grad3.AddColorStop(0.63, colors.RGBAF{1.0, 1.0, 0.52127, 1.0})
	grad3.AddColorStop(0.86, colors.RGBAF{0.54787, 0.93351, 0.56914, 1.0})
	grad3.AddColorStop(1.0, colors.RGBAF{0.80585, 0.81648, 0.8218, 1.0})

	grad4 := gg.NewLinearGradient(0, height/2, width, height/2)
	grad4.AddColorStop(0.0, colors.RGBAF{0.6595, 0.0, 0.0, 1.0})
	grad4.AddColorStop(0.18, colors.RGBAF{0.0, 0.3058, 0.5877, 1.0})
	grad4.AddColorStop(0.39, colors.RGBAF{0.8164, 0.4148, 0.0718, 1.0})
	grad4.AddColorStop(0.57, colors.RGBAF{0.0, 0.4867, 0.1648, 1.0})
	grad4.AddColorStop(0.78, colors.RGBAF{0.2978, 0.1382, 0.75, 1.0})
	grad4.AddColorStop(1.0, colors.RGBAF{0.6595, 0.0, 0.0, 1.0})

	for row := 0; row < 4; row++ {
		for col := 0; col < 8; col++ {
			x := float64(marginSize + col*(marginSize+108))
			y := float64(marginSize + row*(marginSize+108))
			dc.DrawRectangle(x, y, 108, 108)
			dc.Fill()
		}
	}
	mask := dc.AsMask()
	dc.Clear()

	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, 1024, 128)
	dc.SetFillStyle(grad1)
	dc.Fill()
	dc.DrawRectangle(0, 128, 1024, 128)
	dc.SetFillStyle(grad2)
	dc.Fill()
	dc.DrawRectangle(0, 256, 1024, 128)
	dc.SetFillStyle(grad3)
	dc.Fill()
	dc.DrawRectangle(0, 384, 1024, 128)
	dc.SetFillStyle(grad4)
	dc.Fill()

	dc.SavePNG(outFile)
}
