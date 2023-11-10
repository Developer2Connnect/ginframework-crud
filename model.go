package main

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Model struct {
	todos []Todo
}

func NewModel() *Model {
	return &Model{
		todos: []Todo{},
	}
}

func (m *Model) GetTodos() []Todo {
	return m.todos
}

func (m *Model) GetTodoByID(id int) (*Todo, bool) {
	for _, todo := range m.todos {
		if todo.ID == id {
			return &todo, true
		}
	}
	return nil, false
}

func (m *Model) AddTodo(todo Todo) {
	// Assign a unique ID
	todo.ID = len(m.todos) + 1
	m.todos = append(m.todos, todo)
}

func (m *Model) UpdateTodoByID(id int, updatedTodo Todo) bool {
	for i, todo := range m.todos {
		if todo.ID == id {
			m.todos[i] = updatedTodo
			return true
		}
	}
	return false
}

func (m *Model) DeleteTodoByID(id int) bool {
	for i, todo := range m.todos {
		if todo.ID == id {
			// Remove the todo from the slice
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			return true
		}
	}
	return false
}
