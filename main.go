package main


// Import declaration declares library packages referenced in this file.
import (
	"encoding/json" // Package json implements encoding and decoding of JSON
	"html/template"
	"log"  // It defines a type
	"net/http" // a web server!
	"os" // OS functions like working with the file system

	"fmt" // A package in the Go standard library.
	"io/ioutil" // Implements some I/O utility functions.
)




//This function  register our handlers on server routes using the http.HandleFunc convenience function. It sets up the default router in the net/http package and takes a function as an argument.
func main() {  

	fmt.Println("Veuillez patientez encore quelque secondes svp")
	ApiArtist()
	//http.HandleFunc("/", PageEnter)
	//http.HandleFunc("/mainPage", mainPage)
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/ProfilArtist", ProfilArtist)   // HandleFunc registers the handler function for the given pattern in the DefaultServeMux
	http.HandleFunc("/searchApi", searchApi)
	

	fmt.Println()
	fmt.Println("Merci pour votre patience, allez sur ce lien = http://localhost:3000 svp")
	http.ListenAndServe(":3000", nil)  //First parameter of ListenAndServe is TCP address to listen to.
}										//// Second parameter is an interface, specifically http.Handler.


/*
func PageEnter(res http.ResponseWriter, req *http.Request) {
	enter, acc := template.ParseFiles("docs/htmlTemplates/index.html")
	if acc != nil {
		
		return
	}
	if req.URL.Path != "/" {
		
		return
	}
	
	result := []ApiData{}
	for _, v := range allData {
		
	result = append(result, v)
		
	}
	enter.Execute(res, result)


}
*/

// This ApiData struct type has nine fields
type ApiData struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relation     string              `json:"relations"`
	Concerts     map[string][]string `json:"datesLocations"`
	Locations    string   			 `json:"locations"`
}

//This relation struct type has ID and Concerts fields.
type relation struct {
	ID       int                 `json:"id"`
	Concerts map[string][]string `json:"datesLocations"`
}

var allData []ApiData  // Type switch allows switching on the type of something instead of value


// This is a function 
func ApiArtist() {

	style := http.FileServer(http.Dir("Css&Html template"))
	http.Handle("/Css&Html template/", http.StripPrefix("/Css&Html template/", style))  // StripPrefix returns a handler that serves HTTP requests by removing the given prefix from the request URL's Path 

	allData = AllDataUp("https://groupietrackers.herokuapp.com/api/artists")  // We’ll parse this example URL, which includes a scheme, authentication info, host, port, path, query params, and query fragment.
	if allData == nil {
		fmt.Println("Erreur")
		os.Exit(1)
	}


}

// Parsing Templates
func mainPage(res http.ResponseWriter, req *http.Request) {
	main, page := template.ParseFiles("Css&Html template/Html/PageGroupieTracker.html")
	if page  != nil {
		
		return
	}
	if req.URL.Path != "/" {
		
		return
	}
	
	result := []ApiData{}
	for _, v := range allData {
		
	result = append(result, v)
		
	}
	main.Execute(res, result)

}


func ProfilArtist(res http.ResponseWriter, req *http.Request) {  // A common way to write a handler is by using the http.HandlerFunc adapter on functions with the appropriate signature.
	profil, fileArtist := template.ParseFiles("Css&Html template/Html/ProfilArtist.html")
	if fileArtist != nil {
		fmt.Println(fileArtist)
		
		return
	}
	name := req.FormValue("name")
	for _, v := range allData {
		if v.Name == name {
			fmt.Println(v)
			profil.Execute(res, v)
			break
		}
	}
	return
}



// To have a search bar
func searchApi(res http.ResponseWriter, req *http.Request) {
	search, element := template.ParseFiles("Css&Html template/Html/searchapi.html")
	if element != nil {
		
		return
		
	}
	search.Execute(res, allData)
}

//To collect all the data

func AllDataUp(link string) []ApiData {
	dataOne := Data(link)
	Artists := []ApiData{}
	e := json.Unmarshal(dataOne, &Artists)
	if e != nil {
		log.Fatal(e)
		return nil
	}
	for i := 0; i < len(Artists); i++ {
		r := relation{}
		json.Unmarshal(Data(Artists[i].Relation), &r)
		Artists[i].Concerts = r.Concerts
	}
	return Artists
}


To get all the data
func Data(link string) []byte {
	dataOne, urlOne := http.Get(link)
	if urlOne != nil {
		log.Fatal( urlOne)
		return nil
	}
	dataTwo,  urlTwo := ioutil.ReadAll(dataOne.Body)
	if urlTwo != nil {
		log.Fatal(urlTwo)
		return nil
	}
	return dataTwo
}


// To get the character
func String(ar []string, a string) bool {
	for _, v := range ar {
		if v == a {
			return true
		}
	}
	return false
}