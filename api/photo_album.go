package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-demos/db"
	"go-demos/model"
)

// PhotoAlbumList 相册列表
func PhotoAlbumList(ctx *gin.Context) {

	list, err := db.ListPhotoAlbums()
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.JSON(200, list)
}

//http:
//	header 小数据
//	url 小数据
//	body  大的数据

// CreatePhotoAlbumArgs 创建的参数，数据一般存在 http请求的body里
type CreatePhotoAlbumArgs struct {
	Name         string `json:"name"`          // 存body
	CoverPicture string `json:"cover_picture"` // 存body
}

// CreatePhotoAlbum 创建相册
func CreatePhotoAlbum(ctx *gin.Context) {
	var args CreatePhotoAlbumArgs
	err := ctx.ShouldBind(&args)
	if err != nil {
		fmt.Println("CreatePhotoAlbum 读取body错误", err.Error())
		ctx.Status(400)
		return
	}

	fmt.Println("从 body 读取到", args.Name)

	var album = model.PhotoAlbum{AlbumName: args.Name, CoverPicture: args.CoverPicture, Count: 0}

	err = db.CreatePhotoAlbum(album)
	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.Status(200)
	return
}
