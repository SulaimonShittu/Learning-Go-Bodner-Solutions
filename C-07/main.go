package main

import (
	"fmt"
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
	type teamm struct {
		name TeamName
		wins int
	}
	data := make([]teamm, 0, len(l.Wins))
	for k, v := range l.Wins {
		data = append(data, teamm{
			k,
			v,
		})
	}
	sortedTeams := make([]teamm, 0, len(data))
	fmt.Println(sortedTeams)
	return teamRanking
}

func main() {

}
