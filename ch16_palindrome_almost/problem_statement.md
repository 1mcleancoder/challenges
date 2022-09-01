Given a string, determine if it is almost a palindrome. 
A string is almost a palindrome if it becomes a palindrome by removing just one letter. Considering only alpha-numeric characters and ignore case sensitivity.

"race a car"
=>
"raceacar"

here **ea** are the central chars which is not palindrome but removing a single char either e or a, it becomes a palindrome.

Also, for "abccdba"
Here **ccd** do not match but if we replace d then **cc** is a palindrome but on the other hand if we remove **c** then the remaining **cd** is not a palindrome. But since there exists one palindrome with only one letter removed. So it is almost a palindrome.