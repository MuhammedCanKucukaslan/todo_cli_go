package main

/*
the todo is a command line application that allows you to add, delete, and mark items as complete.
todo -h            # help
todo -v            # version
todo -l            # list all items (un-completed)
todo -c            # list completed items
todo -a"Buy Milk"  # add new item
todo -m TODO-Id    # mark as complete
todo -d TODO-Id    # delete item
*/

import (
	"flag"
	"fmt"
	"os"
)

type TodoList map[int]*item

func main() {
	var todoList TodoList = TodoList{}
	// Parsing the arguments ...
	h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")
	l := flag.Bool("l", false, "list all items (un-completed)")
	c := flag.Bool("c", false, "list completed items")
	a := flag.String("a", "", "add new item")
	m := flag.Int("m", -1, "mark as complete")
	d := flag.Int("d", -1, "delete item")
	flag.Parse()

	if *h {
		printHelp()
	} else if *v {
		printVersion()
	} else if *l {
		printlist(todoList)
	} else if *c {
		printlist(todoList)
	} else if *a != "" {
		addItem(&todoList, *a)
	} else if *m != -1 {
		markItem(&todoList, *m)
	} else if *d != 1 {
		deleteItem(&todoList, *d)
	} else {
		fmt.Println("No arguments given")
		printHelp()
	}
}

func printHelp() {
	fmt.Println("-----------------\nHelp for using the todo cli app")
	fmt.Println("-----------------")
	fmt.Println("todo -h            # help")
	fmt.Println("todo -v            # version")
	fmt.Println("todo -l            # list all items (un-completed)")
	fmt.Println("todo -c            # list completed items")
	fmt.Println("todo -a \"Buy Milk\"  # add new item")
	fmt.Println("todo -m TODO-ID    # mark as complete")
	fmt.Println("todo -d TODO-ID    # delete item")
}

func printVersion() {
	v := "0.1.0"
	date := "2021.11.14"
	fmt.Println("todo ", v, "\nYou're currently using the version ", v, " released on", date)
	fmt.Println("There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.")
}

func printlist(todoList TodoList) {
	printMap(todoList)
}

func addItem(todoList *TodoList, Title string) int {
	length := len(*todoList)
	i := item{length, Title, "today", false}
	(*todoList)[length] = &i
	return length
}

func deleteItem(todoList *TodoList, Id int) {
	delete(*todoList, Id)
}

func markItem(tdmap *TodoList, Id int) {
	(*tdmap)[Id].Status = true
}

func debugPrintArgs() {
	fmt.Println("-----------------\nList of the Args")
	fmt.Println("-----------------")
	for i, a := range os.Args {
		fmt.Println(i, " ", a)
	}
	fmt.Println("-----------------")
}

func printMap(m TodoList) {
	for i, item := range m {
		fmt.Println(i, " ", item.toString())
	}
}
