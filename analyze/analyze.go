package analyze

import (
	"fmt"
	"regexp"
	"strings"
)

var regex = ``

func getRegexRes(r *regexp.Regexp, input string) []map[string]string {
	// 这里要用FindAllStringSubmatch，找到所有的匹配
	submatch := r.FindAllStringSubmatch(input, -1)

	groupNames := r.SubexpNames()
	var result []map[string]string // slice of map

	// 循环所有行
	for _, user := range submatch {
		m := make(map[string]string)

		// 对每一行生成一个map
		for j, name := range groupNames {
			if j != 0 && name != "" {
				m[name] = strings.TrimSpace(user[j])
			}
		}
		result = append(result, m)
	}
	return result
}

func Analyze(in string) (string, error) {
	regex, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	command := "#!/bin/sh\n"

	for _, res := range getRegexRes(regex, in) {
		if _, ok := res["package"]; ok {
			command = fmt.Sprintf("%secho \"update %s\"\ncargo update -p %s\n", command, res["package"], res["package"])
		}
	}

	return command, nil
}
