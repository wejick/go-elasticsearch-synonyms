package elasticsearchSynonyms

import (
	"reflect"
	"testing"
)

func TestContract(t *testing.T) {
	type args struct {
		lhs         []string
		replacement string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "contract empty",
			args: args{
				lhs:         []string{""},
				replacement: "",
			},
			wantOutput: " => ",
		},
		{
			name: "contract gb",
			args: args{
				lhs:         []string{"g b", "great britain", "gb"},
				replacement: "gb",
			},
			wantOutput: "g b,great britain,gb => gb",
		},
	}
	for _, tt := range tests {
		if gotOutput := Contract(tt.args.lhs, tt.args.replacement); gotOutput != tt.wantOutput {
			t.Errorf("%q. Contract() = %v, want %v", tt.name, gotOutput, tt.wantOutput)
		}
	}
}

func TestExpand(t *testing.T) {
	type args struct {
		stringArray []string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name: "expand empty",
			args: args{
				stringArray: []string{""},
			},
			wantOutput: " => ",
		},
		{
			name: "expand gb",
			args: args{
				stringArray: []string{"g b", "great britain", "gb"},
			},
			wantOutput: "g b,great britain,gb => g b,great britain,gb",
		},
	}
	for _, tt := range tests {
		if gotOutput := Expand(tt.args.stringArray); gotOutput != tt.wantOutput {
			t.Errorf("%q. Expand() = %v, want %v", tt.name, gotOutput, tt.wantOutput)
		}
	}
}

func TestExpandString(t *testing.T) {
	tests := []struct {
		name       string
		lhs        string
		wantOutput string
	}{
		{
			name:       "expand empty",
			lhs:        "",
			wantOutput: "",
		}, {
			name:       "expand gb",
			lhs:        "gb britain english",
			wantOutput: "gb,britain,english",
		},
	}
	for _, tt := range tests {
		if gotOutput := ExpandString(tt.lhs); gotOutput != tt.wantOutput {
			t.Errorf("%q. ExpandString() = %v, want %v", tt.name, gotOutput, tt.wantOutput)
		}
	}
}

func TestExplicit(t *testing.T) {
	type args struct {
		lhs []string
		rhs []string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name:       "explicit empty",
			wantOutput: " => ",
		},
		{
			name: "explicit rhs gb empty",
			args: args{
				lhs: []string{"gb", "g b", "great britain"},
			},
			wantOutput: "gb,g b,great britain => gb,g b,great britain",
		},
		{
			name: "explicite gb full",
			args: args{
				lhs: []string{"gb", "g b", "great britain"},
				rhs: []string{"britain", "england", "wales"},
			},
			wantOutput: "gb,g b,great britain => britain,england,wales",
		},
	}
	for _, tt := range tests {
		if gotOutput := Explicit(tt.args.lhs, tt.args.rhs); gotOutput != tt.wantOutput {
			t.Errorf("%q. Explicit() = %v, want %v", tt.name, gotOutput, tt.wantOutput)
		}
	}
}

func TestStringToArray(t *testing.T) {
	type args struct {
		toConvert string
		delimiter string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []string
	}{
		{
			name:       "string to array empty",
			wantOutput: []string{},
		},
		{
			name: "stringToArray gb",
			args: args{
				toConvert: "gb|britain|great britain",
				delimiter: "|",
			},
			wantOutput: []string{"gb", "britain", "great britain"},
		},
	}
	for _, tt := range tests {
		if gotOutput := StringToArray(tt.args.toConvert, tt.args.delimiter); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
			t.Errorf("%q. StringToArray() = %v, want %v", tt.name, gotOutput, tt.wantOutput)
		}
	}
}
