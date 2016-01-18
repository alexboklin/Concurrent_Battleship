/*
Here's how the enemy ships are located:

  1 2 3 4 5 6 7 8 9 10
A - - - - - - - - + +
B - + + + + + - - - -
C - - - - - - - - - -
D - - - + - - - + - -
E - - - + - - - + - -
F - - - + - - - + - -
G - - - - - - - + - -
H - - - - + + - - - -
I - - + - - - - - - -
J - - - - - - - - - +
*/

package main

import "fmt"

func main() {
	enemyShipUnitsHit := 0

	enemyShipMap := []string{
		"--------++",
		"-+++++----",
		"----------",
		"---+---+--",
		"---+---+--",
		"---+---+--",
		"-------+--",
		"----++----",
		"--+-------",
		"---------+",
	}

	done := make(chan bool)

	for _, row := range enemyShipMap {
		go func(rowToAnalyze string) {
			fmt.Println(rowToAnalyze)
			for _, symbol := range rowToAnalyze {
				if string(symbol) == "+" {
					enemyShipUnitsHit += 1
					if enemyShipUnitsHit == 18 {
						fmt.Println("Found all the ships!")
						done <- true
					}
				}
			}
		}(row)
	}
	<-done
}
