package tiktok_api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var jsonStr = `{"version":"1.0","type":"video","title":"Scramble up ur name & I’ll try to guess it😍❤️ #foryoupage #petsoftiktok #aesthetic","author_url":"https://www.tiktok.com/@scout2015","author_name":"Scout & Suki","width":"100%","height":"100%","html":"<blockquote class=\"tiktok-embed\" cite=\"https://www.tiktok.com/@scout2015/video/6718335390845095173\" data-video-id=\"6718335390845095173\" style=\"max-width: 605px;min-width: 325px;\" > <section> <a target=\"_blank\" title=\"@scout2015\" href=\"https://www.tiktok.com/@scout2015\">@scout2015</a> <p>Scramble up ur name & I’ll try to guess it😍❤️ <a title=\"foryoupage\" target=\"_blank\" href=\"https://www.tiktok.com/tag/foryoupage\">#foryoupage</a> <a title=\"petsoftiktok\" target=\"_blank\" href=\"https://www.tiktok.com/tag/petsoftiktok\">#petsoftiktok</a> <a title=\"aesthetic\" target=\"_blank\" href=\"https://www.tiktok.com/tag/aesthetic\">#aesthetic</a></p> <a target=\"_blank\" title=\"♬ original sound - 𝐇𝐚𝐰𝐚𝐢𝐢𓆉\" href=\"https://www.tiktok.com/music/original-sound-6689804660171082501\">♬ original sound - 𝐇𝐚𝐰𝐚𝐢𝐢𓆉</a> </section> </blockquote> <script async src=\"https://www.tiktok.com/embed.js\"></script>","thumbnail_width":720,"thumbnail_height":1280,"thumbnail_url":"https://p16.muscdn.com/obj/tos-maliva-p-0068/06kv6rfcesljdjr45ukb0000d844090v0200000a05","provider_url":"https://www.tiktok.com","provider_name":"TikTok"}`

func TestTikTokService_Embed(t *testing.T) {
	tts := NewTikTokService()
	resp, err := tts.Embed(map[string]string{
		"url": "https://www.tiktok.com/@scout2015/video/6718335390845095173",
	})
	assert.NoError(t, err)

	r := &Response{}
	err = json.Unmarshal([]byte(jsonStr), r)
	assert.NoError(t, err)

	assert.Equal(t, resp.Type, r.Type)
	assert.Equal(t, resp.Version, r.Version)
	assert.Equal(t, resp.AuthorUrl, r.AuthorUrl)
	assert.Equal(t, resp.ProviderUrl, r.ProviderUrl)
	assert.Equal(t, resp.ProviderName, r.ProviderName)
}
