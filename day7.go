package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

// type article struct {
// 	ID int
// 	Title string
// 	Content string
// }

// CONSUMING RESTFUL API
type starWarsPeople struct{
	Name string `json:"name"`
	Height string `json:"height"`
	Mass string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender string `json:"gender"`
	Films []string `json:"films"`
}

func day7()  {
	response, _ := http.Get("http://swapi.co/api/people/1")
	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var people starWarsPeople
	json.Unmarshal(responseData, &people)
	fmt.Println("---------Print Field-----------")
	fmt.Println("Name : ", people.Name)
	fmt.Println("Height : ", people.Height)
	fmt.Println("Mass : ",people.Mass)
	fmt.Println("Hair Color : ", people.HairColor)
	fmt.Println("film : ", people.Films[0])
}

// PACKAGE GO
// func articles(w http.ResponseWriter, r *http.Request){
// 	var data = []article{
// 		article{1, "Judul Pertama :", "Isi Pertama"},
// 		article{2, "Judul Kedua :", "Isi Kedua"},
// 	}

// 	w.Header().Set("Content-type", "application/json")

// 	if r.Method == "GET" {
// 		result, err := json.Marshal(data)
// 		if err != nil {
// 			http.Error(w,err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Write(result)
// 		return
// 	}
// 	http.Error(w, "", http.StatusBadRequest)
// 	return
// }

// func Hello(w http.ResponseWriter, r *http.Request)  {
// 	if r.Method==("GET"){
// 		http.Error(w, "Hello", http.StatusOK)
// 		return
// 	}
// 	http.Error(w, "", http.StatusBadRequest)
// 	return
// }

// func day7 (){
// 	http.HandleFunc("/articles", articles)
// 	http.HandleFunc("/Hello", Hello)
// 	fmt.Println("Starting web server at http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }