/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
 var intersection = function(nums1, nums2) {
  console.log('----------')
  
  const n1 = nums1.filter((val, index, self) => self.indexOf(val) === index)
  const n2 = nums2.filter((val, index, self) => self.indexOf(val) === index)
  
  console.log(n1, "=> ", nums1)
  console.log(n2, "=> ", nums2)


  if (n1.length == 0 || n2.length == 0) {
    return []
  }

  

  const merge = [...n1, ...n2]
  console.log("merge : ==> ", merge)

  let ans = []
  for (let i = 0; i < merge.length; i++) {
    for (let j = i + 1; j < merge.length; j++) {
      if (merge[i] == merge[j]) {
        console.log("hahah : ", merge[i], merge[j])
        ans.push(merge[i])
      }
    }
  }
return ans
};