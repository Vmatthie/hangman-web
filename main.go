package main 

func main() {
	wins := 0
	loses := 0
	numletters := rand.Intn(11) + 4 //generates random number from 4 to 15
	again, has_won := play_hangman(numletters)

	for {
		if has_won == true {
			wins++
			numletters = rand.Intn(11) + 4 //generates random number from 4 to 15
		} else {
			loses++
			numletters = rand.Intn(11) + 4 //generates random number from 4 to 15
		}
		if again == "y" {
			clearscreen()
			fmt.Printf("------------------------\n")
			fmt.Printf("    Current Score\n")
			fmt.Printf("  %d: wins, %d: loses\n", wins, loses)
			fmt.Printf("------------------------\n")
			again, has_won = Play_hangman(numletters)
		} else if again == "n" {
			break
		}
	}
}
