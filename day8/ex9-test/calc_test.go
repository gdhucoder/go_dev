package main

import "testing"

func TestAdd(t *testing.T) {
	r := add(3, 4)
	if r != 6 {
		t.Fatalf("expecting %d, but %d\n", 6, r)
	}
	t.Logf("test succ")

	// PS D:\project> cd .\src\go_dev\day8
	// PS D:\project\src\go_dev\day8> cd .\ex9-test\
	// PS D:\project\src\go_dev\day8\ex9-test> go test -v
	// === RUN   TestAdd
	// --- PASS: TestAdd (0.00s)
	// 		calc_test.go:10: test succ
	// PASS
	// ok      go_dev/day8/ex9-test    0.506s
	// PS D:\project\src\go_dev\day8\ex9-test>
}
