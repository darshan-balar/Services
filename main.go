package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/darshan/services/controllers"
	"github.com/darshan/services/inits"
	"github.com/gorilla/mux"
)

var (
	
	ServiceController controllers.ServiceController
	VersionController controllers.VersionController

)

func init(){

	log.SetOutput(os.Stderr)

	// LoadConfig()
	inits.LoadConfig(os.Args[1])
	err := inits.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	ServiceController = controllers.NewServiceController(inits.DB) 
	VersionController = controllers.GetVersionController(inits.DB)


}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/services", ServiceController.Services).Methods("GET")
	router.HandleFunc("/versions", VersionController.GetVersions).Methods("GET")

	fmt.Println("Server is running on port 8081")
	http.ListenAndServe(":8081", router)
	

}