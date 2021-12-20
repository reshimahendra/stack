export const utils = {
    format(n, currency) {
        return new Intl.NumberFormat('id', {
            style: 'currency',
            currency: currency
        }).format(n)
    }
}  
