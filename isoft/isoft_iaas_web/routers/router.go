package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/ifile/fileUpload/", &controllers.IFileController{}, "post:FileUpload")
	beego.Router("/api/ifile/locateShards/", &controllers.IFileController{}, "post:LocateShards")
	beego.Router("/api/ifile/fileDownload/", &controllers.IFileController{}, "get:FileDownload")
	beego.Router("/api/ifile/getImgOrVedioMedia/", &controllers.IFileController{}, "get:GetImgOrVedioMedia")

	beego.Router("/api/heartbeat/sendHeartBeat/", &controllers.HeartBeatController{}, "post:SendHeartBeat")
	beego.Router("/api/heartbeat/queryAllAliveHeartBeat/", &controllers.HeartBeatController{}, "get:QueryAllAliveHeartBeat")

	beego.Router("/api/metadata/searchLatestVersion/", &controllers.MetadataController{}, "post:SearchLatestVersion")
	beego.Router("/api/metadata/getMetadata/", &controllers.MetadataController{}, "post:GetMetadata")
	beego.Router("/api/metadata/putMetadata/", &controllers.MetadataController{}, "post:PutMetadata")
	beego.Router("/api/metadata/addVersion/", &controllers.MetadataController{}, "post:AddVersion")
	beego.Router("/api/metadata/searchAllVersions/", &controllers.MetadataController{}, "post:SearchAllVersions")
	beego.Router("/api/metadata/delMetadata/", &controllers.MetadataController{}, "post:DelMetadata")
	beego.Router("/api/metadata/hasHash/", &controllers.MetadataController{}, "post:HasHash")
	beego.Router("/api/metadata/searchHashSize/", &controllers.MetadataController{}, "post:SearchHashSize")
	beego.Router("/api/metadata/searchVersionStatus/", &controllers.MetadataController{}, "post:SearchVersionStatus")
	beego.Router("/api/metadata/filterPageMetadatas/", &controllers.MetadataController{}, "post:FilterPageMetadatas")
}
