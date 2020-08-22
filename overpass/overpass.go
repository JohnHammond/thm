package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-homedir"
)

type passListEntry struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

//Secure encryption algorithm from https://socketloop.com/tutorials/golang-rotate-47-caesar-cipher-by-47-characters-example
func rot47(input string) string {
	var result []string
	for i := range input[:len(input)] {
		j := int(input[i])
		if (j >= 33) && (j <= 126) {
			result = append(result, string(rune(33+((j+14)%94))))
		} else {
			result = append(result, string(input[i]))
		}
	}
	return strings.Join(result, "")
}

//Encrypt the credentials and write them to a file.
func saveCredsToFile(filepath string, passlist []passListEntry) string {
	file, err := os.OpenFile(filepath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	defer file.Close()
	stringToWrite := rot47(credsToJSON(passlist))
	if _, err := file.WriteString(stringToWrite); err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}
	return "Success"
}

//Load the credentials from the encrypted file
func loadCredsFromFile(filepath string) ([]passListEntry, string) {
	buff, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, "Failed to open or read file"
	}
	//Decrypt passwords
	buff = []byte(rot47(string(buff)))
	//Load decrypted passwords
	var passlist []passListEntry
	err = json.Unmarshal(buff, &passlist)
	if err != nil {
		fmt.Println(err.Error())
		return nil, "Failed to load creds"
	}
	return passlist, "Ok"
}

//Convert the array of credentials to JSON
func credsToJSON(passlist []passListEntry) string {
	jsonBuffer, err := json.Marshal(passlist)
	if err != nil {
		fmt.Println(err.Error())
		return "Something went wrong"
	}
	return string(jsonBuffer)
}

//Python style input function
func input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()

	}
	return ""
}

func serviceSearch(passlist []passListEntry, serviceName string) (int, passListEntry) {
	//A linear search is the best I can do, Steve says it's Oh Log N whatever that means
	for index, entry := range passlist {
		if entry.Name == serviceName {
			return index, entry
		}
	}
	return -1, passListEntry{}
}

func getPwdForService(passlist []passListEntry, serviceName string) string {
	index, entry := serviceSearch(passlist, serviceName)
	if index != -1 {
		return entry.Pass
	}
	return "Pass not found"
}

func setPwdForService(passlist []passListEntry, serviceName string, newPwd string) []passListEntry {
	index, entry := serviceSearch(passlist, serviceName)
	//If service exists, update entry
	if index != -1 {
		entry.Pass = newPwd
		passlist[index] = entry
		return passlist
	}
	//If it doesn't, create an entry
	entry = passListEntry{Name: serviceName, Pass: newPwd}
	passlist = append(passlist, entry)
	return passlist
}

func deletePwdByService(passlist []passListEntry, serviceName string) (resultList []passListEntry, status string) {
	index, _ := serviceSearch(passlist, serviceName)
	if index != -1 {
		//remove Pwd from passlist
		resultList = append(passlist[:index], passlist[index+1:]...)
		status = "Ok"
		return
	}
	return passlist, "Pass not found"
}

func printAllPasswords(passlist []passListEntry) {
	for _, entry := range passlist {
		fmt.Println(entry.Name, "\t", entry.Pass)
	}
}

func main() {
	credsPath, err := homedir.Expand("~/.overpass")
	if err != nil {
		fmt.Println("Error finding home path:", err.Error())
	}
	//Load credentials
	passlist, status := loadCredsFromFile(credsPath)
	if status != "Ok" {
		fmt.Println(status)
		fmt.Println("Continuing with new password file.")
		passlist = make([]passListEntry, 0)
	}

	fmt.Println("Welcome to Overpass")

	//Determine function
	option := -1
	fmt.Print(
		"Options:\n" +
			"1\tRetrieve Password For Service\n" +
			"2\tSet or Update Password For Service\n" +
			"3\tDelete Password For Service\n" +
			"4\tRetrieve All Passwords\n" +
			"5\tExit\n")

	for option > 5 || option < 1 {
		optionString := input("Choose an option:\t")
		optionChoice, err := strconv.Atoi(optionString)
		if err != nil || optionChoice > 5 || optionChoice < 1 {
			fmt.Println("Please enter a valid number")
		}
		option = optionChoice
	}

	switch option {
	case 1:
		service := input("Enter Service Name:\t")
		getPwdForService(passlist, service)
	case 2:
		service := input("Enter Service Name:\t")
		newPwd := input("Enter new password:\t")
		passlist = setPwdForService(passlist, service, newPwd)
		saveCredsToFile(credsPath, passlist)
	case 3:
		service := input("Enter Service Name:\t")
		passlist, status := deletePwdByService(passlist, service)
		if status != "Ok" {
			fmt.Println(status)
		}
		saveCredsToFile(credsPath, passlist)
	case 4:
		printAllPasswords(passlist)
	}
}
