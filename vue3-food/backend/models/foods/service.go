package foods

import "fmt"

/*
   PRODUCT TYPE SERVICE
*/
type ProductTypeService interface {
    Create(pt ProdTypeBase) (ProductType, error)
    Update(id int, pt ProdTypeBase) (ProductType, error)
    Delete(id int) error
    First() (ProductType, error)
    Last() (ProductType, error)
    FindByID(id int) (ProductType, error)
    All() ([]ProductType, error)
}

type productTypeService struct {
    repo ProductTypeRepo
}

func NewProductTypeService(repo ProductTypeRepo) *productTypeService {
    return &productTypeService{repo}
}

func (s *productTypeService) Create(ptr ProdTypeBase) (ProductType, error) {
    pt := ProductType{
        ProdType: ptr,
    }
    return s.repo.Create(pt)
}

func (s *productTypeService) Update(id int, ptr ProdTypeBase) (ProductType, error) {
    return s.repo.Update(id, ProductType{ProdType: ptr})
}

func (s *productTypeService) Delete(id int) (error) {
    return s.repo.Delete(id)
}

func (s *productTypeService) First() (ProductType, error) {
    return s.repo.First()
}

func (s *productTypeService) Last() (ProductType, error) {
    return s.repo.Last()
}

func (s *productTypeService) FindByID(id int) (ProductType, error) {
    return s.repo.FindByID(id)
}

func (s *productTypeService) All() ([]ProductType, error) {
    return s.repo.All()
}


/*
    PRODUCT CATEGORY SERVICE 
*/
type ProductCategoryService interface {
    Create(pcr ProdCatBase) (ProductCategory, error)
    Update(id int, pcr ProdCatBase) (ProductCategory, error)
    Delete(id int) error
    First() (ProductCategory, error)
    Last() (ProductCategory, error)
    FindByID(id int) (ProductCategory, error)
    All() ([]ProductCategory, error)
}

type productCategoryService struct {
    repo ProductCategoryRepo
}

func NewProductCategoryService(repo ProductCategoryRepo) *productCategoryService {
    return &productCategoryService{repo}
}

func (s *productCategoryService) Create(pcr ProdCatBase) (ProductCategory, error) {
    pc := ProductCategory{
        ProdCat: pcr,
    }
    return s.repo.Create(pc)
}

func (s *productCategoryService) Update(id int, pcr ProdCatBase) (ProductCategory, error) {
    return s.repo.Update(id, ProductCategory{ProdCat: pcr})
}

func (s *productCategoryService) Delete(id int) (error) {
    return s.repo.Delete(id)
}

func (s *productCategoryService) First() (ProductCategory, error) {
    return s.repo.First()
}

func (s *productCategoryService) Last() (ProductCategory, error) {
    return s.repo.Last()
}

func (s *productCategoryService) FindByID(id int) (ProductCategory, error) {
    return s.repo.FindByID(id)
}

func (s *productCategoryService) All() ([]ProductCategory, error) {
    return s.repo.All()
}


/*
    PRODUCT SERVICE 
*/
type ProductService interface {
    Create(pr ProductBase) (Product, error)
    Update(id int, pr ProductBase) (Product, error)
    Delete(id int) error
    First() (Product, error)
    Last() (Product, error)
    FindByID(id int) (Product, error)
    All() ([]Product, error)
    Limit(max int) ([]Product, error)
    Offset(offset int) ([]Product, error)
    FindName(name string) ([]Product, error)
}

type productService struct {
    repo ProductRepo
}

func NewProductService(repo ProductRepo) *productService {
    return &productService{repo}
}

// check is there an empty field
func fieldEmpty(field ...interface{}) bool {
    for _,k := range field {
        if k == nil {
            return true
        }
    }
    return false
}

func (s *productService) Create(pr ProductBase) (Product, error) {
    return s.repo.Create(Product{Prod:pr})
}

func (s *productService) Update(id int, pr ProductBase) (Product, error) {
    fmt.Println("Service product update 1:\n", pr)
    prod := Product{Prod:pr}
    fmt.Println("Service product update 2:\n", prod)
    prod.ID = uint(id)
    return s.repo.Update(id, prod)
}

func (s *productService) Delete(id int) (error) {
    return s.repo.Delete(id)
}

func (s *productService) First() (Product, error) {
    return s.repo.First()
}

func (s *productService) Last() (Product, error) {
    return s.repo.Last()
}

func (s *productService) FindByID(id int) (Product, error) {
    return s.repo.FindByID(id)
}

func (s *productService) All() ([]Product, error) {
    return s.repo.All()
}

func (s *productService) Limit(max int) ([]Product, error) {
    return s.repo.Limit(max)
}

func (s *productService) Offset(offset int) ([]Product, error) {
    return s.repo.Offset(offset)
}

func (s *productService) FindName(name string) ([]Product, error) {
    return s.repo.FindName(name)
}
