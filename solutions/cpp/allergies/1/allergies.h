#pragma once

#include <string>
#include <unordered_set>

namespace allergies {

enum class item {
    eggs = 1,
    peanuts = 2,
    shellfish = 4,
    strawberries = 8,
    tomatoes = 16,
    chocolate = 32,
    pollen = 64,
    cats = 128
};

class allergy_test {
public:
    explicit allergy_test(int score);

    bool is_allergic_to(const std::string& item_name) const;
    bool is_allergic_to(item allergy_item) const;
    
    std::unordered_set<std::string> get_allergies() const;

private:
    int score_;
};

}  // namespace allergies