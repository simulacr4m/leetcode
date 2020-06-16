class Solution {
public:
    int strStr(string haystack, string needle) {
        if (needle == "")
            return 0;
        int n = haystack.length();
        int m = needle.length();
        
        int right[256];
        for (int c = 0; c < 256; c++)
            right[c] = -1;
        for (int j = 0; j < m; j++)
            right[needle[j]] = j;
        
        int skip = 0;
        for (int i = 0; i <= n-m; i += skip) {
            skip = 0;
            for (int j = m-1; j >= 0; j--) {
                if (needle[j] != haystack[i+j]) {
                    skip = j - right[haystack[i+j]];
                    if (skip < 1)
                        skip = 1;
                    break;
                }
            }
            if (skip == 0)
                return i;
        }
        return -1;
    }
};
