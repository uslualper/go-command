package services

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"

	"go-command/pkg/utils/http"
	"go-command/pkg/warmup/schema"
)

type SiteMap struct {
	Client *fasthttp.Client
}

func (s *SiteMap) GetSiteMap(url string) (schema.SiteMap, error) {
	client := http.NewClient(s.Client)
	client.SetTimeout(20)

	response, status := client.Get(url)

	siteMap := schema.SiteMap{}

	if status == fasthttp.StatusOK {
		if err := xml.Unmarshal(response, &siteMap); err != nil {
			return siteMap, err
		}
		return siteMap, nil
	}

	return siteMap, errors.New("GET site map error: " + strconv.Itoa(status) + ", " + url)
}

func (s *SiteMap) GetSiteMapUrl(url string) (schema.SiteMapUrl, error) {
	client := http.NewClient(s.Client)
	client.SetTimeout(30)

	response, status := client.Get(url)

	if status != 200 {
		return schema.SiteMapUrl{}, errors.New("GET site map url error: status code " + fmt.Sprintf("%d", status))
	}

	var siteMapUrl schema.SiteMapUrl
	if err := xml.Unmarshal(response, &siteMapUrl); err != nil {
		return schema.SiteMapUrl{}, errors.New("xml unmarshal error: " + err.Error())
	}

	return siteMapUrl, nil
}

func (s *SiteMap) VisitUrl(url string, headers map[string]string) (int, error) {
	client := http.NewClient(s.Client)
	client.SetTimeout(5)

	for key, header := range headers {
		client.AddHeader(key, header)
	}

	_, status := client.Get(url)

	if status != 200 {
		return status, errors.New("GET url error: status code " + fmt.Sprintf("%d", status))
	}

	return status, nil
}
