package schema

import (
	"context"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
	"log"
)

func NewDgrapClient() *dgo.Dgraph {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	log.Println("NewDgrapClient success!")
	return client
}

func CreateSchema(client *dgo.Dgraph) error {
	schema := `name: string @index(term).
               age int.
               type Person {
					name
					age}
						`

	op := &api.Operation{Schema: schema}
	log.Println("CreateSchema success!")
	err := client.Alter(context.Background(), op)
	return err
}
