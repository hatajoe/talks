package main

import "C"

import (
	"fmt"
)

//export Hello
func Hello() {
	// exportする関数に↑のようなコメントを書く
	// //とexportの間にスペースを入れてはいけない
	fmt.Println("Hello!")
}

func init() {
	// ライブラリロード時に呼び出される
	fmt.Println("Loaded!")
}

func main() {
	// パッケージはmainである必要がある
}
