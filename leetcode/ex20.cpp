class Solution {
public:
    bool isValid(string s) {
        stack<char> st;
        for (int i = 0; i < s.length(); i++) {
            if (s[i] == ')' || s[i] == '}' || s[i] == ']') {
                if (st.empty())
                    return false;
                char value = st.top();
                st.pop();
                if (value != '(' && s[i] == ')') return false;
                if (value != '{' && s[i] == '}') return false;
                if (value != '[' && s[i] == ']') return false;
            } else {
                st.push(s[i]);
            }
        }
        return st.empty();
    }
};
