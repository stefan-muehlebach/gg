package main

import (
	"github.com/stefan-muehlebach/gg"
	"github.com/stefan-muehlebach/gg/colors"
	"github.com/stefan-muehlebach/gg/fonts"
)

const TEXT = "Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off—then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me."

const (
	W = 1024
	H = 1024
	P = 16
	outFile = "wrap.png"
)

var (
	BackColor = colors.RGBAF{0.851, 0.811, 0.733, 1.0}
	LineColor = colors.RGBAF{0.153, 0.157, 0.133, 1.0}
)

func main() {
	dc := gg.NewContext(W, H)
    face1, _ := fonts.NewFace(fonts.GoBold, 14.0)
    face2, _ := fonts.NewFace(fonts.LucidaBright, 10.0)

	dc.SetFillColor(BackColor)
	dc.Clear()
	dc.DrawLine(W/2, 0, W/2, H)
	dc.DrawLine(0, H/2, W, H/2)
	dc.DrawRectangle(P, P, W-P-P, H-P-P)
	dc.SetStrokeColor(colors.RGBAF{0, 0, 1.0, 0.25})
	dc.SetStrokeWidth(3)
	dc.Stroke()

	dc.SetTextColor(LineColor)
	dc.SetFontFace(face1)
	dc.DrawStringWrapped("TOP LEFT", P, P, 0, 0, W, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("TOP CENTER", W/2, P, 0.5, 0, 0, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped("TOP RIGHT", W-P, P, 1, 0, W, 1.5, gg.AlignRight)
	dc.DrawStringWrapped("MIDDLE LEFT", P, H/2, 0, 0.5, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("MIDDLE CENTER", W/2, H/2, 0.5, 0.5, 0, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped("MIDDLE RIGHT", W-P, H/2, 1, 0.5, 0, 1.5, gg.AlignRight)
	dc.DrawStringWrapped("BOTTOM LEFT", P, H-P, 0, 1, W, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("BOTTOM CENTER", W/2, H-P, 0.5, 1, 0, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped("BOTTOM RIGHT", W-P, H-P, 1, 1, W, 1.5, gg.AlignRight)

	dc.SetFontFace(face2)
	dc.DrawStringWrapped(TEXT, W/2-P, H/2-P, 1, 1, W/2-3*P, 1.5, gg.AlignRight)
	dc.DrawStringWrapped(TEXT, W/2+P, H/2-P, 0, 1, W/2-3*P, 1.75, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2-P, H/2+P, 1, 0, W/2-3*P, 2.0, gg.AlignRight)
	dc.DrawStringWrapped(TEXT, W/2+P, H/2+P, 0, 0, W/2-3*P, 2.25, gg.AlignLeft)

	dc.SavePNG(outFile)
}
