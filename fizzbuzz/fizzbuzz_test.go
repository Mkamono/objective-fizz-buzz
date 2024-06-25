package fizzbuzz_test

import (
	"testing"

	"github.com/Mkamono/objective-fizz-buzz/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{n: 1},
			want: "1",
		},
		{
			name: "3",
			args: args{n: 3},
			want: "Fizz",
		},
		{
			name: "5",
			args: args{n: 5},
			want: "Buzz",
		},
		{
			name: "15",
			args: args{n: 15},
			want: "FizzBuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzbuzz.FizzBuzz(tt.args.n); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
