package model

type VideoDto struct {
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

type SellerDto struct {
	UID      int    `json:"uid"`
	UserName string `json:"username"`
}

type SaleDto struct {
	Video            *VideoDto  `json:"video"`
	Seller           *SellerDto `json:"seller"`
	LicType          string     `json:"licType"`
	Keyword          *string    `json:"keyword"`
	Timestamp        int        `json:"timestamp"`
	TimestampTimeAgo string     `json:"timestampTimeAgo"`
}

type ResDataDto struct {
	URL   string    `json:"url"`
	State *StateDto `json:"state"`
}

type StateDto struct {
	LoaderData *LoaderDataDto `json:"loaderData"`
}

type LoaderDataDto struct {
	Sales []*SaleDto `json:"routes/_landing.ranking.sales"`
}
