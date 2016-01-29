package main

import( "fmt" )

func main(){
	bestFinish := bestLeagueFinishes(3,10,7,4,16)
	fmt.Println(bestFinish)

}

func bestLeagueFinishes( finishes ...int) int {
best := finishes[0]
for _,i := range finishes {

if i < best{
best = i

}

}
return best

}