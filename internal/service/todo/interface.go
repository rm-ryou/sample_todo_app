package todo

type Getter interface{}

type Modifier interface{}

type Repository interface {
	Getter
	Modifier
}
