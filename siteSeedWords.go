package siteLinkKeyword

import (
	"github.com/kevin-zx/site-info-crawler/sitethrougher"
	"github.com/kevin-zx/siteLinkKeyword/filter"
	"strings"
)

// this GetSiteLinkWords is very strict, put all filter to GetSiteLinkWords, so we'll except get pure word by this function
func GetPureWords(si *sitethrougher.SiteInfo) map[string]int {
	// filter sort is matter, ClearPunctuationsFilter is better in the top of filter
	return GetSiteLinkWords(si, filter.CommStopWordFilter, filter.TrimFilter,
		filter.TailNumFilter, filter.ClearPunctuationsFilter,
		filter.KeywordLenFilter, filter.RestOneFilter, filter.EnvFilter, filter.FormalWordFilter)
}

// get link text word from a site with appear times
// si: site info
// filters: some filters
func GetSiteLinkWords(si *sitethrougher.SiteInfo, filters ...filter.Filter) map[string]int {
	sitethrougher.FillSiteLinksDetailHrefText(si)
	kim := make(map[string]int)
	for _, link := range si.SiteLinks {
		for key, info := range link.DetailHrefTexts {
			key = clear(key)
			if key == "___img___" {
				continue
			}
			kLen := len(strings.Split(key, ""))
			if kLen >= 15 || kLen <= 1 {
				continue
			}
			if info.Count >= 1 {
				if _, ok := kim[key]; ok {
					kim[key] += info.Count
				} else {
					kim[key] = info.Count
				}
			}
		}
	}
	for _, f := range filters {
		kim = f(kim)
	}

	return kim
}
func clear(key string) string {
	key = strings.ReplaceAll(key, " ", "")
	key = strings.ReplaceAll(key, " ", "")
	key = strings.ReplaceAll(key, "|", "")
	key = strings.ReplaceAll(key, "\r", "")
	key = strings.ReplaceAll(key, "\n", "")
	key = strings.ReplaceAll(key, "\t", "")
	return key
}
