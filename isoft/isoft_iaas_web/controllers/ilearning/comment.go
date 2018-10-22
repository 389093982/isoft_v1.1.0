package ilearning

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/models/ilearning"
	"time"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) FilterCommentReply() {
	// 获取 comment_id 和 theme_type
	comment_id, _ := this.GetInt("comment_id")
	theme_type := this.GetString("theme_type")
	reply_comment_type := this.GetString("reply_comment_type")
	// 获取父评论 id
	parent_id, _ := this.GetInt("parent_id")
	comment_replys, err := ilearning.FilterCommentReply(comment_id, theme_type, parent_id, reply_comment_type)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "comment_replys": comment_replys}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": err.Error()}
	}
	this.ServeJSON()
}

func getIncrementDepth(parent_id int) int {
	if parent_id == 0 {
		return 1
	}
	comment_reply, err := ilearning.QueryCommentReplyById(parent_id)
	if err != nil {
		return 1
	}
	return comment_reply.Depth + 1
}

func (this *CommentController) AddCommentReply() {
	user_name := this.Ctx.Input.Session("UserName").(string)
	// 获取 comment_id 和 theme_type
	comment_id, _ := this.GetInt("comment_id")
	theme_type := this.GetString("theme_type")
	reply_comment_type := this.GetString("reply_comment_type")
	// 查询 CommentTheme
	CommentTheme, _ := ilearning.FilterCommentTheme(comment_id, theme_type)
	// 获取父评论 id
	parent_id, _ := this.GetInt("parent_id", 0)
	// 获取评论内容
	reply_content := this.GetString("reply_content")
	// 获取被评论人员
	refer_user_name := this.GetString("refer_user_name")
	// 构造 CommentReply 实例
	var comment_reply ilearning.CommentReply
	comment_reply.ParentId = parent_id
	comment_reply.ReplyThemeType = CommentTheme.ThemeType
	comment_reply.ReplyContent = reply_content
	comment_reply.CommentTheme = &CommentTheme
	comment_reply.ReplyCommentType = reply_comment_type
	comment_reply.ReferUserName = refer_user_name
	comment_reply.SubReplyAmount = 0
	comment_reply.CreatedBy = user_name
	comment_reply.CreatedTime = time.Now()
	comment_reply.LastUpdatedBy = user_name
	comment_reply.LastUpdatedTime = time.Now()
	// 深度 + 1
	comment_reply.Depth = getIncrementDepth(parent_id)
	_, err := ilearning.AddCommentReply(&comment_reply)
	ilearning.ModifySubReplyAmount(parent_id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *CommentController) FilterCommentTheme() {
	// 获取课程 id
	comment_id, _ := this.GetInt("comment_id")
	theme_type := this.GetString("theme_type")
	comment_theme, err := ilearning.FilterCommentTheme(comment_id, theme_type)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "comment_theme": comment_theme}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}
