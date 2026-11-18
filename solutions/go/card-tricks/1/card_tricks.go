package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index>len(slice){
        return -1
    }
    for i,v :=range slice{
        if i==index{
            return v
        }
    }
    return -1
}

// SetItem writes an item to a slice at a given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index>=len(slice)||index<0{
       slice=append(slice,value)
        return slice
    } else{
        slice[index]=value
    
    return slice
    }
    
	
        
 
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
   var newslice = []int{}

        newslice=append(newslice,values...)
        
    
    newslice=append(newslice,slice...)
    return newslice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index<0||index>=len(slice){
        return slice
    }	
	s1:=slice[:index]
	s2:=slice[index+1:]
    slice=append(s1,s2...)
	return slice
}
