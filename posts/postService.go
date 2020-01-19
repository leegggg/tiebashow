package posts

// GetPostsByParentAndFlrFlag ...
func GetPostsByParentAndFlrFlag(parent int64, flr bool, pn int64, nb int64) ([]Post, error) {
	postHeaders, err := GetPostHeadersByParentAndFlrFlag(parent, flr, pn, nb)
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, postHeader := range postHeaders {
		var post Post
		post.PostHeader = postHeader
		content, err := GetContent(*postHeader.Pid)
		if err == nil {
			post.Content = content
		}
		nbFlr, err := GetNbPostFlr(*postHeader.Pid)
		if err == nil {
			post.NbFlr = nbFlr
		}
		posts = append(posts, post)
	}
	return posts, err
}
