package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/load"
)


func main() {
	println("Starting")
	println("Fetching System Preferences and System information...")

	println("--------------------------------")

	displaySystemInfo()

	println("--------------------------------")

	// Operating System and Architecture
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)

	// CPU and Goroutine Information
	fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
	fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())

	// Go Version
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("Go Compiler: %s\n", runtime.Compiler)

	// User Information
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error fetching user information: %v\n", err)
		return
	}
	fmt.Printf("Username: %s\n", currentUser.Username)
	fmt.Printf("Home Directory: %s\n", currentUser.HomeDir)

	// Hostname and IP Address
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error fetching hostname: %v\n", err)
		return
	}
	fmt.Printf("Hostname: %s\n", hostname)

	//uptime estimate
	fmt.Printf("Uptime: %s\n", uptime())

}


// Uptime calculates the system uptime
func uptime() string {
	// Get the system boot time using the syscall package
	var sysinfo syscall.Sysinfo_t
	err := syscall.Sysinfo(&sysinfo)
	if err != nil {
		return fmt.Sprintf("Error getting uptime: %v", err)
	}

	// Calculate the uptime duration
	uptimeDuration := time.Duration(sysinfo.Uptime) * time.Second
	return fmt.Sprintf("%v", uptimeDuration)
}


// A function that obtains the memory usage and information of the system
func memoryInfo() {
	var sysinfo syscall.Sysinfo_t
	err := syscall.Sysinfo(&sysinfo)
	if err != nil {
		fmt.Printf("Error getting memory info: %v\n", err)
		return
	}

	totalRAM := sysinfo.Totalram * uint64(syscall.Getpagesize()) / (1024 * 1024)
	freeRAM := sysinfo.Freeram * uint64(syscall.Getpagesize()) / (1024 * 1024)
	usedRAM := totalRAM - freeRAM

	fmt.Printf("Total RAM: %d MB\n", totalRAM)
	fmt.Printf("Free RAM: %d MB\n", freeRAM)
	fmt.Printf("Used RAM: %d MB\n", usedRAM)
}

// A function that displays the disk space usage and information of the system

func diskUsage(path string) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		fmt.Printf("Error getting disk usage: %v\n", err)
		return
	}

	total := stat.Blocks * uint64(stat.Bsize) / (1024 * 1024)
	free := stat.Bfree * uint64(stat.Bsize) / (1024 * 1024)
	used := total - free

	fmt.Printf("Disk Usage for %s:\n", path)
	fmt.Printf("Total: %d MB\n", total)
	fmt.Printf("Used: %d MB\n", used)
	fmt.Printf("Free: %d MB\n", free)
}


// a function that returns network interface address and port number

func networkInfo() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error getting network interfaces: %v\n", err)
		return
	}

	for _, iface := range interfaces {
		fmt.Printf("Interface: %s\n", iface.Name)
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Printf("Error getting addresses for %s: %v\n", iface.Name, err)
			continue
		}
		for _, addr := range addrs {
			fmt.Printf("  Address: %s\n", addr.String())
		}
	}
}


// CPU load statistics function
func cpuLoad() {
	avg, err := load.Avg()
	if err != nil {
		fmt.Printf("Error getting CPU load: %v\n", err)
		return
	}

	fmt.Printf("Load Average (1m, 5m, 15m): %.2f, %.2f, %.2f\n", avg.Load1, avg.Load5, avg.Load15)
}


// function to return detailed OS information
func osInfo() {
	out, err := exec.Command("uname", "-a").Output()
	if err != nil {
		fmt.Printf("Error getting OS info: %v\n", err)
		return
	}
	fmt.Printf("OS Info: %s\n", string(out))
}



// Final function to print all this information from all those functions
func displaySystemInfo() {
	fmt.Println("Fetching System Information...")
	fmt.Println()
	osInfo()
	fmt.Println()
	memoryInfo()
	fmt.Println()
	diskUsage("/")
	fmt.Println()
	networkInfo()
	fmt.Println()
	cpuLoad()
	fmt.Println()
	fmt.Printf("Uptime: %s\n", uptime())
}
