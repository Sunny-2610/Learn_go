package main

import "fmt"

func main() {

    for i := 1; i <= 5; i++ {

        if i == 3 {
			//it is going to skip the digit 3 . it is not going to include it in the counting.
            continue
        }

        fmt.Println(i)
    }
}