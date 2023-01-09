# Go: Custom String Sorting
Sort a set of strings based on the following factors:

An odd length string should precede an even length string.
If both strings have odd lengths, the shorter of the two should precede. 
If both strings have even lengths, the longer of the two should precede.
If the two strings have equal lengths, they should be in alphabetical order.
 
Example
strArr =['abc', 'ab', 'abcde', 'a', 'abcd']

The result is ['a', 'abc', 'abcde', 'abcd', 'ab'] .

Odd length words are sorted in ascending order of length followed by even length words in descending order of length. There are no matching length words, so no alphabetical sorting is required.

Function Description
Complete the function customSorting in the editor below.

customSorting has the following parameter:

    string strArr[n]:  an array of strings


Returns
    string[n]: the sorted array

Constraints

1 ≤ length of strArr ≤ 1000
1 ≤ length of strArr[i] ≤ 100
Each strArr[i] contains uppercase and lowercase English letters only.
 

Input Format For Custom Testing
Sample Case 0
Sample Input For Custom Testing

STDIN    Function
-----    --------
5     →  strArr[] size n = 5
abc   →  strArr = ['abc', 'ab', 'abcde', 'a', abcd']
ab
abcde
a
abcd
Sample Output

a
abc
abcde
abcd
ab
Explanation

The odd length strings come first, sorted ascending by length. Even length words come last in descending length order.

Sample Case 1
Language: Go

Autocomplete Ready




More
113141516171819202122232425