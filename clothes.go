package main

type Clothes struct {
	Name       string
	Color      string
	MinTemp    int
	MaxTemp    int
	PartOfBody BodyPart
}

type BodyPart int

const (
	Head BodyPart = iota
	Torso
	Leg
	Foot
)
