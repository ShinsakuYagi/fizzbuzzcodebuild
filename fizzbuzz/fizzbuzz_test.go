package fizzbuzz

import (
	"math"
	"testing"
)

func Test_getFizzBuzz(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{1},
			"1",
		},
		{
			"success",
			args{2},
			"2",
		},
		{
			"success",
			args{3},
			"FIZZ!",
		},
		{
			"success",
			args{4},
			"4",
		},
		{
			"success",
			args{5},
			"BUZZ!",
		},
		{
			"success",
			args{15},
			"FIZZ BUZZ!",
		},
		{
			"success",
			args{1545},
			"FIZZ BUZZ!",
		},
		{
			"success",
			args{99999},
			"FIZZ!",
		},
		{
			"success",
			args{100000},
			"BUZZ!",
		},
		{
			"success",
			args{math.MaxInt32},
			"2147483647",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFizzBuzz(tt.args.num); got != tt.want {
				t.Errorf("getFizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
