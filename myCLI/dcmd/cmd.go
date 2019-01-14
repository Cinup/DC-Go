package dcmd

var (
	Commands = make(map[string]func([]string))
)

func init() {
	Commands["ls"] = ListImage
}
