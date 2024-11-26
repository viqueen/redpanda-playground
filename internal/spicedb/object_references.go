package spicedb

import (
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/gofrs/uuid"
)

func organisationObject(orgId uuid.UUID) *pb.ObjectReference {
	return &pb.ObjectReference{
		ObjectType: "organisation",
		ObjectId:   orgId.String(),
	}
}

func userObject(userId uuid.UUID) *pb.ObjectReference {
	return &pb.ObjectReference{
		ObjectType: "user",
		ObjectId:   userId.String(),
	}
}

func groupObject(groupId uuid.UUID) *pb.ObjectReference {
	return &pb.ObjectReference{
		ObjectType: "group",
		ObjectId:   groupId.String(),
	}
}

func documentObject(documentId uuid.UUID) *pb.ObjectReference {
	return &pb.ObjectReference{
		ObjectType: "document",
		ObjectId:   documentId.String(),
	}
}
