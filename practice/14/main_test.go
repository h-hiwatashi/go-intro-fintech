package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1

	if Debug {
		t.Skip("Skip: Debug is true")
	}
	v := IsOne(i)

	if !v {
		t.Errorf("IsOne(%d) = %v; want true", i, v)
	}
}


// go test
// テストの詳細
// go test -v

// go test -v ./...

// カバレッジ
// cover rage
// go test -cover ./...