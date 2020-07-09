package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gen2brain/raylib-go/raymath"
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

	defaultTransform rl.Matrix
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
	// TO-DO: Test rotation axis at y .2 or .8
	rl.DrawModelEx(r.model, r.position, r.rotationAxis, r.angle, r.scale, r.color) // Draw 3d model with texture
}

func (r *Rocket) drawInfo() {
	rl.DrawText("Yaw: "+fmt.Sprintf("%g", r.yaw)+"°", screenWidth-160, 20, 20, rl.DarkGray)
	rl.DrawText("Pitch: "+fmt.Sprintf("%g", r.pitch)+"°", screenWidth-160, 40, 20, rl.DarkGray)
	rl.DrawText("Roll: "+fmt.Sprintf("%g", r.roll)+"°", screenWidth-160, 60, 20, rl.DarkGray)
	rl.DrawText("Fuel: "+fmt.Sprintf("%g", r.fuel)+" L", screenWidth-160, 80, 20, rl.DarkGreen)
	rl.DrawText("[Hit Speed: "+fmt.Sprintf("%.1f", r.vy)+"m/s]", screenWidth-225, screenHeight-20, 20, rl.DarkGreen)
}

func (r *Rocket) setup() {
	r.model.Materials = make([]rl.Material, 1)
	r.model.Materials[0].Maps[rl.MapDiffuse].Texture = r.texture
	r.defaultTransform = r.model.Transform
	r.reset()
}

func (r *Rocket) resetAxis() {
	r.yaw = 0.0
	r.pitch = 0.0
	r.roll = 0.0

	//Resets all to 90 degree default
	r.model.Transform = r.defaultTransform

}

func (r *Rocket) reset() {

	r.vy = 0
	r.vx = 0
	r.vz = 0
	r.fuel = float32(initialFuel)

	time = 0
}

func (r *Rocket) thrust() {

	// fmt.Println(int32(thrust) / mass)
	var v float32 = float32(thrust / mass)

	//this only woks if thrust & rocket are always 90 degrees
	r.vy += float32(v)

	var fr float32 = float32(flowRate) / float32(second)

	r.fuel -= float32(fr)
}

func (r *Rocket) waitForKeys() {

	if rl.IsKeyPressed(rl.KeyR) {
		r.position = rl.NewVector3(rX, rY, rZ)
		r.reset()
	}

	if rl.IsKeyDown(rl.KeyT) {
		r.thrust()
	}

	if rl.IsKeyDown(rl.KeyF) {
		r.resetAxis()
	}

	if rl.IsKeyDown(rl.KeyW) {
		r.yaw++
	} else if rl.IsKeyDown(rl.KeyQ) {
		r.yaw--
	}

}

func (r *Rocket) isVertical() bool {
	return r.pitch == 0 && r.yaw == 0 && r.roll == 0
}

func (r *Rocket) update() {

	//Yaw control
	yawUpdate := raymath.MatrixRotateX(rl.Deg2rad * r.yaw)
	r.model.Transform = raymath.MatrixMultiply(r.model.Transform, yawUpdate)
	//Pitch Control
	pitchUpdate := raymath.MatrixRotateY(rl.Deg2rad * r.pitch)
	r.model.Transform = raymath.MatrixMultiply(r.model.Transform, pitchUpdate)
	//Roll control
	rollUpdate := raymath.MatrixRotateZ(rl.Deg2rad * r.roll)
	r.model.Transform = raymath.MatrixMultiply(r.model.Transform, rollUpdate)

	r.vy += gravity

	r.position.Y += r.vy / float32(second)

	if r.position.Y <= 0 {
		r.position.Y = 0
		r.reset()
	}
}
