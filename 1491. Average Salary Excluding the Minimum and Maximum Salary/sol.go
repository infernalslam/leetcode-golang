func average(salary []int) float64 {
  return sol(salary)
}

func sol (salary []int) float64 {
  
    // find min max
  min, max := findMinMaxSalary(salary)
  
  sum := 0.000
  count := 0
  
  for _, v := range salary {
    if v == min || v == max {
      continue
    }
    sum += float64(v)
    count++
  }
  
  
  ans := sum / float64(count)
  
  
  return ans
  
}

func findMinMaxSalary(s []int) (int, int) {
  min := s[0]
  max := s[0]
  
  for _, v := range s {
    if v < min {
      min = v
    }
    
    if v > max {
      max = v
    }
    
  }
  
  return min, max
}