package StoreData

import "fmt"

// with new key word --> only memory is allocated and no initialization is done
// with make key word --> not only memory is allocated but also initialization is done
func StoreValue() {

	//Map of string keys and int values
	//The below declation basically means that we are allocating the memory and no initialization is done . Its short hand for new key word.
	//The below set of code will give an error (panic) if we try to add a key value pair to the map --> * assignment to entry in nil map.*  This is because the map is not initialized though the memory has been allocated

	// var score map[string]int
	// score["hitesh"] = 100
	// fmt.Println(score)

	userInfo := make(map[string]string)
	userInfo["name"] = "Hitesh"
	userInfo["age"] = "30"
	userInfo["city"] = "Bangalore"
	userInfo["country"] = "India"
	userInfo["email"] = "abc@gma.com"
	userInfo["phone"] = "1234567890"
	fmt.Println(userInfo)

	//delete(userInfo , "age") --> This will delete the key value pair from the map

	// This syntax is used to iterate over the map
	for k, v := range userInfo {
		if k == "phone" {
			delete(userInfo, k)
		}
		fmt.Println(k, v)
	}

	fmt.Println(userInfo)

}
