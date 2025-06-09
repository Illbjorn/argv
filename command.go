package argv

type Command struct {
	Name  string
	Flags map[string][]string
	Args  []string
}

var empty = []string{}

func (c Command) Flag(names ...string) ([]string, bool) {
	for _, name := range names {
		if v, ok := c.Flags[name]; ok {
			return v, true
		}
	}
	return empty, false
}
