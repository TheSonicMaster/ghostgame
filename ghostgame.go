package main
import "fmt"
import "math/rand"
import "os"
import "time"

func main() {
	var version string = "1.0.0"
	var highscore int = 0
	fmt.Println("ghostgame version",version)
	fmt.Println("Copyright (c) 2021 The Sonic Master")
	for true {
		fmt.Println()
		rand.Seed(time.Now().UnixNano())
		var score int = 0
		fmt.Println("Three doors ahead, a ghost behind one...")
		for true {
			var ghostdoor int = rand.Intn(3 - 1 + 1) + 1
			fmt.Print("Which door do you open? 1, 2, or 3? ")
			var choice int
			fmt.Scanln(&choice)
			if (choice != 1) && (choice != 2) && (choice != 3) {
				continue
			} else if (choice == ghostdoor) {
				fmt.Println("Oh no! You chose the ghost door!")
				fmt.Println("Game over! You scored",score)
				break
			} else {
				fmt.Println("No ghost! You move to the next room!")
				score ++
			}
		}
		if (score > highscore) {
			fmt.Println("Congratulations! You got a new high score!")
			highscore = score
		} else {
			fmt.Println("Your highscore remains as",highscore)
		}
		for true {
			fmt.Print("Would you like to play again? [y/n] ")
			var retry string
			fmt.Scanln(&retry)
			if (retry == "y") || (retry == "Y") || (retry == "yes") || (retry == "Yes" ) || (retry == "YES") {
				break
			} else if (retry == "n") || (retry == "N") || (retry == "no") || (retry == "No" ) || (retry == "NO") {
				os.Exit(0)
			} else {
				continue
			}
		}
	}
}
