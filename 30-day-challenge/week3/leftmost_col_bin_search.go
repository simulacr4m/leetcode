/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */

func binarySearch(binaryMatrix BinaryMatrix, rowSize, row int) int {
    lo, hi := 0, rowSize-1
    for lo <= hi {
        mid := lo + (hi - lo) / 2
        element := binaryMatrix.Get(row, mid)
        if element < 1 {
            lo = mid+1
        } else if element > 1 {
            hi = mid-1
        } else if lo != mid {
            hi = mid
        } else {
            return mid
        }
    }
    return -1
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
    dimensionTuple := binaryMatrix.Dimensions()
    n, m := dimensionTuple[0], dimensionTuple[1]
    minCol := -1
    for i := 0; i < n; i++ {
        index := binarySearch(binaryMatrix, m, i)
        if index == -1 {
            continue
        }
        if minCol > index || minCol == -1 {
            minCol = index
        }
    }
    return minCol
}
