# ogg_coverter

用来转换 飞书语音消息到wav格式，以使用openai wisper API

## 使用方法

```golang

import (
	"context"
	"log"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestOggToWavByPath(t *testing.T) {
	oggfile := "ogg.ogg"
	wavfile := "output.wav"

	OggToWavByPath(oggfile, wavfile)

	config := openai.DefaultConfig("token")
	// proxyUrl, err := url.Parse("http://localhost:7890")
	// if err != nil {
	// 	panic(err)
	// }
	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(proxyUrl),
	// }
	// config.HTTPClient = &http.Client{
	// 	Transport: transport,
	// }
	c := openai.NewClientWithConfig(config)

	res, err := c.CreateTranscription(context.Background(), openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: wavfile,
	})
	if err != nil {
		log.Panicln(err)
	}

	println(res.Text)

}
```
