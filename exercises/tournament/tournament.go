package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Scores struct {
	Played int
	Won    int
	Drawn  int
	Lost   int
	Points int
}

func (s *Scores) win() {
	const POINTS = 3

	s.Won++
	s.Played++
	s.Points += POINTS
}

func (s *Scores) lost() {
	s.Lost++
	s.Played++
}

func (s *Scores) draw() {
	s.Drawn++
	s.Played++
	s.Points++
}

func SortScore(scores map[string]Scores) []string {
	var teams []string = []string{}
	for team := range scores {
		teams = append(teams, team)
	}

	sort.SliceStable(teams, func(i, j int) bool {
		if scores[teams[i]].Points != scores[teams[j]].Points {
			return scores[teams[i]].Points > scores[teams[j]].Points
		}
		return teams[i] < teams[j]
	})

	return teams
}

func Tally(reader io.Reader, writer io.Writer) error {
	const (
		HOME_TEAM = 0
		AWAY_TEAM = 1
		OUTCOME   = 2
	)

	var scores map[string]Scores = map[string]Scores{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {

		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' { // ignore comments and newlines
			continue
		}

		fields := strings.Split(line, ";")
		if len(fields) != 3 {
			return fmt.Errorf("error: unexpected")
		}

		homeTeam := scores[fields[HOME_TEAM]]
		awayTeam := scores[fields[AWAY_TEAM]]
		outcome := fields[OUTCOME]

		switch outcome {
		case "win":
			homeTeam.win()
			awayTeam.lost()
		case "loss":
			homeTeam.lost()
			awayTeam.win()
		case "draw":
			homeTeam.draw()
			awayTeam.draw()
		default:
			return fmt.Errorf("error: unexpected outcome %s", outcome)
		}

		scores[fields[HOME_TEAM]] = homeTeam
		scores[fields[AWAY_TEAM]] = awayTeam
	}

	winners := SortScore(scores)

	w := bufio.NewWriter(writer)
	w.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	for _, v := range winners {
		l := fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d\n", v, scores[v].Played, scores[v].Won, scores[v].Drawn, scores[v].Lost, scores[v].Points)
		w.WriteString(l)
	}

	return w.Flush()
}
