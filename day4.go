package main

import (
	// "fmt"
	// "bufio"
	// "os"
	// "strings"
	// "sort"
	// "strconv"
)

func day4()  {
	// // Nomer 1
	// var array1, array2 int

	// fmt.Printf("Jumlah array 1 : ")
	// fmt.Scanf("%v", &array1)

	// fmt.Printf("Jumlah array 2 : ")
	// fmt.Scanf("%v", &array2)

	// arraySlice1 := make([]string, array1)
	// arraySlice2 := make([]string, array2)
	// for a := 0; a < array1; a++ {
	// 	fmt.Printf("Masukkan data array 1 : ")
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	scanner.Scan()
	// 	arraySlice1[a] = scanner.Text()
	// }
	// for b := 0; b < array2; b++ {
	// 	fmt.Print("Masukkan data array 2 : ")
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	scanner.Scan()
	// 	arraySlice2[b] = scanner.Text()
	// }

	// fmt.Println(arraySlice1)
	// fmt.Println(arraySlice2)
	// combinedArray := append(arraySlice1, arraySlice2...)
	// fmt.Println(combinedArray)
	// for c := 0; c < len(combinedArray); c++ {
	// 	for d := c + 1; d < len(combinedArray); d++ {
	// 		if combinedArray[c] == combinedArray[d] {
	// 			combinedArray[d] = combinedArray[len(combinedArray)-1]
	// 			combinedArray[len(combinedArray)-1] = ""
	// 			combinedArray = combinedArray[:len(combinedArray)-1]
	// 		}
	// 	}
	// }
	// fmt.Println("Hasil : ", combinedArray)


	// //Nomer 2
	// var data[100] int
	// var banyakData,sum int
	// var mean float64
	
    // fmt.Print("Masukkan Banyak Data: ")
	// fmt.Scanln(&banyakData)
	
    // for a := 0; a < banyakData; a++ {
    //     fmt.Print("Masukkan data : ")
    //     fmt.Scanln(&data[a])     
    //     sum += data[a]
    // }
	// mean = float64(sum)/float64(banyakData)
	// fmt.Printf("Mean : %v \n",mean)

	// if banyakData % 2 == 0{
	// 	median1 := (banyakData/2) 
	// 	median2 := (banyakData/2) - 1
	// 	median := float64(data[median1] + data[median2])/2
	// 	fmt.Printf("Median : %v \n", median )
	// }else{
	// 	median := (banyakData - 1)/2 
	// 	fmt.Printf("Median : %v \n", data[median] )
	// }

	// //Nomer 3
	// var keys, value string
	// var inputkeys,inputvalue []string
	// var combinedMap = map[string]string{}

	// fmt.Printf("masukkan keys : ")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// keys = scanner.Text()
	// fmt.Printf("masukkan value : ")
	// scanner1 := bufio.NewScanner(os.Stdin)
	// scanner1.Scan()
	// value = scanner1.Text()
	// inputkeys = strings.Split(keys, ",")
	// inputvalue = strings.Split(value, ",")
	// for i := 0; i < len(inputkeys); i++ {
	// 	combinedMap[inputkeys[i]] = inputvalue[i]
	// }
	// fmt.Println(combinedMap)

	// // nomer 4
	// var inputanCommand string
	// var dataArray1, dataArray2 []string
	// dataArray3 := make([]int, 3, 3)

	// fmt.Println("Command string : ")
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// inputanCommand = scanner.Text()
	// dataArray1 = strings.Split(inputanCommand, ",")
	// fmt.Println(len(dataArray1))
	// for i := 0; i < len(dataArray1); i++ {
	// 	dataArray2 = strings.Split(strings.Trim(dataArray1[i], " "), " ")
	// 	switch dataArray2[0] {
	// 	case "insert":
	// 		posisi, _ := strconv.Atoi(dataArray2[1])
	// 		value, _ := strconv.Atoi(dataArray2[2])
	// 		dataArray3[posisi-1] = value
	// 		fmt.Println(dataArray3)
	// 	case "remove":
	// 		posisi, _ := strconv.Atoi(dataArray2[1])
	// 		dataArray3[posisi-1] = dataArray3[len(dataArray3)-1]
	// 		dataArray3[len(dataArray3)-1] = 0
	// 		dataArray3 = dataArray3[:len(dataArray3)-1]
	// 		fmt.Println(dataArray3)
	// 	case "append":
	// 		value, _ := strconv.Atoi(dataArray2[1])
	// 		fmt.Println(value)
	// 		dataArray3 = append(dataArray3, value)
	// 		fmt.Println(dataArray3)
	// 	case "sort":
	// 		sort.Ints(dataArray3)
	// 		fmt.Println(dataArray3)
	// 	case "pop":
	// 		dataArray3[len(dataArray3)-1] = 0
	// 		dataArray3 = dataArray3[:len(dataArray3)-1]
	// 		fmt.Println(dataArray3)
	// 	case "reverse":
	// 		sort.Ints(dataArray3)
	// 		sort.Slice(dataArray3, func(i, j int) bool {
	// 			return dataArray3[i] > dataArray3[j]
	// 		})
	// 		fmt.Println(dataArray3)
	// }
	// }
}
