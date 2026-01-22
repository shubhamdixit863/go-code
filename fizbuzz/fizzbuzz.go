package fizbuzz


func IfPrac(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "fizzbuzz"
	} else if num%3 == 0 {
		return "fizz"
	} else if num%5 == 0 {
		return "buzz"
	} else {
		return "cant be divided by both 3 and 5"
	}
}

func SwitchPrac(nums int)string {
	switch {
	case nums%3==0 && nums%5==0 :
		return "fizzBuzz"
	case nums%3==0:
		return "fizz"
	case nums%5==0:
		return "Buzz"
	default:
		return "The number is not divisible by 3 or 5 or both"
	}

}