package warmup

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/valyala/fasthttp"

	"go-command/pkg/utils/print"
	"go-command/pkg/warmup/schema"
	"go-command/pkg/warmup/services"
)

type urlPool struct {
	Url     string
	Headers map[string]string
}

type Sitemap struct {
	MaxWorker        int
	MaxRequestInTime int
	MaxRunTime       time.Duration
}

func (s *Sitemap) WarmUp(siteMap schema.SiteMap) {

	siteMapService := services.SiteMap{
		Client: &fasthttp.Client{},
	}

	allUrl := make(chan urlPool)
	var wgSiteMap sync.WaitGroup

	for _, url := range siteMap.SiteMaps {
		wgSiteMap.Add(1)
		go func(url string) {
			defer wgSiteMap.Done()
			siteMapUrl, err := siteMapService.GetSiteMapUrl(url)
			if err != nil {
				fmt.Println("Get sitemap error: ", err, ", url:", url)
			}

			for _, url := range siteMapUrl.Urls {
				for _, headers := range s.multiplexerByHeader(url.Loc) {
					allUrl <- urlPool{Url: url.Loc, Headers: headers}
				}
			}
		}(url.Loc)
	}

	var wgWorker sync.WaitGroup

	var startTime time.Time = time.Now()

	count := 0
	for i := 0; i < s.MaxWorker; i++ {
		wgWorker.Add(1)
		go func(i int, count *int) {
			defer wgWorker.Done()
			for pool := range allUrl {
				s.worker(pool, i, count, &startTime, siteMapService)
			}
		}(i, &count)
	}

	wgSiteMap.Wait()
	close(allUrl)
	wgWorker.Wait()

	fmt.Println("Total count:", count)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Passing time : %v\n", elapsedTime)

	fmt.Println("Done.")

}

func (s *Sitemap) worker(pool urlPool, workerId int, count *int, startTime *time.Time, siteMapService services.SiteMap) {
	startRequestTime := time.Now()
	(*count)++

	status, err := siteMapService.VisitUrl(pool.Url, pool.Headers)

	if err != nil {
		fmt.Println("*error request url, status:", status, ", url:", pool.Url, time.Since(startRequestTime))
	} else {
		fmt.Println("worker: "+strconv.Itoa(workerId+1)+", count:", (*count), ", visited:", pool.Url, time.Since(startRequestTime))
	}

	if (*count)%s.MaxRequestInTime == 0 {
		elapsedTime := time.Since(*startTime)
		if elapsedTime < s.MaxRunTime {
			go print.CountDown(s.MaxRunTime-elapsedTime, "worker "+strconv.Itoa(workerId+1))
			time.Sleep(s.MaxRunTime - elapsedTime)
		}
		*startTime = time.Now()
	}
}

func (s *Sitemap) multiplexerByHeader(url string) []map[string]string {
	return []map[string]string{
		{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		},
		{
			"User-Agent": "Iphone",
			"Accept":     "image/webp",
		},
		{
			"User-Agent": "Iphone",
		},
		{
			"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.89 Safari/537.36",
			"Cache-Control":   "max-age=0",
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
			"HTTPS":           "1",
			"DNT":             "1",
			"Referer":         "https://www.google.com/",
			"Accept-Encoding": "gzip",
		},
	}
}
