
#include <algorithm>
#include <fstream>
#include <iostream>
#include <ranges>
#include <string>
#include <string_view>
#include <vector>

#include <cctype>

std::string trim(const std::string &s) {
  size_t start = 0, end = s.size();
  while (start < end && std::isspace(static_cast<unsigned char>(s[start])))
    ++start;
  while (end > start && std::isspace(static_cast<unsigned char>(s[end - 1])))
    --end;
  return s.substr(start, end - start);
}

struct Range {
  long long start, end;
};

struct ParseResult {
  std::vector<long long> ingredients;
  std::vector<Range> ranges;
};

std::vector<std::string> split(std::string_view s) {
  auto split_view = s | std::views::split('-');

  std::vector<std::string> result;

  for (auto &&subrange : split_view) {
    result.emplace_back(subrange.begin(), subrange.end());
  }

  return result;
}

void preprocess(ParseResult &input) {
  // sort
  auto &ranges = input.ranges;
  std::sort(ranges.begin(), ranges.end(),
            [](const Range &a, const Range &b) { return a.start < b.start; });
  // merge overlaps
  std::vector<Range> merged;
  merged.reserve(ranges.size());
  for (const auto &r : ranges) {
    if (merged.empty() || r.start > merged.back().end + 1) {
      merged.push_back(r);
    } else {
      if (r.end > merged.back().end) {
        merged.back().end = r.end;
      }
    }
  }

  ranges.swap(merged);
}

ParseResult parseFile() {
  ParseResult result;

  std::fstream infile("./day05/input.txt");
  std::string line;

  // optimisitc reservation
  result.ranges.reserve(10 * 10);
  result.ingredients.reserve(10 * 10);
  bool afterBlank = false;

  while (std::getline(infile, line)) {
    if (line == "") {
      continue;
    } else if (line.find('-') != std::string::npos) {
      auto splitted = split(line);
      Range range;
      range.start = std::stoll(splitted[0]);
      range.end = std::stoll(splitted[1]);
      result.ranges.push_back(range);
    } else {
      result.ingredients.push_back(std::stoll(line));
    }
  }

  return result;
}

long long part1(const ParseResult &input) {
  const auto &ranges = input.ranges;
  const auto &ingredients = input.ingredients;

  long long sum = 0;

  for (auto val: ingredients) {
      auto it = std::upper_bound(ranges.begin(), ranges.end(), val, [](long long value, const Range& r) {
          return value < r.start;
      });

      if (it == ranges.begin()) {
          continue;
      }
      --it;
      if (val >= it->start && val <= it->end) {
          ++sum;
      }
  }

  return sum;
}

int main() {
  auto input = parseFile();
  auto start = std::chrono::high_resolution_clock::now();
  auto pt1 = part1(input);
  auto end = std::chrono::high_resolution_clock::now();
  auto duration =
      std::chrono::duration_cast<std::chrono::microseconds>(end - start);
  std::cout << "Part 1: " << pt1 << " and took " << duration.count() << "µs"
            << std::endl;
  return 0;
}
