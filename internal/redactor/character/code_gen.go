package character

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Generate() error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	in, err := os.Open("base.txt")
	if err != nil {
		return err
	}
	defer in.Close()

	namesFile, err := os.Open("names.txt")
	if err != nil {
		return err
	}
	defer namesFile.Close()

	data, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	strData := string(data)
	tags := strings.Split(strData, "\n")

	data, err = ioutil.ReadAll(namesFile)
	if err != nil {
		return err
	}
	strData = string(data)
	names := strings.Split(strData, "\n")

	base := `type Charecter struct {
		%s
	}`
	rows := ""
	for i := 0; i < len(tags); i++ {
		rows += fmt.Sprintf("%s string `charecter:\"%s\" json:\"%s\"`\n", names[i], tags[i], names[i])

	}
	base = fmt.Sprintf(base, rows)
	file.Write([]byte(base))

	return nil
}
