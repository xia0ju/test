package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(x []string,n int) int {
if n==0{
    n=2
}
    return n*len(x)


}
// TODO: define the 'Quantities()' function
func Quantities(x []string) (int,float64) {
    var y int
    var y2 float64
    for _,v:=range x{
        if v=="noodles"{
            y+=50
        }
        if v=="sauce"{
            y2+=0.2
        }
    }
    return y,y2
    
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(x []string,y []string)  {
	y[len(y)-1]=x[len(x)-1]
   
    
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(x [] float64,n int) []float64 {
    var s=make([]float64,len(x))
    for i,v:=range x{
       
        s[i]=v/2*float64(n)
    }
    return s
    
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
