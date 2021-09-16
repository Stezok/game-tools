package character

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type CharacterService struct {
	FilePath string
}

func (cs *CharacterService) GetCharacters() (characters []Character, err error) {
	var file *os.File
	file, err = os.Open(cs.FilePath)
	if err != nil {
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	strData := string(data)
	lines := strings.Split(strData, "\r\n")
	if len(lines) == 0 {
		err = errors.New("CHARACTER: empty file")
		return
	}

	head := strings.Split(lines[0], "\t")
	var rows [][]string
	for i := 1; i < len(lines); i++ {
		rows = append(rows, strings.Split(lines[i], "\t"))
	}

	nameIndexMap := make(map[string]int)
	for i, field := range head {
		nameIndexMap[field] = i
	}
	for i := 0; i < len(rows); i++ {
		var character Character
		characterType := reflect.TypeOf(character)
		characterVal := reflect.ValueOf(&character)
		for ii := 0; ii < characterType.NumField(); ii++ {
			itemTag := characterType.Field(ii).Tag.Get("character")

			if index, ok := nameIndexMap[itemTag]; !ok {
			} else {
				val := rows[i][index]
				characterVal.Elem().Field(ii).SetString(val)
			}
		}
		characters = append(characters, *characterVal.Interface().(*Character))
	}
	return
}

func (cs *CharacterService) encode(s string) []byte {
	return []byte(s)
}

func (cs *CharacterService) WriteCharacters(w io.Writer, characters []Character) error {
	bufferArr := make([]byte, 0, 1024*128) // 128 Kb buffer
	buffer := bytes.NewBuffer(bufferArr)

	var emptyCharacter Character
	characterType := reflect.TypeOf(emptyCharacter)
	for ii := 0; ii < characterType.NumField(); ii++ {
		characterTag := characterType.Field(ii).Tag.Get("character")
		if ii != 0 {
			buffer.Write(cs.encode("\t"))
		}
		buffer.Write(cs.encode(characterTag))
	}
	_, err := io.Copy(w, buffer)
	if err != nil {
		return err
	}

	for _, character := range characters {
		buffer.Write(cs.encode("\r\n"))
		vcharacter := reflect.ValueOf(character)
		for i := 0; i < vcharacter.NumField(); i++ {
			data := vcharacter.Field(i).String()
			if i != 0 {
				buffer.Write(cs.encode("\t"))
			}
			buffer.Write(cs.encode(data))
		}
		_, err = io.Copy(w, buffer)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewCharacterService(filePath string) *CharacterService {
	return &CharacterService{
		FilePath: filePath,
	}
}
