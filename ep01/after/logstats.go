package logstats

import "strings"

func Count5xx(logs []string) int {
	count := 0
	for _, line := range logs {
		for range 3 {
			_, line, _ = strings.Cut(line, "|")
		}
		status, _, _ := strings.Cut(line, "|")
		if len(status) == 3 && status[0] == '5' {
			count++
		}
	}
	return count
}
