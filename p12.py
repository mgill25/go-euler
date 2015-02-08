import math, time

# we use this function to make a lookup table of primes
def primes(n):
    """
    compute the primes below n with the sieve of Eratosthenes
    Time complexity: O(n loglog n). The algorithm requires
    O(n) of memory.

    """

    ps = []
    sieve = [True] * (n)
    ps.append(2)
    for i in range(3, n, 2):
        if sieve[i]:
            ps.append(i)
            for j in range(i*i, n, i):
                sieve[j] = False
    return ps


# create lookup list and set of all prime numbers below a million
start = time.clock()
primes_lookup = primes(1000000)
primes_set    = set(primes_lookup)
elapsed_primes = (time.clock() - start)


def number_of_divisors(n):
    """
    returns the number of divisors of n
    """

    global primes_lookup, primes_set
    # if prime immediately return 2
    if n in primes_set:
        return 2

    # else we use formula (1)
    def aux(n , primes, divisors):
        """
        help function computes divisors
        """

        k = 0
        while n != 1:
            count = 1
            while n % primes[k] == 0:
                n /= primes[k]
                count += 1

            divisors *= count
            k += 1

        return divisors

    return aux(n , primes_lookup, 1)


def main(n):
    """
    returns the first triangular number that has more
    than n divisors
    """

    # we use the formula (2)
    def triangular(n): return n * (n + 1) / 2

    num = 1
    triangular_number = triangular(num)
    num_divisors = number_of_divisors(triangular_number)
    while num_divisors <= n:
        num += 1
        triangular_number = triangular(num)
        num_divisors = number_of_divisors(triangular_number)

    return int(triangular_number), num_divisors



#######################################################################
start = time.clock()
triangular_number, num_divisors = main(500)
elapsed_main = (time.clock() - start)
total = elapsed_primes + elapsed_main

print "triangular number:              {: <20}".format(triangular_number)
print "number of divisors:             {: <20}".format(num_divisors)
print "elapsed time for prime list:    {: <20}".format(elapsed_primes)
print "elapsed time for main function: {: <20}".format(elapsed_main)
print "total time spent:               {: <20}".format(total)
