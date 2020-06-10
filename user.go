package sugusama

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	FullName   string `json:"full_name"`
	ProfilePic string `json:"profile_pic_url"`

	Blocked    bool `json:"blocked_by_viewer"`
	Followed   bool `json:"followed_by_viewer"`
	HasBlocked bool `json:"has_blocked_viewer"`
	Private    bool `json:"is_private"`
	Verified   bool `json:"is_verified"`
	Requested  bool `json:"requested_by_viewer"`
	Restricted bool `json:"restricted_by_viewer"`
}
