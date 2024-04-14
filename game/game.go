package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() *Game {
	player := NewPlayer()

	return &Game{
		player: player,
	}
}
