package main

import "fmt"
import "math/rand"

func sumSlices(x[]int) int {
	totalx := 0
	for _, valuex := range x {
		totalx += valuex
	}
    return totalx
}

func shuffle(a []int) []int {
    b := make([]int, 0, 0)
    k := len(a)
    for i := 0; i < k; i++ {
    j := rand.Intn(len(a))
    b = append(b, a[j])
    a = append(a[:j], a[j+1:]...)
    }
    return b
}

//draw function
func draw(number int, hand []int, deck []int, discard []int) ([]int, []int, []int){
    for i := 0; i < number; i++ {
        if len(deck) > 0 {
            hand = append(hand, deck[:1]...)
            deck = deck[1:]
        } else {
            if len(discard) > 0 {
                deck = shuffle(discard)
                hand = append(hand, deck[:1]...)
                deck = deck [1:]
                discard = nil
            } else {
                return hand, deck, discard
            }
        }
    }
    return hand, deck, discard
}


func main() {

    totalturns := 0
    n := 10000
    for i := 0; i < n; i++ {
		turns := 0
		hand := make([]int, 0, 0)
		deck := make([]int, 0, 0)
		discard := []int{1,1,1,1,1,1,1,0,0,0}
		//fmt.Println(hand, deck, discard)
		for provinces := 0; provinces < 5; {
			turns ++
	        hand, deck, discard = draw(5, hand, deck, discard)
	        if sumSlices(hand) > 7 {
				provinces ++
				discard = append([]int{0}, discard...)
			} else if sumSlices(hand) > 5 {
				discard = append([]int{3}, discard...)
			} else if sumSlices(hand) > 2 {
			discard = append([]int{2}, discard...)
			}
	    	//fmt.Println(hand, deck, discard, provinces, turns)
			discard = append(hand, discard...)
			hand = nil
		}
		totalturns += turns
    }

	averageturns := float64(totalturns)/float64(n)
    fmt.Println("Total turns:", averageturns)
}
