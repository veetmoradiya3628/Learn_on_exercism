#include "atbash_cipher.h"
#include <string>
#include <cctype>

namespace atbash_cipher {

char invertChar(char c) {
    if (std::islower(c)) {
        return 'a' + 'z' - c;
    } else if (std::isupper(c)) {
        return 'a' + 'Z' - c; 
    }
    return c;
}

std::string encode(std::string original) {
    std::string processed = "";
    size_t count = 0;

    for (char c : original) {
        if (std::isalnum(static_cast<unsigned char>(c))) {
            if (count > 0 && count % 5 == 0) {
                processed += ' ';
            }
            
            if (std::isalpha(static_cast<unsigned char>(c))) {
                processed += invertChar(c);
            } else {
                processed += c; 
            }
            count++;
        }
    }
    return processed;
}

std::string decode(std::string encoded) {
    std::string decoded = "";
    for (char c : encoded) {
        if (std::isalnum(static_cast<unsigned char>(c))) {
            if (std::isalpha(static_cast<unsigned char>(c))) {
                decoded += invertChar(c);
            } else {
                decoded += c;
            }
        }
    }
    return decoded;
}

}  // namespace atbash_cipher
