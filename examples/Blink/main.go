// tinygo build -target=pico2 -size=short -o blink.uf2 ./main.go
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

	"fmt"
	"machine"

	"led" // ローカルに置いたledの自作パッケージ
)

func main() {
	// オンボードLEDを初期化
	OnboardLED := led.New(machine.LED)
	// 無限ループで点滅
	for {
		OnboardLED.Blink(1000)
		fmt.Println(OnboardLED.Status())
	}
}
