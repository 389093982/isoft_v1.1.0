package ilearning

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/models/ilearning"
	"time"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) FilterTopicReply() {
	// 获取 topic_id 和 topic_type
	topic_id, _ := this.GetInt("topic_id")
	topic_type := this.GetString("topic_type")
	// 获取父评论 id
	parent_id, _ := this.GetInt("parent_id")
	topic_replys, err := ilearning.FilterTopicReply(topic_id, topic_type, parent_id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "topic_replys": topic_replys}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": err.Error()}
	}
	this.ServeJSON()
}

func getIncrementDepth(parent_id int) int {
	if parent_id == 0{
		return 1
	}
	topic_reply, err := ilearning.QueryTopicReplyById(parent_id)
	if err != nil{
		return 1
	}
	return topic_reply.Depth + 1
}

func (this *CommentController) AddTopicReply() {
	user_name := this.Ctx.Input.Session("UserName").(string)
	// 获取 topic_id 和 topic_type
	topic_id, _ := this.GetInt("topic_id")
	topic_type := this.GetString("topic_type")
	// 查询 topicTheme
	topicTheme, _ := ilearning.FilterTopicTheme(topic_id, topic_type)
	// 获取父评论 id
	parent_id, _ := this.GetInt("parent_id", 0)
	// 获取评论内容
	reply_content := this.GetString("reply_content")
	// 获取被评论人员
	refer_user_name := this.GetString("refer_user_name")
	// 构造 TopicReply 实例
	var topic_reply ilearning.TopicReply
	topic_reply.ParentId = parent_id
	topic_reply.ReplyType = "comment"
	topic_reply.ReplyContent = reply_content
	topic_reply.TopicTheme = &topicTheme
	topic_reply.ReferUserName = refer_user_name
	topic_reply.SubReplyAmount = 0
	topic_reply.CreatedBy = user_name
	topic_reply.CreatedTime = time.Now()
	topic_reply.LastUpdatedBy = user_name
	topic_reply.LastUpdatedTime = time.Now()
	// 深度 + 1
	topic_reply.Depth = getIncrementDepth(parent_id)
	_, err := ilearning.AddTopicReply(&topic_reply)
	ilearning.ModifySubReplyAmount(parent_id)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}

func (this *CommentController) FilterTopicTheme() {
	// 获取课程 id
	topic_id, _ := this.GetInt("topic_id")
	topic_type := this.GetString("topic_type")
	topic_theme, err := ilearning.FilterTopicTheme(topic_id, topic_type)
	if err == nil {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "topic_theme": topic_theme}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	}
	this.ServeJSON()
}