#include "reverse_string.h"
#include <vector>
#include <cstring>

namespace reverse_string {

std::string reverse_string(std::string original) {
    int l = original.length();
    
    std::vector<char> arr(l + 1); 
    std::strcpy(arr.data(), original.c_str());

    for(int i = 0; i < l / 2; i++){
        char t = arr[i];
        arr[i] = arr[l - i - 1];
        arr[l - i - 1] = t;
    }
    
    std::string ans(arr.data(), l);
    return ans;
}    

}  // namespace reverse_string
