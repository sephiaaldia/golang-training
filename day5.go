package main

import (
	//"errors"
	"fmt"
	// "os"
	//"sort"
)

func day5()  {
	//Nomer 1
	var j = 2
	var counter1 = 0
	var inputan int
	var databaru int

	fmt.Println("masukkan index bilangan prima : ")
	fmt.Scanf("%v", &inputan)
	for {
		var counter = 0
		for i := 2; i <= j; i++ {
			if j%i == 0 {
				counter++
			}
		}
		if inputan == 1 {
			databaru = 2
			fmt.Printf("data prima ke %v adalah %v \n", inputan, databaru)
			// os.Exit(2)
		} else if counter == 1 {
			counter1++
			databaru = j
		} else if counter1 == inputan {
			fmt.Printf("data prima ke %v adalah %v \n", counter1, databaru)
			// os.Exit(2)
		}
		j++
	}

	//Nomer 2
}