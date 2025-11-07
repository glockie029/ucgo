package structpractice

type File struct {
	Fd   int
	Name string
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}
