package voi

import (
	"net/url"
)

const branchKey = "key_live_ipzWJnGsJAiE6TBGY4NIGjmdADk1MTOx"

var baseUrlApi = url.URL{
	Scheme: "https",
	Host:   "api.voiapp.io",
	Path:   "v1",
}

var baseUrlLink = url.URL{
	Scheme: "https",
	Host:   "link.voiapp.io",
}

var baseUrlBranch = url.URL{
	Scheme: "https",
	Host:   "api2.branch.io",
	Path:   "v1",
}
