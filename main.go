package main

func main() {
	var myTodos Todos
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&myTodos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&myTodos)
	storage.save(myTodos)
}
