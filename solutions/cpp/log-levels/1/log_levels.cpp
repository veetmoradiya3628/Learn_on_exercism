#include <string>
#include <algorithm>
#include <cctype>

namespace log_line {

// Extracts the message after "[LEVEL]: " and trims leading/trailing whitespace
std::string message(std::string line) {
    size_t colon_pos = line.find(':');
    if (colon_pos == std::string::npos) return line;

    std::string msg = line.substr(colon_pos + 1);

    // Trim leading whitespace
    size_t start = msg.find_first_not_of(" \t\r\n");
    if (start != std::string::npos) {
        msg = msg.substr(start);
    }

    // Trim trailing whitespace
    size_t end = msg.find_last_not_of(" \t\r\n");
    if (end != std::string::npos) {
        msg = msg.substr(0, end + 1);
    }

    return msg;
}

// Extracts log level in lowercase (e.g., "error", "warning", "info")
std::string log_level(std::string line) {
    size_t start = line.find('[');
    size_t end = line.find(']');
    
    if (start != std::string::npos && end != std::string::npos && end > start) {
        std::string level = line.substr(start + 1, end - start - 1);
        return level;
    }
    
    return "";
}

// Combines into: "Message content (level)"
std::string reformat(std::string line) {
    return message(line) + " (" + log_level(line) + ")";
}

}  // namespace log_line