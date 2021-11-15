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

- to check in an array of nums which is the only number that appears once, xor everything together, the sum of the odd times occuring numbers survive.
    - https://leetcode.com/problems/single-number/

- for creating all possible bit masks to create all different combinations possible of choosing from n options:
     for (int mask = 0; mask < (1 << n); mask++)
            System.out.println(Integer.toString(mask,2));

- Integer.toString(number,2) converts the num to binary before stringifying it.
- sweet one liner to check if a num is a power of x :  Integer.toString(n,x).matches("^10*");

