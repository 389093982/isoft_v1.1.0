package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"isoft/isoft_iaas_web/controllers"
	"isoft/isoft_iaas_web/controllers/iblog"
	"isoft/isoft_iaas_web/controllers/ifile"
	"isoft/isoft_iaas_web/controllers/ilearning"
	"isoft/isoft_iaas_web/controllers/sso"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/auth/redirectToLogin/", &sso.AuthController{}, "get,post:RedirectToLogin")
	initIFileRouter()
	initIBlogRouter()
	initILearningRouter()
}

func initILearningRouter() {
	beego.Router("/api/ilearning/newCourse", &ilearning.CourseController{}, "get,post:NewCourse")
	beego.Router("/api/ilearning/getMyCourseList", &ilearning.CourseController{}, "get,post:GetMyCourseList")
	beego.Router("/api/ilearning/changeCourseImg", &ilearning.CourseController{}, "get,post:ChangeCourseImg")
	beego.Router("/api/ilearning/uploadVideo", &ilearning.CourseController{}, "get,post:UploadVideo")
	beego.Router("/api/ilearning/endUpdate", &ilearning.CourseController{}, "get,post:EndUpdate")
	beego.Router("/api/ilearning/showCourseDetail", &ilearning.CourseController{}, "get,post:ShowCourseDetail")

	beego.Router("/api/ilearning/toggle_favorite", &ilearning.CourseController{}, "get,post:ToggleFavorite")

	beego.Router("/api/ilearning/filterCommentTheme", &ilearning.CommentController{}, "get,post:FilterCommentTheme")
	beego.Router("/api/ilearning/addCommentReply", &ilearning.CommentController{}, "get,post:AddCommentReply")
	beego.Router("/api/ilearning/filterCommentReply", &ilearning.CommentController{}, "get,post:FilterCommentReply")
}

func initIFileRouter() {
	beego.Router("/api/ifile/fileUpload/", &ifile.IFileController{}, "post:FileUpload")
	beego.Router("/api/ifile/fileUpload2/", &ifile.IFileController{}, "post:FileUpload2")
	beego.Router("/api/ifile/locateShards/", &ifile.IFileController{}, "post:LocateShards")
	beego.Router("/api/ifile/fileDownload/", &ifile.IFileController{}, "get:FileDownload")
	beego.Router("/api/ifile/fileDownload/", &ifile.IFileController{}, "get:FileDownload")

	beego.Router("/api/heartbeat/sendHeartBeat/", &ifile.HeartBeatController{}, "post:SendHeartBeat")
	beego.Router("/api/heartbeat/queryAllAliveHeartBeat/", &ifile.HeartBeatController{}, "get:QueryAllAliveHeartBeat")

	beego.Router("/api/metadata/searchLatestVersion/", &ifile.MetadataController{}, "post:SearchLatestVersion")
	beego.Router("/api/metadata/getMetadata/", &ifile.MetadataController{}, "post:GetMetadata")
	beego.Router("/api/metadata/putMetadata/", &ifile.MetadataController{}, "post:PutMetadata")
	beego.Router("/api/metadata/addVersion/", &ifile.MetadataController{}, "post:AddVersion")
	beego.Router("/api/metadata/searchAllVersions/", &ifile.MetadataController{}, "post:SearchAllVersions")
	beego.Router("/api/metadata/delMetadata/", &ifile.MetadataController{}, "post:DelMetadata")
	beego.Router("/api/metadata/hasHash/", &ifile.MetadataController{}, "post:HasHash")
	beego.Router("/api/metadata/searchHashSize/", &ifile.MetadataController{}, "post:SearchHashSize")
	beego.Router("/api/metadata/searchVersionStatus/", &ifile.MetadataController{}, "post:SearchVersionStatus")
	beego.Router("/api/metadata/filterPageMetadatas/", &ifile.MetadataController{}, "post:FilterPageMetadatas")
}

func initIBlogRouter() {
	beego.Router("/api/catalog/edit", &iblog.CatalogController{}, "get:Edit;post:PostEdit")
	beego.Router("/api/catalog/getMyCatalogs", &iblog.CatalogController{}, "get:GetMyCatalogs")
	beego.Router("/api/catalog/delete", &iblog.CatalogController{}, "post:PostDelete")

	beego.Router("/api/blog/edit", &iblog.BlogController{}, "get:Edit;post:PostEdit")
	beego.Router("/api/blog/blogList", &iblog.BlogController{}, "get:BlogList")
	beego.Router("/api/blog/getMyBlogs", &iblog.BlogController{}, "get:GetMyBlogs")
	beego.Router("/api/blog/delete", &iblog.BlogController{}, "post:PostDelete")
	beego.Router("/api/blog/search", &iblog.BlogController{}, "get:Search")
	beego.Router("/api/blog/publish", &iblog.BlogController{}, "post:PostPublish")
	beego.Router("/api/blog/showBlogDetail", &iblog.BlogController{}, "get:ShowBlogDetail")
}
