package namespace

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
)

type namespace struct {
}

func List(projectId string) error {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectId)

	if err != nil {
		return err
	}
	query := datastore.NewQuery("__namespace__")

	res, err := client.GetAll(ctx, query, nil)

	if err != nil {
		return fmt.Errorf("namespace.List: %v", err)
	}

	for _, v := range res {
		log.Printf("%v\n", v)
	}

	return nil
}
