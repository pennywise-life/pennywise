#!/bin/bash

# Function to check if the OS is Windows
is_windows() {
    [[ "$OSTYPE" == "msys" || "$OSTYPE" == "cygwin" ]]
}

# Function to check if the OS is macOS
is_macos() {
    [[ "$OSTYPE" == "darwin"* ]]
}

# Function to check if the OS is Linux
is_linux() {
    [[ "$OSTYPE" == "linux-gnu"* ]]
}

echo "Building the pennywise executable..."

# Add platform-specific build commands
if is_windows; then
    echo "Windows platform detected. Build commands for Windows not implemented."
    exit 1
elif is_macos; then
    # macOS build command (assuming you have Go installed)
    go build -o pw . || { echo "Build failed!"; exit 1; }
elif is_linux; then
    # Linux build command
    go build -o pw . || { echo "Build failed!"; exit 1; }
else
    echo "Unsupported operating system."
    exit 1
fi

echo "Waiting for build completion..."
sleep 5

# Add platform-specific move commands
if is_windows; then
    echo "Windows platform detected. Move commands for Windows not implemented."
    exit 1
elif is_macos; then
    # macOS move command
    sudo mv pw /usr/local/bin || { echo "Move failed!"; exit 1; }
elif is_linux; then
    # Linux move command
    sudo mv pw /usr/local/bin || { echo "Move failed!"; exit 1; }
fi

# Determine the current shell accurately
shell_path=$(readlink /proc/$$/exe)
shell_name=${shell_path##*/}

echo "Current shell: $shell_name"

# Add platform-specific completion setup commands
if [[ $shell_name = "bash" ]]; then
    if is_windows; then
        echo "Bash completion setup not supported on Windows."
    elif is_macos || is_linux; then
        # Assume 'compage completion' is a command for generating completions
        compage completion bash > ~/.bashrc_pw
        source ~/.bashrc_pw
        echo "Bash completion setup done."
    fi
else
    echo "Not using Bash, skipping completion setup."
fi

echo "pennywise has been successfully installed!"
