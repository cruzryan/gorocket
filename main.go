package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = int32(800)
	screenHeight = int32(450)
	gridSize     = int32(15)

	thrust      = 7607 * (10 ^ 3) // N
	gravity     = -9.8            // m/s
	mass        = 549054          //kg
	initialFuel = 391640          //L
	flowRate    = 1451            // L / s

	rX = float32(0.0)
	rY = float32(10.0)
	rZ = float32(0.0)
)

var ()

func setupCamera() rl.Camera3D {
	c := rl.Camera3D{}
	c.Position = rl.NewVector3(15.0, 15.0, 15.0)
	c.Target = rl.NewVector3(0.0, 0.0, 0.0)
	c.Up = rl.NewVector3(0.0, 1.0, 0.0)
	c.Fovy = 45.0
	c.Type = rl.CameraPerspective

	return c
}

func main() {

	fmt.Println("Rocket started")

	//Window Set-up
	rl.InitWindow(screenWidth, screenHeight, "gorocket")
	rl.SetTargetFPS(60)

	//Camera Set-up
	camera := setupCamera()
	rl.SetCameraMode(camera, rl.CameraFree)

	rocket := Rocket{
		color:        rl.Blue,
		texture:      rl.LoadTexture("black.png"),
		model:        rl.LoadModel("tri_rocketshape.obj"),
		position:     rl.NewVector3(rX, rY, rZ),
		scale:        rl.NewVector3(1.0, 1.0, 1.0),
		rotationAxis: rl.NewVector3(1.0, 1.0, 1.0),
		fuel:         100.0,
	}

	rocket.setup()

	//Rotate angle
	//a := raymath.MatrixRotateX(rl.Deg2rad * 90)
	//rocket.model.Transform = raymath.MatrixMultiply(rocket.model.Transform, a)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera) // Update camera
		rocket.waitForKeys()
		rocket.update()

		rl.BeginDrawing()
		rl.BeginMode3D(camera)
		rl.ClearBackground(rl.RayWhite)

		rocket.draw()
		//Grid
		rl.DrawGrid(gridSize, 1.0)
		rl.EndMode3D()
		rocket.drawInfo()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
