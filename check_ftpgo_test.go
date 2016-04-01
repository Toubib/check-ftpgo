package main

import (
	"testing"
	"time"
)

func TestNagiosExitOk(t *testing.T) {
	nagiosExitOk(time.Duration(1000))
}

func TestNagiosExitError(t *testing.T) {
	nagiosExitError("Error")
}
