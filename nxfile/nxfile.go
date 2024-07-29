package nxfile

import (
	"goms/nx"
	"log/slog"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Hucaru/gonx"
)

const (
	MapFilename       string = "Map.nx"
	CharacterFilename string = "Character.nx"
	ItemFilename      string = "Item.nx"
)

var (
	maps   map[uint32]*nx.MapNX       = make(map[uint32]*nx.MapNX)
	equips map[uint32]*nx.CharacterNX = make(map[uint32]*nx.CharacterNX)
	items  map[uint32]*nx.ItemNX      = make(map[uint32]*nx.ItemNX)
)

func GetMap(id uint32) *nx.MapNX {
	m, ok := maps[id]
	if !ok {
		slog.Error("Failed to get map info", "id", id)
		return nil
	}
	return m
}

func GetEquip(id uint32) *nx.CharacterNX {
	equip, ok := equips[id]
	if !ok {
		slog.Warn("Failed to get equip info", "id", id)
		return nil
	}
	return equip
}

func GetItem(id uint32) *nx.ItemNX {
	item, ok := items[id]
	if !ok {
		slog.Error("Failed to get item info", "id", id)
		return nil
	}
	return item
}

func getIDFromNodeName(nodeName string) uint32 {
	name := strings.TrimSuffix(nodeName, filepath.Ext(nodeName))
	id, err := strconv.Atoi(name)
	if err != nil {
		slog.Warn("Failed to get id from node name", "node", nodeName)
		return 0
	}
	return uint32(id)
}

func ExtractMap(dir string) {
	searchs := []string{"Map/Map0"}
	rootNodes, textLookup := extractFile(dir, MapFilename)
	for _, search := range searchs {
		valid := gonx.FindNode(search, rootNodes, textLookup, func(currentNode *gonx.Node) {
			for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
				subNode := rootNodes[currentNode.ChildID+i]
				name := textLookup[subNode.NameID]
				var data nx.MapNX
				// Info
				infoSearch := search + "/" + name + "/info"
				valid := gonx.FindNode(infoSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
					nx.MappingField(&data.Info, currentNode, rootNodes, textLookup)
				})
				if !valid {
					slog.Warn("Failed to find infoNode", "infoSearch", infoSearch)
					continue
				}
				// Life
				lifeSearch := search + "/" + name + "/life"
				_ = gonx.FindNode(lifeSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
					for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
						subNode := rootNodes[currentNode.ChildID+i]
						name := textLookup[subNode.NameID]
						lifeSubSearch := lifeSearch + "/" + name
						valid = gonx.FindNode(lifeSubSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
							var life nx.Life
							nx.MappingField(&life, currentNode, rootNodes, textLookup)
							data.Lifes = append(data.Lifes, life)
						})
						if !valid {
							slog.Warn("Failed to find lifeSubNode", "lifeSubSearch", lifeSubSearch)
							continue
						}
					}
				})
				// Portal
				portalSearch := search + "/" + name + "/portal"
				_ = gonx.FindNode(portalSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
					for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
						subNode := rootNodes[currentNode.ChildID+i]
						name := textLookup[subNode.NameID]
						portalSubSearch := portalSearch + "/" + name
						valid = gonx.FindNode(portalSubSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
							var portal nx.Portal
							nx.MappingField(&portal, currentNode, rootNodes, textLookup)
							data.Portals = append(data.Portals, portal)
						})
						if !valid {
							slog.Warn("Failed to find portalSubNode", "portalSubSearch", portalSubSearch)
							continue
						}
					}
				})
				id := getIDFromNodeName(name)
				if id == 0 {
					continue
				}
				maps[id] = &data
			}
		})
		if !valid {
			slog.Warn("Failed to find node", "search", search)
			continue
		}
	}
	slog.Info("Load nxfile end", "nxfile", MapFilename)
}

