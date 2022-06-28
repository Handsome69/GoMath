package gomath

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type graph func(float64) float64

func CreateGraph(f graph) {
	rl.InitWindow(450, 450, "Graph")
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		s := fmt.Sprintf("%f", f(5))
		rl.DrawText(s, 0, 0, 10, rl.White)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
