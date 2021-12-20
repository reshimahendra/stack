package foods

/*
    PRODUCT TYPE SERVICE 
*/
type ProductTypeService interface {
    Create(pt ProductTypeRequest) (ProductType, error)
    Update(id int, pt ProductTypeRequest) (ProductType, error)
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

func (s *productTypeService) Create(ptr ProductTypeRequest) (ProductType, error) {
    pt := ProductType{
        Code:        ptr.Code,
        Name:        ptr.Name,
        Description: ptr.Description,
    }
    return s.repo.Create(pt)
}

func (s *productTypeService) Update(id int, ptr ProductTypeRequest) (ProductType, error) {
    pt := ProductType{
        Code:        ptr.Code,
        Name:        ptr.Name,
        Description: ptr.Description,
    }
    return s.repo.Update(id, pt)
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
    Create(pcr ProductCategoryRequest) (ProductCategory, error)
    Update(id int, pcr ProductCategoryRequest) (ProductCategory, error)
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

func (s *productCategoryService) Create(pcr ProductCategoryRequest) (ProductCategory, error) {
    pc := ProductCategory{
        Code:        pcr.Code,
        Name:        pcr.Name,
        Description: pcr.Description,
    }
    return s.repo.Create(pc)
}

func (s *productCategoryService) Update(id int, pcr ProductCategoryRequest) (ProductCategory, error) {
    pc := ProductCategory{
        Code:        pcr.Code,
        Name:        pcr.Name,
        Description: pcr.Description,
    }
    return s.repo.Update(id, pc)
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
    Create(pr ProductRequest) (Product, error)
    Update(id int, pr ProductReqUpdate) (Product, error)
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

func (s *productService) Create(pr ProductRequest) (Product, error) {
    p := Product{
        Sku: pr.Sku,
        Name: pr.Name,
        CategoryID: pr.CategoryID,
        TypeID: pr.TypeID,
        Price: pr.Price,
        IsReady: pr.IsReady,
        IsAvailable: pr.IsAvailable,
        Picture: pr.Picture,
    }
    return s.repo.Create(p)
}

func (s *productService) Update(id int, pr ProductReqUpdate) (Product, error) {
    p := Product{
        Sku: pr.Sku,
        Name: pr.Name,
        CategoryID: pr.CategoryID,
        TypeID: pr.TypeID,
        Price: pr.Price,
        IsReady: pr.IsReady,
        IsAvailable: pr.IsAvailable,
        Picture: pr.Picture,
    }
    return s.repo.Update(id, p)
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
