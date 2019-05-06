//Package korname is Golang port of Korean Name Generator by agemor
//https://github.com/agemor/korean-name-generator
package korname

import (
	"log"
	"math/rand"
)

const (
	//Male means male
	Male = iota
	//Female means female
	Female
)

var namedata dataset

func init() {
	var err error
	namedata, err = loadDataset("dataset.json")
	if err != nil {
		log.Fatal(err)
	}
}

func pick(count int, set func(int) int) int {
	sum := 0
	selected := 0

	for i := 0; i < count; i++ {
		sum += set(i)
	}

	if sum == 0 {
		sum++
	}

	pivot := rand.Intn(int(sum))

	for i := 0; i < count; i++ {
		if pivot -= set(i); pivot <= 0 {
			selected = i
			break
		}
	}

	return selected
}

func pickSyllable(data int, set int) rune {
	var choseong int
	var jungseong int
	var jongseong int
	switch data {
	case Male:
		choseong = pick(19, func(n int) int {
			if n >= len(namedata.MaleFirstName[set][0]) {
				return 0
			}
			return namedata.MaleFirstName[set][0][n]
		})
		jungseong = pick(21, func(n int) int {
			if choseong*21+n >= len(namedata.MaleFirstName[set][1]) {
				return 0
			}
			return namedata.MaleFirstName[set][1][choseong*21+n]
		})
		jongseong = pick(28, func(n int) int {
			if jungseong*28+n >= len(namedata.MaleFirstName[set][2]) || choseong*28+n >= len(namedata.MaleFirstName[set][3]) {
				return 0
			}
			return namedata.MaleFirstName[set][2][jungseong*28+n] * namedata.MaleFirstName[set][3][choseong*28+n]
		})
	case Female:
		choseong = pick(19, func(n int) int {
			if n >= len(namedata.FemaleFirstName[set][0]) {
				return 0
			}
			return namedata.FemaleFirstName[set][0][n]
		})
		jungseong = pick(21, func(n int) int {
			if choseong*21+n >= len(namedata.FemaleFirstName[set][1]) {
				return 0
			}
			return namedata.FemaleFirstName[set][1][choseong*21+n]
		})
		jongseong = pick(28, func(n int) int {
			if jungseong*28+n >= len(namedata.FemaleFirstName[set][2]) || choseong*28+n >= len(namedata.FemaleFirstName[set][3]) {
				return 0
			}
			return namedata.FemaleFirstName[set][2][jungseong*28+n] * namedata.FemaleFirstName[set][3][choseong*28+n]
		})
	}

	return constructJamo(choseong, jungseong, jongseong)
}

func pickLastName() rune {
	lastNameIndex := pick(len(namedata.LastName), func(n int) int { return namedata.LastNameFreq[n] })
	return rune(namedata.LastName[lastNameIndex] + 0xAC00)
}

func constructJamo(cho, jung, jong int) rune {
	return rune(0xAC00 + 28*21*cho + 28*jung + jong)
}

//Generate generates random korean name with given gender
func Generate(gender int) string {
	return string(pickLastName()) + string(pickSyllable(gender, 0)) + string(pickSyllable(gender, 1))
}
