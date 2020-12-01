Usage instructions:
	1. type "make" or "make all"
	2. Run using "./tigerc <file_name>
	
	OR
	1. type "make"
	2. python test_runner.py #to run the entire test suite and give categorized output

Using make should recompile the compiler and the associated grammar.

I precompiled goyacc and included it within tools folder and the makefile uses it as well.
I included goyaccs source code in the vendor folder as well.

External Dependencies: (Mostly for links, you shouldnt have to install these!)

Lexical analysis framework for Go - github.com/timtadh/lexmachine

Ultimately, I learned a lot about go, a lot about myself,
and a ton about lexing, semantics, type checking, and interpretation.

The Lexer doesn't catch unterminated comments, I tried to re work it for awhile to catch them but

All in all 10/10 learning experience, I'm much better at Go and made a compiler, pogchamp.
