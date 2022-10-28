package Game

import (
	"errors"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"GoPong/Entities/Ball"
	"GoPong/Entities/Racket"
)

type GameConfig struct {
	ptRacketSize  sdl.Point
	iAutoSpeed    int32
	colPlayer     sdl.Color
	colBot        sdl.Color
	colBall       sdl.Color
	iBallSize     int32
	ptBallSpeed   sdl.Point
	iWallDistance int32
}

type Game struct {
	gc GameConfig

	ptRes sdl.Point

	pPlayer      *Racket.Racket
	iPlayerScore int

	pBot      *Racket.Racket
	iBotScore int

	pBall      *Ball.Ball
	hasStarted bool

	pFont *ttf.Font
}

func New(windowRes sdl.Point) (*Game, error) {
	game := Game{}
	game.gc = GameConfig{
		ptRacketSize:  sdl.Point{X: 30, Y: 150},
		iAutoSpeed:    8,
		colPlayer:     sdl.Color{R: 41, G: 190, B: 176, A: 255},
		colBot:        sdl.Color{R: 255, G: 64, B: 0, A: 255},
		colBall:       sdl.Color{R: 255, G: 255, B: 255, A: 255},
		iBallSize:     30,
		ptBallSpeed:   sdl.Point{X: 10, Y: 10},
		iWallDistance: 50,
	}

	game.ptRes = windowRes

	game.pPlayer = Racket.New(
		game.gc.ptRacketSize.X,
		game.gc.ptRacketSize.Y,
		game.gc.colPlayer,
		game.ptRes,
		game.gc.iWallDistance,
		true,
	)

	game.iPlayerScore = 0

	game.pBot = Racket.New(
		game.gc.ptRacketSize.X,
		game.gc.ptRacketSize.Y,
		game.gc.colBot,
		game.ptRes,
		game.gc.iWallDistance,
		false,
	)

	game.iBotScore = 0

	game.pBall = Ball.New(game.gc.ptBallSpeed, game.ptRes, game.gc.iBallSize, game.gc.colBall, game.gc.iWallDistance)

	game.hasStarted = false

	ttf.Init()

	err := error(nil)
	game.pFont, err = ttf.OpenFont("./Resources/Roboto.ttf", 32)

	if err != nil {
		return &game, errors.New("there was an issue loading the font")
	}

	return &game, err
}

func (pGame *Game) InitGraph(pRenderer *sdl.Renderer) {
	pGame.pBall.InitGraph(pRenderer)
	pGame.pPlayer.InitGraph(pRenderer)
	pGame.pBot.InitGraph(pRenderer)
}

func (pGame *Game) Draw(pRenderer *sdl.Renderer) {
	if !pGame.hasStarted {
		pGame.renderTextOnScreen(pRenderer, "Press [SPACE] to start the game", pGame.gc.colBall, sdl.Point{X: pGame.ptRes.X / 2, Y: 50})
	}

	pGame.pBall.Draw(pRenderer)
	pGame.pPlayer.Draw(pRenderer)
	pGame.pBot.Draw(pRenderer)

	pGame.renderTextOnScreen(pRenderer, fmt.Sprint(pGame.iPlayerScore), pGame.gc.colPlayer, sdl.Point{X: pGame.ptRes.X / 4, Y: 50})
	pGame.renderTextOnScreen(pRenderer, fmt.Sprint(pGame.iBotScore), pGame.gc.colBot, sdl.Point{X: 3 * (pGame.ptRes.X / 4), Y: 50})
}

func (pGame *Game) MovePlayer(iY int32) {
	pGame.pPlayer.MoveManual(iY)
}

func (pGame *Game) Start() {
	if !pGame.hasStarted {
		pGame.hasStarted = true
		go pGame.pBall.Freeze()
	}
}

func (pGame *Game) Update() {
	if !pGame.hasStarted {
		return
	}

	pGame.pBall.Move([]sdl.Rect{pGame.pPlayer.GetFrame(), pGame.pBot.GetFrame()})
	pGame.pBot.MoveAuto(pGame.gc.iAutoSpeed, pGame.pBall.GetCoords())

	if pGame.pBall.GetCoords().X > pGame.ptRes.X {
		pGame.iPlayerScore++
		pGame.pBall.Reset()
		go pGame.pBall.Freeze()
	} else if pGame.pBall.GetCoords().X < 0 {
		pGame.iBotScore++
		pGame.pBall.Reset()
		go pGame.pBall.Freeze()
	}

}

func (pGame *Game) renderTextOnScreen(pRenderer *sdl.Renderer, text string, col sdl.Color, pos sdl.Point) error {
	surFont, err := pGame.pFont.RenderUTF8Blended(text, col)
	if err != nil {
		return err
	}

	texFont, err := pRenderer.CreateTextureFromSurface(surFont)
	if err != nil {
		return err
	}

	_, _, w, h, err := texFont.Query()

	texRect := sdl.Rect{X: pos.X - w/2, Y: pos.Y - h/2, W: w, H: h}
	pRenderer.Copy(texFont, nil, &texRect)

	return err
}
