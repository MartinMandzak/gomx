/* i was making fun of his problems
   but i can't make a better one for the life of me */

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	var waiting_room = []string{}

	reader := bufio.NewReader(os.Stdin)
	var input string
	
	for input != "done" {
		fmt.Println("Please enter a command:")
		input,_ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input{
		case "done":
			break
		case "next":
			if len(waiting_room) != 0 {
				waiting_room=append(waiting_room[:0],waiting_room[1:]...)
			} else {
				fmt.Println("The waiting room is empty!")
			}
		default:
			waiting_room=append(waiting_room, input)
		}

		for _,el := range(waiting_room){
			fmt.Print(el+" ")
		}
		fmt.Println()
	}
}

