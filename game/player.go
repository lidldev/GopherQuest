package game

import (
	"fmt"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Char struct {
	x  int
	y  int
	vx int
	vy int

	s *ebiten.Image
}

const (
	unit    = 16
	groundY = 510
)

func (c *Char) tryJump() {
	if c.y == groundY*unit {
		c.vy = -10 * unit
	}
}

func (c *Char) Update() {
	c.x += c.vx
	c.y += c.vy

	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.vx > 0 {
		c.vx -= 2
	} else if c.vx < 0 {
		c.vx += 2
	}
	if c.vy < 20*unit {
		c.vy += 8
	}
}

func (c *Char) Draw(screen *ebiten.Image) {
	c.s = assets.IdleSprite
	if c.vx > 0 {
		c.s = assets.RightSprite
	} else if c.vx < 0 {
		c.s = assets.LeftSprite
	}

	msg := fmt.Sprintf("Gopher X: %d Gopher Y: %d", c.x, c.y)
	ebitenutil.DebugPrint(screen, msg)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(c.s, op)
}

type Player struct {
	player *Char
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.player.Draw(screen)
}

func (p *Player) Update() {
	if p.player == nil {
		p.player = &Char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.player.vx = -2 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.player.vx = 2 * unit
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.player.tryJump()
	}
	p.player.Update()
}
