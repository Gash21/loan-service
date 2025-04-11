package document

type DocumentRepository interface {
	Create(document *Document) error
}
