package services

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"

	"go-command/pkg/utils/http"
	"go-command/pkg/warmup/schema"
)

type SiteMap struct {
	Client *fasthttp.Client
}

func (s *SiteMap) GetSiteMap(url string) (schema.SiteMap, error) {
	client := http.Client{Client: s.Client}
	client.Init(url)
	client.SetTimeout(20)
	response, status := client.Get()

	siteMap := schema.SiteMap{}

	if status == 200 {
		if err := xml.Unmarshal(response, &siteMap); err != nil {
			return siteMap, err
		}
		return siteMap, nil
	}

	return siteMap, errors.New("GET site map error" + fmt.Sprintf(" %d", status) + ", " + url)
}

func (s *SiteMap) GetSiteMapUrl(url string) (schema.SiteMapUrl, error) {
	client := http.Client{Client: s.Client}
	client.Init(url)
	client.SetTimeout(30)
	response, status := client.Get()

	siteMapUrl := schema.SiteMapUrl{}

	if status == 200 {
		if err := xml.Unmarshal(response, &siteMapUrl); err != nil {
			return siteMapUrl, errors.New("xml unmarshal error: " + err.Error() + ", " + url)
		}
		return siteMapUrl, nil
	}

	return siteMapUrl, errors.New("GET site map url error" + fmt.Sprintf(" %d", status) + ", " + url)
}

func (s *SiteMap) VisitUrl(url string, headers map[string]string) (status int, err error) {
	client := http.Client{Client: s.Client}
	client.Init(url)
	client.SetTimeout(5)

	for key, header := range headers {
		client.AddHeader(key, header)
	}

	_, status = client.Get()

	if status != 200 {
		return status, errors.New("GET url error")
	}

	return status, nil
}
