package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

const fileName string = "./commandToTemplate.json"

func main() {
	fileUpdated := make(chan bool)
	go checkFileUpdated(fileUpdated)

	for {
		select {
		case <-fileUpdated:
			//load file, read it's content, parse JSON,
			//and return map with parsed values
			cmdToTplMap, err := readJSONFileToMap(fileName)
			if err != nil {
				log.Println("file to JSON to map problem : ", err)
			}

			if cmdToTplMap != nil {
				fmt.Println("\nContent of the map unmarshaled from fileContent :")
				for key, value := range cmdToTplMap {
					fmt.Println("key = ", key, "value = ", value)
				}
			}
		}
	}

}

//JSDONFileToMap Load file, read it's content, parse JSON,
//and return map with parsed values.
func readJSONFileToMap(fileName string) (map[string]string, error) {
	cmdToTplMap := make(map[string]string)

	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("Failed to open file %v\n", err)
		return nil, err
	}

	fileContent, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("Failed reading file %v\n", err)
		return nil, err
	}

	fmt.Println("Content read from file : \n", string(fileContent))

	err = json.Unmarshal(fileContent, &cmdToTplMap)
	if err != nil {
		log.Printf("Failed unmarshaling %v\n", err)
		return nil, err
	}

	return cmdToTplMap, nil
}

func checkFileUpdated(fileUpdated chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Failed fsnotify.NewWatcher")
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		//Give a true value to updated so it reads the file the first time.
		fileUpdated <- true
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					//testing with an update chan to get updates
					fileUpdated <- true
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fileName)
	if err != nil {
		log.Fatal(err)
	}
	<-done

	return
}
