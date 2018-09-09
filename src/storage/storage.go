package storage

import (
	"math"
)

const Female = "F"
const Male = "M"
const UnknownGender = "UNKNWON"
const DiferentiatingPercent = 60
const AgeIntervalsNumber = 9

var AgeIntervals [AgeIntervalsNumber]AgeInterval

func init() {
	AgeIntervals = [AgeIntervalsNumber]AgeInterval{
		{AgeMin: 0, AgeMax: 3},
		{AgeMin: 4, AgeMax: 7},
		{AgeMin: 8, AgeMax: 14},
		{AgeMin: 15, AgeMax: 22},
		{AgeMin: 23, AgeMax: 33},
		{AgeMin: 34, AgeMax: 44},
		{AgeMin: 45, AgeMax: 56},
		{AgeMin: 57, AgeMax: 100},
	}
}

type AgeInterval struct {
	AgeMin int
	AgeMax int
}

type Person struct {
	AgeIdentifier int
	Gender        string
}

type Showroom struct {
	Persons []Person
	Cameras map[int]Person
}

type Storage struct {
	Showrooms map[int]Showroom
}

func (storage Storage) PersonInShowroom(showroomId int, person Person) {
	storage.Showrooms[showroomId] = storage.Showrooms[showroomId].personIn(person)
}

func (storage Storage) PersonOutShowroom(showroomId int, person Person) {
	storage.Showrooms[showroomId] = storage.Showrooms[showroomId].personOut(person)
}

func (storage Storage) PersonInFrontOfCamera(showroomId int, cameraId int, person Person) {
	storage.Showrooms[showroomId].personInFrontOfCamera(cameraId, person)
}

func (storage Storage) GetRelevantAgeAndGender(showroomId int) (int, string) {
	var showroom = storage.Showrooms[showroomId]

	return showroom.getRelevantAge(), showroom.getRelevantGender()
}

func (storage Storage) GetPersonInFrontOfCamera(showroomId int, cameraId int) Person {
	showroom := storage.Showrooms[showroomId]
	cameraPerson, ok := showroom.Cameras[cameraId]

	if !ok {
		return Person{}
	}

	return cameraPerson
}

func (showroom Showroom) personInFrontOfCamera(cameraId int, person Person) {
	showroom.Cameras[cameraId] = person
}

func (showroom Showroom) personIn(person Person) Showroom{
	showroom.Persons =  append(showroom.Persons, person)

	return showroom
}

func (showroom Showroom) personOut(person Person) Showroom{
	OuterLoop:
	for key, showroomPerson := range showroom.Persons {
		if person.AgeIdentifier == showroomPerson.AgeIdentifier && person.Gender == showroomPerson.Gender {
			showroom.Persons = append(showroom.Persons[:key], showroom.Persons[key+1:]...)
			break OuterLoop
		}
	}

	return showroom
}

func (showroom Showroom) getRelevantAge() int {
	var relevantInterval int
	var intervalsCount = [AgeIntervalsNumber]int{}
	var intervalMax = 0

	for _, person := range showroom.Persons {
		intervalsCount[person.AgeIdentifier]++
	}

	for key, count := range intervalsCount {
		if count > intervalMax {
			intervalMax = count
			relevantInterval = key
		}
	}

	return relevantInterval
}

func (showroom Showroom) getRelevantGender() string {
	womens := 0
	men := 0

	for _, person := range showroom.Persons {
		if person.Gender == Male {
			men++
		} else if person.Gender == Female {
			womens++
		}
	}

	totalShowroomPeople := len(showroom.Persons)

	womensPercent := float64(womens * 100 / totalShowroomPeople)
	menPercent := float64(men * 100 / totalShowroomPeople)

	if math.Max(womensPercent, menPercent) < DiferentiatingPercent {
		return UnknownGender
	}

	if womensPercent >= DiferentiatingPercent {
		return Female
	} else {
		return Male
	}
}
