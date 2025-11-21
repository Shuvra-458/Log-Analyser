# golog – A Simple Log Analysis CLI in Go

`golog` is a lightweight, beginner-friendly but production-style **log analysis CLI tool** written in Go. It helps you quickly inspect log files, extract useful information, and tail logs in real time.

This project is perfect for:

* Learning Go
* Learning how to build real CLI tools with Cobra
* DevOps / SRE tool-building experience

---

## 🚀 Features

### ✔ Log Statistics

Analyze a log file and get:

* Total lines
* Number of error entries
* Number of warnings
* Optional: Save stats into a plain text file

### ✔ Error Extraction

Extract all log lines containing errors.
Can optionally save output to a text file.

### ✔ Real-time Log Tailing

Like a mini `tail -f` built in Go.
Shows new log entries as they appear.

### ✔ Clean Project Structure

Uses Go modules, Cobra command structure, and reusable helper functions.

---

## 📦 Project Structure

```
golog/
├── cmd/
│   ├── root.go        # Base command
│   ├── stats.go       # Stats command
│   ├── errors.go      # Errors command
│   ├── tail.go        # Tail command
│   └── utils.go       # Helper functions
├── go.mod
└── main.go
```

---

## 🔧 Installation

### Clone repo:

```
git clone <your repo url>
cd golog
```

### Build binary:

```
go build -o golog
```

Run:

```
./golog --help
```

---

## 🧭 Usage

### 📊 1. View Log Statistics

```
./golog stats -f /var/log/syslog
```

**Save stats to a text file:**

```
./golog stats -f /var/log/syslog -o stats.txt
```

**Save into a directory (auto filename):**

```
./golog stats -f /var/log/syslog -o ./output
```

---

### ❌ 2. Extract Error Logs

```
./golog errors -f /var/log/syslog
```

**Save to file:**

```
./golog errors -f /var/log/syslog -o errors.txt
```

---

### 📡 3. Real-time Log Tailing

```
./golog tail -f /var/log/syslog
```

This works similar to:

```
tail -f /var/log/syslog
```

---

## 📥 Flags

### Global

```
-f, --file          Path to log file
-h, --help          Help menu
```

### stats command

```
-o, --output-file   Save statistics to a text file
```

### errors command

```
-o, --output-file   Save extracted errors to a text file
```

---

## 🛠 Technologies Used

* **Go (Golang)**
* **Cobra CLI framework**
* Standard libraries (`os`, `strings`, `time`, `bufio`)

---

## 🌱 Future Enhancements

This project is designed to be extendable. Potential next features:

* JSON/CSV export
* Log level detection (INFO/WARN/DEBUG/ERROR)
* Date-range filtering
* Follow multiple log files
* Write logs to SQLite or BoltDB for querying
* REST API wrapper

---

## 🧑‍💻 Author

Built with ❤️ by a DevOps-driven Go learner.
Feel free to extend, fork, and use this in your workflow.

---

## 📝 License

MIT License.
Feel free to modify and use.
