package yaml

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	envValReg, allPatternReg,keyPatternReg *regexp.Regexp
)

const (
	envValRegPattern = `\${(.*?)}`
	// all Pattern
	allPattern = `\$\[.*[|]{2}.*\]`
	keyPattern = `\${(.*)}`
)
func init()  {
	envValReg = regexp.MustCompile(envValRegPattern)
	allPatternReg = regexp.MustCompile(allPattern)
	keyPatternReg = regexp.MustCompile(keyPattern)
}
// if string like $[${NAME}||archaius]
// will query environment variable for ${NAME}
// if environment variable is "" return default string `archaius`
func expandValueEnv(value string) (realValue string) {
	realValue = value
	if !allPatternReg.MatchString(value) {
		return realValue
	}
	key := keyPatternReg.FindString(value)
	fmt.Println(key)
	defaultVal := strings.TrimPrefix(value[2:], key)
	defaultVal = defaultVal[2:len(defaultVal)-1]
	envMatch := envValReg.FindAllString(key, -1)

	for i, s := range envMatch {
		fmt.Println(s[2:len(s)-1])
		newVal := os.Getenv(s[2:len(s)-1])
		if newVal == "" {
			return defaultVal
		}
		key = strings.ReplaceAll(key, envMatch[i], newVal)
 	}

	return key
}