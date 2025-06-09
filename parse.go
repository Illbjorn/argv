package argv

import (
	"fmt"
	"strings"
)

func Parse(tokens []string) (Command, error) {
	cmd := Command{
		Flags: map[string][]string{},
	}

	i := -1
	for {
		i += 1
		if i >= len(tokens) {
			return cmd, nil
		}

		this := tokens[i]
		next := ""
		if i+1 < len(tokens) {
			next = tokens[i+1]
		}
		thisIsFlag := strings.HasPrefix(this, "-")
		nextIsFlag := strings.HasPrefix(next, "-")

		switch {
		default:
			// This should be unreachable...
			return cmd, fmt.Errorf("found unexpected token [%s][%s]", this, next)

		case !thisIsFlag:
			// Positional arg
			if cmd.Name == "" {
				cmd.Name = this
			} else {
				cmd.Args = append(cmd.Args, this)
			}

		case next != "" && !nextIsFlag:
			// int/string flag
			name := strings.TrimLeft(this, "-")
			cmd.Flags[name] = append(cmd.Flags[name], next)
			// Offset our position to account for consuming the next2 value
			i++

		case next == "" || nextIsFlag:
			// bool flag
			name := strings.TrimLeft(this, "-")
			cmd.Flags[name] = append(cmd.Flags[name], "true")
		}
	}
}
