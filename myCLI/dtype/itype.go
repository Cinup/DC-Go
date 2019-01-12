package dtype


//type Labels struct {
//	value map[string] string
//}
/*
	目前还无法解析到Labels,等待解决
*/
type Image struct {
	Containers int8
	Created    int64
	Id         string
	//Labels		Labels
	ParentId    string
	RepoDigests []string
	RepoTags    []string
	SharedSize  int8
	Size        int64
	VirtualSize int64
}