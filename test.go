package test

import (
	"fmt"
	"testing"

	jsonparser "github.com/aa/v2/json_parser"
)

func TestParse(t testing.T) {
	type expectation struct {
		path string
		res  bool
	}
	type data struct {
		name  string
		paths []expectation
	}
	testData := []data{
		{
			name: "test 1",
			paths: []expectation{
				{
					path: "tests/step1/invalid.json",
					res:  false,
				},
				{
					path: "tests/step1/valid.json",
					res:  true,
				},
			},
		},
		{
			name: "test 2",
			paths: []expectation{
				{
					path: "json_parser/tests/step2/invalid.json",
					res:  false,
				},
				{
					path: "json_parser/tests/step2/invalid2.json",
					res:  false,
				},
				{
					path: "json_parser/tests/step2/valid.json",
					res:  true,
				},
				{
					path: "json_parser/tests/step2/valid2.json",
					res:  true,
				},
			},
		},
		{
			name: "test 3",
			paths: []expectation{
				{
					path: "json_parser/tests/step3/invalid.json",
					res:  false,
				},
				{
					path: "json_parser/tests/step3/valid.json",
					res:  true,
				},
			},
		},
		{
			name: "test 4",
			paths: []expectation{
				{
					path: "json_parser/tests/step4/invalid.json",
					res:  false,
				},
				{
					path: "json_parser/tests/step4/valid.json",
					res:  true,
				},
				{
					path: "json_parser/tests/step4/valid2.json",
					res:  true,
				},
			},
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			for _, p := range tt.paths {
				res := jsonparser.ParseFromFile(p.path)
				if p.res != res {
					fmt.Printf("t: %v\nName: %s\nPath: %s", t, tt.name, p.path)
					t.Fail()
				}
			}
		})
	}
}
