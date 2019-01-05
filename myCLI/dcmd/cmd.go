package dcmd

var (
	Commands map[string]func()
)

func init() {
	Commands = map[string]func(){}
	Commands["ls"] = ImageList
}
