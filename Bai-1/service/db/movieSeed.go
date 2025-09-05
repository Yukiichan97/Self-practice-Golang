package db

import (
	"awesomeProject6/model"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SeedMoviesFromCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Can not open: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Can not read CSV: %v", err)
	}

	for i, record := range records[1:] {
		if len(record) < 3 {
			log.Printf("Skipping row %d: insufficient data", i+2)
			continue
		}

		title, year := parseTitle(record[1])
		if title == "" {
			log.Printf("Skipping row %d: cannot parse title %s", i+2, record[1])
			continue
		}
		genre := parseGenres(record[2])

		movie := model.Movie{
			Title: title,
			Year:  year,
			Genre: genre,
		}

		if err := DB.Create(&movie).Error; err != nil {
			log.Printf("Error inserting movie %s: %v", title, err)
			continue
		}

		if i%1000 == 0 {
			log.Printf("Processed %d movies", i)
		}
	}
	return nil
}

func parseTitle(titleWithYear string) (string, int) {
	i := regexp.MustCompile(`\((\d+)\)`)
	matches := i.FindStringSubmatch(titleWithYear)

	if len(matches) != 2 {
		return titleWithYear, 0
	}

	titlestr := strings.Replace(titleWithYear, matches[0], "", -1)
	year, err := strconv.Atoi(matches[1])
	if err != nil {
		year = 0
	}
	return titlestr, year
}

func parseGenres(genresStr string) string {
	if genresStr == "" || genresStr == "(no genres listed)" {
		return "Unknow"
	}

	genres := strings.Split(genresStr, "|")

	return strings.Join(genres, ", ")
}
