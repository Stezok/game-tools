package character_test

import (
	"os"
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

func TestGetWriteCharacters(t *testing.T) {
	service := character.NewCharacterService("CharacterInfo.txt")
	chars, err := service.GetCharacters()
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create("CharacterInfo.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	err = service.WriteCharacters(file, chars)
	if err != nil {
		t.Error(err)
	}
}
