#include "allergies.h"

#include <vector>

namespace allergies {

namespace {

const std::vector<std::pair<item, std::string>> ALL_ITEMS = {
    {item::eggs, "eggs"},
    {item::peanuts, "peanuts"},
    {item::shellfish, "shellfish"},
    {item::strawberries, "strawberries"},
    {item::tomatoes, "tomatoes"},
    {item::chocolate, "chocolate"},
    {item::pollen, "pollen"},
    {item::cats, "cats"}
};

}  // namespace

allergy_test::allergy_test(int score) : score_(score) {}

bool allergy_test::is_allergic_to(item allergy_item) const {
    return (score_ & static_cast<int>(allergy_item)) != 0;
}

bool allergy_test::is_allergic_to(const std::string& item_name) const {
    for (const auto& [allergy_enum, name] : ALL_ITEMS) {
        if (name == item_name) {
            return is_allergic_to(allergy_enum);
        }
    }
    return false;
}

std::unordered_set<std::string> allergy_test::get_allergies() const {
    std::unordered_set<std::string> result;
    for (const auto& [allergy_enum, name] : ALL_ITEMS) {
        if (is_allergic_to(allergy_enum)) {
            result.insert(name);
        }
    }
    return result;
}

}  // namespace allergies