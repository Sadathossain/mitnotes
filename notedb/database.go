package notedb

type NoteDB interface {
	GetAllNotes() ([]string, error)
	SaveNote(string) error
	DeleteNote(string) error
	GetHealthStatus() map[string]string
	RegisterMetrics()
}
