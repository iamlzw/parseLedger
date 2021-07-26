package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func main()  {
	f, err := os.Create("test.csv")//创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)//创建一个新的写入文件流
	var data [][]string
	for j:= 0 ; j < 200 ; j++ {
		for i := 1 ; i < 5000 ; i++ {
			var first string
			if i < 10 {
				first = "A000"+strconv.Itoa(i)
			}else if i < 100 {
				first = "A00"+strconv.Itoa(i)
			}else if i < 1000 {
				first = "A0"+strconv.Itoa(i)
			}else if i < 5000 {
				first = "A"+strconv.Itoa(i)
			}
			data = append(data,[]string{first,"A"+strconv.Itoa(5000+i)})
		}
	}

	w.WriteAll(data)//写入数据
	w.Flush()
}

