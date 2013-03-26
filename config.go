package main

type Rule struct {
	Path string
	Run string
}

func ParseConfig(path string) (rules []*Rule, err error) {
	rules = make([]*Rule, 0)
	return
}
