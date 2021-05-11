package epik

import (
	"fmt"
)

func main() {
	Echo("hello")
	node, closer, err := NewFullNodeAPI("ws://116.63.167.70:1234/rpc/v0", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.Spdcp3ssmva1jFKxUQm2ORjG5oskg2OnHHqbqQu4E9w")
	defer closer()
	if err != nil {
		panic(err)
	}
	if node != nil {
		fmt.Println("ok")
	}

}

func Echo(msg string) {
	fmt.Println(msg)
}
