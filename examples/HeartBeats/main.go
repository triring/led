// go get github.com/triring/led
// tinygo build -target=pico2 -size=short -o HeartBeats.uf2 ./main.go
// tinygo flash -target=pico2 -size=short -monitor ./main.go

/*
使用するボードにあわせてtargetを変更すること。
pico
pico-plus2
pico-w
pico2
pico2-ice
pico2-w
*/

package main

import (
	"github.com/triring/led" // githubで公開しているパッケージをインポートする場合の記述
	"machine"
	//	"led"  ローカルのディレクトリに置かれたledのパッケージをインポートする場合の記述
)

func main() {
	// オンボードLEDを初期化
	OnboardLED := led.New(machine.LED)
	// LEDの点滅で心臓の鼓動を表現
	for {
		OnboardLED.Blink(50, 200, 50, 200, 0, 250, 0, 250)
	}
}
