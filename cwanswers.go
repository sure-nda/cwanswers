package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	//"runtime"
	"unicode"
	//"unicode/utf8"
	"strconv"
	"strings"
)

type puzzle struct {
	cells         [][]string
	acrossAnswers []string
	downAnswers   []string
}

func readpuzzle(inputfile string) *puzzle {
	content, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Println("File I/O Error")
		return nil
	}

	splitcontent := strings.Split(string(content), "===\n")

	lines := strings.Split(splitcontent[0], "\n")
	Xsize := len(lines)
	var Ysize int = 0
	var puzz puzzle
	puzz.cells = make([][]string, Xsize)
	for lineno := 0; lineno < Xsize; lineno++ {
		if lines[lineno] == "" {
			Xsize--
			break
		}
		if strings.ContainsAny(lines[lineno], "|") {
			puzz.cells[lineno] = strings.Split(lines[lineno], "|")
		} else {
			puzz.cells[lineno] = strings.Split(lines[lineno], "\t")
		}
		tmpysize := len(puzz.cells[lineno])

		if Ysize == 0 {
			Ysize = tmpysize
		}
		if tmpysize != Ysize {
			fmt.Printf("mismatchedColumns,row=%d,Ysize=%d,tmpysize=%d\n", lineno, Ysize, tmpysize)
		}
	}
	return &puzz
	//fmt.Printf("Ysize=%d\n", Ysize)

	//fmt.Println(cells)
}

