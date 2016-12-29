package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"

	pb "server/rpcdefsql"
)

var sessionID = 1

var enrolledInstructorsID []int32
var enrolledUsersID []int32

var insMap = make(map[int32]*pb.InstructorInfo)
var userMap = make(map[int32]*pb.UserInfo)

func RegisterEnrolledInstructor(instructorID int32,
	iInfo *pb.InstructorInfo) {

	enrolledInstructorsID = append(enrolledInstructorsID, instructorID)
	insMap[instructorID] = iInfo
}

func RegisterEnrolledUser(userID int32,
	uInfo *pb.UserInfo) {

	enrolledUsersID = append(enrolledUsersID, userID)
	userMap[userID] = uInfo
}

func getEnrolledUser(uid int32) *pb.UserInfo {
	uInfo, ok := userMap[uid]
	if !ok {
		panic("invalid user")
	}
	return uInfo
}

func getEnrolledIns(insID int32) *pb.InstructorInfo {
	iInfo, ok := insMap[insID]
	if !ok {
		panic("invalid ins")
	}
	return iInfo
}

func getEnrolledInstructorID() (error, int32) {
	numIns := len(enrolledInstructorsID)
	if numIns == 0 {
		return errors.New("Instructors not registered"), 0
	}
	return nil, enrolledInstructorsID[rand.Intn(len(enrolledInstructorsID))]
}

func getEnrolledUsersID() (error, int32) {
	numIns := len(enrolledUsersID)
	if numIns == 0 {
		return errors.New("Users not registered"), 0
	}
	return nil, enrolledUsersID[rand.Intn(len(enrolledUsersID))]
}

func GetNewSession() (error, pb.SessionInfo) {

	var si pb.SessionInfo
	var err error
	t := time.Now()
	si.SessionTime = t.String()
	si.SessionType = pb.FitnessCategory_YOGA
	err, si.InstructorInfoID = getEnrolledInstructorID()
	if err != nil {
		return err, si
	}

	//si.TagList = []pb.SessionTag{pb.SessionTag_CALMING, pb.SessionTag_RELAXING}
	si.DifficultyLevel = pb.SessionDifficulty(rand.Intn(3)) //pb.SessionDifficulty_MODERATE
	si.SessionDesc = "my session" + strconv.Itoa(sessionID)
	sessionID += 1
	return nil, si
}

var certID = 1

func GetNewInstructor() pb.InstructorInfo {

	var ui pb.InstructorInfo
	ui.FirstName = randomdata.FirstName(randomdata.Female)
	ui.LastName = randomdata.LastName()
	ui.Email = randomdata.Email()
	ui.City = randomdata.City()
	ui.Certification = "FitnessCert" + strconv.Itoa(certID)
	certID += 1

	return ui
}

func GetNewUser() pb.UserInfo {

	var ui pb.UserInfo
	ui.FirstName = randomdata.FirstName(randomdata.Male)
	ui.LastName = randomdata.LastName()
	ui.Email = randomdata.Email()
	ui.PassWord = randomdata.LastName()
	ui.City = randomdata.City()

	return ui
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func GetUrl() string {
	serv := os.Getenv("SERVERIP")
	if serv == "" {
		fmt.Printf("IP not set")
		os.Exit(-1)
	}
	return serv + ":8099"
}
func GetHttpUrl() string {
	serv := os.Getenv("SERVERIP")
	if serv == "" {
		fmt.Printf("SERVERIP not set")
		os.Exit(-1)
	}
	return "http://" + serv + ":8099"
}
