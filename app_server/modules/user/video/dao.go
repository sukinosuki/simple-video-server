package video

import (
	"github.com/jmoiron/sqlx"
	"simple-video-server/db"
	"simple-video-server/models"
)

type Dao struct {
	sdb *sqlx.DB
}

var _dao = &Dao{
	sdb: db.GetSqlxDB(),
}

func GetDao() *Dao {
	return _dao
}

func (d *Dao) GetAll(uid uint, query *VideoQuery) ([]models.Video, error) {
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
