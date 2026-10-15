package logstats

import "strings"

func Count5xx(logs []string) int {
	count := 0
	for _, line := range logs {
		fields := strings.Split(line, "|")
		if len(fields) < 4 {
			continue
		}
		if len(fields[3]) == 3 && fields[3][0] == '5' {
			count++
		}
	}
	return count
}