func ExtractCharacter(dir string) {
	searchs := []string{"Weapon", "Glove", "Face", "Cap", "Shoes", "PetEquip", "Longcoat", "Shield",
		"Android", "Familiar", "Bits", "Mechanic", "Accessory", "Totem", "Hair", "TamingMob", "Coat", "Ring", "MonsterBattle",
		"Dragon", "Pants", "SkillSkin", "Cape",
	}
	rootNodes, textLookup := extractFile(dir, CharacterFilename)
	for _, search := range searchs {
		valid := gonx.FindNode(search, rootNodes, textLookup, func(currentNode *gonx.Node) {
			for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
				subNode := rootNodes[currentNode.ChildID+i]
				name := textLookup[subNode.NameID]
				infoSearch := search + "/" + name + "/info"
				var data nx.CharacterNX
				valid := gonx.FindNode(infoSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
					nx.MappingField(&data, currentNode, rootNodes, textLookup)
				})
				if !valid {
					slog.Warn("Failed to find subNode", "infoSearch", infoSearch)
					continue
				}
				id := getIDFromNodeName(name)
				if id == 0 {
					continue
				}
				equips[id] = &data
			}
		})
		if !valid {
			slog.Warn("Failed to find node", "search", search)
			continue
		}
	}
	slog.Info("Load nxfile end", "nxfile", CharacterFilename)
}

func ExtractItem(dir string) {
	rootNodes, textLookup := extractFile(dir, ItemFilename)
	extracItemPet(rootNodes, textLookup)
	extracItemConsume(rootNodes, textLookup)
	extracItemOther(rootNodes, textLookup)
	slog.Info("Load nxfile end", "nxfile", ItemFilename)
}

func extracItemPet(rootNodes []gonx.Node, textLookup []string) {
	search := "Pet"
	valid := gonx.FindNode(search, rootNodes, textLookup, func(currentNode *gonx.Node) {
		for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
			subNode := rootNodes[currentNode.ChildID+i]
			name := textLookup[subNode.NameID]
			infoSearch := search + "/" + name + "/info"
			var data nx.ItemNX
			valid := gonx.FindNode(infoSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
				nx.MappingField(&data, currentNode, rootNodes, textLookup)
			})
			if !valid {
				slog.Warn("Failed to find subNode", "infoSearch", infoSearch)
				continue
			}
			id := getIDFromNodeName(name)
			if id == 0 {
				continue
			}
			items[id] = &data
		}
	})
	if !valid {
		slog.Warn("Failed to find search", "search", search)
	}
}

func extracItemConsume(rootNodes []gonx.Node, textLookup []string) {
	search := "Consume"
	valid := gonx.FindNode(search, rootNodes, textLookup, func(currentNode *gonx.Node) {
		for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
			subNode := rootNodes[currentNode.ChildID+i]
			groupName := textLookup[subNode.NameID]
			subSearch := search + "/" + groupName
			valid := gonx.FindNode(subSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
				for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
					subNode := rootNodes[currentNode.ChildID+i]
					name := textLookup[subNode.NameID]
					infoSearch := subSearch + "/" + name + "/info"
					var data nx.ItemNX
					// Info Search
					valid := gonx.FindNode(infoSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
						nx.MappingField(&data, currentNode, rootNodes, textLookup)
					})
					if !valid {
						slog.Warn("Failed to find search", "infoSearch", infoSearch)
						continue
					}
					// Spec search
					specSearch := subSearch + "/" + name + "/spec"
					_ = gonx.FindNode(specSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
						nx.MappingField(&data, currentNode, rootNodes, textLookup)
					})
					id := getIDFromNodeName(name)
					if id == 0 {
						continue
					}
					items[id] = &data
				}
			})
			if !valid {
				slog.Warn("Failed to find search", "subSearch", subSearch)
				continue
			}
		}
	})
	if !valid {
		slog.Warn("Failed to find search", "search", search)
	}
}

func extracItemOther(rootNodes []gonx.Node, textLookup []string) {
	searches := []string{"Install", "Etc", "Cash"}
	for _, search := range searches {
		valid := gonx.FindNode(search, rootNodes, textLookup, func(currentNode *gonx.Node) {
			for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
				subNode := rootNodes[currentNode.ChildID+i]
				groupName := textLookup[subNode.NameID]
				subSearch := search + "/" + groupName
				valid := gonx.FindNode(subSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
					for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
						subNode := rootNodes[currentNode.ChildID+i]
						name := textLookup[subNode.NameID]
						infoSearch := subSearch + "/" + name + "/info"
						var data nx.ItemNX
						// Info Search
						valid := gonx.FindNode(infoSearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
							nx.MappingField(&data, currentNode, rootNodes, textLookup)
						})
						if !valid {
							slog.Warn("Failed to find search", "infoSearch", infoSearch)
							continue
						}
						id := getIDFromNodeName(name)
						if id == 0 {
							continue
						}
						items[id] = &data
					}
				})
				if !valid {
					slog.Warn("Failed to find search", "subSearch", subSearch)
					continue
				}
			}
		})
		if !valid {
			slog.Warn("Failed to find search", "search", search)
			continue
		}
	}
}
