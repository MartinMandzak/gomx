/*
Recreating the almigthy bus exericse from my high school...
*/

package main

import (
	"fmt"
	"math/rand"
)

type Bus struct{
	ID string
	capacity int
	passengerCount int
	stopCount int
}

func MinInteger(a int, b int) int{
	if a>b{
		return b;
	}
	return a;
}

func main(){
	var maxChange int = 16;

	bus1 := Bus{"A-001", 50,0,14}
	bus2 := Bus{"A-002", 35,6,10}
	bus3 := Bus{"A-003", 20,0,7}
	bus4 := Bus{"A-004", 50,12,12}

	garage := [4]Bus{bus1,bus2,bus3,bus4};

	for _, bus := range garage{
		fmt.Printf("Bus %s is out!\n",bus.ID);
		for i:=0; i<bus.stopCount; i++ {
			var passengersOut = rand.Intn(maxChange);
			var passengersIn = rand.Intn(maxChange);

			bus.passengerCount -= MinInteger(bus.passengerCount, passengersOut);
			if bus.passengerCount + passengersIn <= bus.capacity{
				bus.passengerCount += passengersIn;
			}

			fmt.Printf(
				"Stop no.%d! Current passenger count: %d/%d ...\n",(i+1), bus.passengerCount,bus.capacity,
			);
		}
		fmt.Printf("End of the ride! %d people are walking home...", bus.passengerCount);
		fmt.Println("\n#############\n");
	}
}


