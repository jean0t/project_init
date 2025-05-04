# **project_init**  

A simple CLI tool to bootstrap new projects with opinionated defaults.  

---

## Features

- Creates base directory structure  
- Adds a `LICENSE` file (MIT or GPL-3)  
- Initializes a Git repository (`git init`)  
- Adds a sensible `.gitignore`  
- Supports **Go** and **Python** out of the box  
- CLI-driven—no GUI
- Adds a `README.md` file  

---

## Installation

1. Clone this repo:  
   ```bash
   git clone https://github.com/jean0t/project_init.git
   cd project_init
   ```
2. Build the binary:  
   ```bash
   go build -o project_init ./project_init
   ```
3. (Optionally) move it into your `$PATH`:  
   ```bash
   mv project_init /usr/local/bin/
   ```

**OBS: you can avoid clone the repo and build the binary by yourself, simply download in the releases the binary and go to step 3 directly**

---

## Usage

```bash
project_init -lang <go|python> [-license <mit|gpl-3>] [-gitignore=<true|false>] <project_name>
```

| Flag          | Description                                              | Default   |
|---------------|----------------------------------------------------------|-----------|
| `-lang`      | Language scaffold to generate (`go` or `python`)         | `python`  |
| `-license`   | License type (`mit` or `gpl-3`)                          | `gpl-3`   |
| `-gitignore` | Include a `.gitignore` file (`true` or `false`)          | `true`    |

### Examples

- Create a Go project with MIT license:  
  ```bash
  project_init myapp --lang go --license mit
  ```
- Create a Python project with GPL-3 and no `.gitignore`:  
  ```bash
  project_init service --lang python --license gpl-3 --gitignore=false
  ```

---

## Generated Structure

```text
myapp/
├── LICENSE        # MIT or GPL-3, per flag
├── .gitignore
├── src/           # source code
│   └── main.go    # or main.py
├── tests/         # test stubs
└── .git/          # Git repo
```

---

## README Skeleton

Each new project also gets a starter `README.md` with:

```markdown
# <project_name>

A brief description of what the project does.

## Installation

Instructions to install or build.

## Usage

Examples of how to run the project.

## License

<MIT or GPL-3> © <year> <your name>
```

---

## To-Do

- [ ] Add support for more languages (e.g., JavaScript, Rust, Java)  
- [x] Enhance templates (customizable README, code samples)  
- [x] Allow custom license text or additional license types  

---

## Contributing

Forks and pull requests are welcome!  

My motivation: I loved Crystal’s built-in `crystal init` skeleton generator and wanted something similar for the languages I use every day.  

Feel free to add new languages, improve templates, or suggest enhancements.
