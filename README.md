# System_Info_Script
Go script that fetches and displays system information,


This project is a Go script that fetches and displays comprehensive system information, such as operating system details, memory usage, CPU load, disk usage, network interfaces, and more. The script is useful for system diagnostics, monitoring, and retrieving essential system information.

# Features
Operating System and Architecture: Displays the operating system name and architecture.
CPU Information: Shows the number of CPU cores and currently active goroutines.
Memory Usage: Displays total, free, and used memory.
Disk Usage: Provides disk space usage for a specified path.
Network Interfaces: Lists network interfaces along with their IP addresses.
CPU Load: Shows the system’s CPU load averages for the last 1, 5, and 15 minutes.
Hostname and IP Address: Retrieves and displays the system hostname.
Go Runtime Information: Displays the Go version and compiler used.
System Uptime: Shows the system’s uptime since the last boot.
User Information: Retrieves and displays the current user’s username and home directory.
Detailed OS Information: Uses the uname command to fetch detailed OS information.


# Installation


## Clone the repository:


    git clone https://github.com/C9b3rD3vi1/System_Info_Script.git

    cd System_Info_Script


## Install dependencies:

Make sure you have Go installed. Additionally, the script uses the gopsutil package for CPU load statistics. You can install it with:

    go get github.com/shirou/gopsutil/load


## Build the project:

You can build the script into a binary:

    go build -o System_Infor_Script


## Run the script:

After building, you can run the script using:

You can run the script directly using Go:

    go run main.go


# Usage

The script will automatically fetch and display the following system information upon execution:

Operating System and architecture details
CPU information (cores and goroutines)
Memory usage (total, free, used)
Disk usage for the root directory (/)
Network interfaces and their IP addresses
CPU load averages (1m, 5m, 15m)
System hostname and uptime
Current user's username and home directory
Detailed OS information using the uname command
Simply run the script, and the output will be printed to your terminal.


# Customization
To check disk usage for a different directory, modify the diskUsage function call in the displaySystemInfo function:

go

    diskUsage("/your/path/here")

To include or exclude specific information, you can modify the displaySystemInfo function to only call the necessary functions.


# Contribution
Contributions are welcome! If you have suggestions, ideas, or want to report a bug, please open an issue or submit a pull request.

# License
This project is licensed under the MIT License. See the LICENSE file for details.

# Acknowledgments
This script uses the gopsutil library to retrieve CPU load statistics.
