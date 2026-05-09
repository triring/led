// go get github.com/triring/led
// tinygo build -target=pico2 -size=short -o HeartBeats.uf2 .
// tinygo flash -target=pico2 -size=short -monitor .

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
	"github.com/triring/led" // githubで公開しているパッケージをインポートする場合
	"machine"
	//	"led"  ローカルのディレクトリに置かれたledのパッケージをインポートする場合
)

func main() {
	// オンボードLEDを初期化
	LED := led.New(machine.LED)
	// LEDの点滅で心臓の鼓動を表現
	for {
		LED.Blink(50, 200, 50, 200, 0, 250, 0, 250)
	}
}
