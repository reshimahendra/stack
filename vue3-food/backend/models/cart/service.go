package cart

type CartService interface {
    Create (c Cart) (Cart, error)
    Update (id int, c Cart) (Cart, error)
    Delete (id int) error
    FindByID (id int) (Cart, error)
    All() ([]Cart, error)
}

type CartItemService interface {
    Create (od CartItem) (CartItem, error)
    Update (id int, c CartItem) (CartItem, error)
    Delete (id int) error
    FindByID (id int) (CartItem, error)
    All() ([]CartItem, error)
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
type cartItemService struct {
    repo CartItemRepo
}

func NewCartItemService(repo CartItemRepo) *cartItemService {
    return &cartItemService{repo}
}

func (cis *cartItemService) Create(ol CartItem) (CartItem, error) {
    return cis.repo.Create(ol)
}

func (cis *cartItemService) Update(id int, ol CartItem) (CartItem, error) {
    return cis.repo.Update(id, ol)
}

func (cis *cartItemService) FindByID(id int) (CartItem, error) {
    return cis.repo.FindByID(id)
}

func (cis *cartItemService) All() ([]CartItem, error) {
    return cis.repo.All()
}
