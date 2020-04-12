func smashStones(x int, y int) int {
    if x <= y {
        return y-x
    }
    return smashStones(y, x)
}

func lastStoneWeight(stones []int) int {
    if len(stones) == 1 {
        return stones[0]
    }
    for i := 0; i < len(stones); i++ {
        sort.Ints(stones)
        smashed := smashStones(stones[len(stones)-1], stones[len(stones)-2])
        stones[len(stones)-1] = smashed
        stones[len(stones)-2] = 0
    }
    return stones[len(stones)-1]
}
