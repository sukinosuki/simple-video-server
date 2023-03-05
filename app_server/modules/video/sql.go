package video

const (
	updateVideo = `
		UPDATE 
		    video 
		set title = ?, cover = ?, updated_at = ? 
		WHERE id = ?
	`
)
