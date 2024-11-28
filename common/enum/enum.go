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
	RelationFollow    Relation = 1
	RelationUnfollow  Relation = 2
	RelationFollowing Relation = 3
	RelationFollower  Relation = 4
)
