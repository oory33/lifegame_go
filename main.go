package main

import (
	"fmt"
	"lifegame/system"
	"math/rand"
	"time"
)

type table struct { //盤面の構造体の定義
	fields [][]bool
}

func (t *table) initialize(n int) { //盤面の初期化
	for i := 0; i < n; i++ {
		t.fields = append(t.fields, make([]bool, n))
		for j := 0; j < n; j++ {
			t.fields[i][j] = rand.Intn(2) == 1
		}
	}
}

func (t *table) init_galaxy() {
	g := [...][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	for i := 0; i < len(g); i++ {
		t.fields = append(t.fields, make([]bool, len(g)))
		for j := 0; j < len(g[0]); j++ {
			t.fields[i][j] = itob(g[i][j])
		}
	}
}

func (t *table) height() int { //高さを取得するメソッド
	return len(t.fields)
}

func (t *table) width() int { //幅を取得するメソッド
	return len(t.fields[0])
}

func (t *table) issame(newt *table) bool { //盤面が同じかどうかを判定するメソッド
	value := true
	for i := 0; i < t.height(); i++ {
		for j := 0; j < t.width(); j++ {
			if t.fields[i][j] != newt.fields[i][j] {
				value = false
			}
		}
	}
	return value
}

func (t *table) replace(x, y int, val bool) { //盤面の要素を書き換えるメソッド
	t.fields[x][y] = val
}

func btoi(b bool) int { //boolを01に変換する関数
	if b {
		return 1
	} else {
		return 0
	}
}

func itob(i int) bool {
	if i != 0 {
		return true
	} else {
		return false
	}
}

func at_cell(t *table, x, y int) bool { //引数の座標が盤面内かどうかの判定
	if x < 0 || y < 0 || x >= t.height() || y >= t.width() {
		return false
	} else {
		return t.fields[x][y]
	}
}

func count_surround(t *table, x, y int) int { //周囲の生きているセルを数える
	count := 0
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if at_cell(t, i, j) {
				count += 1
			}
		}
	}
	return count - btoi(at_cell(t, x, y))
}

func dead_or_alive(t *table, x, y int) bool {
	value := false
	switch at_cell(t, x, y) {
	case true:
		if count_surround(t, x, y) == 2 || count_surround(t, x, y) == 3 {
			value = true
		}
	case false:
		if count_surround(t, x, y) == 3 {
			value = true
		}
	}
	return value
}

func print_table(t *table) {
	system.Clear()
	for i := 0; i < t.width()+2; i++ {
		fmt.Print("-")
	}
	fmt.Println()
	for i := 0; i < t.height(); i++ {
		fmt.Print("|")
		for j := 0; j < t.width(); j++ {
			if t.fields[i][j] {
				fmt.Print("■")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
	for i := 0; i < t.width()+2; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func next_generation(t *table) table {
	var new_gen table
	new_gen.initialize(t.height())
	for i := 0; i < t.height(); i++ {
		for j := 0; j < t.width(); j++ {
			new_gen.replace(j, i, dead_or_alive(t, j, i))
		}
	}
	return new_gen
}

func main() {
	t := &table{}
	// t.initialize(20)
	t.init_galaxy()

	running := true

	for running {
		print_table(t)
		new_t := next_generation(t)
		new_p := &new_t //全て参照渡しなので引数にはポインタ変数を使う
		if new_p.issame(t) {
			running = false
		} else {
			t = new_p
		}
		time.Sleep(200 * time.Millisecond)
	}
}
