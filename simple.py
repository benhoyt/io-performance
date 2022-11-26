# Run with the following command to clear the read cache between runs:
#
# sudo sysctl vm.drop_caches=3; python3 simple.py <kjvbible_x100.txt >/dev/null

import collections
import sys
import time

t0 = time.time()
content = sys.stdin.read()
t1 = time.time()
read_time = t1 - t0

counts = collections.Counter(content.lower().split())
t2 = time.time()
process_time = t2 - t1

most_common = counts.most_common()
t3 = time.time()
sort_time = t3 - t2

for word, count in most_common:
	print(word, count)

t4 = time.time()
output_time = t4 - t3

total_time = t4 - t0

print('Reading   :', read_time, file=sys.stderr)
print('Processing:', process_time, file=sys.stderr)
print('Sorting   :', sort_time, file=sys.stderr)
print('Outputting:', output_time, file=sys.stderr)
print('TOTAL     :', total_time, file=sys.stderr)
