package challenge

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type ChallangeData struct {
	Days []string
	Sum  []int
}

func GetChallangeData(wikiContent string) *ChallangeData {
	body := strings.NewReader(wikiContent)

	z := html.NewTokenizer(body)
	table := [][]string{}
	row := []string{}
	tt := z.Next()
	for tt != html.ErrorToken {
		tok := z.Token().Data
		//fmt.Printf("ForToken %v Type %v\n", tok, tt.String())
		if tt == html.StartTagToken {
			//fmt.Printf("StartToken %v\n", tok)
			if tok == "tr" {
				if len(row) > 0 {
					table = append(table, row)
					row = []string{}
				}
				//fmt.Printf("Table now %v\n", table)
			}
			found := false
			text := ""
			if tok == "td" || tok == "th" {
				//fmt.Printf("found token %v\n",tok)
				for found == false {
					inner := z.Next()
					data := z.Token().Data
					//fmt.Printf("next token %v data -%v-\n",inner, data)
					if inner == html.EndTagToken && (data == "td" || data == "th" || data == "ac:link") {
						//fmt.Printf("found end token %v data -%v-\n",inner, data)
						//		inner = z.Next()
						break
					}
					if inner == html.TextToken {
						text = strings.TrimSpace((string)(data))
						//fmt.Printf("set text %v\n", text)
						break
					}
				}
				//fmt.Printf("append %v\n", text)
				row = append(row, text)
			}

		}
		tt = z.Next()
	}
	//	fmt.Printf("Table now %v", table)

	if len(row) > 0 {
		table = append(table, row)
	}
	fmt.Printf("Table now %v\n\n", table)
	var summe []int
	summe = make([]int, 20)
	for y := 0; y < len(table); y++ {

		for x := 0; x < len(table[0]); x++ {
			//			fmt.Printf("len table y %v x %d\n", len(table[y]), x)
			if len(table[y]) == x {
				break
			}
			fmt.Printf("%10.10s ", table[y][x])
			i, _ := strconv.Atoi(table[y][0])
			if strings.TrimSpace(table[y][0]) == fmt.Sprintf("%d", i) && x > 1 && len(table[y][x]) > 0 {
				//w, _ := strconv.Atoi(table[y][x])
				summe[x] = summe[x] + 1
			}
			fmt.Printf("%d \t", summe[x])
		}
		fmt.Println()
	}
	s := 0
	var res ChallangeData
	for i := 2; i < len(table[0]); i++ {
		if i > len(summe) {
			s = 0
		} else {
			s = summe[i]
		}
		res.Days = append(res.Days, table[0][i])
		res.Sum = append(res.Sum, s)
		fmt.Printf("%s\t%d\n", table[0][i], s)
	}
	fmt.Println(summe)
	return &res
}
