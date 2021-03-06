package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type TodoList map[int]*item

func (td TodoList) writeToFile(filename string) error {
	confDir, e := os.UserConfigDir() // get the config directory
	//fmt.Println("confDir: ", confDir)
	//fmt.Println("write e: ", e)
	if e == nil {
		os.Chdir(confDir)
		os.Mkdir("todo", 0777)
		os.Chdir("todo")
	}
	str := "["
	for _, it := range td {
		str += it.toJSON() + ","
	}
	// we need to remove the last comma
	if len(str) > 1 {
		str = str[:len(str)-1]
	}
	str += "]"
	return ioutil.WriteFile(filename, []byte(str), 0666)

}

func readFromFile(td TodoList, filename string) TodoList {

	// Reading the file
	confDir, e := os.UserConfigDir()
	//fmt.Println("confDir: ", confDir)
	//fmt.Println("read e: ", e)
	if e == nil {
		os.Chdir(confDir)
		os.Mkdir("todo", 0777)
		os.Chdir("todo")
	}
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// maybe we should have a log file, right?
		//fmt.Println("Error: ", err)
		//fmt.Println("Cannot read ", filename)
		return td
	}
	str := string(byteSlice)
	str = strings.Trim(str, "[ ]")
	str = strings.Replace(str, "},{", "}|{", -1)
	arr := strings.Split(str, "|")

	for _, i := range arr {
		//fmt.Println("\n\n", xx, " ", i)
		it := parseJSON(i)
		if it != nil {
			td[it.Id] = it
		}
	}
	return td
}

func (td TodoList) addItem(Title string) int {
	index := 0
	if len(td) == 0 {
		index = 0
	} else {
		for i := range td {
			if i >= index {
				index = i
			}
		}
	}
	index++ // it either starts with 1 or it is 1 more than the largest index
	// store date in the format of ""2021-11-14T11:59:37+03:00", namely time.RFC3339
	it := item{index, Title, time.Now().Format(time.RFC3339), false}
	td[it.Id] = &it

	return index
}

func (td TodoList) markItem(id int) {
	td[id].Status = true
}

func (td TodoList) unMarkItem(id int) {
	td[id].Status = false
}

func (td TodoList) deleteItem(id int) {
	delete(td, id)
}

func (td *TodoList) printInComplete() {
	iW := 6
	tW := 30
	dW := 20
	seperator := "|"
	str := fmt.Sprintf("%-*s"+seperator+" %-*s"+seperator+" %-*s", iW, "Id", tW, "Title", dW, "Date")
	// + 1 for the extra space following the seperator
	str += fmt.Sprintf("\n%s:%s:%s", strings.Repeat("-", iW), strings.Repeat("-", 1+tW), strings.Repeat("-", 1+dW))
	fmt.Println(str)
	for _, item := range *td {
		if !item.Status {
			fmt.Println(item.toFormattedString(seperator, iW, tW, dW))
		}
	}
	//printMap(todoList)
}
func (td *TodoList) printComplete() {
	iW := 6
	tW := 30
	dW := 20
	seperator := "|"
	str := fmt.Sprintf("%-*s"+seperator+" %-*s"+seperator+" %-*s", iW, "Id", tW, "Title", dW, "Date")
	// + 1 for the extra space following the seperator
	str += fmt.Sprintf("\n%s:%s:%s", strings.Repeat("-", iW), strings.Repeat("-", 1+tW), strings.Repeat("-", 1+dW))
	fmt.Println(str)

	for _, item := range *td {
		if item.Status {
			fmt.Println(item.toFormattedString(seperator, iW, tW, dW))
		}
	}
	//printMap(todoList)
}
