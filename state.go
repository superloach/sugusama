package sugusama

type State struct {
	Login      *LoginResp
	SharedData *SharedData

	CSRF     string
	AppID    string
	WWWClaim string
}
