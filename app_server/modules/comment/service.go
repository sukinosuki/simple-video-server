package comment

import (
	"errors"
	"simple-video-server/constants/media_type"
	"simple-video-server/constants/video_status"
	"simple-video-server/core"
	"simple-video-server/models"
	"simple-video-server/pkg/business_code"
	"simple-video-server/pkg/global"
	"strconv"
)

type _service struct {
	dao *_dao
}

var Service = &_service{
	dao: Dao,
}

func (s *_service) Add(c *core.Context, commentAdd *CommentAdd) (uint, error) {
	//校验media type、media id
	uid := *c.UID
	mediaType := media_type.Video
	_mediaID, err := strconv.Atoi(c.Param("media_id"))
	if err != nil {
		return 0, err
	}
	mediaID := uint(_mediaID)

	switch c.Param("media") {
	case "video":
		mediaType = media_type.Video
		exists, video, err := s.dao.IsVideoExists(mediaType.Code, mediaID)
		if err != nil {
			return 0, err
		}
		if !exists {
			return 0, business_code.RecodeNotFound
		}
		//	TODO:校验video的是否合法
		if video.Locked || video.Status != video_status.AuditPermit {
			return 0, errors.New("video已被锁定或审核状态不合法")
		}

	//	TODO
	case "post":
		mediaType = media_type.Post
	//	todo
	default:
		panic(errors.New("不支持的media type"))
	}

	comment := &models.Comment{
		Root:      commentAdd.Root,
		Content:   commentAdd.Content,
		MediaType: mediaType.Code,
		MediaID:   mediaID,
		ReplyID:   commentAdd.ReplyID,
		AtUID:     commentAdd.AtUID,
		UID:       uid,
	}

	err = s.dao.Create(comment)

	return comment.ID, err
}

func (s *_service) Delete(c *core.Context) error {
	uid := *c.UID
	id := c.GetId()
	mediaType := media_type.Video
	_mediaID, err := strconv.Atoi(c.Param("media_id"))
	if err != nil {
		return err
	}
	mediaID := uint(_mediaID)

	switch c.Param("media") {
	case "video":
		mediaType = media_type.Video
	//	TODO
	case "post":
		mediaType = media_type.Post
	//	todo
	default:
		panic(errors.New("不支持的media type"))
	}

	err = s.dao.Delete(uid, mediaType.Code, mediaID, id)
	return err
}

func (s *_service) Get(c *core.Context, query *CommentQuery) ([]*CommentResSimple, error) {
	db := global.MysqlDB

	var comments []*CommentResSimple

	//所有的top n二级评论
	var secondComments []*CommentResSimple

	// 3.2 获取每个一级评论的所有回复数
	// SELECT c.root, count(c.root) reply_count FROM `comment` AS c WHERE c.media_id = 4 GROUP BY c.root
	subQuery := db.
		Table("comment").
		Select("root, count(root) reply_count").
		Where("media_id = ?", 4).
		Group("root")

	subQuery3 := db.Table("comment as c").
		Select("c.*, COALESCE(cc.reply_count, 0) AS reply_count, 0 AS row_num").
		Joins("RIGHT JOIN (?) cc ON c.id = cc.root", subQuery).
		Where("c.media_id = ? AND c.root IS NULL", 4).
		Order("reply_count DESC, id DESC"). //TODO: 按最热、最新来排序
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize())

	subQuery4 := db.
		Table("(?) as c1", subQuery3).
		Select("c1.*, user.id user_id, user.nickname user_nickname, user.cover user_cover").
		Joins("LEFT JOIN user on user.id = c1.uid")

	err := subQuery4.Find(&comments).Error
	if err != nil {
		panic(err)
	}

	subQuery5 := db.Table("comment").
		Select("comment.*, user.id user_id, user.cover user_cover, user.nickname user_nickname, 0 AS reply_count, ROW_NUMBER() OVER (PARTITION BY root ORDER BY created_at DESC) as row_num").
		Joins("LEFT JOIN user on user.id = comment.uid").
		Where("media_id = ? AND root IS NOT NULL", 4)

	err = db.Table("(?) as C0", subQuery5).
		Where("c0.row_num <= ?", 2).
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
