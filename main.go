package main

type MediaStyleT struct {
	DisplayType string
}

func NewMediaStyleT(displayType string) MediaStyleT {
	return MediaStyleT{
		DisplayType: displayType,
	}
}

type ContentStyleT struct {
	MediaStyle MediaStyleT
}

func NewContentStyleT(mediaStyle MediaStyleT) ContentStyleT {
	return ContentStyleT{
		MediaStyle: mediaStyle,
	}
}

type ContentT struct {
	Text          string
	ContentsStyle ContentStyleT
}

func NewContentT(text string, contentStyle ContentStyleT) ContentT {
	return ContentT{
		Text:          text,
		ContentsStyle: contentStyle,
	}
}

type ReadPermissionT struct {
	Gids []string
	Type string
}

func NewReadPermissionT(gids []string, types string) ReadPermissionT {
	return ReadPermissionT{
		Gids: gids,
		Type: types,
	}
}

type Post struct {
	Content  ContentT
	PostInfo ReadPermissionT
}

func NewPost(content ContentT, postInfo ReadPermissionT) Post {
	return Post{
		Content:  content,
		PostInfo: postInfo,
	}
}

func main() {

}
