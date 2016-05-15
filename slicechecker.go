/*

This program consists of one function, which takes a []int type slice and checks its every element for following criteria:
    "If the sum of the elements on the left side of nth element of the slice is equal to the sum of the elements
    on the right side of nth element of the slice, return true or some handy information."
    and an example of that function's usage.
    It was a cool little excercise while learning Go :).

*/

package main

import "fmt"
import "strconv"
import "strings"

func SliceChecker(slice []int) string {

	//Here we check the length of the slice in order to iterate over its elements.
	var length = len(slice)

	//Now we use a for loop to check every element of the slice. Notice the incremented value is checked against
	//the length of the slice minus one. It's done like this, because for example: the length of the three element
	//slice {1, 2, 3} is indeed equal 3, like it's supposed to be, but we have to remember that indexes of slice
	//elements begin with 0, so in the slice {1, 2, 3} the first element has the zeroeth index number, the second
	//element has the first index number and so on. The index number of the last element of a slice is the length
	//of the slice decremented by 1.

	for i := 0; i <= length-1; i++ {

		//We declare a default value for the elements we want to compare - the sums - each time we iterate
		//over the next element, because every element will (mostly) have different sums than any other element.

		sumleft := 0
		sumright := 0

		//Now we iterate with a for loop over the range of our slice, sliced in two parts using "[:]" syntax.

		for _, left_side_element := range slice[:i] {

			//With each iteration (for every element) the sumleft variable, wich we declared with a default
			//value of 0 at the beginning of the main for loop is now incremented by the value of every
			//element on the left side of the current i-th element of the slice.

			sumleft += left_side_element
		}
		for _, right_side_element := range slice[i+1:] {

			//The same goes for the sumright variable, now being incremented by the value of every element
			//on the right side of the current i-th element of the slice.

			sumright += right_side_element
		}
		if sumleft == sumright {

			//This conditional statement checks if the partial sums are equal to each other.
			//If they are, program prints a sentence with some pretty useful information, being:
			//  - the list elements
			//  - the number which met the criteria
			//  - index of the number which met the criteria.
			//If they are not, it just iterates over the next element of the slice.

			fmt.Println("\nGiven the list: " + strings.Trim(fmt.Sprint(slice), "[]") +
				", the number which meets the criteria is: " +
				strconv.Itoa(slice[i]) + " and its index is: " +
				strconv.Itoa(i) + ".")
			fmt.Println("\nThe sum of the elements on both sides of the number is equal to: " +
				strconv.Itoa(sumleft) + ".")
		}
	}
	//We end the program with a polite message.
	return "The program has ended.\n"
}

func main() {

	//We declare an example slice...

	example_slice := []int{1, 3, 5, 2, 9}

	//...and test our function with it

	fmt.Println(SliceChecker(example_slice))
}
