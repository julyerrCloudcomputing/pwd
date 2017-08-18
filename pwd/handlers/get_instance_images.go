package handlers

import (
	"encoding/json"
	"net/http"	
	"log"
	"github.com/play-with-docker/play-with-docker/pwd"
)

func GetInstanceImages(rw http.ResponseWriter, req *http.Request) {
	instanceImages := core.InstanceAllowedImages()
	json.NewEncoder(rw).Encode(instanceImages)
}

func SearchInstanceImages(rw http.ResponseWriter, req *http.Request) {
	body := pwd.ImageSearch{}
	// the image info has been stoted in the body
	json.NewDecoder(req.Body).Decode(&body)
	log.Println("***")
	log.Println("***")
	log.Println("***")
	log.Println(body.Term)
	log.Println(body.LimitNum)
	log.Println("***")
	log.Println("***")
	log.Println("***")
	instanceImages := core.InstanceImageSearch(body)
	if instanceImages == nil{
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(rw).Encode(instanceImages)
}