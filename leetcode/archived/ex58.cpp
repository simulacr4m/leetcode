class Solution {
public:
    int lengthOfLastWord(string s) {
        if (s.length() == 0)
            return s.length();
        int i, k;
        for (i = s.length()-1; i >= 0 && s.at(i) == ' '; i--)
            ;
        for (k = i; k >= 0; k--)
            if (s.at(k) == ' ')
                break;
        return i - k;
    }
};
