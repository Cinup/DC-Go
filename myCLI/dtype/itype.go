package dtype

//type Labels struct {
//	value map[string] string
//}
/*
	目前还无法解析到Labels,等待解决
*/
type Image struct {
	//Containers int8
	//VirtualSize int64
	//SharedSize  int8
	//Labels		Labels
	//ParentId    string
	RepoDigests []string
	RepoTags    []string
	Id          string
	Created     int64
	Size        int64
}
