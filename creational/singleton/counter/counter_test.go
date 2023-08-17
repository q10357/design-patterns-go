package counter

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// Lets do some tests!
// Only one counter, and it should be initialized
// If instance does exist, it should
func TestSingletonInitialization(t *testing.T) {
	counter := GetInstance()
	//counter should not be nil
	if counter == nil {
		t.Error("Counter is nil, not the expected value")
	}

	currentCount := counter.IncrementByOne()
	if currentCount != 1 {
		t.Errorf("Expected value 1 after first incremenet, got %d\n", currentCount)
	}

	expectedValue := counter
	//replica should have same value as counter
	counterReplica := GetInstance()
	if counterReplica != expectedValue {
		t.Error("Expected singleton class to only return one instance, but got two different")
	}

	currentCount = counterReplica.IncrementByOne()
	if currentCount != 2 {
		t.Errorf("The fields of original singleton and replicas are not identical")
	} else {
		assert.Equal(t, currentCount, 2)
	}

}
