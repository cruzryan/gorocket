package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

//Rocket ...
type Rocket struct {

	//How it looks
	color   rl.Color
	model   rl.Model
	texture rl.Texture2D

	//How it is
	position     rl.Vector3
	rotationAxis rl.Vector3
	scale        rl.Vector3
	angle        float32

	//forces
	vy float32
	vx float32
	vz float32

	yaw   float32
	pitch float32
	roll  float32

	//What it has
	fuel float32
}

func (r *Rocket) draw() {
	// TO-DO: Test rotation axis at y .2
	rl.DrawModelEx(r.model, r.position, r.rotationAxis, r.angle, r.scale, r.color) // Draw 3d model with texture
}

func (r *Rocket) drawInfo() {
	rl.DrawText("Yaw: "+fmt.Sprintf("%g", r.yaw), screenWidth-160, 20, 20, rl.DarkGray)
	rl.DrawText("Pitch: "+fmt.Sprintf("%g", r.pitch), screenWidth-160, 40, 20, rl.DarkGray)
	rl.DrawText("Roll: "+fmt.Sprintf("%g", r.roll), screenWidth-160, 60, 20, rl.DarkGray)
	rl.DrawText("Fuel: "+fmt.Sprintf("%g", r.fuel), screenWidth-160, 80, 20, rl.DarkGreen)
}

func (r *Rocket) setup() {
	r.model.Materials = make([]rl.Material, 1)
	r.model.Materials[0].Maps[rl.MapDiffuse].Texture = r.texture
	r.reset()
}

func (r *Rocket) reset() {

	r.vy = 0
	r.vx = 0
	r.vz = 0
	r.yaw = 0
	r.pitch = 0
	r.roll = 0

	r.fuel = initialFuel

}

func (r *Rocket) thrust(angle int) {

	r.fuel -= flowRate / second
}

func (r *Rocket) waitForKeys() {
	if rl.IsKeyPressed(rl.KeyR) {
		r.position = rl.NewVector3(rX, rY, rZ)
	}
}

func (r *Rocket) update() {
	r.vy += gravity

	r.position.Y += float32(r.vy / second)

	if r.position.Y <= 0 {
		r.position.Y = 0
		r.reset()
	}
}
