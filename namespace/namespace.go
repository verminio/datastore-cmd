package namespace

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

type Namespace struct {
	Name string `json:"name"`
}

type Namespaces []Namespace

func List(projectId string) (Namespaces, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectId)

	if err != nil {
		return nil, fmt.Errorf("namespace.List: %v", err)
	}
	query := datastore.NewQuery("__namespace__").KeysOnly()

	res, err := client.GetAll(ctx, query, nil)

	if err != nil {
		return nil, fmt.Errorf("namespace.List: %v", err)
	}

	namespaces := Namespaces{}
	for _, v := range res {
		namespaces = append(namespaces, Namespace{Name: v.Name})
	}

	return namespaces, nil
}
