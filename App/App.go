package App

import (
	"errors"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"

	"GoPong/Game"
)

type App struct {
	pWindow   *sdl.Window
	pRenderer *sdl.Renderer
	strTitle  string

	pGame *Game.Game
}

func New(iWidth int32, iHeight int32, strTitle string) (*App, error) {
	app := App{}
	err := error(nil)

	sdl.Init(sdl.INIT_EVERYTHING)

	app.pWindow, err = sdl.CreateWindow(
		strTitle,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		iWidth, iHeight,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		return &app, errors.New("couldn't create SDL Window, aborting")
	}

	app.pRenderer, err = sdl.CreateRenderer(app.pWindow, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Couldn't create hardware accelerated renderer, reverting to software mode.")
		app.pRenderer, err = sdl.CreateRenderer(app.pWindow, -1, sdl.RENDERER_SOFTWARE)
		if err != nil {
			return &app, errors.New("couldn't create renderer, aborting")
		}
	}
	app.strTitle = strTitle

	app.pGame, err = Game.New(sdl.Point{X: iWidth, Y: iHeight})

	app.pGame.InitGraph(app.pRenderer)
	return &app, err
}

func (pApp *App) Run() {
	quit := false
	ticker := time.NewTicker((1000 / 60) * time.Millisecond)

	for !quit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.KeyboardEvent:
				keycode := t.Keysym.Sym
				switch keycode {
				case sdl.K_ESCAPE:
					quit = true
				case sdl.K_SPACE:
					pApp.pGame.Start()
				}
			case *sdl.MouseMotionEvent:
				pApp.pGame.MovePlayer(t.Y)
			}
		}

		_, update := <-ticker.C

		if update {
			pApp.pGame.Update()
			pApp.pRenderer.SetDrawColor(0, 0, 0, 255)
			pApp.pRenderer.Clear()

			pApp.pGame.Draw(pApp.pRenderer)
			pApp.pRenderer.Present()
		}
	}
}
