package main

import "testing"

func Test_looping(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{ "Default parameter", args { "Code.education Rocks" }, "Code.education Rocks" },
		{ "Another parameter", args{ "Hello World!" }, "Hello World!" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := looping(tt.args.msg); got != tt.want {
				t.Errorf("looping() = %v, want %v", got, tt.want)
			}
		})
	}
}
