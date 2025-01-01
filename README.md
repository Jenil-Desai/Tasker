# 🎯 Tasker - Command-Center Inspired CLI Task Manager  

**Tasker** is a command-line interface (CLI) task manager inspired by a command-center experience. Built with **Go**, it lets you deploy, track, and accomplish tasks efficiently right from your terminal.  

---

## 📑 Table of Contents  

1. [Overview](#-overview)  
2. [Technologies](#-technologies)  
3. [Packages & Libraries Used](#-packages--libraries-used)  
4. [Getting Started](#-getting-started)  
5. [Setup](#-setup)  
6. [Features](#-features)  
7. [Demo & Screenshots](#-demo--screenshots)  
8. [Acknowledgments](#-acknowledgments)  
9. [License](#-license)  

---

## 🌟 Overview  

**Description**:  
Tasker is a powerful CLI task manager designed for efficiency and simplicity, making task management seamless for developers and power users.  

**Command Arsenal**:

| Command | Operation | Description |
|---------|-----------|-------------|
| `tasker add` | Deploy Mission | Deploy new objectives to your command center with priority levels and deadlines |
| `tasker delete` | Abort Mission | Eliminate specific missions from your radar using mission ID |
| `tasker list` | Scan Radar | Survey all active missions in your command center with detailed intel |
| `tasker complete` | Mission Accomplished | Mark objectives as completed and log them in mission history |

Use `tasker --help` to access the tactical manual for detailed command operations.

---

## 💻 Technologies  

This project is built with:  

| Language  | Database |  
|-----------|----------|  
| **Golang** | **SQLite** |  

---

## 📦 Packages / Libraries Used  

Tasker leverages the following Go packages:  

| Package / Library                      | Purpose                                        |  
|----------------------------------------|------------------------------------------------|  
| `Cobra`                                | Build a flexible CLI application framework    |  
| `text/tabwriter`                       | Formatting text into tabular output           |  
| `steebchen/prisma-client-go`           | Prisma ORM for SQLite                         |  
| `fatih/color`                          | Add color to terminal output                  |  

---

## 🚀 Getting Started  

- [Install IDE](https://code.visualstudio.com/)
- [Install Golang](https://go.dev/)

---

## ⚙️ Setup  

1. Clone The Repositry
    ```bash
    git clone https://GitHub.com/Jenil-Desai/Tasker
    ```
2. Go Into Project Folder And Install Dependeinces
    ```bash
    cd tasker && go mod vendor
    ```
3. Run The Tool
    ```bash
    go run main.go <command>
    ```

---

## 🎯 Features  

1. **Add Task**: Create new tasks with a simple command.  
2. **Delete Task**: Remove tasks no longer needed.  
3. **List Task**: View all tasks in an organized manner.  
4. **Mark Task**: Mark tasks as completed with ease.  

---

## 🔗 Demo & Screenshots  

- Coming Soon  

---

## 🙏 Acknowledgments  

Special thanks to the following resources:  

1. [Prisma for Go Documentation](https://goprisma.org/)  
2. [DreamsOfCode Go Projects](https://github.com/dreamsofcode-io/goprojects/tree/main)  
3. [Fatih Color Library](https://github.com/fatih/color)  
4. [JetBrains CLI Apps with Cobra](https://www.jetbrains.com/guide/go/tutorials/cli-apps-go-cobra/creating_cli/)  

---

## 📜 License  

This project is licensed under the [MIT License](LICENSE). See the [LICENSE](LICENSE) file for details.  

---

### 🚀 Task management just got simpler with **Tasker**!
