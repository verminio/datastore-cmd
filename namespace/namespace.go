package namespace

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

type Namespace struct {
	Name string `json:"namespace"`
}

func List(projectId string) (*[]Namespace, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, projectId)

	if err != nil {
		return nil, fmt.Errorf("namespace.List: %v", err)
	}
	query := datastore.NewQuery("__namespace__").KeysOnly()

	keys, err := client.GetAll(ctx, query, nil)

	if err != nil {
		return nil, fmt.Errorf("namespace.List: %v", err)
	}

	res := &[]Namespace{}
	for _, v := range keys {
		*res = append(*res, Namespace{Name: v.Name})
	}

	return res, nil
}
