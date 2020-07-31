package sugusama

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

var (
	sharedDataRegexp = regexp.MustCompile(`window\._sharedData = ([^;]+);`)

	ErrSharedDataNotFound = errors.New("shared data not found")
)

type sharedData struct {
	BundleVariant         string                 `json:"bundle_variant"`
	CacheSchemaVersion    int64                  `json:"cache_schema_version"`
	Config                sharedDataConfig       `json:"config"`
	CountryCode           string                 `json:"country_code"`
	DeploymentStage       string                 `json:"deployment_stage"`
	DeviceID              string                 `json:"device_id"`
	Encryption            sharedDataEncryption   `json:"encryption"`
	EntryData             sharedDataEntryData    `json:"entry_data"`
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
	ServerChecks          sharedDataServerChecks `json:"server_checks"`
	ToCache               sharedDataToCache      `json:"to_cache"`
	ZeroData              map[string]interface{} `json:"zero_data"`
}

func (c *Client) extractShared(body []byte) error {
	data := sharedDataRegexp.Find(body)
	if data == nil {
		return ErrSharedDataNotFound
	}
	data = data[21 : len(data)-1]

	sd := sharedData{}
	err := json.Unmarshal(data, &sd)
	if err != nil {
		err = fmt.Errorf("unmarshal sd: %w", err)
		return err
	}

	c.State.processShared(sd)

	return nil
}

