package event_manager

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBaseManager(t *testing.T) {
	cem := NewCommandEventManager()

	cem.Add("no-command", func(_ *Command) {
		fmt.Println("no command was recieved")
	})

	cem.Add("any-command", func(c *Command) {
		fmt.Printf("the %s command was executed", c.Kind)
	})

	cem.Add("sum", func(c *Command) {
		a, _ := strconv.Atoi(c.Args[0])
		b, _ := strconv.Atoi(c.Args[1])
		fmt.Printf("the sum result is: %d", a+b)
	})

	cem.Run()
}
