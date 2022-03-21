package db

import "go-demos/model"

// ListPhotoAlbums 查询相册列表
func ListPhotoAlbums() (list []model.PhotoAlbum, err error) {
	err = DB.Model(model.PhotoAlbum{}).Find(&list).Error
	return
}

// CreatePhotoAlbum 创建相册
func CreatePhotoAlbum(album model.PhotoAlbum) (err error) {
	err = DB.Model(model.PhotoAlbum{}).Create(&album).Error
	return
}
