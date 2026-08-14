#include "armstrong_numbers.h"
#include <string>
#include <cmath>

namespace armstrong_numbers {

bool is_armstrong_number(int num) {
    int length = std::to_string(num).length();

    int temp = num;
    int sum = 0;
    
    while (temp > 0) {
        int digit = temp % 10;          
        sum += std::pow(digit, length); 
        temp /= 10;                     
    }
    return sum == num;
}

}  // namespace armstrong_numbers
