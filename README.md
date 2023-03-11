# ogg_coverter

用来转换 飞书语音消息到wav格式，以使用openai wisper API

## 使用方法

函数1：`OggToWavByPath(input string, output string) error`
```golang
func TestOggToWavByPath(t *testing.T) {
	oggfile := "ogg.ogg"
	wavfile := "output.wav"
	OggToWavByPath(oggfile, wavfile)
}
```
函数2：`OggToWav(input io.Reader, output io.WriteSeeker) error`
```golang

func main() {
	input, err := os.Open(ogg)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(wav)
	if err != nil {
		return err
	}

	defer output.Close()
	return OggToWav(input, output)
}
```