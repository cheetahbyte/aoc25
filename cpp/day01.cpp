#include <chrono>
#include <fstream>
#include <iostream>
#include <string>

constexpr int START_POSITION = 50;
constexpr int DIAL_SIZE = 100;

std::string readFile(const std::string& path) {
    std::ifstream t(path, std::ios::in | std::ios::binary);
    t.seekg(0, std::ios::end);
    const std::streamoff size = t.tellg();
    std::string buffer(static_cast<size_t>(size), '\0');
    t.seekg(0);
    t.read(buffer.data(), size);
    return buffer;
}

int pt1_from_buffer(const std::string& buffer) {
    int pos = START_POSITION;
    int hits = 0;

    const char* ptr = buffer.data();
    const char* end = ptr + buffer.size();

    while (ptr < end) {
        while (ptr < end && (*ptr == '\n' || *ptr == '\r' || *ptr == ' ' || *ptr == '\t')) {
            ++ptr;
        }
        if (ptr >= end) break;

        char direction = *ptr++;
        int sign = (direction == 'R') ? 1 : -1;

        int val = 0;
        while (ptr < end && *ptr >= '0' && *ptr <= '9') {
            val = val * 10 + (*ptr - '0');
            ++ptr;
        }

        pos += sign * val;

        pos %= DIAL_SIZE;
        if (pos < 0) pos += DIAL_SIZE;

        if (pos == 0) {
            ++hits;
        }

        while (ptr < end && (*ptr == '\n' || *ptr == '\r')) {
            ++ptr;
        }
    }

    return hits;
}

int main() {
    auto rawBuffer = readFile("./day01/input.txt");

    auto start = std::chrono::high_resolution_clock::now();

    int p1 = pt1_from_buffer(rawBuffer);

    auto end = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double, std::milli> elapsed = end - start;

    std::cout << "Part 1 was " << p1 << "\n";
    std::cout << "Time: " << elapsed.count() << " ms\n";

    return 0;
}
