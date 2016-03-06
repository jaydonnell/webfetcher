package webfetcher

import (
	"io"
	"golang.org/x/net/html"
	"strings"
)

type PageMetaData interface {
    Title() string
    Description() string
    ContentType() string
    CanonicalURL() string
    ImageURL() string
}

type PageInfo struct {
	//Title string
	// Description string
	// ContentType string
	// Image string
	// CanonicalURL string
	OGProps map[string]string
    TwitterCardProps map[string]string
    InferredProps map[string]string
}

func (p PageInfo) Title() string {
  v, ok := p.OGProps["og:title"]
  if ok {
      return v
  }
  v, ok = p.TwitterCardProps["twitter:title"]
  if ok {
      return v
  }
  return ""  
}

func (p PageInfo) Description() string {
  v, ok := p.OGProps["og:description"]
  if ok {
      return v
  }
  v, ok = p.TwitterCardProps["twitter:description"]
  if ok {
      return v
  }    
  return ""  
}

func (p PageInfo) ContentType() string {
      v, ok := p.OGProps["og:type"]
  if ok {
      return v
  }
  return ""  
}

func (p PageInfo) CanonicalURL() string {
  v, ok := p.InferredProps["canonical"]
  if ok {
      return v
  }  
  return ""  
}

func (p PageInfo) ImageURL() string {
  v, ok := p.OGProps["og:image"]
  if ok {
      return v
  }
  v, ok = p.TwitterCardProps["twitter:image"]
  if ok {
      return v
  }   
  return ""  
}

func ExtractMetaData(doc io.Reader) (PageInfo, error) {
	twitterPrefix := "twitter"
	ogPrefix := "og"
    pi := PageInfo{}
    pi.InferredProps = make(map[string]string)
    pi.OGProps = make(map[string]string)
    pi.TwitterCardProps = make(map[string]string)
	
	z := html.NewTokenizer(doc)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if z.Err() == io.EOF {
				return pi, nil
			}
			return pi, z.Err()
		}

		t := z.Token()

		if t.Type == html.EndTagToken && t.Data == "head" {
			return pi, nil
		}

		if t.Data == "link" {
			isCanonical := false
			href := ""
			for _, a := range t.Attr {
				if a.Key == "rel" && a.Val == "canonical"{
				 isCanonical = true
				}
				if a.Key == "href" {
					href = a.Val
				}
			}

			if isCanonical {
				pi.InferredProps["canonical"] = href
			}

		}

		if t.Data == "meta" {
			var prop, cont string
			for _, a := range t.Attr {
				switch a.Key {
				case "property":
					prop = a.Val
				case "content":
					cont = a.Val
				}
			}

			if strings.HasPrefix(prop, twitterPrefix) && cont != "" {
                pi.TwitterCardProps[prop] = cont    
            } else if strings.HasPrefix(prop, ogPrefix) && cont != "" {
                pi.OGProps[prop] = cont
			}
		}
	}

	return pi, nil
}

func GetInfo(r io.Reader) (PageMetaData, error) {
	pi, err := ExtractMetaData(r)
	return pi, err
}