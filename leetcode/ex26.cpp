class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.size() <= 1)
            return nums.size();
        int end = 0, pos = 1;
        while (end < nums.size()) {
            int begin = end;
            while (end < nums.size() && nums[end] == nums[begin])
                ++end;
            if (end < nums.size()) {
                nums[pos] = nums[end];
                ++pos;
            }
        }
        return pos;
    }
};
