package main

import ("fmt"
		"gopro/week3"
		"gopro/week3/facebook"
		"gopro/week3/twitter"
		"gopro/week3/linkedin"
		"gopro/week3/exporter"
		"os"
		)


func main() {
	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	err := ExportXml(fb, "facebook.xml")
	err := ExportJson(fb, "facebook.json")
	err := exportXml(twtr, "twitter.xml")
	err := exportJson(twtr, "twitter.json")

	if err ! =  nil{
		panic(err)
	}
	//twtr := new(week3.Twitter)
	//lnkdin := new(week3.Linkedin)
	//ScrollFeeds(fb, twtr, lnkdin)
}

func ScrollFeeds(platforms ...SocialMedia) {
	for _, sm := range platforms {
		for _, fb := range sm.Feed() {
			fmt.Println(fb)
		}
		fmt.Println("****************************")
	}
}

func exportText(u week3.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}