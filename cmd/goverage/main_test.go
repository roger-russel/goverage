package main

import "testing"

func TestReadFile(t *testing.T) {
	type args struct {
		fullFileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Case One",
			args: args{
				fullFileName: "../../tests/fixture/case-one.cov",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadFile(tt.args.fullFileName)
		})
	}
}
