package dbInterface

type User struct {
	UserID   int    `json:"user_id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email_id"`
	Password string `json:"password_id"`
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
