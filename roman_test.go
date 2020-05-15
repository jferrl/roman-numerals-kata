package roman

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"one", args{1}, "I"},
		{"two", args{2}, "II"},
		{"three", args{3}, "III"},
		{"four", args{4}, "IV"},
		{"five", args{5}, "V"},
		{"nine", args{9}, "IX"},
		{"ten", args{10}, "X"},
		{"twenty one", args{21}, "XXI"},
		{"fifty", args{50}, "L"},
		{"sixty", args{60}, "LX"},
		{"sixty one", args{61}, "LXI"},
		{"one hundred", args{100}, "C"},
		{"five hundred", args{500}, "D"},
		{"one thousand", args{1000}, "M"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert(tt.args.v); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
