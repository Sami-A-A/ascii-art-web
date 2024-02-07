package asciiartutil

type AsciiArtData struct {
	Text string
	Banner string
	Color string
	Alignment string
	Export string
}

func PrintAsciiArt(asciiArt AsciiArtData) string {
 return asciiArt.Text
}