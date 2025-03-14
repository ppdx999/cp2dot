package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// GraphFormatter インターフェース
type GraphFormatter interface {
	Header() string
	EdgeFormat(u, v string) string
	Footer() string
}

// DirectedGraphFormatter (有向グラフ)
type DirectedGraphFormatter struct{}

func (d DirectedGraphFormatter) Header() string { return "digraph G {" }
func (d DirectedGraphFormatter) EdgeFormat(u, v string) string {
	return fmt.Sprintf("    %s -> %s;", u, v)
}
func (d DirectedGraphFormatter) Footer() string { return "}" }

// UndirectedGraphFormatter (無向グラフ)
type UndirectedGraphFormatter struct{}

func (u UndirectedGraphFormatter) Header() string { return "graph G {" }
func (u UndirectedGraphFormatter) EdgeFormat(u1, v string) string {
	return fmt.Sprintf("    %s -- %s;", u1, v)
}
func (u UndirectedGraphFormatter) Footer() string { return "}" }

// Factory 関数: グラフの種類に応じた `GraphFormatter` を返す
func NewGraphFormatter(directed bool) GraphFormatter {
	if directed {
		return DirectedGraphFormatter{}
	}
	return UndirectedGraphFormatter{}
}

// グラフを読み取り、DOT 形式で出力する処理
func processGraph(formatter GraphFormatter) {
	fmt.Println(formatter.Header())

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var u, v string
		fmt.Sscanf(scanner.Text(), "%s %s", &u, &v)
		fmt.Println(formatter.EdgeFormat(u, v))
	}

	fmt.Println(formatter.Footer())
}

func main() {
	// CLI オプションの解析
	directed := flag.Bool("d", false, "directed graph")
	flag.Parse()

	// Factory パターンで適切な `GraphFormatter` を生成
	formatter := NewGraphFormatter(*directed)

	// グラフを処理（標準入力から読み取り、DOT 出力）
	processGraph(formatter)
}
