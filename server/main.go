package main

import (
	"log"
	"map-reduce-server/common"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type API int

func (a *API) DoMapStep(item common.MapStep) {
	log.Printf("DoMapStep: %d", item.MapStepNumber)

}

func DoReduceStep(item common.ReduceStep) {
	log.Println("DoReduceStep: ", item.ReduceStepNumber)

}
func HealthCheck(url string, reply *common.Response) {
	log.Println("HealthCheck: ", url)
	reply.Message = "OK"
}

func main() {
	port := ":" + os.Args[1]

	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("listener error", err)
	}

	log.Printf("Server: serving rpc on port %s", port)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving: ", err)
	}
}
