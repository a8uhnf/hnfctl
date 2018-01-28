package cmds

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	wlog "gopkg.in/dixonwille/wlog.v2"

	"github.com/daviddengcn/go-colortext"
	"github.com/dixonwille/wmenu"
	"github.com/fsnotify/fsnotify"
)

var FileNames []string

const (
	TorrentFileSuffix  = ".torrent"
	DownloadFolderPath = "/home/hanifa/Downloads"
)

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use: "delete",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Delete Command...")
			// getTheInput()
		},
	}
}

func watchDownloadFolder() {
	fmt.Println("Hello Watcher!!!!")
	fmt.Println("----------------------")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("Event", event)
			case err := <-watcher.Errors:
				log.Println("********************", err)
			}
		}
	}()
	err = watcher.Add(DownloadFolderPath)
	if err != nil {
		log.Fatalln(err)
	}
	<-done
}

func getTheInput() {
	menu := wmenu.NewMenu("What is your favorite food?")
	menu.AddColor(wlog.Color{Code: ct.Green}, wlog.Color{Code: ct.Yellow}, wlog.Color{Code: ct.Magenta}, wlog.Color{Code: ct.Yellow})
	menu.Action(func(opts []wmenu.Opt) error {
		fmt.Println(opts[0].Text + " is your favorite food.")
		return nil
	})
	menu.Option("Pizza", nil, true, nil)
	menu.Option("Ice Cream", nil, false, nil)
	menu.Option("Tacos", nil, false, func(o wmenu.Opt) error {
		fmt.Println("Tacos are great")
		return nil
	})
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func visitDowaloads() {
	fmt.Println("Hello Visit!!!")
	err := filepath.Walk(DownloadFolderPath, downloadWalkFunc)
	if err != nil {
		log.Fatalln(err)
	}
}

func downloadWalkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		if strings.HasSuffix(path, TorrentFileSuffix) {
			// fmt.Println("*********************")
			FileNames = append(FileNames, info.Name())
		}
	} else {

	}
	return nil
}
