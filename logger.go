package nlogger

import (
	"fmt"
	"os"
	"time"
)

var dirName = "log"

type Logger interface {
	//Error creates an error log describing the provided error and time.
	Error(err error, givenTime time.Time)

	//Warn creates an warning log describing the provided warning message and time.
	Warn(msg string, givenTime time.Time)

	//Update creates a log for an updated item.
	//
	//Here entity indicates what type of item is updated such as: user, product, cart etc.
	//
	//And id indicates the id of updated item and msg for providing a custom message
	Update(entity, id interface{}, msg string, givenTime time.Time)

	//Delete creates a log for a deleted item.
	//
	//Here entity indicates what type of item is deleted such as: user, product, cart etc.
	//
	//And id indicates the id of deleted item and msg for providing a custom message
	Delete(entity, id interface{}, msg string, givenTime time.Time)

	//Endpoint cretes a log describing the provided endpoint that is visited by someone.
	//
	//It also mention the visitor's ip address and the visiting time.
	Endpoint(ep string, ip interface{}, givenTime time.Time)
}

type LoggerService struct {
	File *os.File
	//DefaultFile *os.File
}

//Error creates an error log describing the provided error and time.
func (ls *LoggerService) Error(err error, givenTime time.Time) {
	//preparing data to write
	currTime := time.Now()
	dateTime := currTime.Format("2006.01.02 15:04:05")
	errStr := fmt.Sprintf("%v", err)

	text := dateTime + " [ Error ]\t\t" + errStr + " - " + givenTime.Format("2006.01.02 15:04:05") + "\n"

	filePath := getFilePath(ls.File)
	ls.appendToFile(filePath, text)
}

//Warn creates an warning log describing the provided warning message and time.
func (ls *LoggerService) Warn(msg string, givenTime time.Time) {
	//preparing data to write
	currTime := time.Now()
	dateTime := currTime.Format("2006.01.02 15:04:05")

	text := dateTime + " [ Warn ]\t\t" + msg + " - " + givenTime.Format("2006.01.02 15:04:05") + "\n"

	filePath := getFilePath(ls.File)
	ls.appendToFile(filePath, text)
}

//Update creates a log for an updated item.
//
//Here entity indicates what type of item is updated such as: user, product, cart etc.
//
//And id indicates the id of updated item and msg for providing a custom message
func (ls *LoggerService) Update(entity, id interface{}, msg string, givenTime time.Time) {
	//preparing data to write
	currTime := time.Now()
	dateTime := currTime.Format("2006.01.02 15:04:05")
	ent := fmt.Sprintf("%v", entity)
	idStr := fmt.Sprintf("%v", id)

	text := dateTime + " [ Update ]\t\t" + ent + " - " + idStr + ": " + msg + " - " + givenTime.Format("2006.01.02 15:04:05") + "\n"

	filePath := getFilePath(ls.File)
	ls.appendToFile(filePath, text)
}

//Delete creates a log for a deleted item.
//
//Here entity indicates what type of item is deleted such as: user, product, cart etc.
//
//And id indicates the id of deleted item and msg for providing a custom message
func (ls *LoggerService) Delete(entity, id interface{}, msg string, givenTime time.Time) {
	//preparing data to write
	currTime := time.Now()
	dateTime := currTime.Format("2006.01.02 15:04:05")
	ent := fmt.Sprintf("%v", entity)
	idStr := fmt.Sprintf("%v", id)

	text := dateTime + " [ Delete ]\t\t" + ent + " - " + idStr + ": " + msg + " - " + givenTime.Format("2006.01.02 15:04:05") + "\n"

	filePath := getFilePath(ls.File)
	ls.appendToFile(filePath, text)
}

//Endpoint cretes a log describing the provided endpoint that is visited by someone.
//
//It also mention the visitor's ip address and the visiting time.
func (ls *LoggerService) Endpoint(ep string, ip interface{}, givenTime time.Time) {
	//preparing data to write
	currTime := time.Now()
	dateTime := currTime.Format("2006.01.02 15:04:05")
	ipStr := fmt.Sprintf("%v", ip)

	text := dateTime + " [ Endpoint ]\t" + ep + " " + ipStr + " - " + givenTime.Format("2006.01.02 15:04:05") + "\n"

	filePath := getFilePath(ls.File)
	ls.appendToFile(filePath, text)
}

func getFilePath(file *os.File) string {
	info, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	fName := info.Name()

	return dirName + "/" + fName
}

func (ls *LoggerService) appendToFile(fileName, text string) {
	file, err := os.OpenFile(fileName, os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if _, err := file.WriteString(text); err != nil {
		fmt.Println(err)
	}
}

func getFullFileName(fileName string) string {
	//creating directory if does not exist
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		//fmt.Println("Dir does not exist")
		os.Mkdir(dirName, 0755) //(owner:7=rwx group:5=r-x other:5=r-x) This means that the directory has the default permissions -rwxr-xr-x (represented in octal notation as 0755).
	}

	fullFileName := dirName + "/" + fileName

	return fullFileName
}

//NewLog returns a new LoggerService instance with a default file.
func NewLog() *LoggerService {
	fullFileName := getFullFileName("default.log")

	file, err := os.OpenFile(fullFileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	return &LoggerService{file}
}

//NewLogger returns a new LoggerService instance with the provided file name.
func NewLogger(fileName string) *LoggerService {
	fullFileName := getFullFileName(fileName)

	file, err := os.OpenFile(fullFileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	return &LoggerService{file}
}
