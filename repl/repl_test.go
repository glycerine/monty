package repl

/*
# manual tests:
# a challenge to somebody more clever than me: automate these.

---
>>> """ one """
# should return immediately with the string " one "
---
>>> """ one
long
expression
that
is
a
string"""
# should return immediately with the string.
---
>> """ one



"""
# should return immediately with the string " one\n\n\n\n"
---
def f(a,b):return a+b

# should define f immediately after that no-ident line.
f(
 1,
 2
)
# should *print* 3 immediately, without waiting for another line.
---
def f():
 a= """







"""
 return a

# should define f immediately after that last no-indent line.
b=f()
assert.eq(b, "\n\n\n\n\n\n\n\n")
---
b="""



"""
# should immediately (without waiting for another newline) define b equal to "\n\n\n\n".
*/
