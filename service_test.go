package main

import "testing"

func TestName(t *testing.T) {
	Tunnel("root", "", ":22", ":3306", "localhost:13306")
}
