package Racket

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/veandco/go-sdl2/img"
)

type Racket struct {
	rFrame sdl.Rect
	col    sdl.Color

	ptRes     sdl.Point
	isPlayer1 bool

	tex *sdl.Texture
}

func New(iWidth int32, iHeight int32, color sdl.Color, windowRes sdl.Point, iWallDistance int32, isPlayer1 bool) *Racket {
	racket := Racket{}
	posX := windowRes.X - iWallDistance - iWidth/2
	if isPlayer1 {
		posX = iWallDistance - iWidth/2
	}

	racket.rFrame = sdl.Rect{
		X: posX,
		Y: windowRes.Y/2 - iHeight/2,
		W: iWidth,
		H: iHeight,
	}
	racket.ptRes = windowRes
	racket.col = color
	racket.isPlayer1 = isPlayer1
	racket.tex = nil
	return &racket
}

func (pRacket *Racket) InitGraph(pRenderer *sdl.Renderer) {
	pRacket.tex, _ = img.LoadTexture(pRenderer, "./Resources/racket.png")
	pRacket.tex.SetColorMod(pRacket.col.R, pRacket.col.G, pRacket.col.B)
}

func (pRacket *Racket) MoveManual(iY int32) {
	if iY-(pRacket.rFrame.H/2) > 0 && iY+(pRacket.rFrame.H/2) < pRacket.ptRes.Y {
		pRacket.rFrame.Y = iY - pRacket.rFrame.H/2
	}
}

func (pRacket *Racket) MoveAuto(iSpeed int32, ptBallCoords sdl.Point) {
	if (pRacket.isPlayer1 && (ptBallCoords.X < pRacket.ptRes.X/2)) || (!pRacket.isPlayer1 && (ptBallCoords.X > pRacket.ptRes.X/2)) {
		if ptBallCoords.Y < pRacket.rFrame.Y+pRacket.rFrame.H/2 && pRacket.rFrame.Y > 0 {
			pRacket.rFrame.Y -= iSpeed
		} else if ptBallCoords.Y > pRacket.rFrame.Y+pRacket.rFrame.H/2 && pRacket.rFrame.Y+pRacket.rFrame.H < pRacket.ptRes.Y {
			pRacket.rFrame.Y += iSpeed
		}
	}
}

func (pRacket *Racket) Draw(pRenderer *sdl.Renderer) {
	if pRacket.tex == nil {
		pRenderer.SetDrawColor(pRacket.col.R, pRacket.col.G, pRacket.col.B, pRacket.col.A)
		pRenderer.FillRect(&pRacket.rFrame)
	} else {
		pRenderer.Copy(pRacket.tex, nil, &pRacket.rFrame)
	}
}

func (pRacket *Racket) GetFrame() sdl.Rect {
	return pRacket.rFrame
}
