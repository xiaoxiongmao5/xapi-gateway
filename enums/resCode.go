package enums

type ResCodeCommon int
type ResCodeOfManage int

const (
	// ResCodeCommon
	Success            ResCodeCommon = iota // 0
	ParameterError                          // 1
	AuthenticationFail                      // 2
	Unauthorized
	NotAdminRole
)

const (
	// ResCodeOfManage
	UpdateIPRateLimitConfigFailed ResCodeOfManage = 2000
)
