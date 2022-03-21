package model

// - 图片列表 pictures/ images
//- id 图片id
//- photo_album_id 相册表id 管理相册表
//- picture_url 图片
//- name 图片名称

type Photo struct {
	Id           int
	PhotoAlbumId int
	Name         string
	PictureUrl   string
}
