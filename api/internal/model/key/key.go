package key

import (
	"time"
)

// CacheNotFound 默认缓存占位符
const CacheNotFound = "null"

// GameInfoPrefix 游戏缓存前缀
const GameInfoPrefix = "game_info_"

// GameInfoExpress 游戏缓存过期时间
const GameInfoExpress = time.Second * 10 * 60

// UserTokenPrefix token前缀
const UserTokenPrefix = "user_token_"

// LoginTokenPrefix  token前缀
const LoginTokenPrefix = "login_token_"

// UserTokenExpress token过期时间
const UserTokenExpress = time.Hour * 24 * 3

// CodePrefix 验证码前缀
const CodePrefix = "code_"

// CodeExpress 验证码过期时间
const CodeExpress = time.Minute * 10

// AppStoreServerAPIToken 验证码前缀
const AppStoreServerAPIToken = "AppStoreServerAPIToken"

const AppStoreServerAPITokenExpress = time.Hour
