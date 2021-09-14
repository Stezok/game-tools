package character_test

import (
	"testing"

	"github.com/Stezok/game-tools/internal/redactor/character"
)

func TestGetCharacters(t *testing.T) {
	service := character.NewCharacterService("CharacterInfo.txt")
	chars, err := service.GetCharacters()
	if err != nil {
		t.Error(err)
	}
	t.Error(chars[0])
}
