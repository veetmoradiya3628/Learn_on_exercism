#pragma once
#include <string>

namespace atbash_cipher {
std::string encode(std::string original);
std::string decode(std::string encoded);
}  // namespace atbash_cipher
