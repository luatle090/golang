package main

import (
	"fmt"
	"os"
	"sort"
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

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
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

type ColumnFunc func(x, y *Track) bool

type StableSort struct {
	t       []*Track
	columns []ColumnFunc
}

func (m StableSort) Len() int {
	return len(m.t)
}

func (m StableSort) Less(i, j int) bool {

	if len(m.columns) > 0 {
		fn := m.columns[len(m.columns)-1]
		return fn(m.t[i], m.t[j])
	}
	return false
}
func (m StableSort) Swap(i, j int) {
	m.t[i], m.t[j] = m.t[j], m.t[i]
}

func (m *StableSort) Select(fn ColumnFunc) {
	m.columns = append(m.columns, fn)
}

func sortByTitle(x, y *Track) bool {
	if x.Title != y.Title {
		return x.Title < y.Title
	}
	return false
}

func sortByYear(x, y *Track) bool {
	if x.Year != y.Year {
		return x.Year < y.Year
	}
	return false
}

func sortByArtist(x, y *Track) bool {
	if x.Artist != y.Artist {
		return x.Artist < y.Artist
	}
	return false
}

func sortByLength(x, y *Track) bool {
	if x.Length.Seconds() != y.Length.Seconds() {
		return x.Length.Seconds() < y.Length.Seconds()
	}
	return false
}

func main() {
	fmt.Println()

	trackSorts := StableSort{tracks, nil}
	trackSorts.Select(sortByTitle)
	trackSorts.Select(sortByYear)
	sort.Sort(trackSorts)
	printTracks(tracks)

	//var c []ColumnFunc
	//c = append(c, sortByArtist)

	//fmt.Println(len(c))

	// sort.Stable(customSort{tracks, sortByTitle})
	// sort.Stable(customSort{tracks, sortByYear})

	// //sort.Stable(customSort{tracks, sortByLength})
	//printTracks(tracks)

	// 	people := []struct {
	// 		Name string
	// 		Age  int
	// 	}{
	// 		{"Alice", 25},
	// 		{"Alice", 75},
	// 		{"Alice", 75},
	// 		{"Bob", 75},
	// 		{"Bob", 25},
	// 		{"Colin", 25},
	// 		{"Elizabeth", 25},
	// 		{"Elizabeth", 75},
	// 	}

	// 	// Sort by name, preserving original order
	// 	//sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	// 	//fmt.Println("By name:", people)

	// 	// Sort by age preserving name order
	// 	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	// 	fmt.Println("By age,name:", people)
}
