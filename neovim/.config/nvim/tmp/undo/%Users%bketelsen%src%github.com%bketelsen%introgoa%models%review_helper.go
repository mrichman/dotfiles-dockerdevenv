Vim�UnDo� �@�B�����N�l� �#�wع�EI��   u                                   V�T�    _�                     G        ����                                                                                                                                                                                                                                                                                                                                                             V�T�     �   F   I   r       5�_�                    H       ����                                                                                                                                                                                                                                                                                                                                                             V�T�     �   G   I   s      // START OMITj5�_�                    \        ����                                                                                                                                                                                                                                                                                                                                                             V�T�     �   [   ]   s       5�_�                             ����                                                                                                                                                                                                                                                                                                                                                             V�T�    �               s   L//************************************************************************//   // API "congo": Model Helpers   //   .// Generated with goagen v0.0.1, command line:   // $ goagen   5// --out=$(GOPATH)/src/github.com/gopheracademy/congo   1// --design=github.com/gopheracademy/congo/design   //   <// The content of this file is auto-generated, DO NOT MODIFY   L//************************************************************************//       package models       import (   	"github.com/goadesign/goa"   %	"github.com/gopheracademy/congo/app"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   	"time"   )        // MediaType Retrieval Functions       0// ListReview returns an array of view: default.   ^func (m *ReviewDB) ListReview(ctx context.Context, proposalID int, userID int) []*app.Review {   R	defer goa.MeasureSince([]string{"goa", "db", "review", "listreview"}, time.Now())       	var native []*Review   	var objs []*app.Review   �	err := m.Db.Scopes(ReviewFilterByProposal(proposalID, &m.Db), ReviewFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error       	if err != nil {   >		goa.Error(ctx, "error listing Review", "error", err.Error())   		return objs   	}       	for _, t := range native {   )		objs = append(objs, t.ReviewToReview())   	}       	return objs   }       >// ReviewToReview returns the Review representation of Review.   /func (m *Review) ReviewToReview() *app.Review {   	review := &app.Review{}   	review.Comment = m.Comment   	review.ID = &m.ID   	review.Rating = &m.Rating       	return review   }       /// OneReview returns an array of view: default.   lfunc (m *ReviewDB) OneReview(ctx context.Context, id int, proposalID int, userID int) (*app.Review, error) {   Q	defer goa.MeasureSince([]string{"goa", "db", "review", "onereview"}, time.Now())       	var native Review   �	err := m.Db.Scopes(ReviewFilterByProposal(proposalID, &m.Db), ReviewFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("Proposal").Preload("User").Where("id = ?", id).Find(&native).Error       1	if err != nil && err != gorm.ErrRecordNotFound {   >		goa.Error(ctx, "error getting Review", "error", err.Error())   		return nil, err   	}       !	view := *native.ReviewToReview()   	return &view, err   }        // MediaType Retrieval Functions       // START OMIT   1// ListReviewLink returns an array of view: link.   ffunc (m *ReviewDB) ListReviewLink(ctx context.Context, proposalID int, userID int) []*app.ReviewLink {   V	defer goa.MeasureSince([]string{"goa", "db", "review", "listreviewlink"}, time.Now())       	var native []*Review   	var objs []*app.ReviewLink   �	err := m.Db.Scopes(ReviewFilterByProposal(proposalID, &m.Db), ReviewFilterByUser(userID, &m.Db)).Table(m.TableName()).Find(&native).Error       	if err != nil {   >		goa.Error(ctx, "error listing Review", "error", err.Error())   		return objs   	}       	for _, t := range native {   -		objs = append(objs, t.ReviewToReviewLink())   	}       	return objs   }   // END OMIT   B// ReviewToReviewLink returns the Review representation of Review.   7func (m *Review) ReviewToReviewLink() *app.ReviewLink {   	review := &app.ReviewLink{}   	review.ID = &m.ID       	return review   }       0// OneReviewLink returns an array of view: link.   tfunc (m *ReviewDB) OneReviewLink(ctx context.Context, id int, proposalID int, userID int) (*app.ReviewLink, error) {   U	defer goa.MeasureSince([]string{"goa", "db", "review", "onereviewlink"}, time.Now())       	var native Review   �	err := m.Db.Scopes(ReviewFilterByProposal(proposalID, &m.Db), ReviewFilterByUser(userID, &m.Db)).Table(m.TableName()).Preload("Proposal").Preload("User").Where("id = ?", id).Find(&native).Error       1	if err != nil && err != gorm.ErrRecordNotFound {   >		goa.Error(ctx, "error getting Review", "error", err.Error())   		return nil, err   	}       %	view := *native.ReviewToReviewLink()   	return &view, err   }5��