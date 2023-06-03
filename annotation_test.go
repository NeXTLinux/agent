package main

import (
	"testing"
)

func TestWriteAnnotation(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "writing_annotation", args: args{message: "annotation1"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteAnnotation(tt.args.message)
		})
	}

}
