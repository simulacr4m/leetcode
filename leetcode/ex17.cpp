class Solution {
private:
    double newton_method(double x) {
        double a = x / 2;
        for (int i = 0; i < 100; i++) {
            a = 0.5*(x / a + a);
        }
        return a;
    }
public:
    int mySqrt(int x) {
        if (x == 0) return 0;
        double v = newton_method((double) x);
        return (int) v;
    }
};
