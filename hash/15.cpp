#include <vector>
#include <algorithm>

using namespace std;

vector<vector<int>> threeSum(vector<int>& nums) {
    sort(nums.begin(), nums.end());
    vector<vector<int>> answer;
    int n = nums.size();
    for (int i = 0; i < n; i++) {
        if (nums[i] > 0) {
            return answer;
        }
        while (nums[i] == nums[i-1]) {
            continue;
        }
        int left = i + 1;
        int right = n - 1;
        while (left < right) {
            int sum = nums[i] + nums[left] + nums[right];
            if (sum < 0) {
                while (nums[left] == nums[left+1]) {
                    left++;
                }
                left++;
            } else if (sum > 0) {
                while (nums[right] == nums[right-1]) {
                    right--;
                }
                right--;
            }
        }
    }
    return answer;
}