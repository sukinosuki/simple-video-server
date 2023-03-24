package comment

import (
	"errors"
	"simple-video-server/constants/media_type"
	"simple-video-server/constants/video_status"
	"simple-video-server/core"
	"simple-video-server/models"
	"simple-video-server/pkg/business_code"
	"simple-video-server/pkg/global"
)

type Service struct {
	dao *_dao
}

var _Service = &Service{
	dao: Dao,
}

func (s *Service) Add(c *core.Context, commentAdd *CommentAdd, mediaID uint, mediaType int) (*CommentResSimple, error) {
	uid := *c.AuthUID

	switch {
	case media_type.Video.Is(mediaType):

		exists, video, err := s.dao.IsVideoExists(mediaType, mediaID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, business_code.RecodeNotFound
		}
		//	TODO:校验video的是否合法
		if video.Locked || video.Status != video_status.AuditPermit {
			return nil, errors.New("video已被锁定或审核状态不合法")
		}

	//	TODO
	case media_type.Post.Is(mediaType):
		//mediaType = media_type.Post
	//	todo
	default:
		panic(errors.New("不支持的media type"))
	}

	comment := &models.Comment{
		Root:      commentAdd.Root,
		Content:   commentAdd.Content,
		MediaType: mediaType,
		MediaID:   mediaID,
		ReplyID:   commentAdd.ReplyID,
		AtUID:     commentAdd.AtUID,
		UID:       uid,
	}

	err := s.dao.Create(comment)

	//user, _ := s.dao.GetUserById(uid)

	commentResSimple := &CommentResSimple{
		ID:         comment.ID,
		Root:       comment.Root,
		AtUID:      comment.AtUID,
		Content:    comment.Content,
		MediaType:  comment.MediaType,
		MediaID:    comment.MediaID,
		UID:        comment.UID,
		CreatedAt:  comment.CreatedAt,
		Like:       0,
		Dislike:    0,
		ReplyCount: 0,
		RowNum:     0,
		Replies:    make([]*CommentResSimple, 0),
		User: &CommentResSimpleUser{
			ID:       comment.UID,
			Avatar:   c.Auth.Avatar,
			Nickname: c.Auth.Nickname,
		},
	}

	return commentResSimple, err
}

func (s *Service) Delete(c *core.Context, mediaID uint, mediaType int) error {
	uid := *c.AuthUID
	id := c.GetParamId()
	//mediaType := media_type.Video
	//_mediaID, err := strconv.Atoi(c.Param("media_id"))
	//if err != nil {
	//	return err
	//}
	//mediaID := uint(_mediaID)

	switch {
	case media_type.Video.Is(mediaType):
		//mediaType = media_type.Video
	//	TODO
	case media_type.Post.Is(mediaType):

		//mediaType = media_type.Post
	//	todo
	default:
		panic(errors.New("不支持的media type"))
	}

	err := s.dao.Delete(uid, mediaType, mediaID, id)
	return err
}

func (s *Service) GetAll(c *core.Context, query *CommentQuery) ([]*CommentResSimple, error) {
	db := global.MysqlDB

	var comments []*CommentResSimple

	//所有的top n二级评论
	var secondComments []*CommentResSimple

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
		//Order("id DESC"). //TODO: 按最热、最新来排序
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize())

	subQuery4 := db.
		Table("(?) as c1", subQuery3).
		Select("c1.*, user.id user_id, user.nickname user_nickname, user.avatar user_avatar").
		Joins("LEFT JOIN user on user.id = c1.uid")

	err := subQuery4.Find(&comments).Error
	if err != nil {
		panic(err)
	}

	subQuery5 := db.Table("comment").
		Select("comment.*, user.id user_id, user.avatar user_avatar, user.nickname user_nickname, 0 AS reply_count, ROW_NUMBER() OVER (PARTITION BY root ORDER BY created_at DESC) as row_num").
		Joins("LEFT JOIN user on user.id = comment.uid").
		Where("media_id = ? AND root IS NOT NULL", query.MediaID)

	err = db.Table("(?) as C0", subQuery5).
		Where("c0.row_num <= ?", 2). // top 2
		Find(&secondComments).Error

	if err != nil {
		panic(err)
	}

	var replyMap = make(map[uint][]*CommentResSimple)

	for _, v := range secondComments {
		arr, ok := replyMap[*v.Root]
		if ok {
			replyMap[*v.Root] = append(arr, v)
		} else {
			replyMap[*v.Root] = []*CommentResSimple{v}
		}
	}

	for _, v := range comments {
		arr, ok := replyMap[v.ID]
		if ok {
			v.Replies = arr
		}
	}

	return comments, err
}

func (s *Service) Get(c *core.Context, query *CommentQuery) ([]CommentResSimple, error) {
	id := c.GetParamId()
	db := global.MysqlDB

	var comment []CommentResSimple
	err := db.Model(&models.Comment{}).
		Select("comment.*, user.id user_id, user.nickname user_nickname, user.avatar user_avatar").
		Joins("LEFT JOIN user ON user.id = comment.uid").
		Where("root = ?", id).
		Order("comment.created_at DESC").
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize()).
		Find(&comment).Error

	return comment, err
}
