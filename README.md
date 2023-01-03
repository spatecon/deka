# Deka
DSL (Domain Specific Language) with bytecode interpreter and compiler for it.

## Overview

1. **DSL syntax** description. Allows to generate a parser.
2. **Scanner and parser** are used to convert source code to AST. Could be compiled for Wasm to be used in front-end syntax highlighting.
3. **Byte-code emitter** takes AST as input and generate complete program in byte-code.
4. Interpreter (**Virtual Machine**) executes a byte-code. State of VM can be saved as image and loaded later.


