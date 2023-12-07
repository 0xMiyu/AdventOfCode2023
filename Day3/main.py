import itertools
import math
import re
from collections import defaultdict

with open("input.txt") as f:
    ls = f.read().strip().split("\n")

box = list(itertools.product((-1, 0, 1), (-1, 0, 1)))
symbols = {
    (i, j): x
    for i, l in enumerate(ls)
    for j, x in enumerate(l)
    if x != "." and not x.isdigit()
}

part_sum = 0
parts_by_symbol = defaultdict(list)

for i, l in enumerate(ls):
    for match in re.finditer(r"\d+", l):
        n = int(match.group(0))
        boundary = {
            (i + di, j + dj)
            for di, dj in box
            for j in range(match.start(), match.end())
        }
        if symbols.keys() & boundary:
            part_sum += n
        for symbol in symbols.keys() & boundary:
            parts_by_symbol[symbol].append(n)

# Part 1
print(part_sum)

# Part 2
print(
    sum(
        math.prod(v)
        for k, v in parts_by_symbol.items()
        if len(v) == 2 and symbols[k] == "*"
    )
)