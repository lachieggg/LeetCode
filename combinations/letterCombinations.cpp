#include <iostream>
#include <string>
#include <vector>
#include <map>

#define LENGTH 26  // length of the alphabet
#define FIRST  97  // ascii for 'a'
#define BUTTON 3   // number of letters per button

class Solution {
public:
    std::map<char, int> mapping;

    // Print out a mapping
    void print_mapping(std::map<char, int> mapping) {
        for(auto pair = mapping.cbegin(); pair != mapping.cend(); ++pair) {
            std::cout << pair->first << " " << pair->second << "\n";
        }
    }

    // Constructor
    Solution() {
        int j=1;
        for(int i=FIRST; i<FIRST+LENGTH; i++) {
            this->mapping.insert(std::pair<char, int>((char)i, j));
            if(i%BUTTON == 0 && i > 0) {
                j++;
            }
        }

        print_mapping(mapping);
    }

    void algorithm(std::vector<std::string>& v, int digit) {
        // Take every string in vector
        // Create 3 of them
        // Then append the different characters to each of them
        char characters[] = {'a', 'b', 'c'};
        (void)characters;
    }

    // Solution
    std::vector<std::string> letterCombinations(std::string digits) {
        std::vector<std::string> vect = std::vector<std::string>();
        std::vector<char> set = std::vector<char>();

        for(auto digit: digits) {
            for(auto pair = this->mapping.cbegin(); pair != this->mapping.cend(); ++pair) {
                if(pair->second == digit) {
                    set.push_back(pair->first);
                }

                for(int i=0; i<vect.size(); i++) {
                    vect[i] = vect[i];
                }
            }

            std::cout << digit << std::endl;

        }

        return vect;
    }
};

int main() {
    Solution s = Solution();
    s.letterCombinations("ag");
}
    