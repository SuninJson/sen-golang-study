package _1_statistical_directory

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func WalkDir(dirs ...string) string {
	if len(dirs) == 0 {
		dirs = []string{"."}
	}

	filesizeCh := make(chan int64, 1)

	wg := &sync.WaitGroup{}
	for _, dir := range dirs {
		wg.Add(1)
		go walkDir(dir, filesizeCh, wg)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(filesizeCh)
	}(wg)

	var fileNum, sizeTotal int64
	for filesize := range filesizeCh {
		fileNum++
		sizeTotal += filesize
	}

	return fmt.Sprintf("%d files %.2f MB\n", fileNum, float64(sizeTotal)/1e6)
}
func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, fileInfo := range fileInfos(dir) {
		if fileInfo.IsDir() {
			subDir := filepath.Join(dir, fileInfo.Name())
			wg.Add(1)
			go walkDir(subDir, fileSizes, wg)
		} else {
			fileSizes <- fileInfo.Size()
		}
	}
}
func fileInfos(dir string) []fs.FileInfo {
	//log.Println("The dir witch need to get its file information:", dir)
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("walkdir: %v\n", err)
		return []fs.FileInfo{}
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		infos = append(infos, info)
		//log.Println(dir, " info:", info.Name())
	}

	return infos
}
