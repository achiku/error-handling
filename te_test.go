package main

import "testing"

func TestHandleError(t *testing.T) {
	err := handleError()
	t.Log(err)
}

func TestHandleNormal(t *testing.T) {
	err := handleNormal()
	t.Log(err)
}
