package main

import (
	"../proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net"
	"os"
	"os/signal"
)

type server struct {
	session *mgo.Session
}

type blogItem struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	AuthorID string        `json:"author_id" bson:"author_id"`
	Content  string        `json:"content" bson:"content"`
	Title    string        `json:"title" bson:"title"`
}

func (s server) CreateBlog(ctx context.Context, request *proto.CreateBlogRequest) (*proto.CreateBlogResponse, error) {
	blog := request.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
		ID:       bson.NewObjectId(),
	}

	e := s.session.DB("mydb").C("blog").Insert(data)
	if e != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error! Failed to save to MongoDb: %v", e))
	}

	result := blogItem{}
	if e := s.session.DB("mydb").C("blog").FindId(data.ID).One(&result); e != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error! Failed to read new item from MongoDb: %v", e))
	}

	response := &proto.CreateBlogResponse{
		Blog: &proto.Blog{
			Id:       result.ID.Hex(),
			AuthorId: result.AuthorID,
			Title:    result.Title,
			Content:  result.Content,
		},
	}

	return response, nil
}

func (s server) ReadBlog(ctx context.Context, request *proto.ReadBlogRequest) (*proto.ReadBlogResponse, error) {
	blogId := bson.ObjectIdHex(request.GetBlogId())

	result := blogItem{}
	if e := s.session.DB("mydb").C("blog").FindId(blogId).One(&result); e != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error! Failed to read item from MongoDb: %v", e))
	}

	response := &proto.ReadBlogResponse{
		Blog: &proto.Blog{
			Id:       result.ID.Hex(),
			AuthorId: result.AuthorID,
			Title:    result.Title,
			Content:  result.Content,
		},
	}

	return response, nil
}

func (s server) UpdateBlog(ctx context.Context, request *proto.UpdateBlogRequest) (*proto.UpdateBlogResponse, error) {
	blog := request.GetBlog()
	blogId := bson.ObjectIdHex(blog.GetId())

	selection := bson.M{
		"_id": blogId,
	}
	change := bson.M{
		"$set": bson.M{
			"author_id": blog.GetAuthorId(),
			"title":     blog.GetTitle(),
			"content":   blog.GetContent(),
		},
	}

	if e := s.session.DB("mydb").C("blog").Update(selection, change); e != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error! Failed to update of MongoDb: %v", e))
	}

	result := blogItem{}
	if e := s.session.DB("mydb").C("blog").FindId(blogId).One(&result); e != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error! Failed to read item from MongoDb: %v", e))
	}

	response := &proto.UpdateBlogResponse{
		Blog: &proto.Blog{
			Id:       result.ID.Hex(),
			AuthorId: result.AuthorID,
			Title:    result.Title,
			Content:  result.Content,
		},
	}

	return response, nil
}

func (s server) DeleteBlog(ctx context.Context, request *proto.DeleteBlogRequest) (*proto.DeleteBlogResponse, error) {
	blogId := bson.ObjectIdHex(request.GetBlogId())

	selection := bson.M{
		"_id": blogId,
	}

	if e := s.session.DB("mydb").C("blog").Remove(selection); e != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error! Failed to remove from MongoDb: %v", e))
	}

	response := &proto.DeleteBlogResponse{
		BlogId: blogId.Hex(),
	}

	return response, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("The Blog Service Server is starting!")

	log.Println("Connecting to MongoDb")
	session, e := mgo.Dial("mongodb://localhost")
	if e != nil {
		log.Fatalf("Failed to dial to MongoDb: %v", e)
	}

	listener, e := net.Listen("tcp", "0.0.0.0:50051")
	if e != nil {
		log.Fatalf("Failed to listen: %v", e)
	}

	server := &server{
		session: session,
	}

	s := grpc.NewServer()
	proto.RegisterBlogServiceServer(s, server)

	// Register reflection service on gRPC server
	reflection.Register(s)

	go func() {
		e = s.Serve(listener)
		if e != nil {
			log.Fatalf("Failed to serve: %v", e)
		}
	}()

	// Wait for Control C to exit
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Block until a signal is received
	<-signals

	log.Println("Closing the MongoDb Connection!")
	session.Close()

	log.Println("Stopping the Blog Service Server!")
	s.Stop()

	log.Println("Closing the Listener!")
	e = listener.Close()
	if e != nil {
		log.Fatalf("Failed to close listener: %v", e)
	}

	log.Println("The Blog Service Server is shut down!")
}
