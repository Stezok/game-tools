package character_test

import (
	"testing"

	"github.com/Stezok/game-tools/internal/redactor/character"
)

func TestGeneration(t *testing.T) {
	err := character.Generate()
	if err != nil {
		t.Error(err)
	}
}
