package parser

type Video struct {
	ID            int           `json:"vid"`
	Title         string        `json:"title"`
	CreatedAt     int           `json:"createdAt"`
	DownloadTimes int           `json:"downloadTimes"`
	SoftwareType  *SoftwareType `json:"softwareType"`
	IsRecommended bool          `json:"isRecommended"`
}

type SoftwareType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Seller struct {
	UID      int    `json:"uid"`
	UserName string `json:"username"`
}

type Sale struct {
	Video            *Video  `json:"video"`
	Seller           *Seller `json:"seller"`
	LicType          string  `json:"licType"`
	Keyword          *string `json:"keyword"`
	Timestamp        int     `json:"timestamp"`
	TimestampTimeAgo string  `json:"timestampTimeAgo"`
}

type ResData struct {
	URL   string `json:"url"`
	State *State `json:"state"`
}

type State struct {
	LoaderData *LoaderData `json:"loaderData"`
}

type LoaderData struct {
	Sales []*Sale `json:"routes/_landing.ranking.sales"`
}
