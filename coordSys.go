package gomath

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CreateGraph(f expression) {
	rl.InitWindow(600, 600, "Graph")
	rl.SetTargetFPS(60)
	points := DrawGraph(f)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for i := -600; i < 600; i++ {
			rl.DrawPixel(int32(i)+300, -int32(points[i+600])+300, rl.White)
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func DrawGraph(f expression) []float64 {
	var result []float64
	for i := -600; i < 600; i++ {
		result = append(result, f(float64(i)/120)*120)
	}
	return result
}
