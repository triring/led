// Package LED は、LED用ドライバです。
package led

import (
	"machine"
	"time"
)

const (
	ActiveLow  = iota // Low（0）で点灯する場合
	ActiveHigh        // High（1）で点灯する場合
)

// Device wraps the pins of LED.
//
// ledが接続されているPinの情報を保持します。
type Device struct {
	led     machine.Pin // LEDが接続されているPin情報を保持
	turnOn  bool        // LEDの点灯条件を定義
	turnOff bool        // LEDの消灯条件を定義
}

// New creates a new led device.
//
// ledドライバを初期化します。
// led: LEDを接続するGPIOの番号を設定して下さい。
// OutputType: LEDがLowで点灯する場合は0を、Highで点灯する場合は1を設定して下さい。
func New(led machine.Pin, OutputType int) Device {
	var dev Device
	// Configure sets up the pins.
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	switch OutputType {
	case ActiveLow: // Lowで点灯する場合の設定：シンク駆動・カソード側駆動
		led.Set(true)
		dev = Device{led: led, turnOn: false, turnOff: true}
	case ActiveHigh: // Highで点灯する場合の設定：ソース駆動・アノード側駆動
		led.Set(false)
		dev = Device{led: led, turnOn: true, turnOff: false}
	}
	return dev
}

// Make the LED blink.
//
// このメソッドは、LEDを点滅させます。
// 引数は、任意の数の引数を受け取れる可変個引数になっています。
// 奇数番目は点灯時間、偶数番目は消灯時間なので、長さを組み合わせて並べると様々は発光パターンで点滅させることができます。
// 単位はミリ秒です。
func (d *Device) Blink(DurationTime ...int) {
	for _, dt := range DurationTime {
		d.led.Set(!d.led.Get())
		time.Sleep(time.Millisecond * time.Duration(dt))
	}
}

// Pin mode setting
//
// ピンのモードを設定します。
func pinMode(pin machine.Pin, mode bool) {
	if mode {
		//	fmt.Println("machine.PinInput")
		pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	} else {
		//	fmt.Println("machine.PinOutput")
		pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
}

// Check the status of the leds.
//
// LEDを点灯状態を確認するメソッドです。
func (d *Device) Status() bool {
	if d.turnOn == d.led.Get() {
		return d.turnOn
	} else {

		return d.turnOff
	}
}

// LED turn On
//
//	LEDを点灯するメソッドです。
func (d *Device) On() {
	d.led.Set(d.turnOn)
}

// LED turn Off
//
//	LEDを消灯するメソッドです。
func (d *Device) Off() {
	d.led.Set(d.turnOff)
}

// Toggle LED on and off
//
// LEDのOnとOffを交互に切り替えるメソッドです。
func (d *Device) Toggle() {
	d.led.Set(!d.led.Get())
}
