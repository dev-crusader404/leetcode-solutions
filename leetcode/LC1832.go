package leetcode

func checkIfPangram(sentence string) bool {
    if len(sentence) == 0 || len(sentence) < 26 { return false }
    counter := make(map[rune]struct{})

    for _, c := range sentence {
        counter[c] = struct{}{}
    }
    return len(counter) == 26
}