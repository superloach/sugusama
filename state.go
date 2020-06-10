package sugusama

type State struct {
	Viewer   *User
	Activity *Activity
	Feed     []*Post

	CSRF     string
	AppID    string
	WWWClaim string
}

func (s *State) processAdditional(ad additionalData) {
	if ad.User.ID != "" {
		s.Viewer.ID = ad.User.ID
	}
	if ad.User.Username != "" {
		s.Viewer.Username = ad.User.Username
	}
	if ad.User.FullName != "" {
		s.Viewer.FullName = ad.User.FullName
	}
	if ad.User.ProfilePic != "" {
		s.Viewer.ProfilePic = ad.User.ProfilePic
	}

	s.Viewer.Blocked = ad.User.Blocked
	s.Viewer.Followed = ad.User.Followed
	s.Viewer.HasBlocked = ad.User.HasBlocked
	s.Viewer.Private = ad.User.Private
	s.Viewer.Verified = ad.User.Verified
	s.Viewer.Requested = ad.User.Requested
	s.Viewer.Restricted = ad.User.Restricted

	println("TODO process feed")
}

func (s *State) processShared(sd sharedData) {
	println("TODO process shared")
}

func (s *State) processActivity(ar activityResp) {
	println("TODO process activity")
}
