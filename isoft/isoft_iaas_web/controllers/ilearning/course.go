package ilearning

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/satori/go.uuid"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/ilearning"
	"path"
	"strings"
	"time"
)

var UploadFileSavePathImg string
var UploadFileSavePathVideo string

func init() {
	UploadFileSavePathImg = beego.AppConfig.String("UploadFileSavePathImg")
	UploadFileSavePathVideo = beego.AppConfig.String("UploadFileSavePathVideo")
}

type CourseController struct {
	beego.Controller
}

func (this *CourseController) ToggleFavorite() {
	// 获取课程 id
	favorite_id, _ := this.GetInt("favorite_id")
	favorite_type := this.GetString("favorite_type")
	user_name := this.Ctx.Input.Session("UserName").(string)
	flag := ilearning.IsFavorite(user_name, favorite_id, favorite_type)
	if flag {
		ilearning.DelFavorite(user_name, favorite_id, favorite_type)
	} else {
		ilearning.AddFavorite(user_name, favorite_id, favorite_type)
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}

func (this *CourseController) ShowCourseDetail() {
	// 获取课程 id
	id, err := this.GetInt("course_id")
	if err != nil{
		this.Data["json"] = &map[string]interface{}{"status":"SUCCESS"}
		this.ServeJSON()
		return
	}
	course, err := ilearning.QueryCourseById(id)
	cVideos, err := ilearning.QueryCourseVideo(id)
	user_name := this.Ctx.Input.Session("UserName").(string)
	// 课程是否收藏
	flag1 := ilearning.IsFavorite(user_name, id, "course_collect")
	// 课程是否点赞
	flag2 := ilearning.IsFavorite(user_name, id, "course_praise")
	this.Data["json"] = &map[string]interface{}{"status":"SUCCESS", "course":&course,
		"cVideos":&cVideos, "course_collect":flag1, "course_parise":flag2}
	this.ServeJSON()
}

func (this *CourseController) EndUpdate() {
	// 获取课程 id
	id, err := this.GetInt("course_id")
	if err == nil {
		flag := ilearning.EndUpdate(id)
		if flag == true {
			this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
		}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *CourseController) UploadVideo() {
	user_name := this.Ctx.Input.Session("UserName").(string)
	// 获取课程 id
	id, err1 := this.GetInt("id")
	video_number, err2 := this.GetInt("video_number")
	f, fh, err3 := this.GetFile("file")
	defer f.Close()
	// 检查文件格式是否是视频格式
	if path.Ext(fh.Filename) != ".ogg" && path.Ext(fh.Filename) != ".mp4" && path.Ext(fh.Filename) != ".webm" {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errMsg": "视频格式不合法!"}
		this.ServeJSON()
	}
	if err1 == nil && err2 == nil && err3 == nil {
		// fh.Filename 原始文件名,存储时使用 UUID 进行重命名
		u := uuid.NewV4()
		newFileName := u.String() + path.Ext(fh.Filename)
		saveFilePath := path.Join(UploadFileSavePathVideo, newFileName)
		// 与 this.GetFile("file") 保持一致的名字
		err := this.SaveToFile("file", saveFilePath)
		if err == nil {
			// 刷新 DB 记录
			id, flag := ilearning.UploadVideo(id, video_number, "http://localhost:8086/" + saveFilePath, fh.Filename)
			// 刷新评论主题
			topic_theme := ilearning.TopicTheme{}
			topic_theme.TopicId = int(id)
			topic_theme.TopicType = "course_video_topic_type"
			topic_theme.TopicContent = strings.Join([]string{user_name, "@", fh.Filename,
				"视频更新啦，喜欢该课程的小伙伴们不要错过奥，简洁、直观、免费的课程，能让你更快的掌握知识"}, "")
			topic_theme.CreatedBy = user_name
			topic_theme.CreatedTime = time.Now()
			topic_theme.LastUpdatedBy = user_name
			topic_theme.LastUpdatedTime = time.Now()
			// 增加一条评论主题
			ilearning.AddTopicTheme(&topic_theme)

			if flag == true {
				this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "msg": "保存成功!"}
			} else {
				this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": "保存失败!"}
			}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": "保存失败!"}
		}
		this.ServeJSON()
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": "保存失败!"}
		this.ServeJSON()
	}
}

func (this *CourseController) ChangeCourseImg() {
	id, _ := this.GetInt("id")
	f, fh, err := this.GetFile("file")
	defer f.Close()
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
		this.ServeJSON()
	} else {
		// 与 this.GetFile("file") 保持一致的名字
		// fh.Filename 原始文件名,存储时使用 UUID 进行重命名
		u := uuid.NewV4()
		newFileName := u.String() + path.Ext(fh.Filename)
		saveFilePath := path.Join(UploadFileSavePathImg, newFileName)
		err := this.SaveToFile("file", saveFilePath)
		// 更新图片
		flag := ilearning.ChangeImage(id, "http://localhost:8086/" + saveFilePath)
		if err == nil && flag == true {
			this.Data["json"] = &map[string]interface{}{"path": saveFilePath, "status": "SUCCESS"}
		} else {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
		}
		this.ServeJSON()
	}
}

func (this *CourseController) GetMyCourseList()  {
	condArr := make(map[string]string)
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	condArr["CourseAuthor"] = this.Ctx.Input.Session("UserName").(string)
	courses, count, err := ilearning.QueryCourse(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	//初始化
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status":"ERROR"}
	}else{
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status":"SUCCESS","courses":&courses,"paginator":&paginatorMap}
	}
	this.ServeJSON()
}


func (this *CourseController) NewCourse() {
	user_name := this.Ctx.Input.Session("UserName").(string)
	var course ilearning.Course
	course_name := this.GetString("course_name")
	course_type := this.GetString("course_type")
	course_sub_type := this.GetString("course_sub_type")
	course_short_desc := this.GetString("course_short_desc")
	course.CourseName = course_name
	course.CourseType = course_type
	course.CourseSubType = course_sub_type
	course.CourseShortDes = course_short_desc
	course.CourseStatus = "更新中"
	course.CourseAuthor = user_name
	id, err := ilearning.AddNewCourse(&course)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status":"SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status":"ERROR"}
	}
	topic_theme := ilearning.TopicTheme{}
	topic_theme.TopicId = int(id)
	topic_theme.TopicType = "course_topic_type"
	topic_theme.TopicContent = strings.Join([]string{user_name, "@", course_name,
		"课程更新啦，喜欢该课程的小伙伴们不要错过奥，简洁、直观、免费的课程，能让你更快的掌握知识@", course_short_desc}, "")
	topic_theme.CreatedBy = user_name
	topic_theme.CreatedTime = time.Now()
	topic_theme.LastUpdatedBy = user_name
	topic_theme.LastUpdatedTime = time.Now()
	// 增加一条评论主题
	ilearning.AddTopicTheme(&topic_theme)
	this.ServeJSON()
}
