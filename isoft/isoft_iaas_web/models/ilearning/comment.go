package ilearning

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CommentTheme struct {
	Id              int       `json:"id"`
	CommentId       int       `json:"comment_id"`                       // 评论主题 id
	CommentType     string    `json:"comment_type"`                     // 评论主题类型
	CommentContent  string    `json:"comment_content" orm:"size(4000)"` // 评论主题内容
	CreatedBy       string    `json:"created_by"`                       // 评论主题创建人
	CreatedTime     time.Time `json:"created_time"`                     // 评论主题创建时间
	LastUpdatedBy   string    `json:"last_updated_by"`                  // 评论主题修改人
	LastUpdatedTime time.Time `json:"last_updated_time"`                // 评论主题修改时间
}

type CommentReply struct {
	Id              int           `json:"id"`
	ParentId        int           `json:"parent_id" orm:"default(0)` // 父级评论回复 id
	CommentTheme    *CommentTheme `orm:"rel(fk)" json:"comment_theme"`
	Depth           int           `json:"depth"`                          // 当前评论深度
	ReplyType       string        `json:"reply_type"`                     // 评论类型
	ReplyContent    string        `json:"reply_content" orm:"size(4000)"` // 评论内容
	ReferUserName   string        `json:"refer_user_name"`                // 被评论人
	SubReplyAmount  int           `json:"sub_reply_amount"`               // 子评论数
	CreatedBy       string        `json:"created_by"`                     // 评论回复创建人
	CreatedTime     time.Time     `json:"created_time"`                   // 评论回复创建时间
	LastUpdatedBy   string        `json:"last_updated_by"`                // 评论回复修改人
	LastUpdatedTime time.Time     `json:"last_updated_time"`              // 评论回复修改时间
}

func AddCommentTheme(comment_theme *CommentTheme) (id int64, err error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("comment_theme").Filter("comment_id", comment_theme.CommentId).Filter("comment_type", comment_theme.CommentType).Count()
	if count == 0 {
		id, err = o.Insert(comment_theme)
	}
	return
}

func QueryCommentReplyById(id int) (comment_reply CommentReply, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("comment_reply").Filter("id", id).One(&comment_reply)
	return
}

func FilterCommentTheme(comment_id int, comment_type string) (comment_theme CommentTheme, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("comment_theme").Filter("comment_id", comment_id).Filter("comment_type", comment_type).One(&comment_theme)
	return
}

func AddCommentReply(comment_reply *CommentReply) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(comment_reply)
	return
}

func FilterCommentReply(comment_id int, comment_type string, parent_id int) (comment_replys []CommentReply, err error) {
	o := orm.NewOrm()
	commentTheme, _ := FilterCommentTheme(comment_id, comment_type)
	_, err = o.QueryTable("comment_reply").Filter("comment_theme_id", commentTheme.Id).Filter("parent_id", parent_id).
		OrderBy("-created_time").All(&comment_replys)
	return
}

func ModifySubReplyAmount(id int) {
	o := orm.NewOrm()
	o.QueryTable("comment_reply").Filter("id", id).Update(orm.Params{
		"sub_reply_amount": orm.ColValue(orm.ColAdd, 1),
	})
}