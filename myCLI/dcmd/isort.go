package dcmd

import (
	"myCLI/dtype"
	"myCLI/utils"
)

type ImageSlice []dtype.Image

func (images ImageSlice) Len() int {
	return len(images)
}
func (images ImageSlice) Swap(i, j int) {
	images[i], images[j] = images[j], images[i]
}

type SortByRep struct {
	ImageSlice
}

func (images SortByRep) Less(i, j int) bool {
	return images.ImageSlice[i].RepoDigests[0][:5] <
		images.ImageSlice[j].RepoDigests[0][:5]
}

type SortByRepTag struct {
	ImageSlice
}

func (images SortByRepTag) Less(i, j int) bool {
	return utils.GetTag(images.ImageSlice[i].RepoTags[0]) <=
		utils.GetTag(images.ImageSlice[i].RepoTags[j])
}

type SortById struct {
	ImageSlice
}

func (images SortById) Less(i, j int) bool {
	return images.ImageSlice[i].Id < images.ImageSlice[j].Id
}

type SortByCreate struct {
	ImageSlice
}

func (images SortByCreate) Less(i, j int) bool {
	return images.ImageSlice[i].Created < images.ImageSlice[j].Created
}

type SortBySize struct {
	ImageSlice
}

func (images SortBySize) Less(i, j int) bool {
	return images.ImageSlice[i].Size < images.ImageSlice[j].Size
}
