class Solution {
private:
    int convert2D(int row, int col, int n) {
        return n*row + col;
    }
    
    int *convert1D(int v, int n) {
        static int coords[2];
        int row = v / n;
        int col = v % n;
        coords[0] = row;
        coords[1] = col;
        return coords;
    }
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        if (matrix.size() == 0 || matrix.at(0).size() == 0)
            return false;
        int m = matrix.size();
        int n = matrix.at(0).size();
        
        int lo = 0, hi = convert2D(m-1, n-1, n);
        while (lo <= hi) {
            int mid = (lo + hi) / 2;
            int *coords = convert1D(mid, n);
            int row = coords[0], col = coords[1];
            if (matrix.at(row).at(col) < target)
                lo = mid+1;
            else if (matrix.at(row).at(col) > target)
                hi = mid-1;
            else
                return true;
        }
        return false;
    }
};
