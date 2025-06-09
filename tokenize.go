package argv

import (
	"fmt"
)

func Tokenize(input string) ([]string, error) {
	i := -1
	from := i + 1
	tokens := make([]string, 0, 8)

	capture := func() {
		if i-from <= 0 {
			return
		}
		tokens = append(tokens, input[from:i])
		from = i + 1
	}

	for {
		i += 1
		if i >= len(input) {
			capture()
			return tokens, nil
		}

		// If we see an unquoted string, slice it as the next arg
		if input[i] == ' ' {
			capture()
		}

		// If we find a quote (single or double) - parse it as a string literal
		if input[i] == '\'' || input[i] == '"' {
			// Drop any dangling tokens
			capture()

			// Set the terminator we'll look for
			//
			// TODO: Handle manual escaping
			term := input[i]

			// Relocate the `from` position to inside the quote
			from = i + 1

			// Consume until we hit a terminator
			for {
				i += 1
				if i >= len(input) {
					// TODO: Produce better errors (annotate the bad string position)
					return nil, fmt.Errorf(
						"reached EOL while looking for matching [%c] in string [%s]",
						term, input,
					)
				}

				// We've reached our terminator!
				if input[i] == term {
					capture()
					// We offset `from` here since we want to hop the trailing terminator
					from++
					break
				}
			}
		}
	}
}
