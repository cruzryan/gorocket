package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = int32(800)
	screenHeight = int32(450)
	gridSize     = int32(15)

	thrust      = int32(7607000)   // N
	gravity     = -9.8             // m/s
	mass        = int32(549054)    // kg
	initialFuel = int32(391640)    // L
	flowRate    = float32(34454.4) // L / s

	second = int32(6000) //To keep track of scale

	rX = float32(0.0)
	rY = float32(0.0)
	rZ = float32(0.0)
)

var (
	time = float32(0)
)

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

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera) // Update camera
		rl.BeginDrawing()
		rl.BeginMode3D(camera)
		rl.ClearBackground(rl.RayWhite)
		rocket.waitForKeys()
		rocket.update()

		rocket.draw()
		//Gridtttt
		rl.DrawGrid(gridSize, 1.0)
		rl.EndMode3D()
		rocket.drawInfo()
		//Fix this
		time += (1 / float32(second)) * 1000
		// fmt.Println(rocket.position.Y)
		rl.DrawText("T: "+fmt.Sprintf("%.2f", time)+" s", 10, 40, 20, rl.Gray)
		rl.DrawFPS(10, 10)
		rl.EndDrawing()

	}

	rl.CloseWindow()
}
