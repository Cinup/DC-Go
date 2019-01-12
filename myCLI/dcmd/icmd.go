package dcmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myCLI/API"
	"myCLI/dtype"
	"myCLI/utils"
	"net"
	"os"
	"time"
)

func ImageList() {
	addr := net.UnixAddr{"/var/run/docker.sock", "unix"}
	conn, err := net.DialUnix("unix", nil, &addr)
	if err != nil {
		fmt.Println(API.ConnectError)
		//退出程序
		os.Exit(-1)
	}
	_, err = conn.Write([]byte(API.GetImages))
	if err != nil {
		panic(err)
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
