
// copy the array size (dp memo)
// go backwards through array (we are iterating indices)

// while stack is not empty and current num > num on top
// dp[i] = max(dp[i] + 1, dp[top])
// pop

// every iteration push the index into the stack
// every iteration check for max
 
class Solution {
    public:
    int totalSteps(std::vector<int>& nums) {

        std::vector<int> dp(nums.size());
        std::stack<int> maxStack;
        int numSteps = 0;

        for(int i = nums.size() - 1;  i >= 0; i--) {
            while(!maxStack.empty() && nums[i] > nums[maxStack.top()]) {
                dp[i] = std::max(dp[maxStack.top()], dp[i] + 1);
                maxStack.pop();
            }
            maxStack.push(i);
            numSteps = std::max(numSteps, dp[i]);
        }


        return numSteps;
    }




}