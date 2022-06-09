#include <iostream>
#include <stdio.h>
#include <vector>

#define LENGTH 64
#define DEBUG 0 

class Solution {
public:
    void remove(std::vector<int>& nums, int j) {
        nums.erase(nums.begin()+j, nums.begin()+j+1);
    }

    void print_pair(int i, int j) {
        std::cout << std::endl;
        std::cout << "Comparing: " << i << " " << j;
    }

    int removeDuplicates(std::vector<int>& nums) {
        int j=0;
        for(int i=0; i<(int)nums.size()+j-1; i++) {
            if(nums[i-j] == nums[i-j+1]) {
                if(DEBUG) { 
                    print_pair(nums[i-j], nums[i-j+1]); 
                }
                remove(nums, i-j);        
                j++;
            } 
        }
        
        return nums.size();
    }

};

void print_vector(std::vector<int> v) {
    for(int k: v) {
        std::cout << k << ' ';
    }
}

int main() {
    Solution s = Solution();
    std::vector r = std::vector<int>();
    srand(time(NULL));

    for(int j=0; j<LENGTH; j++) {
        r.push_back(rand() % 10);
    }

    std::vector n = {0,0,1,1,1,2,2,3,3,4};

    print_vector(r);
    s.removeDuplicates(r);
    std::cout << std::endl;
    print_vector(r);
}