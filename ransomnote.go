package main



func CanConstruct(ransomNote, magazine string) bool {
	noteMap := make(map[rune]int)
    magazineMap := make(map[rune]int)

    for _, v := range ransomNote {
        noteMap[v]++
    }

    for _, v := range magazine {
        magazineMap[v]++
    }

    for k, v := range noteMap {
        if magazineMap[k] < v {
            return false
        }
    }

    return true
}