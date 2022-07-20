package model

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

type PostInfoT struct {
	ReadPermission ReadPermissionT `json:"readPermission"`
}

func NewPostInfoT(readPermission ReadPermissionT) PostInfoT {
	return PostInfoT{
		ReadPermission: readPermission,
	}
}

type Post struct {
	Contents ContentT  `json:"contents"`
	PostInfo PostInfoT `json:"postInfo"`
}

func NewPost(content ContentT, postInfo PostInfoT) Post {
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

var PERMISSION = []string{"NONE", "FRIEND", "ALL"}

const Cookie = ``
const Referer = `https://linevoom.line.me/?utm_source=chrome_app&utm_medium=windows`
const Url = "https://linevoom.line.me/api/post/create?sourceType=TIMELINE"
