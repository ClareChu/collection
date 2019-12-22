package core

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	NewConfig("../config/application.yml")
}
