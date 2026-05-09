// go get github.com/triring/led
// tinygo build -target=pico2 -size=short -o Morse.uf2 .
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
	// モールス符号の基本単位の定義
	unit := 100     // 基準単位
	dit := unit     // 短点（ドット、.）
	dah := unit * 3 // 長点（ダッシュ、-）
	spc := unit     // 記号間

	/*
	   項目	長さ（比率）	説明
	   短点 (・)	1	すべての基準となる時間の長さ
	   長点 (－)	3	短点3個分の長さ
	   記号間の間隔	1	同一文字内の「・」や「－」の間の無音時間
	   文字間の間隔	3	文字と文字（例：AとBの間）の無音時間
	   単語間の間隔	7	単語と単語の間の無音時間
	*/

	// オンボードLEDを初期化
	LED := led.New(machine.LED)
	time.Sleep(time.Millisecond * 3000)
	// LEDの点滅でモールス信号を送信
	for {
		fmt.Printf("T\t-\n")
		LED.Blink(dah, spc*3)
		fmt.Printf("I\t..\n")
		LED.Blink(dit, spc, dit, spc*3)
		fmt.Printf("N\t-.\n")
		LED.Blink(dah, spc, dit, spc*3)
		fmt.Printf("Y\t-.--\n")
		LED.Blink(dah, spc, dit, spc, dah, spc, dah, spc*3)
		fmt.Printf("G\t--.\n")
		LED.Blink(dah, spc, dah, spc, dit, spc*3)
		fmt.Printf("O\t---\n\n")
		LED.Blink(dah, spc, dah, spc, dah, spc*7)
	}
}
