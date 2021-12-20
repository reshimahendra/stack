package foods

import "gorm.io/gorm"

// 1. PRODUCT TYPE Repository
type ProductTypeRepo interface {
    Create(pt ProductType) (ProductType, error)
    Update(id int, pt ProductType) (ProductType, error)
    Delete(id int) error
    First() (ProductType, error)
    Last() (ProductType, error)
    FindByID(id int) (ProductType, error)
    All() ([]ProductType, error)
}

// 2. Product Category Repository
type ProductCategoryRepo interface {
    Create(pt ProductCategory) (ProductCategory, error)
    Update(id int, pt ProductCategory) (ProductCategory, error)
    Delete(id int) error
    First() (ProductCategory, error)
    Last() (ProductCategory, error)
    FindByID(id int) (ProductCategory, error)
    All() ([]ProductCategory, error)
}

// 3. Product Repository
type ProductRepo interface {
    Create(pt Product) (Product, error)
    Update(id int, pt Product) (Product, error)
    Delete(id int) error
    First() (Product, error)
    Last() (Product, error)
    FindByID(id int) (Product, error)
    All() ([]Product, error)
    Limit(max int) ([]Product, error)
    Offset(offset int) ([]Product, error)
    FindName(name string) ([]Product, error)
}


/* 
    1. PRODUCT TYPE REPOSITORY 
    implementation of product type repository based on ProductTypeRepo interface 
*/
type productTypeRepo struct {
    db *gorm.DB
}

func NewProductTypeRepo(db *gorm.DB) *productTypeRepo{
    return &productTypeRepo{db}
}

// Create PRODUCT TYPE 
func (r *productTypeRepo) Create(pt ProductType) (ProductType, error) {
    err := r.db.Create(&pt).Error
    return pt, err 
}

// Update PRODUCT TYPE 
func (r *productTypeRepo) Update(id int, pt ProductType) (ProductType, error) {
    err := r.db.Where("deleted_at IS NULL AND id=?", id).Updates(&pt).Error
    return pt, err
}

// Delete PRODUCT TYPE
func (r *productTypeRepo) Delete(id int) error {
    var pt ProductType
    return r.db.Where("deleted_at IS NULL AND id=?", id).Delete(&pt).Error
}

// First PRODUCT TYPE
func (r *productTypeRepo) First() (ProductType, error) {
    var pt ProductType
    err := r.db.First(&pt).Error

    return pt, err
}

// Last PRODUCT TYPE 
func (r *productTypeRepo) Last() (ProductType, error) {
    var pt ProductType
    err := r.db.Last(&pt).Error

    return pt, err
}

// Find By ID PRODUCT TYPE 
func (r *productTypeRepo) FindByID(id int) (ProductType, error) {
    var pt ProductType
    err := r.db.Where("deleted_at IS NULL AND id=?", id).First(&pt).Error

    return pt, err
}

// Get all PRODUCT TYPE
func (r *productTypeRepo) All() ([]ProductType, error) {
    var prodTypes []ProductType
    err := r.db.Where("deleted_at IS NULL").Find(&prodTypes).Error

    return prodTypes, err
}


/* 
    2. PRODUCT CATEGORY REPOSITORY 
    implementation of product category repository based on ProductCategoryRepo interface 
*/
type productCategoryRepo struct {
    db *gorm.DB
}

func NewProductCategoryRepo(db *gorm.DB) *productCategoryRepo {
    return &productCategoryRepo{db}
}

// Create 'product category'
func (r * productCategoryRepo) Create(pc ProductCategory) (ProductCategory, error) {
    err := r.db.Create(&pc).Error
    return pc, err
}

// Update 'product category'
func (r *productCategoryRepo) Update(id int, pc ProductCategory) (ProductCategory, error) {
    err := r.db.Where("deleted_at IS NULL AND id=?", id).Updates(&pc).Error
    return pc, err
}

// Delete 'product category'
func (r *productCategoryRepo) Delete(id int) error {
    var pc ProductCategory
    err := r.db.Where("deleted_at IS NULL AND id=?", id).Delete(&pc).Error

    return err
}

// First 'product category'
func (r *productCategoryRepo) First() (ProductCategory, error) {
    var pc ProductCategory
    err := r.db.Where("deleted_at IS NULL").First(&pc).Error
    return pc, err
}

// Last 'Product category'
func (r *productCategoryRepo) Last() (ProductCategory, error) {
    var pc ProductCategory
    err := r.db.Where("deleted_at IS NULL").Last(&pc).Error
    return pc, err
}

// 'product category' Find By ID 
func (r *productCategoryRepo) FindByID(id int) (ProductCategory, error) {
    var pc ProductCategory
    err := r.db.Where("deleted_at IS NULL AND id=?", id).First(&pc).Error
    return pc, err
}

// Get ALL 'product category'
func (r *productCategoryRepo) All() ([]ProductCategory, error) {
    var pcs []ProductCategory
    err := r.db.Where("deleted_at IS NULL").Find(&pcs).Error
    return pcs, err
}

/* 
    3. PRODUCT REPOSITORY 
*/
type productRepo struct {
    db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepo {
    return &productRepo{db}
}

// Create 'product'
func (r * productRepo) Create(p Product) (Product, error) {
    err := r.db.Create(&p).Error
    return p, err
}

// Update 'product'
func (r *productRepo) Update(id int, p Product) (Product, error) {
    err := r.db.Where("deleted_at IS NULL AND id=?", id).Updates(&p).Error
    return p, err
}

// Delete 'product'
func (r *productRepo) Delete(id int) error {
    var p Product
    err := r.db.Where("deleted_at IS NULL AND id=?", id).Delete(&p).Error

    return err
}

// First 'product'
func (r *productRepo) First() (Product, error) {
    var p Product
    err := r.db.Where("deleted_at IS NULL").First(&p).Error
    return p, err
}

// Last 'product'
func (r *productRepo) Last() (Product, error) {
    var p Product
    err := r.db.Where("deleted_at IS NULL").Last(&p).Error
    return p, err
}

// 'product' Find By ID 
func (r *productRepo) FindByID(id int) (Product, error) {
    var p Product
    err := r.db.Where("deleted_at IS NULL AND id=?", id).First(&p).Error
    return p, err
}

// Get ALL 'product'
func (r *productRepo) All() ([]Product, error) {
    var ps []Product
    err := r.db.Where("deleted_at IS NULL").Find(&ps).Error
    return ps, err
}

// Get Limit 'product'
func (r *productRepo) Limit(max int) ([]Product, error) {
    var ps []Product
    err := r.db.Where("deleted_at IS NULL").Limit(max).Find(&ps).Error
    return ps, err
}

// Get offset 'product'
func (r *productRepo) Offset(offset int) ([]Product, error) {
    var ps []Product
    err := r.db.Where("deleted_at IS NULL").Offset(offset).Find(&ps).Error
    return ps, err
}

// 'product' Find Name 
func (r *productRepo) FindName(name string) ([]Product, error) {
    var p []Product
    err := r.db.Where("deleted_at IS NULL AND name like ?", "%"+name+"%").Find(&p).Error
    return p, err
}
