package storage

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore/apiv1"
	firestorepb "cloud.google.com/go/firestore/apiv1/firestorepb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Firestore struct {
	client           *firestore.Client
	Credentials_file string
	Project_id       string
	Database_id      string
}

func (f *Firestore) Close() error {
	return f.client.Close()
}

func (f *Firestore) Connect(ctx context.Context) error {
	opt := option.WithCredentialsFile(f.Credentials_file)
	client, err := firestore.NewClient(ctx, opt)
	if err != nil {
		return err
	}
	f.client = client
	return nil
}

func (f *Firestore) List(req *firestorepb.ListDocumentsRequest, ctx context.Context) ([]interface{}, error) {
	docs := f.client.ListDocuments(ctx, req)
	var documents []interface{}
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}
	return documents, nil
}

func (f *Firestore) Get(req *firestorepb.GetDocumentRequest, ctx context.Context) (*firestorepb.Document, error) {
	resp, err := f.client.GetDocument(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *Firestore) Create(req *firestorepb.GetDocumentRequest, ctx context.Context) (*firestorepb.Document, error) {
	resp, err := f.client.GetDocument(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *Firestore) Update(req *firestorepb.UpdateDocumentRequest, ctx context.Context) (*firestorepb.Document, error) {
	resp, err := f.client.UpdateDocument(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *Firestore) Delete(req *firestorepb.DeleteDocumentRequest, ctx context.Context) error {
	err := f.client.DeleteDocument(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func NewFirestore(ctx context.Context) (*Firestore, error) {
	credentials_file := ctx.Value("credentials_file").(*string)
	if credentials_file == nil {
		return nil, fmt.Errorf("credentials_file is required")
	}
	return &Firestore{
		Credentials_file: *credentials_file,
	}, nil
}
