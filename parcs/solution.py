# import gmpy2
from Pyro4 import expose
import random


def jacobi(a, n):
        assert(n > a > 0 and n%2 == 1)
        t = 1
        while a != 0:
            while a % 2 == 0:
                a /= 2
                r = n % 8
                if r == 3 or r == 5:
                    t = -t
            a, n = n, a
            if a % 4 == n % 4 == 3:
                t = -t
            a %= n
        if n == 1:
            return t
        else:
            return 0

class Solver:
    def __init__(self, workers=None, input_file_name=None, output_file_name=None):
        self.input_file_name = input_file_name
        self.output_file_name = output_file_name
        self.workers = workers

    def solve(self):
        (a, b, k) = self.read_input()
     
        step_n = (b - a) / len(self.workers)
    
        # map
        mapped = []
        for i in xrange(0, len(self.workers)):
            mapped.append(self.workers[i].mymap(str(a + i * step_n), str(a + (i + 1) * step_n), k))

        # reduce
        primes = self.myreduce(mapped)

        # output
        self.write_output(primes)

        #print("Job Finished")

    @staticmethod
    @expose
    def mymap(a, b, k):
        a = int(a)
        b = int(b)
        primes = []

        if a % 2 == 0:
            a += 1

        while a < b:
            if Solver.is_probable_prime(a,k):
                primes.append(str(a))
            a += 2

        return primes

    @staticmethod
    @expose
    def myreduce(mapped):
        output = []

        for primes in mapped:
            output = output + primes.value
        return output

    def read_input(self):
        f = open(self.input_file_name, 'r')
        a = int(f.readline())
        b = int(f.readline())
        k = int(f.readline())
        f.close()
        return a, b, k

    def write_output(self, output):
        f = open(self.output_file_name, 'w')
        f.write(', '.join(output))
        f.write('\n')
        f.close()
    
    @staticmethod
    @expose
    def gcd(a, b):
        while a != b:
            if a > b:
                a = a - b
            else:
                b = b - a        
        return a

    @staticmethod
    @expose
    def is_probable_prime(n, k) :
        if n == 1:
            return False
        # special case 2
        if n == 2:
            return True
        # ensure n is odd
        if n % 2 == 0:
            return False

        for i in range(1,k):
            a = random.randint(2,n-1)

            if Solver.gcd(a, n) > 1:
                return False

            if (((a ** ((n - 1) / 2))%n) != (jacobi(a,n)%n)) :
                return False
            
        return True
