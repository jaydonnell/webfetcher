package webfetcher


import (
	"testing"
	"io/ioutil"
	"fmt"
	"bytes"
)


func TestGetOGInfo(t *testing.T){
	dat, _ := ioutil.ReadFile("testdata/vox.html")
	r := bytes.NewReader(dat)
	pageInfo, _ := GetInfo(r)

	fmt.Println(pageInfo.Title)
	if pageInfo.Title() != "John Oliver's eye-opening, haunting segment on mandatory minimum sentences for drugs" {
		t.Errorf("Failed to parse correct title, got: %v\n", pageInfo.Title())
	}
	if pageInfo.ContentType() != "article" {
		t.Errorf("Failed to parse correct content type")
	}
	if pageInfo.CanonicalURL() != "http://www.vox.com/2015/7/27/9045643/john-oliver-mandatory-minimums" {
		t.Errorf("Failed to parse correct canonical url, got: %v\n", pageInfo.CanonicalURL())
	}
	if pageInfo.Description() != "We have 2 million people incarcerated. If we keep going this direction, we'll soon have enough to populate an entire new country with prisoners." {
		t.Errorf("Failed to parse correct description")
	}
	if pageInfo.ImageURL() != "https://cdn1.vox-cdn.com/thumbor/kZqxT4iz7t_-2KofHjNKPLvdjcU=/0x3:1277x712/1080x600/cdn0.vox-cdn.com/uploads/chorus_image/image/46836584/John_20Oliver_20mandatory_20minimums.0.png" {
		t.Errorf("Failed to parse correct image")
	}


	dat, _ = ioutil.ReadFile("testdata/nytimes.html")
	r = bytes.NewReader(dat)
	pageInfo, _ = GetInfo(r)

	fmt.Println(pageInfo.Title())
	if pageInfo.Title() != "The Governing Cancer of Our Time" {
		t.Errorf("Failed to parse correct title, got %v\n", pageInfo.Title())
	}
	if pageInfo.ContentType() != "" {
		t.Errorf("Failed to parse correct content type, got: %v", pageInfo.ContentType())
	}
	if pageInfo.CanonicalURL() != "http://www.nytimes.com/2016/02/26/opinion/the-governing-cancer-of-our-time.html" {
		t.Errorf("Failed to parse correct canonical url")
	}
	if pageInfo.Description() != "Donald Trump’s candidacy is the culmination of 30 years of antipolitics." {
		t.Errorf("Failed to parse correct description, got: %v", pageInfo.Description())
	}
	if pageInfo.ImageURL() != "https://cdn1.nyt.com/images/2014/11/01/opinion/brooks-circular/brooks-circular-thumbLarge-v4.png" {
		t.Errorf("Failed to parse correct image, got: %v", pageInfo.ImageURL())
	}
}
