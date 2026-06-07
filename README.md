# led
このリポジトリは、一般的なledをtinygoで制御するために実装したドライバです。  
**「LEDは、とてもメジャーなデバイスなのに、#Tinygo の公式にドライバが登録されていなかったから自作した。」** という訳ではありません。  
tinygoのコードをパッケージとして集約する方法やGithubへの登録方法を学ぶために、テスト用に作成したものです。  
実用的な意味も価値もありません。  

## 使用方法

以下のコマンドで、このリポジトリの内容をローカルにコピーして下さい。
```bash
git clone https://github.com/triring/led.git
```

コピーされたledディレクトリ内のファイル構成
```bash
led
|   .gitignore
|   go.mod
|   led.go
|   LICENSE
|   README.md
|
\---examples
    +---Blink
    |       main.go
    |
    +---GenjiFirefly
    |       main.go
    |
    +---HeartBeats
    |       main.go
    |
    \---Morse
            main.go
```

コピーしたディレクトリ内に、examplesディレクトリがあります。

1. PCとターゲットボードを接続して下さい。
2. 実行したいコードのあるディレクトリ内に移動して下さい。
3. 以下のコマンドで、コンパイル&実行ファイルの転送を行って下さい。  
(-targetオプションは、使用するマイコンボードに合わせて修正して下さい。)

```bash
tinygo flash -target=pico2 -size=short -monitor .
```

## 解説

このドライバを使うと、以下のようなコードで簡単にLチカができます。

1. "github.com/triring/led"をインポートする。
2. ledを初期化する。
3. Blink()メソッドを呼び出す

```go
package main
import (
	"machine"
	"github.com/triring/led"
)
func main() {
	LED := led.New(machine.LED) // オンボードLEDを初期化
	for {   // 無限ループで点滅
		LED.Blink(500)
	}
}
```

以下は、tinygoのチュートリアルに書かれているLチカのコードです。
このコードと比較すると、LEDドライバを使うことで、記述がとてもシンプルになることがわかると思います。

[Blinking LED](https://tinygo.org/docs/tutorials/blinky/)

```go
package main

import (
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.Low()
        time.Sleep(time.Millisecond * 500)

        led.High()
        time.Sleep(time.Millisecond * 500)
    }
}
```

## メソッドの使い方

### Blink()メソッド
このメソッドの引数は、任意の数の引数を受け取れる可変個引数になっています。
点灯時間、消灯時間を1組として、これを複数個ならべていくと、様々な点滅パターンを作成できます。

例1 モールス信号 G
```go
LED.Blink(300, 100, 300, 100, 100)
```

例2 LEDの点滅で心臓の鼓動を表現
```go
LED.Blink(50, 200, 50, 200, 0, 250, 0, 250)
```

## このパッケージのドキュメント

[package ledのドキュメント](https://pkg.go.dev/github.com/triring/led)
