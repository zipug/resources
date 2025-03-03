package utils

// Function to convert bytes to gigabytes
func BytesToGB(bytes uint64) float64 {
	const gb = 1 << 30 // 1 GB = 2^30 bytes
	return float64(bytes) / float64(gb)
}

// Function to convert bytes to megabytes
func BytesToMB(bytes uint64) float64 {
	const gb = 1 << 20 // 1 MB = 2^20 bytes
	return float64(bytes) / float64(gb)
}
