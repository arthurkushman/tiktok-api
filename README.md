# tiktok-api
TikTok api sdk for simple video embedding

[![Go Report Card](https://goreportcard.com/badge/github.com/arthurkushman/buildsqlx)](https://goreportcard.com/report/github.com/arthurkushman/tiktok-api)
[![GoDoc](https://github.com/golang/gddo/blob/c782c79e0a3c3282dacdaaebeff9e6fd99cb2919/gddo-server/assets/status.svg)](https://godoc.org/github.com/arthurkushman/tiktok-api)
[![codecov](https://codecov.io/gh/arthurkushman/tiktok-api/branch/master/graph/badge.svg)](https://codecov.io/gh/arthurkushman/tiktok-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

### Installation
```go
go get -u github.com/arthurkushman/tiktok-api 
```

### Usage
To get an embedded content of any video:
```go
tts := NewTikTokService()
resp, err := tts.Embed(map[string]string{
    "url": "https://www.tiktok.com/@scout2015/video/6718335390845095173",
})

resp.Title // "Scramble up ur name & I’ll try to guess it😍❤️ #foryoupage #petsoftiktok #aesthetic"
resp.AuthorName // "Scout & Suki"
// and more props - see Response structure bellow
```
`resp` is a `Response` struct with all the fields filled-up e.g.:
```go
type Response struct {
  Version: "1.0",
  Type: "video",
  Title: "Scramble up ur name & I’ll try to guess it😍❤️ #foryoupage #petsoftiktok #aesthetic",
  AuthorUrl: "https://www.tiktok.com/@scout2015",
  AuthorName: "Scout & Suki",
  Width: "100%",
  Height: "100%",
  Html: "<blockquote class=\"tiktok-embed\" cite=\"https://www.tiktok.com/@scout2015/video/6718335390845095173\" data-video-id=\"6718335390845095173\" style=\"max-width: 605px;min-width: 325px;\" > <section> <a target=\"_blank\" title=\"@scout2015\" href=\"https://www.tiktok.com/@scout2015\">@scout2015</a> <p>Scramble up ur name & I’ll try to guess it😍❤️ <a title=\"foryoupage\" target=\"_blank\" href=\"https://www.tiktok.com/tag/foryoupage\">#foryoupage</a> <a title=\"petsoftiktok\" target=\"_blank\" href=\"https://www.tiktok.com/tag/petsoftiktok\">#petsoftiktok</a> <a title=\"aesthetic\" target=\"_blank\" href=\"https://www.tiktok.com/tag/aesthetic\">#aesthetic</a></p> <a target=\"_blank\" title=\"♬ original sound - 𝐇𝐚𝐰𝐚𝐢𝐢𓆉\" href=\"https://www.tiktok.com/music/original-sound-6689804660171082501\">♬ original sound - 𝐇𝐚𝐰𝐚𝐢𝐢𓆉</a> </section> </blockquote> <script async src=\"https://www.tiktok.com/embed.js\"></script>",
  ThumbnailWidth: 720,
  ThumbnailHeight: 1280,
  ThumbnailUrl: "https://p16.muscdn.com/obj/tos-maliva-p-0068/06kv6rfcesljdjr45ukb0000d844090v0200010605",
  ProviderUrl: "https://www.tiktok.com",
  ProviderName: "TikTok"
}
```
It allows you to get the embed code and additional information about the video associated with the webpage link provided.