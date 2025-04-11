package document

type DocumentRepository interface {
	CreateDocument(document *Document) error
}
