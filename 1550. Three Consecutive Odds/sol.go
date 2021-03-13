func threeConsecutiveOdds(arr []int) bool {
  return sol(arr)
}

func sol (arr []int) bool {
  
  
  check := 0
  
  for i := 0; i < len(arr); i++ {
    if arr[i] % 2 != 0 {
      check++
    } else {
      check = 0
    }
    if check == 3 {
      return true
    }
  }

  
  return false
  
  
}