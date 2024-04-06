package test

import (
	"fmt"
	"testing"

	jsonparser "github.com/aa/v2/json_parser"
)

func TestParse(t *testing.T) {
	type expectation struct {
		path  string
		valid bool
	}
	type data struct {
		name  string
		paths []expectation
	}
	testData := []data{
		{
			name: "Step 1",
			paths: []expectation{
				{
					path:  "tests/step1/invalid.json",
					valid: false,
				},
				{
					path:  "tests/step1/valid.json",
					valid: true,
				},
			},
		},
		{
			name: "Step 2",
			paths: []expectation{
				{
					path:  "tests/step2/invalid.json",
					valid: false,
				},
				{
					path:  "tests/step2/invalid2.json",
					valid: false,
				},
				{
					path:  "tests/step2/valid.json",
					valid: true,
				},
				{
					path:  "tests/step2/valid2.json",
					valid: true,
				},
			},
		},
		{
			name: "Step 3",
			paths: []expectation{
				{
					path:  "tests/step3/invalid.json",
					valid: false,
				},
				{
					path:  "tests/step3/valid.json",
					valid: true,
				},
			},
		},
		{
			name: "Step 4",
			paths: []expectation{
				{
					path:  "tests/step4/invalid.json",
					valid: false,
				},
				{
					path:  "tests/step4/valid.json",
					valid: true,
				},
				{
					path:  "tests/step4/valid2.json",
					valid: true,
				},
			},
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			for _, p := range tt.paths {
				parser := jsonparser.NewParser(p.path)
				parser.ParseFromReader()

			}
			fmt.Println("\n\nTesting completed for ", tt.name, ".")
		})
	}
}
