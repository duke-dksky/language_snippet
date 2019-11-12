composition
=============================

[methods,interfaces and embedded-types](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html)

When we embed a type,
the methods of that type become methods of the outer type,
but when they are invoked,
the receiver of the method is the inner type,
not the outer one. 
            - Effective Go


### method sets

1. The method set of the corresponding pointer type \*T is the set of all methods with receiver \*T or T
2. The method set of any other type T consists of all methods with receiver type T
3. If S contains an anonymous field T, the method sets of S and \*S both include promoted methods with receiver T
4. If S contains an anonymous field \*T, the method sets of S and \*S both include promoted methods with receiver T or \*T


