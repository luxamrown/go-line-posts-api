package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"mohamadelabror.me/golinepostsapi/model"
)

func StartPost(post string, privacy string) {
	header := model.NewHeader(model.Cookie, model.Referer)

	newMediaStyle := model.NewMediaStyleT("GRID")
	newContentStyle := model.NewContentStyleT(newMediaStyle)
	newContent := model.NewContentT(post, newContentStyle)
	newReadPermission := model.NewReadPermissionT([]string{}, privacy)
	newPostInfo := model.NewPostInfoT(newReadPermission)
	newPost := model.NewPost(newContent, newPostInfo)
	jsonPost, err := json.Marshal(newPost)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, model.Url, bytes.NewBuffer(jsonPost))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", header.Cookie)
	req.Header.Set("Referer", header.Referer)
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input post: ")
	scanner.Scan()
	input := scanner.Text()
	StartPost(input, "FRIEND")
}
