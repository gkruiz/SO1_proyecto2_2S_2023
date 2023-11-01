package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	pb "servergrpc/proto-grpc"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

type Info struct {
	Carne    string `json:"carne"`
	Nombre   string `json:"nombre"`
	Curso    string `json:"curso"`
	Nota     int    `json:"nota"`
	Semestre string `json:"semestre"`
	Year     int    `json:"year"`
}

func almacenar_comentario(comentario string) {

	fmt.Printf("INGRESO SQL INSERT: %s", comentario)

	db, err := sql.Open("mysql", "root:asdf1234.,.,@tcp(34.121.157.97:3306)/tarea7")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var data Info

	if err := json.Unmarshal([]byte(comentario), &data); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(data)

	result, err2 := db.Query("Insert into datos(carnet,nombre,curso,nota,semestre,year)values(?,?,?,?,?,?)", data.Carne, data.Nombre, data.Curso, data.Nota, data.Semestre, data.Year)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(result)

}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	almacenar_comentario(in.GetId())
	fmt.Printf(">> Hemos recibido la data del cliente: %v\n", in.GetId())
	return &pb.ReplyInfo{Info: ">> Hola Cliente, he recibido el comentario: " + in.GetId()}, nil
}

func main() {

	escucha, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Fallo al levantar el servidor: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})
	if err := s.Serve(escucha); err != nil {
		log.Fatalf("Fallo al levantar el servidor: %v", err)
	}
}
