package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	filePath := `C:\Users\lisha\OneDrive\Documents\example_code.go` 
	code, err := readFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	vulnerableString := findVulnerabilities(code)

	outputFilePath := `C:\Users\lisha\OneDrive\Documents\output_example_code.txt`
	err = writeOutputFile(outputFilePath, vulnerableString)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}

	if len(vulnerableString) > 0 {
		fmt.Println("Potential vulnerabilities found. Results written to", outputFilePath)
	} else {
		fmt.Println("No potential vulnerabilities found.")
	}
}

func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var code string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return code, nil
}

func findVulnerabilities(code string) string {
	var vulnerableLines []string

	lines := strings.Split(code, "\n")
	for i, line := range lines {
		// Identify potential SQL injection vulnerabilities
		sqlInjectionPattern := `.*[Dd][Bb].[Ee][Xx][Ee][Cc].*\(.*\)`
		matchedSQLInjection, _ := regexp.MatchString(sqlInjectionPattern, line)

		// Identify potential Cross-Site Scripting (XSS) vulnerabilities
		xssPattern := `.*[Ii][Nn][Nn][Ee][Rr][Hh][Tt][Mm][Ll].*\(.*\)`
		matchedXSS, _ := regexp.MatchString(xssPattern, line)

		// ... (similar checks for other vulnerabilities)
		bufferOverflowPattern := `\b(buffer\s*overflow|heap\s*overflow)\b`
		matchedBufferOverflow, _ := regexp.MatchString(bufferOverflowPattern, line)

		if matchedBufferOverflow {
			vulnerableLines = append(vulnerableLines, fmt.Sprintf("Potential buffer overflow vulnerability found at line %d: %s", i+1, line))
		}

		if matchedSQLInjection {
			vulnerableLines = append(vulnerableLines, fmt.Sprintf("Potential SQL injection vulnerability found at line %d: %s", i+1, line))
		}

		if matchedXSS {
			vulnerableLines = append(vulnerableLines, fmt.Sprintf("Potential XSS vulnerability found at line %d: %s", i+1, line))
		}

		// ... (similar appends for other vulnerabilities)
	}

	return strings.Join(vulnerableLines, "\n")
}

func writeOutputFile(filePath string, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}




