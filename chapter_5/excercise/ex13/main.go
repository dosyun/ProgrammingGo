package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"./links"
)

func getOrigHosts(worklist []string) (map[string]bool, error) {
	hosts := make(map[string]bool)
	for _, rawurl := range worklist {
		parsed, err := url.Parse(rawurl)
		if err != nil {
			return hosts, fmt.Errorf("url.Parse(%s) error [%v]", rawurl, err)
		}
		hosts[parsed.Host] = true
	}
	return hosts, nil
}

func breadthFirst(f func(item string, origHosts map[string]bool) []string, worklist []string) {
	hosts, err := getOrigHosts(worklist)
	if err != nil {
		log.Print(err)
	}
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, hosts)...)
			}
		}
	}
}

func crawl(rawurl string, origHosts map[string]bool) []string {
	fmt.Println(rawurl)
	list, err := links.Extract(rawurl, origHosts)
	if err != nil {
		fmt.Println(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
