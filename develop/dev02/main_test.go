package main

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		str []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "начинается с цифры", args: args{[]rune("34dgreg4\\rf")}, want: false},
		{name: "начинается с буквы", args: args{[]rune("dgreg4\\rf")}, want: true},
		{name: "начинается с escape последовательности", args: args{[]rune("\\3dgreg4f")}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.str); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repitSign(t *testing.T) {
	type args struct {
		sign    rune
		counter int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "нулевой кол-во повторов", args: args{sign: '0', counter: 0}, want: ""},
		{name: "повторение 5 раз", args: args{sign: '\\', counter: 5}, want: "\\\\\\\\\\"},
		{name: "повторени 3 раза", args: args{sign: 'a', counter: 3}, want: "aaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repitSign(tt.args.sign, tt.args.counter); got != tt.want {
				t.Errorf("repitSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringHandler(t *testing.T) {
	type args struct {
		str string
	}
	type expected struct {
		wantStr string
		wantErr bool
	}
	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{args: args {"a4bc2d5e"}, expected: expected {wantStr: "aaaabccddddde", wantErr: false}},
		{args: args {"abcd"}, expected: expected {wantStr: "abcd", wantErr: false}},
		{name: "некорректная строка", args: args {"45"}, expected: expected {wantStr: "", wantErr: true}},
		{name: "пустая строка", args: args {""}, expected: expected {wantStr: "", wantErr: false}},
		{name: "эскейп плследовательность'", args: args {"qwe\\4\\5"}, expected: expected {wantStr: "qwe45", wantErr: false}},
		{name: "эскейп плследовательность'", args: args {"qwe\\45"}, expected: expected {wantStr: "qwe44444", wantErr: false}},
		{name: "эскейп плследовательность'", args: args {"qwe\\\\5"}, expected: expected {wantStr: "qwe\\\\\\\\\\", wantErr: false}},
		{name: "эскейп плследовательность'", args: args {"qwe\n2def"}, expected: expected {wantStr: "qwe\n\ndef", wantErr: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringHandler(tt.args.str)
			if (err != nil) != tt.expected.wantErr {
				t.Errorf("stringHandler() error = %v, wantErr %v", err, tt.expected.wantErr)
				return
			}
			if got != tt.expected.wantStr {
				t.Errorf("stringHandler() = %v, want %v", got, tt.expected.wantStr)
			}
		})
	}
}
