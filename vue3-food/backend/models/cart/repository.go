package cart

import "gorm.io/gorm"

type CartRepo interface {
    Create (c Cart) (Cart, error)
    Update (id int, c Cart) (Cart, error)
    Delete (id int) error
    FindByID (id int) (Cart, error)
    All() ([]Cart, error)
}

type OrderListRepo interface {
    Create (od OrderList) (OrderList, error)
    Update (id int, c OrderList) (OrderList, error)
    Delete (id int) error
    FindByID (id int) (OrderList, error)
    All() ([]OrderList, error)
}


// Cart Repository
type cartRepo struct {
    db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *cartRepo {
    return &cartRepo{db}
}

func (cr *cartRepo) Create(c Cart) (Cart, error) {
    err := cr.db.Create(&c).Error
    return c, err
}

func (cr *cartRepo) Update(id int, c Cart) (Cart, error) {
    err := cr.db.Where("created_at IS NULL and id=?", id).Updates(&c).Error
    return c, err
}

func (cr *cartRepo) Delete(id int) error {
    var c Cart
    err := cr.db.Where("created_at IS NULL and id=?", id).Delete(&c).Error
    return err
}

func (cr *cartRepo) FindByID(id int) (Cart, error) {
    var c Cart
    err := cr.db.Where("created_at IS NULL AND id=?", id).First(&c).Error
    return c, err
}

func (cr *cartRepo) All() ([]Cart, error) {
    var cl []Cart

    err := cr.db.Find(&cl).Error
    return cl, err
}


// Order list repo 
type orderListRepo struct {
    db *gorm.DB
}

func NewOrderListRepo(db *gorm.DB) *orderListRepo{
    return &orderListRepo{db}
}

func (olr *orderListRepo) Create(ol OrderList) (OrderList, error) {
    err := olr.db.Create(&ol).Error
    return ol, err
}

func (olr *orderListRepo) Update(id int, ol OrderList) (OrderList, error) {
    err := olr.db.Where("created_at IS NULL and id=?", id).Updates(&ol).Error
    return ol, err
}

func (olr *orderListRepo) Delete(id int) error {
    var ol OrderList
    err := olr.db.Where("created_at IS NULL and id=?", id).Delete(&ol).Error
    return err
}

func (olr *orderListRepo) FindByID(id int) (OrderList, error) {
    var ol OrderList
    err := olr.db.Where("created_at IS NULL AND id=?", id).First(&ol).Error
    return ol, err
}

func (olr *orderListRepo) All() ([]OrderList, error) {
    var ol []OrderList

    err := olr.db.Find(&ol).Error
    return ol, err
}
