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
	RelationUnfollow  Relation = iota // 未关注/取消关注
	RelationFollow                    // 关注/已关注
	RelationFollowing                 // 关注列表
	RelationFollower                  // 粉丝列表
)
