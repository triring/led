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
	"github.com/triring/led" // githubで公開しているパッケージをインポートする場合
	"machine"
	"time"
	//	"led"  ローカルのディレクトリに置かれたledのパッケージをインポートする場合
)

func main() {

	// オンボードLEDをセットアップ
	LED := led.New(machine.LED)
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
