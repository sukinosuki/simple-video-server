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

func (s *_service) Get(c *core.Context, query *CommentQuery) ([]CommentResSimple, error) {
	db := global.MysqlDB
	//var comment []CommentResSimple
	//commentReplySummaryTable := db.Table("(?) as comment_reply_summary",
	//	//db.Model(&models.Comment{}).
	//	db.Select("root, count(root) reply_count").
	//		Where("comment.media_id = ? AND comment.root IS NOT NULL GROUP BY comment.root", 4),
	//)
	//err := db.Model(&models.Comment{}).
	//	Select("comment.id, comment.root, comment.at_uid, comment.content, comment.media_id, comment.media_type, comment.uid, comment.created_at, comment.like, comment.dislike, comment_reply_summary.reply_count").
	//	Joins("left join (?) on comment.id = comment_reply_summary.root", commentReplySummaryTable).
	//	Find(&comment).Error

	var comments []CommentResSimple

	//err := db.
	//	Table("comment").
	//	Select("comment.*, comment_reply_summary_table.reply_count").
	//	Joins("INNER JOIN (?) AS comment_reply_summary_table ON comment.id = comment_reply_summary_table.root",
	//		db.Table("comment").
	//			Select("root, count(root) AS reply_count").
	//			Where("media_id = ? AND root IS NOT NULL", 4).
	//			//Group("root").SubQuery()).
	//			Group("root")).
	//	Find(&comments).Error

	//2
	//subQuery1 := db.Table("comment").
	//	Select("*, 0 AS reply_count, ROW_NUMBER() OVER (PARTITION BY root ORDER BY created_at DESC) as row_num").
	//	Where("media_id = ? AND root IS NOT NULL", 4)
	//
	//subQuery2 := db.Table("comment").
	//	Select("id").
	//	Where("media_id = ?", 4)
	//
	//subQuery3 := db.Table("comment").
	//	Select("root, count(root) reply_count").
	//	Joins("LEFT JOIN (?) p ON p.id = comment.media_id", subQuery2).
	//	Where("media_id = ? AND root IS NOT NULL", 4).
	//	Group("root")
	//
	//err := db.Raw("(?) UNION ALL (?)",
	//	db.Table("(?) as c0", subQuery1).Where("c0.row_num <= ?", 2),
	//	db.Table("`comment` as c").
	//		Select("c.*, COALESCE(cc.reply_count, 0) AS reply_count, ? AS row_num", subQuery3).
	//		Joins("LEFT JOIN (?) cc ON c.id = cc.root", subQuery3).
	//		Where("c.media_id = ? AND c.root IS NULL", subQuery3).
	//		Order("reply_count DESC, id DESC").
	//		Limit(10)).
	//	Find(&comments).Error

	//3
	//3.1 所有一级评论top n的二级评论
	//SELECT c0.* FROM (
	//	SELECT *, 0 AS reply_count, ROW_NUMBER() OVER (PARTITION BY root ORDER BY created_at DESC) as row_num
	//FROM `comment`
	//WHERE media_id = 4
	//AND root IS NOT NULL
	//) AS c0
	//WHERE c0.row_num <= 2
	subQuery1 := db.Table("comment").
		Select("*, 0 AS reply_count, ROW_NUMBER() OVER (PARTITION BY root ORDER BY created_at DESC) as row_num").
		Where("media_id = ? AND root IS NOT NULL", 4)

	subQuery2 := db.Table("(?) as C0", subQuery1).Where("c0.row_num <= ?", 2)

	// 3.2 获取每个一级评论的所有回复数
	// SELECT c.root, count(c.root) reply_count FROM `comment` AS c WHERE c.media_id = 4 GROUP BY c.root
	subQuery := db.
		Table("comment").
		Select("root, count(root) reply_count").
		Where("media_id = ?", 4).
		Group("root")

	subQuery3 := db.Table("comment as c").
		Select("c.*, COALESCE(cc.reply_count, 0) AS reply_count, 0 AS row_num, user.id user_id, user.nickname user_nickname, user.cover user_cover").
		Joins("RIGHT JOIN (?) cc ON c.id = cc.root", subQuery).
		Joins("LEFT JOIN user on user.id = c.uid").
		Where("c.media_id = ? AND c.root IS NULL", 4).
		Order("reply_count DESC, id DESC"). //TODO: 按最热、最新来排序
		Offset(query.GetSafeOffset()).
		Limit(query.GetSafeSize())
	//Find(&comments).Error

	err := db.Raw("(?) UNION ALL (?)", subQuery2, subQuery3).Find(&comments).Error

	if err != nil {
		return nil, err
	}
	var replyMap = make(map[uint][]CommentResSimple)
	var commentList []CommentResSimple
	for _, v := range comments {
		// 一级评论
		if v.Root == nil {
			commentList = append(commentList, v)
		} else {
			arr, ok := replyMap[*v.Root]
			//replyMap[*v.Root] = append(arr, &v)
			if ok {
				replyMap[*v.Root] = append(arr, v)
			} else {
				replyMap[*v.Root] = []CommentResSimple{v}
			}
		}
	}
	for _, v := range commentList {
		arr, ok := replyMap[v.ID]
		if ok {

			v.Replies = arr
		}
	}
	//return comments, err
	return commentList, err
}
