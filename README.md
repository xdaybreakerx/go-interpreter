# go-interpreter

### Challenges set in book:
#### Lexing
- [ ] lexer attaches the line number, column number and filename to a token
- [ ] support more than just ASCII (ie, unicode and utf-8, emojis, etc)
- [ ] support floats, hex, octal, and any numbers that are not strict integers 

#### Parsing 
- [ ] postfix operators (eg, foobar **++** )

#### Evaluation
- [ ] Boolean operands - adding, subtracting, dividing and multiplying booleans, checking whether true is greater than false with < or >. 
- [ ] Add stack trace to error object

#### Data Structures
- [ ] readString reports an error instead of simply returning when it reaches the end of the input.
- [ ] Add support for character escaping so that string literals like "hello \"world\"", "hello\n world" and "hello\t\t\tworld" work"

### Long term project challenges: 
- [ ] complete mvp of product as outlined in book
- [ ] complete the [lost chapter](https://interpreterbook.com/lost/)
- [ ] implement all additional features discussed in book 
- [ ] honestly i'm not a huge fan of *monkeys* so maybe i'll rewrite this interpreter to be for a language named after my cat *misha* 
- [ ] once above complete, check out the authors follow up book on [writing a compiler](https://compilerbook.com/). 

#### Reflections: 
Early days so far, but...

I’m working through *Writing An Interpreter In Go* as part of my broader journey in CS. I've almost exclusively worked with interpreted, dynamically typed languages (python, and js), so this book helps both:
- increase my understanding of what is happening with the languages I use
- gives me exposure to working with a compiled, statically typed language 

Building an interpreter will not only sharpen my problem-solving and system design skills but also provide me with insights into language constructs and execution models - that's the hope at least 😅 

At the very least, tacking this project is a way to peel back the layers and explore what happens beneath the surface when code is executed. The interpreter lets me see the immediate effects of my decisions, making it easier to grasp core concepts like parsing, tokenization, and evaluation.