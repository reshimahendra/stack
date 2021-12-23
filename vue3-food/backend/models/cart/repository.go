package cart

import "gorm.io/gorm"

type CartRepo interface {
    Create (c Cart) (Cart, error)
    Update (id int, c Cart) (Cart, error)
    Delete (id int) error
    FindByID (id int) (Cart, error)
    All() ([]Cart, error)
}

type CartItemRepo interface {
    Create (od CartItem) (CartItem, error)
    Update (id int, c CartItem) (CartItem, error)
    Delete (id int) error
    FindByID (id int) (CartItem, error)
    All() ([]CartItem, error)
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
    err := cr.db.Where("id=?", id).Updates(&c).Error
    return c, err
}

func (cr *cartRepo) Delete(id int) error {
    var c Cart
    err := cr.db.Where("id=?", id).Delete(&c).Error
    return err
}

func (cr *cartRepo) FindByID(id int) (Cart, error) {
    var c Cart
    err := cr.db.Where("id=?", id).First(&c).Error
    return c, err
}

func (cr *cartRepo) All() ([]Cart, error) {
    var cl []Cart

    err := cr.db.Find(&cl).Error
    return cl, err
}


// Order list repo 
type cartItemRepo struct {
    db *gorm.DB
}

func NewCartItemRepo(db *gorm.DB) *cartItemRepo{
    return &cartItemRepo{db}
}

func (cir *cartItemRepo) Create(ol CartItem) (CartItem, error) {
    err := cir.db.Create(&ol).Error
    return ol, err
}

func (cir *cartItemRepo) Update(id int, ol CartItem) (CartItem, error) {
    err := cir.db.Where("id=?", id).Updates(&ol).Error
    return ol, err
}

func (cir *cartItemRepo) Delete(id int) error {
    var ol CartItem
    err := cir.db.Where("id=?", id).Delete(&ol).Error
    return err
}

func (cir *cartItemRepo) FindByID(id int) (CartItem, error) {
    var ol CartItem
    err := cir.db.Where("id=?", id).First(&ol).Error
    return ol, err
}

func (cir *cartItemRepo) All() ([]CartItem, error) {
    var ol []CartItem

    err := cir.db.Find(&ol).Error
    return ol, err
}
