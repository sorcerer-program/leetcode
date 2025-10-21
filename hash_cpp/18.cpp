#include <vector>
#include <algorithm>

using namespace std;

vector<vector<int>> fourSum(vector<int> nums, int target) {
    int size = nums.size();
    vector<vector<int>> answer;
    if (size < 4) {
        return answer;
    }
    sort(nums.begin(), nums.end());
    int n = 0;
    for (int i = 0; i < size - 3; i++) {
        if (i > 0 && nums[i] == nums[i-1]) {
            continue;
        }
        for (int j = i + 1; j < size - 2; j++) {
            if (j > i + 1 && nums[j] == nums[j-1]) {
                continue;
            }
            int left = j + 1;
            int right = size - 1;
            while (left < right) {
                long long sum = (long long)nums[i] + (long long)nums[j] + (long long)nums[left] + (long long)(nums[right]);
                if (sum < target) {
                    left++;
                }else if (sum > target) {
                    right--;
                }else {
                    answer.push_back({nums[i], nums[j], nums[left], nums[right]});
                    while (left < right && nums[left] == nums[left+1]) {
                        left++;
                    }
                    while (left < right && nums[right] == nums[right-1]) {
                        right--;
                    }
                    left++;
                    right--;
                }
            }
        }
    }
    return answer;
}