type sharedDataViewer struct {
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

func (s sharedDataViewer) BadgeCount() (*DirectBadgeCountResp, error) {
	bc := &DirectBadgeCountResp{}
	return bc, json.Unmarshal([]byte(s.badgeCount), &bc)
}

type sharedDataToCache struct {
	CB             bool            `json:"cb"`
	Gatekeepers    map[string]bool `json:"gatekeepers"`
	ProbablyHasApp bool            `json:"probably_has_app"`
	QE             sharedDataQE    `json:"qe"`
}

type sharedDataConfig struct {
	CSRF     string           `json:"csrf_token"`
	Viewer   sharedDataViewer `json:"viewer"`
	ViewerID string           `json:"viewerId"`
}

type sharedDataEntryData struct {
	FeedPage []map[string]interface{} `json:"FeedPage"`
}

type sharedDataGP struct {
	G string            `json:"g"`
	P map[string]string `json:"p"`
}

type sharedDataServerChecks struct {
	HFE bool `json:"hfe"`
}

type sharedDataEncryption struct {
	KeyID     string `json:"key_id"`
	PublicKey string `json:"public_key"`
	Version   string `json:"version"`
}

type sharedDataLPQEX struct {
	L   map[string]interface{} `json:"l"`
	P   map[string]interface{} `json:"p"`
	QEX bool                   `json:"qex"`
}

type sharedDataQE struct {
	AppUpsell                     sharedDataGP `json:"app_upsell"`
	FelixClearFbCookie            sharedDataGP `json:"felix_clear_fb_cookie"`
	FelixCreationDurationLimits   sharedDataGP `json:"felix_creation_duration_limits"`
	FelixCreationFbCrossposting   sharedDataGP `json:"felix_creation_fb_crossposting"`
	FelixCreationFbCrosspostingV2 sharedDataGP `json:"felix_creation_fb_crossposting_v2"`
	FelixCreationValidation       sharedDataGP `json:"felix_creation_validation"`
	IglAppUpsell                  sharedDataGP `json:"igl_app_upsell"`
	Notif                         sharedDataGP `json:"notif"`
	Onetaplogin                   sharedDataGP `json:"onetaplogin"`
	PostOptions                   sharedDataGP `json:"post_options"`
	StickerTray                   sharedDataGP `json:"sticker_tray"`
	WebSentry                     sharedDataGP `json:"web_sentry"`

	// these are bs, leave them out for now
	/*Zero   sharedDataLPQEX `json:"0"`
	One00  sharedDataLPQEX `json:"100"`
	One01  sharedDataLPQEX `json:"101"`
	One02  sharedDataLPQEX `json:"102"`
	One03  sharedDataLPQEX `json:"103"`
	One04  sharedDataLPQEX `json:"104"`
	One05  sharedDataLPQEX `json:"105"`
	One08  sharedDataLPQEX `json:"108"`
	One09  sharedDataLPQEX `json:"109"`
	One10  sharedDataLPQEX `json:"110"`
	One11  sharedDataLPQEX `json:"111"`
	One12  sharedDataLPQEX `json:"112"`
	One13  sharedDataLPQEX `json:"113"`
	One14  sharedDataLPQEX `json:"114"`
	One15  sharedDataLPQEX `json:"115"`
	One2   sharedDataLPQEX `json:"12"`
	One3   sharedDataLPQEX `json:"13"`
	One6   sharedDataLPQEX `json:"16"`
	One7   sharedDataLPQEX `json:"17"`
	Two    sharedDataLPQEX `json:"2"`
	Two1   sharedDataLPQEX `json:"21"`
	Two2   sharedDataLPQEX `json:"22"`
	Two3   sharedDataLPQEX `json:"23"`
	Two5   sharedDataLPQEX `json:"25"`
	Two6   sharedDataLPQEX `json:"26"`
	Two8   sharedDataLPQEX `json:"28"`
	Two9   sharedDataLPQEX `json:"29"`
	Three1 sharedDataLPQEX `json:"31"`
	Three3 sharedDataLPQEX `json:"33"`
	Three4 sharedDataLPQEX `json:"34"`
	Three6 sharedDataLPQEX `json:"36"`
	Three7 sharedDataLPQEX `json:"37"`
	Three9 sharedDataLPQEX `json:"39"`
	Four1  sharedDataLPQEX `json:"41"`
	Four2  sharedDataLPQEX `json:"42"`
	Four3  sharedDataLPQEX `json:"43"`
	Four4  sharedDataLPQEX `json:"44"`
	Four5  sharedDataLPQEX `json:"45"`
	Four6  sharedDataLPQEX `json:"46"`
	Four7  sharedDataLPQEX `json:"47"`
	Four9  sharedDataLPQEX `json:"49"`
	Five0  sharedDataLPQEX `json:"50"`
	Five4  sharedDataLPQEX `json:"54"`
	Five5  sharedDataLPQEX `json:"55"`
	Five8  sharedDataLPQEX `json:"58"`
	Five9  sharedDataLPQEX `json:"59"`
	Six2   sharedDataLPQEX `json:"62"`
	Six5   sharedDataLPQEX `json:"65"`
	Six6   sharedDataLPQEX `json:"66"`
	Six7   sharedDataLPQEX `json:"67"`
	Six8   sharedDataLPQEX `json:"68"`
	Six9   sharedDataLPQEX `json:"69"`
	Seven0 sharedDataLPQEX `json:"70"`
	Seven1 sharedDataLPQEX `json:"71"`
	Seven2 sharedDataLPQEX `json:"72"`
	Seven3 sharedDataLPQEX `json:"73"`
	Seven4 sharedDataLPQEX `json:"74"`
	Seven5 sharedDataLPQEX `json:"75"`
	Seven7 sharedDataLPQEX `json:"77"`
	Seven8 sharedDataLPQEX `json:"78"`
	Eight0 sharedDataLPQEX `json:"80"`
	Eight4 sharedDataLPQEX `json:"84"`
	Eight5 sharedDataLPQEX `json:"85"`
	Eight7 sharedDataLPQEX `json:"87"`
	Eight9 sharedDataLPQEX `json:"89"`
	Nine2  sharedDataLPQEX `json:"92"`
	Nine3  sharedDataLPQEX `json:"93"`
	Nine5  sharedDataLPQEX `json:"95"`
	Nine6  sharedDataLPQEX `json:"96"`
	Nine7  sharedDataLPQEX `json:"97"`
	Nine8  sharedDataLPQEX `json:"98"`
	Nine9  sharedDataLPQEX `json:"99"`*/
}
