package spicedb

import (
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"github.com/gofrs/uuid"
	"log"
)

func SetupData(ctx context.Context) error {
	return withClient(ctx, setupOrganisation)
}

func setupOrganisation(ctx context.Context, client *authzed.Client) error {
	orgId := uuid.Must(uuid.NewV4())
	users := make([]uuid.UUID, 0, 10)
	for i := 0; i < 10; i++ {
		userId, err := setupUser(ctx, client, orgId)
		if err != nil {
			return err
		}
		users = append(users, userId)
	}
	groups := make([]uuid.UUID, 0, 10)
	for i := 0; i < 10; i++ {
		members := users[i/10 : i+3/10]
		groupId, err := setupGroup(ctx, client, orgId, members)
		if err != nil {
			return err
		}
		groups = append(groups, groupId)
	}
	documents := make([]uuid.UUID, 0, 10)
	for i := 0; i < 10; i++ {
		documentId, err := setupDocument(ctx, client, orgId, users[i], groups[i/10:i+3/10])
		if err != nil {
			return err
		}
		documents = append(documents, documentId)
	}
	return nil
}

func setupUser(ctx context.Context, client *authzed.Client, orgId uuid.UUID) (uuid.UUID, error) {
	userId := uuid.Must(uuid.NewV4())
	request := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &pb.Relationship{
					Resource: organisationObject(orgId),
					Relation: "member",
					Subject:  &pb.SubjectReference{Object: userObject(userId)},
				},
			},
		},
	}
	response, err := client.WriteRelationships(ctx, request)
	if err != nil {
		return uuid.Nil, err
	}
	log.Printf("User created: %s | token=%s", userId, response.GetWrittenAt())
	return userId, nil
}

func setupGroup(ctx context.Context, client *authzed.Client, orgId uuid.UUID, members []uuid.UUID) (uuid.UUID, error) {
	groupId := uuid.Must(uuid.NewV4())
	request := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &pb.Relationship{
					Resource: groupObject(groupId),
					Relation: "org",
					Subject:  &pb.SubjectReference{Object: organisationObject(orgId)},
				},
			},
		},
	}
	for _, member := range members {
		request.Updates = append(request.Updates, &pb.RelationshipUpdate{
			Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &pb.Relationship{
				Resource: groupObject(groupId),
				Relation: "member",
				Subject:  &pb.SubjectReference{Object: userObject(member)},
			},
		})
	}
	response, err := client.WriteRelationships(ctx, request)
	if err != nil {
		return uuid.Nil, err
	}
	log.Printf("Group created: %s | token=%s", groupId, response.GetWrittenAt())
	return groupId, nil
}

func setupDocument(ctx context.Context, client *authzed.Client, orgId uuid.UUID, owner uuid.UUID, sharedWith []uuid.UUID) (uuid.UUID, error) {
	documentId := uuid.Must(uuid.NewV4())
	request := &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &pb.Relationship{
					Resource: documentObject(documentId),
					Relation: "org",
					Subject:  &pb.SubjectReference{Object: organisationObject(orgId)},
				},
			},
			{
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &pb.Relationship{
					Resource: documentObject(documentId),
					Relation: "owner",
					Subject:  &pb.SubjectReference{Object: userObject(owner)},
				},
			},
		},
	}
	for _, group := range sharedWith {
		request.Updates = append(request.Updates, &pb.RelationshipUpdate{
			Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &pb.Relationship{
				Resource: documentObject(documentId),
				Relation: "shared_with",
				Subject:  &pb.SubjectReference{Object: groupObject(group), OptionalRelation: "member"},
			},
		})
	}
	response, err := client.WriteRelationships(ctx, request)
	if err != nil {
		return uuid.Nil, err
	}
	log.Printf("Document created: %s | token=%s", documentId, response.GetWrittenAt())
	return uuid.Nil, nil
}
