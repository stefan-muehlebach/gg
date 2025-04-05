package colors

import (
	"testing"
)

// Check if every colorname in the slice 'Names' is also found
// in the map 'Map'.
func TestNames(t *testing.T) {
	for _, colorName := range Names {
		if _, ok := Map[colorName]; !ok {
			t.Errorf("Colorname '%s': missing in 'Map'.", colorName)
		}
	}
}

// Check if every colorname in the map 'Map' is also found
// in the slice 'Names'.
func TestMap(t *testing.T) {
	for colorName, _ := range Map {
		nameFound := false
		for _, name := range Names {
			if colorName == name {
				nameFound = true
				break
			}
		}
		if !nameFound {
			t.Errorf("Colorname '%s': missing in 'Names'.", colorName)
		}
	}
}

// Check if every colorname in the map 'Groups' is also found
// in the map 'Map'.
func TestGroups(t *testing.T) {
	for group, colorList := range Groups {
		for _, colorName := range colorList {
			if _, ok := Map[colorName]; !ok {
				t.Errorf("Colorname '%s' in group '%v': missing in 'Map'.",
					colorName, group)
			}
		}
	}
}

// Check if every group name can be converted between string and ColorGroup.
func TestColorGroup(t *testing.T) {
    var group, g ColorGroup
    var str string
    var err error

    for group = 0; group < NumColorGroups; group++ {
        str = group.String()
        if err = g.Set(str); err != nil {
            t.Errorf("Color group name '%s': not convertable back", str)
        }
        if g != group {
            t.Errorf("Color group name '%s': wrong value", str)
        }
    }
}
