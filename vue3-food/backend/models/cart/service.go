package cart

type CartService interface {
    Create (c Cart) (Cart, error)
    Update (id int, c Cart) (Cart, error)
    Delete (id int) error
    FindByID (id int) (Cart, error)
    All() ([]Cart, error)
}

type OrderListService interface {
    Create (od OrderList) (OrderList, error)
    Update (id int, c OrderList) (OrderList, error)
    Delete (id int) error
    FindByID (id int) (OrderList, error)
    All() ([]OrderList, error)
}

// Cart Service 
type cartService struct {
    repo CartRepo
}

func NewCartService(repo CartRepo) *cartService{
    return &cartService{repo}
}

func (cs *cartService) Create(c Cart) (Cart, error) {
    return cs.repo.Create(c)
}

func (cs *cartService) Update(id int, c Cart) (Cart, error) {
    return cs.repo.Update(id, c)
}

func (cs *cartService) Delete(id int) error {
    return cs.repo.Delete(id)
}

func (cs *cartService) FindByID(id int) (Cart, error) {
    return cs.repo.FindByID(id)
}

func (cs *cartService) All() ([]Cart, error) {
    return cs.repo.All()
}

// Order List Service
type orderListService struct {
    repo OrderListRepo
}

func NewOrderListService(repo OrderListRepo) *orderListService {
    return &orderListService{repo}
}

func (ols *orderListService) Create(ol OrderList) (OrderList, error) {
    return ols.repo.Create(ol)
}

func (ols *orderListService) Update(id int, ol OrderList) (OrderList, error) {
    return ols.repo.Update(id, ol)
}

func (ols *orderListService) FindByID(id int) (OrderList, error) {
    return ols.repo.FindByID(id)
}

func (ols *orderListService) All() ([]OrderList, error) {
    return ols.repo.All()
}
