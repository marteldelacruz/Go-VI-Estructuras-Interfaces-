package Multimedia

import "fmt"

type Multimedia interface {
	Show()
}

type DigitalContent struct {
	Data []Multimedia
}

type Image struct {
	Title    string
	Format   string
	Channels string
}

type Audio struct {
	Title  string
	Format string
	Lenght float64
}

type Video struct {
	Title  string
	Format string
	Frames float64
}

func (image *Image) Show() {
	fmt.Println("--------------IMAGE----------------------")
	fmt.Println("Title:\t\t ", image.Title)
	fmt.Println("Format:\t\t ", image.Format)
	fmt.Println("Channels:\t ", image.Channels)
}

func (audio *Audio) Show() {
	fmt.Println("---------------AUDIO----------------------")
	fmt.Println("Title:\t ", audio.Title)
	fmt.Println("Format:\t ", audio.Format)
	fmt.Println("Lenght:\t ", audio.Lenght)
}

func (video *Video) Show() {
	fmt.Println("---------------VIDEO-----------------------")
	fmt.Println("Title:\t ", video.Title)
	fmt.Println("Format:\t ", video.Format)
	fmt.Println("Frames:\t ", video.Frames)
}
