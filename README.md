# Renom

[![Go CI](https://github.com/CHashtager/Renom/actions/workflows/go-tests.yml/badge.svg)](https://github.com/CHashtager/Renom/actions)


A simple CLI tool to recursively rename files, directories, and replace text inside files.

## 🚀 Features
- Rename directories that contain a specific substring.
- Rename files that contain a specific substring.
- Replace text inside files.

## 🛠️ Installation

### **1️⃣ Clone the Repository**
```bash
git clone https://github.com/CHashtager/Renom.git
cd renom
```

### **2️⃣ Build the CLI**
Compile the Go binary:
```bash
go build -o renom ./cmd/main.go
```
This will generate an executable file named `renom`.

---

## 📌 Adding to PATH for Global Use

To use `renom` from anywhere in the terminal, you need to add it to your system's **PATH**.

### **🔹 Linux / macOS**
1. Move the binary to `/usr/local/bin`:
   ```bash
   sudo mv renom /usr/local/bin/
   ```
2. Verify installation:
   ```bash
   renom --help
   ```

### **🔹 Windows**
1. Move `renom.exe` to a permanent location, e.g., `C:\renom\`
2. Add this directory to the **system PATH**:
   - Open **Start Menu** → Search **"Environment Variables"** → Open **"Edit the system environment variables"**
   - Under **System Properties**, click **Environment Variables**
   - Under **System variables**, select **Path** → **Edit**
   - Click **New**, then add `C:\renom\`
   - Click **OK** to save

3. Verify installation:
   ```powershell
   renom --help
   ```

---

## 🏃 Usage
```bash
renom old_string new_string
```
Example:
```bash
renom "test" "example"
```
This renames all occurrences of `"test"` with `"example"` in files, directories, and file contents.

---

## 🛠 Uninstalling
To remove the tool, delete the binary from the PATH:
```bash
sudo rm /usr/local/bin/renom  # Linux/macOS
del C:\renom\renom.exe  # Windows
```

---

## 📝 License
MIT License
