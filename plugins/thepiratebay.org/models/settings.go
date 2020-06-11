package models

import (
	"github.com/jinzhu/gorm"
)

type SettingBloomberg struct {
	gorm.Model
	Enabled         bool
	SitemapURL      string
	AllowedDomains  []Domain
	CacheDir        string
	ConsumerThreads int
	QueueMaxSize    int
}

type Domain struct {
	Address string
}
