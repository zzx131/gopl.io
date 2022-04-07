package channeltest

import (
	"fmt"
	"testing"
	"time"
)

func TestForChannel(t *testing.T) {
	suck(pump())
	time.Sleep(1 * time.Second)

}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

func TestChannelTwo(t *testing.T) {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}
