#include <vector>

int Part2(std::vector<std::string> lines) {
    return lines.size();
}

extern "C" {
    int Part2_Bridge(char** rawStrings, int count) {
        std::vector<std::string> vec;
        vec.reserve(count);
        for (int i = 0; i < count; i++)
            vec.push_back(std::string(rawStrings[i]));

        return Part2(vec);
    }
}
