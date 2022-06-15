# Code Kata for Golang #1

[A code kata is an exercise in programming which helps a programmer hone their skills through practice and repetition.](<https://en.wikipedia.org/wiki/Kata_(programming)>)c.

- [Topic](#topic)
- [Frame conditions](#frame-conditions)
- [Tasks](#tasks)
- [Procedure](#procedure)
- [License](#license)

## Topic

You have to implement an abstracted and simple application to manage a library of books and magazines.

## Frame conditions

1. You have 1.5 hours for this Kata - not a minute longer.

   If you reach this time limit stop your work immediately.
   It is an important part of the kata to respect the time limit.

2. There are no restrictions on how to use the provided time.
   If you want to code the entire time or take breaks, etc. you are free to.

3. This is a real world situation. You are allowed to consult the Internet, import any package you like, call a friend, et

   **BUT:** You are not allowed to do [pair programming](https://en.wikipedia.org/wiki/Pair_programming).
   **AND** If you have already done this kata before, don't review your previous implementation or steal code from it.

4. Develop your code based using Go v1.18.x

5. Keep the following priorities in mind while you implementing - in the mentioned order:

   1. Code quality
   2. Readability
   3. Feature Set
   4. Test Coverage

6. Given resources:

   > **Hint:** There is a reason why there are so many books and authors with [umlauts](https://en.wikipedia.org/wiki/Germanic_umlaut).

   - [`authors.csv`](resources/authors.csv): Contains authors with its `email`, `firstName` and `lastName`.
   - [`books.csv`](resources/books.csv): Contains books with its `title`, `description`, one or more `authors` and an `isbn`.
   - [`magazines.csv`](resources/magazines.csv): Contains magazines with its `title`, one or more `authors`, a `publishedAt` and an `isbn`.

## Tasks

- [Main tasks](#main-tasks)
- [Optional tasks](#optional-tasks)

### Main tasks

1. Your software should read all data from the given CSV files in a meaningful structure.

2. Print out all books and magazines (could be a GUI, console, …) with all their details (with a meaningful output format).

   > **Hint**: Do not call `printBooks(...)` first and then `printMagazines(...)` ;-)

3. Find a book or magazine by its `isbn`.

4. Find all books and magazines by their `authors`’ email.

5. Print out all books and magazines with all their details sorted by `title`.
   This sort should be done for books and magazines together.

### Optional tasks

> **Hint:** Optional means optional.

1. Write Unit tests for one or more methods.

2. Implement an interactive user interface for one or more of the Main tasks. (GUI or CLI is fine)

3. Add a book and a magazine to the data structure of your software and export it to a new CSV files.

## Procedure

1. Get the code: [Fork this repository](https://github.com/fathom5/code-kata-1/fork), and clone it.
2. Edit go.mod `module` directive to point to your fork.
3. Start the timer and start working on the kata.
4. Discuss your solution.
5. Repeat this kata again in a few days, weeks or next month.

## Run It

    make start

## Test It

    make test

## License

See [LICENSE](LICENSE) file.
