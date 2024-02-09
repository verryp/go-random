package main

import (
	"fmt"
	"strings"
	"testing"
)

// func main() {
// 	// bot := `Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML; like Gecko) Chrome/104.0.5112.79 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`
// 	//
// 	// fmt.Println(strings.Contains(bot, `bot`))
// }

func TestBotMatcher(t *testing.T) {
	t.Parallel()

	macher := `bot`
	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; AhrefsBot/7.0; +http://ahrefs.com/robot/)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; SemrushBot/7~bl; +http://www.semrush.com/bot.html)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0+(compatible; UptimeRobot/2.0; http://www.uptimerobot.com/)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible;PetalBot;+https://webmaster.petalsearch.com/site/petalbot)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML; like Gecko) Chrome/104.0.5112.79 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; MJ12bot/v1.4.8; http://mj12bot.com/)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML; like Gecko) Chrome/104.0.5112.101 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/605.1.15 (KHTML; like Gecko) Version/13.1.1 Safari/605.1.15 (Applebot/0.1; +http://www.apple.com/go/applebot)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 AppleWebKit/537.36 (KHTML; like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Chrome/104.0.5112.101 Safari/537.36`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; AhrefsSiteAudit/6.1; +http://ahrefs.com/robot/)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 AppleWebKit/537.36 (KHTML; like Gecko; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) Chrome/103.0.5060.134 Safari/537.36`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; DotBot/1.2; +https://opensiteexplorer.org/dotbot; help@moz.com)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML; like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/601.2.4 (KHTML; like Gecko) Version/9.0.1 Safari/601.2.4 facebookexternalhit/1.1 Facebot Twitterbot/1.0`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML; like Gecko) Chrome/105.0.5195.102 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML; like Gecko) Chrome/105.0.5195.52 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 AppleWebKit/537.36 (KHTML; like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Chrome/104.0.5112.79 Safari/537.36`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; SeekportBot; +https://bot.seekport.com)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots) AppleWebKit/537.36 (KHTML; like Gecko) Chrome/81.0.4044.268`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 AppleWebKit/537.36 (KHTML; like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Safari/537.36`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; AhrefsSiteAudit/6.1; +http://ahrefs.com/robot/site-audit)`,
			want:  true,
		},
		{
			input: `Mozilla/5.0 (compatible; SemrushBot-BA; +http://www.semrush.com/bot.html)`,
			want:  true,
		},
		{
			input: `lorem ipsum`,
			want:  false,
		},
		{
			input: `bto tob tbo b-o-t bo t boot`,
			want:  false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf(`test cases-%d`, i), func(t *testing.T) {
			got := strings.Contains(tc.input, macher)
			if got != tc.want {
				t.Errorf(`want %v but got %v`, tc.want, got)
			}
		})
	}

}
