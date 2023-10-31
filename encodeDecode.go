package main

import ("fmt")

func encode(s []string) []string {
    for i, str := range s {
        s[i] = fmt.Sprintf("%d#%s", len(str), str)
    }
    return s
}