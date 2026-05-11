// go get github.com/triring/led
// tinygo build -target=pico2 -size=short -o GenjiFirefly.uf2 .
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
	"fmt"
	"machine"
	"time"

	"github.com/triring/led" // githubで公開しているパッケージをインポートする場合
	//	"led"  ローカルのディレクトリに置かれたledのパッケージをインポートする場合
)

func main() {
	// オンボードLEDを初期化
	// 第1引数: LEDを接続しているGPIOの番号を設定して下さい。
	// 第2引数: LEDがLowで点灯する場合は0を、Highで点灯する場合は1を設定して下さい。
	LED := led.New(machine.LED, 1)

	/* ゲンジボタルの明滅パターンは、生息地域によって異なっている。
	| 地域     | 点灯(秒) | 消灯(秒) |
	|:--------:|:--------:|:--------:|
	| 五島列島 |    1     |    1     |
	| 西日本   |    2     |    1     |
	| 東日本   |    4     |    1     |
	*/
	time.Sleep(time.Millisecond * 3000)
	var i int = 0
	for {
		fmt.Printf("五島列島のゲンジボタル\n")
		for i = 0; i < 15; i++ {
			LED.Blink(1000, 1000)
		}
		fmt.Printf("西日本のゲンジボタル\n")
		for i = 0; i < 5; i++ {
			LED.Blink(2000, 1000)
		}
		fmt.Printf("東日本のゲンジボタル\n")
		for i = 0; i < 3; i++ {
			LED.Blink(4000, 1000)
		}
	}
}
