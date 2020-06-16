class Solution {
public:
    int mySqrt(int x) {
        if (x == 0)
            return 0;
        double a = (double)x / 2;
        for (int i = 0; i < 50; i++) {
            a = 0.5*(x / a + a);
        }
        return (int)a;
    }
};
