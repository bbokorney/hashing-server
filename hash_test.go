package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	// Normally I'd use an assertion library like testify (https://github.com/stretchr/testify)
	// to make these assertions more concise
	base64Encoded := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	hashedPassword := Hash("angryMonkey")
	if hashedPassword != base64Encoded {
		t.Fatalf("Expected %s to match %s", base64Encoded, hashedPassword)
	}
}
