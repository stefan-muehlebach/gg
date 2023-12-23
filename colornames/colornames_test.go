package colornames

import (
    "testing"
)

// Check if every colorname in the slice 'Names' is also found
// in the map 'Map'.
func TestSlice(t *testing.T) {
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

