[Quiz Game from Gophercises](https://courses.calhoun.io/lessons/les_goph_01)

# Part 1 - Read CSV and offer up CLI for answering
Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect.

The CSV file should default to problems.csv (example shown below), but the user should be able to customize the filename via a flag.

The first column is a question and the second column in the same row is the answer to that question.

### Packages:
- flags
- csv
- os
- time

# Part 2 - Add a Timer

The default time limit should be 30 seconds, but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has exceeded.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

# Bonuses
- Add an option (a new flag) to shuffle the quiz order each time it is run.
- Clean up the text (remove whitespace etc.) use the `strings` package
