module main

go 1.26.2

require (
	led v0.0.0	// パッケージ名とバージョン
)

replace (
	// パッケージ名 => project から led までのパス
	led => ../../../../led
)
