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

	// ?????a ch??? c???a m???ng trackPointers
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

// ki???m tra m???ng con tr???
// k???t lu???n b??n d?????i
func main() {

	// ??p ph???n t??? ?????u ti??n ho???c c??c ph???n t??? mong mu???n c???a track (con tr???)
	// v??o ki???u trackPointerList
	trackPointerList222 := trackPointerList(tracks[:1])
	fmt.Printf("%v", trackPointerList222)

	trackPointers := trackPointerList(tracks)
	appliedSortPointer(trackPointers)
	fmt.Printf("\n------------------------\n")
	//appliedSortValue()

	// t???o m???ng con tr??? tr??? v??o ?????a ch??? t???ng ph???n t??? c???a m???ng
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

// m???ng l?? con tr??? (con tr??? ??ang tr??? v??o 1 m???ng) th?? s??? swap c??c con tr???.
// T???c con tr??? s??? tr??? ?????n c??c v??? tr?? swap m???i, c??n m???ng g???c (m???ng ??ang b??? tr???) th?? ?????a ch??? v???n gi??? nguy??n v??? tr?? v?? ko swap value c???a n??

// n???u swap l?? value th?? s??? swap value, c??n ?????a ch??? c???a n?? v???n gi??? nguy??n.
// Khuy???t ??i???m l?? do swap value n??n s??? copy value g??y ch???m vi???c ch???y

// => c??? 2 swap (value v?? pointer) ?????u gi??? nguy??n ?????a ch??? m???ng nh??ng kh??c ??? vi???c s??? l?? swap value hay pointer
