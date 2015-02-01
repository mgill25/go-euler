# Euler p9: Pythagorean Triples
# Naive iterative solution with a few optimizations.
# a < b < c
import time

start = time.time()

for a in range(1, 1000):
    # b will be at least as big as a, and at most as big as 1000 - a.
    for b in range(a, 1000 - a):
        # since sum is 1000; c will be...
        c = 1000 - b - a
        if a * a + b * b == c * c:
            print(a, b, c)
            print("Product: {}".format(a * b * c))

elapsed = time.time() - start
print("Time: {:.5f} seconds".format(elapsed))
