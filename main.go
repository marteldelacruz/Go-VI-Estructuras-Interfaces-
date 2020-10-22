package main

import (
	"fmt"

	Multimedia "./Multimedia"
)

const (
	IMAGE string = "image"
	AUDIO string = "audio"
	VIDEO string = "video"
)

func addContent(media *Multimedia.WebContent, _type string) {
	var title, format, channels string
	var lenght, frames float64

	// Same info
	fmt.Println("------------------------------------")
	fmt.Print("Title: ")
	fmt.Scanln(&title)
	fmt.Print("Format: ")
	fmt.Scanln(&format)

	// Different info
	switch _type {
	case IMAGE:
		fmt.Print("Channels: ")
		fmt.Scanln(&channels)

		img := Multimedia.Image{Title: title, Format: format, Channels: channels}
		media.Data = append(media.Data, &img)

		break
	case AUDIO:
		fmt.Print("Lenght: ")
		fmt.Scanln(&lenght)

		audio := Multimedia.Audio{Title: title, Format: format, Lenght: lenght}
		media.Data = append(media.Data, &audio)

		break
	case VIDEO:
		fmt.Print("Frames: ")
		fmt.Scanln(&frames)

		vid := Multimedia.Video{Title: title, Format: format, Frames: frames}
		media.Data = append(media.Data, &vid)

		break
	}

}

func showContent(media *Multimedia.WebContent) {
	for _, c := range media.Data {
		c.Show()
	}
	fmt.Print("Press enter to continue...")
	fmt.Scanln()
}

func main() {
	var opt string
	media := Multimedia.WebContent{}

	for opt != "0" {
		fmt.Println("\n------------------------------")
		fmt.Println("1-. Add image")
		fmt.Println("2-. Add audio")
		fmt.Println("3-. Add video")
		fmt.Println("4-. Show content")
		fmt.Println("0-. Exit")
		fmt.Print("\nSelect an option: ")
		fmt.Scanln(&opt)

		switch opt {
		case "1":
			addContent(&media, IMAGE)
			break
		case "2":
			addContent(&media, AUDIO)
			break
		case "3":
			addContent(&media, VIDEO)
			break
		case "4":
			showContent(&media)
			break
		}
	}
}
