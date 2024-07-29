package nxfile

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/Hucaru/gonx"
)

func extractFile(dir, filename string) ([]gonx.Node, []string) {
	path := path.Join(dir, filename)
	slog.Info("Loading nxfile...", "path", path)
	rootNodes, textLookup, _, _, err := gonx.Parse(path)
	if err != nil {
		slog.Error("Faild to parse nx file", "path", path)
		return nil, nil
	}
	return rootNodes, textLookup
}

func FieldStatistics(dir, filename string, searches []string, keyword string, isArray bool) {
	rootNodes, textLookup := extractFile(dir, filename)
	if rootNodes == nil {
		return
	}
	done := false
	temp := make(map[string]int)
	for _, search := range searches {
		deepSearch(search, keyword, rootNodes, textLookup, isArray, func(keywordNode *gonx.Node) {
			for i := uint32(0); i < uint32(keywordNode.ChildCount); i++ {
				option := rootNodes[keywordNode.ChildID+i]
				optionName := textLookup[option.NameID]
				if optionName == "id" {
					value := gonx.DataToInt32(option.Data)
					if value == 18191 {
						done = true
					}
				}
				if done {
					value := gonx.DataToInt32(option.Data)
					slog.Info("test", "optionName", optionName, "value", value)
				}
				temp[optionName]++
			}
			done = false

		})
	}
	length := len(temp)
	words := make([]string, length)
	index := 0
	for name, count := range temp {
		field := strings.ToUpper(name[:1]) + name[1:]
		words[index] = fmt.Sprintf("%s any // count=%d", field, count)
		index++
	}
	sort.Strings(words)
	path := filename + ".txt"
	file, err := os.Create(path)
	if err != nil {
		slog.Error("Failed to create file", "err", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(strings.Join(words, "\r\n"))
	if err != nil {
		slog.Error("Failed to write string", "err", err)
		return
	}
	slog.Debug("Test statistics ok")
}

func deepSearch(search string, keyword string, rootNodes []gonx.Node, textLookup []string, isArray bool, callback func(n *gonx.Node)) {
	trySearch := search + "/" + keyword
	valid := gonx.FindNode(trySearch, rootNodes, textLookup, func(currentNode *gonx.Node) {
		if !isArray {
			callback(currentNode)
			return
		} else {
			// Expand all child nodes
			for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
				subNode := rootNodes[currentNode.ChildID+i]
				name := textLookup[subNode.NameID]
				nextSearch := trySearch + "/" + name
				valid := gonx.FindNode(nextSearch, rootNodes, textLookup, func(childNode *gonx.Node) {
					callback(childNode)
				})
				if !valid {
					slog.Warn("Failed to find more child node", "nextSearch", nextSearch)
					continue
				}
			}

		}
	})
	if !valid {
		valid = gonx.FindNode(search, rootNodes, textLookup, func(currentNode *gonx.Node) {
			for i := uint32(0); i < uint32(currentNode.ChildCount); i++ {
				subNode := rootNodes[currentNode.ChildID+i]
				name := textLookup[subNode.NameID]
				nextSearch := search + "/" + name
				deepSearch(nextSearch, keyword, rootNodes, textLookup, isArray, func(n *gonx.Node) {
					callback(n)
				})
			}
		})
		if !valid {
			slog.Warn("Failed to find more node", "search", search)
		}
	}
}
