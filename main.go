package main

import(
  "fmt"
  "time"
)

type Ball struct{ hits int }

func main() {
  table := make(chan *Ball)
  go player("ping", table)
  go player("pong", table)


  fmt.Println("The game is afoot...")
  table <- new(Ball) // game on; toss the ball
  time.Sleep(1 * time.Second)
  <-table // game over; grab the ball
  fmt.Println("The game is over.")
}

func player(name string, table chan *Ball) {
  for {
    ball := <-table
    ball.hits++
    fmt.Println(name, ball.hits)
    time.Sleep(50 * time.Millisecond)
    table <- ball
  }
}