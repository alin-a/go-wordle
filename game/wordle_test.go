package game

import (
	"reflect"
	"testing"
)

func Test_compareInputAndWord(t *testing.T) {
	type args struct {
		input string
		word  string
	}
	tests := []struct {
		name string
		args args
		want [5]int
	}{
		{
			name: "Puppe/Stopp",
			args: struct {
				input string
				word  string
			}{input: "puppe", word: "stopp"},
			want: [5]int{1, 0, 0, 2, 0},
		},
		{
			name: "Nadel/Pasta",
			args: struct {
				input string
				word  string
			}{input: "nadel", word: "pasta"},
			want: [5]int{0, 2, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareInputAndWord(tt.args.input, tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareInputAndWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		elems []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Array contains element",
			args: struct {
				elems []int
				value int
			}{
				elems: []int{1, 2, 3},
				value: 2,
			},
			want: true,
		},

		{
			name: "Array does not element",
			args: struct {
				elems []int
				value int
			}{
				elems: []int{1, 2, 3},
				value: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.elems, tt.args.value); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_matchOccurrence(t *testing.T) {
	type args struct {
		inputMap map[string][]int
		wordMap  map[string][]int
	}
	tests := []struct {
		name string
		args args
		want [5]int
	}{
		{
			name: "All characters are different",
			args: struct {
				inputMap map[string][]int
				wordMap  map[string][]int
			}{
				inputMap: map[string][]int{
					"a": {0},
					"b": {1},
					"c": {2},
					"d": {3},
					"e": {4},
				},
				wordMap: map[string][]int{
					"f": {0},
					"g": {1},
					"h": {2},
					"i": {3},
					"j": {4},
				},
			},
			want: [5]int{0, 0, 0, 0, 0},
		},
		{
			name: "All characters occur in both words but in different positions",
			args: struct {
				inputMap map[string][]int
				wordMap  map[string][]int
			}{
				inputMap: map[string][]int{
					"a": {0},
					"b": {1},
					"c": {2},
					"d": {3},
					"e": {4},
				},
				wordMap: map[string][]int{
					"b": {0},
					"c": {1},
					"d": {2},
					"e": {3},
					"a": {4},
				},
			},
			want: [5]int{1, 1, 1, 1, 1},
		},
		{
			name: "All characters occur in both words in the same positions",
			args: struct {
				inputMap map[string][]int
				wordMap  map[string][]int
			}{
				inputMap: map[string][]int{
					"a": {0},
					"b": {1},
					"c": {2},
					"d": {3},
					"e": {4},
				},
				wordMap: map[string][]int{
					"a": {0},
					"b": {1},
					"c": {2},
					"d": {3},
					"e": {4},
				},
			},
			want: [5]int{2, 2, 2, 2, 2},
		},
		{
			name: "One characters occurs in lesser frequency in the word than in the input with one correct position",
			args: struct {
				inputMap map[string][]int
				wordMap  map[string][]int
			}{
				inputMap: map[string][]int{
					"a": {0, 4},
					"b": {1},
					"c": {2},
					"d": {3},
				},
				wordMap: map[string][]int{
					"a": {0},
					"b": {1},
					"c": {2},
					"d": {3},
					"e": {4},
				},
			},
			want: [5]int{2, 2, 2, 2, 0},
		},
		{
			name: "One characters occurs in lesser frequency in the word than in the input with no correct position",
			args: struct {
				inputMap map[string][]int
				wordMap  map[string][]int
			}{
				inputMap: map[string][]int{
					"a": {3, 4},
					"b": {0},
					"c": {1},
					"d": {2},
				},
				wordMap: map[string][]int{
					"a": {0},
					"b": {1},
					"c": {2},
					"d": {3},
					"e": {4},
				},
			},
			want: [5]int{1, 1, 1, 1, 0},
		},

		{
			name: "A character occurs in lesser frequency in input than in word",
			args: struct {
				inputMap map[string][]int
				wordMap  map[string][]int
			}{
				inputMap: map[string][]int{
					"a": {0},
					"b": {1},
					"c": {2},
					"d": {3},
					"e": {4},
				},
				wordMap: map[string][]int{
					"a": {0, 4},
					"b": {1},
					"c": {2},
					"d": {3},
				},
			},
			want: [5]int{2, 2, 2, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchOccurrence(tt.args.inputMap, tt.args.wordMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchOccurrence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toMapOfCharacterAndIndices(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want map[string][]int
	}{
		{
			name: "Word with five different characters",
			args: struct{ word string }{word: "nudel"},
			want: map[string][]int{
				"n": {0},
				"u": {1},
				"d": {2},
				"e": {3},
				"l": {4},
			},
		},
		{
			name: "Word with three different characters",
			args: struct{ word string }{word: "puppe"},
			want: map[string][]int{
				"p": {0, 2, 3},
				"u": {1},
				"e": {4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toMapOfCharacterAndIndices(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toMapOfCharacterAndIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
