package model

//  表结构设计
//- 相册表 photo_albums
//    - id
//    - cover_picture 封面
//    - album_name 名称
//    - count 数量
type PhotoAlbum struct {
	Id           int
	CoverPicture string
	AlbumName    string
	Count        int
}

type Persion struct {
	Name string
}
