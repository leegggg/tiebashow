package threads

import (
	"github.com/leegggg/tiebashow/posts"
)

// GetThreads ...
func GetThreads(kw string, pn int64, nb int64, good bool) ([]Thread, error) {
	var err error
	var headers []ThreadHeader
	if !good {
		headers, err = GetThreadHeaders(kw, pn, nb)
	} else {
		headers, err = GetGoodThreadHeaders(kw, pn, nb)
	}
	var threads []Thread
	if err != nil {
		return threads, err
	}
	for _, header := range headers {
		thread, err := GetThreadByHeader(header)
		if err == nil {
			threads = append(threads, thread)
		}
	}
	return threads, err
}

// GetThreadByHeader ...
func GetThreadByHeader(header ThreadHeader) (Thread, error) {
	var thread Thread
	thread.ThreadHeader = header
	firstHeader, err := posts.GetFirstPostHeaderByKz(header.Kz)
	if err == nil {
		first, err := posts.GetPostByPostHeader(firstHeader)
		if err == nil {
			thread.First = first
		}
	}
	lastHeader, err := posts.GetLastPostHeaderByKz(header.Kz)
	if err == nil {
		last, err := posts.GetPostByPostHeader(lastHeader)
		if err == nil {
			thread.Last = last
		}
	}
	return thread, err
}
