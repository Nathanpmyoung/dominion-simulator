package main

import "fmt"
import "math/rand"
//import "math"
import "time"

type Game struct{
	Hand []string `json:"hand"`
	InPlay []string `json:"inplay"`
	Deck []string `json:"deck"`
	Discard []string `json:"discard"`
}

func sumCoins(cards[]string) int {
	totalx := 0
	for _, value := range cards {
		switch value {
		case "Copper":
			totalx += 1
		case "Silver":
			totalx += 2
		case "Gold":
			totalx += 3
		}
	}
    return totalx
}

func shuffle(a []string) []string {
    b := []string{}
    k := len(a)
    for i := 0; i < k; i++ {
    j := rand.Intn(len(a))
    b = append(b, a[j])
    a = append(a[:j], a[j+1:]...)
    }
    return b
}

func lookThrough(name string,cards []string) int {
	howMany := 0
	for _, card := range cards {
		if card == name {
			howMany += 1
		}
	}
	return howMany
}

func smithyCondition(ratio float64,game Game) bool {

	decks := [][]string{
		game.Hand,
		game.InPlay,
		game.Deck,
		game.Discard,
	}
	totalCards := make([]string,0,0)
	for _, c := range decks {
    totalCards = append(totalCards, c...)
	}
	//fmt.Println("Total Cards:", len(totalCards), totalCards)
	if lookThrough("Smithy", totalCards) < 10 && float64(lookThrough("Smithy", totalCards))/ float64(len(totalCards)) < ratio {
		return true
	}

	return false
}

//draw function
func draw(number int,game Game) (Game){

    for i := 0; i < number; i++ {
        if len(game.Deck) > 0 {
            game.Hand = append(game.Hand, game.Deck[:1]...)
            game.Deck = game.Deck[1:]
        } else {
            if len(game.Discard) > 0 {
                game.Deck = shuffle(game.Discard)
                game.Hand = append(game.Hand, game.Deck[:1]...)
                game.Deck = game.Deck [1:]
                game.Discard = nil
            } else {
                return game
            }
        }
    }
    return game
}

/*
func ratioCondition(minRatio [2]float64, midRatio [2]float64, maxRatio [2]float64, close float64) bool {
	if (midRatio[1] - minRatio[1] < close) || (maxRatio[1] - midRatio[1] < close) {
		return false
	}
	return true
}
*/

func fullGame(ratio float64) float64{
	totalturns := 0
	//totalstdev := 0.0
	n := 100000
	//strat := {"","","","Silver","Smithy","","Gold","","Province"}

	for i := 0; i < n; i++ {
		turns := 0
		game := Game{Hand:[]string{}, InPlay:[]string{}, Deck:[]string{}, Discard:[]string{"Copper","Copper","Copper","Copper","Copper","Copper","Copper","Estate","Estate","Estate"}}
		//fmt.Println(game.Hand)
		for provinces := 0; provinces < 5; {
			turns ++
			game = draw(5, game)
			if lookThrough("Smithy", game.Hand) > 0 {
				//fmt.Println("Play Smithy:", game)
				// put smithy in play game.Discard = append([]string{"Smithy"}, game.Discard...)
				// remove Smithy from game.Hand
				game = draw(3, game)
			}
			if sumCoins(game.Hand) > 7 {
				provinces ++
				game.Discard = append([]string{"Province"}, game.Discard...)
			} else if sumCoins(game.Hand) > 5 {
				game.Discard = append([]string{"Gold"}, game.Discard...)
			} else if sumCoins(game.Hand) > 3 && smithyCondition(ratio, game) {
				game.Discard = append([]string{"Smithy"}, game.Discard...)
				//If the card density is higher
				//card totals?
			} else if sumCoins(game.Hand) > 2 {
				game.Discard = append([]string{"Silver"}, game.Discard...)
			}
			//fmt.Println(game)
			game.Discard = append(game.Hand, game.Discard...)
			game.Hand = nil
		}
		totalturns += turns
		//totalstdev += math.Pow((float64(turns)-18.15),2)
	}

	averageTurns := float64(totalturns)/float64(n)
	//fullstdev := math.Pow((totalstdev/(float64(n)-1)),.5)
	fmt.Println("Total turns:", totalturns, "n:", n, "Average turns:", averageTurns, "Ratio:", ratio)
	return averageTurns
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	/*
	minRatio := [2]float64{0.0, 0}
	midRatio := [2]float64{0.25, 0}
	maxRatio := [2]float64{0.5, 0}
	//close := 0.1

	minRatio[1] = fullGame(minRatio[0])
	fmt.Println("minRatio:", minRatio)
	midRatio[1] = fullGame(midRatio[0])
	fmt.Println("midRatio:", midRatio)
	maxRatio[1] = fullGame(maxRatio[0])
	fmt.Println("maxRatio:", maxRatio)

	for i := 0; i < 50 ; i++ {

		currentRatio := midRatio
		if midRatio[0] - minRatio[0] == maxRatio[0] - midRatio[0] {
			if minRatio[1] > maxRatio[1]{
				currentRatio[0] = minRatio[0] + (midRatio[0] - minRatio[0]) * 0.5
			}else{
				currentRatio[0] = midRatio[0] + (maxRatio[0] - midRatio[0]) * 0.5
			}
		} else{
			if midRatio[0] - minRatio[0] > maxRatio[0] - midRatio[0] {
				currentRatio[0] = minRatio[0] + (midRatio[0] - minRatio[0]) * 0.5
			}else{
				currentRatio[0] = midRatio[0] + (maxRatio[0] - midRatio[0]) * 0.5
			}
		}

		currentRatio[1] = fullGame(currentRatio[0])
		fmt.Println("minRatio:", minRatio, "midRatio:", midRatio, "maxRatio:", maxRatio, "currentRatio:", currentRatio)

		if currentRatio[0] < midRatio[0]{
			if currentRatio[1] > minRatio[1]{
				//fmt.Println("Error: minRatio:", minRatio, "midRatio:", midRatio, "maxRatio:", maxRatio, "currentRatio:", currentRatio)
				//break
			}else if currentRatio[1] < midRatio[1] {
				maxRatio = midRatio
				midRatio = [2]float64{currentRatio[0],currentRatio[1]}

			}else{
				minRatio = [2]float64{currentRatio[0],currentRatio[1]}
			}
		}else{
			if currentRatio[1] > maxRatio[1]{

				//fmt.Println("Error: minRatio:", minRatio, "midRatio:", midRatio, "maxRatio:", maxRatio, "currentRatio:", currentRatio)
				//break
			}else if currentRatio[1] < midRatio[1] {
				minRatio = midRatio
				midRatio = [2]float64{currentRatio[0],currentRatio[1]}
			}else{
				maxRatio = [2]float64{currentRatio[0],currentRatio[1]}
			}
		}
	}
	*/

	for i := 0.0; i<100; i++{
		fullGame(i/100)
	}
}
