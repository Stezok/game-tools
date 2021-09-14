package item

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
	"reflect"
)

type ItemService struct {
	FilePath string
}

func (service *ItemService) GetItems() (items []Item, err error) {
	var file *os.File
	file, err = os.Open(service.FilePath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	var rows [][]string
	rows, err = reader.ReadAll()
	if err != nil {
		return
	}

	nameIndexMap := make(map[string]int)
	for i, field := range rows[0] {
		nameIndexMap[field] = i
	}

	for i := 1; i < len(rows); i++ {
		var item Item
		itemType := reflect.TypeOf(item)
		itemVal := reflect.ValueOf(&item)
		for ii := 0; ii < itemType.NumField(); ii++ {
			itemTag := itemType.Field(ii).Tag.Get("item")

			if index, ok := nameIndexMap[itemTag]; !ok {
				log.Print(itemTag)
			} else {
				val := rows[i][index]
				itemVal.Elem().Field(ii).SetString(val)
			}
		}
		items = append(items, *itemVal.Interface().(*Item))
	}
	return
}

func (service *ItemService) WriteItems(w io.Writer, items []Item) error {
	bufferArr := make([]byte, 0, 1024*128) // 128 Kb buffer
	buffer := bytes.NewBuffer(bufferArr)

	var emptyItem Item
	itemType := reflect.TypeOf(emptyItem)
	for ii := 0; ii < itemType.NumField(); ii++ {
		itemTag := itemType.Field(ii).Tag.Get("item")
		if ii != 0 {
			buffer.Write([]byte("\t"))
		}
		buffer.Write([]byte(itemTag))
	}
	_, err := io.Copy(w, buffer)
	if err != nil {
		return err
	}

	for _, item := range items {
		buffer.Write([]byte("\n"))
		vitem := reflect.ValueOf(item)
		for i := 0; i < vitem.NumField(); i++ {
			data := vitem.Field(i).String()
			if i != 0 {
				buffer.Write([]byte("\t"))
			}
			buffer.Write([]byte(data))
		}
		_, err = io.Copy(w, buffer)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewItemService(filepath string) *ItemService {
	return &ItemService{
		FilePath: filepath,
	}
}
