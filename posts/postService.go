package posts

// GetPostsByParentAndFlrFlag ...
func GetPostsByParentAndFlrFlag(parent int64, flr bool, pn int64, nb int64) ([]Post, error) {
	postHeaders, err := GetPostHeadersByParentAndFlrFlag(parent, flr, pn, nb)
	var posts []Post
	if err != nil {
		return posts, err
	}
	for _, postHeader := range postHeaders {
		post, _ := GetPostByPostHeader(postHeader)
		posts = append(posts, post)
	}
	return posts, err
}

// GetPostByPostHeader ...
func GetPostByPostHeader(header PostHeader) (Post, error) {
	var post Post
	post.PostHeader = header
	content, err := GetContent(*header.Pid)
	if err == nil {
		post.Content = content
	}
	nbFlr, err := GetNbPostFlr(*header.Pid)
	if err == nil {
		post.NbFlr = nbFlr
	}
	urls, err := getAttachementUrlsByPost(post)
	if err == nil {
		post.Attachements = urls
	}
	return post, err
}

func getAttachementUrlsByPost(post Post) ([]string, error) {
	attachementPrefix := "/img/"
	var urls []string
	attachements, err := GetAttachementHeadersByPostHeader(post.PostHeader)
	if err != nil {
		return urls, err
	}
	for _, attachement := range attachements {
		if attachement.Status == 200 {
			urls = append(urls, attachementPrefix+*attachement.Path)
		} else {
			urls = append(urls, *attachement.Link)
		}
	}
	return urls, err
}
