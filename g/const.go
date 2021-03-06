package g

const (
	DOMAIN     = "https://emsystem.com"
	DOMAIN_API = "https://api.emsystem.com"
	VERSION    = "0.0.1"

	AUTH_TOKEN_NS = "auth:" //认证token存取前缀

	RUN_MODE_DEV     = "dev"
	RUN_MODE_PRODUCT = "prod"

	LOG_ENGINE_FILE    = "file"
	LOG_ENGINE_MONGODB = "mongodb"

	THUMB_FOLDER_NUM = 1024
	THUMB_FORMAT     = ".png"

	LOCAL_TIME_ZONE_OFFSET = 8 * 60 * 60 //Beijing(UTC+8:00)

	DEFAULT_FILE_MODE = 0755
)

//todo log type
const (
	LOG_TYPE_DEFAULT = iota
	LOG_TYPE_REGISTER
	LOG_TYPE_ACTIVATION
	LOG_TYPE_LOGIN
)
