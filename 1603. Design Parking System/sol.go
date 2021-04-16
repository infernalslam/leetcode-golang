type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{Big: big, Medium: medium, Small: small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && this.Big > 0 {
		this.Big -= 1
		return true
	} else if carType == 2 && this.Medium > 0 {
		this.Medium -= 1
		return true
	} else if carType == 3 && this.Small > 0 {
		this.Small -= 1
		return true
	} else {
		return false
	}
}

/**
* Your ParkingSystem object will be instantiated and called as such:
* obj := Constructor(big, medium, small);
* param_1 := obj.AddCar(carType);
 */