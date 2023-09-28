package schema

type SiteMap struct {
	SiteMaps []Loc `xml:"sitemap"`
}

type SiteMapUrl struct {
	Urls []Loc `xml:"url"`
}

type Loc struct {
	Loc      string `xml:"loc"`
	Priority string `xml:"priority"`
}
