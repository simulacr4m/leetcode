class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        vector<int> pair;
        for (int i = 0; i < numbers.size(); i++) {
            int v = target - numbers[i];
            int lo = i+1, hi = numbers.size()-1;
            while (lo <= hi) {
                int mid = (lo + hi) / 2;
                if (numbers[mid] < v)
                    lo = mid+1;
                else if (numbers[mid] > v)
                    hi = mid-1;
                else {
                    pair.push_back(i+1);
                    pair.push_back(mid+1);
                    return pair;
                }
            }
        }
        return pair;
    }
};
