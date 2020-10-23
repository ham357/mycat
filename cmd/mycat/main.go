package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var n bool
	flag.BoolVar(&n, "n", false, "行番号表示")
	flag.Parse()
	f := readFile(n)

	for _, fn := range flag.Args() {
		err := filepath.Walk(".",
			func(path string, info os.FileInfo, err error) error {
				//オプションで指定しているファイル名と同じだったら処理
				if filepath.Base(path) == fn {
					if err := f(path); err != nil {
						fmt.Fprintln(os.Stderr, "os.Open Err:", err)
					}
				}
				return nil
			})
		if err != nil {
			fmt.Fprintln(os.Stderr, "filepath.Walk Err:", err)
		}
	}
}

func readFile(opt bool) func(fn string) error {
	n := 0

	return func(fn string) error {
		f, err := os.Open(fn)
		if err != nil {
			return err
		}
		defer f.Close()

		// 標準入力から読み込む
		scanner := bufio.NewScanner(f)
		// 1行ずつ読み込んで繰り返す
		for scanner.Scan() {
			//-nオプションがTrueの場合、行番号表示
			if opt {
				n++
				fmt.Printf("%v: ",n)
			}
			//1行分を出力する
			fmt.Println(scanner.Text())
		}
		// まとめてエラー処理をする
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner.Err:", err)
		}

		return nil
	}
}
