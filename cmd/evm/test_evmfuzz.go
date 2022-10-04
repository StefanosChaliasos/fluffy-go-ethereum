package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/tests"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	_ "os"
	_ "testing"
	_ "time"
)

func test() {
	//dirname := "../openethereum/crates/evmfuzz/fuzz/corpus/fuzz_target_1"
	dirname := "../openethereum/crates/evmfuzz/go-errors"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		println("PROCESSING: " + dirname + "/" + file.Name())
		fuzzedInput, err := ioutil.ReadFile(dirname + "/" + file.Name())
		if err != nil {
			panic(err)
		}

		fuzzedInputProto := tests.Fuzzed{}
		err = proto.Unmarshal(fuzzedInput, &fuzzedInputProto)
		if err != nil {
			panic(err)
		}

		fuzzResults := tests.RunFuzz(fuzzedInputProto)
		if len(fuzzResults.Roots) == 0 {
			fmt.Println("Empty result")
		} else {
			marshalOptions := prototext.MarshalOptions{}
			fuzzResultText := marshalOptions.Format(&fuzzResults)
			fmt.Println(fuzzResultText)
		}
	}
}
