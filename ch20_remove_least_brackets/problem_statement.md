Given a string only containing round brackets '(' and ')' and lowercase characters, remove the least amount of brackets so that string is valid.

The string is valid if it is empty or if there are brackets, they are all close.

"a)bc(d)" => Here valid string is: abc(d) 
"(ab(c)a" => Here valid strings can be: (abc)a or ab(c)a
"))((" => Here valid strings can only be an empty string: ""

Constraints:
- What do we return from our algorithm?
  - A valid string with fewer brackets removed.
- Will there be spaces in the strings?
  - Assume no. The string only contians lowercase chars and '(' and ')'