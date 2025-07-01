package exercise

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SearchRules struct {
	enableLineNumber bool
	fileNameOnly     bool
	caseInsensitive  bool
	invert           bool
	entireLine       bool
}

func Search(pattern string, flags, files []string) []string {
	if len(files) == 0 || len(pattern) == 0 {
		return []string{}
	}
	sr := SearchRules{}
	for _, f := range flags {
		switch f {
		case "-n":
			sr.enableLineNumber = true
		case "-l":
			sr.fileNameOnly = true
		case "-i":
			sr.caseInsensitive = true
		case "-v":
			sr.invert = true
		case "-x":
			sr.entireLine = true
		}
	}
	return sr.searchPattern(pattern, files)
}

func (s SearchRules) searchPattern(pattern string, files []string) []string {
	var result []string
	for _, fName := range files {
		file, err := os.Open(fName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return result
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		lineNumber := 1
		copyOfPattern := pattern

		for scanner.Scan() {
			var match, prefix string
			if len(files) > 1 {
				prefix = fName + ":"
			}
			line := strings.TrimSpace(scanner.Text())
			copyLine := line
			if s.caseInsensitive {
				copyLine = strings.ToLower(line)
				copyOfPattern = strings.ToLower(pattern)
			}
			if strings.Contains(copyLine, copyOfPattern) && !s.invert {
				if s.entireLine && copyLine == copyOfPattern {
					match = line
				} else if !s.entireLine {
					match = line
				}
				if s.enableLineNumber {
					match = fmt.Sprintf("%d:%s", lineNumber, match)
				}
				if s.fileNameOnly {
					result = append(result, fName)
					break
				}
			} else if s.invert && !strings.Contains(copyLine, copyOfPattern) {
				match = line
			}

			if len(match) > 0 {
				if len(prefix) > 0 {
					match = prefix + match
				}
				result = append(result, match)
			}
			lineNumber++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	}
	return result
}
