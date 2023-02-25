package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/eXoterr/SimpleFileOrganizerBackend/internal/config"
	"github.com/eXoterr/SimpleFileOrganizerBackend/pkg/utils"
	"github.com/julienschmidt/httprouter"
)

type FileList struct {
	Files []File `json:"files"`
}

type File struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	IsDir    bool   `json:"directory"`
	Preview  string `json:"preview,omitempty"`
}

func ListFiles(cfg *config.Config) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		path := cfg.BasePath + r.URL.Query().Get("path")

		fs, err := os.ReadDir(path)
		utils.LogCheckError(err)
		list := FileList{}
		for _, file := range fs {

			info, err := file.Info()
			utils.LogCheckError(err)

			list.Files = append(list.Files, File{
				Filename: info.Name(),
				Size:     info.Size(),
				IsDir:    info.Mode().IsDir(),
			})
		}
		resp, err := json.Marshal(list)
		utils.HttpCheckError(err, w)
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "*")
		header.Set("Access-Control-Allow-Origin", "*")
		w.Write(resp)
	}
}

func ListLibrary(cfg *config.Config) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		path := cfg.LibraryPath + r.URL.Query().Get("path")

		fs, err := os.ReadDir(path)
		utils.LogCheckError(err)
		list := FileList{}
		for _, file := range fs {

			info, err := file.Info()
			utils.LogCheckError(err)

			list.Files = append(list.Files, File{
				Filename: info.Name(),
				Size:     info.Size(),
				IsDir:    info.Mode().IsDir(),
			})
		}
		resp, err := json.Marshal(list)
		utils.HttpCheckError(err, w)
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "*")
		header.Set("Access-Control-Allow-Origin", "*")
		w.Write(resp)
	}
}
