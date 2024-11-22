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

const UserTokenPrefix = "user_token_"

const UserTokenExpress = time.Hour * 24 * 3
