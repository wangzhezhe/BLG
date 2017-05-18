package store

import (
	"testing"
)

func TestStore(t *testing.T) {

	go Start()
	Watch()
}
