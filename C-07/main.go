package main

import (
	"fmt"
	"io"
	"maps"
	"os"
)

type TeamName string

type Team struct {
	Name         TeamName
	PlayersNames []string
}

type League struct {
	Teams []Team
	Wins  map[TeamName]int
}

func (l *League) MatchResult(firstTeam TeamName,
	firstTeamScore int, secondTeam TeamName, secondTeamScore int) {
	if l.Wins == nil {
		l.Wins = map[TeamName]int{}
	}
	if firstTeamScore > secondTeamScore {
		l.Wins[firstTeam]++
		return
	}
	if secondTeamScore > firstTeamScore {
		l.Wins[secondTeam]++
		return
	}
}

func (l *League) Ranking() []TeamName {
	teamRanking := make([]TeamName, 0, len(l.Teams))
	data := maps.Clone(l.Wins)
	for len(data) > 0 {
		var maxWin int
		var teamName TeamName
		for k, v := range data {
			if v > maxWin {
				maxWin = v
				teamName = k
			}
		}
		teamRanking = append(teamRanking, teamName)
		delete(data, teamName)
	}
	return teamRanking
}

type Ranker interface {
	Ranking() []TeamName
}

func RankPrinter(r Ranker, writer io.Writer) {
	for _, v := range r.Ranking() {
		io.WriteString(writer, fmt.Sprintln(string(v)))
	}
}

func main() {
	// E-02
	epl := &League{
		Teams: []Team{
			{
				Name: "Chelsea",
				PlayersNames: []string{
					"Delap", "Palmer", "Joao Pedro", "Caicedo", "James",
				},
			},
			{
				Name: "Arsenal",
				PlayersNames: []string{
					"Saka", "Madueke", "Eze", "Gyokeres", "Saliba",
				},
			},
			{
				Name: "Man City",
				PlayersNames: []string{
					"Cherki", "Marmoush", "Haaland", "Rodri", "Semenyo",
				},
			},
			{
				Name: "Spurs",
				PlayersNames: []string{
					"Gray", "Spence", "Simons", "Odobert", "Udogie",
				},
			},
			{
				Name: "Nottingham Forrest",
				PlayersNames: []string{
					"Hudson-Odoi", "Gibbs White", "Aina", "Woods", "Ndoye",
				},
			},
		},
	}
	epl.MatchResult("Chelsea", 3, "Spurs", 0)
	epl.MatchResult("Chelsea", 4, "Arsenal", 2)
	epl.MatchResult("Chelsea", 4, "Man City", 2)
	epl.MatchResult("Chelsea", 2, "Nottingham Forrest", 0)
	epl.MatchResult("Nottingham Forrest", 3, "Spurs", 0)
	epl.MatchResult("Arsenal", 4, "Man City", 5)
	epl.MatchResult("Spurs", 0, "Man City", 5)
	epl.MatchResult("Arsenal", 2, "Nottingham Forrest", 0)

	RankPrinter(epl, os.Stdout)
}
