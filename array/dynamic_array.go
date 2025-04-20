package array

import "errors"

const (
	DefaultCapacity = 10
)

type DynamicArray struct {
	Size 			int
	Capacity   		int
	ElementData     []any
}

// Put function is change/update the value in array with the index and new value
func (da *DynamicArray) Put(index int, data any) error {
	// 1. Check index out of range?
	if err := da.CheckRangeFromIndex(index); err != nil {
		return err
	}
	// 2. Change data
	da.ElementData[index] = data
	return nil
}

// Add function is add new element to our array
func (da *DynamicArray) Add(data any) {
	// 1. Check if array has enough space or not?
	if da.Size == da.Capacity {
		da.NewCapacity()
	}

	// 2. Add to to the end & increase size
	da.ElementData[da.Size] = data
	da.Size += 1 
}

// Remove function is remove an element with the index
func (da *DynamicArray) Remove(index int) error {
	// 1. Check if array is empty
	if da.IsEmpty() {
		return errors.New("array is empty.")
	}

	// 2. check if index out of range or not?
	if err := da.CheckRangeFromIndex(index); err != nil {
		return err
	}

	// 3. Remove element by copy from index+1 to index & reduce size
	copy(da.ElementData[index:], da.ElementData[index+1:])
	da.Size--
	da.ElementData[da.Size] = nil 	// if we set variable to nil, it will be cleared by GC
	return nil
}

// Get function is return one element with the index of array
func (da *DynamicArray) Get(index int) (any, error) {
	// 1. Check if array is empty
	if da.IsEmpty() {
		return nil, errors.New("array is empty.")
	}

	// 2. check if index out of range or not?
	if err := da.CheckRangeFromIndex(index); err != nil {
		return nil, err
	}

	// 3. Get data and return
	return da.ElementData[index], nil
}

// IsEmpty function is check that the array has value or not
func (da *DynamicArray) IsEmpty() bool {
	return da.Size == 0
}

// GetAll function return all value of array
func (da *DynamicArray) GetAll() []any {
	return da.ElementData[:da.Size]
}

// CheckRangeFromIndex function will check the range from the index
func (da *DynamicArray) CheckRangeFromIndex(index int) error {
	if index < 0 || index >= da.Size {
		return errors.New("index out of range")
	}
	return nil
}

// NewCapacity function increase the Capacity
func (da *DynamicArray) NewCapacity() {
	// Allocating more space
	if da.Capacity == 0 {
		da.Capacity = DefaultCapacity
	} else {
		da.Capacity = da.Capacity << 1 // multiple by 2
	}
	newDataElement := make([]any, da.Capacity)

	// copy old data to new data
	copy(newDataElement, da.ElementData)

	// assign Element to new data (note that old empty will be clear by GC)
	da.ElementData = newDataElement	
} 
