package sugusama

import (
	"encoding/json"
	"errors"
	"regexp"
)

var SharedDataRegexp = regexp.MustCompile(`window\._sharedData = ([^;]+);`)

var (
	ErrSharedDataNotFound = errors.New("shared data not found")
)

type SharedData struct {
	BundleVariant         string                 `json:"bundle_variant"`
	CacheSchemaVersion    int64                  `json:"cache_schema_version"`
	Config                SharedDataConfig       `json:"config"`
	CountryCode           string                 `json:"country_code"`
	DeploymentStage       string                 `json:"deployment_stage"`
	DeviceID              string                 `json:"device_id"`
	Encryption            SharedDataEncryption   `json:"encryption"`
	EntryData             SharedDataEntryData    `json:"entry_data"`
	Hostname              string                 `json:"hostname"`
	IsCanary              bool                   `json:"is_canary"`
	IsDev                 bool                   `json:"is_dev"`
	IsWhitelistedCrawlBot bool                   `json:"is_whitelisted_crawl_bot"`
	Knobx                 map[string]bool        `json:"knobx"`
	LanguageCode          string                 `json:"language_code"`
	Locale                string                 `json:"locale"`
	MidPct                float64                `json:"mid_pct"`
	Nonce                 string                 `json:"nonce"`
	Platform              string                 `json:"platform"`
	RolloutHash           string                 `json:"rollout_hash"`
	ServerChecks          SharedDataServerChecks `json:"server_checks"`
	ToCache               SharedDataToCache      `json:"to_cache"`
	ZeroData              map[string]interface{} `json:"zero_data"`
}

func (c *Client) ExtractSharedData(body []byte) (*SharedData, error) {
	data := SharedDataRegexp.Find(body)
	if data == nil {
		return nil, ErrSharedDataNotFound
	}
	data = data[21 : len(data)-1]

	sd := &SharedData{}
	err := json.Unmarshal(data, &sd)
	if err != nil {
		return nil, err
	}

	c.State.SharedData = sd
	return sd, nil
}

type SharedDataViewer struct {
	badgeCount       string `json:"badge_count"`
	Biography        string `json:"biography"`
	ExternalURL      string `json:"external_url"`
	FullName         string `json:"full_name"`
	HasPhoneNumber   bool   `json:"has_phone_number"`
	HasProfilePic    bool   `json:"has_profile_pic"`
	HasTabbedInbox   bool   `json:"has_tabbed_inbox"`
	ID               string `json:"id"`
	IsJoinedRecently bool   `json:"is_joined_recently"`
	IsPrivate        bool   `json:"is_private"`
	ProfilePicURL    string `json:"profile_pic_url"`
	ProfilePicURLHD  string `json:"profile_pic_url_hd"`
	Username         string `json:"username"`
}

func (s SharedDataViewer) BadgeCount() (*DirectBadgeCountResp, error) {
	bc := &DirectBadgeCountResp{}
	return bc, json.Unmarshal([]byte(s.badgeCount), &bc)
}

type SharedDataToCache struct {
	CB             bool            `json:"cb"`
	Gatekeepers    map[string]bool `json:"gatekeepers"`
	ProbablyHasApp bool            `json:"probably_has_app"`
	QE             SharedDataQE    `json:"qe"`
}

type SharedDataConfig struct {
	CSRF     string           `json:"csrf_token"`
	Viewer   SharedDataViewer `json:"viewer"`
	ViewerID string           `json:"viewerId"`
}

type SharedDataEntryData struct {
	FeedPage []map[string]interface{} `json:"FeedPage"`
}

type SharedDataGP struct {
	G string            `json:"g"`
	P map[string]string `json:"p"`
}

type SharedDataServerChecks struct {
	HFE bool `json:"hfe"`
}

type SharedDataEncryption struct {
	KeyID     string `json:"key_id"`
	PublicKey string `json:"public_key"`
	Version   string `json:"version"`
}

type SharedDataLPQEX struct {
	L   map[string]interface{} `json:"l"`
	P   map[string]interface{} `json:"p"`
	QEX bool                   `json:"qex"`
}

