package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"isoft/isoft_iaas_web/controllers"
	"isoft/isoft_iaas_web/controllers/cms"
	"isoft/isoft_iaas_web/controllers/common"
	"isoft/isoft_iaas_web/controllers/iblog"
	"isoft/isoft_iaas_web/controllers/ifile"
	"isoft/isoft_iaas_web/controllers/ilearning"
	"isoft/isoft_iaas_web/controllers/monitor"
	"isoft/isoft_iaas_web/controllers/share"
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
	initIBlogRouter()
	initILearningRouter()
	initCMSRouter()
	initShareRouter()
	initCommonRouter()
	initMonitorRouter()
	initIFileRouter()
}

func initIFileRouter() {
	beego.Router("/api/ifile/fileUpload", &ifile.IFileController{}, "get,post:FileUpload")
	beego.Router("/api/ifile/filterPageIFiles", &ifile.IFileController{}, "get,post:FilterPageIFiles")
}

func initMonitorRouter() {
	beego.Router("/api/monitor/registerHeartBeat", &monitor.HeartBeatController{}, "get,post:RegisterHeartBeat")
	beego.Router("/api/monitor/filterPageHeartBeat", &monitor.HeartBeatController{}, "get,post:FilterPageHeartBeat")
}

func initCommonRouter() {
	beego.Router("/api/common/showCourseHistory", &common.HistoryController{}, "get,post:ShowCourseHistory")
}

func initShareRouter() {
	beego.Router("/api/share/filterShareList", &share.ShareController{}, "get,post:FilterShareList")
	beego.Router("/api/share/addNewShare", &share.ShareController{}, "get,post:AddNewShare")
	beego.Router("/api/share/showShareDetail", &share.ShareController{}, "get,post:ShowShareDetail")
}

func initCMSRouter() {
	beego.Router("/api/cms/queryAllConfigurations", &cms.CMSController{}, "get,post:QueryAllConfigurations")
	beego.Router("/api/cms/addConfiguration", &cms.CMSController{}, "get,post:AddConfiguration")
	beego.Router("/api/cms/filterConfigurations", &cms.CMSController{}, "get,post:FilterConfigurations")
	beego.Router("/api/cms/queryRandomFrinkLink", &cms.CMSController{}, "get,post:QueryRandomFrinkLink")
}

func initILearningRouter() {
	beego.Router("/api/ilearning/newCourse", &ilearning.CourseController{}, "get,post:NewCourse")
	beego.Router("/api/ilearning/getMyCourseList", &ilearning.CourseController{}, "get,post:GetMyCourseList")
	beego.Router("/api/ilearning/changeCourseImg", &ilearning.CourseController{}, "get,post:ChangeCourseImg")
	beego.Router("/api/ilearning/uploadVideo", &ilearning.CourseController{}, "get,post:UploadVideo")
	beego.Router("/api/ilearning/endUpdate", &ilearning.CourseController{}, "get,post:EndUpdate")
	beego.Router("/api/ilearning/showCourseDetail", &ilearning.CourseController{}, "get,post:ShowCourseDetail")
	beego.Router("/api/ilearning/getAllCourseType", &ilearning.CourseController{}, "get,post:GetAllCourseType")
	beego.Router("/api/ilearning/getAllCourseSubType", &ilearning.CourseController{}, "get,post:GetAllCourseSubType")
	beego.Router("/api/ilearning/getHotCourseRecommend", &ilearning.CourseController{}, "get,post:GetHotCourseRecommend")
	beego.Router("/api/ilearning/searchCourseList", &ilearning.CourseController{}, "get,post:SearchCourseList")

	beego.Router("/api/ilearning/toggle_favorite", &ilearning.CourseController{}, "get,post:ToggleFavorite")

	beego.Router("/api/ilearning/filterCommentTheme", &ilearning.CommentController{}, "get,post:FilterCommentTheme")
	beego.Router("/api/ilearning/addCommentReply", &ilearning.CommentController{}, "get,post:AddCommentReply")
	beego.Router("/api/ilearning/filterCommentReply", &ilearning.CommentController{}, "get,post:FilterCommentReply")
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
