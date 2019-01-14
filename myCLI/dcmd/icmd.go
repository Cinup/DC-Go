package dcmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myCLI/api"
	"myCLI/dtype"
	"myCLI/utils"
	"net"
	"os"
	"sort"
	"time"
)

var (
	name    = "/var/run/docker.sock"
	network = "unix"
	para    = make(map[string]int,5)
)

func init() {
	//para["-s"] = 0
	para["-s1"] = 1
	para["-s2"] = 2
	para["-s3"] = 3
	para["-s4"] = 4
	para["-s5"] = 5
	//println("init")
}
func ListImage(args []string) {
	addr := net.UnixAddr{name, network}
	conn, err := net.DialUnix("unix", nil, &addr)
	if err != nil {
		fmt.Println(api.ConnectError)
		//退出程序
		os.Exit(-1)
	}
	_, err = conn.Write([]byte(api.GetImages))
	if err != nil {
		fmt.Println(api.WriteError)
		os.Exit(-1)
	}
	result, errConn := ioutil.ReadAll(conn)
	if errConn != nil {
		panic(errConn)
	}
	body := utils.GetBody(result[:])
	var images []dtype.Image
	errJson := json.Unmarshal(body, &images)
	if errJson != nil {
		panic(errJson)
	}
	//如果只有一个参数,说明只有ls指令
	if len(args) == 1 {
		listImage(images)
	} else {
		//fmt.Println(para)
		value := para[args[1]]
		switch value {
		case 0, 1:
			sort.Sort(SortByRep{images})
		case 2:
			sort.Sort(SortByRepTag{images})
		case 3:
			sort.Sort(SortById{images})
		case 4:
			sort.Sort(SortByCreate{images})
		case 5:
			sort.Sort(SortBySize{images})
		default:
			sort.Sort(SortByRep{images})
		}
		listImage(images)
	}
}

func listImage(images []dtype.Image) {
	fmt.Println("Number of Images: ", len(images))
	fmt.Printf("%-20s%-20s%-20s%-20s%s",
		"REPOSITORY", "TAG", "IMAGE ID", "CREATED", "SIZE")
	fmt.Println()
	for i := 0; i < len(images); i++ {
		fmt.Printf("%-20s%-20s%-20s%-20s%s",
			images[i].RepoDigests[0][:5],
			utils.GetTag(images[i].RepoTags[0]),
			images[i].Id[7:19],
			time.Unix(images[i].Created, 0).Format("2006-01-02"),
			utils.CalSize(images[i].Size))
		fmt.Println()
	}
}
