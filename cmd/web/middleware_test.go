package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mH myHandler
	h := NoSurf(&mH)
	switch h.(type) {
	case http.Handler:
	default:
		t.Error("Type is not http.Handler")
	}
}

func TestSessionLoad(t *testing.T) {
	var mH myHandler
	h := SessionLoad(&mH)
	switch h.(type) {
	case http.Handler:
	default:
		t.Error("Type is not http.Handler")
	}
}
