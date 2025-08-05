package main

import (
	"fmt"
	"regexp"
)

// Regex for raw YouTube links
var youtubeRegex = regexp.MustCompile(`(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/watch\?v=|youtu\.be\/)([A-Za-z0-9_-]{11})`)
var youtubeAnchorRegex = regexp.MustCompile(`<a[^>]+href="(?:https?:\/\/)?(?:www\.)?(?:youtube\.com\/watch\?v=|youtu\.be\/)([A-Za-z0-9_-]{11})"[^>]*>.*?<\/a>`)

func embedYouTubeVideos(content string) string {
	// Replace anchor-wrapped YouTube links
	content = youtubeAnchorRegex.ReplaceAllStringFunc(content, func(anchor string) string {
		matches := youtubeAnchorRegex.FindStringSubmatch(anchor)
		if len(matches) < 2 {
			return anchor
		}
		videoID := matches[1]
		return fmt.Sprintf(
			`<div class="max-w-[560px] aspect-video w-full mx-auto my-4"><iframe class="w-full h-full" src="https://www.youtube.com/embed/%s" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe></div>`,
			videoID,
		)
	})

	// Replace raw YouTube URLs
	content = youtubeRegex.ReplaceAllStringFunc(content, func(url string) string {
		matches := youtubeRegex.FindStringSubmatch(url)
		if len(matches) < 2 {
			return url
		}
		videoID := matches[1]
		return fmt.Sprintf(
			`<div class="max-w-[560px] aspect-video w-full mx-auto my-4"><iframe class="w-full h-full" src="https://www.youtube.com/embed/%s" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe></div>`,
			videoID,
		)
	})

	return content
}
