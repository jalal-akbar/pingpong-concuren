package main

import (
	"log"
	"math/rand"
	"time"
)

type ball struct {
	hits       int
	lastPlayer string
}

func main() {
	table := make(chan *ball)
	done := make(chan *ball)

	go player("jalal", table, done)
	go player("akbar", table, done)

	referee(table, done)
}

func referee(table, done chan *ball) {
	ballToSend := new(ball)
	table <- ballToSend

	ballReceived := <-done
	log.Println("winner is", ballReceived.lastPlayer)

}

func player(name string, table, done chan *ball) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	for {
		select {
		case ball := <-table:
			v := r.Intn(1000)
			if v%11 == 0 {
				log.Println(name, "drop the ball")
				done <- ball
				return
			}
			ball.hits++
			ball.lastPlayer = name
			log.Println(name, "hits the ball", ball.hits)
			time.Sleep(50 * time.Millisecond)
			table <- ball
		case <-time.After(2 * time.Second):
			return
		}
	}
}
