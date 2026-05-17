package redis

const (
	KEY_ALL_USER      = "user:all"
	KEY_USER_BY_ID    = "user:id:%s"
	KEY_USER_BY_EMAIL = "user:email:%s"

	KEY_ALL_POST            = "post:all"
	KEY_ALL_POST_WITH_USERS = "post:all:users"
	KEY_POST_BY_ID          = "post:id:%s"
	KEY_POST_BY_USER        = "post:user:%s"
)
