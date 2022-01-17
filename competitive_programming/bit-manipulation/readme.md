# Bit Manipulation
- import java.math.BigInteger; some places this might be necessary

- If the input is very big and therefore given in a string, might want to use BigInteger.
    - BigInteger class logical operators : 
        - a.and(b)
        - a.or(b)
        - a.shiftLeft(1)
        - a.toString(2) : converts the bigInteger to string with base 2

- For adding a number without using addition :
    - break the operation of calculating sum and carry for each bit, and then add the sum and carry :
        - sum = a ^ b
        - carry =  (a & b ) << 1; carry is left shifted once as carry from this bit will be used for next bit, not this
        - ans = sum + carry;

- XOR of any number with itself is 0. Think of it as a filter for even number of occurences.

- XOR of anything with 0 is the number itself. Think of this as the number being added to a variable with value 0.

- XOR of any nuber with all 1s is the compliment of that number. It flips all bits in the number.

- to find the all-bits-1 number greater than a number : int n = 1; while(n < k) n = (n << 1) | 1;

- to find power of 2greater than this number : int n = 1; while(n < k) n = n << 1;

- to check in an array of nums which is the only number that appears once, xor everything together, the sum of the odd times occuring numbers survive.
    - https://leetcode.com/problems/single-number/

- for creating all possible bit masks to create all different combinations possible of choosing from n options:
     for (int mask = 0; mask < (1 << n); mask++)
            System.out.println(Integer.toString(mask,2));

- Integer.toString(number,2) converts the num to binary before stringifying it.
- sweet one liner to check if a num is a power of x :  Integer.toString(n,x).matches("^10*");
- check if a num is a power of 2 : n & n-1 should be 0. e.g. 4: 100 & 011 = 0. 5: 101 & 100 != 0
- check if any number can be represented as powers of 3 (or anything uptill 9)
    - Bits of any number when represented in x-nary form can take values upto x-1. e.g.
        - 7 with base 3 is 21 = 3^1 (2) + 3^0 (1).
- -x = ~x + 1
- x & -x results in a sequence where only the first (from least sig. side) high bit of x remains high. all else is 0.

