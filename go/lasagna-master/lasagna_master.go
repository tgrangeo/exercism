package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, n int) int{
    if ( n == 0 ){
        n = 2
    }
	return len(layers) * n
}

// TODO: define the 'Quantities()' function
func Quantities(t []string) (int,float64){
    noodles := 0
    sauce := 0.0
    for i:= 0; i < len(t); i++{
        if (t[i] == "noodles"){
            noodles += 50
        }
    	if (t[i] == "sauce"){
            sauce += 0.2
        }
	}
	return noodles,sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend []string, me []string){
    maxFriend := len(friend)
    maxMe := len(me)
    me[maxMe - 1] = friend[maxFriend - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(q []float64, n int) []float64{
    size := len(q)
    ret := make([]float64,size)
    for i := 0 ; i < len(q); i++{
        ret[i] = (q[i] / 2.0) * float64(n)
    }
	return ret
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
