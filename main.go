package main

import (
	"google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	opts := protogen.Options{}
	opts.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = internal_gengo.SupportedFeatures
		return New(gen).Generate()
	})
}
