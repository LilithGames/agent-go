package main

import "github.com/LilithGames/agent-go/pkg/agent"

type Player struct {
	agent.One
	id string
}

func NewPlayer(id string) *Player {
	return &Player{id: id}
}

func (p *Player) ID() string {
	return p.id
}
