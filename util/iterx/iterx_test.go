// Copyright JAMF Software, LLC

package iterx

import (
	"iter"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFrom(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{items: []int{}},
		},
		{
			name: "single item",
			args: args{items: []int{1}},
			want: []int{1},
		},
		{
			name: "multiple items",
			args: args{items: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, Collect(From(tt.args.items...)))
		})
	}
}

func TestConsume(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			args: args{items: []int{}},
		},
		{
			name: "single item",
			args: args{items: []int{1}},
			want: []int{1},
		},
		{
			name: "multiple items",
			args: args{items: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := 0
			Consume(From(tt.args.items...), func(item int) {
				require.Equal(t, tt.want[i], item)
				i++
			})
		})
	}
}

func TestFirst(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty",
			args: args{items: []int{}},
		},
		{
			name: "single item",
			args: args{items: []int{1}},
			want: 1,
		},
		{
			name: "multiple items",
			args: args{items: []int{1, 2, 3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, First(From(tt.args.items...)))
		})
	}
}

func TestMap(t *testing.T) {
	type args[S any, R any] struct {
		seq iter.Seq[S]
		fn  func(S) R
	}
	type testCase[S any, R any] struct {
		name string
		args args[S, R]
		want iter.Seq[R]
	}
	tests := []testCase[int, string]{
		{
			name: "empty",
			args: args[int, string]{
				seq: From[int](),
				fn: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			want: From[string](),
		},
		{
			name: "single item",
			args: args[int, string]{
				seq: From(1),
				fn: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			want: From("1"),
		},
		{
			name: "multiple items",
			args: args[int, string]{
				seq: From(1, 2, 3, 4, 5),
				fn: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			want: From("1", "2", "3", "4", "5"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.args.seq, tt.args.fn)
			require.Equal(t, Collect(tt.want), Collect(got))
		})
	}
}
