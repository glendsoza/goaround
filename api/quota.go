package api

import "sync"

var currentQuota = APIQuota{}
var lock = &sync.RWMutex{}

// Sets the current quota
func SetCurrentQuota(quotaRemaining, quotaMax int) {
	// acquire lock
	lock.Lock()
	defer lock.Unlock()
	currentQuota.QuotaRemaining = quotaRemaining
	currentQuota.QuotaMax = quotaMax
}

// Get the current quota
func GetCurrentQuota() (int, int) {
	lock.RLock()
	defer lock.RUnlock()
	return currentQuota.QuotaRemaining, currentQuota.QuotaMax
}
