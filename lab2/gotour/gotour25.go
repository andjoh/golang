package main
import "fmt"
import "math"

func main(){

var z,z1,x float64=1,0,2
for  math.Abs(z-z1)>0.001 && z!=z1{
z1=z
z=newton(z,x)
}

fmt.Printf("With Newtons method the square  root of %f is %f \n",x,z)
fmt.Printf("With math.Sqrt method the square root of %f is %f \n",x,math.Sqrt(x))


}
func newton(z, x float64) float64{
return z-(z*z-x)/2*z
}
