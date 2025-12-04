#include <fstream>
#include <iostream>
#include <vector>
#include <string>
#include <chrono>

#define vec2char std::vector<std::vector<char>>

vec2char parseIntoGrid() {
    std::ifstream file("./day04/input.txt");
    std::string line;
    vec2char vec;
    while(std::getline(file, line)) {
        if (!line.empty()) {
            vec.emplace_back(line.begin(), line.end());
        }
    }
    return vec;
}

int part1(const vec2char& grid) {
    int rows = grid.size();
    int cols = grid[0].size();
    int hits = 0;
    int dr[] = {-1, -1, -1,  0,  0,  1,  1,  1};
    int dc[] = {-1,  0,  1, -1,  1, -1,  0,  1};

    for (int r = 0; r < rows; ++r) {
        for (int c = 0; c < cols; ++c) {
            if (grid[r][c] != '@') continue;

            int neighborCount = 0;
            for (int i = 0; i<8; ++i) {
                int nr = r + dr[i];
                int nc = c + dc[i];
                if (nr >= 0 && nr < rows && nc >= 0 && nc < cols) {
                    if (grid[nr][nc] == '@') neighborCount++;
                }
            }
            if (neighborCount < 4) hits++;
        }
    }
    return hits;
}

int part2_original(vec2char grid) {
    int rows = grid.size();
    int cols = grid[0].size();
    int totalRemoved = 0;
    int dr[] = {-1, -1, -1,  0,  0,  1,  1,  1};
    int dc[] = {-1,  0,  1, -1,  1, -1,  0,  1};

    std::vector<std::pair<int, int>> candidates;
    for(int r = 0; r < rows; ++r) {
        for(int c = 0; c < cols; ++c) {
            if(grid[r][c] == '@') candidates.push_back({r, c});
        }
    }

    std::vector<std::vector<int>> checkTags(rows, std::vector<int>(cols, 0));
    int currentToken = 1;

    while (!candidates.empty()) {
        std::vector<std::pair<int, int>> toRemove;
        std::vector<std::pair<int, int>> nextCandidates;

        for (const auto& p : candidates) {
            int r = p.first;
            int c = p.second;
            int neighborCount = 0;
            for (int i = 0; i < 8; ++i) {
                int nr = r + dr[i];
                int nc = c + dc[i];
                if (nr >= 0 && nr < rows && nc >= 0 && nc < cols) {
                    if (grid[nr][nc] == '@') neighborCount++;
                }
            }
            if (neighborCount < 4) toRemove.push_back(p);
        }

        if (toRemove.empty()) break;
        totalRemoved += toRemove.size();
        currentToken++;

        for (const auto& p : toRemove) {
            int r = p.first;
            int c = p.second;
            grid[r][c] = '.';
            for (int i = 0; i < 8; ++i) {
                int nr = r + dr[i];
                int nc = c + dc[i];
                if (nr >= 0 && nr < rows && nc >= 0 && nc < cols) {
                    if (grid[nr][nc] == '@' && checkTags[nr][nc] != currentToken) {
                        checkTags[nr][nc] = currentToken;
                        nextCandidates.push_back({nr, nc});
                    }
                }
            }
        }
        candidates = std::move(nextCandidates);
    }
    return totalRemoved;
}

int main() {
    auto vec = parseIntoGrid();

    auto start1 = std::chrono::high_resolution_clock::now();
    auto p1 = part1(vec);
    auto end1 = std::chrono::high_resolution_clock::now();

    auto start2 = std::chrono::high_resolution_clock::now();
    auto p2 = part2_original(vec);
    auto end2 = std::chrono::high_resolution_clock::now();

    auto dur1 = std::chrono::duration_cast<std::chrono::microseconds>(end1 - start1);
    auto dur2_orig = std::chrono::duration_cast<std::chrono::microseconds>(end2 - start2);

    std::cout << "Part 1: " << p1 << " (" << dur1.count() << " us)" << std::endl;
    std::cout << "----------------" << std::endl;
    std::cout << "Part 2 (Original) : " << p2 << " (" << dur2_orig.count() << " us)" << std::endl;

    return 0;
}
