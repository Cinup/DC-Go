package dcmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myCLI/dtype"
	"net"
	"time"
)

func ImageList() {
	addr := net.UnixAddr{"/var/run/docker.sock", "unix"}
	conn, err := net.DialUnix("unix", nil, &addr)
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte("GET /images/json HTTP/1.0\r\n\r\n"))
	if err != nil {
		panic(err)
	}
	result, err_conn := ioutil.ReadAll(conn)
	if err_conn != nil {
		panic(err_conn)
	}
	body := getBody(result[:])

	var images []dtype.Image

	err_js := json.Unmarshal(body, &images)
	if err_js != nil {
		panic(err_js)
	}
	fmt.Println("Number of Images: ", len(images))
	fmt.Printf("%-20s%-20s%-20s%-20s%s",
		"REPOSITORY", "TAG", "IMAGE ID", "CREATED", "SIZE")
	fmt.Println()
	fmt.Printf("%-20s%-20s%-20s%-20s%d%s",
		images[0].RepoDigests[0][:5],
		images[0].RepoTags[0],
		images[0].Id[7:19],
		time.Unix(images[0].Created, 0).Format("2006-01-02"),
		images[0].Size/1000/1000, "MB")
	fmt.Println()

}
func getBody(result []byte) (body []byte) {
	for i := 0; i <= len(result)-4; i++ {
		if result[i] == 91 && result[i+1] == 123 {
			body = result[i:]
			break
		}
	}
	return
}