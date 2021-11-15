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

- 