func cwanswers(inputfile string, isHtml bool) {

	content, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Println("File I/O Error")
		return
	}
	var w *bufio.Writer
	var f *os.File

	splitcontent := strings.Split(string(content), "===\n")

	lines := strings.Split(splitcontent[0], "\n")
	Xsize := len(lines)
	if isHtml {
		outputfile := strings.Replace(inputfile, ".txt", ".html", 1)
		f, err = os.OpenFile(outputfile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("File I/O Error")
			return
		}
		w = bufio.NewWriter(f)
		fmt.Fprintf(w, "<style type=\"text/css\" media=\"all\">\n\t@import url( http://eemaata.com/em/wp-content/themes/hive/crosswordprint.css );\n </style>\n")

	}
	//fmt.Printf("Xsize=%d\n", Xsize)
	var Ysize int = 0
	cells := make([][]string, Xsize)
	var acrossAnswers [100]string
	var downAnswers [100]string
	var id, clueidstr string
	var clueid int = 0
	var newclue bool
	for lineno := 0; lineno < Xsize; lineno++ {
		if lines[lineno] == "" {
			Xsize--
			break
		}
		if strings.ContainsAny(lines[lineno], "|") {
			cells[lineno] = strings.Split(lines[lineno], "|")
		} else {
			cells[lineno] = strings.Split(lines[lineno], "\t")
		}
		tmpysize := len(cells[lineno])

		if Ysize == 0 {
			Ysize = tmpysize
		}
		if tmpysize != Ysize {
			fmt.Printf("mismatchedColumns,row=%d,Ysize=%d,tmpysize=%d\n", lineno, Ysize, tmpysize)
		}
	}
	//fmt.Printf("Ysize=%d\n", Ysize)

	//fmt.Println(cells)

	if isHtml {
		fmt.Fprintln(w, "<table id=\"table15367\" class=\"crossword\"><tbody>")
	}
	for lineno := 0; lineno < Xsize; lineno++ {
		if isHtml {
			fmt.Fprintln(w, "\t<tr>")
		}
		for i := range cells[lineno] {
			newclue = false
			if cells[lineno][i] != "#" {
				if ((i == 0 || cells[lineno][i-1] == "#") && (i+1) < Ysize && cells[lineno][i+1] != "#") ||
					((i == 0 || cells[lineno][i-1] == "#") && (lineno == 0 || cells[lineno-1][i] == "#") && !((lineno+1) < Xsize && cells[lineno+1][i] != "#")) {
					newclue = true
					coly := i
					//fmt.Printf("#### lineno=%d coly=%d Ysize=%d Xsize=%d\n", lineno, coly, Xsize, Ysize);
					//In a loop, read the complete answer
					for coly < Ysize && cells[lineno][coly] != "#" {
						acrossAnswers[clueid+1] += cells[lineno][coly]
						coly++
					}
				}
				if (lineno == 0 || cells[lineno-1][i] == "#") && (lineno+1) < Xsize && cells[lineno+1][i] != "#" {
					newclue = true
					colx := lineno
					//fmt.Printf("#### lineno=%d colx=%d Ysize=%d Xsize=%d i=%d\n", lineno, colx, Xsize, Ysize,i);
					//In a loop, read the complete answer
					for colx < Xsize && cells[colx][i] != "#" {
						//fmt.Printf("#### lineno=%d colx=%d Ysize=%d Xsize=%d i=%d\n", lineno, colx, Xsize, Ysize,i);
						downAnswers[clueid+1] += cells[colx][i]
						colx++
					}
				}

				if newclue {
					clueid++
					clueidstr = strconv.Itoa(clueid)
					id = " id=d" + clueidstr
				} else {
					id = ""
					clueidstr = ""
				}
				if isHtml {
					fmt.Fprintf(w, "\t\t<td%s><input type=\"text\" value=\"%s\"/></td>\n", id, cells[lineno][i])
				} else {
					fmt.Printf("\t%s%s", clueidstr, cells[lineno][i])
				}
			} else {
				if isHtml {
					fmt.Fprintf(w, "\t\t<td></td>\n")
				} else {
					fmt.Printf("\t#")
				}
			}
		}
		if isHtml {
			fmt.Fprintln(w, "\t</tr>")
		} else {
			fmt.Println("")
		}
	}
	if isHtml {
		fmt.Fprintln(w, "<tbody></table>")
	}

	if len(splitcontent) > 2 {
		if isHtml {
			fmt.Fprintln(w, "<h3>అడ్డం</h3> <ol id=\"hor\" class=\"pointer\">")
		}
		acrossclues := strings.Split(splitcontent[1], "\n")
		var clueno = 0
		for i := 1; i <= clueid; i++ {
			if acrossAnswers[i] != "" {
				if isHtml {
					if acrossclues[clueno][0] >= '0' && acrossclues[clueno][0] <= '9' {
						indexByte := strings.IndexFunc(acrossclues[clueno], func(c rune) bool { return (unicode.IsSpace(c)) })
						fmt.Fprintf(w, "<li value=\"%d\"><strong>%s</strong>\n సమాధానం: %s</li>\n", i, acrossclues[clueno][indexByte:], acrossAnswers[i])
					} else {
						fmt.Fprintf(w, "<li value=\"%d\"><strong>%s</strong>\n సమాధానం: %s</li>\n", i, acrossclues[clueno], acrossAnswers[i])
					}
				} else {
					fmt.Println(i, acrossclues[clueno], "\nసమాధానం: ", acrossAnswers[i])
				}
				clueno++
			}
		}
		if isHtml {
			fmt.Fprintln(w, "</ol>")
			fmt.Fprintln(w, "<h3>నిలువు</h3> <ol id=\"ver\" class=\"pointer\">")
		}

		downclues := strings.Split(splitcontent[2], "\n")
		clueno = 0
		for i := 1; i <= clueid; i++ {
			if downAnswers[i] != "" {
				if isHtml {
					if downclues[clueno][0] >= '0' && downclues[clueno][0] <= '9' {
						indexByte := strings.IndexFunc(downclues[clueno], func(c rune) bool { return (unicode.IsSpace(c)) })
						fmt.Fprintf(w, "<li value=\"%d\"><strong>%s</strong>\n సమాధానం: %s</li>\n", i, downclues[clueno][indexByte:], downAnswers[i])
					} else {
						fmt.Fprintf(w, "<li value=\"%d\"><strong>%s</strong>\n సమాధానం: %s</li>\n", i, downclues[clueno], downAnswers[i])
					}
				} else {
					fmt.Println(i, downclues[clueno], "\nసమాధానం: ", downAnswers[i])
				}
				clueno++
			}
		}
		if isHtml {
			fmt.Fprintln(w, "</ol>")
			w.Flush()
			f.Close()
		}
	}
	//outfile.Sync()
}

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Error: Missing parameter. Usage: %s [--html] <input-unicode-file>\n", os.Args[0])
		return
	}

	var isHtml bool = false
	var inputfile string
	if len(os.Args) == 3 {
		if os.Args[1] == "--html" {
			isHtml = true
			inputfile = os.Args[2]
		}
	}
	if isHtml == false {
		inputfile = os.Args[1]
	}
	cwanswers(inputfile, isHtml)
}
