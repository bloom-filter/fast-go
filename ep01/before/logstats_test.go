package logstats

import "testing"

func TestCount5xx(t *testing.T) {
	logs := []string{
		"2026-10-02 10:00:01|GET|/api/users|200|12",
		"2026-10-02 10:00:02|GET|/api/users|500|120",
		"2026-10-02 10:00:03|POST|/api/orders|503|310",
	}
	got := Count5xx(logs)
	if got != 2 {
		t.Fatalf("got %d, want 2", got)
	}
}

func BenchmarkCount5xx(b *testing.B) {
	logs := makeTestLogs(10000)

	for b.Loop() {
		Count5xx(logs)
	}
}

func makeTestLogs(n int) []string {
	logs := make([]string, n)
	for i := range logs {
		status := "200"
		if i%20 == 0 {
			status = "500"
		}
		logs[i] = "2026-10-02 10:00:01|GET|/api/users|" +
			status + "|12"
	}
	return logs
}
