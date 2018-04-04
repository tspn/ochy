package main

import (
	"fmt"
	"github.com/tspn/ochy/project"
)

func doInit(args *ProgramArgs){
	if len(args.Args) != 3 {
		panic(fmt.Sprintln("could not processes command", args.Args))
	}

	pn := args.Args[2]
	ps := project.New(pn)
	ps.Save()
}