type SharedDataQE struct {
	AppUpsell                     SharedDataGP `json:"app_upsell"`
	FelixClearFbCookie            SharedDataGP `json:"felix_clear_fb_cookie"`
	FelixCreationDurationLimits   SharedDataGP `json:"felix_creation_duration_limits"`
	FelixCreationFbCrossposting   SharedDataGP `json:"felix_creation_fb_crossposting"`
	FelixCreationFbCrosspostingV2 SharedDataGP `json:"felix_creation_fb_crossposting_v2"`
	FelixCreationValidation       SharedDataGP `json:"felix_creation_validation"`
	IglAppUpsell                  SharedDataGP `json:"igl_app_upsell"`
	Notif                         SharedDataGP `json:"notif"`
	Onetaplogin                   SharedDataGP `json:"onetaplogin"`
	PostOptions                   SharedDataGP `json:"post_options"`
	StickerTray                   SharedDataGP `json:"sticker_tray"`
	WebSentry                     SharedDataGP `json:"web_sentry"`

	// these are bs, leave them out for now
	/*Zero   SharedDataLPQEX `json:"0"`
	One00  SharedDataLPQEX `json:"100"`
	One01  SharedDataLPQEX `json:"101"`
	One02  SharedDataLPQEX `json:"102"`
	One03  SharedDataLPQEX `json:"103"`
	One04  SharedDataLPQEX `json:"104"`
	One05  SharedDataLPQEX `json:"105"`
	One08  SharedDataLPQEX `json:"108"`
	One09  SharedDataLPQEX `json:"109"`
	One10  SharedDataLPQEX `json:"110"`
	One11  SharedDataLPQEX `json:"111"`
	One12  SharedDataLPQEX `json:"112"`
	One13  SharedDataLPQEX `json:"113"`
	One14  SharedDataLPQEX `json:"114"`
	One15  SharedDataLPQEX `json:"115"`
	One2   SharedDataLPQEX `json:"12"`
	One3   SharedDataLPQEX `json:"13"`
	One6   SharedDataLPQEX `json:"16"`
	One7   SharedDataLPQEX `json:"17"`
	Two    SharedDataLPQEX `json:"2"`
	Two1   SharedDataLPQEX `json:"21"`
	Two2   SharedDataLPQEX `json:"22"`
	Two3   SharedDataLPQEX `json:"23"`
	Two5   SharedDataLPQEX `json:"25"`
	Two6   SharedDataLPQEX `json:"26"`
	Two8   SharedDataLPQEX `json:"28"`
	Two9   SharedDataLPQEX `json:"29"`
	Three1 SharedDataLPQEX `json:"31"`
	Three3 SharedDataLPQEX `json:"33"`
	Three4 SharedDataLPQEX `json:"34"`
	Three6 SharedDataLPQEX `json:"36"`
	Three7 SharedDataLPQEX `json:"37"`
	Three9 SharedDataLPQEX `json:"39"`
	Four1  SharedDataLPQEX `json:"41"`
	Four2  SharedDataLPQEX `json:"42"`
	Four3  SharedDataLPQEX `json:"43"`
	Four4  SharedDataLPQEX `json:"44"`
	Four5  SharedDataLPQEX `json:"45"`
	Four6  SharedDataLPQEX `json:"46"`
	Four7  SharedDataLPQEX `json:"47"`
	Four9  SharedDataLPQEX `json:"49"`
	Five0  SharedDataLPQEX `json:"50"`
	Five4  SharedDataLPQEX `json:"54"`
	Five5  SharedDataLPQEX `json:"55"`
	Five8  SharedDataLPQEX `json:"58"`
	Five9  SharedDataLPQEX `json:"59"`
	Six2   SharedDataLPQEX `json:"62"`
	Six5   SharedDataLPQEX `json:"65"`
	Six6   SharedDataLPQEX `json:"66"`
	Six7   SharedDataLPQEX `json:"67"`
	Six8   SharedDataLPQEX `json:"68"`
	Six9   SharedDataLPQEX `json:"69"`
	Seven0 SharedDataLPQEX `json:"70"`
	Seven1 SharedDataLPQEX `json:"71"`
	Seven2 SharedDataLPQEX `json:"72"`
	Seven3 SharedDataLPQEX `json:"73"`
	Seven4 SharedDataLPQEX `json:"74"`
	Seven5 SharedDataLPQEX `json:"75"`
	Seven7 SharedDataLPQEX `json:"77"`
	Seven8 SharedDataLPQEX `json:"78"`
	Eight0 SharedDataLPQEX `json:"80"`
	Eight4 SharedDataLPQEX `json:"84"`
	Eight5 SharedDataLPQEX `json:"85"`
	Eight7 SharedDataLPQEX `json:"87"`
	Eight9 SharedDataLPQEX `json:"89"`
	Nine2  SharedDataLPQEX `json:"92"`
	Nine3  SharedDataLPQEX `json:"93"`
	Nine5  SharedDataLPQEX `json:"95"`
	Nine6  SharedDataLPQEX `json:"96"`
	Nine7  SharedDataLPQEX `json:"97"`
	Nine8  SharedDataLPQEX `json:"98"`
	Nine9  SharedDataLPQEX `json:"99"`*/
}
