package main

import (
	"context"
	"fmt"
	"log"

	immuclient "github.com/codenotary/immuedb/pkg/client"
	"google.golang.org/grpc/metadata"
)

func Run() error {
	client, err := immuclient.NewImmuClient(immuclient.DefaultOptions())
	if err != nil {
		log.Fatal(err)
		return err
	}

	context := context.Background()
	lr, err := client.Login(context, []byte(`immudb`), []byte(`immudb`))
	if err != nil {
		log.Fatal(err)
		return err
	}

	//create new context
	md := metadata.Pairs("authorization", lr.Token)
	context = metadata.NewOutgoingContext(context.Background(), md)

	//key: []byte() val: []byte("another slice of byte)")
	vtx, err := client.VerifiedSet(context, []byte(`hi there`), []byte(`gopher`))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Successfully set key with transaction ID: %d\n", vtx.Id)

	ventry, err := client.VerifiedGet(context, []byte(`welcom`))
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Successfully retrived and verified entry %v\n", ventry)

	return nil
}

func main() {
	fmt.Println("running")

	if err := Run(); err != nil {
		fmt.Println("error ..!")
	}
}
