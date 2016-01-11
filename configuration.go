package main

type conf struct {
	dict string
}

var config conf

func initConfiguration() {
	config = conf{dict: "/usr/share/dict/words"}
}
