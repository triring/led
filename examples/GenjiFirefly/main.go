// go get github.com/triring/led
// tinygo build -target=pico2 -size=short -o GenjiFirefly.uf2 ./main.go
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
	//	"fmt"
	"machine"

	"github.com/triring/led" // githubで公開しているパッケージをインポートする場合の記述
	//	"led"  ローカルのディレクトリに置かれたledのパッケージをインポートする場合の記述
)

func main() {

	// オンボードLEDをセットアップ
	OnboardLED := led.New(machine.LED)
	// ゲンジボタルの明滅パターン
	// 地域によって、発光パターンが異なる。
	/*
		| 地域     | 点灯(秒) | 消灯(秒) |
		|:--------:|:--------:|:--------:|
		| 五島列島 |    1     |    1     |
		| 西日本   |    2     |    1     |
		| 東日本   |    4     |    1     |
	*/
	var i int = 0
	for {
		for i = 0; i < 15; i++ {
			OnboardLED.Blink(1000, 1000) //	五島列島のゲンジボタル
		}
		for i = 0; i < 5; i++ {
			OnboardLED.Blink(2000, 1000) //	西日本のゲンジボタル
		}
		for i = 0; i < 3; i++ {
			OnboardLED.Blink(4000, 1000) //	東日本のゲンジボタル
		}
	}
}
