package event_manager

import (
	"fmt"
	"strings"
)

type Command struct {
	Kind string
	Args []string
}

type CommandEventManager struct {
	BaseManager[*Command]
}

func (m *CommandEventManager) Run() {
	var (
		inp  string
		args Command
	)

	fmt.Scanln(&inp)

	cmd := strings.Split(inp, ":")

	if l := len(cmd); l == 0 {
		m.Invoke("no-command", nil)
	} else if l > 1 {
		args.Args = strings.Split(cmd[1], " ")
	}

	args.Kind = cmd[0]

	m.Invoke("any-command", &args)
	m.Invoke(args.Kind, &args)
}

func NewCommandEventManager() Manager[*Command] {
	return &CommandEventManager{
		BaseManager: BaseManager[*Command]{lst: make(map[string][]Listener[*Command])},
	}
}
