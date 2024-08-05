package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"userclientservice/handler"
	pb "userclientservice/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8084"
)

type Request struct {
	Id string `json:"id"`
}

func GrpcConnection() (pb.Client2RequestClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection failed %v", err)
	}

	return pb.NewClient2RequestClient(conn), conn
}

func main() {
	client2, conn := GrpcConnection()
	defer conn.Close()
	http.HandleFunc("/fetch", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var req Request
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userData := &pb.Id{
			Id: req.Id,
		}
		res, err := handler.FetchUser(client2, userData)
		if err != nil {
			log.Printf("cannot fetch the user %v", err)
		}

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(res)
	})

	fmt.Println("server running at 8031")

	if err := http.ListenAndServe(":8031", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Println("its working")
}

// func main() {

// 	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("Connection failed %v", err)
// 	}
// 	defer conn.Close()
// 	client1 := pb.NewClient2RequestClient(conn)
// 	user := &pb.Id{
// 		Id: "9",
// 	}
// 	res, err := handler.FetchUser(client1, user)
// 	if err != nil {
// 		log.Printf("cannot fetch the user %v", err)
// 	}
// 	fmt.Print(res)

// }
