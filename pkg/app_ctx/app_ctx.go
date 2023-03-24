package app_ctx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simple-video-server/models"
)

const UidKey = "uid"

const AuthErrKey = "auth_err"

const HeaderAuthorizeKey = "Authorization"

const TraceIdKey = "trace-id"

func GetHeaderAuthorize(c *gin.Context) string {
	return c.GetHeader(HeaderAuthorizeKey)
}

func SetTraceId(c *gin.Context, traceId string) {
	set[string](c, TraceIdKey, traceId)
}

func GetTraceId(c *gin.Context) (*string, bool) {
	value, ok := get[string](c, TraceIdKey)
	if ok {
		return value, ok
	}

	return nil, false
}

func GetUid(c *gin.Context) (*uint, bool) {
	value, ok := get[uint](c, UidKey)
	if ok {
		return value, ok
	}

	return nil, ok
}

func GetAuth(c *gin.Context) (*models.User, bool) {
	uid, ok := GetUid(c)
	if ok {

		key := fmt.Sprintf("user:%d:info", *uid)
		user, ok := get[*models.User](c, key)

		return *user, ok
	}

	return nil, false
}

func SetAuth(c *gin.Context, user *models.User) {
	key := fmt.Sprintf("user:%d:info", user.ID)
	set(c, key, user)
}

func SetUid(c *gin.Context, uid uint) {
	set(c, UidKey, uid)
}

func GetAuthorizeErr(c *gin.Context) (error, bool) {
	value, ok := get[error](c, AuthErrKey)
	if ok {
		return *value, ok
	}

	return nil, ok
}

func SetAuthorizeErr(c *gin.Context, err error) {
	set(c, AuthErrKey, err)
}

func get[T any](c *gin.Context, key string) (*T, bool) {
	value, exists := c.Get(key)
	if exists {
		uid, ok := value.(T)
		if ok {
			return &uid, true
		}
	}

	return nil, false
}

func set[T any](c *gin.Context, key string, value T) {
	c.Set(key, value)
}
