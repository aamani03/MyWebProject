package DBInterface

type User struct {
	UserID   int
	Name     string
	Email    string
	Password string
}

func (u User) f() (int32, int8) {
	//TODO implement me
	panic("implement me")
}

type Task struct {
	TaskID   int
	TaskName string
	TaskDesc string
	UserID   int
	StatusID int
}

type Status struct {
	StatusID   int
	StatusName string
	Color      string
}

type Notification struct {
	NotificationID int
	UserID         int
	TaskID         int
}

type UserService interface {
}

type TaskService interface {
}

type StatusService interface {
}

type NotificationService interface {
}
