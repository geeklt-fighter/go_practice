package main
import "fmt"

func main() {
	
	var monsters []map[string]string
	monsters = make([]map[string]string,2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "timlo"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "martin"
		monsters[1]["age"] = "400"
	}


	// 動態新增
	newMonster := map[string]string {
		"name":"new monster",
		"age":"200",
	}
	monsters = append(monsters,newMonster)

	fmt.Println(monsters)
}