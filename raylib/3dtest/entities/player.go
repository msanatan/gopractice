package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position rl.Vector3
	Size     float32
}

func (p *Player) Draw() {
	rl.DrawCube(p.Position, p.Size, p.Size, p.Size, rl.Red)
	rl.DrawCubeWires(p.Position, p.Size, p.Size, p.Size, rl.Maroon)
}

func (p *Player) Move(velocity rl.Vector3) {
	p.Position = rl.Vector3Add(p.Position, velocity)
}
