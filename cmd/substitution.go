package main

import (
	"math/rand"
	"regexp"
	"strings"
)

var SUBSTITUTION_MAP = map[string][]string{
	"$bodyArmour": BODY_ARMOURS,
	"$weapon":     WEAPONS,
	"$ammo":       AMMO,
}

func process(s string) string {
	matcher := regexp.MustCompile(`\$\w+`)
	subs := matcher.FindAllString(s, -1)
	ret := s
	for _, s := range subs {
		sub := SUBSTITUTION_MAP[s][rand.Intn(len(SUBSTITUTION_MAP[s]))]
		ret = strings.Replace(ret, s, sub, 1)
	}
	return ret
}
