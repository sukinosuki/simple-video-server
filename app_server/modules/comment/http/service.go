package http

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"simple-video-server/app_server/modules/comment"
	"simple-video-server/constants/media_type"
	"simple-video-server/constants/video_status"
	"simple-video-server/core"
	"simple-video-server/db"
	"simple-video-server/models"
	"simple-video-server/pkg/business_code"
	"simple-video-server/pkg/global"
)

type Service struct {
	dao *comment.Dao
	db  *gorm.DB
}

var _Service = &Service{
	dao: comment.GetDao(),
	db:  db.GetOrmDB(),
}

// Add 新增
func (s *Service) Add(c *core.Context, commentAdd *comment.CommentAdd, mediaID uint, mediaType int) *comment.CommentResSimple {
	handlerName := "Add"
	uid := *c.AuthUID

	log := c.Log.With(zap.String("handler", handlerName))

	tx := s.db.Begin()

	defer func() {
		err := recover()
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
			panic(err)
		}
	}()

	switch {
	case media_type.Video.Is(mediaType):

		exists, video, err := s.dao.IsVideoExists(mediaType, mediaID)
		c.PanicIfErr(err, handlerName, "获取媒体失败")

		if !exists {
			c.Panic(business_code.RecodeNotFound, handlerName, "media不存在")
			return nil
		}

		//	TODO: 校验video的是否合法
		if video.Locked || video.Status != video_status.AuditPermit {
			c.Panic(errors.New("video已被锁定或审核状态不合法"), handlerName, "视频不合法不可以评论")
			return nil
		}

	//	TODO
	case media_type.Post.Is(mediaType):
		//mediaType = media_type.Post
	//	todo
	default:
		c.Panic(errors.New("不支持的media type"), handlerName, "不支持的media type参数")
		return nil
	}

	_comment := &models.Comment{
		Root:      commentAdd.Root,
		Content:   commentAdd.Content,
		MediaType: mediaType,
		MediaID:   mediaID,
		ReplyID:   commentAdd.ReplyID,
		AtUID:     commentAdd.AtUID,
		UID:       uid,
	}

	err := s.dao.Create(tx, _comment)
	c.PanicIfErr(err, handlerName, "新增评论失败")

	commentResSimple := &comment.CommentResSimple{
		ID:         _comment.ID,
		Root:       _comment.Root,
		AtUID:      _comment.AtUID,
		Content:    _comment.Content,
		MediaType:  _comment.MediaType,
		MediaID:    _comment.MediaID,
		UID:        _comment.UID,
		CreatedAt:  _comment.CreatedAt,
		Like:       0,
		Dislike:    0,
		ReplyCount: 0,
		RowNum:     0,
		Replies:    make([]*comment.CommentResSimple, 0),
		User: &comment.CommentResSimpleUser{
			ID:       _comment.UID,
			Avatar:   c.Auth.Avatar,
			Nickname: c.Auth.Nickname,
		},
	}

	log.Info("用户新增视频")
	return commentResSimple
}

// Delete 删除
func (s *Service) Delete(c *core.Context, mediaID uint, mediaType int) {
	tx := s.db.Begin()

	defer func() {
		err := recover()
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
			panic(err)
		}
	}()

	uid := *c.AuthUID
	id := c.GetParamId()
	handlerName := "Delete"

	log := c.Log.With(zap.String("handler", handlerName))

	switch {
	case media_type.Video.Is(mediaType):
	//	TODO
	case media_type.Post.Is(mediaType):

	//	todo
	default:
		panic(errors.New("不支持的media type"))
	}

	err := s.dao.DeleteByIdAndUidAndIdAndType(tx, uid, mediaType, mediaID, id)
	c.PanicIfErr(err, handlerName, "删除media失败")

	log.Info("用户删除media成功")
}

// GetAll 获取评论
func (s *Service) GetAll(c *core.Context, query *comment.CommentQuery) []*comment.CommentResSimple {
	db := global.MysqlDB
	handlerName := "GetAll"

	var comments []*comment.CommentResSimple

	//所有的top n二级评论
	var replies []*comment.CommentResSimple

	// 3.2 获取每个一级评论的所有回复数
	// SELECT c.root, count(c.root) reply_count FROM `comment` AS c WHERE c.media_id = 4 GROUP BY c.root
	subQuery := db.
		Table("comment").
		Select("root, count(root) reply_count").
		Where("media_id = ?", query.MediaID).
		Group("root")

	subQuery3 := db.Table("comment as c").
		Select("c.*, COALESCE(cc.reply_count, 0) AS reply_count, 0 AS row_num").
		Joins("LEFT JOIN (?) cc ON c.id = cc.root", subQuery).
		Where("c.media_id = ? AND c.root IS NULL", query.MediaID).
		Order("reply_count DESC, id DESC"). //TODO: 按最热、最新来排序
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize())

	subQuery4 := db.
		Table("(?) as c1", subQuery3).
		Select("c1.*, user.id user_id, user.nickname user_nickname, user.avatar user_avatar").
		Joins("LEFT JOIN user on user.id = c1.uid")

	err := subQuery4.Find(&comments).Error
	c.PanicIfErr(err, handlerName, "获取评论失败")

	subQuery5 := db.Table("comment").
		Select("comment.*, user.id user_id, user.avatar user_avatar, user.nickname user_nickname, 0 AS reply_count, ROW_NUMBER() OVER (PARTITION BY root ORDER BY created_at DESC) as row_num").
		Joins("LEFT JOIN user on user.id = comment.uid").
		Where("media_id = ? AND root IS NOT NULL", query.MediaID)

	err = db.Table("(?) as C0", subQuery5).
		Where("c0.row_num <= ?", 2). // top 2
		Find(&replies).Error

	c.PanicIfErr(err, handlerName, "获取评论失败")

	var replyMap = make(map[uint][]*comment.CommentResSimple)

	for _, v := range replies {
		arr, ok := replyMap[*v.Root]
		if ok {
			replyMap[*v.Root] = append(arr, v)
		} else {
			replyMap[*v.Root] = []*comment.CommentResSimple{v}
		}
	}

	for _, v := range comments {
		arr, ok := replyMap[v.ID]
		if ok {
			v.Replies = arr
		}
	}

	return comments
}

// Get 获取评论回复
func (s *Service) Get(c *core.Context, query *comment.CommentQuery) []comment.CommentResSimple {
	handlerName := "Get"
	id := c.GetParamId()
	db := global.MysqlDB

	var comments []comment.CommentResSimple
	err := db.Model(&models.Comment{}).
		Select("comment.*, user.id user_id, user.nickname user_nickname, user.avatar user_avatar").
		Joins("LEFT JOIN user ON user.id = comment.uid").
		Where("root = ?", id).
		Order("comment.created_at DESC").
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize()).
		Find(&comments).Error

	c.PanicIfErr(err, handlerName, "获取评论回复失败")

	return comments
}
