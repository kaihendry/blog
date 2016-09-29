package blog

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Metacrap = `<!DOCTYPE html>
<html lang=en>
<head>
<meta charset="utf-8" />
<link href="/style.css" rel="stylesheet">
<meta name=viewport content="width=device-width, initial-scale=1">
<meta name="twitter:creator" content="@kaihendry">
<meta name='mobile-web-app-capable' content='yes'>
<meta name='apple-mobile-web-app-capable' content='yes'>
<meta name='application-name' content='natalian'>
<meta name='apple-mobile-web-app-status-bar-style' content='gold'>
`

func parseMetaLine(line string) (string, string, error) {
	line = strings.TrimSpace(line)
	item := strings.TrimPrefix(line, `[[!meta `)
	splitItem := strings.Split(item, "=\"")
	//fmt.Printf("I: %q\n", splitItem)
	//fmt.Println("one", splitItem[0], "two", splitItem[1])
	//fmt.Println("S:", splitItem[1])
	if len(splitItem) != 2 {
		return "", "", fmt.Errorf("Error parsing line")
	}
	splitItem[1] = strings.TrimSuffix(splitItem[1], "\"]]")
	splitItem[1] = strings.TrimSuffix(splitItem[1], "\" ]]")
	return splitItem[0], splitItem[1], nil
}

func GetKey(fileName string, keys ...string) map[string]string {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	result := map[string]string{}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return result
		}

		if !strings.HasPrefix(line, "[[!meta") {
			break
		}

		key, value, err := parseMetaLine(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		for _, k := range keys {
			if key == k {
				result[key] = value
			}
		}

	}
	return result
}
