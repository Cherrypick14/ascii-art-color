
# Ascii-art-color

This command-line tool allows users to colorize text strings, either entirely or by specified substrings, using various color code systems. Ideal for developers and command-line enthusiasts, it improves text visibility and readability in terminal applications.



## Installation

Prerequisites:

Go 1.18 or higher installed Git installed Operating System: Windows, macOS, Linux

Download/Clone:

```bash
  $ git clone https://github.com/Cherrypick14/ascii-art-color.git
```

## Features

- It colors a specified substring or the entire string if no substring is provided.

- It accepts colors in different notations (RGB, HSL, ANSI, etc.).

- It displays a usage message if the flag format is incorrect.

- The program accepts additional ASCII-art optional projects.

- It does handle errors  to manage incorrect or improperly formatted inputs.

- Validates the program's handling of multiple inputs by testing various combinations and edge cases.


## Usage
To run the program, navigate to the directory where the program is installed and use the following command:

### Basic Command
```bash
go run . --color=<color> <substring to be colored> "your string here"
```
### Examples

- Color a Substring
```bash
go run . --color=red kit "a king kitten have kit"
```
### Output
![red substring](<ascii/screenshots/redsubstring.png>)

In the example above, the substring "kit" in "kitten" and the word "kit" at the end will be colored red.

- Color the entire string:
```bash
go run . --color=blue "a king kitten have kit"
```
### Output
![blue string](<ascii/screenshots/bluestring.png>)

In this example, the entire string "a king kitten have kit" will be colored blue.

## Usage Message
If the flag format is incorrect, the following usage message will be displayed:

```bash
Usage: go run . [OPTION] [STRING]

Example: go run . --color=<color> <substring to be colored> "something"
```

## Running Tests

To run tests, navigate to the ascii directory and run the following command

```bash
go test -v
```

## License

[MIT](https://choosealicense.com/licenses/mit/)


## Authors

- [@Cherrypick14](https://github.com/Cherrypick14)
- [@stkisengese](https://github.com/stkisengese)

## References
For more different custom color combinations that you can't find in the project, you can check out using this format:

```bash
go run . --color=#ffffff hello

go run . "--color=rgb(255, 0, 0)" hi
```

- **Color Picker Tools**:
  - [HTML Color Picker - W3Schools](https://www.w3schools.com/colors/colors_picker.asp)
- **Color Palette Generators**:
  - [Paletton](https://paletton.com/)
  - [Color Hunt](https://colorhunt.co/)

