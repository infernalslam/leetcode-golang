func singleNumber(nums []int) int {
  return sol(nums)
}

func sol(nums []int) int {
  
  ans := 0
  
  for i := 0; i < len(nums); i++ {
    check := 0
    
    for j := 0; j < len(nums); j++ {
      if nums[i] == nums[j] {
        check++
      }
    }
    
    if check == 1 {
      ans = nums[i]
    }
  }
  
  
  return ans
  
}