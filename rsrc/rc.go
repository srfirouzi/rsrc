package rsrc

import (
	"strconv"
	"strings"
)

// RCItem is item of rc file
type RCItem struct {
	ID       uint16
	Mode     string
	FileName string
}

// RCFile is rc file
type RCFile struct {
	Arch       string
	FileOutput string
	Items      []*RCItem
}

// NewRCFile make new RCFile
func NewRCFile(arch string, file string) *RCFile {
	out := new(RCFile)
	out.Arch = arch
	out.FileOutput = file
	out.Items = make([]*RCItem, 0)
	return out
}

// Add add item to rcfile
func (obj *RCFile) Add(item *RCItem) {
	if item != nil {
		obj.Items = append(obj.Items, item)
	}
}

// AddLine add item to rcfile
func (obj *RCFile) AddLine(str string) {
	item := NewRCitem(str)
	if item != nil {
		obj.Items = append(obj.Items, item)
	}
}

// AddLines add item to rcfile
func (obj *RCFile) AddLines(strs []string) {
	for i := 0; i < len(strs); i++ {
		obj.AddLine(strs[i])
	}
}

// NewRCitem make new RCitem
func NewRCitem(line string) *RCItem {
	parts := strings.Fields(line)
	if len(parts) == 3 {
		number, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil
		}
		mode := "DATA"
		if parts[1] == "ICON" || parts[1] == "3" {
			mode = "ICON"
		}
		if parts[1] == "RT_MANIFEST" || parts[1] == "MANIFEST" || parts[1] == "24" {
			mode = "MANIFEST"
		}
		out := new(RCItem)
		out.ID = (uint16)(number)
		out.Mode = mode
		out.FileName = strings.Trim(parts[2], "\"")
		return out
	}
	return nil

}
