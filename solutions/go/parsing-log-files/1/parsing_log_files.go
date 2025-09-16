package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(?:TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<\W*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`".*(?i)password.*"`)
    matches := 0
    for _, line := range lines {
        match := re.FindStringSubmatch(line)
        if nil != match {
            matches += len(match)
        }
    }
    return matches
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+\w+`)
    re2 := regexp.MustCompile(`User\s+`)
    for i, line := range lines {
        if re.MatchString(line) {
            sub := re.FindString(line)
            user := re2.Split(sub, 2)
            lines[i] = "[USR] " + user[1] + " " + line
        }
    }
    return lines
}
