package upload

type UploadData struct {
	Type int `json:"type" form:"type" binding:"required"`

	Class int `json:"class" form:"class" binding:"required"`
}
