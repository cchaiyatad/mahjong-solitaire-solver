package utils

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/assert"
)

func TestAbs(t *testing.T) {
	cases := []struct {
		Given    int
		Expected int
	}{
		{1, 1},
		{0, 0},
		{-1, 1},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("Abs(%d) should get %d", tc.Given, tc.Expected), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, Abs(tc.Given))
		})
	}
}
func TestJSON(t *testing.T) {
	cases := []struct {
		Given    interface{}
		Expected []byte
	}{
		{1, []byte("1")},
		{3.14, []byte("3.14")},
		{[]string{"hello", "world"}, []byte(`["hello","world"]`)},
		{nil, []byte("null")},
		{make(chan int), []byte("null")},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("ToJSON(%v) should get %s", tc.Given, tc.Expected), func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.Expected, ToJSON(tc.Given))
		})
	}
}

func TestFlag(t *testing.T) {
	t.Run("args.Port should equal 8080 when not provide", func(t *testing.T) {
		os.Args = []string{"app"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		args := GetArgs()
		assert.Equal(t, "8080", args.Port)
	})

	t.Run("args.Port should get from args -port flag", func(t *testing.T) {
		os.Args = []string{"app", "-port", "9090"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		args := GetArgs()
		assert.Equal(t, "9090", args.Port)
	})

	t.Run("args.Port should get from env", func(t *testing.T) {
		os.Args = []string{"app"}
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		os.Setenv("PORT", "1234")

		args := GetArgs()
		assert.Equal(t, "1234", args.Port)
	})
}
