package main

import (
	"fmt"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	myFigure := figure.NewColorFigure("munchicken is Awesome!!!", "larry3d", "green", true)
	myFigure.Print()

	if secret := os.Getenv("SECRET"); secret != "" {
		mySecretFigure := figure.NewColorFigure(fmt.Sprintf("Secret value is: %s", secret), "larry3d", "red", true)
		mySecretFigure.Print()
	}
	
	time.Sleep(10 * time.Hour)
}
