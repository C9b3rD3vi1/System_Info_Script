
package main

import (
    "fmt"
	"time"
	"os/user"
	"runtime"
	"syscall"
	"os"

)

func main() {
	println("Starting")
	println("Fetching System Preferences and System information...")

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