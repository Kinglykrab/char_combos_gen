package main

import (
	"fmt"
	"os"
)

func main() {
	casterClassIDs := []int{11, 12, 13, 14}
	raceIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 128, 130, 330, 522}
	monkRaceIDs := []int{1, 128, 522}
	nonCasterRaceIDs := []int{2, 4, 8, 9, 10, 11, 130}
	const startZoneID = 202
	const startX = -76.0
	const startY = 70.0
	const startZ = -155.7
	const startHeading = 0.0

	combinationsData := "INSERT INTO char_create_combinations VALUES\n"
	startZonesData := "INSERT INTO start_zones VALUES\n"

	allocationData := "INSERT INTO char_create_point_allocations VALUES\n"

	allocationData += "(0, 100, 100, 100, 100, 100, 100, 100, 0, 0, 0, 0, 0, 0, 0);"

	for _, raceID := range raceIDs {
		fmt.Printf("Race %v\n", raceID)

		for classID := 1; classID <= 16; classID++ {
			if classID == 7 && !contains(monkRaceIDs, raceID) {
				continue
			}

			if contains(casterClassIDs, classID) && contains(nonCasterRaceIDs, raceID) {
				continue
			}

			className := getClassName(classID)
			raceName := getRaceName(raceID)

			fmt.Printf("Class %v\n", classID)

			queryEnding := ","
			lineEnding := "\n"

			if raceID == 522 && classID == 16 {
				queryEnding = ";"
				lineEnding = ""
			}

			combinationsData += fmt.Sprintf(
				"(0, %v, %v, 396, %v, 0)%v -- %v %v%v",
				raceID,
				classID,
				startZoneID,
				queryEnding,
				raceName,
				className,
				lineEnding,
			)

			startZonesData += fmt.Sprintf(
				"(%v, %v, %v, %v, %v, 0, 0, %v, 396, %v, %v, %v, %v, %v, 50, -1, -1, NULL, NULL)%v -- %v %v%v",
				startX,
				startY,
				startZ,
				startHeading,
				startZoneID,
				classID,
				raceID,
				startZoneID,
				startX,
				startY,
				startZ,
				queryEnding,
				raceName,
				className,
				lineEnding,
			)
		}

		fmt.Println()
	}

	truncateData := "TRUNCATE char_create_combinations;\n"

	truncateData += "TRUNCATE char_create_point_allocations;\n"

	truncateData += "TRUNCATE start_zones;"

	dataGap := "\n\n"

	// Truncate Data
	fileData := fmt.Sprintf("%v%v", truncateData, dataGap)

	// Allocations Data
	fileData += fmt.Sprintf("%v%v", allocationData, dataGap)

	// Combinations Data
	fileData += fmt.Sprintf("%v%v", combinationsData, dataGap)

	// Start Zones Data
	fileData += fmt.Sprintf("%v%v", startZonesData, dataGap)

	err := os.WriteFile("combinations.sql", []byte(fileData), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func contains(i []int, n int) bool {
	for _, v := range i {
		if v == n {
			return true
		}
	}

	return false
}

func getClassName(classID int) string {
	m := map[int]string{
		1:  "Warrior",
		2:  "Cleric",
		3:  "Paladin",
		4:  "Ranger",
		5:  "Shadow Knight",
		6:  "Druid",
		7:  "Monk",
		8:  "Bard",
		9:  "Rogue",
		10: "Shaman",
		11: "Necromancer",
		12: "Wizard",
		13: "Magician",
		14: "Enchanter",
		15: "Beastlord",
		16: "Berserker",
	}

	if _, ok := m[classID]; !ok {
		return "Unknown Class"
	}

	return m[classID]
}

func getRaceName(raceID int) string {
	m := map[int]string{
		1:   "Human",
		2:   "Barbarian",
		3:   "Erudite",
		4:   "Wood Elf",
		5:   "High Elf",
		6:   "Dark Elf",
		7:   "Half Elf",
		8:   "Dwarf",
		9:   "Troll",
		10:  "Ogre",
		11:  "Halfling",
		12:  "Gnome",
		128: "Iksar",
		130: "Vah Shir",
		330: "Froglok",
		522: "Drakkin",
	}

	if _, ok := m[raceID]; !ok {
		return "Unknown Race"
	}

	return m[raceID]
}
