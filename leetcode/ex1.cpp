class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> pair;
        unordered_map<int, int> table;
        
        for (int i = 0; i < nums.size(); i++) {
            if (table.find(target - nums[i]) != table.end()) {
                pair.push_back(i);
                pair.push_back(table[target - nums[i]]);
                break;
            }
            table[nums[i]] = i;
        }
        
        return pair;
    }
};
