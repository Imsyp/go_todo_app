package main

func main() {
	var id int = 1
	// TaskID 타입에 변환한 후에 대입하고 있으므로 문제 X
	_ = Task{ID: TaskID(id)}
	//빌드 오류
	
}