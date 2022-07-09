package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type trackPointerList []*Track

type trackValues []Track

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var valueTracks = []Track{
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
}

var tracks = []*Track{
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func (m trackPointerList) SwapPointer(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func sortPointer(trackPointers trackPointerList) {
	for i := 0; i < len(tracks)-1; i++ {
		for j := i + 1; j < len(tracks); j++ {
			if trackPointers[i].Year > trackPointers[j].Year {
				trackPointers.SwapPointer(i, j)
			}
		}
	}
}

func (m trackValues) swapValue(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func sortValue(trackValue trackValues) {
	for i := 0; i < len(tracks)-1; i++ {
		for j := i + 1; j < len(tracks); j++ {
			if trackValue[i].Year > trackValue[j].Year {
				trackValue.swapValue(i, j)
			}
		}
	}
}

func appliedSortPointer(trackPointers trackPointerList) {
	fmt.Println("Swap pointer")
	fmt.Println("Truoc khi swap")
	for _, value := range trackPointers {
		fmt.Printf("%p  ", value)
	}

	fmt.Printf("\ngia tri dau cua mang: %v", trackPointers[0])

	// địa chỉ của mảng trackPointers
	fmt.Printf("\ndia chi cua mang: %p\n\n", &trackPointers[0])

	fmt.Println("Sau khi swap")

	sortPointer(trackPointers)
	for _, value := range trackPointers {
		fmt.Printf("%p  ", value)
	}
}

func appliedSortValue() {
	fmt.Println("Swap value")
	fmt.Println("Truoc khi swap")
	for i := range valueTracks {
		fmt.Printf("%p  ", &valueTracks[i])
	}
	fmt.Printf("\n%v\n", valueTracks[0])
	fmt.Println("\n\nSau khi swap")

	sortValue(valueTracks)
	for i := range valueTracks {
		fmt.Printf("%p  ", &valueTracks[i])
	}
	fmt.Printf("\n%v\n", valueTracks[0])
	fmt.Println()
}

// kiểm tra mảng con trỏ
// kết luận bên dưới
func main() {

	// áp phần tử đầu tiên hoặc các phần tử mong muốn của track (con trỏ)
	// vào kiểu trackPointerList
	trackPointerList222 := trackPointerList(tracks[:1])
	fmt.Printf("%v", trackPointerList222)

	trackPointers := trackPointerList(tracks)
	appliedSortPointer(trackPointers)
	fmt.Printf("\n------------------------\n")
	//appliedSortValue()

	// tạo mảng con trỏ trỏ vào địa chỉ từng phần tử của mảng
	var trackPointers2 trackPointerList
	for i := range valueTracks {
		trackPointers2 = append(trackPointers2, &valueTracks[i])
	}
	//printTracks(trackPointers2)
	appliedSortPointer(trackPointers2)
	fmt.Println()
	fmt.Println("gia tri cua mang valueTracks[0]: ", valueTracks[0])
	fmt.Printf("dia chi cua valueTracks[0]: %p\n", &valueTracks[0])
	fmt.Println("In ra dia chi cua mang")
	for i := range valueTracks {
		fmt.Printf("%p  ", &valueTracks[i])
	}
}

// mảng là con trỏ (con trỏ đang trỏ vào 1 mảng) thì sẽ swap các con trỏ.
// Tức con trỏ sẽ trỏ đến các vị trí swap mới, còn mảng gốc (mảng đang bị trỏ) thì địa chỉ vẫn giữ nguyên vị trí và ko swap value của nó

// nếu swap là value thì sẽ swap value, còn địa chỉ của nó vẫn giữ nguyên.
// Khuyết điểm là do swap value nên sẽ copy value gây chậm việc chạy

// => cả 2 swap (value và pointer) đều giữ nguyên địa chỉ mảng nhưng khác ở việc sẽ là swap value hay pointer
