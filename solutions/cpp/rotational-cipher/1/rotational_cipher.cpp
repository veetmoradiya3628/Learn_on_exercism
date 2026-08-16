#include "rotational_cipher.h"
#include <cctype>
namespace rotational_cipher {

std::string rotate(const std::string& input, int shift_key) {
    int shift = (shift_key % 26 + 26) % 26;
    std::string result = "";
    result.reserve(input.length()); 
    
    for (char c : input) {
        if (std::islower(static_cast<unsigned char>(c))) {
            result += static_cast<char>('a' + (c - 'a' + shift) % 26);
        } 
        else if (std::isupper(static_cast<unsigned char>(c))) {
            result += static_cast<char>('A' + (c - 'A' + shift) % 26);
        } 
        else {
            result += c;
        }
    }
    return result;
}
    
}  // namespace rotational_cipher
