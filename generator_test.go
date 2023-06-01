package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func runProtoc(protoFiles []string, parameter string) error {
	args := append([]string{
		"-I", "testdata",
		"-I", "protobuf",
		"--plugin=bin/protoc-gen-graphql",
		fmt.Sprintf("--graphql_out=%s:testdata", parameter),
	}, protoFiles...)
	cmd := exec.Command("protoc", args...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func itGeneratesTheCorrectOutput(t *testing.T, name, parameter string) {
	protoFiles, err := filepath.Glob(filepath.Join("testdata", name, "*.proto"))
	if err != nil {
		t.Error(err)
	}

	if err := runProtoc(protoFiles, parameter); err != nil {
		t.Error(err)
	}

	for _, proto := range protoFiles {
		graphql, err := ioutil.ReadFile(strings.TrimSuffix(proto, ".proto") + "_pb.graphql")
		if err != nil {
			t.Error(err)
		}

		expected, err := ioutil.ReadFile(strings.TrimSuffix(proto, ".proto") + ".golden")
		if err != nil {
			t.Error(err)
		}

		if string(graphql) != string(expected) {
			t.Errorf("expected %s to equal %s", graphql, expected)
		}
	}
}

func TestBasicProtobufTypes(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "basic", "")
}

func TestMessagesWithCycles(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "cycle", "")
}

func TestInputTypesForGrpcServices(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "grpc", "")
}

func TestWrappersParameter(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "wrappers", "null_wrappers,input_mode=all")
}

func TestProtobufExtensions(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "extensions", "root_type_prefix,input_mode=all")
}

func TestFieldNamePreserve(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "field_name_preserve", "field_name=preserve,root_type_prefix=")
}

func TestRootTypePrefix(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "root_type_prefix", "root_type_prefix=")
}

func TestProto2(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "proto2", "")
}

func TestTrimPrefix(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "trim_prefix", "trim_prefix=ProtocGenGraphql")
}

func TestEmptyProto(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "empty", "")
}

func TestNullableListTypes(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "nullable_list_types", "nullable_list_types")
}

func TestOptionalFields(t *testing.T) {
	itGeneratesTheCorrectOutput(t, "optional", "")
}
