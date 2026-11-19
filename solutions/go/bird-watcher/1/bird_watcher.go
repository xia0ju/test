package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var s=0
for i:=0;i<len(birdsPerDay);i++{
    s+=birdsPerDay[i]
}
    return s
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var s=0
    var count=birdsPerDay[(week-1)*7:week*7]
	for _,v:=range count{                   
       s+=v
    }
    return s
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i:=0;i<len(birdsPerDay);i++{
        if i%2==0{
            birdsPerDay[i]++
        }
    }
    return birdsPerDay
}
