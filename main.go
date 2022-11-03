package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("commnad line args error")
		os.Exit(1)
	}
	bucketName := args[0]
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()
	bucket := client.Bucket(bucketName)
	iter := bucket.Objects(ctx, nil)
	for {
		objAttrs, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(objAttrs.Name)
	}
}
