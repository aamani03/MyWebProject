package dbInterface

type User struct {
	UserID   int
	Name     string
	Email    string
	Password string
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

// type StatusService interface {
// }

// type NotificationService interface {
// }
