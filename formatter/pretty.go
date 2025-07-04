package Formatter

import (
	"regexp"
	"strings"
)

const (
	Reset = "\033[0m"
	Cyan  = "\033[36m"
	Green = "\033[32m"
)

// עיצוב SQL עם שורות חדשות והזחות
func PrettySQL(sql string) string {
	keywords := []string{"SELECT", "FROM", "WHERE", "GROUP BY", "HAVING", "ORDER BY", "LIMIT", "/*"}
	for _, kw := range keywords {
		sql = strings.ReplaceAll(sql, kw, "\n"+kw)
	}
	sql = strings.ReplaceAll(sql, ",", ",\n  ")
	sql = strings.ReplaceAll(sql, "AND", "\n    AND")
	sql = strings.ReplaceAll(sql, "OR", "\n    OR")
	return strings.TrimSpace(sql)
}

// צביעה של מילות מפתח
func ColorizeSQL(sql string) string {
	keywords := []string{
		"SELECT", "FROM", "WHERE", "GROUP BY", "HAVING",
		"LIMIT", "ORDER BY", "AND", "OR", "NOT", "AS",
	}
	for _, kw := range keywords {
		pattern := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(kw) + `\b`)
		sql = pattern.ReplaceAllStringFunc(sql, func(match string) string {
			return Cyan + strings.ToUpper(match) + Reset
		})
	}
	return sql
}
