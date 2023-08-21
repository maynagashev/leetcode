package can_place_flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
      return true
    }
    var prev, next int
    for i, current := range flowerbed {
      prev, next = 0, 0
      if i>0 {
        prev = flowerbed[i-1]
      }
      if i < len(flowerbed) -1 {
        next = flowerbed[i+1]
      }
      if current == 0 && prev == 0 && next == 0 {
         n--
         if n==0 {
           return true
         }
         flowerbed[i] = 1 // помечаем как занятое место
      }
    }
    return false
}