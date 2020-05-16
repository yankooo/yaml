package yaml

import (
	"os"
	"regexp"
	"strings"
)

var (
	envValReg, allPatternReg *regexp.Regexp
)

const (
	envValRegPattern = `\${(.*?)}`
	// all Pattern
	allPattern = `\$\[([^\|]+)\|\|([^\]]+)`
)
func init()  {
	envValReg = regexp.MustCompile(envValRegPattern)
	allPatternReg = regexp.MustCompile(allPattern)
}
// if string like $[${NAME}||archaius]
// will query environment variable for ${NAME}
// if environment variable is "" return default string `archaius`
func expandValueEnv(value string) (realValue string) {
	realValue = value
	if !allPatternReg.MatchString(value) {
		return realValue
	}

	match := allPatternReg.FindAllStringSubmatch(value, -1)
	envVal, realValue := match[0][1],match[0][2]

	// convert env val
	envMatch := envValReg.FindAllString(envVal, -1)
	if len(envMatch) == 0 {
		return
	}
	for i, s := range envMatch {
		newValFiled := os.Getenv(s[2:len(s)-1])
		if newValFiled == "" {
			return
		}
		envVal = strings.ReplaceAll(envVal, envMatch[i], newValFiled)
 	}

	return envVal
}