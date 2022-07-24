package consts

const (
	JwtClaimDataKey = "JwtClaimDataKey"
	SuperAdmin      = 1
)

type JwtClaimData struct {
	UserId       uint64
	Username     string
	IsSuperAdmin bool
}
