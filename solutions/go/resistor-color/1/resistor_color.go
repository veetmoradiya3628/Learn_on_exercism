package resistorcolor

var colors = map[string]int{
    "black": 0,
    "brown": 1,
    "red": 2,
    "orange": 3,
    "yellow": 4,
    "green": 5,
    "blue": 6,
    "violet": 7,
    "grey": 8,
    "white": 9,
}

// Colors returns the list of all colors.
func Colors() []string {
	keys := make([]string, 0, len(colors))
    for k := range colors {
        keys = append(keys, k)
    }
    return keys
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for k, v := range colors {
        if k == color {
            return v
        }
    }
    return -1
}
