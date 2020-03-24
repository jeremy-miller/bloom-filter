package filter

import (
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := newFilter(1024)
	bf.add([]byte("hello"))
	bf.add([]byte("world"))
	bf.add([]byte("sir"))
	bf.add([]byte("madam"))
	bf.add([]byte("io"))

	testCases := []struct {
		arg  []byte
		want bool
	}{
		{[]byte("hello"), true},
		{[]byte("world"), true},
		{[]byte("hi"), false},
	}

	for _, tc := range testCases {
		tc := tc // copy value for parallel tests
		t.Run(fmt.Sprintf("arg=%s", tc.arg), func(t *testing.T) {
			t.Parallel()
			got := bf.test(tc.arg)
			if got != tc.want {
				t.Errorf("bf.test(%s) != %t", tc.arg, tc.want)
			}
		})
	}
}
