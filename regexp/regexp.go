package regexp

import "regexp"

// Groups is used to find the capture groups in an expression
func Groups(target, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindStringSubmatch(target)[1:]
}

// Matches is used to see if a regexp pattern matches in string
func Matches(target, pattern string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(target)
}

// Find is used to find regexp matches
func Find(target, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(target, -1)
}
