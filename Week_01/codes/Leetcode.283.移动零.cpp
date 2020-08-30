#include <iostream>
#include <vector>

using namespace std;

void moveZeroes(vector<int>& nums) {
    int notZeroPosition = 0;
    for (int i = 0; i < nums.size(); i++) {
        if (nums[i] != 0) { // Condition.1
            nums[notZeroPosition] = nums[i];

            // i != notZeroPosition即发生了0和非0的位置调换
            // 由于此题特殊性, 调换的另一个数据总是0, 可以省去一个用来做变量调换的中间变量(仅使用常量0)
            if (i != notZeroPosition) {
                nums[i] = 0;
            } 
            
            notZeroPosition++; // Condition.1 满足即确定了一个非0位置, 目标idx后移 
        }
    }
}

int main() {

}