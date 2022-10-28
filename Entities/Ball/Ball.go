package Ball

import (
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Ball struct {
	ptSpeed sdl.Point
	rFrame  sdl.Rect
	col     sdl.Color

	ptRes         sdl.Point
	iWallDistance int32
	isFrozen      bool

	tex *sdl.Texture
}

func New(speed sdl.Point, windowRes sdl.Point, iSize int32, col sdl.Color, iWallDistance int32) *Ball {
	ball := Ball{}

	ball.ptSpeed = speed
	ball.rFrame = sdl.Rect{
		X: windowRes.X/2 - iSize/2,
		Y: windowRes.Y/2 - iSize/2,
		W: iSize,
		H: iSize,
	}
	ball.ptRes = windowRes
	ball.col = col
	ball.iWallDistance = iWallDistance
	ball.isFrozen = false
	ball.tex = nil

	return &ball
}

func (pBall *Ball) InitGraph(pRenderer *sdl.Renderer) {
	pBall.tex, _ = img.LoadTexture(pRenderer, "./Resources/ball.png")
	pBall.tex.SetColorMod(pBall.col.R, pBall.col.G, pBall.col.B)
}

func (pBall *Ball) Move(rects []sdl.Rect) {

	mapIsColliding := map[string]bool{
		"up":   pBall.rFrame.Y < 0,
		"down": pBall.rFrame.Y+pBall.rFrame.H > pBall.ptRes.Y,
	}

	if mapIsColliding["up"] || mapIsColliding["down"] {
		pBall.ptSpeed.Y *= -1
	}

	for i := 0; i < len(rects); i++ {
		if pBall.rFrame.HasIntersection(&rects[i]) && pBall.rFrame.X > pBall.iWallDistance && pBall.rFrame.X < pBall.ptRes.X+pBall.iWallDistance {
			pBall.ptSpeed.X *= -1
		}
	}

	if !pBall.isFrozen {
		pBall.rFrame.X += pBall.ptSpeed.X
		pBall.rFrame.Y += pBall.ptSpeed.Y
	}
}

func (pBall *Ball) Draw(pRenderer *sdl.Renderer) {
	if pBall.tex == nil {
		pRenderer.SetDrawColor(pBall.col.R, pBall.col.G, pBall.col.B, pBall.col.A)
		pRenderer.FillRect(&pBall.rFrame)
	} else {
		pRenderer.Copy(pBall.tex, nil, &pBall.rFrame)
	}
}

func (pBall *Ball) GetCoords() sdl.Point {
	return sdl.Point{X: pBall.rFrame.X + (pBall.rFrame.W / 2), Y: pBall.rFrame.Y + (pBall.rFrame.H / 2)}
}

func (pBall *Ball) Reset() {
	pBall.rFrame.X, pBall.rFrame.Y = pBall.ptRes.X/2-pBall.rFrame.W/2, pBall.ptRes.Y/2-pBall.rFrame.H/2
}

func (pBall *Ball) Freeze() {
	pBall.isFrozen = true
	timer := time.NewTimer(1500 * time.Millisecond)
	<-timer.C
	pBall.isFrozen = false
}
