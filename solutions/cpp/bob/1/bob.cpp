#include "bob.h"
#include <algorithm>
#include <cctype>

namespace bob {

std::string hey(const std::string& utterance) {
    size_t first = utterance.find_first_not_of(" \t\n\r");
    if (first == std::string::npos) {
        return "Fine. Be that way!";
    }
    
    size_t last = utterance.find_last_not_of(" \t\n\r");
    std::string cleaned = utterance.substr(first, last - first + 1);

    bool is_question = (cleaned.back() == '?');
    bool has_letters = std::any_of(cleaned.begin(), cleaned.end(), [](unsigned char c) { 
        return std::isalpha(c); 
    });
    bool is_yelling = has_letters && std::none_of(cleaned.begin(), cleaned.end(), [](unsigned char c) { 
        return std::islower(c); 
    });

    if (is_yelling && is_question) {
        return "Calm down, I know what I'm doing!";
    }
    if (is_yelling) {
        return "Whoa, chill out!";
    }
    if (is_question) {
        return "Sure.";
    }

    return "Whatever.";
}

}  // namespace bob
