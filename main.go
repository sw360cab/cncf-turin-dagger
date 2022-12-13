package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()
	path := "build/"

	// get reference to the local project
	src := client.Host().Directory(".")

	// get `golang` image
	golang := client.Container().From("golang:latest").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src")

	// get reference to build output directory in container
	output := golang.
		WithExec([]string{"go", "build", "-o", path}).
		Directory(path)

	// write contents of container build/ directory to the host
	_, err = output.Export(ctx, path)
	if err != nil {
		return err
	}

	golangExec := golang.WithExec([]string{"go", "test", "./..."})
	code, err := golangExec.ExitCode(ctx)
	if err != nil || code != 0 {
		return fmt.Errorf("go test failed")
	}

	fmt.Println("Building and Testing Looks Good 🤠")
	return nil
}
