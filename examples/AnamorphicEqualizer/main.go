package main

// tinygo flash -target=pico -size=short -monitor .
// tinygo build -o GoroutineBlink.uf2 -target=pico -size short .
// Maker Nano RP2040
// https://docs.google.com/document/d/15IMxlESQE43sP7brZpqlfHVTAdX1N_XenzpvfP8cT_8/edit?tab=t.0

import (
	"fmt"
	"machine"
	"time"

	// "led" // ローカルのディレクトリに置かれたledのパッケージをインポートする場合
	"github.com/triring/led" // githubで公開しているパッケージをインポートする場合
)

// Message は goroutine 間でやり取りするデータ構造
type Message struct {
	Direction int // 移動方向: 1 = 順方向, -1 = 逆方向
}

// TaskNode は双方向リンクリストのノードを表す構造体
type TaskNode struct {
	Id     int
	Led    led.Device
	MyChan chan Message // 自分の指示待ちチャネル
	Next   *TaskNode    // 次のタスクへのポインタ
	Prev   *TaskNode    // 前のタスクへのポインタ
}

// LEDの初期状態
// LEDの実装方法により、GPIO ピンの出力での点灯状態が異なってきます。
// * High()の時に点灯する
// * Low()の時に点灯する
// 全LEDの初期状態は、全て消灯状態にしておく必要があるので、LEDが消灯状態になるように、以下の定数を0か1を設定して下さい。
const (
	InitialState int = 1
)

// 点滅パターンを設定する。
// 各ノードは、このタイミングで1度だけ点滅する。
// 実機での点滅状態を確認しながら調整すること。
/* for Maker Nano RP2040 LED array 11 LEDs */
const (
	LightOnTime  time.Duration = 70 // 点灯時間
	LightOffTime time.Duration = 21 // 消灯時間
)

/* for Maker PI RP2040 LED array 13 LEDs
const (
	LightOnTime  time.Duration = 60 // 点灯時間
	LightOffTime time.Duration = 17 // 消灯時間
)
*/

// 使用するGPIOピンのリスト
// LEDの配置順にGPIOを書き込んでいく。
// リストを変更するだけでノードの増減が可能
/* for Maker Nano RP2040 LED array 11 LEDs */
var pins = [...]machine.Pin{
	machine.GPIO2, machine.GPIO3, machine.GPIO4, machine.GPIO5,
	machine.GPIO6, machine.GPIO7, machine.GPIO8, machine.GPIO9,
	machine.GPIO17, machine.GPIO19, machine.GPIO16,
}

/* for Maker PI RP2040 LED array 13 LEDs
var pins = [...]machine.Pin{
	machine.GPIO0, machine.GPIO1, machine.GPIO2, machine.GPIO3,
	machine.GPIO4, machine.GPIO5, machine.GPIO6, machine.GPIO7,
	machine.GPIO16, machine.GPIO17, machine.GPIO26, machine.GPIO27, machine.GPIO28,
}
*/
// run は各タスク（goroutine）のメインループ
func (tn *TaskNode) run() {
	// 無限ループで自分宛ての指示を待ち続ける
	for msg := range tn.MyChan {
		fmt.Printf("[%2d] Blink ... (方向: %2d)\n", tn.Id, msg.Direction)

		// 1. 決められた処理を行う (点灯・消灯)
		tn.Led.Toggle()
		time.Sleep(LightOnTime * time.Millisecond)
		tn.Led.Toggle()
		time.Sleep(LightOffTime * time.Millisecond)

		// 2. 次にメッセージを送る相手と方向を決定する
		var target *TaskNode
		nextDir := msg.Direction

		if msg.Direction == 1 {
			// 順方向の進路
			target = tn.Next
			if target == nil { // 端に到達したら折り返す
				target = tn.Prev
				nextDir = -1
				fmt.Printf("[%2d] 端に達したため、逆方向へ折り返します。\n", tn.Id)
			}
		} else {
			// 逆方向の進路
			target = tn.Prev
			if target == nil { // 端に到達したら折り返す
				target = tn.Next
				nextDir = 1
				fmt.Printf("[%2d] 端に達したため、順方向へ折り返します。\n", tn.Id)
			}
		}

		// 3. 次のタスクにメッセージを送信する
		if target != nil {
			target.MyChan <- Message{Direction: nextDir}
		}
	}
}

func main() {
	var firstNode, lastNode *TaskNode
	time.Sleep(2 * time.Second) // シリアル接続等の初期化待ち

	// --- 1. リンクリストの構築と goroutine の起動 ---
	for i, pin := range pins {
		node := &TaskNode{
			Id:     i,
			Led:    led.New(pin, InitialState), // LEDが消灯状態になるように、InitialStateを0か1に設定しておくこと。
			MyChan: make(chan Message),
		}
		// led.New(pin, ?)
		// オンボードLEDを初期化
		// 第1引数: LEDを接続しているGPIOの番号を設定する。
		// 第2引数: LEDがLowで点灯する場合は0を、Highで点灯する場合は1を設定する。

		// 双方向リンクリストを構築
		if firstNode == nil {
			firstNode = node
		} else {
			lastNode.Next = node // 前のノードの次を現在のノードに
			node.Prev = lastNode // 現在のノードの前を前のノードに
		}
		lastNode = node

		// タスク用の goroutine を即座に起動
		go node.run()
	}

	// --- 2. 最初のトリガーを送信 ---
	fmt.Println("[Main] 端から往復無限ループを開始します。")
	firstNode.MyChan <- Message{Direction: 1}

	// メインルーチンが終了しないようにブロック
	select {}
}
