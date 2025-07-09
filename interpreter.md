# Basics of interpreter

- What we're going to build: Code -> intermediate representation (AST) -> evaluation

Major parts of the algorithm:

- Lexer - tokenizer - identifying words (lex - greek for word)
- Parser - Takes the words/tokens to build a tree/AST which can be interpreted to create sentences
- AST - ^
- Static analysis - Reference resolution for variables (scoping), for statically types languages (type checking), giving semantics to AST
  (Everything uptil now is Frontend)
- Intermediate representations - not tightly coupled with source or destination forms
  - different forms of IR presentation: CFG, SSA, Continuation passing style, three address code
  - Bridge b/w FE and BE - can be ported to different architecture of machines
- Optimization - mostly compile time optimizations transforming the code
- Code generation
- Virtual machine
- Runtime

## Shortcuts & Alternate Routes

- Single pass compilers - combines parsingm, analysis and code generation
- Tree-walk interpreters
- Transpilers
- JIT compilation
