package main

import (
	"encoding/csv"
	"log"
	"os"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	// O_WRONLY:書き込みモード開く、O_CREATE:無かったらファイル作成
	file, err := os.OpenFile("/tmp/people.csv", os.O_WRONLY|os.O_CREATE, 0600)
	failOnError(err)
	defer file.Close()

	err = file.Truncate(0) // ファイルを空っぽにする（実行2回目以降用)
	failOnError(err)

	writer := csv.NewWriter(file)
	writer.Write([]string{"Alice", "20"})
	writer.Write([]string{"Bob", "21"})
	writer.Write([]string{"Carol", "22"})
	writer.Flush()
}
