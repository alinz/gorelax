package util

//MapArray Comment TODO
type MapArray struct {
	collection map[string][]interface{}
}

//Add Comment TODO
func (ma *MapArray) Add(key string, value interface{}) {
	//example can be found at http://play.golang.org/p/nKljWsVZs2
	ma.collection[key] = append(ma.collection[key], value)
}

//Get Comment TODO
func (ma *MapArray) Get(key string) ([]interface{}, bool) {
	result, ok := ma.collection[key]
	return result, ok
}

//Iterate Comment TODO
func (ma *MapArray) Iterate(key string, callback func(value interface{}, index int) bool) {
	if array, ok := ma.collection[key]; ok {
		for index, value := range array {
			if !callback(value, index) {
				break
			}
		}
	}
}

//MakeMapArray Comment TODO
func MakeMapArray() MapArray {
	return MapArray{make(map[string][]interface{})}
}
