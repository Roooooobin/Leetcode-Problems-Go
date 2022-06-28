package main

import (
	"strconv"
	"strings"
)

type Codec struct {
	mp map[int]string
	id int
}

func Constructor() Codec {
	mp := make(map[int]string)
	return Codec{mp, 0}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {

	this.id++
	this.mp[this.id] = longUrl
	return "http://tinyurl.com/" + strconv.Itoa(this.id)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {

	idx := strings.LastIndexByte(shortUrl, '/')
	id, _ := strconv.Atoi(shortUrl[idx+1:])
	return this.mp[id]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
