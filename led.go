// Package LED driver
package led
import (
//	"fmt"
	"machine"
	"time"
)

// Device wraps the pins of LED.
//
// ledが接続されているPinの情報を保持します。
type Device struct {
	led     machine.Pin
}

// New creates a new led device.
//
// ledドライバを初期化します。
func New(led machine.Pin) Device {
	// Configure sets up the pins.
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.Low()	// LED Lights off
	return Device{led: led}
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
	return d.led.Get()
}

// LED turn On
//
// 	LEDを点灯するメソッドです。
func (d *Device) On() {
	d.led.High()
}

// LED turn Off
//
// 	LEDを消灯するメソッドです。
func (d *Device) Off() {
	d.led.Low()
}

// Toggle LED on and off
//
// LEDのOnとOffを交互に切り替えるメソッドです。
func (d *Device) Toggle() {
	d.led.Set(!d.led.Get())
}
