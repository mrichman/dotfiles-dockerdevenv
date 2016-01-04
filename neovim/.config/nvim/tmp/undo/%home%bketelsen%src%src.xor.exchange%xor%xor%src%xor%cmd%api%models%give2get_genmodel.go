Vim�UnDo� <�����eط��§p0�!~�M�>���    q                                   V���     _�                             ����                                                                                                                                                                                                                                                                                                                                                             V��?     �               �   L//************************************************************************//   // xor: Models   //   .// Generated with goagen v0.0.1, command line:   // $ goagen   "// --out=$(GOPATH)/src/xor/cmd/api   // --design=xor/cmd/api/design   //   <// The content of this file is auto-generated, DO NOT MODIFY   L//************************************************************************//       package models       import (   	"time"   	"xor/cmd/api/app"       	"github.com/jinzhu/copier"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   )       !// app.Give2getModel storage type   // Identifier:   type Give2get struct {   )	ID              int `gorm:"primary_key"`   	CreatedAt       time.Time   	UpdatedAt       time.Time   	DeletedAt       *time.Time   	PortfolioID     int   	DatatypeID      int   	DatatypefieldID int   1	Provided        bool `json:"provided,omitempty"`   }       Ifunc Give2getFromCreatePayload(ctx *app.CreateGive2getContext) Give2get {   	payload := ctx.Payload   	m := Give2get{}   	copier.Copy(&m, payload)       %	m.PortfolioID = int(ctx.PortfolioID)   #	m.DatatypeID = int(ctx.DatatypeID)   -	m.DatatypefieldID = int(ctx.DatatypefieldID)   		return m   }       Ifunc Give2getFromUpdatePayload(ctx *app.UpdateGive2getContext) Give2get {   	payload := ctx.Payload   	m := Give2get{}   	copier.Copy(&m, payload)   		return m   }       )func (m Give2get) ToApp() *app.Give2get {   	target := app.Give2get{}   	copier.Copy(&target, &m)   	return &target   }        type Give2getStorage interface {   %	List(ctx context.Context) []Give2get   3	One(ctx context.Context, id int) (Give2get, error)   7	Add(ctx context.Context, o Give2get) (Give2get, error)   .	Update(ctx context.Context, o Give2get) error   *	Delete(ctx context.Context, id int) error   }       type Give2getDB struct {   	DB gorm.DB   }       S// would prefer to just pass a context in here, but they're all different, so can't   ^func Give2getFilterByPortfolio(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   0			return db.Where("portfolio_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }       S// would prefer to just pass a context in here, but they're all different, so can't   ]func Give2getFilterByDatatype(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   /			return db.Where("datatype_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }       S// would prefer to just pass a context in here, but they're all different, so can't   bfunc Give2getFilterByDatatypefield(parentid int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {   	if parentid > 0 {   %		return func(db *gorm.DB) *gorm.DB {   4			return db.Where("datatypefield_id = ?", parentid)   		}   		} else {   %		return func(db *gorm.DB) *gorm.DB {   			return db   		}   	}   }       ,func NewGive2getDB(db gorm.DB) *Give2getDB {   	return &Give2getDB{DB: db}   }       ;func (m *Give2getDB) List(ctx context.Context) []Give2get {       	var objs []Give2get   	m.DB.Find(&objs)   	return objs   }       Ifunc (m *Give2getDB) One(ctx context.Context, id int) (Give2get, error) {       	var obj Give2get       !	err := m.DB.Find(&obj, id).Error   	return obj, err   }       Qfunc (m *Give2getDB) Add(ctx context.Context, model Give2get) (Give2get, error) {   !	err := m.DB.Create(&model).Error   	return model, err   }       Hfunc (m *Give2getDB) Update(ctx context.Context, model Give2get) error {   !	obj, err := m.One(ctx, model.ID)   	if err != nil {   		return err   	}   ,	err = m.DB.Model(&obj).Updates(model).Error   	return err   }       @func (m *Give2getDB) Delete(ctx context.Context, id int) error {   	var obj Give2get   #	err := m.DB.Delete(&obj, id).Error   	if err != nil {   		return err   	}   	return nil   }       Hfunc FilterGive2getByPortfolio(parent int, list []Give2get) []Give2get {    	filtered := make([]Give2get, 0)   	for _, o := range list {   #		if o.PortfolioID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }       Gfunc FilterGive2getByDatatype(parent int, list []Give2get) []Give2get {    	filtered := make([]Give2get, 0)   	for _, o := range list {   "		if o.DatatypeID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }       Lfunc FilterGive2getByDatatypefield(parent int, list []Give2get) []Give2get {    	filtered := make([]Give2get, 0)   	for _, o := range list {   '		if o.DatatypefieldID == int(parent) {   !			filtered = append(filtered, o)   		}   	}   	return filtered   }5�_�                             ����                                                                                                                                                                                                                                                                                                                                                             V���     �               q   L//************************************************************************//   // xor: Models   //   .// Generated with goagen v0.0.1, command line:   // $ goagen   "// --out=$(GOPATH)/src/xor/cmd/api   // --design=xor/cmd/api/design   //   <// The content of this file is auto-generated, DO NOT MODIFY   L//************************************************************************//       package models       import (   	"database/sql"   	"time"   	"xor/cmd/api/app"       	"github.com/jinzhu/copier"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   )       !// app.Give2getModel storage type   // Identifier:   type Give2get struct {   )	ID              int `gorm:"primary_key"`   	CreatedAt       time.Time   	UpdatedAt       time.Time   	DeletedAt       *time.Time   	Portfolio       Portfolio   	PortfolioID     *sql.NullInt64   	Datatype        Datatype   	DatatypeID      *sql.NullInt64   	Datatypefield   Datatypefield   	DatatypefieldID *sql.NullInt64   1	Provided        bool `json:"provided,omitempty"`   }       Ifunc Give2getFromCreatePayload(ctx *app.CreateGive2getContext) Give2get {   	payload := ctx.Payload   	m := Give2get{}   	copier.Copy(&m, payload)       		return m   }       Ifunc Give2getFromUpdatePayload(ctx *app.UpdateGive2getContext) Give2get {   	payload := ctx.Payload   	m := Give2get{}   	copier.Copy(&m, payload)   		return m   }       )func (m Give2get) ToApp() *app.Give2get {   	target := app.Give2get{}   	copier.Copy(&target, &m)   	return &target   }        type Give2getStorage interface {   %	List(ctx context.Context) []Give2get   3	One(ctx context.Context, id int) (Give2get, error)   7	Add(ctx context.Context, o Give2get) (Give2get, error)   .	Update(ctx context.Context, o Give2get) error   *	Delete(ctx context.Context, id int) error   }       type Give2getDB struct {   	DB gorm.DB   }       ,func NewGive2getDB(db gorm.DB) *Give2getDB {   	return &Give2getDB{DB: db}   }       ;func (m *Give2getDB) List(ctx context.Context) []Give2get {       	var objs []Give2get   	m.DB.Find(&objs)   	return objs   }       Ifunc (m *Give2getDB) One(ctx context.Context, id int) (Give2get, error) {       	var obj Give2get       !	err := m.DB.Find(&obj, id).Error   	return obj, err   }       Qfunc (m *Give2getDB) Add(ctx context.Context, model Give2get) (Give2get, error) {   !	err := m.DB.Create(&model).Error   	return model, err   }       Hfunc (m *Give2getDB) Update(ctx context.Context, model Give2get) error {   !	obj, err := m.One(ctx, model.ID)   	if err != nil {   		return err   	}   ,	err = m.DB.Model(&obj).Updates(model).Error   	return err   }       @func (m *Give2getDB) Delete(ctx context.Context, id int) error {   	var obj Give2get   #	err := m.DB.Delete(&obj, id).Error   	if err != nil {   		return err   	}   	return nil   }5��