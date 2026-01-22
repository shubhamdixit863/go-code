package ifElse
func VoteEligibility(age int)string{
if age>18{
	return "You are eligible to vote"
}else if age<18{
	return "come back when you turn 18"
}else{
	return "Invalid age"
}
}