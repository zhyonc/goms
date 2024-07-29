package test

import (
	"fmt"
	"goms/nxfile"
	"testing"
)

const (
	nxFileDir string = "../"
)

var (
	MapSearches       []string = []string{"Map/Map0"}
	characterSearches []string = []string{"Weapon", "Glove", "Face", "Cap", "Shoes", "PetEquip", "Longcoat", "Shield",
		"Android", "Familiar", "Bits", "Mechanic", "Accessory", "Totem", "Hair", "TamingMob", "Coat", "Ring", "MonsterBattle",
		"Dragon", "Pants", "SkillSkin", "Cape",
	}
	itemSearches []string = []string{"Pet", "Consume", "Install", "Etc", "Cash"}
)

func TestFieldStatistics(t *testing.T) {
	done := make(chan bool, 1)
	go func() {
		nxfile.FieldStatistics(nxFileDir, nxfile.MapFilename, MapSearches, "portal", true)
		done <- true
	}()
	<-done
}

func TestExtractMap(t *testing.T) {
	done := make(chan bool, 1)
	go func() {
		nxfile.ExtractMap(nxFileDir)
		done <- true
	}()
	<-done
	m := nxfile.GetMap(10000)
	fmt.Println(m)
}
func TestExtractCharacter(t *testing.T) {
	done := make(chan bool, 1)
	go func() {
		nxfile.ExtractCharacter(nxFileDir)
		done <- true
	}()
	<-done
	equip := nxfile.GetEquip(1202212)
	fmt.Println(equip)
}

func TestExtractItem(t *testing.T) {
	done := make(chan bool, 1)
	go func() {
		nxfile.ExtractItem(nxFileDir)
		done <- true
	}()
	<-done
	item := nxfile.GetItem(5000055)
	fmt.Println(item)
}
