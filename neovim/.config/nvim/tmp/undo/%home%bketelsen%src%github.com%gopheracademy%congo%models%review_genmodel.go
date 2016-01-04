Vim�UnDo� 7��L^�.2�b����Y.%���!s7��   �                                   V���     _�                             ����                                                                                                                                                                                                                                                                                                                                                             Vsn�     �               �   L//************************************************************************//   // congo: Models   //   .// Generated with goagen v0.0.1, command line:   // $ goagen   5// --out=$(GOPATH)/src/github.com/gopheracademy/congo   1// --design=github.com/gopheracademy/congo/design   //   <// The content of this file is auto-generated, DO NOT MODIFY   L//************************************************************************//       package models       import (   	"time"       %	"github.com/gopheracademy/congo/app"   	"github.com/jinzhu/copier"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   )       // app.ReviewModel storage type   // Identifier:   type Review struct {   *	ID               int `gorm:"primary_key"`   	CreatedAt        time.Time   	UpdatedAt        time.Time   	DeletedAt        *time.Time   	Proposal, UserID int   	// Auth   	Password string       
	// OAuth2   	Oauth2Uid      string   	Oauth2Provider string   	Oauth2Token    string   	Oauth2Refresh  string   	Oauth2Expiry   time.Time       	// Confirm   	ConfirmToken string   	Confirmed    bool       	// Lock   	AttemptNumber int64   	AttemptTime   time.Time   	Locked        time.Time       	// Recover   	RecoverToken       string   	RecoverTokenExpiry time.Time   5	Comment            string `json:"comment,omitempty"`   4	Rating             int    `json:"rating,omitempty"`   }       Cfunc ReviewFromCreatePayload(ctx *app.CreateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   /	m.Proposal, UserID = int(ctx.Proposal, UserID)   		return m   }       Cfunc ReviewFromUpdatePayload(ctx *app.UpdateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   		return m   }   %func (m Review) ToApp() *app.Review {   	target := app.Review{}   	copier.Copy(&target, &m)   	return &target   }       "func (m Review) GetRole() string {   	return m.Role   }       type ReviewStorage interface {   #	List(ctx context.Context) []Review   1	One(ctx context.Context, id int) (Review, error)   3	Add(ctx context.Context, o Review) (Review, error)   ,	Update(ctx context.Context, o Review) error   *	Delete(ctx context.Context, id int) error   }       type ReviewDB struct {   	DB gorm.DB   }       /*   S// would prefer to just pass a context in here, but they're all different, so can't   Qfunc ReviewFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   /			return db.Where("proposal_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }   S// would prefer to just pass a context in here, but they're all different, so can't   Qfunc ReviewFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   +			return db.Where("user_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }   */   (func NewReviewDB(db gorm.DB) *ReviewDB {   	return &ReviewDB{DB: db}   }       7func (m *ReviewDB) List(ctx context.Context) []Review {       	var objs []Review   	m.DB.Find(&objs)   	return objs   }       Efunc (m *ReviewDB) One(ctx context.Context, id int) (Review, error) {       	var obj Review       !	err := m.DB.Find(&obj, id).Error   	return obj, err   }       Kfunc (m *ReviewDB) Add(ctx context.Context, model Review) (Review, error) {   !	err := m.DB.Create(&model).Error   	return model, err   }   Dfunc (m *ReviewDB) Update(ctx context.Context, model Review) error {   !	obj, err := m.One(ctx, model.ID)   	if err != nil {   		return err   	}   ,	err = m.DB.Model(&obj).Updates(model).Error   	return err   }   >func (m *ReviewDB) Delete(ctx context.Context, id int) error {   	var obj Review   #	err := m.DB.Delete(&obj, id).Error   	if err != nil {   		return err   	}   	return nil   }       Afunc FilterReviewByProposal(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   "		if o.ProposalID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }       =func FilterReviewByUser(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   		if o.UserID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Vso     �               �   L//************************************************************************//   // congo: Models   //   .// Generated with goagen v0.0.1, command line:   // $ goagen   5// --out=$(GOPATH)/src/github.com/gopheracademy/congo   1// --design=github.com/gopheracademy/congo/design   //   <// The content of this file is auto-generated, DO NOT MODIFY   L//************************************************************************//       package models       import (   	"time"       %	"github.com/gopheracademy/congo/app"   	"github.com/jinzhu/copier"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   )       // app.ReviewModel storage type   // Identifier:   type Review struct {   *	ID               int `gorm:"primary_key"`   	CreatedAt        time.Time   	UpdatedAt        time.Time   	DeletedAt        *time.Time   	Proposal, UserID int   	// Auth   	Password string       
	// OAuth2   	Oauth2Uid      string   	Oauth2Provider string   	Oauth2Token    string   	Oauth2Refresh  string   	Oauth2Expiry   time.Time       	// Confirm   	ConfirmToken string   	Confirmed    bool       	// Lock   	AttemptNumber int64   	AttemptTime   time.Time   	Locked        time.Time       	// Recover   	RecoverToken       string   	RecoverTokenExpiry time.Time   5	Comment            string `json:"comment,omitempty"`   4	Rating             int    `json:"rating,omitempty"`   }       Cfunc ReviewFromCreatePayload(ctx *app.CreateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   /	m.Proposal, UserID = int(ctx.Proposal, UserID)   		return m   }       Cfunc ReviewFromUpdatePayload(ctx *app.UpdateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   		return m   }   %func (m Review) ToApp() *app.Review {   	target := app.Review{}   	copier.Copy(&target, &m)   	return &target   }       "func (m Review) GetRole() string {   	return m.Role   }       type ReviewStorage interface {   #	List(ctx context.Context) []Review   1	One(ctx context.Context, id int) (Review, error)   3	Add(ctx context.Context, o Review) (Review, error)   ,	Update(ctx context.Context, o Review) error   *	Delete(ctx context.Context, id int) error   }       type ReviewDB struct {   	DB gorm.DB   }       /*   S// would prefer to just pass a context in here, but they're all different, so can't   Qfunc ReviewFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   /			return db.Where("proposal_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }   S// would prefer to just pass a context in here, but they're all different, so can't   Qfunc ReviewFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   +			return db.Where("user_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }   */   (func NewReviewDB(db gorm.DB) *ReviewDB {   	return &ReviewDB{DB: db}   }       7func (m *ReviewDB) List(ctx context.Context) []Review {       	var objs []Review   	m.DB.Find(&objs)   	return objs   }       Efunc (m *ReviewDB) One(ctx context.Context, id int) (Review, error) {       	var obj Review       !	err := m.DB.Find(&obj, id).Error   	return obj, err   }       Kfunc (m *ReviewDB) Add(ctx context.Context, model Review) (Review, error) {   !	err := m.DB.Create(&model).Error   	return model, err   }   Dfunc (m *ReviewDB) Update(ctx context.Context, model Review) error {   !	obj, err := m.One(ctx, model.ID)   	if err != nil {   		return err   	}   ,	err = m.DB.Model(&obj).Updates(model).Error   	return err   }   >func (m *ReviewDB) Delete(ctx context.Context, id int) error {   	var obj Review   #	err := m.DB.Delete(&obj, id).Error   	if err != nil {   		return err   	}   	return nil   }       Afunc FilterReviewByProposal(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   "		if o.ProposalID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }       =func FilterReviewByUser(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   		if o.UserID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             VsoL     �               �   L//************************************************************************//   // congo: Models   //   .// Generated with goagen v0.0.1, command line:   // $ goagen   5// --out=$(GOPATH)/src/github.com/gopheracademy/congo   1// --design=github.com/gopheracademy/congo/design   //   <// The content of this file is auto-generated, DO NOT MODIFY   L//************************************************************************//       package models       import (   	"time"       %	"github.com/gopheracademy/congo/app"   	"github.com/jinzhu/copier"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   )       // app.ReviewModel storage type   // Identifier:   type Review struct {   *	ID               int `gorm:"primary_key"`   	CreatedAt        time.Time   	UpdatedAt        time.Time   	DeletedAt        *time.Time   	Proposal, UserID int   3	Comment          string `json:"comment,omitempty"`   2	Rating           int    `json:"rating,omitempty"`   }       Cfunc ReviewFromCreatePayload(ctx *app.CreateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   /	m.Proposal, UserID = int(ctx.Proposal, UserID)   		return m   }       Cfunc ReviewFromUpdatePayload(ctx *app.UpdateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   		return m   }   %func (m Review) ToApp() *app.Review {   	target := app.Review{}   	copier.Copy(&target, &m)   	return &target   }       type ReviewStorage interface {   #	List(ctx context.Context) []Review   1	One(ctx context.Context, id int) (Review, error)   3	Add(ctx context.Context, o Review) (Review, error)   ,	Update(ctx context.Context, o Review) error   *	Delete(ctx context.Context, id int) error   }       type ReviewDB struct {   	DB gorm.DB   }       /*   S// would prefer to just pass a context in here, but they're all different, so can't   Qfunc ReviewFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   /			return db.Where("proposal_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }   S// would prefer to just pass a context in here, but they're all different, so can't   Qfunc ReviewFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   +			return db.Where("user_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }   */   (func NewReviewDB(db gorm.DB) *ReviewDB {   	return &ReviewDB{DB: db}   }       7func (m *ReviewDB) List(ctx context.Context) []Review {       	var objs []Review   	m.DB.Find(&objs)   	return objs   }       Efunc (m *ReviewDB) One(ctx context.Context, id int) (Review, error) {       	var obj Review       !	err := m.DB.Find(&obj, id).Error   	return obj, err   }       Kfunc (m *ReviewDB) Add(ctx context.Context, model Review) (Review, error) {   !	err := m.DB.Create(&model).Error   	return model, err   }   Dfunc (m *ReviewDB) Update(ctx context.Context, model Review) error {   !	obj, err := m.One(ctx, model.ID)   	if err != nil {   		return err   	}   ,	err = m.DB.Model(&obj).Updates(model).Error   	return err   }   >func (m *ReviewDB) Delete(ctx context.Context, id int) error {   	var obj Review   #	err := m.DB.Delete(&obj, id).Error   	if err != nil {   		return err   	}   	return nil   }       Afunc FilterReviewByProposal(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   "		if o.ProposalID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }       =func FilterReviewByUser(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   		if o.UserID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }5�_�                    D        ����                                                                                                                                                                                                                                                                                                                                                             Vso�     �               �   L//************************************************************************//   // congo: Models   //   .// Generated with goagen v0.0.1, command line:   // $ goagen   5// --out=$(GOPATH)/src/github.com/gopheracademy/congo   1// --design=github.com/gopheracademy/congo/design   //   <// The content of this file is auto-generated, DO NOT MODIFY   L//************************************************************************//       package models       import (   	"time"       %	"github.com/gopheracademy/congo/app"   	"github.com/jinzhu/copier"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   )       // app.ReviewModel storage type   // Identifier:   type Review struct {   $	ID         int `gorm:"primary_key"`   	CreatedAt  time.Time   	UpdatedAt  time.Time   	DeletedAt  *time.Time   	ProposalID int   	UserID     int   -	Comment    string `json:"comment,omitempty"`   ,	Rating     int    `json:"rating,omitempty"`   }       Cfunc ReviewFromCreatePayload(ctx *app.CreateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   /	m.Proposal, UserID = int(ctx.Proposal, UserID)   		return m   }       Cfunc ReviewFromUpdatePayload(ctx *app.UpdateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   		return m   }   %func (m Review) ToApp() *app.Review {   	target := app.Review{}   	copier.Copy(&target, &m)   	return &target   }       type ReviewStorage interface {   #	List(ctx context.Context) []Review   1	One(ctx context.Context, id int) (Review, error)   3	Add(ctx context.Context, o Review) (Review, error)   ,	Update(ctx context.Context, o Review) error   *	Delete(ctx context.Context, id int) error   }       type ReviewDB struct {   	DB gorm.DB   }       /*   S// would prefer to just pass a context in here, but they're all different, so can't   Qfunc ReviewFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   /			return db.Where("proposal_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }   S// would prefer to just pass a context in here, but they're all different, so can't   Qfunc ReviewFilter(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   +			return db.Where("user_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }   */   (func NewReviewDB(db gorm.DB) *ReviewDB {   	return &ReviewDB{DB: db}   }       7func (m *ReviewDB) List(ctx context.Context) []Review {       	var objs []Review   	m.DB.Find(&objs)   	return objs   }       Efunc (m *ReviewDB) One(ctx context.Context, id int) (Review, error) {       	var obj Review       !	err := m.DB.Find(&obj, id).Error   	return obj, err   }       Kfunc (m *ReviewDB) Add(ctx context.Context, model Review) (Review, error) {   !	err := m.DB.Create(&model).Error   	return model, err   }   Dfunc (m *ReviewDB) Update(ctx context.Context, model Review) error {   !	obj, err := m.One(ctx, model.ID)   	if err != nil {   		return err   	}   ,	err = m.DB.Model(&obj).Updates(model).Error   	return err   }   >func (m *ReviewDB) Delete(ctx context.Context, id int) error {   	var obj Review   #	err := m.DB.Delete(&obj, id).Error   	if err != nil {   		return err   	}   	return nil   }       Afunc FilterReviewByProposal(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   "		if o.ProposalID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }       =func FilterReviewByUser(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   		if o.UserID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }5�_�                    (        ����                                                                                                                                                                                                                                                                                                                                                             Vspr     �               �   L//************************************************************************//   // congo: Models   //   .// Generated with goagen v0.0.1, command line:   // $ goagen   5// --out=$(GOPATH)/src/github.com/gopheracademy/congo   1// --design=github.com/gopheracademy/congo/design   //   <// The content of this file is auto-generated, DO NOT MODIFY   L//************************************************************************//       package models       import (   	"time"       %	"github.com/gopheracademy/congo/app"   	"github.com/jinzhu/copier"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   )       // app.ReviewModel storage type   // Identifier:   type Review struct {   $	ID         int `gorm:"primary_key"`   	CreatedAt  time.Time   	UpdatedAt  time.Time   	DeletedAt  *time.Time   	ProposalID int   	UserID     int   -	Comment    string `json:"comment,omitempty"`   ,	Rating     int    `json:"rating,omitempty"`   }       Cfunc ReviewFromCreatePayload(ctx *app.CreateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   /	m.Proposal, UserID = int(ctx.Proposal, UserID)   		return m   }       Cfunc ReviewFromUpdatePayload(ctx *app.UpdateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   		return m   }   %func (m Review) ToApp() *app.Review {   	target := app.Review{}   	copier.Copy(&target, &m)   	return &target   }       type ReviewStorage interface {   #	List(ctx context.Context) []Review   1	One(ctx context.Context, id int) (Review, error)   3	Add(ctx context.Context, o Review) (Review, error)   ,	Update(ctx context.Context, o Review) error   *	Delete(ctx context.Context, id int) error   }       type ReviewDB struct {   	DB gorm.DB   }       S// would prefer to just pass a context in here, but they're all different, so can't   [func ReviewFilterByProposal(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   /			return db.Where("proposal_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }       S// would prefer to just pass a context in here, but they're all different, so can't   Wfunc ReviewFilterByUser(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   +			return db.Where("user_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }       (func NewReviewDB(db gorm.DB) *ReviewDB {   	return &ReviewDB{DB: db}   }       7func (m *ReviewDB) List(ctx context.Context) []Review {       	var objs []Review   	m.DB.Find(&objs)   	return objs   }       Efunc (m *ReviewDB) One(ctx context.Context, id int) (Review, error) {       	var obj Review       !	err := m.DB.Find(&obj, id).Error   	return obj, err   }       Kfunc (m *ReviewDB) Add(ctx context.Context, model Review) (Review, error) {   !	err := m.DB.Create(&model).Error   	return model, err   }   Dfunc (m *ReviewDB) Update(ctx context.Context, model Review) error {   !	obj, err := m.One(ctx, model.ID)   	if err != nil {   		return err   	}   ,	err = m.DB.Model(&obj).Updates(model).Error   	return err   }   >func (m *ReviewDB) Delete(ctx context.Context, id int) error {   	var obj Review   #	err := m.DB.Delete(&obj, id).Error   	if err != nil {   		return err   	}   	return nil   }       Afunc FilterReviewByProposal(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   "		if o.ProposalID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }       =func FilterReviewByUser(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   		if o.UserID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }5�_�                     4        ����                                                                                                                                                                                                                                                                                                                                                             V���     �               �   L//************************************************************************//   // congo: Models   //   .// Generated with goagen v0.0.1, command line:   // $ goagen   5// --out=$(GOPATH)/src/github.com/gopheracademy/congo   1// --design=github.com/gopheracademy/congo/design   //   <// The content of this file is auto-generated, DO NOT MODIFY   L//************************************************************************//       package models       import (   	"time"       %	"github.com/gopheracademy/congo/app"   	"github.com/jinzhu/copier"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   )       // app.ReviewModel storage type   // Identifier:   type Review struct {   $	ID         int `gorm:"primary_key"`   	CreatedAt  time.Time   	UpdatedAt  time.Time   	DeletedAt  *time.Time   	ProposalID int   	UserID     int   -	Comment    string `json:"comment,omitempty"`   ,	Rating     int    `json:"rating,omitempty"`   }       Cfunc ReviewFromCreatePayload(ctx *app.CreateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)       #	m.ProposalID = int(ctx.ProposalID)   	m.UserID = int(ctx.UserID)   		return m   }       Cfunc ReviewFromUpdatePayload(ctx *app.UpdateReviewContext) Review {   	payload := ctx.Payload   	m := Review{}   	copier.Copy(&m, payload)   		return m   }   %func (m Review) ToApp() *app.Review {   	target := app.Review{}   	copier.Copy(&target, &m)   	return &target   }       type ReviewStorage interface {   #	List(ctx context.Context) []Review   1	One(ctx context.Context, id int) (Review, error)   3	Add(ctx context.Context, o Review) (Review, error)   ,	Update(ctx context.Context, o Review) error   *	Delete(ctx context.Context, id int) error   }       type ReviewDB struct {   	DB gorm.DB   }       S// would prefer to just pass a context in here, but they're all different, so can't   [func ReviewFilterByProposal(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   /			return db.Where("proposal_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }       S// would prefer to just pass a context in here, but they're all different, so can't   Wfunc ReviewFilterByUser(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   +			return db.Where("user_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }       (func NewReviewDB(db gorm.DB) *ReviewDB {   	return &ReviewDB{DB: db}   }       7func (m *ReviewDB) List(ctx context.Context) []Review {       	var objs []Review   	m.DB.Find(&objs)   	return objs   }       Efunc (m *ReviewDB) One(ctx context.Context, id int) (Review, error) {       	var obj Review       !	err := m.DB.Find(&obj, id).Error   	return obj, err   }       Kfunc (m *ReviewDB) Add(ctx context.Context, model Review) (Review, error) {   !	err := m.DB.Create(&model).Error   	return model, err   }   Dfunc (m *ReviewDB) Update(ctx context.Context, model Review) error {   !	obj, err := m.One(ctx, model.ID)   	if err != nil {   		return err   	}   ,	err = m.DB.Model(&obj).Updates(model).Error   	return err   }   >func (m *ReviewDB) Delete(ctx context.Context, id int) error {   	var obj Review   #	err := m.DB.Delete(&obj, id).Error   	if err != nil {   		return err   	}   	return nil   }       Afunc FilterReviewByProposal(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   "		if o.ProposalID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }       =func FilterReviewByUser(parent int, list []Review) []Review {   	filtered := make([]Review, 0)   	for _, o := range list {   		if o.UserID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }5��