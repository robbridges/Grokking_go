package main

func LongestConsecutiveLetters(s string) int {
	if len(s) == 0 {return 0}
    
    var res int
    // set your counter to 1, every letter will have a rank of at least one also set the precharacter to s[0]
    counter, preChar := 1, s[0]
    // initialize the pointer to the second digit in the string, check to see if hte one before it is a match
    for i := 1; i < len(s); i++ {
        if s[i] == preChar {
            counter++
            // if your current counter is greater than your res set the response to the current count
            if counter > res {res = counter}
        } else {
			// if not a match reset the counter to one, and move the precharacter forward one, i will be increted
            counter = 1
            preChar = s[i]
        }
    }
    // again check to see if the counter is greater than res before returning
    if counter > res {res = counter}
    
    return res
	
}