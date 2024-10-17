package main

// go run main.go

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const clockWidth = 40
const clockHeight = 20
const clockRadius = 9

const xScale = 2.0

var clockFace [clockHeight][clockWidth]string

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func drawClockFace() {
	for i := range clockFace {
		for j := range clockFace[i] {
			clockFace[i][j] = " " // Clear clock face
		}
	}
	centerX, centerY := clockWidth/2, clockHeight/2

	for hour := 12; hour > 0; hour-- {
		angle := float64(hour) * (math.Pi / 6) // 360° / 12 = 30° = π / 6 radians
		x := int(math.Round(float64(centerX) + xScale*clockRadius*math.Sin(angle)))
		y := int(math.Round(float64(centerY) - clockRadius*math.Cos(angle)))
		clockFace[y][x] = strconv.FormatInt(int64(hour), 10)
	}
}

func plotHand(length int, angle float64, symbol string) {
	centerX, centerY := clockWidth/2, clockHeight/2
	for i := 1; i <= length; i++ {
		x := int(math.Round(float64(centerX) + xScale*float64(i)*math.Sin(angle)))
		y := int(math.Round(float64(centerY) - float64(i)*math.Cos(angle)))
		clockFace[y][x] = symbol
	}
}

func drawClock() {
	clearScreen()
	now := time.Now()

	hourAngle := float64(now.Hour()%12) * (math.Pi / 6)
	minuteAngle := float64(now.Minute()) * (math.Pi / 30)
	secondAngle := float64(now.Second()) * (math.Pi / 30)

	drawClockFace()

	plotHand(clockRadius-4, hourAngle, "#")
	plotHand(clockRadius-2, minuteAngle, "=")
	plotHand(clockRadius-1, secondAngle, "-")

	// Display the clock face
	for i := 0; i < clockHeight; i++ {
		for j := 0; j < clockWidth; j++ {
			fmt.Print(string(clockFace[i][j]))
		}
		fmt.Println()
	}
}

func main() {
	// Set up channel to listen for interrupt signal (Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Run the clock until interrupted
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-signalChan:
			fmt.Println("Clock interrupted. Exiting.")
			return
		case <-ticker.C:
			drawClock()
		}
	}
}
