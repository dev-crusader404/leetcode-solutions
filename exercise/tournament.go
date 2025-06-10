package exercise

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	Name        string
	MatchPlayed int
	Win         int
	Lose        int
	Draw        int
	Points      int
}

func Tally(reader io.Reader, writer io.Writer) error {
	readChan := make(chan string)
	errChan := make(chan error)
	teams := initMap()
	go Read(reader, readChan, errChan)
	for {
		select {
		case line, ok := <-readChan:
			if !ok {
				readChan = nil
			} else {
				err := Parse(line, teams)
				if err != nil {
					return err
				}
			}
		case err, ok := <-errChan:
			if !ok {
				errChan = nil
			}
			if err != nil {
				return err
			}
		}
		if readChan == nil && errChan == nil {
			break
		}
	}
	teamDetail := make([]team, 0)
	for _, ele := range teams {
		teamDetail = append(teamDetail, *ele)
	}
	sort.Slice(teamDetail, func(i, j int) bool {
		if teamDetail[i].Points == teamDetail[j].Points {
			return teamDetail[i].Name < teamDetail[j].Name
		}
		return teamDetail[i].Points > teamDetail[j].Points
	})
	printResult(writer, teamDetail)
	return nil
}

func initMap() map[string]*team {
	t := map[string]*team{
		"Devastating Donkeys":     {Name: "Devastating Donkeys"},
		"Allegoric Alaskians":     {Name: "Allegoric Alaskians"},
		"Blithering Badgers":      {Name: "Blithering Badgers"},
		"Courageous Californians": {Name: "Courageous Californians"},
	}
	return t
}

func Read(reader io.Reader, readChan chan<- string, errChan chan<- error) {
	b := make([]byte, 1024)
	for {
		n, err := reader.Read(b)
		if err != nil {
			if err == io.EOF {
				if n > 0 {
					readChan <- string(b[:n])
				}
				closure(readChan, errChan)
				break
			}
			errChan <- err
			closure(readChan, errChan)
			break
		}
		readChan <- string(b[:n])
	}
}

func closure(r chan<- string, err chan<- error) {
	close(r)
	close(err)
}

func Parse(s string, teams map[string]*team) error {
	if len(s) == 0 {
		return errors.New("empty input")
	}
	s = strings.TrimSpace(s)
	line := strings.Split(s, "\n")
	for _, l := range line {
		l = strings.TrimSpace(l)
		data := strings.Split(l, ";")
		if len(data) == 0 || strings.HasPrefix(data[0], "#") || data[0] == "" {
			continue
		}
		if _, ok := teams[data[0]]; !ok {
			return errors.New("invalid team name")
		}
		if _, ok := teams[data[1]]; !ok {
			return errors.New("invalid team name")
		}
		err := updateTeamMap(teams, data[0], data[1], data[2])
		if err != nil {
			return err
		}
	}
	return nil
}

func updateTeamMap(m map[string]*team, t1, t2, result string) error {
	team1 := m[t1]
	team1.MatchPlayed++
	team2 := m[t2]
	team2.MatchPlayed++
	if result == "win" {
		team1.Win++
		team1.Points += 3
		team2.Lose++
	} else if result == "draw" {
		team1.Draw++
		team1.Points++
		team2.Draw++
		team2.Points++
	} else if result == "loss" {
		team2.Win++
		team2.Points += 3
		team1.Lose++
	} else {
		return errors.New("invalid result")
	}
	return nil
}

func printResult(w io.Writer, data []team) {
	fmt.Fprintf(w, "%-30s | MP |  W |  D |  L |  P\n", "Team")
	for _, ele := range data {
		fmt.Fprintf(w, "%-30s |  %d |  %d |  %d |  %d |  %d\n",
			ele.Name, ele.MatchPlayed, ele.Win, ele.Draw, ele.Lose, ele.Points)
	}
}
