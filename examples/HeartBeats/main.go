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
	"machine"

	"github.com/triring/led" // githubで公開しているパッケージをインポートする場合
	//	"led"  ローカルのディレクトリに置かれたledのパッケージをインポートする場合
)

func main() {
	// オンボードLEDを初期化
	// 第1引数: LEDを接続しているGPIOの番号を設定して下さい。
	// 第2引数: LEDがLowで点灯する場合は0を、Highで点灯する場合は1を設定して下さい。
	LED := led.New(machine.LED, 1)
	// LEDの点滅で心臓の鼓動を表現
	for {
		LED.Blink(50, 200, 50, 200, 0, 250, 0, 250)
	}
}
