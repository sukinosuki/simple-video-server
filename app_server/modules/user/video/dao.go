package video

import (
	"github.com/jmoiron/sqlx"
	"simple-video-server/db"
	"simple-video-server/models"
)

type _dao struct {
	sdb *sqlx.DB
}

var dao = &_dao{
	sdb: db.GetSqlxDB(),
}

func (d *_dao) GetAll(uid uint, query *VideoQuery) ([]models.Video, error) {
	sql := `
		SELECT
			* 
		FROM
			video
		WHERE
			uid = ?
		LIMIT ?, ?
	`

	var videos []models.Video
	err := db.SqlSelectLog(d.sdb.Select, &videos, sql, uid, (query.Page-1)*query.Size, query.Size)

	return videos, err
}
