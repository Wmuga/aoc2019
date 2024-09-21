package fileparser

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wmuga/aoc2019/pkg/models"
)

type ParseError struct {
	file   string
	line   int
	reason error
}

func (p ParseError) Error() string {
	return fmt.Sprintf("Error parsing %s on line %d: %s", p.file, p.line, p.reason)
}

func (p ParseError) Unwrap() error {
	return p.reason
}

func NewParseErrorCreator(filename string) func(int, error) ParseError {
	return func(line int, reason error) ParseError {
		return ParseError{
			file:   filename,
			line:   line,
			reason: reason,
		}
	}
}

const (
	prefixTest   = "@Test"
	prefixAnswer = "@Answer"
	prefixPart   = "@Part"
)

var (
	ErrNoTestName = errors.New("No test name")
	ErrNoPart     = errors.New("No part number given")
)

// ReadTests reads parses custom test file format
func ReadTests(filename string) ([]models.Test, error) {
	res := make([]models.Test, 0)
	testFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(testFile), "\n")
	data := models.Test{}
	newError := NewParseErrorCreator(filename)

	var isData bool
	for num, line := range lines {
		// data section - just add lines
		if isData {
			data.Data = append(data.Data, line)
			continue
		}
		line = strings.TrimSpace(line)
		// skip empty lines
		if line == "" {
			continue
		}

		// new test
		if strings.HasPrefix(line, prefixTest) {
			isData = false
			// check for name
			if data.Name == "" && len(data.Data) != 0 {
				return nil, newError(num+1, ErrNoTestName)
			}
			// Non empty name - add old test
			if data.Name != "" {
				res = append(res, data)
				data = models.Test{}
			}
			_, name, ok := getPair(line)
			if !ok {
				return nil, newError(num+1, ErrNoTestName)
			}

			data.Name = name
			continue
		}

		if strings.HasPrefix(line, prefixPart) {
			_, part, ok := getPair(line)
			if !ok {
				return nil, newError(num+1, ErrNoPart)
			}

			partNum, err := strconv.Atoi(part)
			if err != nil {
				return nil, newError(num+1, err)
			}
			data.Part = partNum
			continue
		}

		if strings.HasPrefix(line, prefixAnswer) {
			_, ans, ok := getPair(line)
			if ok {
				data.Answer = ans
			}
			continue
		}

		isData = true
		data.Data = append(data.Data, line)

		if data.Name == "" {
			return nil, newError(num+1, ErrNoTestName)
		}

		if data.Part == 0 {
			return nil, newError(num+1, ErrNoPart)
		}
	}

	return res, nil
}

func GetInput(inFile string) ([]string, error) {
	data, err := os.ReadFile(inFile)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

// getPair splits line by space. returns first two words with ok = true
// else ok = false
func getPair(line string) (first, second string, ok bool) {
	data := strings.Split(line, " ")
	if len(data) < 2 {
		return
	}
	return data[0], data[1], true
}
