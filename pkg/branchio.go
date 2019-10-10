package voi

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

// VerifyLoginMail takes the Link or token behind the "Open the app" button in the login email and returns a token that can
// be used to start a user session. Valid values for `link` are either:
// * https://link.voiapp.io/ohJ7i9ahr0
// * ohJ7i9ahr0
func (c *Client) VerifyLoginMail(link string) (string, error) {

	universalLinkURL := baseUrlLink
	if u, err := url.Parse(link); err == nil {
		universalLinkURL.Path = path.Join(universalLinkURL.Path, u.Path)
	} else {
		universalLinkURL.Path = path.Join(universalLinkURL.Path, link)
	}

	nowUnixMillis := time.Now().UnixNano() / int64(time.Millisecond)

	payload := &BranchIoOpenRequest{
		AdTrackingEnabled:         true,
		AppVersion:                "5.5.2",
		AppleAdAttributionChecked: false,
		BranchKey:                 branchKey,
		Brand:                     "Apple",
		CD: CD{
			Mv: "-1",
			Pn: "com.voiapp.voi",
		},
		Country:                "DE",
		Debug:                  false,
		DeviceFingerprintID:    "184776525281839033",
		FacebookAppLinkChecked: false,
		FirstInstallTime:       nowUnixMillis,
		HardwareID:             uuid.New().String(),
		HardwareIDType:         "idfa",
		IdentityID:             "707679665586682175",
		Instrumentation: Instrumentation{
			V1CloseBrtt: "236",
		},
		IosBundleID:        "com.voiapp.voi",
		IosVendorID:        uuid.New().String(),
		IsHardwareIDReal:   true,
		Language:           "en",
		LastestUpdateTime:  nowUnixMillis,
		LatestInstallTime:  nowUnixMillis,
		LocalIP:            "100.76.32.213",
		Model:              "iPhone10,6",
		OS:                 "iOS",
		OSVersion:          "13.1.2",
		PreviousUpdateTime: nowUnixMillis,
		RetryNumber:        0,
		ScreenHeight:       2436,
		ScreenWidth:        1125,
		SDK:                "ios0.26.0",
		UniversalLinkURL:   universalLinkURL.String(),
		Update:             1,
		URIScheme:          "voiapp",
		UserAgent:          "Mozilla/5.0 (iPhone; CPU iPhone OS 13_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	u := baseUrlBranch
	u.Path = path.Join(baseUrlLink.Path, "/v1/open")

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	req.Header.Set("accept-language", "en-us")
	req.Header.Set("content-type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		break
	case http.StatusBadRequest:
		return "", &ErrorVoiBadRequest{Response: resp}
	default:
		return "", &ErrorVoiUnexpectedResponseCode{Response: resp}
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	r := &BranchIoOpenResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return "", err
	}

	data := &BranchIoOpenData{}
	err = json.Unmarshal([]byte(r.Data), data)
	if err != nil {
		return "", err
	}

	return data.Token, nil
}

type BranchIoOpenResponse struct {
	Data                string `json:"data"`
	DeviceFingerprintID string `json:"device_fingerprint_id"`
	IdentityID          string `json:"identity_id"`
	Link                string `json:"link"`
	SessionID           string `json:"session_id"`
}

type BranchIoOpenData struct {
	MatchGuaranteed   bool    `json:"+match_guaranteed"`
	CreationSource    int64   `json:"~creation_source"`
	ClickedBranchLink bool    `json:"+clicked_branch_link"`
	ClickTimestamp    int64   `json:"+click_timestamp"`
	OneTimeUse        bool    `json:"$one_time_use"`
	ID                float64 `json:"~id"`
	IsFirstSession    bool    `json:"+is_first_session"`
	ReferringLink     string  `json:"~referring_link"`
	Token             string  `json:"token"`
}

type BranchIoOpenRequest struct {
	AdTrackingEnabled         bool            `json:"ad_tracking_enabled"`
	AppVersion                string          `json:"app_version"`
	AppleAdAttributionChecked bool            `json:"apple_ad_attribution_checked"`
	BranchKey                 string          `json:"branch_key"`
	Brand                     string          `json:"brand"`
	CD                        CD              `json:"cd"`
	Country                   string          `json:"country"`
	Debug                     bool            `json:"debug"`
	DeviceFingerprintID       string          `json:"device_fingerprint_id"`
	FacebookAppLinkChecked    bool            `json:"facebook_app_link_checked"`
	FirstInstallTime          int64           `json:"first_install_time"`
	HardwareID                string          `json:"hardware_id"`
	HardwareIDType            string          `json:"hardware_id_type"`
	IdentityID                string          `json:"identity_id"`
	Instrumentation           Instrumentation `json:"instrumentation"`
	IosBundleID               string          `json:"ios_bundle_id"`
	IosVendorID               string          `json:"ios_vendor_id"`
	IsHardwareIDReal          bool            `json:"is_hardware_id_real"`
	Language                  string          `json:"language"`
	LastestUpdateTime         int64           `json:"lastest_update_time"`
	LatestInstallTime         int64           `json:"latest_install_time"`
	LocalIP                   string          `json:"local_ip"`
	Model                     string          `json:"model"`
	OS                        string          `json:"os"`
	OSVersion                 string          `json:"os_version"`
	PreviousUpdateTime        int64           `json:"previous_update_time"`
	RetryNumber               int64           `json:"retryNumber"`
	ScreenHeight              int64           `json:"screen_height"`
	ScreenWidth               int64           `json:"screen_width"`
	SDK                       string          `json:"sdk"`
	UniversalLinkURL          string          `json:"universal_link_url"`
	Update                    int64           `json:"update"`
	URIScheme                 string          `json:"uri_scheme"`
	UserAgent                 string          `json:"user_agent"`
}

type CD struct {
	Mv string `json:"mv"`
	Pn string `json:"pn"`
}

type Instrumentation struct {
	V1CloseBrtt string `json:"/v1/close-brtt"`
}
