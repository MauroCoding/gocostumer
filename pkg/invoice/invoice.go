package invoice

type Invoice struct {
	country, city string
	total         float64
	client        customer.Costumer
}
