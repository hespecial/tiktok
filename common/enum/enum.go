package enum

const (
	ConfigFile = "./config/config.yaml"

	JwtTokenName = "Token"
	JwtIssuer    = "tiktok"
	JwtTtl       = 24 * 7 // hour
	JwtSecret    = "tiktok"

	ContextUserId = "user"
)

type Relation uint8

const (
	RelationFollow    Relation = 1 // 已关注
	RelationUnfollow  Relation = 2 // 未关注
	RelationFollowing Relation = 3 // 关注对象
	RelationFollower  Relation = 4 // 粉丝
)
