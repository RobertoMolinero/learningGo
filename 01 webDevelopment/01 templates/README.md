# Templates

The use of html templates is examined in this part of my repository.

## Data

First, data is transferred to the template.

### Writing Template

In the first example, the template is still a variable within the program. The value of a second variable is inserted and the result is output as html file.

### Parsing Template

Here the individual possibilities are indicated to load an existing template.

* Single file
* Multiple files
* Entire directory with all files of a certain pattern ("*.gohtml")

### Variable

Different types of variables are each transferred to a template.

### Data Structures

Here are the examples with the more complex data types. First lists and maps and then structures which in turn contain lists and maps at the end.

The last example is very close to a complex real-world example in which the data comes from a relational database.

## Functions

The uses of predefined and self-written functions are explained here with examples.

### String & Date

The first two sections use functions to manipulate dates and strings.

### Pipeline

In the next section, functions will be interconnected with pipes.

### Predefined

The topic here are the predefined functions that can be used directly in the template.

* The functions 'gt', 'lt', 'eq' to compare transferred values
* The function 'index' to enumerate transferred values
* The function 'if' to for evaluating transferred values

### Nested Templates

Here the combination of several templates is demonstrated. A typical use case would be a header/footer that should be the same on all pages.

### Methods

In the last section, functions of one type are written and used in the template.