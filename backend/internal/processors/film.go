package processors

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func OrganizeFilm(path string) string {
	quality := []string{
		"bdrip",
		"webdl",
		"webrip",
		"bdremux",
		"camrip",
		"web-dl",
	}
	reYear := regexp.MustCompile("[0-9]{4}")
	reResolution := regexp.MustCompile("[0-9]{3,4}p")
	reName := regexp.MustCompile(`[^[^(|\d]*`)
	reExtension := regexp.MustCompile(`[$\w]*$`)

	filmQuality := "unk"
	filmResolution := "unkResolution"
	filmName := ""
	filmYear := ""

	log.Printf("organizing %s\n", path)

	fileNameRaw := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]
	fileName := strings.ToLower(fileNameRaw)

	fileExtension := reExtension.FindString(fileName)

	for _, q := range quality { // Determine film source quality
		if strings.Contains(fileName, q) {
			filmQuality = q
			break
		}
	}

	filmResolutionTmp := reResolution.FindString(fileName) // Determine film resolution
	if len(filmResolutionTmp) > 0 {
		filmResolution = filmResolutionTmp
	}

	filmYearTmp := reYear.FindString(fileName) // Determine film year
	if len(filmYearTmp) > 0 {
		filmYear = filmYearTmp
	}

	filmNameTmp := reName.FindString(fileName) // Determine film name
	if len(filmYearTmp) > 0 {
		filmName = strings.ToUpper(string(filmNameTmp[0])) + filmNameTmp[1:]
		filmName = strings.ReplaceAll(filmName, ".", " ")
	}

	finalName := fmt.Sprintf(
		"%s (%s) - %s %s.%s",
		filmName,
		filmYear,
		filmQuality,
		filmResolution,
		fileExtension,
	)

	log.Printf("organized %s as\n %s\n", fileName, finalName)

	return finalName
}
