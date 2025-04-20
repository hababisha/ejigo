# ejigo 

**Skip classes with freedom.**  
`ejigo` is a simple command-line tool to track your university courses and absence counts — nothing fancy, made it to learn go.


---

##  Idea Behind the Tool

One day I was sitting at dorm, bored out of my mind, thinking about skipping the 2:30 (afternoon) class if you know what I mean. But I had no clue how many absences I already had. Was I close to the 20 percent limit? Would I still be allowed to sit for finals? I couldn't even remember which course had how many absences.

That's why ejigo — a lightweight CLI tool to help me keep track of absences per course. Not to avoid responsibility, but to skip classes efficiently lol.

---

## Commands available now

- `add`: Add a new course
- `absent`: Increment absence count for a course
- `reset`: Reset all courses (at the end of the semester)
- `summary`: View your courses and their absence counts

use ejigo --help for more details

---

## 📦 Installation (Manual)

> ⚠️ Requires [Go](https://go.dev/doc/install) to be installed.

### 1. Clone the repository

```bash
git clone https://github.com/hababisha/ejigo.git
cd ejigo
```

### 2. Build the executable

```bash
go build -o ejigo
```

This will generate a binary file called `ejigo` in the current directory.

---

## ✅ Optional: Add `ejigo` to your PATH

If you want to use `ejigo` from anywhere in your terminal:

### Linux/macOS

Move the binary to a directory in your PATH, like `~/.local/bin`:

```bash
mkdir -p ~/.local/bin
cp ejigo ~/.local/bin/ejigo
```

Then make sure `~/.local/bin` is in your shell config (`~/.bashrc`, `~/.zshrc`, or `~/.bash_profile`):

```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

Now you can use:

```bash
ejigo summary
```

from any directory 

---

## Contribute

PRs and suggestions are welcome! If you have ideas to make `ejigo` better, open an issue or fork and improve it.

