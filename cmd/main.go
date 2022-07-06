package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/domenicomastrangelo/taskmanager/internal/db"
	"github.com/domenicomastrangelo/taskmanager/internal/task"
)

func main() {
	db.ConnectDefault()

	if db.DB != nil {
		defer db.DB.Close()
	} else {
		log.Println("Cannot connect to db")
	}

	if len(os.Args) <= 1 {
		log.Println("You need to provide some command")
		return
	}

	var command = os.Args[1]
	tmpArgs := os.Args
	os.Args = []string{}
	os.Args = append(os.Args, tmpArgs[0])
	os.Args = append(os.Args, tmpArgs[2:]...)

	switch command {
	case "login":
	case "list":
		maxResults := 10
		minDays := 10
		minDate := time.Now().Add(time.Duration(-minDays * (int(time.Hour) * 24))).UTC()

		ts := task.List(maxResults, minDate, false)

		for _, t := range ts {
			fmt.Println("------------------------")
			fmt.Println("ID: " + strconv.Itoa(t.ID))
			fmt.Println(t.CreatedAt.Local().Format("Mon Jan _2 15:04:05 MST 2006"))
			fmt.Println(t.Title)
			fmt.Println(t.Message)
			fmt.Println("------------------------")
			fmt.Println()
		}
	case "add":
		titleArg := flag.String("t", "", "Task title")
		messageArg := flag.String("m", "", "Task message")
		flag.Parse()

		if strings.Compare(*titleArg, "") == 0 || (strings.Compare(*messageArg, "") == 0) {
			flag.PrintDefaults()
			return
		}

		t := task.Task{
			Title:   *titleArg,
			Message: *messageArg,
		}
		t.Add()
	case "delete":
	case "set-done":
		IDArg := flag.Int("id", 0, "Task ID")
		doneArg := flag.Bool("done", false, "true or false")
		flag.Parse()

		if *IDArg == 0 {
			flag.PrintDefaults()
			return
		}

		t := task.Task{
			ID:   *IDArg,
			Done: *doneArg,
		}
		t.SetDone()
	case "save":
	case "update":
	}
}
