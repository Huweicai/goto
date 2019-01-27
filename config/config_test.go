package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_typeOf(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want ModelType
	}{
		{
			args: args{make(map[string]interface{})},
			want: TypeVector,
		}, {
			want: TypeUnknown,
		}, {
			args: args{""},
			want: TypeScalar,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typeOf(tt.args.i); got != tt.want {
				t.Errorf("typeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNest_GetURL(t *testing.T) {
	nest, err := NewNest("../config.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	u, _ := nest.GetScalar([]string{"tce", "stream"})
	print(u)
}

func Test_toLower(t *testing.T) {
	assert := require.New(t)
	m := map[string]interface{}{
		"A": "A",
		"b": "B",
		"C": "c",
		"D": map[string]interface{}{
			"E": "e",
			"f": "F",
		},
	}
	assert.NotPanics(func() {
		toLower(m)
		assert.Equal("a", m["a"])
		assert.Equal("f", m["d"].(map[string]interface{})["f"])
	})
}

func TestNest_AddScalar(t *testing.T) {
	nest, err := NewNest("../config.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	nest.AddScalar([]string{"tce", "stream", "a"}, "http://www.baidu.com")
	nest.Flush()
	println(nest.Data)
}
