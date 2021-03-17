func kidsWithCandies(candies []int, extraCandies int) []bool {
  return sol(candies, extraCandies)
}

func sol(c []int, ex int) []bool {
  
  max := findMax(c)
  ans := []bool{}
  
  for _, v := range c {
    res := false
    if ex + v >= max {
      res = true
    }
    ans = append(ans, res)
  }
  
  return ans
}

func findMax(list []int) int {
  max := list[0]
  for _, v := range list {
    if max < v {
      max = v
    }
  }
  return max
}