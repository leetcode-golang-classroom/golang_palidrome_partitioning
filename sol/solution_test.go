package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	s := "aab"
	for idx := 0; idx < b.N; idx++ {
		partition(s)
	}
}
func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "s = \"aab\"",
			args: args{s: "aab"},
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "s = \"a\"",
			args: args{s: "a"},
			want: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
