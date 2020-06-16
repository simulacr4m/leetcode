class Solution {
public:
    int reverse(int x) {
        int r = 0;
        while (x != 0) {
            int d = x % 10;
            x /= 10;
            if (r > INT_MAX/10 || (r == INT_MAX / 10 && d > 7))
                return 0;
            if (r < INT_MIN/10 || (r == INT_MIN / 10 && d < -8))
                return 0;
            r = 10*r + d;
        }
        return r;
    }
};
