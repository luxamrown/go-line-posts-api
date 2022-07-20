package main

import (
	// "bytes"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MediaStyleT struct {
	DisplayType string `json:"displayType"`
}

func NewMediaStyleT(displayType string) MediaStyleT {
	return MediaStyleT{
		DisplayType: displayType,
	}
}

type ContentStyleT struct {
	MediaStyle MediaStyleT `json:"mediaStyle"`
}

func NewContentStyleT(mediaStyle MediaStyleT) ContentStyleT {
	return ContentStyleT{
		MediaStyle: mediaStyle,
	}
}

type ContentT struct {
	Text          string        `json:"text"`
	ContentsStyle ContentStyleT `json:"contentsStyle"`
}

func NewContentT(text string, contentStyle ContentStyleT) ContentT {
	return ContentT{
		Text:          text,
		ContentsStyle: contentStyle,
	}
}

type ReadPermissionT struct {
	Gids []string `json:"gids"`
	Type string   `json:"type"`
}

func NewReadPermissionT(gids []string, types string) ReadPermissionT {
	return ReadPermissionT{
		Gids: gids,
		Type: types,
	}
}

type Post struct {
	Contents ContentT        `json:"contents"`
	PostInfo ReadPermissionT `json:"postInfo"`
}

func NewPost(content ContentT, postInfo ReadPermissionT) Post {
	return Post{
		Contents: content,
		PostInfo: postInfo,
	}
}

type Header struct {
	Cookie  string `header:"Cookie"`
	Referer string `header:"Referer"`
}

func NewHeader(cookie string, referer string) Header {
	return Header{
		Cookie:  cookie,
		Referer: referer,
	}
}

func main() {

	const cookie = ``

	const referer = `https://linevoom.line.me/?utm_source=chrome_app&utm_medium=windows`

	const url = "https://linevoom.line.me/api/post/create?sourceType=TIMELINE"

	// header := NewHeader(cookie, referer)

	newMediaStyle := NewMediaStyleT("GRID")
	newContentStyle := NewContentStyleT(newMediaStyle)
	newContent := NewContentT("testing this is post from golang http", newContentStyle)
	newReadPermission := NewReadPermissionT([]string{}, "ALL")
	newPost := Post{
		newContent,
		newReadPermission,
	}
	jsonPost, err := json.Marshal(newPost)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPost))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Referer", referer)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
