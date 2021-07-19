package model

import "time"

type BingHPImageArchive struct {
	Images   []BingImage
	ToolTips ToolTips
	UpdateAt time.Time
}

type BingImage struct {
	StartDate     string
	FullstartDate string
	EndDate       string
	URL           string
	UrlBase       string
	Copyright     string
	CopyrightLink string
	Title         string
	Quiz          string
	Wp            bool
	Hsh           string
	Drk           int
	Top           int
	Bot           int
	Hs            []interface{}
}

type ToolTips struct {
	Loading  string
	Previous string
	Next     string
	Walle    string
	Walls    string
}
