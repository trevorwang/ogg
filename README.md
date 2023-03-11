# ogg_coverter

用来转换 飞书语音消息到wav格式，以使用openai wisper API

## 使用方法

```golang
func TestOggToWavByPath(t *testing.T) {
	oggfile := "ogg.ogg"
	wavfile := "output.wav"
	OggToWavByPath(oggfile, wavfile)
}
```